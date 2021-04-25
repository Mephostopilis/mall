package model

import (
	"time"
)

const (
	VipModeLimit   = 1
	VipModeInfinit = 2
)

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int64
	Content string
	Author  string
}

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type Zone struct {
	Key     string // key用来表示区域key，如tk001
	IP      string
	Service string // 每个区域对应一个服务
	Target  string
}

// Vip vip商品
type Vip struct {
	Level int32
	Title string
	Price int32
	// Deadline xtime.Time
}

type Device struct {
	Mac  string
	Ips  []string
	Lock bool
}

type User struct {
	ID          uint64
	Name        string
	Vip         int32
	VipMode     int32
	VipDeadline time.Time
	Devs        map[string]Device
	Zones       map[string]uint64
}

type Port struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	Cipher               string   `protobuf:"bytes,2,opt,name=Cipher,proto3" json:"cipher"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"password"`
	Host                 string   `protobuf:"bytes,4,opt,name=Host,proto3" json:"host"`
	Port                 string   `protobuf:"bytes,5,opt,name=Port,proto3" json:"port"`
	Remaining            string   `protobuf:"bytes,6,opt,name=Remaining,proto3" json:"remaining"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type GetPortListResp struct {
	List []Port `json:"list"`
}

type GetPortResp struct {
	P                    *Port    `protobuf:"bytes,1,opt,name=P,proto3" json:"p"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type NewPortResp struct {
	P                    *Port    `protobuf:"bytes,1,opt,name=P,proto3" json:"p"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type StatPortResp struct {
	Checkcnt uint64 `json:"checkcnt"`
}
