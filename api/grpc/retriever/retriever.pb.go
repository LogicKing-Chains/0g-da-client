// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: retriever/retriever.proto

package retriever

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

type BlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash of the ReducedBatchHeader defined onchain, see:
	// https://github.com/0glabs/0g-da-client/blob/master/contracts/src/interfaces/IZGDAServiceManager.sol#L43
	// This identifies the batch that this blob belongs to.
	BatchHeaderHash []byte `protobuf:"bytes,1,opt,name=batch_header_hash,json=batchHeaderHash,proto3" json:"batch_header_hash,omitempty"`
	// Which blob in the batch this is requesting for (note: a batch is logically an
	// ordered list of blobs).
	BlobIndex uint32 `protobuf:"varint,2,opt,name=blob_index,json=blobIndex,proto3" json:"blob_index,omitempty"`
	// The Ethereum block number at which the batch for this blob was constructed.
	ReferenceBlockNumber uint32 `protobuf:"varint,3,opt,name=reference_block_number,json=referenceBlockNumber,proto3" json:"reference_block_number,omitempty"`
	// Which quorum of the blob this is requesting for (note a blob can participate in
	// multiple quorums).
	QuorumId uint32 `protobuf:"varint,4,opt,name=quorum_id,json=quorumId,proto3" json:"quorum_id,omitempty"`
}

func (x *BlobRequest) Reset() {
	*x = BlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retriever_retriever_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobRequest) ProtoMessage() {}

func (x *BlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_retriever_retriever_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobRequest.ProtoReflect.Descriptor instead.
func (*BlobRequest) Descriptor() ([]byte, []int) {
	return file_retriever_retriever_proto_rawDescGZIP(), []int{0}
}

func (x *BlobRequest) GetBatchHeaderHash() []byte {
	if x != nil {
		return x.BatchHeaderHash
	}
	return nil
}

func (x *BlobRequest) GetBlobIndex() uint32 {
	if x != nil {
		return x.BlobIndex
	}
	return 0
}

func (x *BlobRequest) GetReferenceBlockNumber() uint32 {
	if x != nil {
		return x.ReferenceBlockNumber
	}
	return 0
}

func (x *BlobRequest) GetQuorumId() uint32 {
	if x != nil {
		return x.QuorumId
	}
	return 0
}

type BlobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The blob retrieved and reconstructed from the ZGDA Nodes per BlobRequest.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlobReply) Reset() {
	*x = BlobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retriever_retriever_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobReply) ProtoMessage() {}

func (x *BlobReply) ProtoReflect() protoreflect.Message {
	mi := &file_retriever_retriever_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobReply.ProtoReflect.Descriptor instead.
func (*BlobReply) Descriptor() ([]byte, []int) {
	return file_retriever_retriever_proto_rawDescGZIP(), []int{1}
}

func (x *BlobReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_retriever_retriever_proto protoreflect.FileDescriptor

var file_retriever_retriever_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x22, 0xab, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x62, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x14, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x6f, 0x72, 0x75,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x6f, 0x72,
	0x75, 0x6d, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x4b, 0x0a, 0x09, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x42, 0x6c,
	0x6f, 0x62, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x2e, 0x42,
	0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x30, 0x67, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x30, 0x67, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_retriever_retriever_proto_rawDescOnce sync.Once
	file_retriever_retriever_proto_rawDescData = file_retriever_retriever_proto_rawDesc
)

func file_retriever_retriever_proto_rawDescGZIP() []byte {
	file_retriever_retriever_proto_rawDescOnce.Do(func() {
		file_retriever_retriever_proto_rawDescData = protoimpl.X.CompressGZIP(file_retriever_retriever_proto_rawDescData)
	})
	return file_retriever_retriever_proto_rawDescData
}

var file_retriever_retriever_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_retriever_retriever_proto_goTypes = []interface{}{
	(*BlobRequest)(nil), // 0: retriever.BlobRequest
	(*BlobReply)(nil),   // 1: retriever.BlobReply
}
var file_retriever_retriever_proto_depIdxs = []int32{
	0, // 0: retriever.Retriever.RetrieveBlob:input_type -> retriever.BlobRequest
	1, // 1: retriever.Retriever.RetrieveBlob:output_type -> retriever.BlobReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_retriever_retriever_proto_init() }
func file_retriever_retriever_proto_init() {
	if File_retriever_retriever_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_retriever_retriever_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobRequest); i {
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
		file_retriever_retriever_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobReply); i {
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
			RawDescriptor: file_retriever_retriever_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_retriever_retriever_proto_goTypes,
		DependencyIndexes: file_retriever_retriever_proto_depIdxs,
		MessageInfos:      file_retriever_retriever_proto_msgTypes,
	}.Build()
	File_retriever_retriever_proto = out.File
	file_retriever_retriever_proto_rawDesc = nil
	file_retriever_retriever_proto_goTypes = nil
	file_retriever_retriever_proto_depIdxs = nil
}
