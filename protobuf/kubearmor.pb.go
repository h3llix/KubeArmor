// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: kubearmor.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Health check
type NonceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce int32 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *NonceMessage) Reset() {
	*x = NonceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubearmor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonceMessage) ProtoMessage() {}

func (x *NonceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_kubearmor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonceMessage.ProtoReflect.Descriptor instead.
func (*NonceMessage) Descriptor() ([]byte, []int) {
	return file_kubearmor_proto_rawDescGZIP(), []int{0}
}

func (x *NonceMessage) GetNonce() int32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

// message struct
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp   int64  `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	UpdatedTime string `protobuf:"bytes,2,opt,name=UpdatedTime,proto3" json:"UpdatedTime,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=ClusterName,proto3" json:"ClusterName,omitempty"`
	HostName    string `protobuf:"bytes,4,opt,name=HostName,proto3" json:"HostName,omitempty"`
	HostIP      string `protobuf:"bytes,5,opt,name=HostIP,proto3" json:"HostIP,omitempty"`
	Type        string `protobuf:"bytes,6,opt,name=Type,proto3" json:"Type,omitempty"`
	Level       string `protobuf:"bytes,7,opt,name=Level,proto3" json:"Level,omitempty"`
	Message     string `protobuf:"bytes,8,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubearmor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_kubearmor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_kubearmor_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *Message) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Message) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *Message) GetHostIP() string {
	if x != nil {
		return x.HostIP
	}
	return ""
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// alert struct
type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp     int64  `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	UpdatedTime   string `protobuf:"bytes,2,opt,name=UpdatedTime,proto3" json:"UpdatedTime,omitempty"`
	ClusterName   string `protobuf:"bytes,3,opt,name=ClusterName,proto3" json:"ClusterName,omitempty"`
	HostName      string `protobuf:"bytes,4,opt,name=HostName,proto3" json:"HostName,omitempty"`
	NamespaceName string `protobuf:"bytes,5,opt,name=NamespaceName,proto3" json:"NamespaceName,omitempty"`
	PodName       string `protobuf:"bytes,6,opt,name=PodName,proto3" json:"PodName,omitempty"`
	ContainerID   string `protobuf:"bytes,7,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"`
	ContainerName string `protobuf:"bytes,8,opt,name=ContainerName,proto3" json:"ContainerName,omitempty"`
	HostPID       int32  `protobuf:"varint,9,opt,name=HostPID,proto3" json:"HostPID,omitempty"`
	PPID          int32  `protobuf:"varint,10,opt,name=PPID,proto3" json:"PPID,omitempty"`
	PID           int32  `protobuf:"varint,11,opt,name=PID,proto3" json:"PID,omitempty"`
	UID           int32  `protobuf:"varint,12,opt,name=UID,proto3" json:"UID,omitempty"`
	PolicyName    string `protobuf:"bytes,13,opt,name=PolicyName,proto3" json:"PolicyName,omitempty"`
	Severity      string `protobuf:"bytes,14,opt,name=Severity,proto3" json:"Severity,omitempty"`
	Tags          string `protobuf:"bytes,15,opt,name=Tags,proto3" json:"Tags,omitempty"`
	Message       string `protobuf:"bytes,16,opt,name=Message,proto3" json:"Message,omitempty"`
	Type          string `protobuf:"bytes,17,opt,name=Type,proto3" json:"Type,omitempty"`
	Source        string `protobuf:"bytes,18,opt,name=Source,proto3" json:"Source,omitempty"`
	Operation     string `protobuf:"bytes,19,opt,name=Operation,proto3" json:"Operation,omitempty"`
	Resource      string `protobuf:"bytes,20,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Data          string `protobuf:"bytes,21,opt,name=Data,proto3" json:"Data,omitempty"`
	Action        string `protobuf:"bytes,22,opt,name=Action,proto3" json:"Action,omitempty"`
	Result        string `protobuf:"bytes,23,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubearmor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_kubearmor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_kubearmor_proto_rawDescGZIP(), []int{2}
}

func (x *Alert) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Alert) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *Alert) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Alert) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *Alert) GetNamespaceName() string {
	if x != nil {
		return x.NamespaceName
	}
	return ""
}

func (x *Alert) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *Alert) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *Alert) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *Alert) GetHostPID() int32 {
	if x != nil {
		return x.HostPID
	}
	return 0
}

func (x *Alert) GetPPID() int32 {
	if x != nil {
		return x.PPID
	}
	return 0
}

func (x *Alert) GetPID() int32 {
	if x != nil {
		return x.PID
	}
	return 0
}

func (x *Alert) GetUID() int32 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *Alert) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *Alert) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *Alert) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *Alert) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Alert) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Alert) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Alert) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Alert) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Alert) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Alert) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Alert) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// log struct
type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp     int64  `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	UpdatedTime   string `protobuf:"bytes,2,opt,name=UpdatedTime,proto3" json:"UpdatedTime,omitempty"`
	ClusterName   string `protobuf:"bytes,3,opt,name=ClusterName,proto3" json:"ClusterName,omitempty"`
	HostName      string `protobuf:"bytes,4,opt,name=HostName,proto3" json:"HostName,omitempty"`
	NamespaceName string `protobuf:"bytes,5,opt,name=NamespaceName,proto3" json:"NamespaceName,omitempty"`
	PodName       string `protobuf:"bytes,6,opt,name=PodName,proto3" json:"PodName,omitempty"`
	ContainerID   string `protobuf:"bytes,7,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"`
	ContainerName string `protobuf:"bytes,8,opt,name=ContainerName,proto3" json:"ContainerName,omitempty"`
	HostPID       int32  `protobuf:"varint,9,opt,name=HostPID,proto3" json:"HostPID,omitempty"`
	PPID          int32  `protobuf:"varint,10,opt,name=PPID,proto3" json:"PPID,omitempty"`
	PID           int32  `protobuf:"varint,11,opt,name=PID,proto3" json:"PID,omitempty"`
	UID           int32  `protobuf:"varint,12,opt,name=UID,proto3" json:"UID,omitempty"`
	Type          string `protobuf:"bytes,13,opt,name=Type,proto3" json:"Type,omitempty"`
	Source        string `protobuf:"bytes,14,opt,name=Source,proto3" json:"Source,omitempty"`
	Operation     string `protobuf:"bytes,15,opt,name=Operation,proto3" json:"Operation,omitempty"`
	Resource      string `protobuf:"bytes,16,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Data          string `protobuf:"bytes,17,opt,name=Data,proto3" json:"Data,omitempty"`
	Result        string `protobuf:"bytes,18,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubearmor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_kubearmor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_kubearmor_proto_rawDescGZIP(), []int{3}
}

func (x *Log) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Log) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

func (x *Log) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Log) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *Log) GetNamespaceName() string {
	if x != nil {
		return x.NamespaceName
	}
	return ""
}

func (x *Log) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *Log) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *Log) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *Log) GetHostPID() int32 {
	if x != nil {
		return x.HostPID
	}
	return 0
}

func (x *Log) GetPPID() int32 {
	if x != nil {
		return x.PPID
	}
	return 0
}

func (x *Log) GetPID() int32 {
	if x != nil {
		return x.PID
	}
	return 0
}

func (x *Log) GetUID() int32 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *Log) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Log) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Log) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Log) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Log) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Log) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// request message
type RequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter string `protobuf:"bytes,1,opt,name=Filter,proto3" json:"Filter,omitempty"`
}

func (x *RequestMessage) Reset() {
	*x = RequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubearmor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMessage) ProtoMessage() {}

func (x *RequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_kubearmor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMessage.ProtoReflect.Descriptor instead.
func (*RequestMessage) Descriptor() ([]byte, []int) {
	return file_kubearmor_proto_rawDescGZIP(), []int{4}
}

func (x *RequestMessage) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// reply message
type ReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retval int32 `protobuf:"varint,1,opt,name=Retval,proto3" json:"Retval,omitempty"`
}

func (x *ReplyMessage) Reset() {
	*x = ReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kubearmor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMessage) ProtoMessage() {}

func (x *ReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_kubearmor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMessage.ProtoReflect.Descriptor instead.
func (*ReplyMessage) Descriptor() ([]byte, []int) {
	return file_kubearmor_proto_rawDescGZIP(), []int{5}
}

func (x *ReplyMessage) GetRetval() int32 {
	if x != nil {
		return x.Retval
	}
	return 0
}

var File_kubearmor_proto protoreflect.FileDescriptor

var file_kubearmor_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6b, 0x75, 0x62, 0x65, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x0c, 0x4e, 0x6f, 0x6e,
	0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22,
	0xe3, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x6f, 0x73,
	0x74, 0x49, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x48, 0x6f, 0x73, 0x74, 0x49,
	0x50, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xf3, 0x04, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x49, 0x44,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x50, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50,
	0x50, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x50, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x17, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xef, 0x03, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x6f, 0x73,
	0x74, 0x50, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x48, 0x6f, 0x73, 0x74,
	0x50, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x50, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x50, 0x50, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x28, 0x0a,
	0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x74, 0x76, 0x61, 0x6c, 0x32,
	0xef, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x14, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0b, 0x57, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x30, 0x01, 0x12, 0x32, 0x0a,
	0x09, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x30,
	0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x2f, 0x4b, 0x75, 0x62, 0x65, 0x41, 0x72,
	0x6d, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kubearmor_proto_rawDescOnce sync.Once
	file_kubearmor_proto_rawDescData = file_kubearmor_proto_rawDesc
)

func file_kubearmor_proto_rawDescGZIP() []byte {
	file_kubearmor_proto_rawDescOnce.Do(func() {
		file_kubearmor_proto_rawDescData = protoimpl.X.CompressGZIP(file_kubearmor_proto_rawDescData)
	})
	return file_kubearmor_proto_rawDescData
}

var file_kubearmor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kubearmor_proto_goTypes = []interface{}{
	(*NonceMessage)(nil),   // 0: feeder.NonceMessage
	(*Message)(nil),        // 1: feeder.Message
	(*Alert)(nil),          // 2: feeder.Alert
	(*Log)(nil),            // 3: feeder.Log
	(*RequestMessage)(nil), // 4: feeder.RequestMessage
	(*ReplyMessage)(nil),   // 5: feeder.ReplyMessage
}
var file_kubearmor_proto_depIdxs = []int32{
	0, // 0: feeder.LogService.HealthCheck:input_type -> feeder.NonceMessage
	4, // 1: feeder.LogService.WatchMessages:input_type -> feeder.RequestMessage
	4, // 2: feeder.LogService.WatchAlerts:input_type -> feeder.RequestMessage
	4, // 3: feeder.LogService.WatchLogs:input_type -> feeder.RequestMessage
	5, // 4: feeder.LogService.HealthCheck:output_type -> feeder.ReplyMessage
	1, // 5: feeder.LogService.WatchMessages:output_type -> feeder.Message
	2, // 6: feeder.LogService.WatchAlerts:output_type -> feeder.Alert
	3, // 7: feeder.LogService.WatchLogs:output_type -> feeder.Log
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kubearmor_proto_init() }
func file_kubearmor_proto_init() {
	if File_kubearmor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kubearmor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonceMessage); i {
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
		file_kubearmor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_kubearmor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_kubearmor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_kubearmor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMessage); i {
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
		file_kubearmor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMessage); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kubearmor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kubearmor_proto_goTypes,
		DependencyIndexes: file_kubearmor_proto_depIdxs,
		MessageInfos:      file_kubearmor_proto_msgTypes,
	}.Build()
	File_kubearmor_proto = out.File
	file_kubearmor_proto_rawDesc = nil
	file_kubearmor_proto_goTypes = nil
	file_kubearmor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogServiceClient interface {
	HealthCheck(ctx context.Context, in *NonceMessage, opts ...grpc.CallOption) (*ReplyMessage, error)
	WatchMessages(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchMessagesClient, error)
	WatchAlerts(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchAlertsClient, error)
	WatchLogs(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchLogsClient, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) HealthCheck(ctx context.Context, in *NonceMessage, opts ...grpc.CallOption) (*ReplyMessage, error) {
	out := new(ReplyMessage)
	err := c.cc.Invoke(ctx, "/feeder.LogService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) WatchMessages(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogService_serviceDesc.Streams[0], "/feeder.LogService/WatchMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceWatchMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_WatchMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type logServiceWatchMessagesClient struct {
	grpc.ClientStream
}

func (x *logServiceWatchMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logServiceClient) WatchAlerts(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchAlertsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogService_serviceDesc.Streams[1], "/feeder.LogService/WatchAlerts", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceWatchAlertsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_WatchAlertsClient interface {
	Recv() (*Alert, error)
	grpc.ClientStream
}

type logServiceWatchAlertsClient struct {
	grpc.ClientStream
}

func (x *logServiceWatchAlertsClient) Recv() (*Alert, error) {
	m := new(Alert)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logServiceClient) WatchLogs(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (LogService_WatchLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogService_serviceDesc.Streams[2], "/feeder.LogService/WatchLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceWatchLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogService_WatchLogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type logServiceWatchLogsClient struct {
	grpc.ClientStream
}

func (x *logServiceWatchLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServiceServer is the server API for LogService service.
type LogServiceServer interface {
	HealthCheck(context.Context, *NonceMessage) (*ReplyMessage, error)
	WatchMessages(*RequestMessage, LogService_WatchMessagesServer) error
	WatchAlerts(*RequestMessage, LogService_WatchAlertsServer) error
	WatchLogs(*RequestMessage, LogService_WatchLogsServer) error
}

// UnimplementedLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (*UnimplementedLogServiceServer) HealthCheck(context.Context, *NonceMessage) (*ReplyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (*UnimplementedLogServiceServer) WatchMessages(*RequestMessage, LogService_WatchMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchMessages not implemented")
}
func (*UnimplementedLogServiceServer) WatchAlerts(*RequestMessage, LogService_WatchAlertsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchAlerts not implemented")
}
func (*UnimplementedLogServiceServer) WatchLogs(*RequestMessage, LogService_WatchLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchLogs not implemented")
}

func RegisterLogServiceServer(s *grpc.Server, srv LogServiceServer) {
	s.RegisterService(&_LogService_serviceDesc, srv)
}

func _LogService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeder.LogService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).HealthCheck(ctx, req.(*NonceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_WatchMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).WatchMessages(m, &logServiceWatchMessagesServer{stream})
}

type LogService_WatchMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type logServiceWatchMessagesServer struct {
	grpc.ServerStream
}

func (x *logServiceWatchMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _LogService_WatchAlerts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).WatchAlerts(m, &logServiceWatchAlertsServer{stream})
}

type LogService_WatchAlertsServer interface {
	Send(*Alert) error
	grpc.ServerStream
}

type logServiceWatchAlertsServer struct {
	grpc.ServerStream
}

func (x *logServiceWatchAlertsServer) Send(m *Alert) error {
	return x.ServerStream.SendMsg(m)
}

func _LogService_WatchLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServiceServer).WatchLogs(m, &logServiceWatchLogsServer{stream})
}

type LogService_WatchLogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type logServiceWatchLogsServer struct {
	grpc.ServerStream
}

func (x *logServiceWatchLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

var _LogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feeder.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _LogService_HealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchMessages",
			Handler:       _LogService_WatchMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchAlerts",
			Handler:       _LogService_WatchAlerts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchLogs",
			Handler:       _LogService_WatchLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kubearmor.proto",
}
