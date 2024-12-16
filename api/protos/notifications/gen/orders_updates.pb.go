// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: orders_updates.proto

package notifications

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

type GetOrdersUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetOrdersUpdatesRequest) Reset() {
	*x = GetOrdersUpdatesRequest{}
	mi := &file_orders_updates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersUpdatesRequest) ProtoMessage() {}

func (x *GetOrdersUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orders_updates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersUpdatesRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_orders_updates_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrdersUpdatesRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type OrderUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId   string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	NewStatus string `protobuf:"bytes,2,opt,name=new_status,json=newStatus,proto3" json:"new_status,omitempty"`
}

func (x *OrderUpdateMessage) Reset() {
	*x = OrderUpdateMessage{}
	mi := &file_orders_updates_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdateMessage) ProtoMessage() {}

func (x *OrderUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_orders_updates_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdateMessage.ProtoReflect.Descriptor instead.
func (*OrderUpdateMessage) Descriptor() ([]byte, []int) {
	return file_orders_updates_proto_rawDescGZIP(), []int{1}
}

func (x *OrderUpdateMessage) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderUpdateMessage) GetNewStatus() string {
	if x != nil {
		return x.NewStatus
	}
	return ""
}

type GetOrdersUpdatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrdersUpdates []*OrderUpdateMessage `protobuf:"bytes,2,rep,name=orders_updates,json=ordersUpdates,proto3" json:"orders_updates,omitempty"`
}

func (x *GetOrdersUpdatesResponse) Reset() {
	*x = GetOrdersUpdatesResponse{}
	mi := &file_orders_updates_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersUpdatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersUpdatesResponse) ProtoMessage() {}

func (x *GetOrdersUpdatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orders_updates_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersUpdatesResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersUpdatesResponse) Descriptor() ([]byte, []int) {
	return file_orders_updates_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrdersUpdatesResponse) GetOrdersUpdates() []*OrderUpdateMessage {
	if x != nil {
		return x.OrdersUpdates
	}
	return nil
}

var File_orders_updates_proto protoreflect.FileDescriptor

var file_orders_updates_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x32, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x12, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65,
	0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x64, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x42,
	0x4c, 0x5a, 0x4a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x61, 0x72, 0x6b, 0x2d, 0x6d, 0x61,
	0x69, 0x6c, 0x2d, 0x72, 0x75, 0x2f, 0x32, 0x30, 0x32, 0x34, 0x5f, 0x32, 0x5f, 0x6b, 0x6f, 0x74,
	0x79, 0x61, 0x72, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_updates_proto_rawDescOnce sync.Once
	file_orders_updates_proto_rawDescData = file_orders_updates_proto_rawDesc
)

func file_orders_updates_proto_rawDescGZIP() []byte {
	file_orders_updates_proto_rawDescOnce.Do(func() {
		file_orders_updates_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_updates_proto_rawDescData)
	})
	return file_orders_updates_proto_rawDescData
}

var file_orders_updates_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orders_updates_proto_goTypes = []any{
	(*GetOrdersUpdatesRequest)(nil),  // 0: notifications.GetOrdersUpdatesRequest
	(*OrderUpdateMessage)(nil),       // 1: notifications.OrderUpdateMessage
	(*GetOrdersUpdatesResponse)(nil), // 2: notifications.GetOrdersUpdatesResponse
}
var file_orders_updates_proto_depIdxs = []int32{
	1, // 0: notifications.GetOrdersUpdatesResponse.orders_updates:type_name -> notifications.OrderUpdateMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orders_updates_proto_init() }
func file_orders_updates_proto_init() {
	if File_orders_updates_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_updates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orders_updates_proto_goTypes,
		DependencyIndexes: file_orders_updates_proto_depIdxs,
		MessageInfos:      file_orders_updates_proto_msgTypes,
	}.Build()
	File_orders_updates_proto = out.File
	file_orders_updates_proto_rawDesc = nil
	file_orders_updates_proto_goTypes = nil
	file_orders_updates_proto_depIdxs = nil
}