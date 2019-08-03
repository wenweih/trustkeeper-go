// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dashboard.proto

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

type GetGroupsRequest struct {
	NamespaceID          string   `protobuf:"bytes,1,opt,name=namespaceID" json:"namespaceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupsRequest) Reset()         { *m = GetGroupsRequest{} }
func (m *GetGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupsRequest) ProtoMessage()    {}
func (*GetGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{0}
}
func (m *GetGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupsRequest.Unmarshal(m, b)
}
func (m *GetGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupsRequest.Marshal(b, m, deterministic)
}
func (dst *GetGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupsRequest.Merge(dst, src)
}
func (m *GetGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupsRequest.Size(m)
}
func (m *GetGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupsRequest proto.InternalMessageInfo

func (m *GetGroupsRequest) GetNamespaceID() string {
	if m != nil {
		return m.NamespaceID
	}
	return ""
}

type GetGroupsReply struct {
	Groups               []*Group `protobuf:"bytes,1,rep,name=groups" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupsReply) Reset()         { *m = GetGroupsReply{} }
func (m *GetGroupsReply) String() string { return proto.CompactTextString(m) }
func (*GetGroupsReply) ProtoMessage()    {}
func (*GetGroupsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{1}
}
func (m *GetGroupsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupsReply.Unmarshal(m, b)
}
func (m *GetGroupsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupsReply.Marshal(b, m, deterministic)
}
func (dst *GetGroupsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupsReply.Merge(dst, src)
}
func (m *GetGroupsReply) XXX_Size() int {
	return xxx_messageInfo_GetGroupsReply.Size(m)
}
func (m *GetGroupsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupsReply proto.InternalMessageInfo

func (m *GetGroupsReply) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type Group struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc" json:"desc,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{2}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (dst *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(dst, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Group) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateGroupRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	NamespaceID          string   `protobuf:"bytes,4,opt,name=namespaceID" json:"namespaceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupRequest) Reset()         { *m = CreateGroupRequest{} }
func (m *CreateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGroupRequest) ProtoMessage()    {}
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{3}
}
func (m *CreateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupRequest.Unmarshal(m, b)
}
func (m *CreateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupRequest.Marshal(b, m, deterministic)
}
func (dst *CreateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupRequest.Merge(dst, src)
}
func (m *CreateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGroupRequest.Size(m)
}
func (m *CreateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupRequest proto.InternalMessageInfo

func (m *CreateGroupRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *CreateGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateGroupRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *CreateGroupRequest) GetNamespaceID() string {
	if m != nil {
		return m.NamespaceID
	}
	return ""
}

type CreateGroupReply struct {
	Group                *Group   `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGroupReply) Reset()         { *m = CreateGroupReply{} }
func (m *CreateGroupReply) String() string { return proto.CompactTextString(m) }
func (*CreateGroupReply) ProtoMessage()    {}
func (*CreateGroupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{4}
}
func (m *CreateGroupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGroupReply.Unmarshal(m, b)
}
func (m *CreateGroupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGroupReply.Marshal(b, m, deterministic)
}
func (dst *CreateGroupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGroupReply.Merge(dst, src)
}
func (m *CreateGroupReply) XXX_Size() int {
	return xxx_messageInfo_CreateGroupReply.Size(m)
}
func (m *CreateGroupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGroupReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGroupReply proto.InternalMessageInfo

func (m *CreateGroupReply) GetGroup() *Group {
	if m != nil {
		return m.Group
	}
	return nil
}

type UpdateGroupRequest struct {
	Groupid              string   `protobuf:"bytes,1,opt,name=groupid" json:"groupid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupRequest) Reset()         { *m = UpdateGroupRequest{} }
func (m *UpdateGroupRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupRequest) ProtoMessage()    {}
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{5}
}
func (m *UpdateGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupRequest.Unmarshal(m, b)
}
func (m *UpdateGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupRequest.Merge(dst, src)
}
func (m *UpdateGroupRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupRequest.Size(m)
}
func (m *UpdateGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupRequest proto.InternalMessageInfo

func (m *UpdateGroupRequest) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

func (m *UpdateGroupRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateGroupRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type UpdateGroupReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateGroupReply) Reset()         { *m = UpdateGroupReply{} }
func (m *UpdateGroupReply) String() string { return proto.CompactTextString(m) }
func (*UpdateGroupReply) ProtoMessage()    {}
func (*UpdateGroupReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{6}
}
func (m *UpdateGroupReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGroupReply.Unmarshal(m, b)
}
func (m *UpdateGroupReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGroupReply.Marshal(b, m, deterministic)
}
func (dst *UpdateGroupReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGroupReply.Merge(dst, src)
}
func (m *UpdateGroupReply) XXX_Size() int {
	return xxx_messageInfo_UpdateGroupReply.Size(m)
}
func (m *UpdateGroupReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGroupReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGroupReply proto.InternalMessageInfo

type GetGroupAssetRequest struct {
	Groupid              string   `protobuf:"bytes,1,opt,name=groupid" json:"groupid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupAssetRequest) Reset()         { *m = GetGroupAssetRequest{} }
func (m *GetGroupAssetRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupAssetRequest) ProtoMessage()    {}
func (*GetGroupAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{7}
}
func (m *GetGroupAssetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupAssetRequest.Unmarshal(m, b)
}
func (m *GetGroupAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupAssetRequest.Marshal(b, m, deterministic)
}
func (dst *GetGroupAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupAssetRequest.Merge(dst, src)
}
func (m *GetGroupAssetRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupAssetRequest.Size(m)
}
func (m *GetGroupAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupAssetRequest proto.InternalMessageInfo

func (m *GetGroupAssetRequest) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

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
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{8}
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

type ChainAsset struct {
	ChainID              string         `protobuf:"bytes,1,opt,name=chainID" json:"chainID,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Coin                 string         `protobuf:"bytes,3,opt,name=coin" json:"coin,omitempty"`
	Status               bool           `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	Decimal              uint64         `protobuf:"varint,5,opt,name=Decimal" json:"Decimal,omitempty"`
	SimpleAssets         []*SimpleAsset `protobuf:"bytes,6,rep,name=SimpleAssets" json:"SimpleAssets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChainAsset) Reset()         { *m = ChainAsset{} }
func (m *ChainAsset) String() string { return proto.CompactTextString(m) }
func (*ChainAsset) ProtoMessage()    {}
func (*ChainAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{9}
}
func (m *ChainAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainAsset.Unmarshal(m, b)
}
func (m *ChainAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainAsset.Marshal(b, m, deterministic)
}
func (dst *ChainAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainAsset.Merge(dst, src)
}
func (m *ChainAsset) XXX_Size() int {
	return xxx_messageInfo_ChainAsset.Size(m)
}
func (m *ChainAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainAsset.DiscardUnknown(m)
}

var xxx_messageInfo_ChainAsset proto.InternalMessageInfo

func (m *ChainAsset) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *ChainAsset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChainAsset) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

func (m *ChainAsset) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *ChainAsset) GetDecimal() uint64 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *ChainAsset) GetSimpleAssets() []*SimpleAsset {
	if m != nil {
		return m.SimpleAssets
	}
	return nil
}

type GetGroupAssetReply struct {
	Chainassets          []*ChainAsset `protobuf:"bytes,1,rep,name=chainassets" json:"chainassets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetGroupAssetReply) Reset()         { *m = GetGroupAssetReply{} }
func (m *GetGroupAssetReply) String() string { return proto.CompactTextString(m) }
func (*GetGroupAssetReply) ProtoMessage()    {}
func (*GetGroupAssetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{10}
}
func (m *GetGroupAssetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupAssetReply.Unmarshal(m, b)
}
func (m *GetGroupAssetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupAssetReply.Marshal(b, m, deterministic)
}
func (dst *GetGroupAssetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupAssetReply.Merge(dst, src)
}
func (m *GetGroupAssetReply) XXX_Size() int {
	return xxx_messageInfo_GetGroupAssetReply.Size(m)
}
func (m *GetGroupAssetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupAssetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupAssetReply proto.InternalMessageInfo

func (m *GetGroupAssetReply) GetChainassets() []*ChainAsset {
	if m != nil {
		return m.Chainassets
	}
	return nil
}

type ChangeGroupAssetsRequest struct {
	Chainassets          []*ChainAsset `protobuf:"bytes,1,rep,name=chainassets" json:"chainassets,omitempty"`
	Groupid              string        `protobuf:"bytes,2,opt,name=groupid" json:"groupid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ChangeGroupAssetsRequest) Reset()         { *m = ChangeGroupAssetsRequest{} }
func (m *ChangeGroupAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeGroupAssetsRequest) ProtoMessage()    {}
func (*ChangeGroupAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{11}
}
func (m *ChangeGroupAssetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeGroupAssetsRequest.Unmarshal(m, b)
}
func (m *ChangeGroupAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeGroupAssetsRequest.Marshal(b, m, deterministic)
}
func (dst *ChangeGroupAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeGroupAssetsRequest.Merge(dst, src)
}
func (m *ChangeGroupAssetsRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeGroupAssetsRequest.Size(m)
}
func (m *ChangeGroupAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeGroupAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeGroupAssetsRequest proto.InternalMessageInfo

func (m *ChangeGroupAssetsRequest) GetChainassets() []*ChainAsset {
	if m != nil {
		return m.Chainassets
	}
	return nil
}

func (m *ChangeGroupAssetsRequest) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

type ChangeGroupAssetsReply struct {
	Chainassets          []*ChainAsset `protobuf:"bytes,1,rep,name=chainassets" json:"chainassets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ChangeGroupAssetsReply) Reset()         { *m = ChangeGroupAssetsReply{} }
func (m *ChangeGroupAssetsReply) String() string { return proto.CompactTextString(m) }
func (*ChangeGroupAssetsReply) ProtoMessage()    {}
func (*ChangeGroupAssetsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{12}
}
func (m *ChangeGroupAssetsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeGroupAssetsReply.Unmarshal(m, b)
}
func (m *ChangeGroupAssetsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeGroupAssetsReply.Marshal(b, m, deterministic)
}
func (dst *ChangeGroupAssetsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeGroupAssetsReply.Merge(dst, src)
}
func (m *ChangeGroupAssetsReply) XXX_Size() int {
	return xxx_messageInfo_ChangeGroupAssetsReply.Size(m)
}
func (m *ChangeGroupAssetsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeGroupAssetsReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeGroupAssetsReply proto.InternalMessageInfo

func (m *ChangeGroupAssetsReply) GetChainassets() []*ChainAsset {
	if m != nil {
		return m.Chainassets
	}
	return nil
}

type AddAssetRequest struct {
	Groupid              string   `protobuf:"bytes,1,opt,name=Groupid" json:"Groupid,omitempty"`
	Chainid              string   `protobuf:"bytes,2,opt,name=Chainid" json:"Chainid,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=Symbol" json:"Symbol,omitempty"`
	Identify             string   `protobuf:"bytes,4,opt,name=Identify" json:"Identify,omitempty"`
	Decimal              string   `protobuf:"bytes,5,opt,name=Decimal" json:"Decimal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAssetRequest) Reset()         { *m = AddAssetRequest{} }
func (m *AddAssetRequest) String() string { return proto.CompactTextString(m) }
func (*AddAssetRequest) ProtoMessage()    {}
func (*AddAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{13}
}
func (m *AddAssetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAssetRequest.Unmarshal(m, b)
}
func (m *AddAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAssetRequest.Marshal(b, m, deterministic)
}
func (dst *AddAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAssetRequest.Merge(dst, src)
}
func (m *AddAssetRequest) XXX_Size() int {
	return xxx_messageInfo_AddAssetRequest.Size(m)
}
func (m *AddAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddAssetRequest proto.InternalMessageInfo

func (m *AddAssetRequest) GetGroupid() string {
	if m != nil {
		return m.Groupid
	}
	return ""
}

func (m *AddAssetRequest) GetChainid() string {
	if m != nil {
		return m.Chainid
	}
	return ""
}

func (m *AddAssetRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *AddAssetRequest) GetIdentify() string {
	if m != nil {
		return m.Identify
	}
	return ""
}

func (m *AddAssetRequest) GetDecimal() string {
	if m != nil {
		return m.Decimal
	}
	return ""
}

type AddAssetReply struct {
	Asset                *SimpleAsset `protobuf:"bytes,1,opt,name=Asset" json:"Asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AddAssetReply) Reset()         { *m = AddAssetReply{} }
func (m *AddAssetReply) String() string { return proto.CompactTextString(m) }
func (*AddAssetReply) ProtoMessage()    {}
func (*AddAssetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_1aef1bd888682fdc, []int{14}
}
func (m *AddAssetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAssetReply.Unmarshal(m, b)
}
func (m *AddAssetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAssetReply.Marshal(b, m, deterministic)
}
func (dst *AddAssetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAssetReply.Merge(dst, src)
}
func (m *AddAssetReply) XXX_Size() int {
	return xxx_messageInfo_AddAssetReply.Size(m)
}
func (m *AddAssetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAssetReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddAssetReply proto.InternalMessageInfo

func (m *AddAssetReply) GetAsset() *SimpleAsset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func init() {
	proto.RegisterType((*GetGroupsRequest)(nil), "pb.GetGroupsRequest")
	proto.RegisterType((*GetGroupsReply)(nil), "pb.GetGroupsReply")
	proto.RegisterType((*Group)(nil), "pb.Group")
	proto.RegisterType((*CreateGroupRequest)(nil), "pb.CreateGroupRequest")
	proto.RegisterType((*CreateGroupReply)(nil), "pb.CreateGroupReply")
	proto.RegisterType((*UpdateGroupRequest)(nil), "pb.UpdateGroupRequest")
	proto.RegisterType((*UpdateGroupReply)(nil), "pb.UpdateGroupReply")
	proto.RegisterType((*GetGroupAssetRequest)(nil), "pb.GetGroupAssetRequest")
	proto.RegisterType((*SimpleAsset)(nil), "pb.SimpleAsset")
	proto.RegisterType((*ChainAsset)(nil), "pb.ChainAsset")
	proto.RegisterType((*GetGroupAssetReply)(nil), "pb.GetGroupAssetReply")
	proto.RegisterType((*ChangeGroupAssetsRequest)(nil), "pb.ChangeGroupAssetsRequest")
	proto.RegisterType((*ChangeGroupAssetsReply)(nil), "pb.ChangeGroupAssetsReply")
	proto.RegisterType((*AddAssetRequest)(nil), "pb.AddAssetRequest")
	proto.RegisterType((*AddAssetReply)(nil), "pb.AddAssetReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DashboardClient is the client API for Dashboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DashboardClient interface {
	GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsReply, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupReply, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupReply, error)
	GetGroupAsset(ctx context.Context, in *GetGroupAssetRequest, opts ...grpc.CallOption) (*GetGroupAssetReply, error)
	ChangeGroupAssets(ctx context.Context, in *ChangeGroupAssetsRequest, opts ...grpc.CallOption) (*ChangeGroupAssetsReply, error)
	AddAsset(ctx context.Context, in *AddAssetRequest, opts ...grpc.CallOption) (*AddAssetReply, error)
}

type dashboardClient struct {
	cc *grpc.ClientConn
}

func NewDashboardClient(cc *grpc.ClientConn) DashboardClient {
	return &dashboardClient{cc}
}

func (c *dashboardClient) GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsReply, error) {
	out := new(GetGroupsReply)
	err := c.cc.Invoke(ctx, "/pb.Dashboard/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupReply, error) {
	out := new(CreateGroupReply)
	err := c.cc.Invoke(ctx, "/pb.Dashboard/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupReply, error) {
	out := new(UpdateGroupReply)
	err := c.cc.Invoke(ctx, "/pb.Dashboard/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetGroupAsset(ctx context.Context, in *GetGroupAssetRequest, opts ...grpc.CallOption) (*GetGroupAssetReply, error) {
	out := new(GetGroupAssetReply)
	err := c.cc.Invoke(ctx, "/pb.Dashboard/GetGroupAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) ChangeGroupAssets(ctx context.Context, in *ChangeGroupAssetsRequest, opts ...grpc.CallOption) (*ChangeGroupAssetsReply, error) {
	out := new(ChangeGroupAssetsReply)
	err := c.cc.Invoke(ctx, "/pb.Dashboard/ChangeGroupAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) AddAsset(ctx context.Context, in *AddAssetRequest, opts ...grpc.CallOption) (*AddAssetReply, error) {
	out := new(AddAssetReply)
	err := c.cc.Invoke(ctx, "/pb.Dashboard/AddAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServer is the server API for Dashboard service.
type DashboardServer interface {
	GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsReply, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupReply, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*UpdateGroupReply, error)
	GetGroupAsset(context.Context, *GetGroupAssetRequest) (*GetGroupAssetReply, error)
	ChangeGroupAssets(context.Context, *ChangeGroupAssetsRequest) (*ChangeGroupAssetsReply, error)
	AddAsset(context.Context, *AddAssetRequest) (*AddAssetReply, error)
}

func RegisterDashboardServer(s *grpc.Server, srv DashboardServer) {
	s.RegisterService(&_Dashboard_serviceDesc, srv)
}

func _Dashboard_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Dashboard/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).GetGroups(ctx, req.(*GetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Dashboard/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Dashboard/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_GetGroupAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).GetGroupAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Dashboard/GetGroupAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).GetGroupAsset(ctx, req.(*GetGroupAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_ChangeGroupAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeGroupAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).ChangeGroupAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Dashboard/ChangeGroupAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).ChangeGroupAssets(ctx, req.(*ChangeGroupAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_AddAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).AddAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Dashboard/AddAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).AddAsset(ctx, req.(*AddAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dashboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Dashboard",
	HandlerType: (*DashboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroups",
			Handler:    _Dashboard_GetGroups_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _Dashboard_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _Dashboard_UpdateGroup_Handler,
		},
		{
			MethodName: "GetGroupAsset",
			Handler:    _Dashboard_GetGroupAsset_Handler,
		},
		{
			MethodName: "ChangeGroupAssets",
			Handler:    _Dashboard_ChangeGroupAssets_Handler,
		},
		{
			MethodName: "AddAsset",
			Handler:    _Dashboard_AddAsset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dashboard.proto",
}

func init() { proto.RegisterFile("dashboard.proto", fileDescriptor_dashboard_1aef1bd888682fdc) }

var fileDescriptor_dashboard_1aef1bd888682fdc = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x8f, 0xda, 0x30,
	0x10, 0x55, 0xc2, 0x47, 0x97, 0xa1, 0x0b, 0xec, 0x14, 0xa1, 0x28, 0xaa, 0x54, 0x6a, 0xa9, 0x12,
	0x27, 0xb4, 0x82, 0xb6, 0x97, 0x1e, 0x2a, 0x04, 0x2a, 0xda, 0x4a, 0xbd, 0x64, 0xd5, 0xde, 0x03,
	0xf6, 0x2e, 0x91, 0x02, 0x71, 0x71, 0x38, 0x70, 0xed, 0xbd, 0xa7, 0xfe, 0x96, 0xfe, 0xbf, 0xca,
	0x1f, 0x49, 0x9c, 0x00, 0x6a, 0xb7, 0xb7, 0xbc, 0xb1, 0xfd, 0xe6, 0x8d, 0xe7, 0x79, 0x02, 0x5d,
	0x1a, 0x8a, 0xcd, 0x2a, 0x09, 0xf7, 0x74, 0xcc, 0xf7, 0x49, 0x9a, 0xa0, 0xcb, 0x57, 0xe4, 0x2d,
	0xf4, 0x96, 0x2c, 0x5d, 0xee, 0x93, 0x03, 0x17, 0x01, 0xfb, 0x7e, 0x60, 0x22, 0xc5, 0x21, 0xb4,
	0x77, 0xe1, 0x96, 0x09, 0x1e, 0xae, 0xd9, 0xdd, 0xc2, 0x73, 0x86, 0xce, 0xa8, 0x15, 0xd8, 0x21,
	0x32, 0x85, 0x8e, 0x75, 0x8a, 0xc7, 0x47, 0x7c, 0x0d, 0xcd, 0x47, 0x05, 0x3d, 0x67, 0x58, 0x1b,
	0xb5, 0x27, 0xad, 0x31, 0x5f, 0x8d, 0xd5, 0x86, 0xc0, 0x2c, 0x90, 0x8f, 0xd0, 0x50, 0x01, 0x44,
	0xa8, 0x4b, 0x32, 0x43, 0xac, 0xbe, 0x65, 0x8c, 0x32, 0xb1, 0xf6, 0x5c, 0x1d, 0x93, 0xdf, 0xd8,
	0x01, 0x37, 0xa2, 0x5e, 0x4d, 0x45, 0xdc, 0x88, 0x92, 0x3d, 0xe0, 0x7c, 0xcf, 0xc2, 0x94, 0x69,
	0x5e, 0xa3, 0x16, 0xa1, 0x7e, 0x38, 0x44, 0x34, 0x63, 0x93, 0xdf, 0x79, 0x06, 0xf7, 0x4c, 0x86,
	0x9a, 0x95, 0xa1, 0x52, 0x69, 0xfd, 0x5c, 0xa5, 0xbd, 0x52, 0x4e, 0x59, 0xeb, 0x2b, 0x68, 0xa8,
	0x92, 0x54, 0xca, 0x52, 0xa9, 0x3a, 0x4e, 0xbe, 0x01, 0x7e, 0xe5, 0xb4, 0x2a, 0xd4, 0x83, 0x67,
	0x6a, 0x39, 0xd7, 0x9a, 0xc1, 0x7f, 0x95, 0x4b, 0x10, 0x7a, 0x25, 0x5e, 0x1e, 0x1f, 0xc9, 0x2d,
	0xf4, 0xb3, 0x56, 0xcc, 0x84, 0x60, 0xe9, 0x5f, 0xb3, 0x91, 0x9f, 0x0e, 0xb4, 0xef, 0xa3, 0x2d,
	0x8f, 0x99, 0x3a, 0x20, 0x77, 0xaa, 0x8f, 0xbc, 0xd5, 0x19, 0xc4, 0x01, 0x34, 0xc5, 0x71, 0xbb,
	0x4a, 0x62, 0xa3, 0xcc, 0x20, 0x15, 0x4f, 0xc3, 0xf4, 0x20, 0x94, 0xba, 0xab, 0xc0, 0x20, 0xf4,
	0xe1, 0xea, 0x8e, 0xb2, 0x5d, 0x1a, 0x3d, 0x1c, 0xcd, 0x5d, 0xe6, 0x58, 0x66, 0x59, 0xb0, 0x75,
	0xb4, 0x0d, 0x63, 0xaf, 0x31, 0x74, 0x46, 0xf5, 0x20, 0x83, 0xe4, 0xb7, 0x03, 0x30, 0xdf, 0x84,
	0xd1, 0x2e, 0x97, 0xb3, 0x96, 0xa8, 0x90, 0x63, 0xe0, 0xa5, 0x6b, 0x5a, 0x27, 0xd1, 0x2e, 0xbb,
	0x26, 0xf9, 0x6d, 0xc9, 0xab, 0x97, 0xe4, 0x5d, 0x94, 0x80, 0x53, 0x78, 0x6e, 0xdd, 0x88, 0xf0,
	0x9a, 0xca, 0xc3, 0x5d, 0xd9, 0x58, 0x2b, 0x1e, 0x94, 0x36, 0x91, 0x4f, 0x80, 0x95, 0x9b, 0x97,
	0xe6, 0xb8, 0x85, 0xb6, 0xd2, 0x1b, 0x6a, 0x26, 0xfd, 0x1a, 0x3a, 0x92, 0xa9, 0xa8, 0x31, 0xb0,
	0xb7, 0x90, 0x07, 0xf0, 0xe6, 0x9b, 0x70, 0xf7, 0xc8, 0x0a, 0xaa, 0xfc, 0x29, 0x3e, 0x99, 0xcd,
	0xee, 0xbb, 0x5b, 0xee, 0xfb, 0x67, 0x18, 0x9c, 0xc9, 0xf3, 0x7f, 0x9a, 0x7f, 0x39, 0xd0, 0x9d,
	0x51, 0x5a, 0x75, 0xdc, 0xb2, 0xec, 0x38, 0x03, 0xe5, 0x8a, 0x22, 0x2a, 0x34, 0x19, 0x28, 0x5b,
	0x75, 0xaf, 0x1d, 0xa6, 0x1b, 0x68, 0xd0, 0x53, 0x9c, 0xd4, 0x2a, 0x9c, 0xf4, 0x1e, 0xae, 0x0b,
	0x51, 0xb2, 0xb0, 0x37, 0xd0, 0x50, 0xc8, 0xbc, 0xd4, 0x93, 0x86, 0xea, 0xd5, 0xc9, 0x8f, 0x1a,
	0xb4, 0x16, 0xd9, 0x70, 0xc4, 0x77, 0xd0, 0xca, 0x87, 0x1b, 0xf6, 0xd5, 0xe3, 0xae, 0x4c, 0x48,
	0x1f, 0x2b, 0x51, 0x99, 0xeb, 0x03, 0xb4, 0xad, 0x49, 0x81, 0x03, 0x75, 0x7d, 0x27, 0xe3, 0xca,
	0xef, 0x9f, 0xc4, 0xcd, 0x61, 0xeb, 0x65, 0xeb, 0xc3, 0xa7, 0x23, 0x44, 0x1f, 0xae, 0x8e, 0x00,
	0x9c, 0xc1, 0x75, 0xc9, 0x88, 0xe8, 0xd9, 0xf2, 0xec, 0x1e, 0xf9, 0x83, 0x33, 0x2b, 0x92, 0xe2,
	0x0b, 0xdc, 0x9c, 0x78, 0x03, 0x5f, 0x1a, 0x07, 0x9c, 0xb5, 0xa6, 0xef, 0x5f, 0x58, 0x95, 0x74,
	0x13, 0xb8, 0xca, 0x1a, 0x81, 0x2f, 0xe4, 0xbe, 0x8a, 0x57, 0xfc, 0x9b, 0x72, 0x90, 0xc7, 0xc7,
	0x55, 0x53, 0xfd, 0x94, 0xa6, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x72, 0x59, 0x50, 0xa7,
	0x06, 0x00, 0x00,
}
