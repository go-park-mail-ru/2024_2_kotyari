// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: init_rating_updater.proto

package rating_updater

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_init_rating_updater_proto protoreflect.FileDescriptor

var file_init_rating_updater_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5c,
	0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12,
	0x4b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x23, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x4d, 0x5a, 0x4b,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x2d,
	0x72, 0x75, 0x2f, 0x32, 0x30, 0x32, 0x34, 0x5f, 0x32, 0x5f, 0x6b, 0x6f, 0x74, 0x79, 0x61, 0x72,
	0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_init_rating_updater_proto_goTypes = []any{
	(*UpdateRatingRequest)(nil), // 0: rating_updater.UpdateRatingRequest
	(*empty.Empty)(nil),         // 1: google.protobuf.Empty
}
var file_init_rating_updater_proto_depIdxs = []int32{
	0, // 0: rating_updater.RatingUpdater.UpdateRating:input_type -> rating_updater.UpdateRatingRequest
	1, // 1: rating_updater.RatingUpdater.UpdateRating:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_init_rating_updater_proto_init() }
func file_init_rating_updater_proto_init() {
	if File_init_rating_updater_proto != nil {
		return
	}
	file_rating_updater_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_init_rating_updater_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_init_rating_updater_proto_goTypes,
		DependencyIndexes: file_init_rating_updater_proto_depIdxs,
	}.Build()
	File_init_rating_updater_proto = out.File
	file_init_rating_updater_proto_rawDesc = nil
	file_init_rating_updater_proto_goTypes = nil
	file_init_rating_updater_proto_depIdxs = nil
}
