// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	PushTxRequest
	PushTxResponse
	RegisterRequest
	AccountInfo
	UserInfo
	RegisterResponse
	LoginRequest
	LoginResponse
	GetBlockHeaderRequest
	GetBlockHeaderResponse
	BlockHeader
	GetAccountInfoRequest
	GetAccountInfoResponse
	AccountInfoData
	FavoriteRequest
	FavoriteResponse
	GetFavoriteRequest
	GetFavoriteResponse
	FavoriteArr
	FavoriteData
	GetBalanceRequest
	GetBalanceResponse
	QueryMyBuyRequest
	QueryMyBuyResponse
	BuyData
	Buy
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PushTxRequest struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

func (m *PushTxRequest) Reset()                    { *m = PushTxRequest{} }
func (m *PushTxRequest) String() string            { return proto.CompactTextString(m) }
func (*PushTxRequest) ProtoMessage()               {}
func (*PushTxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PushTxRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PushTxRequest) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *PushTxRequest) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *PushTxRequest) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *PushTxRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PushTxRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *PushTxRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *PushTxRequest) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *PushTxRequest) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *PushTxRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type PushTxResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
}

func (m *PushTxResponse) Reset()                    { *m = PushTxResponse{} }
func (m *PushTxResponse) String() string            { return proto.CompactTextString(m) }
func (*PushTxResponse) ProtoMessage()               {}
func (*PushTxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PushTxResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PushTxResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type RegisterRequest struct {
	Account     *AccountInfo `protobuf:"bytes,1,opt,name=account" json:"account"`
	User        *UserInfo    `protobuf:"bytes,2,opt,name=user" json:"user"`
	VerifyId    string       `protobuf:"bytes,3,opt,name=verify_id,json=verifyId" json:"verify_id"`
	VerifyValue string       `protobuf:"bytes,4,opt,name=verify_value,json=verifyValue" json:"verify_value"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterRequest) GetAccount() *AccountInfo {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *RegisterRequest) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RegisterRequest) GetVerifyId() string {
	if m != nil {
		return m.VerifyId
	}
	return ""
}

func (m *RegisterRequest) GetVerifyValue() string {
	if m != nil {
		return m.VerifyValue
	}
	return ""
}

type AccountInfo struct {
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name"`
	Pubkey string `protobuf:"bytes,3,opt,name=pubkey" json:"pubkey"`
}

func (m *AccountInfo) Reset()                    { *m = AccountInfo{} }
func (m *AccountInfo) String() string            { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()               {}
func (*AccountInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AccountInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountInfo) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

type UserInfo struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserInfo) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *UserInfo) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *UserInfo) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *UserInfo) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *UserInfo) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *UserInfo) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *UserInfo) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *UserInfo) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *UserInfo) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *UserInfo) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type RegisterResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
}

func (m *RegisterResponse) Reset()                    { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()               {}
func (*RegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegisterResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LoginRequest struct {
	Username    string `protobuf:"bytes,1,opt,name=username" json:"username"`
	Random      string `protobuf:"bytes,2,opt,name=random" json:"random"`
	VerifyId    string `protobuf:"bytes,3,opt,name=verify_id,json=verifyId" json:"verify_id"`
	VerifyValue string `protobuf:"bytes,4,opt,name=verify_value,json=verifyValue" json:"verify_value"`
	Signture    string `protobuf:"bytes,5,opt,name=signture" json:"signture"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func (m *LoginRequest) GetVerifyId() string {
	if m != nil {
		return m.VerifyId
	}
	return ""
}

func (m *LoginRequest) GetVerifyValue() string {
	if m != nil {
		return m.VerifyValue
	}
	return ""
}

func (m *LoginRequest) GetSignture() string {
	if m != nil {
		return m.Signture
	}
	return ""
}

type LoginResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LoginResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetBlockHeaderRequest struct {
}

func (m *GetBlockHeaderRequest) Reset()                    { *m = GetBlockHeaderRequest{} }
func (m *GetBlockHeaderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockHeaderRequest) ProtoMessage()               {}
func (*GetBlockHeaderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type GetBlockHeaderResponse struct {
	Code uint32       `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *BlockHeader `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string       `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetBlockHeaderResponse) Reset()                    { *m = GetBlockHeaderResponse{} }
func (m *GetBlockHeaderResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlockHeaderResponse) ProtoMessage()               {}
func (*GetBlockHeaderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetBlockHeaderResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetBlockHeaderResponse) GetData() *BlockHeader {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetBlockHeaderResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type BlockHeader struct {
	HeadBlockNum          uint32 `protobuf:"varint,1,opt,name=head_block_num,json=headBlockNum" json:"head_block_num"`
	HeadBlockHash         string `protobuf:"bytes,2,opt,name=head_block_hash,json=headBlockHash" json:"head_block_hash"`
	HeadBlockTime         uint64 `protobuf:"varint,3,opt,name=head_block_time,json=headBlockTime" json:"head_block_time"`
	HeadBlockDelegate     string `protobuf:"bytes,4,opt,name=head_block_delegate,json=headBlockDelegate" json:"head_block_delegate"`
	CursorLabel           uint32 `protobuf:"varint,5,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	LastConsensusBlockNum uint32 `protobuf:"varint,6,opt,name=last_consensus_block_num,json=lastConsensusBlockNum" json:"last_consensus_block_num"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BlockHeader) GetHeadBlockNum() uint32 {
	if m != nil {
		return m.HeadBlockNum
	}
	return 0
}

func (m *BlockHeader) GetHeadBlockHash() string {
	if m != nil {
		return m.HeadBlockHash
	}
	return ""
}

func (m *BlockHeader) GetHeadBlockTime() uint64 {
	if m != nil {
		return m.HeadBlockTime
	}
	return 0
}

func (m *BlockHeader) GetHeadBlockDelegate() string {
	if m != nil {
		return m.HeadBlockDelegate
	}
	return ""
}

func (m *BlockHeader) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *BlockHeader) GetLastConsensusBlockNum() uint32 {
	if m != nil {
		return m.LastConsensusBlockNum
	}
	return 0
}

type GetAccountInfoRequest struct {
	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName" json:"account_name"`
}

func (m *GetAccountInfoRequest) Reset()                    { *m = GetAccountInfoRequest{} }
func (m *GetAccountInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountInfoRequest) ProtoMessage()               {}
func (*GetAccountInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetAccountInfoRequest) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

type GetAccountInfoResponse struct {
	Code uint32           `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *AccountInfoData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string           `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetAccountInfoResponse) Reset()                    { *m = GetAccountInfoResponse{} }
func (m *GetAccountInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAccountInfoResponse) ProtoMessage()               {}
func (*GetAccountInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetAccountInfoResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetAccountInfoResponse) GetData() *AccountInfoData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetAccountInfoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type AccountInfoData struct {
	AccountName   string `protobuf:"bytes,1,opt,name=account_name,json=accountName" json:"account_name"`
	Pubkey        string `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey"`
	Balance       uint64 `protobuf:"varint,3,opt,name=balance" json:"balance"`
	StakedBalance uint64 `protobuf:"varint,4,opt,name=staked_balance,json=stakedBalance" json:"staked_balance"`
}

func (m *AccountInfoData) Reset()                    { *m = AccountInfoData{} }
func (m *AccountInfoData) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoData) ProtoMessage()               {}
func (*AccountInfoData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AccountInfoData) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *AccountInfoData) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *AccountInfoData) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountInfoData) GetStakedBalance() uint64 {
	if m != nil {
		return m.StakedBalance
	}
	return 0
}

type FavoriteRequest struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

func (m *FavoriteRequest) Reset()                    { *m = FavoriteRequest{} }
func (m *FavoriteRequest) String() string            { return proto.CompactTextString(m) }
func (*FavoriteRequest) ProtoMessage()               {}
func (*FavoriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *FavoriteRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FavoriteRequest) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *FavoriteRequest) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *FavoriteRequest) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *FavoriteRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *FavoriteRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *FavoriteRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *FavoriteRequest) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *FavoriteRequest) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *FavoriteRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type FavoriteResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *FavoriteResponse) Reset()                    { *m = FavoriteResponse{} }
func (m *FavoriteResponse) String() string            { return proto.CompactTextString(m) }
func (*FavoriteResponse) ProtoMessage()               {}
func (*FavoriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *FavoriteResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *FavoriteResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *FavoriteResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetFavoriteRequest struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username"`
	GoodsType string `protobuf:"bytes,2,opt,name=goods_type,json=goodsType" json:"goods_type"`
	PageSize  uint32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size"`
	PageNum   uint32 `protobuf:"varint,4,opt,name=page_num,json=pageNum" json:"page_num"`
}

func (m *GetFavoriteRequest) Reset()                    { *m = GetFavoriteRequest{} }
func (m *GetFavoriteRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFavoriteRequest) ProtoMessage()               {}
func (*GetFavoriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GetFavoriteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetFavoriteRequest) GetGoodsType() string {
	if m != nil {
		return m.GoodsType
	}
	return ""
}

func (m *GetFavoriteRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetFavoriteRequest) GetPageNum() uint32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type GetFavoriteResponse struct {
	Code uint32       `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *FavoriteArr `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string       `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetFavoriteResponse) Reset()                    { *m = GetFavoriteResponse{} }
func (m *GetFavoriteResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFavoriteResponse) ProtoMessage()               {}
func (*GetFavoriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GetFavoriteResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetFavoriteResponse) GetData() *FavoriteArr {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetFavoriteResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type FavoriteArr struct {
	PageNum  uint64          `protobuf:"varint,1,opt,name=page_num,json=pageNum" json:"page_num"`
	RowCount uint64          `protobuf:"varint,2,opt,name=row_count,json=rowCount" json:"row_count"`
	Row      []*FavoriteData `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *FavoriteArr) Reset()                    { *m = FavoriteArr{} }
func (m *FavoriteArr) String() string            { return proto.CompactTextString(m) }
func (*FavoriteArr) ProtoMessage()               {}
func (*FavoriteArr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *FavoriteArr) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *FavoriteArr) GetRowCount() uint64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *FavoriteArr) GetRow() []*FavoriteData {
	if m != nil {
		return m.Row
	}
	return nil
}

type FavoriteData struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username"`
	GoodsId   string `protobuf:"bytes,2,opt,name=goods_id,json=goodsId" json:"goods_id"`
	GoodsName string `protobuf:"bytes,3,opt,name=goods_name,json=goodsName" json:"goods_name"`
	Price     uint64 `protobuf:"varint,4,opt,name=price" json:"price"`
	Time      uint64 `protobuf:"varint,5,opt,name=time" json:"time"`
}

func (m *FavoriteData) Reset()                    { *m = FavoriteData{} }
func (m *FavoriteData) String() string            { return proto.CompactTextString(m) }
func (*FavoriteData) ProtoMessage()               {}
func (*FavoriteData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *FavoriteData) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *FavoriteData) GetGoodsId() string {
	if m != nil {
		return m.GoodsId
	}
	return ""
}

func (m *FavoriteData) GetGoodsName() string {
	if m != nil {
		return m.GoodsName
	}
	return ""
}

func (m *FavoriteData) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *FavoriteData) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type GetBalanceRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username"`
	Random   string `protobuf:"bytes,2,opt,name=random" json:"random"`
}

func (m *GetBalanceRequest) Reset()                    { *m = GetBalanceRequest{} }
func (m *GetBalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceRequest) ProtoMessage()               {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *GetBalanceRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetBalanceRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

type GetBalanceResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetBalanceResponse) Reset()                    { *m = GetBalanceResponse{} }
func (m *GetBalanceResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceResponse) ProtoMessage()               {}
func (*GetBalanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *GetBalanceResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetBalanceResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *GetBalanceResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryMyBuyRequest struct {
	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size"`
	PageNum   int32  `protobuf:"varint,2,opt,name=page_num,json=pageNum" json:"page_num"`
	Username  string `protobuf:"bytes,3,opt,name=username" json:"username"`
	Random    string `protobuf:"bytes,4,opt,name=random" json:"random"`
	Signature string `protobuf:"bytes,5,opt,name=signature" json:"signature"`
}

func (m *QueryMyBuyRequest) Reset()                    { *m = QueryMyBuyRequest{} }
func (m *QueryMyBuyRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryMyBuyRequest) ProtoMessage()               {}
func (*QueryMyBuyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *QueryMyBuyRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *QueryMyBuyRequest) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryMyBuyRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryMyBuyRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func (m *QueryMyBuyRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type QueryMyBuyResponse struct {
	Code uint32   `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *BuyData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string   `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *QueryMyBuyResponse) Reset()                    { *m = QueryMyBuyResponse{} }
func (m *QueryMyBuyResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryMyBuyResponse) ProtoMessage()               {}
func (*QueryMyBuyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *QueryMyBuyResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryMyBuyResponse) GetData() *BuyData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueryMyBuyResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type BuyData struct {
	PageNum  int32  `protobuf:"varint,2,opt,name=page_num,json=pageNum" json:"page_num"`
	RowCount int32  `protobuf:"varint,1,opt,name=row_count,json=rowCount" json:"row_count"`
	Row      []*Buy `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *BuyData) Reset()                    { *m = BuyData{} }
func (m *BuyData) String() string            { return proto.CompactTextString(m) }
func (*BuyData) ProtoMessage()               {}
func (*BuyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *BuyData) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *BuyData) GetRowCount() int32 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *BuyData) GetRow() []*Buy {
	if m != nil {
		return m.Row
	}
	return nil
}

type Buy struct {
	ExchangeId  string `protobuf:"bytes,1,opt,name=exchange_id,json=exchangeId" json:"exchange_id"`
	Username    string `protobuf:"bytes,2,opt,name=username" json:"username"`
	AssetId     string `protobuf:"bytes,3,opt,name=asset_id,json=assetId" json:"asset_id"`
	AssetName   string `protobuf:"bytes,4,opt,name=asset_name,json=assetName" json:"asset_name"`
	AssetType   string `protobuf:"bytes,5,opt,name=asset_type,json=assetType" json:"asset_type"`
	FeatureTag  string `protobuf:"bytes,6,opt,name=feature_tag,json=featureTag" json:"feature_tag"`
	Price       uint64 `protobuf:"varint,7,opt,name=price" json:"price"`
	SampleHash  string `protobuf:"bytes,8,opt,name=sample_hash,json=sampleHash" json:"sample_hash"`
	StorageHash string `protobuf:"bytes,9,opt,name=storage_hash,json=storageHash" json:"storage_hash"`
	Expiretime  uint64 `protobuf:"varint,10,opt,name=expiretime" json:"expiretime"`
}

func (m *Buy) Reset()                    { *m = Buy{} }
func (m *Buy) String() string            { return proto.CompactTextString(m) }
func (*Buy) ProtoMessage()               {}
func (*Buy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

func (m *Buy) GetExchangeId() string {
	if m != nil {
		return m.ExchangeId
	}
	return ""
}

func (m *Buy) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Buy) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *Buy) GetAssetName() string {
	if m != nil {
		return m.AssetName
	}
	return ""
}

func (m *Buy) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *Buy) GetFeatureTag() string {
	if m != nil {
		return m.FeatureTag
	}
	return ""
}

func (m *Buy) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Buy) GetSampleHash() string {
	if m != nil {
		return m.SampleHash
	}
	return ""
}

func (m *Buy) GetStorageHash() string {
	if m != nil {
		return m.StorageHash
	}
	return ""
}

func (m *Buy) GetExpiretime() uint64 {
	if m != nil {
		return m.Expiretime
	}
	return 0
}

func init() {
	proto.RegisterType((*PushTxRequest)(nil), "PushTxRequest")
	proto.RegisterType((*PushTxResponse)(nil), "PushTxResponse")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*AccountInfo)(nil), "AccountInfo")
	proto.RegisterType((*UserInfo)(nil), "UserInfo")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*GetBlockHeaderRequest)(nil), "GetBlockHeaderRequest")
	proto.RegisterType((*GetBlockHeaderResponse)(nil), "GetBlockHeaderResponse")
	proto.RegisterType((*BlockHeader)(nil), "BlockHeader")
	proto.RegisterType((*GetAccountInfoRequest)(nil), "GetAccountInfoRequest")
	proto.RegisterType((*GetAccountInfoResponse)(nil), "GetAccountInfoResponse")
	proto.RegisterType((*AccountInfoData)(nil), "AccountInfoData")
	proto.RegisterType((*FavoriteRequest)(nil), "FavoriteRequest")
	proto.RegisterType((*FavoriteResponse)(nil), "FavoriteResponse")
	proto.RegisterType((*GetFavoriteRequest)(nil), "GetFavoriteRequest")
	proto.RegisterType((*GetFavoriteResponse)(nil), "GetFavoriteResponse")
	proto.RegisterType((*FavoriteArr)(nil), "FavoriteArr")
	proto.RegisterType((*FavoriteData)(nil), "FavoriteData")
	proto.RegisterType((*GetBalanceRequest)(nil), "GetBalanceRequest")
	proto.RegisterType((*GetBalanceResponse)(nil), "GetBalanceResponse")
	proto.RegisterType((*QueryMyBuyRequest)(nil), "QueryMyBuyRequest")
	proto.RegisterType((*QueryMyBuyResponse)(nil), "QueryMyBuyResponse")
	proto.RegisterType((*BuyData)(nil), "BuyData")
	proto.RegisterType((*Buy)(nil), "Buy")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x8e, 0xf3, 0xb3, 0x89, 0x8f, 0x93, 0xdd, 0xec, 0x6c, 0xbb, 0x75, 0x43, 0xcb, 0x6e, 0xad,
	0x52, 0xed, 0x0d, 0x96, 0x58, 0x04, 0x85, 0xde, 0x75, 0x5b, 0xd1, 0xae, 0x54, 0x56, 0x10, 0x16,
	0x04, 0x17, 0x28, 0xcc, 0xc6, 0x13, 0xc7, 0x6a, 0x62, 0x87, 0x19, 0x7b, 0xdb, 0xf4, 0x05, 0xb8,
	0x02, 0x6e, 0x91, 0x10, 0x08, 0xf1, 0x10, 0x3c, 0x01, 0x0f, 0x86, 0xe6, 0xcc, 0xd8, 0x19, 0x27,
	0x9b, 0x88, 0xaa, 0xb7, 0xbd, 0xf3, 0xf9, 0x99, 0x9f, 0x73, 0xce, 0x77, 0xbe, 0x39, 0x06, 0xc8,
	0x04, 0xe3, 0xfe, 0x8c, 0x27, 0x69, 0xe2, 0xfd, 0x55, 0x85, 0xce, 0x17, 0x99, 0x18, 0x9f, 0xbf,
	0xec, 0xb3, 0x1f, 0x33, 0x26, 0x52, 0xe2, 0x42, 0xf3, 0x92, 0x71, 0x11, 0x25, 0xb1, 0x6b, 0x1d,
	0x5a, 0x47, 0x9d, 0x7e, 0x2e, 0x92, 0xdb, 0x00, 0xc3, 0x8c, 0x8b, 0x84, 0x0f, 0xe2, 0x6c, 0xea,
	0x56, 0xd1, 0x68, 0x2b, 0xcd, 0x59, 0x36, 0x25, 0x77, 0xa0, 0xad, 0xcd, 0x13, 0x7a, 0xc1, 0x26,
	0x6e, 0x0d, 0x1d, 0x1c, 0xa5, 0x7b, 0x26, 0x55, 0xa4, 0x07, 0xad, 0x49, 0x34, 0x62, 0x69, 0x34,
	0x65, 0x6e, 0xfd, 0xd0, 0x3a, 0xaa, 0xf7, 0x0b, 0x99, 0xec, 0xc3, 0x96, 0x60, 0x71, 0xc0, 0xb8,
	0xdb, 0x38, 0xb4, 0x8e, 0xec, 0xbe, 0x96, 0xe4, 0x9a, 0x61, 0x12, 0xa7, 0x9c, 0x0e, 0x53, 0x77,
	0x0b, 0x2d, 0x85, 0x2c, 0xd7, 0x4c, 0x59, 0x3a, 0x4e, 0x02, 0xb7, 0xa9, 0xd6, 0x28, 0x89, 0x5c,
	0x83, 0xc6, 0x8c, 0x72, 0x3a, 0x75, 0x5b, 0xa8, 0x56, 0x02, 0xb9, 0x01, 0x4d, 0x11, 0x85, 0x03,
	0x3a, 0x09, 0x5d, 0x1b, 0xef, 0xb6, 0x25, 0xa2, 0xf0, 0xe1, 0x24, 0x24, 0xb7, 0xc0, 0x16, 0x51,
	0x18, 0xd3, 0x34, 0xe3, 0xcc, 0x05, 0x5c, 0xb2, 0x50, 0x78, 0x1f, 0xc3, 0x76, 0x9e, 0x21, 0x31,
	0x4b, 0x62, 0xc1, 0x08, 0x81, 0xfa, 0x30, 0x09, 0x98, 0xce, 0x0f, 0x7e, 0x93, 0x2e, 0xd4, 0xa6,
	0x22, 0xc4, 0xac, 0xd8, 0x7d, 0xf9, 0xe9, 0xfd, 0x66, 0xc1, 0x4e, 0x9f, 0x85, 0x91, 0x48, 0x19,
	0xcf, 0x93, 0x7b, 0x0f, 0x9a, 0x74, 0x38, 0x4c, 0xb2, 0x38, 0xc5, 0xc5, 0xce, 0x71, 0xdb, 0x7f,
	0xa8, 0xe4, 0xd3, 0x78, 0x94, 0xf4, 0x73, 0x23, 0xb9, 0x0d, 0x75, 0x59, 0x24, 0xdc, 0xce, 0x39,
	0xb6, 0xfd, 0xaf, 0x05, 0xe3, 0xe8, 0x81, 0x6a, 0xf2, 0x0e, 0xd8, 0x97, 0x8c, 0x47, 0xa3, 0xf9,
	0x20, 0x0a, 0x30, 0xcf, 0x76, 0xbf, 0xa5, 0x14, 0xa7, 0x81, 0xac, 0x83, 0x36, 0x5e, 0xd2, 0x49,
	0xa6, 0x12, 0x6d, 0xf7, 0x1d, 0xa5, 0xfb, 0x46, 0xaa, 0xbc, 0x4f, 0xc1, 0x31, 0x8e, 0x95, 0xf1,
	0xc4, 0x74, 0xca, 0xf4, 0xe5, 0xf1, 0x5b, 0xa6, 0x76, 0x96, 0x5d, 0x3c, 0x67, 0x73, 0xbd, 0xbf,
	0x96, 0xbc, 0xdf, 0xab, 0xd0, 0xca, 0x6f, 0xf3, 0x16, 0x2b, 0xcb, 0x58, 0xf9, 0x04, 0xba, 0x8b,
	0x92, 0xbf, 0x16, 0x5a, 0xfe, 0xb4, 0xa0, 0xfd, 0x2c, 0x09, 0xa3, 0x38, 0x87, 0x4a, 0x0f, 0x5a,
	0xb2, 0xd6, 0x58, 0x18, 0x4b, 0xc5, 0x92, 0xcb, 0x32, 0x16, 0x4e, 0xe3, 0x20, 0x99, 0xea, 0x1d,
	0xb4, 0xf4, 0xa6, 0xb8, 0x90, 0x67, 0xca, 0x58, 0x30, 0x36, 0x95, 0xd9, 0x42, 0xf6, 0x3e, 0x82,
	0x8e, 0xbe, 0xdf, 0x6b, 0xc5, 0x75, 0x03, 0xae, 0x3f, 0x61, 0xe9, 0xc9, 0x24, 0x19, 0x3e, 0x7f,
	0xca, 0x68, 0x50, 0xb4, 0x82, 0xf7, 0x03, 0xec, 0x2f, 0x1b, 0x36, 0x6c, 0x7c, 0x08, 0xf5, 0x80,
	0xa6, 0x54, 0x37, 0x44, 0xdb, 0x37, 0xd7, 0xa1, 0x25, 0x3f, 0xba, 0xb6, 0x38, 0xfa, 0xd7, 0x2a,
	0x38, 0x86, 0x1f, 0xb9, 0x0b, 0xdb, 0x63, 0x46, 0x83, 0xc1, 0x85, 0xd4, 0x21, 0x2e, 0xd5, 0x09,
	0x6d, 0xa9, 0x45, 0x47, 0x09, 0xcd, 0x7b, 0xb0, 0x63, 0x78, 0x8d, 0xa9, 0x18, 0xeb, 0x70, 0x3a,
	0x85, 0xdb, 0x53, 0x2a, 0xc6, 0x4b, 0x7e, 0x08, 0xd3, 0x1a, 0xc2, 0x74, 0xe1, 0x77, 0x2e, 0xb1,
	0xea, 0xc3, 0x9e, 0xe1, 0x17, 0xb0, 0x09, 0x0b, 0x69, 0x9a, 0x67, 0x7f, 0xb7, 0xf0, 0x7d, 0xac,
	0x0d, 0x2b, 0xad, 0xd1, 0x58, 0x6d, 0x8d, 0xfb, 0xe0, 0x4e, 0xa8, 0x48, 0x07, 0x43, 0x99, 0xae,
	0x58, 0x64, 0xc2, 0x08, 0x69, 0x0b, 0xdd, 0xaf, 0x4b, 0xfb, 0xa3, 0xdc, 0x9c, 0xc7, 0xe6, 0x3d,
	0xc0, 0x62, 0x98, 0x8c, 0xa3, 0xc1, 0x76, 0x07, 0xda, 0x9a, 0x7a, 0x06, 0x06, 0xe0, 0x1c, 0xad,
	0x3b, 0xa3, 0x53, 0xe6, 0x05, 0x58, 0xaf, 0xd2, 0xda, 0x0d, 0xf5, 0xba, 0x5b, 0xaa, 0x57, 0xd7,
	0x64, 0xb9, 0xc7, 0x34, 0xa5, 0x6b, 0x6b, 0xf6, 0xb3, 0x05, 0x3b, 0x4b, 0xbe, 0xff, 0xe3, 0x72,
	0x06, 0x5b, 0x55, 0x4d, 0xb6, 0x92, 0x04, 0x75, 0x41, 0x27, 0x34, 0x1e, 0xe6, 0xc5, 0xc9, 0x45,
	0xf2, 0x1e, 0x6c, 0x8b, 0x94, 0x3e, 0x67, 0xc1, 0x20, 0x77, 0x50, 0x24, 0xd3, 0x51, 0xda, 0x13,
	0xa5, 0xf4, 0xfe, 0xae, 0xc2, 0xce, 0x67, 0xf4, 0x32, 0xe1, 0x51, 0xca, 0xde, 0xbe, 0x90, 0x6b,
	0x58, 0xef, 0x19, 0x74, 0x17, 0x39, 0xda, 0x00, 0x0a, 0x62, 0x80, 0xc2, 0x5e, 0x0b, 0x81, 0x9f,
	0x2c, 0x20, 0x4f, 0x58, 0xba, 0x9c, 0xf5, 0x4d, 0x7c, 0x78, 0x1b, 0x20, 0x4c, 0x92, 0x40, 0x0c,
	0xd2, 0xf9, 0x2c, 0x7f, 0xc6, 0x6c, 0xd4, 0x9c, 0xcf, 0x67, 0x4c, 0xd2, 0xe2, 0x8c, 0x86, 0x6c,
	0x20, 0xa2, 0x57, 0x4c, 0x27, 0xbd, 0x25, 0x15, 0x5f, 0x45, 0xaf, 0x18, 0xb9, 0x09, 0xf8, 0x8d,
	0x15, 0xab, 0xab, 0x72, 0x4a, 0x59, 0xb6, 0xcb, 0xf7, 0xb0, 0x57, 0xba, 0xc8, 0x6b, 0xf0, 0x53,
	0xbe, 0xe8, 0x21, 0x5f, 0xcf, 0x4f, 0x23, 0x70, 0x0c, 0xb7, 0xd2, 0x45, 0x2c, 0x05, 0x56, 0x7d,
	0x11, 0x19, 0x00, 0x4f, 0x5e, 0x0c, 0xd4, 0xe0, 0x50, 0x55, 0xb0, 0xe0, 0xc9, 0x8b, 0x47, 0x38,
	0x2b, 0x1c, 0x40, 0x8d, 0x27, 0x2f, 0xdc, 0xda, 0x61, 0xed, 0xc8, 0x39, 0xee, 0x14, 0x27, 0x63,
	0x9b, 0x49, 0x8b, 0xf7, 0x8b, 0x05, 0x6d, 0x53, 0xbb, 0x31, 0x95, 0x37, 0xa1, 0xa5, 0x52, 0x19,
	0x05, 0x3a, 0x91, 0x4d, 0x94, 0x4f, 0x83, 0x45, 0x96, 0x71, 0x61, 0xcd, 0xc8, 0x32, 0xf6, 0xa0,
	0x84, 0x14, 0x8f, 0x8a, 0x46, 0x52, 0x82, 0x4c, 0x16, 0x82, 0xb9, 0x81, 0x4a, 0xfc, 0xf6, 0x9e,
	0xc0, 0xae, 0xa4, 0x7e, 0xd5, 0x62, 0x6f, 0xf0, 0xde, 0x79, 0x67, 0x88, 0x94, 0x62, 0xa3, 0x37,
	0x86, 0xde, 0x1f, 0x16, 0xec, 0x7e, 0x99, 0x31, 0x3e, 0xff, 0x7c, 0x7e, 0x92, 0xcd, 0xf3, 0x9b,
	0x95, 0xe0, 0x23, 0x37, 0x6d, 0xac, 0x81, 0x4f, 0x15, 0x6d, 0x45, 0xd5, 0xcc, 0x88, 0x6a, 0x6b,
	0x23, 0xaa, 0x97, 0x5e, 0xf0, 0x52, 0xa3, 0x35, 0x96, 0x1b, 0xed, 0x5b, 0x20, 0xe6, 0xf5, 0x36,
	0xc4, 0x7b, 0xab, 0x84, 0xc7, 0x96, 0x7f, 0x92, 0xcd, 0x37, 0xf2, 0xee, 0x77, 0xd0, 0xd4, 0x2e,
	0x9b, 0x22, 0x2a, 0xe1, 0x50, 0x67, 0xa2, 0xc0, 0xe1, 0xbe, 0x89, 0xc3, 0xba, 0x3c, 0x51, 0xc1,
	0xef, 0x9f, 0x2a, 0xd4, 0x4e, 0xb2, 0x39, 0x39, 0x00, 0x87, 0xbd, 0x1c, 0x8e, 0x69, 0x1c, 0x32,
	0x09, 0x2e, 0x55, 0x63, 0xc8, 0x55, 0xa7, 0x41, 0x29, 0x5f, 0xd5, 0x55, 0x58, 0x52, 0x21, 0x58,
	0xba, 0x18, 0x6c, 0x9a, 0x28, 0x2b, 0x58, 0x2a, 0x13, 0x2e, 0x54, 0xe9, 0xb4, 0x51, 0x73, 0xa6,
	0xb9, 0x41, 0x99, 0x91, 0x1b, 0x1a, 0x86, 0x19, 0xb9, 0xe1, 0x00, 0x9c, 0x11, 0xc3, 0xec, 0x0e,
	0x52, 0x1a, 0x6a, 0xfe, 0x04, 0xad, 0x3a, 0xa7, 0xe1, 0x02, 0xd6, 0x4d, 0x13, 0xd6, 0x07, 0xe0,
	0x08, 0x3a, 0x9d, 0x4d, 0x98, 0x9a, 0x10, 0x14, 0x8b, 0x82, 0x52, 0xe1, 0x78, 0x70, 0x07, 0xda,
	0x22, 0x4d, 0xb8, 0x4c, 0x24, 0x7a, 0xd8, 0xea, 0xd1, 0xd2, 0x3a, 0x74, 0x79, 0x17, 0x80, 0xbd,
	0x9c, 0x45, 0x5c, 0xb1, 0x3d, 0xe0, 0xf6, 0x86, 0xe6, 0xf8, 0xdf, 0x1a, 0xd4, 0xe5, 0xa8, 0x4d,
	0x3e, 0x80, 0x56, 0x3e, 0x55, 0x92, 0xae, 0xbf, 0xf4, 0x4f, 0xd1, 0xdb, 0xf5, 0x97, 0x47, 0x4e,
	0xaf, 0x42, 0x8e, 0xa0, 0x81, 0xd3, 0x1a, 0xe9, 0xf8, 0xe6, 0x54, 0xd9, 0xdb, 0xf6, 0x4b, 0x43,
	0x9c, 0x57, 0x21, 0x8f, 0x60, 0xbb, 0x3c, 0x87, 0x91, 0x7d, 0xff, 0xca, 0x89, 0xad, 0x77, 0xc3,
	0xbf, 0x7a, 0x60, 0x2b, 0x36, 0x31, 0xff, 0x29, 0x70, 0x93, 0xd5, 0x49, 0x43, 0x6d, 0x72, 0xc5,
	0x14, 0xe1, 0x55, 0x64, 0x98, 0x39, 0x4d, 0x91, 0xae, 0xbf, 0xc4, 0xff, 0xbd, 0x5d, 0x7f, 0x99,
	0x88, 0xbd, 0x0a, 0x79, 0x00, 0x8e, 0xc1, 0xd0, 0x64, 0xcf, 0x5f, 0x7d, 0x38, 0x7a, 0xd7, 0xfc,
	0x2b, 0x48, 0xdc, 0xab, 0x90, 0xf7, 0xa1, 0x75, 0xce, 0x69, 0x2c, 0x46, 0x8c, 0x93, 0x6d, 0xbf,
	0xf4, 0x13, 0xdc, 0xdb, 0xf1, 0xcb, 0xbf, 0x7c, 0x5e, 0x85, 0xdc, 0x07, 0x58, 0xf4, 0x1e, 0x21,
	0xfe, 0x0a, 0x4f, 0xf4, 0xf6, 0xfc, 0xd5, 0xe6, 0xf4, 0x2a, 0x17, 0x5b, 0xf8, 0xa7, 0xfd, 0xe1,
	0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x5a, 0x43, 0xa9, 0x77, 0x0f, 0x00, 0x00,
}
