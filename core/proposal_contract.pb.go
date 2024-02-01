// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: protocol/core/contract/proposal_contract.proto

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

type ProposalApproveContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress  []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ProposalId    int64  `protobuf:"varint,2,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	IsAddApproval bool   `protobuf:"varint,3,opt,name=is_add_approval,json=isAddApproval,proto3" json:"is_add_approval,omitempty"` // add or remove approval
}

func (x *ProposalApproveContract) Reset() {
	*x = ProposalApproveContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_proposal_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalApproveContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalApproveContract) ProtoMessage() {}

func (x *ProposalApproveContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_proposal_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalApproveContract.ProtoReflect.Descriptor instead.
func (*ProposalApproveContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_proposal_contract_proto_rawDescGZIP(), []int{0}
}

func (x *ProposalApproveContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ProposalApproveContract) GetProposalId() int64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

func (x *ProposalApproveContract) GetIsAddApproval() bool {
	if x != nil {
		return x.IsAddApproval
	}
	return false
}

type ProposalCreateContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte          `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Parameters   map[int64]int64 `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ProposalCreateContract) Reset() {
	*x = ProposalCreateContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_proposal_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalCreateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalCreateContract) ProtoMessage() {}

func (x *ProposalCreateContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_proposal_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalCreateContract.ProtoReflect.Descriptor instead.
func (*ProposalCreateContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_proposal_contract_proto_rawDescGZIP(), []int{1}
}

func (x *ProposalCreateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ProposalCreateContract) GetParameters() map[int64]int64 {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type ProposalDeleteContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ProposalId   int64  `protobuf:"varint,2,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (x *ProposalDeleteContract) Reset() {
	*x = ProposalDeleteContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_core_contract_proposal_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalDeleteContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalDeleteContract) ProtoMessage() {}

func (x *ProposalDeleteContract) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_core_contract_proposal_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalDeleteContract.ProtoReflect.Descriptor instead.
func (*ProposalDeleteContract) Descriptor() ([]byte, []int) {
	return file_protocol_core_contract_proposal_contract_proto_rawDescGZIP(), []int{2}
}

func (x *ProposalDeleteContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ProposalDeleteContract) GetProposalId() int64 {
	if x != nil {
		return x.ProposalId
	}
	return 0
}

var File_protocol_core_contract_proposal_contract_proto protoreflect.FileDescriptor

var file_protocol_core_contract_proposal_contract_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f,
	0x69, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x22, 0xce, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x50, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x49, 0x64, 0x42, 0x3b, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x31,
	0x6d, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x72, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_core_contract_proposal_contract_proto_rawDescOnce sync.Once
	file_protocol_core_contract_proposal_contract_proto_rawDescData = file_protocol_core_contract_proposal_contract_proto_rawDesc
)

func file_protocol_core_contract_proposal_contract_proto_rawDescGZIP() []byte {
	file_protocol_core_contract_proposal_contract_proto_rawDescOnce.Do(func() {
		file_protocol_core_contract_proposal_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_core_contract_proposal_contract_proto_rawDescData)
	})
	return file_protocol_core_contract_proposal_contract_proto_rawDescData
}

var file_protocol_core_contract_proposal_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocol_core_contract_proposal_contract_proto_goTypes = []interface{}{
	(*ProposalApproveContract)(nil), // 0: protocol.ProposalApproveContract
	(*ProposalCreateContract)(nil),  // 1: protocol.ProposalCreateContract
	(*ProposalDeleteContract)(nil),  // 2: protocol.ProposalDeleteContract
	nil,                             // 3: protocol.ProposalCreateContract.ParametersEntry
}
var file_protocol_core_contract_proposal_contract_proto_depIdxs = []int32{
	3, // 0: protocol.ProposalCreateContract.parameters:type_name -> protocol.ProposalCreateContract.ParametersEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protocol_core_contract_proposal_contract_proto_init() }
func file_protocol_core_contract_proposal_contract_proto_init() {
	if File_protocol_core_contract_proposal_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_core_contract_proposal_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalApproveContract); i {
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
		file_protocol_core_contract_proposal_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalCreateContract); i {
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
		file_protocol_core_contract_proposal_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalDeleteContract); i {
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
			RawDescriptor: file_protocol_core_contract_proposal_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protocol_core_contract_proposal_contract_proto_goTypes,
		DependencyIndexes: file_protocol_core_contract_proposal_contract_proto_depIdxs,
		MessageInfos:      file_protocol_core_contract_proposal_contract_proto_msgTypes,
	}.Build()
	File_protocol_core_contract_proposal_contract_proto = out.File
	file_protocol_core_contract_proposal_contract_proto_rawDesc = nil
	file_protocol_core_contract_proposal_contract_proto_goTypes = nil
	file_protocol_core_contract_proposal_contract_proto_depIdxs = nil
}
