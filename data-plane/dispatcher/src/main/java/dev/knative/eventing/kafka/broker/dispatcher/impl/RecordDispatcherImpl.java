/*
 * Copyright © 2018 Knative Authors (knative-dev@googlegroups.com)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package dev.knative.eventing.kafka.broker.dispatcher.impl;

import dev.knative.eventing.kafka.broker.core.AsyncCloseable;
import dev.knative.eventing.kafka.broker.dispatcher.CloudEventSender;
import dev.knative.eventing.kafka.broker.dispatcher.Filter;
import dev.knative.eventing.kafka.broker.dispatcher.RecordDispatcher;
import dev.knative.eventing.kafka.broker.dispatcher.RecordDispatcherListener;
import dev.knative.eventing.kafka.broker.dispatcher.ResponseHandler;
import io.cloudevents.CloudEvent;
import io.vertx.core.Context;
import io.vertx.core.Future;
import io.vertx.core.Promise;
import io.vertx.core.Vertx;
import io.vertx.kafka.client.common.tracing.ConsumerTracer;
import io.vertx.kafka.client.consumer.KafkaConsumerRecord;
import java.util.Objects;
import java.util.function.Function;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import static dev.knative.eventing.kafka.broker.core.utils.Logging.keyValue;

/**
 * This class implements the core algorithm of the Dispatcher (see {@link
 * RecordDispatcherImpl#dispatch(KafkaConsumerRecord)}).
 *
 * @see RecordDispatcherImpl#dispatch(KafkaConsumerRecord)
 */
public class RecordDispatcherImpl implements RecordDispatcher {

  private static final Logger logger = LoggerFactory.getLogger(RecordDispatcherImpl.class);

  private final Filter filter;
  private final Function<KafkaConsumerRecord<Object, CloudEvent>, Future<Void>> subscriberSender;
  private final Function<KafkaConsumerRecord<Object, CloudEvent>, Future<Void>> dlsSender;
  private final RecordDispatcherListener recordDispatcherListener;
  private final AsyncCloseable closeable;
  private final ConsumerTracer consumerTracer;

  /**
   * All args constructor.
   *
   * @param filter                   event filter
   * @param subscriberSender         sender to trigger subscriber
   * @param deadLetterSinkSender     sender to dead letter sink
   * @param responseHandler          handler of the response from {@code subscriberSender}
   * @param recordDispatcherListener hook receiver {@link RecordDispatcherListener}. It allows to plug in custom offset
   * @param consumerTracer           consumer tracer
   */
  public RecordDispatcherImpl(
    final Filter filter,
    final CloudEventSender subscriberSender,
    final CloudEventSender deadLetterSinkSender,
    final ResponseHandler responseHandler,
    final RecordDispatcherListener recordDispatcherListener,
    final ConsumerTracer consumerTracer) {
    Objects.requireNonNull(filter, "provide filter");
    Objects.requireNonNull(subscriberSender, "provide subscriberSender");
    Objects.requireNonNull(deadLetterSinkSender, "provide deadLetterSinkSender");
    Objects.requireNonNull(recordDispatcherListener, "provide offsetStrategy");
    Objects.requireNonNull(responseHandler, "provide sinkResponseHandler");

    this.filter = filter;
    this.subscriberSender = composeSenderAndSinkHandler(subscriberSender, responseHandler, "subscriber");
    this.dlsSender = composeSenderAndSinkHandler(deadLetterSinkSender, responseHandler, "dead letter sink");
    this.recordDispatcherListener = recordDispatcherListener;
    this.closeable = AsyncCloseable.compose(responseHandler, deadLetterSinkSender, subscriberSender);
    this.consumerTracer = consumerTracer;
  }

  /**
   * Handle the given record and returns a future that completes when the dispatch is completed and the offset is committed
   *
   * @param record record to handle.
   */
  @Override
  public Future<Void> dispatch(KafkaConsumerRecord<Object, CloudEvent> record) {
    Promise<Void> promise = Promise.promise();

    /*
    That's pretty much what happens here:

    +------------+    +------------+     +-------------+
    |            |    |            |     |             |
    |  record    |    |  filter    |     |  subscriber |
    |  received  +--->+  matching  +---->+  success    +-->end
    |            |    |            |     |             |
    +-+-------+--+    +-----+------+     +-----+-------+
      |       |             |                  |
      |       v             |                  |             +-------------+
      |    +--+---------+   |           +------v-------+     |             |
      |    |            |   |           |  subscriber  +---->+  dls        |
      |    | filter not |   +---------->+  failure     |     |  success    +---->end
      |    | matching   |               +----+---------+     |             |
      |    |            |                    |               +-----+-------+
      |    +---+--------+              +-----v-------+             |
      |        |                       | dls failure |             v
      |        |                       +-------------+----------> end
      +->end<--+
     */

    onRecordReceived(record, promise);

    return promise.future();
  }

  private void onRecordReceived(final KafkaConsumerRecord<Object, CloudEvent> record, Promise<Void> finalProm) {
    logDebug("Handling record", record);

    // Trace record received event
    // This needs to be done manually in order to work properly with both ordered and unordered delivery.
    if (consumerTracer != null) {
      Context context = Vertx.currentContext();
      ConsumerTracer.StartedSpan span = consumerTracer.prepareMessageReceived(context, record.record());
      if (span != null) {
        finalProm.future()
          .onComplete(ar -> {
            if (ar.succeeded()) {
              span.finish(context);
            } else {
              span.fail(context, ar.cause());
            }
          });
      }
    }

    recordDispatcherListener.recordReceived(record)
      .onSuccess(v -> {
        // Execute filtering
        if (filter.test(record.value())) {
          onFilterMatching(record, finalProm);
        } else {
          onFilterNotMatching(record, finalProm);
        }
      })
      .onFailure(finalProm::fail); // This should really never happen
  }

  private void onFilterMatching(final KafkaConsumerRecord<Object, CloudEvent> record, final Promise<Void> finalProm) {
    logDebug("Record match filtering", record);
    subscriberSender.apply(record)
      .onSuccess(res -> onSubscriberSuccess(record, finalProm))
      .onFailure(ex -> onSubscriberFailure(record, finalProm));
  }

  private void onFilterNotMatching(final KafkaConsumerRecord<Object, CloudEvent> record,
                                   final Promise<Void> finalProm) {
    logDebug("Record doesn't match filtering", record);
    recordDispatcherListener.recordDiscarded(record)
      .onComplete(finalProm);
  }

  private void onSubscriberSuccess(final KafkaConsumerRecord<Object, CloudEvent> record,
                                   final Promise<Void> finalProm) {
    logDebug("Successfully sent event to subscriber", record);
    recordDispatcherListener.successfullySentToSubscriber(record)
      .onComplete(finalProm);
  }

  private void onSubscriberFailure(final KafkaConsumerRecord<Object, CloudEvent> record,
                                   final Promise<Void> finalProm) {
    dlsSender.apply(record)
      .onSuccess(v -> onDeadLetterSinkSuccess(record, finalProm))
      .onFailure(ex -> onDeadLetterSinkFailure(record, ex, finalProm));
  }

  private void onDeadLetterSinkSuccess(final KafkaConsumerRecord<Object, CloudEvent> record,
                                       final Promise<Void> finalProm) {
    logDebug("Successfully sent event to the dead letter sink", record);
    recordDispatcherListener.successfullySentToDeadLetterSink(record)
      .onComplete(finalProm);
  }


  private void onDeadLetterSinkFailure(final KafkaConsumerRecord<Object, CloudEvent> record, final Throwable exception,
                                       final Promise<Void> finalProm) {
    recordDispatcherListener.failedToSendToDeadLetterSink(record, exception)
      .onComplete(finalProm);
  }

  private static Function<KafkaConsumerRecord<Object, CloudEvent>, Future<Void>> composeSenderAndSinkHandler(
    CloudEventSender sender, ResponseHandler sinkHandler, String senderType) {
    return rec -> sender.send(rec.value())
      .onFailure(ex -> logError("Failed to send event to " + senderType, rec, ex))
      .compose(res ->
        sinkHandler.handle(res)
          .onFailure(ex -> logError("Failed to handle " + senderType + " response", rec, ex))
      );
  }

  private static void logError(
    final String msg,
    final KafkaConsumerRecord<Object, CloudEvent> record,
    final Throwable cause) {

    if (logger.isDebugEnabled()) {
      logger.error(msg + " {} {} {} {} {}",
        keyValue("topic", record.topic()),
        keyValue("partition", record.partition()),
        keyValue("headers", record.headers()),
        keyValue("offset", record.offset()),
        keyValue("event", record.value()),
        cause
      );
    } else {
      logger.error(msg + " {} {} {}",
        keyValue("topic", record.topic()),
        keyValue("partition", record.partition()),
        keyValue("offset", record.offset()),
        cause
      );
    }
  }

  private static void logDebug(
    final String msg,
    final KafkaConsumerRecord<Object, CloudEvent> record) {

    logger.debug(msg + " {} {} {} {} {} {}",
      keyValue("topic", record.topic()),
      keyValue("partition", record.partition()),
      keyValue("headers", record.headers()),
      keyValue("offset", record.offset()),
      keyValue("key", record.key()),
      keyValue("event", record.value())
    );
  }

  public Future<Void> close() {
    return this.closeable.close();
  }
}
