// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: wishlist_model.proto

package grpc_gen

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type WishlistItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint32               `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	AddedAt   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=added_at,json=addedAt,proto3" json:"added_at,omitempty"`
}

func (x *WishlistItem) Reset() {
	*x = WishlistItem{}
	mi := &file_wishlist_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WishlistItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistItem) ProtoMessage() {}

func (x *WishlistItem) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishlistItem.ProtoReflect.Descriptor instead.
func (*WishlistItem) Descriptor() ([]byte, []int) {
	return file_wishlist_model_proto_rawDescGZIP(), []int{0}
}

func (x *WishlistItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *WishlistItem) GetAddedAt() *timestamp.Timestamp {
	if x != nil {
		return x.AddedAt
	}
	return nil
}

type Wishlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Link  string          `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Items []*WishlistItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Wishlist) Reset() {
	*x = Wishlist{}
	mi := &file_wishlist_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Wishlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wishlist) ProtoMessage() {}

func (x *Wishlist) ProtoReflect() protoreflect.Message {
	mi := &file_wishlist_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wishlist.ProtoReflect.Descriptor instead.
func (*Wishlist) Descriptor() ([]byte, []int) {
	return file_wishlist_model_proto_rawDescGZIP(), []int{1}
}

func (x *Wishlist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Wishlist) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Wishlist) GetItems() []*WishlistItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_wishlist_model_proto protoreflect.FileDescriptor

var file_wishlist_model_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x64, 0x0a, 0x0c, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x35, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x60, 0x0a, 0x08, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x2c, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x45, 0x5a, 0x43, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f,
	0x32, 0x30, 0x32, 0x34, 0x5f, 0x32, 0x5f, 0x6b, 0x6f, 0x74, 0x79, 0x61, 0x72, 0x69, 0x2f, 0x77,
	0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wishlist_model_proto_rawDescOnce sync.Once
	file_wishlist_model_proto_rawDescData = file_wishlist_model_proto_rawDesc
)

func file_wishlist_model_proto_rawDescGZIP() []byte {
	file_wishlist_model_proto_rawDescOnce.Do(func() {
		file_wishlist_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_wishlist_model_proto_rawDescData)
	})
	return file_wishlist_model_proto_rawDescData
}

var file_wishlist_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wishlist_model_proto_goTypes = []any{
	(*WishlistItem)(nil),        // 0: wishlist.WishlistItem
	(*Wishlist)(nil),            // 1: wishlist.Wishlist
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_wishlist_model_proto_depIdxs = []int32{
	2, // 0: wishlist.WishlistItem.added_at:type_name -> google.protobuf.Timestamp
	0, // 1: wishlist.Wishlist.items:type_name -> wishlist.WishlistItem
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_wishlist_model_proto_init() }
func file_wishlist_model_proto_init() {
	if File_wishlist_model_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wishlist_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wishlist_model_proto_goTypes,
		DependencyIndexes: file_wishlist_model_proto_depIdxs,
		MessageInfos:      file_wishlist_model_proto_msgTypes,
	}.Build()
	File_wishlist_model_proto = out.File
	file_wishlist_model_proto_rawDesc = nil
	file_wishlist_model_proto_goTypes = nil
	file_wishlist_model_proto_depIdxs = nil
}