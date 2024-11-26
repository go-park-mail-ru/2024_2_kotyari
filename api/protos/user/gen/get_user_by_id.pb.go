// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: get_user_by_id.proto

package grpc_gen

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

type GetUserByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserByIdRequest) Reset() {
	*x = GetUserByIdRequest{}
	mi := &file_get_user_by_id_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRequest) ProtoMessage() {}

func (x *GetUserByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_user_by_id_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return file_get_user_by_id_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByIdRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_get_user_by_id_proto protoreflect.FileDescriptor

var file_get_user_by_id_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x41, 0x5a, 0x3f, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x72,
	0x75, 0x2f, 0x32, 0x30, 0x32, 0x34, 0x5f, 0x32, 0x5f, 0x6b, 0x6f, 0x74, 0x79, 0x61, 0x72, 0x69,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_user_by_id_proto_rawDescOnce sync.Once
	file_get_user_by_id_proto_rawDescData = file_get_user_by_id_proto_rawDesc
)

func file_get_user_by_id_proto_rawDescGZIP() []byte {
	file_get_user_by_id_proto_rawDescOnce.Do(func() {
		file_get_user_by_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_user_by_id_proto_rawDescData)
	})
	return file_get_user_by_id_proto_rawDescData
}

var file_get_user_by_id_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_get_user_by_id_proto_goTypes = []any{
	(*GetUserByIdRequest)(nil), // 0: user.GetUserByIdRequest
}
var file_get_user_by_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_get_user_by_id_proto_init() }
func file_get_user_by_id_proto_init() {
	if File_get_user_by_id_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_get_user_by_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_user_by_id_proto_goTypes,
		DependencyIndexes: file_get_user_by_id_proto_depIdxs,
		MessageInfos:      file_get_user_by_id_proto_msgTypes,
	}.Build()
	File_get_user_by_id_proto = out.File
	file_get_user_by_id_proto_rawDesc = nil
	file_get_user_by_id_proto_goTypes = nil
	file_get_user_by_id_proto_depIdxs = nil
}
