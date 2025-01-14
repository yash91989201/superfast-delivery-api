// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: inventory.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ItemStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemId     string                 `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity   int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	RestockQty int32                  `protobuf:"varint,4,opt,name=restock_qty,json=restockQty,proto3" json:"restock_qty,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ItemStock) Reset() {
	*x = ItemStock{}
	mi := &file_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemStock) ProtoMessage() {}

func (x *ItemStock) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemStock.ProtoReflect.Descriptor instead.
func (*ItemStock) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *ItemStock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemStock) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *ItemStock) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ItemStock) GetRestockQty() int32 {
	if x != nil {
		return x.RestockQty
	}
	return 0
}

func (x *ItemStock) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type VariantStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VariantId  string                 `protobuf:"bytes,2,opt,name=variant_id,json=variantId,proto3" json:"variant_id,omitempty"`
	Quantity   int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	RestockQty int32                  `protobuf:"varint,4,opt,name=restock_qty,json=restockQty,proto3" json:"restock_qty,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *VariantStock) Reset() {
	*x = VariantStock{}
	mi := &file_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VariantStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantStock) ProtoMessage() {}

func (x *VariantStock) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantStock.ProtoReflect.Descriptor instead.
func (*VariantStock) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *VariantStock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VariantStock) GetVariantId() string {
	if x != nil {
		return x.VariantId
	}
	return ""
}

func (x *VariantStock) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *VariantStock) GetRestockQty() int32 {
	if x != nil {
		return x.RestockQty
	}
	return 0
}

func (x *VariantStock) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type AddonStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AddonId    string                 `protobuf:"bytes,2,opt,name=addon_id,json=addonId,proto3" json:"addon_id,omitempty"`
	Quantity   int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	RestockQty int32                  `protobuf:"varint,4,opt,name=restock_qty,json=restockQty,proto3" json:"restock_qty,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *AddonStock) Reset() {
	*x = AddonStock{}
	mi := &file_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddonStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddonStock) ProtoMessage() {}

func (x *AddonStock) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddonStock.ProtoReflect.Descriptor instead.
func (*AddonStock) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *AddonStock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddonStock) GetAddonId() string {
	if x != nil {
		return x.AddonId
	}
	return ""
}

func (x *AddonStock) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddonStock) GetRestockQty() int32 {
	if x != nil {
		return x.RestockQty
	}
	return 0
}

func (x *AddonStock) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateItemStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CreateItemStockReq) Reset() {
	*x = CreateItemStockReq{}
	mi := &file_inventory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateItemStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateItemStockReq) ProtoMessage() {}

func (x *CreateItemStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateItemStockReq.ProtoReflect.Descriptor instead.
func (*CreateItemStockReq) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *CreateItemStockReq) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *CreateItemStockReq) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateVariantStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VariantId string `protobuf:"bytes,1,opt,name=variant_id,json=variantId,proto3" json:"variant_id,omitempty"`
	Quantity  int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CreateVariantStockReq) Reset() {
	*x = CreateVariantStockReq{}
	mi := &file_inventory_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVariantStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVariantStockReq) ProtoMessage() {}

func (x *CreateVariantStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVariantStockReq.ProtoReflect.Descriptor instead.
func (*CreateVariantStockReq) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVariantStockReq) GetVariantId() string {
	if x != nil {
		return x.VariantId
	}
	return ""
}

func (x *CreateVariantStockReq) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateAddonStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddonId  string `protobuf:"bytes,1,opt,name=addon_id,json=addonId,proto3" json:"addon_id,omitempty"`
	Quantity int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CreateAddonStockReq) Reset() {
	*x = CreateAddonStockReq{}
	mi := &file_inventory_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAddonStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAddonStockReq) ProtoMessage() {}

func (x *CreateAddonStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAddonStockReq.ProtoReflect.Descriptor instead.
func (*CreateAddonStockReq) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAddonStockReq) GetAddonId() string {
	if x != nil {
		return x.AddonId
	}
	return ""
}

func (x *CreateAddonStockReq) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x71, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x51, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb5, 0x01, 0x0a, 0x0c, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x71, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x51,
	0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xaf, 0x01,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x64, 0x64, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x71,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x51, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x49, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x52, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x4c,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x32, 0xd2, 0x01, 0x0a,
	0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x6f,
	0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22,
	0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x73, 0x68, 0x39, 0x31, 0x39, 0x38, 0x39, 0x32, 0x30, 0x31, 0x2f, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x66, 0x61, 0x73, 0x74, 0x2d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_inventory_proto_goTypes = []any{
	(*ItemStock)(nil),             // 0: pb.ItemStock
	(*VariantStock)(nil),          // 1: pb.VariantStock
	(*AddonStock)(nil),            // 2: pb.AddonStock
	(*CreateItemStockReq)(nil),    // 3: pb.CreateItemStockReq
	(*CreateVariantStockReq)(nil), // 4: pb.CreateVariantStockReq
	(*CreateAddonStockReq)(nil),   // 5: pb.CreateAddonStockReq
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_inventory_proto_depIdxs = []int32{
	6, // 0: pb.ItemStock.updated_at:type_name -> google.protobuf.Timestamp
	6, // 1: pb.VariantStock.updated_at:type_name -> google.protobuf.Timestamp
	6, // 2: pb.AddonStock.updated_at:type_name -> google.protobuf.Timestamp
	3, // 3: pb.InventoryService.CreateItemStock:input_type -> pb.CreateItemStockReq
	4, // 4: pb.InventoryService.CreateVariantStock:input_type -> pb.CreateVariantStockReq
	5, // 5: pb.InventoryService.CreateAddonStock:input_type -> pb.CreateAddonStockReq
	0, // 6: pb.InventoryService.CreateItemStock:output_type -> pb.ItemStock
	1, // 7: pb.InventoryService.CreateVariantStock:output_type -> pb.VariantStock
	2, // 8: pb.InventoryService.CreateAddonStock:output_type -> pb.AddonStock
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}