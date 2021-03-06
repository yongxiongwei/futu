// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetSuspend.proto

package Qot_GetSuspend

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
	SecurityList         []*Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`
	BeginTime            *string     `protobuf:"bytes,2,req,name=beginTime" json:"beginTime,omitempty"`
	EndTime              *string     `protobuf:"bytes,3,req,name=endTime" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_eca503b045c68f08, []int{0}
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

func (m *C2S) GetSecurityList() []*Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

func (m *C2S) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *C2S) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

type Suspend struct {
	Time                 *string  `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Suspend) Reset()         { *m = Suspend{} }
func (m *Suspend) String() string { return proto.CompactTextString(m) }
func (*Suspend) ProtoMessage()    {}
func (*Suspend) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_eca503b045c68f08, []int{1}
}
func (m *Suspend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suspend.Unmarshal(m, b)
}
func (m *Suspend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suspend.Marshal(b, m, deterministic)
}
func (dst *Suspend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suspend.Merge(dst, src)
}
func (m *Suspend) XXX_Size() int {
	return xxx_messageInfo_Suspend.Size(m)
}
func (m *Suspend) XXX_DiscardUnknown() {
	xxx_messageInfo_Suspend.DiscardUnknown(m)
}

var xxx_messageInfo_Suspend proto.InternalMessageInfo

func (m *Suspend) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

type SecuritySuspend struct {
	Security             *Security  `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	SuspendList          []*Suspend `protobuf:"bytes,2,rep,name=suspendList" json:"suspendList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SecuritySuspend) Reset()         { *m = SecuritySuspend{} }
func (m *SecuritySuspend) String() string { return proto.CompactTextString(m) }
func (*SecuritySuspend) ProtoMessage()    {}
func (*SecuritySuspend) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_eca503b045c68f08, []int{2}
}
func (m *SecuritySuspend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecuritySuspend.Unmarshal(m, b)
}
func (m *SecuritySuspend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecuritySuspend.Marshal(b, m, deterministic)
}
func (dst *SecuritySuspend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecuritySuspend.Merge(dst, src)
}
func (m *SecuritySuspend) XXX_Size() int {
	return xxx_messageInfo_SecuritySuspend.Size(m)
}
func (m *SecuritySuspend) XXX_DiscardUnknown() {
	xxx_messageInfo_SecuritySuspend.DiscardUnknown(m)
}

var xxx_messageInfo_SecuritySuspend proto.InternalMessageInfo

func (m *SecuritySuspend) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *SecuritySuspend) GetSuspendList() []*Suspend {
	if m != nil {
		return m.SuspendList
	}
	return nil
}

type S2C struct {
	SecuritySuspendList  []*SecuritySuspend `protobuf:"bytes,1,rep,name=SecuritySuspendList" json:"SecuritySuspendList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_eca503b045c68f08, []int{3}
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

func (m *S2C) GetSecuritySuspendList() []*SecuritySuspend {
	if m != nil {
		return m.SecuritySuspendList
	}
	return nil
}

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
	return fileDescriptor_Qot_GetSuspend_eca503b045c68f08, []int{4}
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
	return fileDescriptor_Qot_GetSuspend_eca503b045c68f08, []int{5}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetSuspend.C2S")
	proto.RegisterType((*Suspend)(nil), "Qot_GetSuspend.Suspend")
	proto.RegisterType((*SecuritySuspend)(nil), "Qot_GetSuspend.SecuritySuspend")
	proto.RegisterType((*S2C)(nil), "Qot_GetSuspend.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetSuspend.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetSuspend.Response")
}

func init() {
	proto.RegisterFile("Qot_GetSuspend.proto", fileDescriptor_Qot_GetSuspend_eca503b045c68f08)
}

var fileDescriptor_Qot_GetSuspend_eca503b045c68f08 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x95, 0xa4, 0xfd, 0xa6, 0xbd, 0x56, 0x5f, 0x90, 0x5b, 0x41, 0x54, 0xf1, 0x23, 0x8a,
	0x84, 0x94, 0x85, 0x2a, 0xb2, 0x18, 0x80, 0x35, 0x03, 0x0b, 0x0c, 0xb5, 0x3b, 0xb0, 0x21, 0xd1,
	0x9e, 0xaa, 0x0c, 0x8d, 0x83, 0xed, 0x0a, 0x75, 0x61, 0xe1, 0x1f, 0x47, 0x76, 0x63, 0x48, 0x43,
	0x99, 0x92, 0x7b, 0xef, 0x63, 0xfb, 0xdd, 0x1d, 0x8c, 0x67, 0x42, 0xbf, 0x3c, 0xa0, 0xe6, 0x1b,
	0x55, 0x61, 0xb9, 0x9c, 0x56, 0x52, 0x68, 0x41, 0xfe, 0xef, 0xab, 0x93, 0x61, 0x2e, 0xd6, 0x6b,
	0x51, 0xee, 0xdc, 0xc9, 0xb1, 0x71, 0x9b, 0x4a, 0xf2, 0x0e, 0x41, 0x4e, 0x39, 0xb9, 0x85, 0xa1,
	0xc2, 0xc5, 0x46, 0x16, 0x7a, 0xfb, 0x58, 0x28, 0x1d, 0x79, 0x71, 0x90, 0x0e, 0xe8, 0x78, 0xda,
	0xe0, 0x79, 0xed, 0xb3, 0x3d, 0x92, 0x9c, 0x41, 0xff, 0x15, 0x57, 0x45, 0x39, 0x2f, 0xd6, 0x18,
	0xf9, 0xb1, 0x9f, 0xf6, 0xd9, 0x8f, 0x40, 0x22, 0x08, 0xb1, 0x5c, 0x5a, 0x2f, 0xb0, 0x9e, 0x2b,
	0x93, 0x73, 0x08, 0xeb, 0x8c, 0x84, 0x40, 0x47, 0x1b, 0xc2, 0xb3, 0x84, 0xfd, 0x4f, 0x3e, 0xe0,
	0xc8, 0x3d, 0xe8, 0xb0, 0x0c, 0x7a, 0xee, 0x65, 0x8b, 0xfe, 0x95, 0xef, 0x9b, 0x22, 0x77, 0x30,
	0x50, 0xbb, 0xc3, 0xb6, 0x29, 0xdf, 0x36, 0x75, 0x3a, 0x6d, 0x0d, 0xae, 0xfe, 0xb2, 0x26, 0x9b,
	0x3c, 0x43, 0xc0, 0x69, 0x4e, 0x66, 0x30, 0x6a, 0xc5, 0x68, 0x8c, 0xe7, 0xf2, 0xd7, 0x4d, 0xfb,
	0x28, 0x3b, 0x74, 0x36, 0xc9, 0x20, 0x64, 0xf8, 0xb6, 0x41, 0xa5, 0xc9, 0x15, 0x04, 0x0b, 0xaa,
	0xea, 0x66, 0x46, 0xed, 0xdb, 0x72, 0xca, 0x99, 0xf1, 0x93, 0x4f, 0x0f, 0x7a, 0x0c, 0x55, 0x25,
	0x4a, 0x85, 0xe4, 0x02, 0x42, 0x89, 0x7a, 0xbe, 0xad, 0x76, 0xf3, 0xea, 0xde, 0x77, 0xae, 0x6f,
	0xb2, 0x8c, 0x39, 0x91, 0x9c, 0xc0, 0x3f, 0x89, 0xfa, 0x49, 0xad, 0x22, 0x3f, 0xf6, 0xd2, 0x3e,
	0xab, 0x2b, 0xbb, 0x09, 0x29, 0x73, 0xb1, 0x34, 0x9b, 0xf0, 0xd2, 0x2e, 0x73, 0xa5, 0x49, 0xa1,
	0xe8, 0x22, 0xea, 0xc4, 0xde, 0xa1, 0x14, 0x9c, 0xe6, 0xcc, 0xf8, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x9d, 0xc1, 0xae, 0x1b, 0x70, 0x02, 0x00, 0x00,
}
