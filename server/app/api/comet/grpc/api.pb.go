// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/comet/grpc/api.proto

package grpc

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/go-kratos/kratos/v2/errors"
	_ "github.com/longXboy/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

//
// v1.0.0
// protocol
type Proto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ver  int32  `protobuf:"varint,1,opt,name=ver,proto3" json:"ver,omitempty"`
	Op   int32  `protobuf:"varint,2,opt,name=op,proto3" json:"op,omitempty"`
	Seq  int32  `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Proto) Reset() {
	*x = Proto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proto) ProtoMessage() {}

func (x *Proto) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proto.ProtoReflect.Descriptor instead.
func (*Proto) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{0}
}

func (x *Proto) GetVer() int32 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *Proto) GetOp() int32 {
	if x != nil {
		return x.Op
	}
	return 0
}

func (x *Proto) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Proto) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{1}
}

type PushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys    []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	ProtoOp int32    `protobuf:"varint,3,opt,name=protoOp,proto3" json:"protoOp,omitempty"`
	Proto   *Proto   `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
}

func (x *PushMsgReq) Reset() {
	*x = PushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReq) ProtoMessage() {}

func (x *PushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReq.ProtoReflect.Descriptor instead.
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{2}
}

func (x *PushMsgReq) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *PushMsgReq) GetProtoOp() int32 {
	if x != nil {
		return x.ProtoOp
	}
	return 0
}

func (x *PushMsgReq) GetProto() *Proto {
	if x != nil {
		return x.Proto
	}
	return nil
}

type PushMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushMsgReply) Reset() {
	*x = PushMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReply) ProtoMessage() {}

func (x *PushMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReply.ProtoReflect.Descriptor instead.
func (*PushMsgReply) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{3}
}

type BroadcastReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoOp int32  `protobuf:"varint,1,opt,name=protoOp,proto3" json:"protoOp,omitempty"`
	Proto   *Proto `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	Speed   int32  `protobuf:"varint,3,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *BroadcastReq) Reset() {
	*x = BroadcastReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReq) ProtoMessage() {}

func (x *BroadcastReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReq.ProtoReflect.Descriptor instead.
func (*BroadcastReq) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{4}
}

func (x *BroadcastReq) GetProtoOp() int32 {
	if x != nil {
		return x.ProtoOp
	}
	return 0
}

func (x *BroadcastReq) GetProto() *Proto {
	if x != nil {
		return x.Proto
	}
	return nil
}

func (x *BroadcastReq) GetSpeed() int32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

type BroadcastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastReply) Reset() {
	*x = BroadcastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReply) ProtoMessage() {}

func (x *BroadcastReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReply.ProtoReflect.Descriptor instead.
func (*BroadcastReply) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{5}
}

type BroadcastRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Proto  *Proto `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
}

func (x *BroadcastRoomReq) Reset() {
	*x = BroadcastRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRoomReq) ProtoMessage() {}

func (x *BroadcastRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRoomReq.ProtoReflect.Descriptor instead.
func (*BroadcastRoomReq) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{6}
}

func (x *BroadcastRoomReq) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *BroadcastRoomReq) GetProto() *Proto {
	if x != nil {
		return x.Proto
	}
	return nil
}

type BroadcastRoomReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastRoomReply) Reset() {
	*x = BroadcastRoomReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRoomReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRoomReply) ProtoMessage() {}

func (x *BroadcastRoomReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRoomReply.ProtoReflect.Descriptor instead.
func (*BroadcastRoomReply) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{7}
}

type RoomsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomsReq) Reset() {
	*x = RoomsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomsReq) ProtoMessage() {}

func (x *RoomsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomsReq.ProtoReflect.Descriptor instead.
func (*RoomsReq) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{8}
}

type RoomsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms map[string]bool `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *RoomsReply) Reset() {
	*x = RoomsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_comet_grpc_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomsReply) ProtoMessage() {}

func (x *RoomsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_comet_grpc_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomsReply.ProtoReflect.Descriptor instead.
func (*RoomsReply) Descriptor() ([]byte, []int) {
	return file_api_comet_grpc_api_proto_rawDescGZIP(), []int{9}
}

func (x *RoomsReply) GetRooms() map[string]bool {
	if x != nil {
		return x.Rooms
	}
	return nil
}

var File_api_comet_grpc_api_proto protoreflect.FileDescriptor

var file_api_comet_grpc_api_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4f, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x67, 0x0a, 0x0a, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4f, 0x70, 0x12, 0x2b, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x6b, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4f, 0x70, 0x12, 0x2b,
	0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x22, 0x10, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x57, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12,
	0x2b, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x22, 0x83,
	0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a,
	0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x9a, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x65, 0x74, 0x12, 0x34,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x07, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x49, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x0d, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x65,
	0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x19, 0x5a, 0x17, 0x65, 0x64, 0x75, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x65, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_comet_grpc_api_proto_rawDescOnce sync.Once
	file_api_comet_grpc_api_proto_rawDescData = file_api_comet_grpc_api_proto_rawDesc
)

func file_api_comet_grpc_api_proto_rawDescGZIP() []byte {
	file_api_comet_grpc_api_proto_rawDescOnce.Do(func() {
		file_api_comet_grpc_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_comet_grpc_api_proto_rawDescData)
	})
	return file_api_comet_grpc_api_proto_rawDescData
}

var file_api_comet_grpc_api_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_comet_grpc_api_proto_goTypes = []interface{}{
	(*Proto)(nil),              // 0: api.comet.grpc.Proto
	(*Empty)(nil),              // 1: api.comet.grpc.Empty
	(*PushMsgReq)(nil),         // 2: api.comet.grpc.PushMsgReq
	(*PushMsgReply)(nil),       // 3: api.comet.grpc.PushMsgReply
	(*BroadcastReq)(nil),       // 4: api.comet.grpc.BroadcastReq
	(*BroadcastReply)(nil),     // 5: api.comet.grpc.BroadcastReply
	(*BroadcastRoomReq)(nil),   // 6: api.comet.grpc.BroadcastRoomReq
	(*BroadcastRoomReply)(nil), // 7: api.comet.grpc.BroadcastRoomReply
	(*RoomsReq)(nil),           // 8: api.comet.grpc.RoomsReq
	(*RoomsReply)(nil),         // 9: api.comet.grpc.RoomsReply
	nil,                        // 10: api.comet.grpc.RoomsReply.RoomsEntry
}
var file_api_comet_grpc_api_proto_depIdxs = []int32{
	0,  // 0: api.comet.grpc.PushMsgReq.proto:type_name -> api.comet.grpc.Proto
	0,  // 1: api.comet.grpc.BroadcastReq.proto:type_name -> api.comet.grpc.Proto
	0,  // 2: api.comet.grpc.BroadcastRoomReq.proto:type_name -> api.comet.grpc.Proto
	10, // 3: api.comet.grpc.RoomsReply.rooms:type_name -> api.comet.grpc.RoomsReply.RoomsEntry
	1,  // 4: api.comet.grpc.Comet.Ping:input_type -> api.comet.grpc.Empty
	1,  // 5: api.comet.grpc.Comet.Close:input_type -> api.comet.grpc.Empty
	2,  // 6: api.comet.grpc.Comet.PushMsg:input_type -> api.comet.grpc.PushMsgReq
	4,  // 7: api.comet.grpc.Comet.Broadcast:input_type -> api.comet.grpc.BroadcastReq
	6,  // 8: api.comet.grpc.Comet.BroadcastRoom:input_type -> api.comet.grpc.BroadcastRoomReq
	8,  // 9: api.comet.grpc.Comet.Rooms:input_type -> api.comet.grpc.RoomsReq
	1,  // 10: api.comet.grpc.Comet.Ping:output_type -> api.comet.grpc.Empty
	1,  // 11: api.comet.grpc.Comet.Close:output_type -> api.comet.grpc.Empty
	3,  // 12: api.comet.grpc.Comet.PushMsg:output_type -> api.comet.grpc.PushMsgReply
	5,  // 13: api.comet.grpc.Comet.Broadcast:output_type -> api.comet.grpc.BroadcastReply
	7,  // 14: api.comet.grpc.Comet.BroadcastRoom:output_type -> api.comet.grpc.BroadcastRoomReply
	9,  // 15: api.comet.grpc.Comet.Rooms:output_type -> api.comet.grpc.RoomsReply
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_comet_grpc_api_proto_init() }
func file_api_comet_grpc_api_proto_init() {
	if File_api_comet_grpc_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_comet_grpc_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proto); i {
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
		file_api_comet_grpc_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_comet_grpc_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReq); i {
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
		file_api_comet_grpc_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReply); i {
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
		file_api_comet_grpc_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReq); i {
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
		file_api_comet_grpc_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReply); i {
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
		file_api_comet_grpc_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRoomReq); i {
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
		file_api_comet_grpc_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastRoomReply); i {
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
		file_api_comet_grpc_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomsReq); i {
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
		file_api_comet_grpc_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomsReply); i {
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
			RawDescriptor: file_api_comet_grpc_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_comet_grpc_api_proto_goTypes,
		DependencyIndexes: file_api_comet_grpc_api_proto_depIdxs,
		MessageInfos:      file_api_comet_grpc_api_proto_msgTypes,
	}.Build()
	File_api_comet_grpc_api_proto = out.File
	file_api_comet_grpc_api_proto_rawDesc = nil
	file_api_comet_grpc_api_proto_goTypes = nil
	file_api_comet_grpc_api_proto_depIdxs = nil
}
