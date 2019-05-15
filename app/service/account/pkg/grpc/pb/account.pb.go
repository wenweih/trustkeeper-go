// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_9a7c9f519f921888, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReply) Reset()         { *m = CreateReply{} }
func (m *CreateReply) String() string { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()    {}
func (*CreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_9a7c9f519f921888, []int{1}
}
func (m *CreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReply.Unmarshal(m, b)
}
func (m *CreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReply.Marshal(b, m, deterministic)
}
func (dst *CreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReply.Merge(dst, src)
}
func (m *CreateReply) XXX_Size() int {
	return xxx_messageInfo_CreateReply.Size(m)
}
func (m *CreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReply proto.InternalMessageInfo

func (m *CreateReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type SigninRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigninRequest) Reset()         { *m = SigninRequest{} }
func (m *SigninRequest) String() string { return proto.CompactTextString(m) }
func (*SigninRequest) ProtoMessage()    {}
func (*SigninRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_9a7c9f519f921888, []int{2}
}
func (m *SigninRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigninRequest.Unmarshal(m, b)
}
func (m *SigninRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigninRequest.Marshal(b, m, deterministic)
}
func (dst *SigninRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigninRequest.Merge(dst, src)
}
func (m *SigninRequest) XXX_Size() int {
	return xxx_messageInfo_SigninRequest.Size(m)
}
func (m *SigninRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SigninRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SigninRequest proto.InternalMessageInfo

func (m *SigninRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SigninRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SigninReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SigninReply) Reset()         { *m = SigninReply{} }
func (m *SigninReply) String() string { return proto.CompactTextString(m) }
func (*SigninReply) ProtoMessage()    {}
func (*SigninReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_9a7c9f519f921888, []int{3}
}
func (m *SigninReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SigninReply.Unmarshal(m, b)
}
func (m *SigninReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SigninReply.Marshal(b, m, deterministic)
}
func (dst *SigninReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigninReply.Merge(dst, src)
}
func (m *SigninReply) XXX_Size() int {
	return xxx_messageInfo_SigninReply.Size(m)
}
func (m *SigninReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SigninReply.DiscardUnknown(m)
}

var xxx_messageInfo_SigninReply proto.InternalMessageInfo

func (m *SigninReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignoutRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignoutRequest) Reset()         { *m = SignoutRequest{} }
func (m *SignoutRequest) String() string { return proto.CompactTextString(m) }
func (*SignoutRequest) ProtoMessage()    {}
func (*SignoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_9a7c9f519f921888, []int{4}
}
func (m *SignoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignoutRequest.Unmarshal(m, b)
}
func (m *SignoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignoutRequest.Marshal(b, m, deterministic)
}
func (dst *SignoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignoutRequest.Merge(dst, src)
}
func (m *SignoutRequest) XXX_Size() int {
	return xxx_messageInfo_SignoutRequest.Size(m)
}
func (m *SignoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignoutRequest proto.InternalMessageInfo

func (m *SignoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SignoutReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignoutReply) Reset()         { *m = SignoutReply{} }
func (m *SignoutReply) String() string { return proto.CompactTextString(m) }
func (*SignoutReply) ProtoMessage()    {}
func (*SignoutReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_9a7c9f519f921888, []int{5}
}
func (m *SignoutReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignoutReply.Unmarshal(m, b)
}
func (m *SignoutReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignoutReply.Marshal(b, m, deterministic)
}
func (dst *SignoutReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignoutReply.Merge(dst, src)
}
func (m *SignoutReply) XXX_Size() int {
	return xxx_messageInfo_SignoutReply.Size(m)
}
func (m *SignoutReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignoutReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignoutReply proto.InternalMessageInfo

func (m *SignoutReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "pb.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "pb.CreateReply")
	proto.RegisterType((*SigninRequest)(nil), "pb.SigninRequest")
	proto.RegisterType((*SigninReply)(nil), "pb.SigninReply")
	proto.RegisterType((*SignoutRequest)(nil), "pb.SignoutRequest")
	proto.RegisterType((*SignoutReply)(nil), "pb.SignoutReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error)
	Signin(ctx context.Context, in *SigninRequest, opts ...grpc.CallOption) (*SigninReply, error)
	Signout(ctx context.Context, in *SignoutRequest, opts ...grpc.CallOption) (*SignoutReply, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/pb.Account/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Signin(ctx context.Context, in *SigninRequest, opts ...grpc.CallOption) (*SigninReply, error) {
	out := new(SigninReply)
	err := c.cc.Invoke(ctx, "/pb.Account/Signin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Signout(ctx context.Context, in *SignoutRequest, opts ...grpc.CallOption) (*SignoutReply, error) {
	out := new(SignoutReply)
	err := c.cc.Invoke(ctx, "/pb.Account/Signout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	Create(context.Context, *CreateRequest) (*CreateReply, error)
	Signin(context.Context, *SigninRequest) (*SigninReply, error)
	Signout(context.Context, *SignoutRequest) (*SignoutReply, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Signin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigninRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Signin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Signin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Signin(ctx, req.(*SigninRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Signout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Signout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Signout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Signout(ctx, req.(*SignoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Account_Create_Handler,
		},
		{
			MethodName: "Signin",
			Handler:    _Account_Signin_Handler,
		},
		{
			MethodName: "Signout",
			Handler:    _Account_Signout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_account_9a7c9f519f921888) }

var fileDescriptor_account_9a7c9f519f921888 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe4,
	0xe2, 0x75, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe1, 0x62, 0x4d, 0xcd, 0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70,
	0x84, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24, 0x98, 0xc0, 0x12,
	0x70, 0xbe, 0x92, 0x2a, 0x17, 0x37, 0xcc, 0x88, 0x82, 0x9c, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xa2,
	0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xb0, 0x09, 0x1c, 0x41, 0x50, 0x1e, 0xc8, 0xa6, 0xe0, 0xcc, 0xf4,
	0xbc, 0xcc, 0x3c, 0xf2, 0x6d, 0x52, 0xe6, 0xe2, 0x86, 0x19, 0x01, 0xb2, 0x49, 0x84, 0x8b, 0xb5,
	0x24, 0x3f, 0x3b, 0x35, 0x0f, 0x66, 0x00, 0x98, 0xa3, 0xa4, 0xc6, 0xc5, 0x07, 0x52, 0x94, 0x5f,
	0x5a, 0x82, 0x64, 0x11, 0x56, 0x75, 0x3c, 0x70, 0x75, 0x78, 0xdc, 0x6d, 0x34, 0x8d, 0x91, 0x8b,
	0xdd, 0x11, 0x12, 0x6e, 0x42, 0x3a, 0x5c, 0x6c, 0x10, 0xaf, 0x0a, 0x09, 0xea, 0x15, 0x24, 0xe9,
	0xa1, 0x84, 0x9c, 0x14, 0x3f, 0xb2, 0x10, 0xc8, 0x44, 0x1d, 0x2e, 0x36, 0x88, 0x73, 0x21, 0xaa,
	0x51, 0x7c, 0x0f, 0x51, 0x8d, 0xec, 0x1b, 0x7d, 0x2e, 0x76, 0xa8, 0x7b, 0x84, 0x84, 0x60, 0x72,
	0x08, 0x4f, 0x48, 0x09, 0xa0, 0x88, 0x15, 0xe4, 0x54, 0x26, 0xb1, 0x81, 0x63, 0xd1, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x1a, 0xe0, 0xb3, 0xc3, 0xd6, 0x01, 0x00, 0x00,
}
