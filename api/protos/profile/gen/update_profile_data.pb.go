// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: update_profile_data.proto

package profile_grpc

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

type UpdateProfileDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Gender   string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *UpdateProfileDataRequest) Reset() {
	*x = UpdateProfileDataRequest{}
	mi := &file_update_profile_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileDataRequest) ProtoMessage() {}

func (x *UpdateProfileDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_update_profile_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileDataRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileDataRequest) Descriptor() ([]byte, []int) {
	return file_update_profile_data_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateProfileDataRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateProfileDataRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateProfileDataRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateProfileDataRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

var File_update_profile_data_proto protoreflect.FileDescriptor

var file_update_profile_data_proto_rawDesc = []byte{
	0x0a, 0x19, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x7d, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x42, 0x4b, 0x5a, 0x49, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x61, 0x72,
	0x6b, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f, 0x32, 0x30, 0x32, 0x34, 0x5f, 0x32,
	0x5f, 0x6b, 0x6f, 0x74, 0x79, 0x61, 0x72, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_update_profile_data_proto_rawDescOnce sync.Once
	file_update_profile_data_proto_rawDescData = file_update_profile_data_proto_rawDesc
)

func file_update_profile_data_proto_rawDescGZIP() []byte {
	file_update_profile_data_proto_rawDescOnce.Do(func() {
		file_update_profile_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_update_profile_data_proto_rawDescData)
	})
	return file_update_profile_data_proto_rawDescData
}

var file_update_profile_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_update_profile_data_proto_goTypes = []any{
	(*UpdateProfileDataRequest)(nil), // 0: profile.UpdateProfileDataRequest
}
var file_update_profile_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_update_profile_data_proto_init() }
func file_update_profile_data_proto_init() {
	if File_update_profile_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_update_profile_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_update_profile_data_proto_goTypes,
		DependencyIndexes: file_update_profile_data_proto_depIdxs,
		MessageInfos:      file_update_profile_data_proto_msgTypes,
	}.Build()
	File_update_profile_data_proto = out.File
	file_update_profile_data_proto_rawDesc = nil
	file_update_profile_data_proto_goTypes = nil
	file_update_profile_data_proto_depIdxs = nil
}
