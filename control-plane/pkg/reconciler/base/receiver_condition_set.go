/*
 * Copyright 2020 The Knative Authors
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

// receiver_condition_set.go contains Broker and Kafka Sink logic for status conditions handling.
package base

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/network"
	"knative.dev/pkg/reconciler"

	"knative.dev/eventing-kafka-broker/control-plane/pkg/config"
)

const (
	ConditionAddressable             apis.ConditionType = "Addressable"
	ConditionDataPlaneAvailable      apis.ConditionType = "DataPlaneAvailable"
	ConditionTopicReady              apis.ConditionType = "TopicReady"
	ConditionConfigMapUpdated        apis.ConditionType = "ConfigMapUpdated"
	ConditionConfigParsed            apis.ConditionType = "ConfigParsed"
	ConditionInitialOffsetsCommitted apis.ConditionType = "InitialOffsetsCommitted"
)

var IngressConditionSet = apis.NewLivingConditionSet(
	ConditionAddressable,
	ConditionDataPlaneAvailable,
	ConditionTopicReady,
	ConditionConfigMapUpdated,
	ConditionConfigParsed,
)

var EgressConditionSet = apis.NewLivingConditionSet(
	ConditionDataPlaneAvailable,
	ConditionTopicReady,
	ConditionConfigMapUpdated,
	ConditionInitialOffsetsCommitted,
)

const (
	TopicOwnerAnnotation = "eventing.knative.dev/topic.owner"

	ReasonDataPlaneNotAvailable  = "Data plane not available"
	MessageDataPlaneNotAvailable = "Did you install the data plane for this component?"

	ReasonTopicNotPresentOrInvalid = "Topic is not present or invalid"
)

type Object interface {
	duckv1.KRShaped
	runtime.Object
}

type StatusConditionManager struct {
	Object Object

	SetAddress func(u *apis.URL)

	Env              *config.Env
	BootstrapServers string

	Recorder record.EventRecorder
}

func (manager *StatusConditionManager) DataPlaneAvailable() {
	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkTrue(ConditionDataPlaneAvailable)
}

func (manager *StatusConditionManager) DataPlaneNotAvailable() reconciler.Event {

	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionDataPlaneAvailable,
		ReasonDataPlaneNotAvailable,
		MessageDataPlaneNotAvailable,
	)

	return fmt.Errorf("%s: %s", ReasonDataPlaneNotAvailable, MessageDataPlaneNotAvailable)
}

func (manager *StatusConditionManager) FailedToGetConfigMap(err error) reconciler.Event {

	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionConfigMapUpdated,
		fmt.Sprintf(
			"Failed to get ConfigMap: %s",
			manager.Env.DataPlaneConfigMapAsString(),
		),
		"%v",
		err,
	)

	return fmt.Errorf("failed to get contract config map %s: %w", manager.Env.DataPlaneConfigMapAsString(), err)
}

func (manager *StatusConditionManager) FailedToGetDataFromConfigMap(err error) reconciler.Event {

	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionConfigMapUpdated,
		fmt.Sprintf(
			"Failed to get contract data from ConfigMap: %s",
			manager.Env.DataPlaneConfigMapAsString(),
		),
		"%v",
		err,
	)

	return fmt.Errorf("failed to get broker and triggers data from config map %s: %w", manager.Env.DataPlaneConfigMapAsString(), err)
}

func (manager *StatusConditionManager) FailedToUpdateConfigMap(err error) reconciler.Event {

	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionConfigMapUpdated,
		fmt.Sprintf("Failed to update ConfigMap: %s", manager.Env.DataPlaneConfigMapAsString()),
		"%s",
		err,
	)

	return fmt.Errorf("failed to update contract config map %s: %w", manager.Env.DataPlaneConfigMapAsString(), err)
}

func (manager *StatusConditionManager) ConfigMapUpdated() {

	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkTrueWithReason(
		ConditionConfigMapUpdated,
		fmt.Sprintf("Config map %s updated", manager.Env.DataPlaneConfigMapAsString()),
		"",
	)
}

func (manager *StatusConditionManager) FailedToCreateTopic(topic string, err error) reconciler.Event {

	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionTopicReady,
		fmt.Sprintf("Failed to create topic: %s", topic),
		"%v",
		err,
	)

	return fmt.Errorf("failed to create topic: %s: %w", topic, err)
}

func (manager *StatusConditionManager) TopicReady(topic string) {

	if owner, ok := manager.Object.GetStatus().Annotations[TopicOwnerAnnotation]; ok {
		manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkTrueWithReason(
			ConditionTopicReady,
			fmt.Sprintf("Topic %s (owner %s)", topic, owner),
			"",
		)

		return
	}

	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkTrueWithReason(
		ConditionTopicReady,
		fmt.Sprintf("Topic %s created", topic),
		"",
	)
}

func (manager *StatusConditionManager) Reconciled() reconciler.Event {

	if manager.SetAddress != nil {
		object := manager.Object

		manager.SetAddress(&apis.URL{
			Scheme: "http",
			Host:   network.GetServiceHostname(manager.Env.IngressName, manager.Env.SystemNamespace),
			Path:   fmt.Sprintf("/%s/%s", object.GetNamespace(), object.GetName()),
		})
		object.GetConditionSet().Manage(object.GetStatus()).MarkTrue(ConditionAddressable)
	}

	return nil
}

func (manager *StatusConditionManager) FailedToUpdateDispatcherPodsAnnotation(err error) {

	// We don't set status conditions for dispatcher pods updates.

	// Record the event.
	manager.Recorder.Eventf(
		manager.Object,
		corev1.EventTypeWarning,
		"failed to update dispatcher pods annotation",
		"%v",
		err,
	)
}

func (manager *StatusConditionManager) FailedToUpdateReceiverPodsAnnotation(err error) reconciler.Event {

	return fmt.Errorf("failed to update receiver pods annotation: %w", err)
}

func (manager *StatusConditionManager) FailedToGetConfig(err error) reconciler.Event {

	return fmt.Errorf("failed to get contract configuration: %w", err)
}

func (manager *StatusConditionManager) FailedToResolveConfig(err error) reconciler.Event {

	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionConfigParsed,
		fmt.Sprintf("%v", err),
		"",
	)

	return fmt.Errorf("failed to get contract configuration: %w", err)
}

func (manager *StatusConditionManager) ConfigResolved() {
	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkTrue(ConditionConfigParsed)
}

func (manager *StatusConditionManager) TopicsNotPresentOrInvalidErr(topics []string, err error) error {
	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionTopicReady,
		ReasonTopicNotPresentOrInvalid,
		"topics %v: %s",
		topics,
		err.Error(),
	)

	return fmt.Errorf("topics %v not present or invalid: %w", topics, err)
}

func (manager *StatusConditionManager) TopicsNotPresentOrInvalid(topics []string) error {
	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionTopicReady,
		ReasonTopicNotPresentOrInvalid,
		"Check topics %v configuration",
		topics,
	)
	return fmt.Errorf("topics %v not present or invalid: check topic configuration", topics)
}

func (manager *StatusConditionManager) InitialOffsetNotCommitted(err error) error {
	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkFalse(
		ConditionInitialOffsetsCommitted,
		"InitialOffsetsNotCommitted",
		err.Error(),
	)
	return err
}

func (manager *StatusConditionManager) InitialOffsetsCommitted() {
	manager.Object.GetConditionSet().Manage(manager.Object.GetStatus()).MarkTrue(ConditionInitialOffsetsCommitted)
}
