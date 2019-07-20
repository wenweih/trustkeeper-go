// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet_management.proto

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

type CreateChainRequest struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Bitid                string   `protobuf:"bytes,2,opt,name=bitid" json:"bitid,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChainRequest) Reset()         { *m = CreateChainRequest{} }
func (m *CreateChainRequest) String() string { return proto.CompactTextString(m) }
func (*CreateChainRequest) ProtoMessage()    {}
func (*CreateChainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{0}
}
func (m *CreateChainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChainRequest.Unmarshal(m, b)
}
func (m *CreateChainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChainRequest.Marshal(b, m, deterministic)
}
func (dst *CreateChainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChainRequest.Merge(dst, src)
}
func (m *CreateChainRequest) XXX_Size() int {
	return xxx_messageInfo_CreateChainRequest.Size(m)
}
func (m *CreateChainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChainRequest proto.InternalMessageInfo

func (m *CreateChainRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *CreateChainRequest) GetBitid() string {
	if m != nil {
		return m.Bitid
	}
	return ""
}

func (m *CreateChainRequest) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type CreateChainReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChainReply) Reset()         { *m = CreateChainReply{} }
func (m *CreateChainReply) String() string { return proto.CompactTextString(m) }
func (*CreateChainReply) ProtoMessage()    {}
func (*CreateChainReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{1}
}
func (m *CreateChainReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChainReply.Unmarshal(m, b)
}
func (m *CreateChainReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChainReply.Marshal(b, m, deterministic)
}
func (dst *CreateChainReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChainReply.Merge(dst, src)
}
func (m *CreateChainReply) XXX_Size() int {
	return xxx_messageInfo_CreateChainReply.Size(m)
}
func (m *CreateChainReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChainReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChainReply proto.InternalMessageInfo

func (m *CreateChainReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type AssignedXpubToGroupRequest struct {
	Groupid              string   `protobuf:"bytes,1,opt,name=groupid" json:"groupid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignedXpubToGroupRequest) Reset()         { *m = AssignedXpubToGroupRequest{} }
func (m *AssignedXpubToGroupRequest) String() string { return proto.CompactTextString(m) }
func (*AssignedXpubToGroupRequest) ProtoMessage()    {}
func (*AssignedXpubToGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{2}
}
func (m *AssignedXpubToGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignedXpubToGroupRequest.Unmarshal(m, b)
}
func (m *AssignedXpubToGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignedXpubToGroupRequest.Marshal(b, m, deterministic)
}
func (dst *AssignedXpubToGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignedXpubToGroupRequest.Merge(dst, src)
}
func (m *AssignedXpubToGroupRequest) XXX_Size() int {
	return xxx_messageInfo_AssignedXpubToGroupRequest.Size(m)
}
func (m *AssignedXpubToGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignedXpubToGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignedXpubToGroupRequest proto.InternalMessageInfo

func (m *AssignedXpubToGroupRequest) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

type AssignedXpubToGroupReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignedXpubToGroupReply) Reset()         { *m = AssignedXpubToGroupReply{} }
func (m *AssignedXpubToGroupReply) String() string { return proto.CompactTextString(m) }
func (*AssignedXpubToGroupReply) ProtoMessage()    {}
func (*AssignedXpubToGroupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{3}
}
func (m *AssignedXpubToGroupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignedXpubToGroupReply.Unmarshal(m, b)
}
func (m *AssignedXpubToGroupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignedXpubToGroupReply.Marshal(b, m, deterministic)
}
func (dst *AssignedXpubToGroupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignedXpubToGroupReply.Merge(dst, src)
}
func (m *AssignedXpubToGroupReply) XXX_Size() int {
	return xxx_messageInfo_AssignedXpubToGroupReply.Size(m)
}
func (m *AssignedXpubToGroupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignedXpubToGroupReply.DiscardUnknown(m)
}

var xxx_messageInfo_AssignedXpubToGroupReply proto.InternalMessageInfo

type GetChainsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetChainsRequest) Reset()         { *m = GetChainsRequest{} }
func (m *GetChainsRequest) String() string { return proto.CompactTextString(m) }
func (*GetChainsRequest) ProtoMessage()    {}
func (*GetChainsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{4}
}
func (m *GetChainsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChainsRequest.Unmarshal(m, b)
}
func (m *GetChainsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChainsRequest.Marshal(b, m, deterministic)
}
func (dst *GetChainsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChainsRequest.Merge(dst, src)
}
func (m *GetChainsRequest) XXX_Size() int {
	return xxx_messageInfo_GetChainsRequest.Size(m)
}
func (m *GetChainsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChainsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChainsRequest proto.InternalMessageInfo

type SimpleChain struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Bip44Id              int32    `protobuf:"varint,3,opt,name=bip44id" json:"bip44id,omitempty"`
	Coin                 string   `protobuf:"bytes,4,opt,name=coin" json:"coin,omitempty"`
	Status               bool     `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleChain) Reset()         { *m = SimpleChain{} }
func (m *SimpleChain) String() string { return proto.CompactTextString(m) }
func (*SimpleChain) ProtoMessage()    {}
func (*SimpleChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{5}
}
func (m *SimpleChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleChain.Unmarshal(m, b)
}
func (m *SimpleChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleChain.Marshal(b, m, deterministic)
}
func (dst *SimpleChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleChain.Merge(dst, src)
}
func (m *SimpleChain) XXX_Size() int {
	return xxx_messageInfo_SimpleChain.Size(m)
}
func (m *SimpleChain) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleChain.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleChain proto.InternalMessageInfo

func (m *SimpleChain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SimpleChain) GetBip44Id() int32 {
	if m != nil {
		return m.Bip44Id
	}
	return 0
}

func (m *SimpleChain) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

func (m *SimpleChain) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type GetChainsReply struct {
	Chains               []*SimpleChain `protobuf:"bytes,1,rep,name=chains" json:"chains,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetChainsReply) Reset()         { *m = GetChainsReply{} }
func (m *GetChainsReply) String() string { return proto.CompactTextString(m) }
func (*GetChainsReply) ProtoMessage()    {}
func (*GetChainsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{6}
}
func (m *GetChainsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChainsReply.Unmarshal(m, b)
}
func (m *GetChainsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChainsReply.Marshal(b, m, deterministic)
}
func (dst *GetChainsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChainsReply.Merge(dst, src)
}
func (m *GetChainsReply) XXX_Size() int {
	return xxx_messageInfo_GetChainsReply.Size(m)
}
func (m *GetChainsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChainsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetChainsReply proto.InternalMessageInfo

func (m *GetChainsReply) GetChains() []*SimpleChain {
	if m != nil {
		return m.Chains
	}
	return nil
}

type CreateWalletRequest struct {
	Groupid              string   `protobuf:"bytes,1,opt,name=groupid" json:"groupid,omitempty"`
	Chainname            string   `protobuf:"bytes,2,opt,name=chainname" json:"chainname,omitempty"`
	Bip44Change          int32    `protobuf:"varint,3,opt,name=bip44change" json:"bip44change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWalletRequest) Reset()         { *m = CreateWalletRequest{} }
func (m *CreateWalletRequest) String() string { return proto.CompactTextString(m) }
func (*CreateWalletRequest) ProtoMessage()    {}
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{7}
}
func (m *CreateWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWalletRequest.Unmarshal(m, b)
}
func (m *CreateWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWalletRequest.Marshal(b, m, deterministic)
}
func (dst *CreateWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWalletRequest.Merge(dst, src)
}
func (m *CreateWalletRequest) XXX_Size() int {
	return xxx_messageInfo_CreateWalletRequest.Size(m)
}
func (m *CreateWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWalletRequest proto.InternalMessageInfo

func (m *CreateWalletRequest) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

func (m *CreateWalletRequest) GetChainname() string {
	if m != nil {
		return m.Chainname
	}
	return ""
}

func (m *CreateWalletRequest) GetBip44Change() int32 {
	if m != nil {
		return m.Bip44Change
	}
	return 0
}

type Wallet struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address" json:"Address,omitempty"`
	ChainName            string   `protobuf:"bytes,3,opt,name=ChainName" json:"ChainName,omitempty"`
	Status               bool     `protobuf:"varint,4,opt,name=Status" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Wallet) Reset()         { *m = Wallet{} }
func (m *Wallet) String() string { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()    {}
func (*Wallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{8}
}
func (m *Wallet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wallet.Unmarshal(m, b)
}
func (m *Wallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wallet.Marshal(b, m, deterministic)
}
func (dst *Wallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wallet.Merge(dst, src)
}
func (m *Wallet) XXX_Size() int {
	return xxx_messageInfo_Wallet.Size(m)
}
func (m *Wallet) XXX_DiscardUnknown() {
	xxx_messageInfo_Wallet.DiscardUnknown(m)
}

var xxx_messageInfo_Wallet proto.InternalMessageInfo

func (m *Wallet) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Wallet) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Wallet) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *Wallet) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type CreateWalletReply struct {
	Wallet               *Wallet  `protobuf:"bytes,1,opt,name=wallet" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateWalletReply) Reset()         { *m = CreateWalletReply{} }
func (m *CreateWalletReply) String() string { return proto.CompactTextString(m) }
func (*CreateWalletReply) ProtoMessage()    {}
func (*CreateWalletReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_7907d0a21409ef20, []int{9}
}
func (m *CreateWalletReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateWalletReply.Unmarshal(m, b)
}
func (m *CreateWalletReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateWalletReply.Marshal(b, m, deterministic)
}
func (dst *CreateWalletReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateWalletReply.Merge(dst, src)
}
func (m *CreateWalletReply) XXX_Size() int {
	return xxx_messageInfo_CreateWalletReply.Size(m)
}
func (m *CreateWalletReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateWalletReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateWalletReply proto.InternalMessageInfo

func (m *CreateWalletReply) GetWallet() *Wallet {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateChainRequest)(nil), "pb.CreateChainRequest")
	proto.RegisterType((*CreateChainReply)(nil), "pb.CreateChainReply")
	proto.RegisterType((*AssignedXpubToGroupRequest)(nil), "pb.AssignedXpubToGroupRequest")
	proto.RegisterType((*AssignedXpubToGroupReply)(nil), "pb.AssignedXpubToGroupReply")
	proto.RegisterType((*GetChainsRequest)(nil), "pb.GetChainsRequest")
	proto.RegisterType((*SimpleChain)(nil), "pb.SimpleChain")
	proto.RegisterType((*GetChainsReply)(nil), "pb.GetChainsReply")
	proto.RegisterType((*CreateWalletRequest)(nil), "pb.CreateWalletRequest")
	proto.RegisterType((*Wallet)(nil), "pb.Wallet")
	proto.RegisterType((*CreateWalletReply)(nil), "pb.CreateWalletReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletManagementClient is the client API for WalletManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletManagementClient interface {
	CreateChain(ctx context.Context, in *CreateChainRequest, opts ...grpc.CallOption) (*CreateChainReply, error)
	AssignedXpubToGroup(ctx context.Context, in *AssignedXpubToGroupRequest, opts ...grpc.CallOption) (*AssignedXpubToGroupReply, error)
	GetChains(ctx context.Context, in *GetChainsRequest, opts ...grpc.CallOption) (*GetChainsReply, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletReply, error)
}

type walletManagementClient struct {
	cc *grpc.ClientConn
}

func NewWalletManagementClient(cc *grpc.ClientConn) WalletManagementClient {
	return &walletManagementClient{cc}
}

func (c *walletManagementClient) CreateChain(ctx context.Context, in *CreateChainRequest, opts ...grpc.CallOption) (*CreateChainReply, error) {
	out := new(CreateChainReply)
	err := c.cc.Invoke(ctx, "/pb.WalletManagement/CreateChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletManagementClient) AssignedXpubToGroup(ctx context.Context, in *AssignedXpubToGroupRequest, opts ...grpc.CallOption) (*AssignedXpubToGroupReply, error) {
	out := new(AssignedXpubToGroupReply)
	err := c.cc.Invoke(ctx, "/pb.WalletManagement/AssignedXpubToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletManagementClient) GetChains(ctx context.Context, in *GetChainsRequest, opts ...grpc.CallOption) (*GetChainsReply, error) {
	out := new(GetChainsReply)
	err := c.cc.Invoke(ctx, "/pb.WalletManagement/GetChains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletManagementClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletReply, error) {
	out := new(CreateWalletReply)
	err := c.cc.Invoke(ctx, "/pb.WalletManagement/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagementServer is the server API for WalletManagement service.
type WalletManagementServer interface {
	CreateChain(context.Context, *CreateChainRequest) (*CreateChainReply, error)
	AssignedXpubToGroup(context.Context, *AssignedXpubToGroupRequest) (*AssignedXpubToGroupReply, error)
	GetChains(context.Context, *GetChainsRequest) (*GetChainsReply, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletReply, error)
}

func RegisterWalletManagementServer(s *grpc.Server, srv WalletManagementServer) {
	s.RegisterService(&_WalletManagement_serviceDesc, srv)
}

func _WalletManagement_CreateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletManagementServer).CreateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletManagement/CreateChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletManagementServer).CreateChain(ctx, req.(*CreateChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletManagement_AssignedXpubToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignedXpubToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletManagementServer).AssignedXpubToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletManagement/AssignedXpubToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletManagementServer).AssignedXpubToGroup(ctx, req.(*AssignedXpubToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletManagement_GetChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletManagementServer).GetChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletManagement/GetChains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletManagementServer).GetChains(ctx, req.(*GetChainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletManagement_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletManagementServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletManagement/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletManagementServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WalletManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WalletManagement",
	HandlerType: (*WalletManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChain",
			Handler:    _WalletManagement_CreateChain_Handler,
		},
		{
			MethodName: "AssignedXpubToGroup",
			Handler:    _WalletManagement_AssignedXpubToGroup_Handler,
		},
		{
			MethodName: "GetChains",
			Handler:    _WalletManagement_GetChains_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _WalletManagement_CreateWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet_management.proto",
}

func init() {
	proto.RegisterFile("wallet_management.proto", fileDescriptor_wallet_management_7907d0a21409ef20)
}

var fileDescriptor_wallet_management_7907d0a21409ef20 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0xb2, 0xad, 0xc6, 0xa3, 0x92, 0xba, 0x93, 0x34, 0x11, 0x22, 0x14, 0xb3, 0x97, 0x9a,
	0x1e, 0x7c, 0x70, 0xd3, 0x96, 0x52, 0x28, 0x84, 0x04, 0x42, 0x0e, 0xed, 0x41, 0x2e, 0xb4, 0xf4,
	0x52, 0x24, 0x6b, 0xb1, 0x17, 0xf4, 0xb1, 0xd5, 0xae, 0x08, 0xfe, 0x2d, 0xfd, 0xb3, 0x65, 0x76,
	0x57, 0x8e, 0x1c, 0xc7, 0xf4, 0xa6, 0x37, 0x5f, 0x6f, 0xe6, 0xed, 0x13, 0x9c, 0xdf, 0x27, 0x79,
	0xce, 0xf5, 0xef, 0x22, 0x29, 0x93, 0x15, 0x2f, 0x78, 0xa9, 0x67, 0xb2, 0xae, 0x74, 0x85, 0x9e,
	0x4c, 0xd9, 0x2f, 0xc0, 0xeb, 0x9a, 0x27, 0x9a, 0x5f, 0xaf, 0x13, 0x51, 0xc6, 0xfc, 0x4f, 0xc3,
	0x95, 0xc6, 0x33, 0xf0, 0xd5, 0xa6, 0x48, 0xab, 0x3c, 0xec, 0x4d, 0x7a, 0xd3, 0x51, 0xec, 0x10,
	0x9e, 0xc2, 0x30, 0x15, 0x5a, 0x64, 0xa1, 0x67, 0xc2, 0x16, 0x98, 0x6a, 0x9d, 0xe8, 0x46, 0x85,
	0xfd, 0x49, 0x6f, 0x7a, 0x14, 0x3b, 0xc4, 0xde, 0xc2, 0x78, 0x67, 0xb6, 0xcc, 0x37, 0x54, 0x5b,
	0x73, 0xd5, 0xe4, 0xda, 0x4c, 0x3e, 0x8a, 0x1d, 0x62, 0x1f, 0x20, 0xba, 0x52, 0x4a, 0xac, 0x4a,
	0x9e, 0xfd, 0x94, 0x4d, 0xfa, 0xbd, 0xba, 0xad, 0xab, 0x46, 0xb6, 0xfb, 0x84, 0xf0, 0x6c, 0x45,
	0x58, 0x64, 0x6e, 0xa1, 0x16, 0xb2, 0x08, 0xc2, 0x27, 0xfb, 0x64, 0xbe, 0x61, 0x08, 0xe3, 0x5b,
	0xae, 0x0d, 0xb9, 0x72, 0x93, 0xd8, 0x3d, 0x04, 0x0b, 0x51, 0xc8, 0xdc, 0xee, 0x84, 0xc7, 0xe0,
	0x6d, 0x67, 0x7a, 0x22, 0x43, 0x84, 0x41, 0x99, 0x14, 0xdc, 0xdd, 0x67, 0xbe, 0x89, 0x3c, 0x15,
	0xf2, 0xf2, 0x52, 0x64, 0xe6, 0xbe, 0x61, 0xdc, 0x42, 0xaa, 0x5e, 0x56, 0xa2, 0x0c, 0x07, 0xb6,
	0x9a, 0xbe, 0x3b, 0x62, 0x0c, 0x77, 0xc4, 0xf8, 0x04, 0xc7, 0x9d, 0x65, 0x48, 0x8a, 0x37, 0xe0,
	0x2f, 0x0d, 0x0c, 0x7b, 0x93, 0xfe, 0x34, 0x98, 0xbf, 0x98, 0xc9, 0x74, 0xd6, 0x59, 0x2e, 0x76,
	0x69, 0x56, 0xc1, 0x89, 0xd5, 0xf1, 0x87, 0x79, 0xc8, 0xff, 0x8a, 0x82, 0x17, 0x30, 0x32, 0xad,
	0x9d, 0x53, 0x1e, 0x02, 0x38, 0x81, 0xc0, 0x1c, 0xb0, 0x5c, 0x27, 0xe5, 0x8a, 0xbb, 0x9b, 0xba,
	0x21, 0xb6, 0x06, 0xdf, 0x52, 0x91, 0x3e, 0x77, 0x37, 0xad, 0x3e, 0x77, 0x37, 0xc4, 0x79, 0x95,
	0x65, 0x35, 0x57, 0xca, 0xcd, 0x6d, 0x21, 0x71, 0x9a, 0xad, 0xbf, 0x11, 0x67, 0xdf, 0x72, 0x6e,
	0x03, 0xa4, 0xca, 0xc2, 0xaa, 0x32, 0xb0, 0xaa, 0x58, 0xc4, 0x3e, 0xc2, 0xcb, 0xdd, 0xd3, 0x48,
	0x18, 0x06, 0xbe, 0xb5, 0xac, 0x21, 0x0e, 0xe6, 0x40, 0xc2, 0xb8, 0x02, 0x97, 0x99, 0xff, 0xf5,
	0x60, 0x6c, 0x43, 0x5f, 0xb7, 0xb6, 0xc6, 0xcf, 0x10, 0x74, 0x0c, 0x87, 0x67, 0xd4, 0xb7, 0xef,
	0xee, 0xe8, 0x74, 0x2f, 0x4e, 0xac, 0x0b, 0x38, 0x79, 0xc2, 0x49, 0xf8, 0x9a, 0x8a, 0x0f, 0x5b,
	0x33, 0xba, 0x38, 0x98, 0xa7, 0xa1, 0xef, 0x61, 0xb4, 0x7d, 0x75, 0x34, 0xbc, 0x8f, 0x1d, 0x19,
	0xe1, 0xa3, 0x28, 0xb5, 0x7d, 0x81, 0xe7, 0x5d, 0x59, 0xf0, 0xfc, 0x61, 0xe3, 0x1d, 0x0f, 0x44,
	0xaf, 0xf6, 0x13, 0x32, 0xdf, 0xa4, 0xbe, 0xf9, 0xc1, 0xdf, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x35, 0xa2, 0xec, 0x28, 0xfb, 0x03, 0x00, 0x00,
}
