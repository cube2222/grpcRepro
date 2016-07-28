// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Member
	MemberList
*/
package message

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Member struct {
	Address   string `protobuf:"bytes,1,opt,name=Address,json=address" json:"Address,omitempty"`
	Heartbeat int64  `protobuf:"varint,2,opt,name=Heartbeat,json=heartbeat" json:"Heartbeat,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
	Alive     bool   `protobuf:"varint,4,opt,name=Alive,json=alive" json:"Alive,omitempty"`
}

func (m *Member) Reset()                    { *m = Member{} }
func (m *Member) String() string            { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()               {}
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MemberList struct {
	List map[string]*Member `protobuf:"bytes,1,rep,name=list" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MemberList) Reset()                    { *m = MemberList{} }
func (m *MemberList) String() string            { return proto.CompactTextString(m) }
func (*MemberList) ProtoMessage()               {}
func (*MemberList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MemberList) GetList() map[string]*Member {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Member)(nil), "message.Member")
	proto.RegisterType((*MemberList)(nil), "message.MemberList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MyService service

type MyServiceClient interface {
	GetResponse(ctx context.Context, in *Member, opts ...grpc.CallOption) (*MemberList, error)
}

type myServiceClient struct {
	cc *grpc.ClientConn
}

func NewMyServiceClient(cc *grpc.ClientConn) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) GetResponse(ctx context.Context, in *Member, opts ...grpc.CallOption) (*MemberList, error) {
	out := new(MemberList)
	err := grpc.Invoke(ctx, "/message.myService/GetResponse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyService service

type MyServiceServer interface {
	GetResponse(context.Context, *Member) (*MemberList, error)
}

func RegisterMyServiceServer(s *grpc.Server, srv MyServiceServer) {
	s.RegisterService(&_MyService_serviceDesc, srv)
}

func _MyService_GetResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServiceServer).GetResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.myService/GetResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServiceServer).GetResponse(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.myService",
	HandlerType: (*MyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResponse",
			Handler:    _MyService_GetResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0xcd, 0x76, 0xbb, 0x35, 0x53, 0x44, 0x89, 0x1e, 0xc2, 0xa2, 0x50, 0x0a, 0x42, 0x4f,
	0x05, 0x2b, 0x82, 0x78, 0x5b, 0x41, 0xdc, 0x83, 0x5e, 0xa2, 0x2f, 0x90, 0xba, 0x83, 0x06, 0x9b,
	0x6d, 0x49, 0xc6, 0x42, 0x1f, 0xc1, 0xb7, 0x96, 0xb6, 0x5b, 0x85, 0xc5, 0x4b, 0xc8, 0xff, 0x23,
	0x99, 0x1f, 0x03, 0x47, 0x16, 0xbd, 0xd7, 0xef, 0x98, 0x37, 0xae, 0xa6, 0x5a, 0x44, 0x3b, 0x99,
	0x12, 0x2c, 0x9e, 0xd1, 0x96, 0xe8, 0x84, 0x84, 0x68, 0xb5, 0xd9, 0x38, 0xf4, 0x5e, 0xb2, 0x84,
	0x65, 0x5c, 0x45, 0x7a, 0x94, 0xe2, 0x1c, 0xf8, 0x1a, 0xb5, 0xa3, 0x12, 0x35, 0xc9, 0x59, 0xc2,
	0xb2, 0x40, 0xf1, 0x8f, 0xc9, 0xe8, 0xd3, 0x57, 0x63, 0xd1, 0x93, 0xb6, 0x8d, 0x0c, 0xc6, 0x94,
	0x26, 0x43, 0x9c, 0x41, 0xb8, 0xaa, 0x4c, 0x8b, 0x72, 0x9e, 0xb0, 0xec, 0x50, 0x85, 0xba, 0x17,
	0xe9, 0x37, 0x03, 0x18, 0xc7, 0x3e, 0x19, 0x4f, 0xe2, 0x0a, 0xe6, 0x95, 0xf1, 0x24, 0x59, 0x12,
	0x64, 0x71, 0x71, 0x91, 0x4f, 0xac, 0x7f, 0x95, 0xbc, 0x3f, 0x1e, 0xb6, 0xe4, 0x3a, 0x35, 0x54,
	0x97, 0x6b, 0xe0, 0xbf, 0x96, 0x38, 0x81, 0xe0, 0x13, 0xbb, 0x1d, 0x76, 0x7f, 0x15, 0x97, 0x10,
	0xb6, 0xba, 0xfa, 0xc2, 0x01, 0x37, 0x2e, 0x8e, 0xf7, 0xbe, 0x54, 0x63, 0x7a, 0x37, 0xbb, 0x65,
	0xc5, 0x3d, 0x70, 0xdb, 0xbd, 0xa0, 0x6b, 0xcd, 0x1b, 0x8a, 0x1b, 0x88, 0x1f, 0x91, 0x14, 0xfa,
	0xa6, 0xde, 0x7a, 0x14, 0xfb, 0xef, 0x96, 0xa7, 0xff, 0xb0, 0xa5, 0x07, 0xe5, 0x62, 0xd8, 0xea,
	0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0xfa, 0xfb, 0x95, 0x66, 0x01, 0x00, 0x00,
}