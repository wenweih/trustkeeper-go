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
	Orgname              string   `protobuf:"bytes,3,opt,name=orgname" json:"orgname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_c3be93700ef927f8, []int{0}
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

func (m *CreateRequest) GetOrgname() string {
	if m != nil {
		return m.Orgname
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
	return fileDescriptor_account_c3be93700ef927f8, []int{1}
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
	return fileDescriptor_account_c3be93700ef927f8, []int{2}
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
	return fileDescriptor_account_c3be93700ef927f8, []int{3}
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
	return fileDescriptor_account_c3be93700ef927f8, []int{4}
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
	return fileDescriptor_account_c3be93700ef927f8, []int{5}
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
	return fileDescriptor_account_c3be93700ef927f8, []int{6}
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
	return fileDescriptor_account_c3be93700ef927f8, []int{7}
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
	return fileDescriptor_account_c3be93700ef927f8, []int{8}
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
	NamespaceID          uint32   `protobuf:"varint,2,opt,name=namespaceID" json:"namespaceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthReply) Reset()         { *m = AuthReply{} }
func (m *AuthReply) String() string { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()    {}
func (*AuthReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_c3be93700ef927f8, []int{9}
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

func (m *AuthReply) GetNamespaceID() uint32 {
	if m != nil {
		return m.NamespaceID
	}
	return 0
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

func init() { proto.RegisterFile("account.proto", fileDescriptor_account_c3be93700ef927f8) }

var fileDescriptor_account_c3be93700ef927f8 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x49, 0x3f, 0xd2, 0xf6, 0xa6, 0xe9, 0xff, 0xef, 0xa5, 0x48, 0xc8, 0xaa, 0x8e, 0x52,
	0x14, 0x4a, 0x05, 0x7d, 0x82, 0xa0, 0x1b, 0xb7, 0x71, 0xe9, 0x2a, 0x6d, 0x87, 0x1a, 0x4c, 0x33,
	0x63, 0x66, 0x06, 0xe9, 0x23, 0xfb, 0x16, 0x32, 0x1f, 0x69, 0xa7, 0xa0, 0x2e, 0xdc, 0xe5, 0xdc,
	0x7b, 0x72, 0x4e, 0xf8, 0xdd, 0x40, 0x5c, 0xac, 0xd7, 0x4c, 0xd5, 0x72, 0xc9, 0x1b, 0x26, 0x19,
	0x76, 0xf8, 0x8a, 0xbc, 0x40, 0xfc, 0xd0, 0xd0, 0x42, 0xd2, 0x9c, 0xbe, 0x2b, 0x2a, 0x24, 0x4e,
	0xa1, 0x4f, 0x77, 0x45, 0x59, 0x25, 0xc1, 0x2c, 0xb8, 0x1e, 0xe5, 0x56, 0x60, 0x0a, 0x43, 0x5e,
	0x08, 0xf1, 0xc1, 0x9a, 0x4d, 0xd2, 0x31, 0x8b, 0x83, 0xc6, 0x04, 0x06, 0xac, 0xd9, 0xd6, 0xc5,
	0x8e, 0x26, 0x5d, 0xb3, 0x6a, 0x25, 0xb9, 0x80, 0xa8, 0x0d, 0xe7, 0xd5, 0x1e, 0x11, 0x7a, 0x4a,
	0x95, 0x1b, 0x97, 0x6c, 0x9e, 0x49, 0x06, 0xf1, 0x73, 0xb9, 0xad, 0xcb, 0xfa, 0xcf, 0xfd, 0xe4,
	0x12, 0xa2, 0x36, 0x42, 0xb7, 0x4c, 0xa1, 0x2f, 0xd9, 0x1b, 0xad, 0xdb, 0x00, 0x23, 0xc8, 0x1c,
	0x26, 0xda, 0xc4, 0x94, 0xf4, 0x8a, 0xbe, 0xf5, 0x8d, 0x0f, 0x3e, 0x9d, 0x76, 0x0e, 0x61, 0x43,
	0x85, 0xaa, 0xa4, 0xb1, 0x0d, 0x73, 0xa7, 0xc8, 0x15, 0x8c, 0x73, 0x56, 0x51, 0xf1, 0x7b, 0x1a,
	0x01, 0x70, 0x2e, 0xf7, 0x65, 0x8d, 0x56, 0x49, 0x30, 0xeb, 0x6a, 0x8f, 0x11, 0x24, 0x86, 0x28,
	0x53, 0xf2, 0xd5, 0x05, 0x91, 0x0c, 0x46, 0x56, 0xfe, 0x40, 0x0c, 0x67, 0x10, 0x69, 0xb8, 0x82,
	0x17, 0x6b, 0xfa, 0xf4, 0x68, 0x68, 0xc4, 0xb9, 0x3f, 0xba, 0xfb, 0x0c, 0x60, 0x90, 0xd9, 0x4b,
	0xe3, 0x02, 0x42, 0x7b, 0x02, 0x3c, 0x5b, 0xf2, 0xd5, 0xf2, 0xe4, 0xd6, 0xe9, 0x3f, 0x7f, 0xa4,
	0xfb, 0x16, 0x10, 0x5a, 0x94, 0xd6, 0x7d, 0x72, 0x19, 0xeb, 0xf6, 0x49, 0xdf, 0xc2, 0xc0, 0xb1,
	0x42, 0x6c, 0x77, 0x47, 0xc0, 0xe9, 0xff, 0x93, 0x99, 0x7e, 0xe1, 0x06, 0xfa, 0x06, 0x07, 0x9a,
	0x95, 0xcf, 0x2f, 0x9d, 0x78, 0x13, 0x6d, 0x9d, 0x43, 0x4f, 0x63, 0x40, 0x53, 0xea, 0xf1, 0x49,
	0xe3, 0xe3, 0x80, 0x57, 0xfb, 0x55, 0x68, 0x7e, 0xe5, 0xfb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x33, 0x31, 0x14, 0x5a, 0xdb, 0x02, 0x00, 0x00,
}
