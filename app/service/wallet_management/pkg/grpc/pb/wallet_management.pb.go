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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{0}
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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{1}
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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{2}
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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{3}
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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{4}
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
	Decimal              uint64   `protobuf:"varint,6,opt,name=Decimal" json:"Decimal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleChain) Reset()         { *m = SimpleChain{} }
func (m *SimpleChain) String() string { return proto.CompactTextString(m) }
func (*SimpleChain) ProtoMessage()    {}
func (*SimpleChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{5}
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

func (m *SimpleChain) GetDecimal() uint64 {
	if m != nil {
		return m.Decimal
	}
	return 0
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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{6}
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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{7}
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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{8}
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

type ChainWithWallets struct {
	Wallets              []*Wallet `protobuf:"bytes,1,rep,name=Wallets" json:"Wallets,omitempty"`
	TotalSize            int32     `protobuf:"varint,2,opt,name=TotalSize" json:"TotalSize,omitempty"`
	ChainName            string    `protobuf:"bytes,3,opt,name=ChainName" json:"ChainName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ChainWithWallets) Reset()         { *m = ChainWithWallets{} }
func (m *ChainWithWallets) String() string { return proto.CompactTextString(m) }
func (*ChainWithWallets) ProtoMessage()    {}
func (*ChainWithWallets) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{9}
}
func (m *ChainWithWallets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainWithWallets.Unmarshal(m, b)
}
func (m *ChainWithWallets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainWithWallets.Marshal(b, m, deterministic)
}
func (dst *ChainWithWallets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainWithWallets.Merge(dst, src)
}
func (m *ChainWithWallets) XXX_Size() int {
	return xxx_messageInfo_ChainWithWallets.Size(m)
}
func (m *ChainWithWallets) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainWithWallets.DiscardUnknown(m)
}

var xxx_messageInfo_ChainWithWallets proto.InternalMessageInfo

func (m *ChainWithWallets) GetWallets() []*Wallet {
	if m != nil {
		return m.Wallets
	}
	return nil
}

func (m *ChainWithWallets) GetTotalSize() int32 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *ChainWithWallets) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
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
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{10}
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

type GetWalletsRequest struct {
	Groupid              string   `protobuf:"bytes,1,opt,name=groupid" json:"groupid,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Bip44Change          int32    `protobuf:"varint,4,opt,name=bip44Change" json:"bip44Change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWalletsRequest) Reset()         { *m = GetWalletsRequest{} }
func (m *GetWalletsRequest) String() string { return proto.CompactTextString(m) }
func (*GetWalletsRequest) ProtoMessage()    {}
func (*GetWalletsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{11}
}
func (m *GetWalletsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWalletsRequest.Unmarshal(m, b)
}
func (m *GetWalletsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWalletsRequest.Marshal(b, m, deterministic)
}
func (dst *GetWalletsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletsRequest.Merge(dst, src)
}
func (m *GetWalletsRequest) XXX_Size() int {
	return xxx_messageInfo_GetWalletsRequest.Size(m)
}
func (m *GetWalletsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletsRequest proto.InternalMessageInfo

func (m *GetWalletsRequest) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

func (m *GetWalletsRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetWalletsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetWalletsRequest) GetBip44Change() int32 {
	if m != nil {
		return m.Bip44Change
	}
	return 0
}

type GetWalletsReply struct {
	ChainWithWallets     []*ChainWithWallets `protobuf:"bytes,1,rep,name=ChainWithWallets" json:"ChainWithWallets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetWalletsReply) Reset()         { *m = GetWalletsReply{} }
func (m *GetWalletsReply) String() string { return proto.CompactTextString(m) }
func (*GetWalletsReply) ProtoMessage()    {}
func (*GetWalletsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{12}
}
func (m *GetWalletsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWalletsReply.Unmarshal(m, b)
}
func (m *GetWalletsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWalletsReply.Marshal(b, m, deterministic)
}
func (dst *GetWalletsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWalletsReply.Merge(dst, src)
}
func (m *GetWalletsReply) XXX_Size() int {
	return xxx_messageInfo_GetWalletsReply.Size(m)
}
func (m *GetWalletsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWalletsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetWalletsReply proto.InternalMessageInfo

func (m *GetWalletsReply) GetChainWithWallets() []*ChainWithWallets {
	if m != nil {
		return m.ChainWithWallets
	}
	return nil
}

type QueryWalletsForGroupByChainNameRequest struct {
	Groupid              string   `protobuf:"bytes,1,opt,name=groupid" json:"groupid,omitempty"`
	ChainName            string   `protobuf:"bytes,2,opt,name=chainName" json:"chainName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryWalletsForGroupByChainNameRequest) Reset() {
	*m = QueryWalletsForGroupByChainNameRequest{}
}
func (m *QueryWalletsForGroupByChainNameRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWalletsForGroupByChainNameRequest) ProtoMessage()    {}
func (*QueryWalletsForGroupByChainNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{13}
}
func (m *QueryWalletsForGroupByChainNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryWalletsForGroupByChainNameRequest.Unmarshal(m, b)
}
func (m *QueryWalletsForGroupByChainNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryWalletsForGroupByChainNameRequest.Marshal(b, m, deterministic)
}
func (dst *QueryWalletsForGroupByChainNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWalletsForGroupByChainNameRequest.Merge(dst, src)
}
func (m *QueryWalletsForGroupByChainNameRequest) XXX_Size() int {
	return xxx_messageInfo_QueryWalletsForGroupByChainNameRequest.Size(m)
}
func (m *QueryWalletsForGroupByChainNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWalletsForGroupByChainNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWalletsForGroupByChainNameRequest proto.InternalMessageInfo

func (m *QueryWalletsForGroupByChainNameRequest) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

func (m *QueryWalletsForGroupByChainNameRequest) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

type QueryWalletsForGroupByChainNameReply struct {
	Wallets              []*Wallet `protobuf:"bytes,1,rep,name=Wallets" json:"Wallets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueryWalletsForGroupByChainNameReply) Reset()         { *m = QueryWalletsForGroupByChainNameReply{} }
func (m *QueryWalletsForGroupByChainNameReply) String() string { return proto.CompactTextString(m) }
func (*QueryWalletsForGroupByChainNameReply) ProtoMessage()    {}
func (*QueryWalletsForGroupByChainNameReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wallet_management_cbc3f86bb7f131e7, []int{14}
}
func (m *QueryWalletsForGroupByChainNameReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryWalletsForGroupByChainNameReply.Unmarshal(m, b)
}
func (m *QueryWalletsForGroupByChainNameReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryWalletsForGroupByChainNameReply.Marshal(b, m, deterministic)
}
func (dst *QueryWalletsForGroupByChainNameReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWalletsForGroupByChainNameReply.Merge(dst, src)
}
func (m *QueryWalletsForGroupByChainNameReply) XXX_Size() int {
	return xxx_messageInfo_QueryWalletsForGroupByChainNameReply.Size(m)
}
func (m *QueryWalletsForGroupByChainNameReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWalletsForGroupByChainNameReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWalletsForGroupByChainNameReply proto.InternalMessageInfo

func (m *QueryWalletsForGroupByChainNameReply) GetWallets() []*Wallet {
	if m != nil {
		return m.Wallets
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
	proto.RegisterType((*ChainWithWallets)(nil), "pb.ChainWithWallets")
	proto.RegisterType((*CreateWalletReply)(nil), "pb.CreateWalletReply")
	proto.RegisterType((*GetWalletsRequest)(nil), "pb.GetWalletsRequest")
	proto.RegisterType((*GetWalletsReply)(nil), "pb.GetWalletsReply")
	proto.RegisterType((*QueryWalletsForGroupByChainNameRequest)(nil), "pb.QueryWalletsForGroupByChainNameRequest")
	proto.RegisterType((*QueryWalletsForGroupByChainNameReply)(nil), "pb.QueryWalletsForGroupByChainNameReply")
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
	GetWallets(ctx context.Context, in *GetWalletsRequest, opts ...grpc.CallOption) (*GetWalletsReply, error)
	QueryWalletsForGroupByChainName(ctx context.Context, in *QueryWalletsForGroupByChainNameRequest, opts ...grpc.CallOption) (*QueryWalletsForGroupByChainNameReply, error)
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

func (c *walletManagementClient) GetWallets(ctx context.Context, in *GetWalletsRequest, opts ...grpc.CallOption) (*GetWalletsReply, error) {
	out := new(GetWalletsReply)
	err := c.cc.Invoke(ctx, "/pb.WalletManagement/GetWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletManagementClient) QueryWalletsForGroupByChainName(ctx context.Context, in *QueryWalletsForGroupByChainNameRequest, opts ...grpc.CallOption) (*QueryWalletsForGroupByChainNameReply, error) {
	out := new(QueryWalletsForGroupByChainNameReply)
	err := c.cc.Invoke(ctx, "/pb.WalletManagement/QueryWalletsForGroupByChainName", in, out, opts...)
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
	GetWallets(context.Context, *GetWalletsRequest) (*GetWalletsReply, error)
	QueryWalletsForGroupByChainName(context.Context, *QueryWalletsForGroupByChainNameRequest) (*QueryWalletsForGroupByChainNameReply, error)
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

func _WalletManagement_GetWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletManagementServer).GetWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletManagement/GetWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletManagementServer).GetWallets(ctx, req.(*GetWalletsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletManagement_QueryWalletsForGroupByChainName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWalletsForGroupByChainNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletManagementServer).QueryWalletsForGroupByChainName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.WalletManagement/QueryWalletsForGroupByChainName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletManagementServer).QueryWalletsForGroupByChainName(ctx, req.(*QueryWalletsForGroupByChainNameRequest))
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
		{
			MethodName: "GetWallets",
			Handler:    _WalletManagement_GetWallets_Handler,
		},
		{
			MethodName: "QueryWalletsForGroupByChainName",
			Handler:    _WalletManagement_QueryWalletsForGroupByChainName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet_management.proto",
}

func init() {
	proto.RegisterFile("wallet_management.proto", fileDescriptor_wallet_management_cbc3f86bb7f131e7)
}

var fileDescriptor_wallet_management_cbc3f86bb7f131e7 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0x45, 0xb6, 0xac, 0x24, 0xe3, 0x8f, 0xc4, 0xd9, 0xfc, 0x09, 0x11, 0xbe, 0x9a, 0x25, 0xb4,
	0x26, 0x17, 0xb9, 0x48, 0xd3, 0x3f, 0x0a, 0xa5, 0x69, 0x42, 0x43, 0xa0, 0x0d, 0x54, 0x0e, 0xa4,
	0xf4, 0xa6, 0x95, 0xac, 0xc5, 0x5e, 0xd0, 0x5f, 0xa5, 0x15, 0x41, 0x7d, 0x8b, 0xbe, 0x58, 0x9f,
	0xa9, 0xcc, 0xee, 0x4a, 0x96, 0xed, 0xa4, 0xf6, 0x9d, 0xce, 0xec, 0xec, 0x9c, 0x33, 0x67, 0x76,
	0x57, 0x70, 0x70, 0xef, 0x85, 0x21, 0x13, 0xdf, 0x23, 0x2f, 0xf6, 0xc6, 0x2c, 0x62, 0xb1, 0x38,
	0x49, 0xb3, 0x44, 0x24, 0xa4, 0x95, 0xfa, 0xf4, 0x1b, 0x90, 0x8b, 0x8c, 0x79, 0x82, 0x5d, 0x4c,
	0x3c, 0x1e, 0xbb, 0xec, 0x67, 0xc1, 0x72, 0x41, 0xf6, 0xc1, 0xca, 0xcb, 0xc8, 0x4f, 0x42, 0xdb,
	0xe8, 0x1b, 0x83, 0x0d, 0x57, 0x23, 0xb2, 0x0b, 0x1d, 0x9f, 0x0b, 0x1e, 0xd8, 0x2d, 0x19, 0x56,
	0x40, 0x66, 0x0b, 0x4f, 0x14, 0xb9, 0xdd, 0xee, 0x1b, 0x83, 0x75, 0x57, 0x23, 0x7a, 0x0c, 0xbd,
	0x99, 0xda, 0x69, 0x58, 0x62, 0x6e, 0xc6, 0xf2, 0x22, 0x14, 0xb2, 0xf2, 0xba, 0xab, 0x11, 0x7d,
	0x09, 0xce, 0x79, 0x9e, 0xf3, 0x71, 0xcc, 0x82, 0xaf, 0x69, 0xe1, 0xdf, 0x26, 0x57, 0x59, 0x52,
	0xa4, 0x95, 0x1e, 0x1b, 0xd6, 0xc6, 0x88, 0x79, 0xa0, 0x05, 0x55, 0x90, 0x3a, 0x60, 0x3f, 0xb8,
	0x2f, 0x0d, 0x4b, 0x4a, 0xa0, 0x77, 0xc5, 0x84, 0x24, 0xcf, 0x75, 0x25, 0xfa, 0xdb, 0x80, 0xee,
	0x90, 0x47, 0x69, 0xa8, 0x44, 0x91, 0x4d, 0x68, 0xd5, 0x45, 0x5b, 0x3c, 0x20, 0x04, 0xcc, 0xd8,
	0x8b, 0x98, 0x6e, 0x50, 0x7e, 0x23, 0xbb, 0xcf, 0xd3, 0xb3, 0x33, 0x1e, 0xc8, 0x06, 0x3b, 0x6e,
	0x05, 0x31, 0x7b, 0x94, 0xf0, 0xd8, 0x36, 0x55, 0x36, 0x7e, 0x37, 0xdc, 0xe8, 0x34, 0xdd, 0xc0,
	0x2a, 0x97, 0x6c, 0xc4, 0x23, 0x2f, 0xb4, 0xad, 0xbe, 0x31, 0x30, 0xdd, 0x0a, 0xd2, 0x37, 0xb0,
	0xd9, 0xd0, 0x89, 0x2e, 0x3d, 0x03, 0x6b, 0x24, 0xa1, 0x6d, 0xf4, 0xdb, 0x83, 0xee, 0xe9, 0xd6,
	0x49, 0xea, 0x9f, 0x34, 0x64, 0xbb, 0x7a, 0x99, 0x26, 0xb0, 0xa3, 0x2c, 0xbe, 0x93, 0x33, 0x5e,
	0xea, 0x17, 0x39, 0x84, 0x0d, 0xb9, 0xb5, 0xd1, 0xe4, 0x34, 0x40, 0xfa, 0xd0, 0x95, 0xad, 0x8d,
	0x26, 0x5e, 0x3c, 0x66, 0xba, 0xdb, 0x66, 0x88, 0x4e, 0xc0, 0x52, 0x54, 0xe8, 0xdc, 0xf5, 0x65,
	0xe5, 0xdc, 0xf5, 0x25, 0x72, 0x9e, 0x07, 0x41, 0xc6, 0xf2, 0x5c, 0xd7, 0xad, 0x20, 0x72, 0x4a,
	0xd5, 0x37, 0xc8, 0xd9, 0x56, 0x9c, 0x75, 0x00, 0xfd, 0x1a, 0x2a, 0xbf, 0x4c, 0xe5, 0x97, 0x42,
	0x54, 0x40, 0x4f, 0x26, 0xdd, 0x71, 0x31, 0x51, 0x94, 0x39, 0x39, 0x82, 0x35, 0xfd, 0xa9, 0x8d,
	0x01, 0x34, 0x46, 0xf7, 0x5e, 0x2d, 0x21, 0xdf, 0x6d, 0x22, 0xbc, 0x70, 0xc8, 0x7f, 0xa9, 0x1e,
	0x3b, 0xee, 0x34, 0xf0, 0x6f, 0x35, 0xf4, 0x15, 0x6c, 0xcf, 0x1a, 0x8a, 0xe3, 0xa0, 0x60, 0xa9,
	0x3b, 0x24, 0xdb, 0x9d, 0x65, 0xd5, 0x2b, 0xb4, 0x84, 0xed, 0x2b, 0x26, 0xb4, 0x84, 0xe5, 0x73,
	0x20, 0x60, 0xa6, 0xde, 0xb8, 0x92, 0x27, 0xbf, 0xf1, 0x76, 0x85, 0x3c, 0xe2, 0x42, 0xfb, 0xae,
	0x40, 0x3d, 0x93, 0x0b, 0x35, 0x13, 0xb3, 0x31, 0x13, 0x15, 0xa2, 0x43, 0xd8, 0x6a, 0x52, 0xa3,
	0xe2, 0xf7, 0x8b, 0xe6, 0x69, 0xc7, 0x76, 0x51, 0xfb, 0xfc, 0x9a, 0xbb, 0x90, 0x4d, 0x7f, 0xc0,
	0xd3, 0x2f, 0x05, 0xcb, 0x4a, 0x8d, 0x3f, 0x26, 0x99, 0xbc, 0x59, 0x1f, 0xca, 0xda, 0xab, 0xd5,
	0x0f, 0xdb, 0xcd, 0xfc, 0x61, 0x93, 0x56, 0x7f, 0x82, 0xa3, 0xa5, 0x0c, 0xd8, 0xcb, 0x4a, 0x43,
	0x3f, 0xfd, 0xd3, 0x86, 0x9e, 0xfa, 0xfe, 0x5c, 0xbf, 0x73, 0xe4, 0x2d, 0x74, 0x1b, 0x2f, 0x10,
	0xd9, 0x97, 0xbd, 0x2f, 0x3c, 0x77, 0xce, 0xee, 0x42, 0x1c, 0x79, 0x87, 0xb0, 0xf3, 0xc0, 0xd3,
	0x42, 0xfe, 0xc7, 0xe4, 0xc7, 0xdf, 0x2a, 0xe7, 0xf0, 0xd1, 0x75, 0x2c, 0xfa, 0x02, 0x36, 0xea,
	0xbb, 0x4e, 0x24, 0xef, 0xfc, 0x13, 0xe5, 0x90, 0xb9, 0x28, 0x6e, 0x7b, 0x07, 0xff, 0x35, 0x8f,
	0x25, 0x39, 0x98, 0x2a, 0x9e, 0xb9, 0xf9, 0xce, 0xde, 0xe2, 0x02, 0xee, 0x7f, 0x0d, 0x30, 0x3d,
	0x22, 0x64, 0x4f, 0x33, 0xcc, 0x9e, 0x56, 0x67, 0x67, 0x3e, 0x8c, 0x3b, 0xef, 0xe1, 0xc9, 0x92,
	0x29, 0x91, 0x63, 0xdc, 0xb7, 0xda, 0x61, 0x71, 0x06, 0x2b, 0xe5, 0xa6, 0x61, 0xe9, 0x5b, 0xf2,
	0x27, 0xf5, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x06, 0x83, 0x77, 0xbf, 0x06, 0x00,
	0x00,
}
