// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet_key.proto

package pb

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

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

type GenerateMnemonicRequest struct {
	Namespaceid          string   `protobuf:"bytes,1,opt,name=namespaceid" json:"namespaceid,omitempty"`
	Bip44Ids             []int32  `protobuf:"varint,2,rep,packed,name=bip44ids" json:"bip44ids,omitempty"`
	Bip44AccountSize     int32    `protobuf:"varint,3,opt,name=bip44accountSize" json:"bip44accountSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateMnemonicRequest) Reset()         { *m = GenerateMnemonicRequest{} }
func (m *GenerateMnemonicRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateMnemonicRequest) ProtoMessage()    {}
func (*GenerateMnemonicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{0}
}
func (m *GenerateMnemonicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateMnemonicRequest.Unmarshal(m, b)
}
func (m *GenerateMnemonicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateMnemonicRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateMnemonicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateMnemonicRequest.Merge(dst, src)
}
func (m *GenerateMnemonicRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateMnemonicRequest.Size(m)
}
func (m *GenerateMnemonicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateMnemonicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateMnemonicRequest proto.InternalMessageInfo

func (m *GenerateMnemonicRequest) GetNamespaceid() string {
	if m != nil {
		return m.Namespaceid
	}
	return ""
}

func (m *GenerateMnemonicRequest) GetBip44Ids() []int32 {
	if m != nil {
		return m.Bip44Ids
	}
	return nil
}

func (m *GenerateMnemonicRequest) GetBip44AccountSize() int32 {
	if m != nil {
		return m.Bip44AccountSize
	}
	return 0
}

type Bip44AccountKey struct {
	Account              int32    `protobuf:"varint,1,opt,name=account" json:"account,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bip44AccountKey) Reset()         { *m = Bip44AccountKey{} }
func (m *Bip44AccountKey) String() string { return proto.CompactTextString(m) }
func (*Bip44AccountKey) ProtoMessage()    {}
func (*Bip44AccountKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{1}
}
func (m *Bip44AccountKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bip44AccountKey.Unmarshal(m, b)
}
func (m *Bip44AccountKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bip44AccountKey.Marshal(b, m, deterministic)
}
func (dst *Bip44AccountKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bip44AccountKey.Merge(dst, src)
}
func (m *Bip44AccountKey) XXX_Size() int {
	return xxx_messageInfo_Bip44AccountKey.Size(m)
}
func (m *Bip44AccountKey) XXX_DiscardUnknown() {
	xxx_messageInfo_Bip44AccountKey.DiscardUnknown(m)
}

var xxx_messageInfo_Bip44AccountKey proto.InternalMessageInfo

func (m *Bip44AccountKey) GetAccount() int32 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *Bip44AccountKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Bip44ThirdXpubsForChain struct {
	Chain                int32              `protobuf:"varint,1,opt,name=chain" json:"chain,omitempty"`
	Xpubs                []*Bip44AccountKey `protobuf:"bytes,2,rep,name=xpubs" json:"xpubs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Bip44ThirdXpubsForChain) Reset()         { *m = Bip44ThirdXpubsForChain{} }
func (m *Bip44ThirdXpubsForChain) String() string { return proto.CompactTextString(m) }
func (*Bip44ThirdXpubsForChain) ProtoMessage()    {}
func (*Bip44ThirdXpubsForChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{2}
}
func (m *Bip44ThirdXpubsForChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bip44ThirdXpubsForChain.Unmarshal(m, b)
}
func (m *Bip44ThirdXpubsForChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bip44ThirdXpubsForChain.Marshal(b, m, deterministic)
}
func (dst *Bip44ThirdXpubsForChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bip44ThirdXpubsForChain.Merge(dst, src)
}
func (m *Bip44ThirdXpubsForChain) XXX_Size() int {
	return xxx_messageInfo_Bip44ThirdXpubsForChain.Size(m)
}
func (m *Bip44ThirdXpubsForChain) XXX_DiscardUnknown() {
	xxx_messageInfo_Bip44ThirdXpubsForChain.DiscardUnknown(m)
}

var xxx_messageInfo_Bip44ThirdXpubsForChain proto.InternalMessageInfo

func (m *Bip44ThirdXpubsForChain) GetChain() int32 {
	if m != nil {
		return m.Chain
	}
	return 0
}

func (m *Bip44ThirdXpubsForChain) GetXpubs() []*Bip44AccountKey {
	if m != nil {
		return m.Xpubs
	}
	return nil
}

type GenerateMnemonicReply struct {
	Chainsxpubs          []*Bip44ThirdXpubsForChain `protobuf:"bytes,1,rep,name=chainsxpubs" json:"chainsxpubs,omitempty"`
	Version              string                     `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GenerateMnemonicReply) Reset()         { *m = GenerateMnemonicReply{} }
func (m *GenerateMnemonicReply) String() string { return proto.CompactTextString(m) }
func (*GenerateMnemonicReply) ProtoMessage()    {}
func (*GenerateMnemonicReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{3}
}
func (m *GenerateMnemonicReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateMnemonicReply.Unmarshal(m, b)
}
func (m *GenerateMnemonicReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateMnemonicReply.Marshal(b, m, deterministic)
}
func (dst *GenerateMnemonicReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateMnemonicReply.Merge(dst, src)
}
func (m *GenerateMnemonicReply) XXX_Size() int {
	return xxx_messageInfo_GenerateMnemonicReply.Size(m)
}
func (m *GenerateMnemonicReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateMnemonicReply.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateMnemonicReply proto.InternalMessageInfo

func (m *GenerateMnemonicReply) GetChainsxpubs() []*Bip44ThirdXpubsForChain {
	if m != nil {
		return m.Chainsxpubs
	}
	return nil
}

func (m *GenerateMnemonicReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type WalletHD struct {
	CoinType             int32    `protobuf:"varint,1,opt,name=CoinType" json:"CoinType,omitempty"`
	Account              int32    `protobuf:"varint,2,opt,name=Account" json:"Account,omitempty"`
	Change               int32    `protobuf:"varint,3,opt,name=Change" json:"Change,omitempty"`
	AddressIndex         uint32   `protobuf:"varint,4,opt,name=AddressIndex" json:"AddressIndex,omitempty"`
	MnemonicVersion      string   `protobuf:"bytes,5,opt,name=MnemonicVersion" json:"MnemonicVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletHD) Reset()         { *m = WalletHD{} }
func (m *WalletHD) String() string { return proto.CompactTextString(m) }
func (*WalletHD) ProtoMessage()    {}
func (*WalletHD) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{4}
}
func (m *WalletHD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletHD.Unmarshal(m, b)
}
func (m *WalletHD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletHD.Marshal(b, m, deterministic)
}
func (dst *WalletHD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletHD.Merge(dst, src)
}
func (m *WalletHD) XXX_Size() int {
	return xxx_messageInfo_WalletHD.Size(m)
}
func (m *WalletHD) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletHD.DiscardUnknown(m)
}

var xxx_messageInfo_WalletHD proto.InternalMessageInfo

func (m *WalletHD) GetCoinType() int32 {
	if m != nil {
		return m.CoinType
	}
	return 0
}

func (m *WalletHD) GetAccount() int32 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *WalletHD) GetChange() int32 {
	if m != nil {
		return m.Change
	}
	return 0
}

func (m *WalletHD) GetAddressIndex() uint32 {
	if m != nil {
		return m.AddressIndex
	}
	return 0
}

func (m *WalletHD) GetMnemonicVersion() string {
	if m != nil {
		return m.MnemonicVersion
	}
	return ""
}

type SignedBitcoincoreTxRequest struct {
	WalletHD             *WalletHD `protobuf:"bytes,1,opt,name=WalletHD" json:"WalletHD,omitempty"`
	TxHex                string    `protobuf:"bytes,2,opt,name=TxHex" json:"TxHex,omitempty"`
	VinAmount            int64     `protobuf:"varint,3,opt,name=VinAmount" json:"VinAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SignedBitcoincoreTxRequest) Reset()         { *m = SignedBitcoincoreTxRequest{} }
func (m *SignedBitcoincoreTxRequest) String() string { return proto.CompactTextString(m) }
func (*SignedBitcoincoreTxRequest) ProtoMessage()    {}
func (*SignedBitcoincoreTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{5}
}
func (m *SignedBitcoincoreTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedBitcoincoreTxRequest.Unmarshal(m, b)
}
func (m *SignedBitcoincoreTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedBitcoincoreTxRequest.Marshal(b, m, deterministic)
}
func (dst *SignedBitcoincoreTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedBitcoincoreTxRequest.Merge(dst, src)
}
func (m *SignedBitcoincoreTxRequest) XXX_Size() int {
	return xxx_messageInfo_SignedBitcoincoreTxRequest.Size(m)
}
func (m *SignedBitcoincoreTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedBitcoincoreTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignedBitcoincoreTxRequest proto.InternalMessageInfo

func (m *SignedBitcoincoreTxRequest) GetWalletHD() *WalletHD {
	if m != nil {
		return m.WalletHD
	}
	return nil
}

func (m *SignedBitcoincoreTxRequest) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

func (m *SignedBitcoincoreTxRequest) GetVinAmount() int64 {
	if m != nil {
		return m.VinAmount
	}
	return 0
}

type SignedBitcoincoreTxReply struct {
	SignedTxHex          string   `protobuf:"bytes,1,opt,name=SignedTxHex" json:"SignedTxHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedBitcoincoreTxReply) Reset()         { *m = SignedBitcoincoreTxReply{} }
func (m *SignedBitcoincoreTxReply) String() string { return proto.CompactTextString(m) }
func (*SignedBitcoincoreTxReply) ProtoMessage()    {}
func (*SignedBitcoincoreTxReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{6}
}
func (m *SignedBitcoincoreTxReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedBitcoincoreTxReply.Unmarshal(m, b)
}
func (m *SignedBitcoincoreTxReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedBitcoincoreTxReply.Marshal(b, m, deterministic)
}
func (dst *SignedBitcoincoreTxReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedBitcoincoreTxReply.Merge(dst, src)
}
func (m *SignedBitcoincoreTxReply) XXX_Size() int {
	return xxx_messageInfo_SignedBitcoincoreTxReply.Size(m)
}
func (m *SignedBitcoincoreTxReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedBitcoincoreTxReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignedBitcoincoreTxReply proto.InternalMessageInfo

func (m *SignedBitcoincoreTxReply) GetSignedTxHex() string {
	if m != nil {
		return m.SignedTxHex
	}
	return ""
}

type SignedEthereumTxRequest struct {
	TxHex                string    `protobuf:"bytes,1,opt,name=TxHex" json:"TxHex,omitempty"`
	ChainID              string    `protobuf:"bytes,2,opt,name=ChainID" json:"ChainID,omitempty"`
	WalletHD             *WalletHD `protobuf:"bytes,3,opt,name=WalletHD" json:"WalletHD,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SignedEthereumTxRequest) Reset()         { *m = SignedEthereumTxRequest{} }
func (m *SignedEthereumTxRequest) String() string { return proto.CompactTextString(m) }
func (*SignedEthereumTxRequest) ProtoMessage()    {}
func (*SignedEthereumTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{7}
}
func (m *SignedEthereumTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedEthereumTxRequest.Unmarshal(m, b)
}
func (m *SignedEthereumTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedEthereumTxRequest.Marshal(b, m, deterministic)
}
func (dst *SignedEthereumTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedEthereumTxRequest.Merge(dst, src)
}
func (m *SignedEthereumTxRequest) XXX_Size() int {
	return xxx_messageInfo_SignedEthereumTxRequest.Size(m)
}
func (m *SignedEthereumTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedEthereumTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignedEthereumTxRequest proto.InternalMessageInfo

func (m *SignedEthereumTxRequest) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

func (m *SignedEthereumTxRequest) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *SignedEthereumTxRequest) GetWalletHD() *WalletHD {
	if m != nil {
		return m.WalletHD
	}
	return nil
}

type SignedEthereumTxReply struct {
	SignedTxHex          string   `protobuf:"bytes,1,opt,name=SignedTxHex" json:"SignedTxHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedEthereumTxReply) Reset()         { *m = SignedEthereumTxReply{} }
func (m *SignedEthereumTxReply) String() string { return proto.CompactTextString(m) }
func (*SignedEthereumTxReply) ProtoMessage()    {}
func (*SignedEthereumTxReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_key_aca8896db1c90657, []int{8}
}
func (m *SignedEthereumTxReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedEthereumTxReply.Unmarshal(m, b)
}
func (m *SignedEthereumTxReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedEthereumTxReply.Marshal(b, m, deterministic)
}
func (dst *SignedEthereumTxReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedEthereumTxReply.Merge(dst, src)
}
func (m *SignedEthereumTxReply) XXX_Size() int {
	return xxx_messageInfo_SignedEthereumTxReply.Size(m)
}
func (m *SignedEthereumTxReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedEthereumTxReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignedEthereumTxReply proto.InternalMessageInfo

func (m *SignedEthereumTxReply) GetSignedTxHex() string {
	if m != nil {
		return m.SignedTxHex
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateMnemonicRequest)(nil), "pb.GenerateMnemonicRequest")
	proto.RegisterType((*Bip44AccountKey)(nil), "pb.Bip44AccountKey")
	proto.RegisterType((*Bip44ThirdXpubsForChain)(nil), "pb.Bip44ThirdXpubsForChain")
	proto.RegisterType((*GenerateMnemonicReply)(nil), "pb.GenerateMnemonicReply")
	proto.RegisterType((*WalletHD)(nil), "pb.WalletHD")
	proto.RegisterType((*SignedBitcoincoreTxRequest)(nil), "pb.SignedBitcoincoreTxRequest")
	proto.RegisterType((*SignedBitcoincoreTxReply)(nil), "pb.SignedBitcoincoreTxReply")
	proto.RegisterType((*SignedEthereumTxRequest)(nil), "pb.SignedEthereumTxRequest")
	proto.RegisterType((*SignedEthereumTxReply)(nil), "pb.SignedEthereumTxReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletKeyClient is the client API for WalletKey service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletKeyClient interface {
	GenerateMnemonic(ctx context.Context, in *GenerateMnemonicRequest, opts ...grpc.CallOption) (*GenerateMnemonicReply, error)
	SignedBitcoincoreTx(ctx context.Context, in *SignedBitcoincoreTxRequest, opts ...grpc.CallOption) (*SignedBitcoincoreTxReply, error)
	SignedEthereumTx(ctx context.Context, in *SignedEthereumTxRequest, opts ...grpc.CallOption) (*SignedEthereumTxReply, error)
}

type walletKeyClient struct {
	cc *grpc.ClientConn
}

func NewWalletKeyClient(cc *grpc.ClientConn) WalletKeyClient {
	return &walletKeyClient{cc}
}

func (c *walletKeyClient) GenerateMnemonic(ctx context.Context, in *GenerateMnemonicRequest, opts ...grpc.CallOption) (*GenerateMnemonicReply, error) {
	out := new(GenerateMnemonicReply)
	err := c.cc.Invoke(ctx, "/pb.WalletKey/GenerateMnemonic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKeyClient) SignedBitcoincoreTx(ctx context.Context, in *SignedBitcoincoreTxRequest, opts ...grpc.CallOption) (*SignedBitcoincoreTxReply, error) {
	out := new(SignedBitcoincoreTxReply)
	err := c.cc.Invoke(ctx, "/pb.WalletKey/SignedBitcoincoreTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKeyClient) SignedEthereumTx(ctx context.Context, in *SignedEthereumTxRequest, opts ...grpc.CallOption) (*SignedEthereumTxReply, error) {
	out := new(SignedEthereumTxReply)
	err := c.cc.Invoke(ctx, "/pb.WalletKey/SignedEthereumTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyServer is the server API for WalletKey service.
type WalletKeyServer interface {
	GenerateMnemonic(context.Context, *GenerateMnemonicRequest) (*GenerateMnemonicReply, error)
	SignedBitcoincoreTx(context.Context, *SignedBitcoincoreTxRequest) (*SignedBitcoincoreTxReply, error)
	SignedEthereumTx(context.Context, *SignedEthereumTxRequest) (*SignedEthereumTxReply, error)
}

func RegisterWalletKeyServer(s *grpc.Server, srv WalletKeyServer) {
	s.RegisterService(&_WalletKey_serviceDesc, srv)
}

func _WalletKey_GenerateMnemonic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateMnemonicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKeyServer).GenerateMnemonic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletKey/GenerateMnemonic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKeyServer).GenerateMnemonic(ctx, req.(*GenerateMnemonicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKey_SignedBitcoincoreTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedBitcoincoreTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKeyServer).SignedBitcoincoreTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletKey/SignedBitcoincoreTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKeyServer).SignedBitcoincoreTx(ctx, req.(*SignedBitcoincoreTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKey_SignedEthereumTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedEthereumTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKeyServer).SignedEthereumTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletKey/SignedEthereumTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKeyServer).SignedEthereumTx(ctx, req.(*SignedEthereumTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WalletKey_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.WalletKey",
	HandlerType: (*WalletKeyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateMnemonic",
			Handler:    _WalletKey_GenerateMnemonic_Handler,
		},
		{
			MethodName: "SignedBitcoincoreTx",
			Handler:    _WalletKey_SignedBitcoincoreTx_Handler,
		},
		{
			MethodName: "SignedEthereumTx",
			Handler:    _WalletKey_SignedEthereumTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet_key.proto",
}

func init() { proto.RegisterFile("wallet_key.proto", fileDescriptor_wallet_key_aca8896db1c90657) }

var fileDescriptor_wallet_key_aca8896db1c90657 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x13, 0xdc, 0x36, 0x93, 0xa2, 0x46, 0xdb, 0x96, 0x98, 0x50, 0x21, 0x6b, 0x4f, 0x86,
	0x43, 0x0e, 0xa1, 0x17, 0x24, 0x7a, 0x48, 0x53, 0xa0, 0x2d, 0xe2, 0xb2, 0x89, 0x0a, 0xe2, 0x82,
	0xfc, 0x31, 0x6a, 0x56, 0x4d, 0xd6, 0xc6, 0x76, 0xc0, 0xee, 0x95, 0xff, 0xc2, 0x6f, 0xe4, 0x88,
	0xf6, 0xc3, 0x49, 0x9a, 0x0f, 0xd4, 0xdb, 0xbe, 0xb7, 0xeb, 0xb7, 0xef, 0xcd, 0xcc, 0x1a, 0x5a,
	0xbf, 0xfc, 0xc9, 0x04, 0xf3, 0xef, 0x77, 0x58, 0x76, 0x93, 0x34, 0xce, 0x63, 0x52, 0x4b, 0x02,
	0xfa, 0xdb, 0x82, 0xf6, 0x47, 0x14, 0x98, 0xfa, 0x39, 0x7e, 0x16, 0x38, 0x8d, 0x05, 0x0f, 0x19,
	0xfe, 0x98, 0x61, 0x96, 0x13, 0x17, 0x9a, 0xc2, 0x9f, 0x62, 0x96, 0xf8, 0x21, 0xf2, 0xc8, 0xb1,
	0x5c, 0xcb, 0x6b, 0xb0, 0x65, 0x8a, 0x74, 0x60, 0x2f, 0xe0, 0xc9, 0xe9, 0x29, 0x8f, 0x32, 0xa7,
	0xe6, 0xd6, 0x3d, 0x9b, 0xcd, 0x31, 0x79, 0x0d, 0x2d, 0xb5, 0xf6, 0xc3, 0x30, 0x9e, 0x89, 0x7c,
	0xc8, 0xef, 0xd1, 0xa9, 0xbb, 0x96, 0x67, 0xb3, 0x35, 0x9e, 0x9e, 0xc1, 0xc1, 0xb9, 0xe4, 0xfa,
	0x9a, 0xfb, 0x84, 0x25, 0x71, 0x60, 0xd7, 0x9c, 0x50, 0x17, 0xdb, 0xac, 0x82, 0xa4, 0x05, 0xf5,
	0x3b, 0x2c, 0x9d, 0x9a, 0xb2, 0x23, 0x97, 0xf4, 0x1b, 0xb4, 0xd5, 0xe7, 0xa3, 0x31, 0x4f, 0xa3,
	0xaf, 0xc9, 0x2c, 0xc8, 0x3e, 0xc4, 0xe9, 0x60, 0xec, 0x73, 0x41, 0x8e, 0xc0, 0x0e, 0xe5, 0xc2,
	0x88, 0x68, 0x40, 0x5e, 0x81, 0x5d, 0xc8, 0x63, 0xca, 0x74, 0xb3, 0x77, 0xd8, 0x4d, 0x82, 0xee,
	0x8a, 0x01, 0xa6, 0x4f, 0xd0, 0x04, 0x8e, 0xd7, 0xeb, 0x93, 0x4c, 0x4a, 0x72, 0x06, 0x4d, 0x25,
	0x96, 0x69, 0x25, 0x4b, 0x29, 0xbd, 0x98, 0x2b, 0xad, 0x7b, 0x61, 0xcb, 0xe7, 0x65, 0xbe, 0x9f,
	0x98, 0x66, 0x3c, 0x16, 0x26, 0x49, 0x05, 0xe9, 0x1f, 0x0b, 0xf6, 0xbe, 0xa8, 0x5e, 0x5d, 0x5e,
	0xc8, 0x0a, 0x0f, 0x62, 0x2e, 0x46, 0x65, 0x82, 0x26, 0xc2, 0x1c, 0x4b, 0x09, 0xe3, 0x57, 0x49,
	0xd8, 0xac, 0x82, 0xe4, 0x19, 0xec, 0x0c, 0xc6, 0xbe, 0xb8, 0xad, 0x2a, 0x6e, 0x10, 0xa1, 0xb0,
	0xdf, 0x8f, 0xa2, 0x14, 0xb3, 0xec, 0x4a, 0x44, 0x58, 0x38, 0x4f, 0x5c, 0xcb, 0x7b, 0xca, 0x1e,
	0x70, 0xc4, 0x83, 0x83, 0x2a, 0xe8, 0x8d, 0x31, 0x68, 0x2b, 0x83, 0xab, 0x34, 0xbd, 0x87, 0xce,
	0x90, 0xdf, 0x0a, 0x8c, 0xce, 0x79, 0x1e, 0xc6, 0x5c, 0x84, 0x71, 0x8a, 0xa3, 0xa2, 0x9a, 0x1e,
	0x6f, 0x91, 0x42, 0x39, 0x6f, 0xf6, 0xf6, 0x65, 0x71, 0x2a, 0x8e, 0x2d, 0x32, 0x1e, 0x81, 0x3d,
	0x2a, 0x2e, 0xb1, 0x30, 0x85, 0xd0, 0x80, 0x9c, 0x40, 0xe3, 0x86, 0x8b, 0xfe, 0x54, 0xe5, 0x93,
	0x31, 0xea, 0x6c, 0x41, 0xd0, 0x77, 0xe0, 0x6c, 0xbc, 0x5b, 0x76, 0xc6, 0x85, 0xa6, 0xde, 0xd3,
	0xaa, 0x66, 0x6e, 0x97, 0x28, 0x9a, 0x41, 0x5b, 0xc3, 0xf7, 0xf9, 0x18, 0x53, 0x9c, 0x4d, 0x17,
	0xb6, 0xe7, 0x66, 0xac, 0x65, 0x33, 0x0e, 0xec, 0xaa, 0x1e, 0x5e, 0x5d, 0x54, 0xdd, 0x32, 0xf0,
	0x41, 0xcc, 0xfa, 0xff, 0x62, 0xd2, 0xb7, 0x70, 0xbc, 0x7e, 0xe9, 0xa3, 0xfc, 0xf6, 0xfe, 0x5a,
	0xd0, 0xd0, 0x3a, 0xf2, 0x69, 0x5c, 0x43, 0x6b, 0x75, 0x24, 0x89, 0x1a, 0xbc, 0x2d, 0x0f, 0xb9,
	0xf3, 0x7c, 0xf3, 0xa6, 0xbc, 0x7b, 0x08, 0x87, 0x1b, 0xea, 0x48, 0x5e, 0xca, 0x2f, 0xb6, 0x37,
	0xb7, 0x73, 0xb2, 0x75, 0x5f, 0x8a, 0x5e, 0x43, 0x6b, 0x35, 0xa9, 0x36, 0xb8, 0xa5, 0xe8, 0xda,
	0xe0, 0xc6, 0xe2, 0x04, 0x3b, 0xea, 0x5f, 0xf5, 0xe6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0xb4, 0xbd, 0x24, 0xbf, 0x04, 0x00, 0x00,
}
