// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

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

type SimpleAsset struct {
	AssetID              string   `protobuf:"bytes,1,opt,name=AssetID" json:"AssetID,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	Identify             string   `protobuf:"bytes,4,opt,name=Identify" json:"Identify,omitempty"`
	Decimal              uint64   `protobuf:"varint,5,opt,name=Decimal" json:"Decimal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleAsset) Reset()         { *m = SimpleAsset{} }
func (m *SimpleAsset) String() string { return proto.CompactTextString(m) }
func (*SimpleAsset) ProtoMessage()    {}
func (*SimpleAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_17c0faa9de04a42f, []int{0}
}
func (m *SimpleAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleAsset.Unmarshal(m, b)
}
func (m *SimpleAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleAsset.Marshal(b, m, deterministic)
}
func (dst *SimpleAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleAsset.Merge(dst, src)
}
func (m *SimpleAsset) XXX_Size() int {
	return xxx_messageInfo_SimpleAsset.Size(m)
}
func (m *SimpleAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleAsset.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleAsset proto.InternalMessageInfo

func (m *SimpleAsset) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

func (m *SimpleAsset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *SimpleAsset) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *SimpleAsset) GetIdentify() string {
	if m != nil {
		return m.Identify
	}
	return ""
}

func (m *SimpleAsset) GetDecimal() uint64 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

type Wallet struct {
	Address              string   `protobuf:"bytes,1,opt,name=Address" json:"Address,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=Status" json:"Status,omitempty"`
	ChainName            string   `protobuf:"bytes,3,opt,name=ChainName" json:"ChainName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Wallet) Reset()         { *m = Wallet{} }
func (m *Wallet) String() string { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()    {}
func (*Wallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_17c0faa9de04a42f, []int{1}
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

func (m *Wallet) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Wallet) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Wallet) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

type AssignAssetsToWalletRequest struct {
	Address              string         `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	SimpleAssets         []*SimpleAsset `protobuf:"bytes,2,rep,name=SimpleAssets" json:"SimpleAssets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AssignAssetsToWalletRequest) Reset()         { *m = AssignAssetsToWalletRequest{} }
func (m *AssignAssetsToWalletRequest) String() string { return proto.CompactTextString(m) }
func (*AssignAssetsToWalletRequest) ProtoMessage()    {}
func (*AssignAssetsToWalletRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_17c0faa9de04a42f, []int{2}
}
func (m *AssignAssetsToWalletRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignAssetsToWalletRequest.Unmarshal(m, b)
}
func (m *AssignAssetsToWalletRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignAssetsToWalletRequest.Marshal(b, m, deterministic)
}
func (dst *AssignAssetsToWalletRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignAssetsToWalletRequest.Merge(dst, src)
}
func (m *AssignAssetsToWalletRequest) XXX_Size() int {
	return xxx_messageInfo_AssignAssetsToWalletRequest.Size(m)
}
func (m *AssignAssetsToWalletRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignAssetsToWalletRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignAssetsToWalletRequest proto.InternalMessageInfo

func (m *AssignAssetsToWalletRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AssignAssetsToWalletRequest) GetSimpleAssets() []*SimpleAsset {
	if m != nil {
		return m.SimpleAssets
	}
	return nil
}

type AssignAssetsToWalletReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignAssetsToWalletReply) Reset()         { *m = AssignAssetsToWalletReply{} }
func (m *AssignAssetsToWalletReply) String() string { return proto.CompactTextString(m) }
func (*AssignAssetsToWalletReply) ProtoMessage()    {}
func (*AssignAssetsToWalletReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_17c0faa9de04a42f, []int{3}
}
func (m *AssignAssetsToWalletReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignAssetsToWalletReply.Unmarshal(m, b)
}
func (m *AssignAssetsToWalletReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignAssetsToWalletReply.Marshal(b, m, deterministic)
}
func (dst *AssignAssetsToWalletReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignAssetsToWalletReply.Merge(dst, src)
}
func (m *AssignAssetsToWalletReply) XXX_Size() int {
	return xxx_messageInfo_AssignAssetsToWalletReply.Size(m)
}
func (m *AssignAssetsToWalletReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignAssetsToWalletReply.DiscardUnknown(m)
}

var xxx_messageInfo_AssignAssetsToWalletReply proto.InternalMessageInfo

type CreateBalancesForAssetRequest struct {
	Asset                *SimpleAsset `protobuf:"bytes,1,opt,name=Asset" json:"Asset,omitempty"`
	Wallets              []*Wallet    `protobuf:"bytes,2,rep,name=Wallets" json:"Wallets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateBalancesForAssetRequest) Reset()         { *m = CreateBalancesForAssetRequest{} }
func (m *CreateBalancesForAssetRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBalancesForAssetRequest) ProtoMessage()    {}
func (*CreateBalancesForAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_17c0faa9de04a42f, []int{4}
}
func (m *CreateBalancesForAssetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBalancesForAssetRequest.Unmarshal(m, b)
}
func (m *CreateBalancesForAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBalancesForAssetRequest.Marshal(b, m, deterministic)
}
func (dst *CreateBalancesForAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBalancesForAssetRequest.Merge(dst, src)
}
func (m *CreateBalancesForAssetRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBalancesForAssetRequest.Size(m)
}
func (m *CreateBalancesForAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBalancesForAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBalancesForAssetRequest proto.InternalMessageInfo

func (m *CreateBalancesForAssetRequest) GetAsset() *SimpleAsset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *CreateBalancesForAssetRequest) GetWallets() []*Wallet {
	if m != nil {
		return m.Wallets
	}
	return nil
}

type CreateBalancesForAssetReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBalancesForAssetReply) Reset()         { *m = CreateBalancesForAssetReply{} }
func (m *CreateBalancesForAssetReply) String() string { return proto.CompactTextString(m) }
func (*CreateBalancesForAssetReply) ProtoMessage()    {}
func (*CreateBalancesForAssetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_17c0faa9de04a42f, []int{5}
}
func (m *CreateBalancesForAssetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBalancesForAssetReply.Unmarshal(m, b)
}
func (m *CreateBalancesForAssetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBalancesForAssetReply.Marshal(b, m, deterministic)
}
func (dst *CreateBalancesForAssetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBalancesForAssetReply.Merge(dst, src)
}
func (m *CreateBalancesForAssetReply) XXX_Size() int {
	return xxx_messageInfo_CreateBalancesForAssetReply.Size(m)
}
func (m *CreateBalancesForAssetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBalancesForAssetReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBalancesForAssetReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SimpleAsset)(nil), "pb.SimpleAsset")
	proto.RegisterType((*Wallet)(nil), "pb.Wallet")
	proto.RegisterType((*AssignAssetsToWalletRequest)(nil), "pb.AssignAssetsToWalletRequest")
	proto.RegisterType((*AssignAssetsToWalletReply)(nil), "pb.AssignAssetsToWalletReply")
	proto.RegisterType((*CreateBalancesForAssetRequest)(nil), "pb.CreateBalancesForAssetRequest")
	proto.RegisterType((*CreateBalancesForAssetReply)(nil), "pb.CreateBalancesForAssetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	AssignAssetsToWallet(ctx context.Context, in *AssignAssetsToWalletRequest, opts ...grpc.CallOption) (*AssignAssetsToWalletReply, error)
	CreateBalancesForAsset(ctx context.Context, in *CreateBalancesForAssetRequest, opts ...grpc.CallOption) (*CreateBalancesForAssetReply, error)
}

type transactionClient struct {
	cc *grpc.ClientConn
}

func NewTransactionClient(cc *grpc.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) AssignAssetsToWallet(ctx context.Context, in *AssignAssetsToWalletRequest, opts ...grpc.CallOption) (*AssignAssetsToWalletReply, error) {
	out := new(AssignAssetsToWalletReply)
	err := c.cc.Invoke(ctx, "/pb.Transaction/AssignAssetsToWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) CreateBalancesForAsset(ctx context.Context, in *CreateBalancesForAssetRequest, opts ...grpc.CallOption) (*CreateBalancesForAssetReply, error) {
	out := new(CreateBalancesForAssetReply)
	err := c.cc.Invoke(ctx, "/pb.Transaction/CreateBalancesForAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	AssignAssetsToWallet(context.Context, *AssignAssetsToWalletRequest) (*AssignAssetsToWalletReply, error)
	CreateBalancesForAsset(context.Context, *CreateBalancesForAssetRequest) (*CreateBalancesForAssetReply, error)
}

func RegisterTransactionServer(s *grpc.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_AssignAssetsToWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignAssetsToWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).AssignAssetsToWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transaction/AssignAssetsToWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).AssignAssetsToWallet(ctx, req.(*AssignAssetsToWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_CreateBalancesForAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBalancesForAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).CreateBalancesForAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transaction/CreateBalancesForAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).CreateBalancesForAsset(ctx, req.(*CreateBalancesForAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignAssetsToWallet",
			Handler:    _Transaction_AssignAssetsToWallet_Handler,
		},
		{
			MethodName: "CreateBalancesForAsset",
			Handler:    _Transaction_CreateBalancesForAsset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_transaction_17c0faa9de04a42f) }

var fileDescriptor_transaction_17c0faa9de04a42f = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xc7, 0x53, 0xfe, 0x33, 0xfc, 0x92, 0x5f, 0xdc, 0x18, 0x52, 0x41, 0x02, 0x36, 0x9a, 0xf4,
	0xc4, 0x01, 0x9e, 0x00, 0x21, 0x26, 0x5c, 0x3c, 0x14, 0xa2, 0x1e, 0xbc, 0x6c, 0x61, 0xd5, 0x26,
	0xd3, 0x76, 0xed, 0x2c, 0x87, 0x3e, 0x84, 0x4f, 0xe5, 0x8b, 0x99, 0xee, 0xb6, 0xfc, 0x49, 0x0a,
	0xb7, 0x7e, 0x76, 0x66, 0xe7, 0xf3, 0xed, 0x64, 0xe1, 0x4a, 0x25, 0x3c, 0x22, 0xbe, 0x51, 0x41,
	0x1c, 0x8d, 0x65, 0x12, 0xab, 0x98, 0x55, 0xa4, 0xef, 0xfc, 0x58, 0xd0, 0x59, 0x05, 0xa1, 0x44,
	0x31, 0x23, 0x12, 0x8a, 0xd9, 0xd0, 0xd4, 0x1f, 0xcb, 0x85, 0x6d, 0x8d, 0x2c, 0xb7, 0xed, 0x15,
	0xc8, 0xba, 0xd0, 0xa0, 0x34, 0xf4, 0x63, 0xb4, 0x2b, 0xba, 0x90, 0x93, 0x3e, 0x57, 0x5c, 0xed,
	0xc8, 0xae, 0x8e, 0x2c, 0xb7, 0xe5, 0xe5, 0xc4, 0x7a, 0xd0, 0x5a, 0x6e, 0x45, 0xa4, 0x82, 0x8f,
	0xd4, 0xae, 0xe9, 0x1b, 0x7b, 0xce, 0x2c, 0x0b, 0xb1, 0x09, 0x42, 0x8e, 0x76, 0x7d, 0x64, 0xb9,
	0x35, 0xaf, 0x40, 0xe7, 0x0d, 0x1a, 0xaf, 0x1c, 0x31, 0x4f, 0xb2, 0xdd, 0x26, 0x82, 0x68, 0x9f,
	0xc4, 0x60, 0x66, 0x5c, 0x19, 0x63, 0xc5, 0x18, 0x0d, 0xb1, 0x5b, 0x68, 0xcf, 0xbf, 0x78, 0x10,
	0x3d, 0xf3, 0x50, 0xe8, 0x30, 0x6d, 0xef, 0x70, 0xe0, 0x20, 0xf4, 0x67, 0x44, 0xc1, 0x67, 0xa4,
	0x7f, 0x88, 0xd6, 0xb1, 0xf1, 0x78, 0xe2, 0x7b, 0x27, 0x48, 0xeb, 0xf8, 0xa9, 0x2e, 0x47, 0x36,
	0x85, 0x7f, 0x47, 0x1b, 0xca, 0xa4, 0x55, 0xb7, 0x33, 0xf9, 0x3f, 0x96, 0xfe, 0xf8, 0xe8, 0xdc,
	0x3b, 0x69, 0x72, 0xfa, 0x70, 0x53, 0x6e, 0x93, 0x98, 0x3a, 0x08, 0x83, 0x79, 0x22, 0xb8, 0x12,
	0x8f, 0x1c, 0x79, 0xb4, 0x11, 0xf4, 0x14, 0x27, 0x66, 0x48, 0x1e, 0xe6, 0x01, 0xea, 0x9a, 0x75,
	0x94, 0x12, 0x97, 0xa9, 0xb2, 0x7b, 0x68, 0x9a, 0xb1, 0x45, 0x28, 0xc8, 0x1a, 0x73, 0x53, 0x51,
	0x72, 0x06, 0xd0, 0x3f, 0x67, 0x93, 0x98, 0x4e, 0x7e, 0x2d, 0xe8, 0xac, 0x0f, 0x6f, 0x83, 0xbd,
	0xc0, 0x75, 0x59, 0x72, 0x36, 0xcc, 0x66, 0x5f, 0xd8, 0x60, 0x6f, 0x70, 0xbe, 0x41, 0x62, 0xca,
	0xde, 0xa1, 0x5b, 0x1e, 0x83, 0xdd, 0x65, 0x17, 0x2f, 0x2e, 0xa4, 0x37, 0xbc, 0xd4, 0x22, 0x31,
	0xf5, 0x1b, 0xfa, 0x49, 0x4f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x51, 0x96, 0xa2, 0xe7,
	0x02, 0x00, 0x00,
}
