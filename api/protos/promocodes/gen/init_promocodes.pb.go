// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: init_promocodes.proto

package promocodes

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

var File_init_promocodes_proto protoreflect.FileDescriptor

var file_init_promocodes_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f,
	0x64, 0x65, 0x73, 0x1a, 0x14, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x65, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x91, 0x02, 0x0a, 0x0a,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x60, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x43, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x49, 0x5a, 0x47, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x6d, 0x61,
	0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f, 0x32, 0x30, 0x32, 0x34, 0x5f, 0x32, 0x5f, 0x6b, 0x6f, 0x74,
	0x79, 0x61, 0x72, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_init_promocodes_proto_goTypes = []any{
	(*GetUserPromoCodesRequest)(nil),  // 0: promocodes.GetUserPromoCodesRequest
	(*GetPromoCodeRequest)(nil),       // 1: promocodes.GetPromoCodeRequest
	(*DeletePromoCodesRequest)(nil),   // 2: promocodes.DeletePromoCodesRequest
	(*GetUserPromoCodesResponse)(nil), // 3: promocodes.GetUserPromoCodesResponse
	(*GetPromoCodeResponse)(nil),      // 4: promocodes.GetPromoCodeResponse
	(*empty.Empty)(nil),               // 5: google.protobuf.Empty
}
var file_init_promocodes_proto_depIdxs = []int32{
	0, // 0: promocodes.PromoCodes.GetUserPromoCodes:input_type -> promocodes.GetUserPromoCodesRequest
	1, // 1: promocodes.PromoCodes.GetPromoCode:input_type -> promocodes.GetPromoCodeRequest
	2, // 2: promocodes.PromoCodes.DeletePromoCode:input_type -> promocodes.DeletePromoCodesRequest
	3, // 3: promocodes.PromoCodes.GetUserPromoCodes:output_type -> promocodes.GetUserPromoCodesResponse
	4, // 4: promocodes.PromoCodes.GetPromoCode:output_type -> promocodes.GetPromoCodeResponse
	5, // 5: promocodes.PromoCodes.DeletePromoCode:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_init_promocodes_proto_init() }
func file_init_promocodes_proto_init() {
	if File_init_promocodes_proto != nil {
		return
	}
	file_get_promocodes_proto_init()
	file_get_promocode_by_name_proto_init()
	file_delete_promocode_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_init_promocodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_init_promocodes_proto_goTypes,
		DependencyIndexes: file_init_promocodes_proto_depIdxs,
	}.Build()
	File_init_promocodes_proto = out.File
	file_init_promocodes_proto_rawDesc = nil
	file_init_promocodes_proto_goTypes = nil
	file_init_promocodes_proto_depIdxs = nil
}
