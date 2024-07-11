// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: reservation.proto

package reservation

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

type ReservationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RestaurantId    string `protobuf:"bytes,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	ReservationTime string `protobuf:"bytes,3,opt,name=reservation_time,json=reservationTime,proto3" json:"reservation_time,omitempty"`
}

func (x *ReservationDetails) Reset() {
	*x = ReservationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationDetails) ProtoMessage() {}

func (x *ReservationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationDetails.ProtoReflect.Descriptor instead.
func (*ReservationDetails) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *ReservationDetails) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReservationDetails) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *ReservationDetails) GetReservationTime() string {
	if x != nil {
		return x.ReservationTime
	}
	return ""
}

type ReservationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RestaurantId    string `protobuf:"bytes,3,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	ReservationTime string `protobuf:"bytes,4,opt,name=reservation_time,json=reservationTime,proto3" json:"reservation_time,omitempty"`
	Status          string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReservationInfo) Reset() {
	*x = ReservationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationInfo) ProtoMessage() {}

func (x *ReservationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationInfo.ProtoReflect.Descriptor instead.
func (*ReservationInfo) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *ReservationInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReservationInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReservationInfo) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *ReservationInfo) GetReservationTime() string {
	if x != nil {
		return x.ReservationTime
	}
	return ""
}

func (x *ReservationInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{3}
}

type ReservationOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuItemId string `protobuf:"bytes,1,opt,name=menu_item_id,json=menuItemId,proto3" json:"menu_item_id,omitempty"`
	Quantity   int32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ReservationOrder) Reset() {
	*x = ReservationOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationOrder) ProtoMessage() {}

func (x *ReservationOrder) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationOrder.ProtoReflect.Descriptor instead.
func (*ReservationOrder) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{4}
}

func (x *ReservationOrder) GetMenuItemId() string {
	if x != nil {
		return x.MenuItemId
	}
	return ""
}

func (x *ReservationOrder) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ReservationOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Order []*ReservationOrder `protobuf:"bytes,2,rep,name=order,proto3" json:"order,omitempty"`
}

func (x *ReservationOrders) Reset() {
	*x = ReservationOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationOrders) ProtoMessage() {}

func (x *ReservationOrders) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationOrders.ProtoReflect.Descriptor instead.
func (*ReservationOrders) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{5}
}

func (x *ReservationOrders) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReservationOrders) GetOrder() []*ReservationOrder {
	if x != nil {
		return x.Order
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RestaurantId string `protobuf:"bytes,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	Limit        int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset       int32  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{7}
}

func (x *Filter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Filter) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *Filter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Filter) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Reservations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*ReservationInfo `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *Reservations) Reset() {
	*x = Reservations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservations) ProtoMessage() {}

func (x *Reservations) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservations.ProtoReflect.Descriptor instead.
func (*Reservations) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{8}
}

func (x *Reservations) GetReservations() []*ReservationInfo {
	if x != nil {
		return x.Reservations
	}
	return nil
}

var File_reservation_proto protoreflect.FileDescriptor

var file_reservation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x7d, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xa2, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x50, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x58, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x28,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x22, 0x74, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x50,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x40,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0x81, 0x04, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x45, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0f, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x44, 0x1a, 0x1c,
	0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x44, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x11, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x13, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x49, 0x44, 0x12, 0x2b, 0x0a, 0x03, 0x50, 0x61, 0x79, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x43, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservation_proto_rawDescOnce sync.Once
	file_reservation_proto_rawDescData = file_reservation_proto_rawDesc
)

func file_reservation_proto_rawDescGZIP() []byte {
	file_reservation_proto_rawDescOnce.Do(func() {
		file_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_proto_rawDescData)
	})
	return file_reservation_proto_rawDescData
}

var file_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_reservation_proto_goTypes = []interface{}{
	(*ReservationDetails)(nil), // 0: reservation.ReservationDetails
	(*ReservationInfo)(nil),    // 1: reservation.ReservationInfo
	(*ID)(nil),                 // 2: reservation.ID
	(*Void)(nil),               // 3: reservation.Void
	(*ReservationOrder)(nil),   // 4: reservation.ReservationOrder
	(*ReservationOrders)(nil),  // 5: reservation.ReservationOrders
	(*Status)(nil),             // 6: reservation.Status
	(*Filter)(nil),             // 7: reservation.Filter
	(*Reservations)(nil),       // 8: reservation.Reservations
}
var file_reservation_proto_depIdxs = []int32{
	4,  // 0: reservation.ReservationOrders.order:type_name -> reservation.ReservationOrder
	1,  // 1: reservation.Reservations.reservations:type_name -> reservation.ReservationInfo
	0,  // 2: reservation.Reservation.CreateReservation:input_type -> reservation.ReservationDetails
	2,  // 3: reservation.Reservation.GetReservationByID:input_type -> reservation.ID
	1,  // 4: reservation.Reservation.UpdateReservation:input_type -> reservation.ReservationInfo
	2,  // 5: reservation.Reservation.DeleteReservation:input_type -> reservation.ID
	2,  // 6: reservation.Reservation.ValidateReservation:input_type -> reservation.ID
	5,  // 7: reservation.Reservation.Order:input_type -> reservation.ReservationOrders
	2,  // 8: reservation.Reservation.Pay:input_type -> reservation.ID
	7,  // 9: reservation.Reservation.FetchReservations:input_type -> reservation.Filter
	2,  // 10: reservation.Reservation.CreateReservation:output_type -> reservation.ID
	1,  // 11: reservation.Reservation.GetReservationByID:output_type -> reservation.ReservationInfo
	3,  // 12: reservation.Reservation.UpdateReservation:output_type -> reservation.Void
	3,  // 13: reservation.Reservation.DeleteReservation:output_type -> reservation.Void
	6,  // 14: reservation.Reservation.ValidateReservation:output_type -> reservation.Status
	2,  // 15: reservation.Reservation.Order:output_type -> reservation.ID
	6,  // 16: reservation.Reservation.Pay:output_type -> reservation.Status
	8,  // 17: reservation.Reservation.FetchReservations:output_type -> reservation.Reservations
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_reservation_proto_init() }
func file_reservation_proto_init() {
	if File_reservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationDetails); i {
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
		file_reservation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationInfo); i {
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
		file_reservation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_reservation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_reservation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationOrder); i {
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
		file_reservation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationOrders); i {
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
		file_reservation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_reservation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_reservation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservations); i {
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
			RawDescriptor: file_reservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservation_proto_goTypes,
		DependencyIndexes: file_reservation_proto_depIdxs,
		MessageInfos:      file_reservation_proto_msgTypes,
	}.Build()
	File_reservation_proto = out.File
	file_reservation_proto_rawDesc = nil
	file_reservation_proto_goTypes = nil
	file_reservation_proto_depIdxs = nil
}
