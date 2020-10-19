// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: proto/def/contract.proto

package contract

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// BackoffPolicyType is the type for backoff policies
type BackoffPolicy int32

const (
	// Exponential backoff policy
	BackoffPolicy_Exponential BackoffPolicy = 0
	// Linear backoff policy
	BackoffPolicy_Linear BackoffPolicy = 1
)

// Enum value maps for BackoffPolicy.
var (
	BackoffPolicy_name = map[int32]string{
		0: "Exponential",
		1: "Linear",
	}
	BackoffPolicy_value = map[string]int32{
		"Exponential": 0,
		"Linear":      1,
	}
)

func (x BackoffPolicy) Enum() *BackoffPolicy {
	p := new(BackoffPolicy)
	*p = x
	return p
}

func (x BackoffPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackoffPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_def_contract_proto_enumTypes[0].Descriptor()
}

func (BackoffPolicy) Type() protoreflect.EnumType {
	return &file_proto_def_contract_proto_enumTypes[0]
}

func (x BackoffPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackoffPolicy.Descriptor instead.
func (BackoffPolicy) EnumDescriptor() ([]byte, []int) {
	return file_proto_def_contract_proto_rawDescGZIP(), []int{0}
}

// CloudEvent content mode
type ContentMode int32

const (
	ContentMode_BINARY     ContentMode = 0
	ContentMode_STRUCTURED ContentMode = 1
)

// Enum value maps for ContentMode.
var (
	ContentMode_name = map[int32]string{
		0: "BINARY",
		1: "STRUCTURED",
	}
	ContentMode_value = map[string]int32{
		"BINARY":     0,
		"STRUCTURED": 1,
	}
)

func (x ContentMode) Enum() *ContentMode {
	p := new(ContentMode)
	*p = x
	return p
}

func (x ContentMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentMode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_def_contract_proto_enumTypes[1].Descriptor()
}

func (ContentMode) Type() protoreflect.EnumType {
	return &file_proto_def_contract_proto_enumTypes[1]
}

func (x ContentMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentMode.Descriptor instead.
func (ContentMode) EnumDescriptor() ([]byte, []int) {
	return file_proto_def_contract_proto_rawDescGZIP(), []int{1}
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// attributes filters events by exact match on event context attributes.
	// Each key in the map is compared with the equivalent key in the event
	// context. An event passes the filter if all values are equal to the
	// specified values.
	//
	// Nested context attributes are not supported as keys. Only string values are supported.
	Attributes map[string]string `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_proto_def_contract_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type EgressConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dead letter is where the event is sent when something goes wrong
	DeadLetter string `protobuf:"bytes,1,opt,name=deadLetter,proto3" json:"deadLetter,omitempty"`
	// retry is the minimum number of retries the sender should attempt when
	// sending an event before moving it to the dead letter sink.
	//
	// Setting retry to 0 means don't retry.
	Retry uint32 `protobuf:"varint,2,opt,name=retry,proto3" json:"retry,omitempty"`
	// backoffPolicy is the retry backoff policy (linear, exponential).
	BackoffPolicy BackoffPolicy `protobuf:"varint,3,opt,name=backoffPolicy,proto3,enum=BackoffPolicy" json:"backoffPolicy,omitempty"`
	// backoffDelay is the delay before retrying in milliseconds.
	BackoffDelay uint64 `protobuf:"varint,4,opt,name=backoffDelay,proto3" json:"backoffDelay,omitempty"`
}

func (x *EgressConfig) Reset() {
	*x = EgressConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EgressConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EgressConfig) ProtoMessage() {}

func (x *EgressConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EgressConfig.ProtoReflect.Descriptor instead.
func (*EgressConfig) Descriptor() ([]byte, []int) {
	return file_proto_def_contract_proto_rawDescGZIP(), []int{1}
}

func (x *EgressConfig) GetDeadLetter() string {
	if x != nil {
		return x.DeadLetter
	}
	return ""
}

func (x *EgressConfig) GetRetry() uint32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

func (x *EgressConfig) GetBackoffPolicy() BackoffPolicy {
	if x != nil {
		return x.BackoffPolicy
	}
	return BackoffPolicy_Exponential
}

func (x *EgressConfig) GetBackoffDelay() uint64 {
	if x != nil {
		return x.BackoffDelay
	}
	return 0
}

type Egress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// consumer group name
	ConsumerGroup string `protobuf:"bytes,1,opt,name=consumerGroup,proto3" json:"consumerGroup,omitempty"`
	// destination is the sink where events are sent.
	Destination string `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Types that are assignable to ReplyStrategy:
	//	*Egress_ReplyUrl
	//	*Egress_ReplyToOriginalTopic
	ReplyStrategy isEgress_ReplyStrategy `protobuf_oneof:"replyStrategy"`
	Filter        *Filter                `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *Egress) Reset() {
	*x = Egress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Egress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Egress) ProtoMessage() {}

func (x *Egress) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Egress.ProtoReflect.Descriptor instead.
func (*Egress) Descriptor() ([]byte, []int) {
	return file_proto_def_contract_proto_rawDescGZIP(), []int{2}
}

func (x *Egress) GetConsumerGroup() string {
	if x != nil {
		return x.ConsumerGroup
	}
	return ""
}

func (x *Egress) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

func (m *Egress) GetReplyStrategy() isEgress_ReplyStrategy {
	if m != nil {
		return m.ReplyStrategy
	}
	return nil
}

func (x *Egress) GetReplyUrl() string {
	if x, ok := x.GetReplyStrategy().(*Egress_ReplyUrl); ok {
		return x.ReplyUrl
	}
	return ""
}

func (x *Egress) GetReplyToOriginalTopic() *empty.Empty {
	if x, ok := x.GetReplyStrategy().(*Egress_ReplyToOriginalTopic); ok {
		return x.ReplyToOriginalTopic
	}
	return nil
}

func (x *Egress) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type isEgress_ReplyStrategy interface {
	isEgress_ReplyStrategy()
}

type Egress_ReplyUrl struct {
	// Send the response to an url
	ReplyUrl string `protobuf:"bytes,3,opt,name=replyUrl,proto3,oneof"`
}

type Egress_ReplyToOriginalTopic struct {
	// Send the response to a Kafka topic
	ReplyToOriginalTopic *empty.Empty `protobuf:"bytes,4,opt,name=replyToOriginalTopic,proto3,oneof"`
}

func (*Egress_ReplyUrl) isEgress_ReplyStrategy() {}

func (*Egress_ReplyToOriginalTopic) isEgress_ReplyStrategy() {}

type Ingress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional content mode to use when pushing messages to Kafka
	ContentMode ContentMode `protobuf:"varint,1,opt,name=contentMode,proto3,enum=ContentMode" json:"contentMode,omitempty"`
	// Ingress can both listen on a specific HTTP path
	// or listen to the / path but match the Host header
	//
	// Types that are assignable to IngressType:
	//	*Ingress_Path
	//	*Ingress_Host
	IngressType isIngress_IngressType `protobuf_oneof:"ingressType"`
}

func (x *Ingress) Reset() {
	*x = Ingress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ingress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ingress) ProtoMessage() {}

func (x *Ingress) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ingress.ProtoReflect.Descriptor instead.
func (*Ingress) Descriptor() ([]byte, []int) {
	return file_proto_def_contract_proto_rawDescGZIP(), []int{3}
}

func (x *Ingress) GetContentMode() ContentMode {
	if x != nil {
		return x.ContentMode
	}
	return ContentMode_BINARY
}

func (m *Ingress) GetIngressType() isIngress_IngressType {
	if m != nil {
		return m.IngressType
	}
	return nil
}

func (x *Ingress) GetPath() string {
	if x, ok := x.GetIngressType().(*Ingress_Path); ok {
		return x.Path
	}
	return ""
}

func (x *Ingress) GetHost() string {
	if x, ok := x.GetIngressType().(*Ingress_Host); ok {
		return x.Host
	}
	return ""
}

type isIngress_IngressType interface {
	isIngress_IngressType()
}

type Ingress_Path struct {
	// path to listen for incoming events.
	Path string `protobuf:"bytes,2,opt,name=path,proto3,oneof"`
}

type Ingress_Host struct {
	// host header to match
	Host string `protobuf:"bytes,3,opt,name=host,proto3,oneof"`
}

func (*Ingress_Path) isIngress_IngressType() {}

func (*Ingress_Host) isIngress_IngressType() {}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the resource
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Topics name
	// Note: If there is an ingress configured, then this field must have exactly 1 element otherwise,
	//  if the resource does just dispatch from Kafka, then this topic list can contain multiple elements
	Topics []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	// A comma separated list of host/port pairs to use for establishing the initial connection to the Kafka cluster.
	// Note: we're using a comma separated list simply because that's how java kafka client likes it.
	BootstrapServers string `protobuf:"bytes,3,opt,name=bootstrapServers,proto3" json:"bootstrapServers,omitempty"`
	// Optional ingress for this topic
	Ingress *Ingress `protobuf:"bytes,4,opt,name=ingress,proto3" json:"ingress,omitempty"`
	// Optional configuration of egress valid for the whole resource
	EgressConfig *EgressConfig `protobuf:"bytes,5,opt,name=egressConfig,proto3" json:"egressConfig,omitempty"`
	// Optional egresses for this topic
	Egresses []*Egress `protobuf:"bytes,6,rep,name=egresses,proto3" json:"egresses,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_proto_def_contract_proto_rawDescGZIP(), []int{4}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Resource) GetBootstrapServers() string {
	if x != nil {
		return x.BootstrapServers
	}
	return ""
}

func (x *Resource) GetIngress() *Ingress {
	if x != nil {
		return x.Ingress
	}
	return nil
}

func (x *Resource) GetEgressConfig() *EgressConfig {
	if x != nil {
		return x.EgressConfig
	}
	return nil
}

func (x *Resource) GetEgresses() []*Egress {
	if x != nil {
		return x.Egresses
	}
	return nil
}

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Count each contract update.
	// Make sure each data plane pod has the same contract generation number.
	Generation uint64      `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
	Resources  []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_def_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_proto_def_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_proto_def_contract_proto_rawDescGZIP(), []int{5}
}

func (x *Contract) GetGeneration() uint64 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *Contract) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_proto_def_contract_proto protoreflect.FileDescriptor

var file_proto_def_contract_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x12, 0x34, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x6f,
	0x66, 0x66, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66,
	0x66, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x6f,
	0x66, 0x66, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x62,
	0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x22, 0xee, 0x01, 0x0a, 0x06,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x4c, 0x0a, 0x14,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x48, 0x00, 0x52, 0x14, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1f, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x22, 0x74, 0x0a, 0x07,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xda, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x0c, 0x65, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x65, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22,
	0x53, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2a, 0x2c, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72,
	0x10, 0x01, 0x2a, 0x29, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x55, 0x52, 0x45, 0x44, 0x10, 0x01, 0x42, 0x5b, 0x0a,
	0x2a, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x11, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5a, 0x1a,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_def_contract_proto_rawDescOnce sync.Once
	file_proto_def_contract_proto_rawDescData = file_proto_def_contract_proto_rawDesc
)

func file_proto_def_contract_proto_rawDescGZIP() []byte {
	file_proto_def_contract_proto_rawDescOnce.Do(func() {
		file_proto_def_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_def_contract_proto_rawDescData)
	})
	return file_proto_def_contract_proto_rawDescData
}

var file_proto_def_contract_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_def_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_def_contract_proto_goTypes = []interface{}{
	(BackoffPolicy)(0),   // 0: BackoffPolicy
	(ContentMode)(0),     // 1: ContentMode
	(*Filter)(nil),       // 2: Filter
	(*EgressConfig)(nil), // 3: EgressConfig
	(*Egress)(nil),       // 4: Egress
	(*Ingress)(nil),      // 5: Ingress
	(*Resource)(nil),     // 6: Resource
	(*Contract)(nil),     // 7: Contract
	nil,                  // 8: Filter.AttributesEntry
	(*empty.Empty)(nil),  // 9: google.protobuf.Empty
}
var file_proto_def_contract_proto_depIdxs = []int32{
	8, // 0: Filter.attributes:type_name -> Filter.AttributesEntry
	0, // 1: EgressConfig.backoffPolicy:type_name -> BackoffPolicy
	9, // 2: Egress.replyToOriginalTopic:type_name -> google.protobuf.Empty
	2, // 3: Egress.filter:type_name -> Filter
	1, // 4: Ingress.contentMode:type_name -> ContentMode
	5, // 5: Resource.ingress:type_name -> Ingress
	3, // 6: Resource.egressConfig:type_name -> EgressConfig
	4, // 7: Resource.egresses:type_name -> Egress
	6, // 8: Contract.resources:type_name -> Resource
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_proto_def_contract_proto_init() }
func file_proto_def_contract_proto_init() {
	if File_proto_def_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_def_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_def_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EgressConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_def_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Egress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_def_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ingress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_def_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_def_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_def_contract_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Egress_ReplyUrl)(nil),
		(*Egress_ReplyToOriginalTopic)(nil),
	}
	file_proto_def_contract_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Ingress_Path)(nil),
		(*Ingress_Host)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_def_contract_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_def_contract_proto_goTypes,
		DependencyIndexes: file_proto_def_contract_proto_depIdxs,
		EnumInfos:         file_proto_def_contract_proto_enumTypes,
		MessageInfos:      file_proto_def_contract_proto_msgTypes,
	}.Build()
	File_proto_def_contract_proto = out.File
	file_proto_def_contract_proto_rawDesc = nil
	file_proto_def_contract_proto_goTypes = nil
	file_proto_def_contract_proto_depIdxs = nil
}
