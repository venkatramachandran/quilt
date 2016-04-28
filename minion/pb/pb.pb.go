// Code generated by protoc-gen-go.
// source: pb/pb.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/pb.proto

It has these top-level messages:
	MinionConfig
	Reply
	Request
	EtcdMembers
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type MinionConfig_Role int32

const (
	MinionConfig_NONE   MinionConfig_Role = 0
	MinionConfig_WORKER MinionConfig_Role = 1
	MinionConfig_MASTER MinionConfig_Role = 2
)

var MinionConfig_Role_name = map[int32]string{
	0: "NONE",
	1: "WORKER",
	2: "MASTER",
}
var MinionConfig_Role_value = map[string]int32{
	"NONE":   0,
	"WORKER": 1,
	"MASTER": 2,
}

func (x MinionConfig_Role) String() string {
	return proto.EnumName(MinionConfig_Role_name, int32(x))
}
func (MinionConfig_Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type MinionConfig struct {
	ID        string            `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Role      MinionConfig_Role `protobuf:"varint,2,opt,name=role,enum=MinionConfig_Role" json:"role,omitempty"`
	PrivateIP string            `protobuf:"bytes,3,opt,name=PrivateIP,json=privateIP" json:"PrivateIP,omitempty"`
	Spec      string            `protobuf:"bytes,4,opt,name=Spec,json=spec" json:"Spec,omitempty"`
	Provider  string            `protobuf:"bytes,5,opt,name=Provider,json=provider" json:"Provider,omitempty"`
	Size      string            `protobuf:"bytes,6,opt,name=Size,json=size" json:"Size,omitempty"`
	Region    string            `protobuf:"bytes,7,opt,name=Region,json=region" json:"Region,omitempty"`
	PublicIP  string            `protobuf:"bytes,8,opt,name=PublicIP,json=publicIP" json:"PublicIP,omitempty"`
}

func (m *MinionConfig) Reset()                    { *m = MinionConfig{} }
func (m *MinionConfig) String() string            { return proto.CompactTextString(m) }
func (*MinionConfig) ProtoMessage()               {}
func (*MinionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Reply struct {
	Success bool   `protobuf:"varint,1,opt,name=Success,json=success" json:"Success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=Error,json=error" json:"Error,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Request struct {
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type EtcdMembers struct {
	IPs []string `protobuf:"bytes,1,rep,name=IPs,json=iPs" json:"IPs,omitempty"`
}

func (m *EtcdMembers) Reset()                    { *m = EtcdMembers{} }
func (m *EtcdMembers) String() string            { return proto.CompactTextString(m) }
func (*EtcdMembers) ProtoMessage()               {}
func (*EtcdMembers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*MinionConfig)(nil), "MinionConfig")
	proto.RegisterType((*Reply)(nil), "Reply")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*EtcdMembers)(nil), "EtcdMembers")
	proto.RegisterEnum("MinionConfig_Role", MinionConfig_Role_name, MinionConfig_Role_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Minion service

type MinionClient interface {
	SetMinionConfig(ctx context.Context, in *MinionConfig, opts ...grpc.CallOption) (*Reply, error)
	GetMinionConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*MinionConfig, error)
	BootEtcd(ctx context.Context, in *EtcdMembers, opts ...grpc.CallOption) (*Reply, error)
}

type minionClient struct {
	cc *grpc.ClientConn
}

func NewMinionClient(cc *grpc.ClientConn) MinionClient {
	return &minionClient{cc}
}

func (c *minionClient) SetMinionConfig(ctx context.Context, in *MinionConfig, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/Minion/SetMinionConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minionClient) GetMinionConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*MinionConfig, error) {
	out := new(MinionConfig)
	err := grpc.Invoke(ctx, "/Minion/GetMinionConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minionClient) BootEtcd(ctx context.Context, in *EtcdMembers, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/Minion/BootEtcd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Minion service

type MinionServer interface {
	SetMinionConfig(context.Context, *MinionConfig) (*Reply, error)
	GetMinionConfig(context.Context, *Request) (*MinionConfig, error)
	BootEtcd(context.Context, *EtcdMembers) (*Reply, error)
}

func RegisterMinionServer(s *grpc.Server, srv MinionServer) {
	s.RegisterService(&_Minion_serviceDesc, srv)
}

func _Minion_SetMinionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MinionConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(MinionServer).SetMinionConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Minion_GetMinionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(MinionServer).GetMinionConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Minion_BootEtcd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(EtcdMembers)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(MinionServer).BootEtcd(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Minion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Minion",
	HandlerType: (*MinionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMinionConfig",
			Handler:    _Minion_SetMinionConfig_Handler,
		},
		{
			MethodName: "GetMinionConfig",
			Handler:    _Minion_GetMinionConfig_Handler,
		},
		{
			MethodName: "BootEtcd",
			Handler:    _Minion_BootEtcd_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x6c, 0xbe, 0x93, 0xd7, 0xda, 0x96, 0x87, 0xc8, 0x52, 0x04, 0x25, 0x07, 0x29, 0x22, 0x11,
	0xea, 0xc1, 0xb3, 0xda, 0x20, 0x45, 0xda, 0x86, 0x8d, 0xe0, 0xd9, 0xa4, 0x6b, 0x59, 0xa8, 0xdd,
	0xb8, 0x49, 0x0b, 0xfa, 0x03, 0xfa, 0xbb, 0xdd, 0xac, 0x11, 0x5b, 0x6f, 0x6f, 0xde, 0x9b, 0x9d,
	0x19, 0x66, 0xa1, 0x5d, 0x64, 0xd7, 0x45, 0x16, 0x15, 0x52, 0x54, 0x22, 0xdc, 0x99, 0xd0, 0x99,
	0xf2, 0x35, 0x17, 0xeb, 0x07, 0xb1, 0x7e, 0xe3, 0x4b, 0xec, 0x82, 0x39, 0x19, 0x13, 0xe3, 0xdc,
	0x18, 0x06, 0xd4, 0xe4, 0x63, 0xbc, 0x00, 0x5b, 0x8a, 0x15, 0x23, 0xa6, 0xda, 0x74, 0x47, 0x18,
	0xed, 0x93, 0x23, 0xaa, 0x2e, 0x54, 0xdf, 0xf1, 0x14, 0x82, 0x44, 0xf2, 0xed, 0x6b, 0xc5, 0x26,
	0x09, 0xb1, 0xf4, 0xf3, 0xa0, 0xf8, 0x5d, 0x20, 0x82, 0x9d, 0x16, 0x2c, 0x27, 0xb6, 0x3e, 0xd8,
	0xa5, 0x9a, 0x71, 0x00, 0x7e, 0x22, 0xc5, 0x96, 0x2f, 0x98, 0x24, 0x8e, 0xde, 0xfb, 0x45, 0x83,
	0x35, 0x9f, 0x7f, 0x31, 0xe2, 0x36, 0x7c, 0x35, 0xe3, 0x09, 0xb8, 0x94, 0x2d, 0x95, 0x39, 0xf1,
	0xf4, 0xd6, 0x95, 0x1a, 0x69, 0x9d, 0x4d, 0xb6, 0xe2, 0xb9, 0x32, 0xf6, 0x1b, 0x9d, 0x06, 0x87,
	0x43, 0xb0, 0xeb, 0x8c, 0xe8, 0x83, 0x3d, 0x9b, 0xcf, 0xe2, 0x7e, 0x0b, 0x01, 0xdc, 0x97, 0x39,
	0x7d, 0x8a, 0x69, 0xdf, 0xa8, 0xe7, 0xe9, 0x5d, 0xfa, 0xac, 0x66, 0x33, 0xbc, 0x05, 0x87, 0xb2,
	0x62, 0xf5, 0x89, 0x04, 0xbc, 0x74, 0x93, 0xe7, 0xac, 0x2c, 0x75, 0x0b, 0x3e, 0xf5, 0xca, 0x1f,
	0x88, 0xc7, 0xe0, 0xc4, 0x52, 0x0a, 0xa9, 0xbb, 0x08, 0xa8, 0xc3, 0x6a, 0x10, 0x06, 0xe0, 0x51,
	0xf6, 0xb1, 0x61, 0x65, 0x15, 0x9e, 0x41, 0x3b, 0xae, 0xf2, 0xc5, 0x94, 0xbd, 0x67, 0x4c, 0x96,
	0xd8, 0x07, 0x6b, 0x92, 0xd4, 0x2a, 0x96, 0x62, 0x5b, 0x3c, 0x29, 0x47, 0x3b, 0x43, 0x39, 0xea,
	0x02, 0xf1, 0x12, 0x7a, 0x29, 0xab, 0x0e, 0xaa, 0x3f, 0x3a, 0x28, 0x77, 0xe0, 0x46, 0x3a, 0x50,
	0xd8, 0xc2, 0x2b, 0xe8, 0x3d, 0xfe, 0xe3, 0xfa, 0x51, 0x63, 0x3a, 0x38, 0x7c, 0xa5, 0xd8, 0x21,
	0xf8, 0xf7, 0x42, 0x54, 0x75, 0x12, 0xec, 0x44, 0x7b, 0x81, 0xfe, 0x14, 0x33, 0x57, 0xff, 0xfe,
	0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xf3, 0xf3, 0x80, 0x0c, 0x02, 0x00, 0x00,
}
