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
	return fileDescriptor_transaction_4698c10da3adb9a7, []int{0}
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
	return fileDescriptor_transaction_4698c10da3adb9a7, []int{1}
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
	return fileDescriptor_transaction_4698c10da3adb9a7, []int{2}
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

func init() {
	proto.RegisterType((*SimpleAsset)(nil), "pb.SimpleAsset")
	proto.RegisterType((*AssignAssetsToWalletRequest)(nil), "pb.AssignAssetsToWalletRequest")
	proto.RegisterType((*AssignAssetsToWalletReply)(nil), "pb.AssignAssetsToWalletReply")
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

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	AssignAssetsToWallet(context.Context, *AssignAssetsToWalletRequest) (*AssignAssetsToWalletReply, error)
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

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignAssetsToWallet",
			Handler:    _Transaction_AssignAssetsToWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_transaction_4698c10da3adb9a7) }

var fileDescriptor_transaction_4698c10da3adb9a7 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x49, 0x37, 0xe7, 0x7c, 0x15, 0xc4, 0x20, 0x12, 0x37, 0xc4, 0xd2, 0x53, 0x4f, 0x3d,
	0x6c, 0x9f, 0x60, 0xb0, 0xcb, 0xae, 0x71, 0xe8, 0x39, 0x5d, 0x9f, 0x12, 0x48, 0x9b, 0xd8, 0x97,
	0x1d, 0xf2, 0x21, 0xfc, 0xce, 0xb2, 0xd0, 0xea, 0x84, 0xea, 0x2d, 0xbf, 0x97, 0x7f, 0xf8, 0xfd,
	0x93, 0xc0, 0xad, 0xef, 0x54, 0x4b, 0xea, 0xe0, 0xb5, 0x6d, 0x4b, 0xd7, 0x59, 0x6f, 0x79, 0xe2,
	0xaa, 0xfc, 0x93, 0x41, 0xfa, 0xac, 0x1b, 0x67, 0x70, 0x43, 0x84, 0x9e, 0x0b, 0xb8, 0x8c, 0x8b,
	0xdd, 0x56, 0xb0, 0x8c, 0x15, 0x57, 0x72, 0x40, 0x7e, 0x0f, 0x33, 0x0a, 0x4d, 0x65, 0x8d, 0x48,
	0xe2, 0x46, 0x4f, 0x71, 0xee, 0x95, 0x3f, 0x92, 0x98, 0x64, 0xac, 0x98, 0xcb, 0x9e, 0xf8, 0x02,
	0xe6, 0xbb, 0x1a, 0x5b, 0xaf, 0xdf, 0x82, 0x98, 0xc6, 0x13, 0xdf, 0x7c, 0xb2, 0x6c, 0xf1, 0xa0,
	0x1b, 0x65, 0xc4, 0x45, 0xc6, 0x8a, 0xa9, 0x1c, 0x30, 0x37, 0xb0, 0xdc, 0x10, 0xe9, 0xf7, 0x36,
	0x6a, 0x69, 0x6f, 0x5f, 0x95, 0x31, 0xe8, 0x25, 0x7e, 0x1c, 0x91, 0x62, 0x3d, 0x55, 0xd7, 0x1d,
	0x12, 0x0d, 0xf5, 0x7a, 0xe4, 0x6b, 0xb8, 0x3e, 0xbb, 0x07, 0x89, 0x24, 0x9b, 0x14, 0xe9, 0xea,
	0xa6, 0x74, 0x55, 0x79, 0x36, 0x97, 0xbf, 0x42, 0xf9, 0x12, 0x1e, 0xc6, 0x6d, 0xce, 0x84, 0x15,
	0x42, 0xba, 0xff, 0x79, 0x33, 0xfe, 0x02, 0x77, 0x63, 0x59, 0xfe, 0x74, 0x52, 0xfc, 0xd3, 0x79,
	0xf1, 0xf8, 0x77, 0xc0, 0x99, 0x50, 0xcd, 0xe2, 0x67, 0xac, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x12, 0x02, 0xe7, 0xff, 0xa1, 0x01, 0x00, 0x00,
}
