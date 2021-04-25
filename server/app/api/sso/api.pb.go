// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: api/sso/api.proto

package sso

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type ApiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason    string       `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message   string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Ttl       int32        `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	PageIndex int32        `protobuf:"varint,5,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize  int32        `protobuf:"varint,6,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Count     int32        `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"` // total
	Data      []*anypb.Any `protobuf:"bytes,8,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ApiReply) Reset() {
	*x = ApiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sso_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiReply) ProtoMessage() {}

func (x *ApiReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sso_api_proto_msgTypes[0]
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
	return file_api_sso_api_proto_rawDescGZIP(), []int{0}
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

func (x *ApiReply) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *ApiReply) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ApiReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ApiReply) GetData() []*anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type DataPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataScope string `protobuf:"bytes,1,opt,name=dataScope,proto3" json:"dataScope,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	RoleId    uint64 `protobuf:"varint,3,opt,name=roleId,proto3" json:"roleId,omitempty"`
	RoleKey   string `protobuf:"bytes,4,opt,name=roleKey,proto3" json:"roleKey,omitempty"`
	DeptId    uint64 `protobuf:"varint,5,opt,name=deptId,proto3" json:"deptId,omitempty"`
	PostId    uint64 `protobuf:"varint,6,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *DataPermission) Reset() {
	*x = DataPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sso_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPermission) ProtoMessage() {}

func (x *DataPermission) ProtoReflect() protoreflect.Message {
	mi := &file_api_sso_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPermission.ProtoReflect.Descriptor instead.
func (*DataPermission) Descriptor() ([]byte, []int) {
	return file_api_sso_api_proto_rawDescGZIP(), []int{1}
}

func (x *DataPermission) GetDataScope() string {
	if x != nil {
		return x.DataScope
	}
	return ""
}

func (x *DataPermission) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DataPermission) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *DataPermission) GetRoleKey() string {
	if x != nil {
		return x.RoleKey
	}
	return ""
}

func (x *DataPermission) GetDeptId() uint64 {
	if x != nil {
		return x.DeptId
	}
	return 0
}

func (x *DataPermission) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type HttpReqInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserAgent  string `protobuf:"bytes,2,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	RemoteAddr string `protobuf:"bytes,3,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
}

func (x *HttpReqInfo) Reset() {
	*x = HttpReqInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sso_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpReqInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpReqInfo) ProtoMessage() {}

func (x *HttpReqInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_sso_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpReqInfo.ProtoReflect.Descriptor instead.
func (*HttpReqInfo) Descriptor() ([]byte, []int) {
	return file_api_sso_api_proto_rawDescGZIP(), []int{2}
}

func (x *HttpReqInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *HttpReqInfo) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *HttpReqInfo) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

type Dept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeptId    int64                  `protobuf:"varint,1,opt,name=deptId,proto3" json:"deptId,omitempty"`
	ParentId  int64                  `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	DeptPath  string                 `protobuf:"bytes,3,opt,name=deptPath,proto3" json:"deptPath,omitempty"`
	DeptName  string                 `protobuf:"bytes,4,opt,name=deptName,proto3" json:"deptName,omitempty"`
	Sort      int32                  `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Leader    string                 `protobuf:"bytes,6,opt,name=leader,proto3" json:"leader,omitempty"`
	Phone     string                 `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Email     string                 `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Status    string                 `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Dept) Reset() {
	*x = Dept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sso_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dept) ProtoMessage() {}

func (x *Dept) ProtoReflect() protoreflect.Message {
	mi := &file_api_sso_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dept.ProtoReflect.Descriptor instead.
func (*Dept) Descriptor() ([]byte, []int) {
	return file_api_sso_api_proto_rawDescGZIP(), []int{3}
}

func (x *Dept) GetDeptId() int64 {
	if x != nil {
		return x.DeptId
	}
	return 0
}

func (x *Dept) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Dept) GetDeptPath() string {
	if x != nil {
		return x.DeptPath
	}
	return ""
}

func (x *Dept) GetDeptName() string {
	if x != nil {
		return x.DeptName
	}
	return ""
}

func (x *Dept) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Dept) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

func (x *Dept) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Dept) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Dept) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Dept) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId    int32                  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
	PostName  string                 `protobuf:"bytes,2,opt,name=postName,proto3" json:"postName,omitempty"`
	PostCode  string                 `protobuf:"bytes,3,opt,name=postCode,proto3" json:"postCode,omitempty"`
	Sort      int32                  `protobuf:"varint,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Status    string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Remark    string                 `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sso_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_api_sso_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_api_sso_api_proto_rawDescGZIP(), []int{4}
}

func (x *Post) GetPostId() int32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *Post) GetPostName() string {
	if x != nil {
		return x.PostName
	}
	return ""
}

func (x *Post) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *Post) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Post) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Post) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_api_sso_api_proto protoreflect.FileDescriptor

var file_api_sso_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x73, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x08, 0x41, 0x70, 0x69,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x22, 0x61, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x9c, 0x02, 0x0a, 0x04, 0x44, 0x65, 0x70, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x64, 0x65, 0x70, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x70, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xd4, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x1c, 0x0a, 0x07, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x73, 0x6f, 0x50, 0x01, 0x5a, 0x0f, 0x65, 0x64, 0x75, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x73, 0x6f, 0x3b, 0x73, 0x73, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_sso_api_proto_rawDescOnce sync.Once
	file_api_sso_api_proto_rawDescData = file_api_sso_api_proto_rawDesc
)

func file_api_sso_api_proto_rawDescGZIP() []byte {
	file_api_sso_api_proto_rawDescOnce.Do(func() {
		file_api_sso_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sso_api_proto_rawDescData)
	})
	return file_api_sso_api_proto_rawDescData
}

var file_api_sso_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_sso_api_proto_goTypes = []interface{}{
	(*ApiReply)(nil),              // 0: api.sso.ApiReply
	(*DataPermission)(nil),        // 1: api.sso.DataPermission
	(*HttpReqInfo)(nil),           // 2: api.sso.HttpReqInfo
	(*Dept)(nil),                  // 3: api.sso.Dept
	(*Post)(nil),                  // 4: api.sso.Post
	(*anypb.Any)(nil),             // 5: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_api_sso_api_proto_depIdxs = []int32{
	5, // 0: api.sso.ApiReply.data:type_name -> google.protobuf.Any
	6, // 1: api.sso.Dept.createdAt:type_name -> google.protobuf.Timestamp
	6, // 2: api.sso.Post.createdAt:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_sso_api_proto_init() }
func file_api_sso_api_proto_init() {
	if File_api_sso_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sso_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_sso_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPermission); i {
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
		file_api_sso_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpReqInfo); i {
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
		file_api_sso_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dept); i {
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
		file_api_sso_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
			RawDescriptor: file_api_sso_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_sso_api_proto_goTypes,
		DependencyIndexes: file_api_sso_api_proto_depIdxs,
		MessageInfos:      file_api_sso_api_proto_msgTypes,
	}.Build()
	File_api_sso_api_proto = out.File
	file_api_sso_api_proto_rawDesc = nil
	file_api_sso_api_proto_goTypes = nil
	file_api_sso_api_proto_depIdxs = nil
}
