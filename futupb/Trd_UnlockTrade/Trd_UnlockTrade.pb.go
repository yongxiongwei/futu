// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_UnlockTrade.proto

package Trd_UnlockTrade

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

type C2S struct {
	Unlock               *bool    `protobuf:"varint,1,req,name=unlock" json:"unlock,omitempty"`
	PwdMD5               *string  `protobuf:"bytes,2,opt,name=pwdMD5" json:"pwdMD5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_UnlockTrade_59134a598c38892c, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetUnlock() bool {
	if m != nil && m.Unlock != nil {
		return *m.Unlock
	}
	return false
}

func (m *C2S) GetPwdMD5() string {
	if m != nil && m.PwdMD5 != nil {
		return *m.PwdMD5
	}
	return ""
}

type S2C struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_UnlockTrade_59134a598c38892c, []int{1}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_UnlockTrade_59134a598c38892c, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_UnlockTrade_59134a598c38892c, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Trd_UnlockTrade.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_UnlockTrade.S2C")
	proto.RegisterType((*Request)(nil), "Trd_UnlockTrade.Request")
	proto.RegisterType((*Response)(nil), "Trd_UnlockTrade.Response")
}

func init() {
	proto.RegisterFile("Trd_UnlockTrade.proto", fileDescriptor_Trd_UnlockTrade_59134a598c38892c)
}

var fileDescriptor_Trd_UnlockTrade_59134a598c38892c = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0xb1, 0x4b, 0x03, 0x31,
	0x14, 0xc6, 0x49, 0xaf, 0xe7, 0xd5, 0xa7, 0x20, 0x04, 0x95, 0xe0, 0x20, 0x21, 0x83, 0xdc, 0x62,
	0xa9, 0xc1, 0x2e, 0xae, 0x71, 0xed, 0xf2, 0x72, 0xce, 0x22, 0x97, 0x87, 0x83, 0xf6, 0x12, 0x93,
	0x14, 0x71, 0xf7, 0x0f, 0x97, 0x5c, 0x7b, 0x4b, 0x71, 0xfc, 0x7d, 0xef, 0x7b, 0x7c, 0x3f, 0xb8,
	0xea, 0xa2, 0x7b, 0x7d, 0x19, 0x3e, 0x7d, 0xff, 0xd1, 0xc5, 0x37, 0x47, 0xcb, 0x10, 0x7d, 0xf6,
	0xfc, 0xe2, 0x28, 0xbe, 0x39, 0x37, 0x7e, 0xbb, 0xf5, 0xc3, 0xfe, 0xac, 0xd6, 0x50, 0x19, 0x6d,
	0xf9, 0x35, 0x9c, 0xec, 0xc6, 0x8e, 0x60, 0x72, 0xd6, 0x2e, 0xf0, 0x40, 0x25, 0x0f, 0xdf, 0x6e,
	0xf3, 0xbc, 0x16, 0x33, 0xc9, 0xda, 0x53, 0x3c, 0x90, 0xaa, 0xa1, 0xb2, 0xda, 0xa8, 0x07, 0x68,
	0x90, 0xbe, 0x76, 0x94, 0x32, 0xbf, 0x83, 0xaa, 0xd7, 0x69, 0x7c, 0x3f, 0xd3, 0x97, 0xcb, 0x63,
	0x19, 0xa3, 0x2d, 0x96, 0x82, 0xfa, 0x65, 0xb0, 0x40, 0x4a, 0xc1, 0x0f, 0x89, 0xf8, 0x2d, 0x34,
	0x91, 0x72, 0xf7, 0x13, 0x68, 0x7c, 0xac, 0x9f, 0xe6, 0xf7, 0x8f, 0xab, 0x15, 0x4e, 0x61, 0x99,
	0x8f, 0x94, 0x37, 0xe9, 0x7d, 0x9a, 0xdf, 0x13, 0x17, 0xd0, 0x50, 0x8c, 0xc6, 0x3b, 0x12, 0x95,
	0x64, 0x6d, 0x8d, 0x13, 0x16, 0x8d, 0xa4, 0x7b, 0x31, 0x97, 0xec, 0x5f, 0x0d, 0xab, 0x0d, 0x96,
	0xc2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x89, 0xb8, 0xee, 0x2e, 0x01, 0x00, 0x00,
}