// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: service/uuid/internal/conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Server  *Server  `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Data    *Data    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Reg     *Reg     `protobuf:"bytes,4,opt,name=reg,proto3" json:"reg,omitempty"`
	Logger  *Logger  `protobuf:"bytes,5,opt,name=logger,proto3" json:"logger,omitempty"`
	App     *App     `protobuf:"bytes,6,opt,name=app,proto3" json:"app,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetReg() *Reg {
	if x != nil {
		return x.Reg
	}
	return nil
}

func (x *Bootstrap) GetLogger() *Logger {
	if x != nil {
		return x.Logger
	}
	return nil
}

func (x *Bootstrap) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis    *Data_Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	Migrate  *Data_Migrate  `protobuf:"bytes,3,opt,name=migrate,proto3" json:"migrate,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Data) GetMigrate() *Data_Migrate {
	if x != nil {
		return x.Migrate
	}
	return nil
}

type Reg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etcd *Reg_Etcd `protobuf:"bytes,1,opt,name=etcd,proto3" json:"etcd,omitempty"`
}

func (x *Reg) Reset() {
	*x = Reg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reg) ProtoMessage() {}

func (x *Reg) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reg.ProtoReflect.Descriptor instead.
func (*Reg) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Reg) GetEtcd() *Reg_Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

type Logger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Stdout bool   `protobuf:"varint,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Level  int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Logger) Reset() {
	*x = Logger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger) ProtoMessage() {}

func (x *Logger) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger.ProtoReflect.Descriptor instead.
func (*Logger) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{5}
}

func (x *Logger) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Logger) GetStdout() bool {
	if x != nil {
		return x.Stdout
	}
	return false
}

func (x *Logger) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid *App_UUID `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{6}
}

func (x *App) GetUuid() *App_UUID {
	if x != nil {
		return x.Uuid
	}
	return nil
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Address string               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Address string               `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	DialTimeout  *durationpb.Duration `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Data_Redis) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

type Data_Migrate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Data_Migrate) Reset() {
	*x = Data_Migrate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Migrate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Migrate) ProtoMessage() {}

func (x *Data_Migrate) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Migrate.ProtoReflect.Descriptor instead.
func (*Data_Migrate) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{3, 2}
}

func (x *Data_Migrate) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Reg_Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints   []string             `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	DialTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=dialTimeout,proto3" json:"dialTimeout,omitempty"`
}

func (x *Reg_Etcd) Reset() {
	*x = Reg_Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reg_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reg_Etcd) ProtoMessage() {}

func (x *Reg_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reg_Etcd.ProtoReflect.Descriptor instead.
func (*Reg_Etcd) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Reg_Etcd) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Reg_Etcd) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

type App_UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId uint32 `protobuf:"varint,1,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
}

func (x *App_UUID) Reset() {
	*x = App_UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App_UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App_UUID) ProtoMessage() {}

func (x *App_UUID) ProtoReflect() protoreflect.Message {
	mi := &file_service_uuid_internal_conf_conf_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App_UUID.ProtoReflect.Descriptor instead.
func (*App_UUID) Descriptor() ([]byte, []int) {
	return file_service_uuid_internal_conf_conf_proto_rawDescGZIP(), []int{6, 0}
}

func (x *App_UUID) GetMachineId() uint32 {
	if x != nil {
		return x.MachineId
	}
	return 0
}

var File_service_uuid_internal_conf_conf_proto protoreflect.FileDescriptor

var file_service_uuid_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x0a, 0x03, 0x72, 0x65, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x75,
	0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x52, 0x03, 0x72, 0x65, 0x67,
	0x12, 0x29, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x03, 0x61,
	0x70, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x22, 0x37, 0x0a,
	0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc2, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2a, 0x0a,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x75,
	0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x6f, 0x0a, 0x04, 0x48, 0x54, 0x54,
	0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x6f, 0x0a, 0x04, 0x47, 0x52,
	0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xeb, 0x03, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x75, 0x69, 0x64,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x52, 0x07, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0xf1, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x3c,
	0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3c, 0x0a, 0x0c,
	0x64, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64,
	0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x1d, 0x0a, 0x07, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x91, 0x01, 0x0a, 0x03, 0x52, 0x65,
	0x67, 0x12, 0x27, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x67, 0x2e,
	0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x1a, 0x61, 0x0a, 0x04, 0x45, 0x74,
	0x63, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x4a, 0x0a,
	0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x64,
	0x6f, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x55, 0x0a, 0x03, 0x41, 0x70, 0x70,
	0x12, 0x27, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x1a, 0x25, 0x0a, 0x04, 0x55, 0x55, 0x49,
	0x44, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x64, 0x75, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_uuid_internal_conf_conf_proto_rawDescOnce sync.Once
	file_service_uuid_internal_conf_conf_proto_rawDescData = file_service_uuid_internal_conf_conf_proto_rawDesc
)

func file_service_uuid_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_service_uuid_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_service_uuid_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_uuid_internal_conf_conf_proto_rawDescData)
	})
	return file_service_uuid_internal_conf_conf_proto_rawDescData
}

var file_service_uuid_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_service_uuid_internal_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: uuid.conf.Bootstrap
	(*Service)(nil),             // 1: uuid.conf.Service
	(*Server)(nil),              // 2: uuid.conf.Server
	(*Data)(nil),                // 3: uuid.conf.Data
	(*Reg)(nil),                 // 4: uuid.conf.Reg
	(*Logger)(nil),              // 5: uuid.conf.Logger
	(*App)(nil),                 // 6: uuid.conf.App
	(*Server_HTTP)(nil),         // 7: uuid.conf.Server.HTTP
	(*Server_GRPC)(nil),         // 8: uuid.conf.Server.GRPC
	(*Data_Database)(nil),       // 9: uuid.conf.Data.Database
	(*Data_Redis)(nil),          // 10: uuid.conf.Data.Redis
	(*Data_Migrate)(nil),        // 11: uuid.conf.Data.Migrate
	(*Reg_Etcd)(nil),            // 12: uuid.conf.Reg.Etcd
	(*App_UUID)(nil),            // 13: uuid.conf.App.UUID
	(*durationpb.Duration)(nil), // 14: google.protobuf.Duration
}
var file_service_uuid_internal_conf_conf_proto_depIdxs = []int32{
	1,  // 0: uuid.conf.Bootstrap.service:type_name -> uuid.conf.Service
	2,  // 1: uuid.conf.Bootstrap.server:type_name -> uuid.conf.Server
	3,  // 2: uuid.conf.Bootstrap.data:type_name -> uuid.conf.Data
	4,  // 3: uuid.conf.Bootstrap.reg:type_name -> uuid.conf.Reg
	5,  // 4: uuid.conf.Bootstrap.logger:type_name -> uuid.conf.Logger
	6,  // 5: uuid.conf.Bootstrap.app:type_name -> uuid.conf.App
	7,  // 6: uuid.conf.Server.http:type_name -> uuid.conf.Server.HTTP
	8,  // 7: uuid.conf.Server.grpc:type_name -> uuid.conf.Server.GRPC
	9,  // 8: uuid.conf.Data.database:type_name -> uuid.conf.Data.Database
	10, // 9: uuid.conf.Data.redis:type_name -> uuid.conf.Data.Redis
	11, // 10: uuid.conf.Data.migrate:type_name -> uuid.conf.Data.Migrate
	12, // 11: uuid.conf.Reg.etcd:type_name -> uuid.conf.Reg.Etcd
	13, // 12: uuid.conf.App.uuid:type_name -> uuid.conf.App.UUID
	14, // 13: uuid.conf.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	14, // 14: uuid.conf.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	14, // 15: uuid.conf.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	14, // 16: uuid.conf.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	14, // 17: uuid.conf.Data.Redis.dial_timeout:type_name -> google.protobuf.Duration
	14, // 18: uuid.conf.Reg.Etcd.dialTimeout:type_name -> google.protobuf.Duration
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_service_uuid_internal_conf_conf_proto_init() }
func file_service_uuid_internal_conf_conf_proto_init() {
	if File_service_uuid_internal_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_uuid_internal_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reg); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Database); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Migrate); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reg_Etcd); i {
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
		file_service_uuid_internal_conf_conf_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App_UUID); i {
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
			RawDescriptor: file_service_uuid_internal_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_uuid_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_service_uuid_internal_conf_conf_proto_depIdxs,
		MessageInfos:      file_service_uuid_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_service_uuid_internal_conf_conf_proto = out.File
	file_service_uuid_internal_conf_conf_proto_rawDesc = nil
	file_service_uuid_internal_conf_conf_proto_goTypes = nil
	file_service_uuid_internal_conf_conf_proto_depIdxs = nil
}
