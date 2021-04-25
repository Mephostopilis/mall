// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/oms/api.proto

package oms

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason  string       `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Ttl     int32        `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Count   int32        `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	Details []*anypb.Any `protobuf:"bytes,6,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *ApiReply) Reset() {
	*x = ApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_oms_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiReply) ProtoMessage() {}

func (x *ApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiReply.ProtoReflect.Descriptor instead.
func (*ApiReply) Descriptor() ([]byte, []int) {
	return file_api_oms_api_proto_rawDescGZIP(), []int{0}
}

func (x *ApiReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ApiReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ApiReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiReply) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *ApiReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ApiReply) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_oms_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_api_oms_api_proto_rawDescGZIP(), []int{1}
}

type CompanyAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompanyAddress) Reset() {
	*x = CompanyAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_oms_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyAddress) ProtoMessage() {}

func (x *CompanyAddress) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyAddress.ProtoReflect.Descriptor instead.
func (*CompanyAddress) Descriptor() ([]byte, []int) {
	return file_api_oms_api_proto_rawDescGZIP(), []int{2}
}

var File_api_oms_api_proto protoreflect.FileDescriptor

var file_api_oms_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6d, 0x73, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x08, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x1c, 0x0a,
	0x07, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x6d, 0x73, 0x50, 0x01, 0x5a, 0x0f, 0x65, 0x64, 0x75, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x73, 0x3b, 0x6f, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_oms_api_proto_rawDescOnce sync.Once
	file_api_oms_api_proto_rawDescData = file_api_oms_api_proto_rawDesc
)

func file_api_oms_api_proto_rawDescGZIP() []byte {
	file_api_oms_api_proto_rawDescOnce.Do(func() {
		file_api_oms_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_oms_api_proto_rawDescData)
	})
	return file_api_oms_api_proto_rawDescData
}

var file_api_oms_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_oms_api_proto_goTypes = []interface{}{
	(*ApiReply)(nil),       // 0: api.oms.ApiReply
	(*Order)(nil),          // 1: api.oms.Order
	(*CompanyAddress)(nil), // 2: api.oms.CompanyAddress
	(*anypb.Any)(nil),      // 3: google.protobuf.Any
}
var file_api_oms_api_proto_depIdxs = []int32{
	3, // 0: api.oms.ApiReply.details:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_oms_api_proto_init() }
func file_api_oms_api_proto_init() {
	if File_api_oms_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_oms_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiReply); i {
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
		file_api_oms_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_api_oms_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyAddress); i {
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
			RawDescriptor: file_api_oms_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_oms_api_proto_goTypes,
		DependencyIndexes: file_api_oms_api_proto_depIdxs,
		MessageInfos:      file_api_oms_api_proto_msgTypes,
	}.Build()
	File_api_oms_api_proto = out.File
	file_api_oms_api_proto_rawDesc = nil
	file_api_oms_api_proto_goTypes = nil
	file_api_oms_api_proto_depIdxs = nil
}
