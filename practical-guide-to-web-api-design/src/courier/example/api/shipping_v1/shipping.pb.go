// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.28.3
// source: proto/shipping/v1/shipping.proto

package shipping_v1

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

type ShippingStatusResponse_Status int32

const (
	ShippingStatusResponse_UNSPECIFIED ShippingStatusResponse_Status = 0
	// 発送待ち
	ShippingStatusResponse_AWAITING_SHIPMENT ShippingStatusResponse_Status = 1
	// 発送済み（配送中）
	ShippingStatusResponse_OUT_FOR_DELIVERY ShippingStatusResponse_Status = 2
	// 配達済み
	ShippingStatusResponse_DELIVERED ShippingStatusResponse_Status = 3
)

// Enum value maps for ShippingStatusResponse_Status.
var (
	ShippingStatusResponse_Status_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "AWAITING_SHIPMENT",
		2: "OUT_FOR_DELIVERY",
		3: "DELIVERED",
	}
	ShippingStatusResponse_Status_value = map[string]int32{
		"UNSPECIFIED":       0,
		"AWAITING_SHIPMENT": 1,
		"OUT_FOR_DELIVERY":  2,
		"DELIVERED":         3,
	}
)

func (x ShippingStatusResponse_Status) Enum() *ShippingStatusResponse_Status {
	p := new(ShippingStatusResponse_Status)
	*p = x
	return p
}

func (x ShippingStatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShippingStatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_shipping_v1_shipping_proto_enumTypes[0].Descriptor()
}

func (ShippingStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_shipping_v1_shipping_proto_enumTypes[0]
}

func (x ShippingStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShippingStatusResponse_Status.Descriptor instead.
func (ShippingStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_shipping_v1_shipping_proto_rawDescGZIP(), []int{3, 0}
}

// *
// CreateShippingは、発送する商品をまとめたパッケージを受け取り発送します。
// - 発送のステータスを返すためのshipping_idを返します。
type CreateShippingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CreateShippingRequest) Reset() {
	*x = CreateShippingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shipping_v1_shipping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShippingRequest) ProtoMessage() {}

func (x *CreateShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_v1_shipping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShippingRequest.ProtoReflect.Descriptor instead.
func (*CreateShippingRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_v1_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShippingRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CreateShippingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShippingId string `protobuf:"bytes,1,opt,name=shipping_id,json=shippingId,proto3" json:"shipping_id,omitempty"`
}

func (x *CreateShippingResponse) Reset() {
	*x = CreateShippingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shipping_v1_shipping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShippingResponse) ProtoMessage() {}

func (x *CreateShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_v1_shipping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShippingResponse.ProtoReflect.Descriptor instead.
func (*CreateShippingResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_v1_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShippingResponse) GetShippingId() string {
	if x != nil {
		return x.ShippingId
	}
	return ""
}

// *
// ShippingStatusは、商品の発送のステータスを返します。
//
// [エラー]
// - InvalidArgument:
//   - shipping_idが空文字列
//
// - NotFound:
//   - shipping_idで指定された発送がない
type ShippingStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShippingId string `protobuf:"bytes,1,opt,name=shipping_id,json=shippingId,proto3" json:"shipping_id,omitempty"`
}

func (x *ShippingStatusRequest) Reset() {
	*x = ShippingStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shipping_v1_shipping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingStatusRequest) ProtoMessage() {}

func (x *ShippingStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_v1_shipping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingStatusRequest.ProtoReflect.Descriptor instead.
func (*ShippingStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_v1_shipping_proto_rawDescGZIP(), []int{2}
}

func (x *ShippingStatusRequest) GetShippingId() string {
	if x != nil {
		return x.ShippingId
	}
	return ""
}

type ShippingStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShippidId string                        `protobuf:"bytes,1,opt,name=shippid_id,json=shippidId,proto3" json:"shippid_id,omitempty"`
	Status    ShippingStatusResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=yoshikishibata.courier.example.api.shipping.v1.ShippingStatusResponse_Status" json:"status,omitempty"`
}

func (x *ShippingStatusResponse) Reset() {
	*x = ShippingStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shipping_v1_shipping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShippingStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingStatusResponse) ProtoMessage() {}

func (x *ShippingStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_v1_shipping_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingStatusResponse.ProtoReflect.Descriptor instead.
func (*ShippingStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_v1_shipping_proto_rawDescGZIP(), []int{3}
}

func (x *ShippingStatusResponse) GetShippidId() string {
	if x != nil {
		return x.ShippidId
	}
	return ""
}

func (x *ShippingStatusResponse) GetStatus() ShippingStatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return ShippingStatusResponse_UNSPECIFIED
}

var File_proto_shipping_v1_shipping_proto protoreflect.FileDescriptor

var file_proto_shipping_v1_shipping_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6b, 0x69, 0x73, 0x68, 0x69, 0x62, 0x61,
	0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x22, 0x32, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x22, 0x38, 0x0a, 0x15, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0xf5, 0x01, 0x0a, 0x16,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x64, 0x49, 0x64, 0x12, 0x65, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6b, 0x69, 0x73,
	0x68, 0x69, 0x62, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x57, 0x41, 0x49, 0x54,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x4f, 0x55, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45,
	0x52, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45,
	0x44, 0x10, 0x03, 0x32, 0xc2, 0x02, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x99, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x45, 0x2e, 0x79, 0x6f,
	0x73, 0x68, 0x69, 0x6b, 0x69, 0x73, 0x68, 0x69, 0x62, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x46, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6b, 0x69, 0x73, 0x68, 0x69, 0x62,
	0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x99, 0x01, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x45, 0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6b,
	0x69, 0x73, 0x68, 0x69, 0x62, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46,
	0x2e, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x6b, 0x69, 0x73, 0x68, 0x69, 0x62, 0x61, 0x74, 0x61, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shipping_v1_shipping_proto_rawDescOnce sync.Once
	file_proto_shipping_v1_shipping_proto_rawDescData = file_proto_shipping_v1_shipping_proto_rawDesc
)

func file_proto_shipping_v1_shipping_proto_rawDescGZIP() []byte {
	file_proto_shipping_v1_shipping_proto_rawDescOnce.Do(func() {
		file_proto_shipping_v1_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shipping_v1_shipping_proto_rawDescData)
	})
	return file_proto_shipping_v1_shipping_proto_rawDescData
}

var file_proto_shipping_v1_shipping_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_shipping_v1_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_shipping_v1_shipping_proto_goTypes = []interface{}{
	(ShippingStatusResponse_Status)(0), // 0: yoshikishibata.courier.example.api.shipping.v1.ShippingStatusResponse.Status
	(*CreateShippingRequest)(nil),      // 1: yoshikishibata.courier.example.api.shipping.v1.CreateShippingRequest
	(*CreateShippingResponse)(nil),     // 2: yoshikishibata.courier.example.api.shipping.v1.CreateShippingResponse
	(*ShippingStatusRequest)(nil),      // 3: yoshikishibata.courier.example.api.shipping.v1.ShippingStatusRequest
	(*ShippingStatusResponse)(nil),     // 4: yoshikishibata.courier.example.api.shipping.v1.ShippingStatusResponse
}
var file_proto_shipping_v1_shipping_proto_depIdxs = []int32{
	0, // 0: yoshikishibata.courier.example.api.shipping.v1.ShippingStatusResponse.status:type_name -> yoshikishibata.courier.example.api.shipping.v1.ShippingStatusResponse.Status
	1, // 1: yoshikishibata.courier.example.api.shipping.v1.Shipping.Create:input_type -> yoshikishibata.courier.example.api.shipping.v1.CreateShippingRequest
	3, // 2: yoshikishibata.courier.example.api.shipping.v1.Shipping.Status:input_type -> yoshikishibata.courier.example.api.shipping.v1.ShippingStatusRequest
	2, // 3: yoshikishibata.courier.example.api.shipping.v1.Shipping.Create:output_type -> yoshikishibata.courier.example.api.shipping.v1.CreateShippingResponse
	4, // 4: yoshikishibata.courier.example.api.shipping.v1.Shipping.Status:output_type -> yoshikishibata.courier.example.api.shipping.v1.ShippingStatusResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_shipping_v1_shipping_proto_init() }
func file_proto_shipping_v1_shipping_proto_init() {
	if File_proto_shipping_v1_shipping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shipping_v1_shipping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShippingRequest); i {
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
		file_proto_shipping_v1_shipping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShippingResponse); i {
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
		file_proto_shipping_v1_shipping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingStatusRequest); i {
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
		file_proto_shipping_v1_shipping_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShippingStatusResponse); i {
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
			RawDescriptor: file_proto_shipping_v1_shipping_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shipping_v1_shipping_proto_goTypes,
		DependencyIndexes: file_proto_shipping_v1_shipping_proto_depIdxs,
		EnumInfos:         file_proto_shipping_v1_shipping_proto_enumTypes,
		MessageInfos:      file_proto_shipping_v1_shipping_proto_msgTypes,
	}.Build()
	File_proto_shipping_v1_shipping_proto = out.File
	file_proto_shipping_v1_shipping_proto_rawDesc = nil
	file_proto_shipping_v1_shipping_proto_goTypes = nil
	file_proto_shipping_v1_shipping_proto_depIdxs = nil
}