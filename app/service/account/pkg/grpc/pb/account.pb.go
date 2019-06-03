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
	return fileDescriptor_account_fa98bfa66adcea4e, []int{0}
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
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateReply) Reset()         { *m = CreateReply{} }
func (m *CreateReply) String() string { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()    {}
func (*CreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_fa98bfa66adcea4e, []int{1}
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

func (m *CreateReply) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
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
	return fileDescriptor_account_fa98bfa66adcea4e, []int{2}
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
	return fileDescriptor_account_fa98bfa66adcea4e, []int{3}
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
	return fileDescriptor_account_fa98bfa66adcea4e, []int{4}
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
	return fileDescriptor_account_fa98bfa66adcea4e, []int{5}
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

type RolesRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolesRequest) Reset()         { *m = RolesRequest{} }
func (m *RolesRequest) String() string { return proto.CompactTextString(m) }
func (*RolesRequest) ProtoMessage()    {}
func (*RolesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_fa98bfa66adcea4e, []int{6}
}
func (m *RolesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesRequest.Unmarshal(m, b)
}
func (m *RolesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesRequest.Marshal(b, m, deterministic)
}
func (dst *RolesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesRequest.Merge(dst, src)
}
func (m *RolesRequest) XXX_Size() int {
	return xxx_messageInfo_RolesRequest.Size(m)
}
func (m *RolesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RolesRequest proto.InternalMessageInfo

func (m *RolesRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RolesReply struct {
	Roles                []string `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolesReply) Reset()         { *m = RolesReply{} }
func (m *RolesReply) String() string { return proto.CompactTextString(m) }
func (*RolesReply) ProtoMessage()    {}
func (*RolesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_fa98bfa66adcea4e, []int{7}
}
func (m *RolesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolesReply.Unmarshal(m, b)
}
func (m *RolesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolesReply.Marshal(b, m, deterministic)
}
func (dst *RolesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolesReply.Merge(dst, src)
}
func (m *RolesReply) XXX_Size() int {
	return xxx_messageInfo_RolesReply.Size(m)
}
func (m *RolesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RolesReply.DiscardUnknown(m)
}

var xxx_messageInfo_RolesReply proto.InternalMessageInfo

func (m *RolesReply) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type AuthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_fa98bfa66adcea4e, []int{8}
}
func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (dst *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(dst, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

type AuthReply struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReply) Reset()         { *m = AuthReply{} }
func (m *AuthReply) String() string { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()    {}
func (*AuthReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_fa98bfa66adcea4e, []int{9}
}
func (m *AuthReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthReply.Unmarshal(m, b)
}
func (m *AuthReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthReply.Marshal(b, m, deterministic)
}
func (dst *AuthReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthReply.Merge(dst, src)
}
func (m *AuthReply) XXX_Size() int {
	return xxx_messageInfo_AuthReply.Size(m)
}
func (m *AuthReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthReply proto.InternalMessageInfo

func (m *AuthReply) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "pb.CreateRequest")
	proto.RegisterType((*CreateReply)(nil), "pb.CreateReply")
	proto.RegisterType((*SigninRequest)(nil), "pb.SigninRequest")
	proto.RegisterType((*SigninReply)(nil), "pb.SigninReply")
	proto.RegisterType((*SignoutRequest)(nil), "pb.SignoutRequest")
	proto.RegisterType((*SignoutReply)(nil), "pb.SignoutReply")
	proto.RegisterType((*RolesRequest)(nil), "pb.RolesRequest")
	proto.RegisterType((*RolesReply)(nil), "pb.RolesReply")
	proto.RegisterType((*AuthRequest)(nil), "pb.AuthRequest")
	proto.RegisterType((*AuthReply)(nil), "pb.AuthReply")
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
	Roles(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*RolesReply, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
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

func (c *accountClient) Roles(ctx context.Context, in *RolesRequest, opts ...grpc.CallOption) (*RolesReply, error) {
	out := new(RolesReply)
	err := c.cc.Invoke(ctx, "/pb.Account/Roles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/pb.Account/Auth", in, out, opts...)
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
	Roles(context.Context, *RolesRequest) (*RolesReply, error)
	Auth(context.Context, *AuthRequest) (*AuthReply, error)
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

func _Account_Roles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Roles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Roles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Roles(ctx, req.(*RolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Auth(ctx, req.(*AuthRequest))
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
		{
			MethodName: "Roles",
			Handler:    _Account_Roles_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _Account_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_account_fa98bfa66adcea4e) }

var fileDescriptor_account_fa98bfa66adcea4e = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x4e, 0x83, 0x30,
	0x14, 0xc6, 0xc3, 0x1c, 0x6c, 0x3b, 0x8c, 0xa9, 0x27, 0x8b, 0x21, 0xdc, 0x38, 0xab, 0x59, 0x34,
	0x59, 0x30, 0xd1, 0x27, 0x20, 0xbe, 0x01, 0x3e, 0x01, 0x6c, 0x8d, 0x12, 0x91, 0x56, 0x68, 0x63,
	0xf6, 0xc8, 0xbe, 0x85, 0xe9, 0x1f, 0xb6, 0x92, 0xa8, 0x17, 0xde, 0xf1, 0x9d, 0xf3, 0xeb, 0x77,
	0x38, 0x5f, 0x0b, 0x51, 0xb1, 0xdd, 0x32, 0xd9, 0x88, 0x94, 0xb7, 0x4c, 0x30, 0x1c, 0xf1, 0x92,
	0x64, 0x10, 0x3d, 0xb5, 0xb4, 0x10, 0x34, 0xa7, 0x1f, 0x92, 0x76, 0x02, 0x97, 0xe0, 0xd3, 0xf7,
	0xa2, 0xaa, 0x63, 0x6f, 0xe5, 0xdd, 0xce, 0x72, 0x23, 0x30, 0x81, 0x29, 0x2f, 0xba, 0xee, 0x93,
	0xb5, 0xbb, 0x78, 0xa4, 0x1b, 0x07, 0x4d, 0xae, 0x20, 0xec, 0x2d, 0x78, 0xbd, 0x47, 0x84, 0xb1,
	0x94, 0xd5, 0xce, 0x9e, 0xd7, 0xdf, 0x6a, 0xca, 0x73, 0xf5, 0xd2, 0x54, 0xcd, 0xff, 0xa7, 0x5c,
	0x43, 0xd8, 0x5b, 0xa8, 0x29, 0x4b, 0xf0, 0x05, 0x7b, 0xa3, 0x4d, 0x6f, 0xa0, 0x05, 0x59, 0xc3,
	0x42, 0x41, 0x4c, 0x0a, 0x67, 0xd0, 0x8f, 0xdc, 0xfc, 0xc0, 0x29, 0xb7, 0x0b, 0x08, 0x5a, 0xda,
	0xc9, 0x5a, 0x68, 0x6c, 0x9a, 0x5b, 0x45, 0x6e, 0x60, 0x9e, 0xb3, 0x9a, 0x76, 0x7f, 0xbb, 0x11,
	0x00, 0x4b, 0xd9, 0x3f, 0x6b, 0x95, 0x8a, 0xbd, 0xd5, 0x89, 0x62, 0xb4, 0x20, 0x11, 0x84, 0x99,
	0x14, 0xaf, 0xd6, 0x88, 0x5c, 0xc2, 0xcc, 0xc8, 0x5f, 0x12, 0x7b, 0xf8, 0xf2, 0x60, 0x92, 0x99,
	0xdb, 0xc2, 0x0d, 0x04, 0x26, 0x60, 0x3c, 0x4f, 0x79, 0x99, 0x0e, 0xee, 0x2b, 0x39, 0x75, 0x4b,
	0xca, 0x6d, 0x03, 0x81, 0x09, 0xca, 0xd0, 0x83, 0xdc, 0x0d, 0xed, 0xe6, 0x78, 0x0f, 0x13, 0x9b,
	0x04, 0x62, 0xdf, 0x3b, 0xc6, 0x97, 0x9c, 0x0d, 0x6a, 0xea, 0xc0, 0x1d, 0xf8, 0x7a, 0x59, 0xd4,
	0x2d, 0x37, 0x9d, 0x64, 0xe1, 0x54, 0x14, 0xba, 0x86, 0xb1, 0x5a, 0x12, 0xf5, 0x50, 0x67, 0xfb,
	0x24, 0x3a, 0x16, 0x78, 0xbd, 0x2f, 0x03, 0xfd, 0x1c, 0x1f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0xbd, 0x9b, 0x4b, 0x9f, 0x02, 0x00, 0x00,
}
