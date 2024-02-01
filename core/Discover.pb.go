// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: protocol/core/Discover.proto

package core

import (
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

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	NodeId  []byte `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_Discover_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_Discover_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_protocol_core_Discover_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Endpoint) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Endpoint) GetNodeId() []byte {
	if x != nil {
		return x.NodeId
	}
	return nil
}

type PingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        *Endpoint `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Version   int32     `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp int64     `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PingMessage) Reset() {
	*x = PingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_Discover_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessage) ProtoMessage() {}

func (x *PingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_Discover_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingMessage.ProtoReflect.Descriptor instead.
func (*PingMessage) Descriptor() ([]byte, []int) {
	return file_protocol_core_Discover_proto_rawDescGZIP(), []int{1}
}

func (x *PingMessage) GetFrom() *Endpoint {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *PingMessage) GetTo() *Endpoint {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *PingMessage) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PingMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type PongMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Echo      int32     `protobuf:"varint,2,opt,name=echo,proto3" json:"echo,omitempty"`
	Timestamp int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *PongMessage) Reset() {
	*x = PongMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_Discover_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongMessage) ProtoMessage() {}

func (x *PongMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_Discover_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongMessage.ProtoReflect.Descriptor instead.
func (*PongMessage) Descriptor() ([]byte, []int) {
	return file_protocol_core_Discover_proto_rawDescGZIP(), []int{2}
}

func (x *PongMessage) GetFrom() *Endpoint {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *PongMessage) GetEcho() int32 {
	if x != nil {
		return x.Echo
	}
	return 0
}

func (x *PongMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type FindNeighbours struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	TargetId  []byte    `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Timestamp int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *FindNeighbours) Reset() {
	*x = FindNeighbours{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_Discover_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNeighbours) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNeighbours) ProtoMessage() {}

func (x *FindNeighbours) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_Discover_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNeighbours.ProtoReflect.Descriptor instead.
func (*FindNeighbours) Descriptor() ([]byte, []int) {
	return file_protocol_core_Discover_proto_rawDescGZIP(), []int{3}
}

func (x *FindNeighbours) GetFrom() *Endpoint {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *FindNeighbours) GetTargetId() []byte {
	if x != nil {
		return x.TargetId
	}
	return nil
}

func (x *FindNeighbours) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Neighbours struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From       *Endpoint   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Neighbours []*Endpoint `protobuf:"bytes,2,rep,name=neighbours,proto3" json:"neighbours,omitempty"`
	Timestamp  int64       `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Neighbours) Reset() {
	*x = Neighbours{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_Discover_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Neighbours) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Neighbours) ProtoMessage() {}

func (x *Neighbours) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_Discover_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Neighbours.ProtoReflect.Descriptor instead.
func (*Neighbours) Descriptor() ([]byte, []int) {
	return file_protocol_core_Discover_proto_rawDescGZIP(), []int{4}
}

func (x *Neighbours) GetFrom() *Endpoint {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Neighbours) GetNeighbours() []*Endpoint {
	if x != nil {
		return x.Neighbours
	}
	return nil
}

func (x *Neighbours) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type BackupMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag     bool  `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Priority int32 `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *BackupMessage) Reset() {
	*x = BackupMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_Discover_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupMessage) ProtoMessage() {}

func (x *BackupMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_Discover_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupMessage.ProtoReflect.Descriptor instead.
func (*BackupMessage) Descriptor() ([]byte, []int) {
	return file_protocol_core_Discover_proto_rawDescGZIP(), []int{5}
}

func (x *BackupMessage) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

func (x *BackupMessage) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

var File_protocol_core_Discover_proto protoreflect.FileDescriptor

var file_protocol_core_Discover_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x50, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x50,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x67,
	0x0a, 0x0b, 0x50, 0x6f, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x72, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x86, 0x01, 0x0a, 0x0a,
	0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x32, 0x0a, 0x0a, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x75, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x6e, 0x65, 0x69, 0x67,
	0x68, 0x62, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x3f, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x3c, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x42, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x31, 0x6d, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_core_Discover_proto_rawDescOnce sync.Once
	file_protocol_core_Discover_proto_rawDescData = file_protocol_core_Discover_proto_rawDesc
)

func file_protocol_core_Discover_proto_rawDescGZIP() []byte {
	file_protocol_core_Discover_proto_rawDescOnce.Do(func() {
		file_protocol_core_Discover_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_core_Discover_proto_rawDescData)
	})
	return file_protocol_core_Discover_proto_rawDescData
}

var file_protocol_core_Discover_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protocol_core_Discover_proto_goTypes = []interface{}{
	(*Endpoint)(nil),       // 0: protocol.Endpoint
	(*PingMessage)(nil),    // 1: protocol.PingMessage
	(*PongMessage)(nil),    // 2: protocol.PongMessage
	(*FindNeighbours)(nil), // 3: protocol.FindNeighbours
	(*Neighbours)(nil),     // 4: protocol.Neighbours
	(*BackupMessage)(nil),  // 5: protocol.BackupMessage
}
var file_protocol_core_Discover_proto_depIdxs = []int32{
	0, // 0: protocol.PingMessage.from:type_name -> protocol.Endpoint
	0, // 1: protocol.PingMessage.to:type_name -> protocol.Endpoint
	0, // 2: protocol.PongMessage.from:type_name -> protocol.Endpoint
	0, // 3: protocol.FindNeighbours.from:type_name -> protocol.Endpoint
	0, // 4: protocol.Neighbours.from:type_name -> protocol.Endpoint
	0, // 5: protocol.Neighbours.neighbours:type_name -> protocol.Endpoint
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protocol_core_Discover_proto_init() }
func file_protocol_core_Discover_proto_init() {
	if File_protocol_core_Discover_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_core_Discover_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_protocol_core_Discover_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingMessage); i {
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
		file_protocol_core_Discover_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongMessage); i {
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
		file_protocol_core_Discover_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNeighbours); i {
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
		file_protocol_core_Discover_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Neighbours); i {
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
		file_protocol_core_Discover_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupMessage); i {
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
			RawDescriptor: file_protocol_core_Discover_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_core_Discover_proto_goTypes,
		DependencyIndexes: file_protocol_core_Discover_proto_depIdxs,
		MessageInfos:      file_protocol_core_Discover_proto_msgTypes,
	}.Build()
	File_protocol_core_Discover_proto = out.File
	file_protocol_core_Discover_proto_rawDesc = nil
	file_protocol_core_Discover_proto_goTypes = nil
	file_protocol_core_Discover_proto_depIdxs = nil
}
