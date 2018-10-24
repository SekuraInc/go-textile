// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thread.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ThreadBlock_Type int32

const (
	ThreadBlock_MERGE      ThreadBlock_Type = 0
	ThreadBlock_IGNORE     ThreadBlock_Type = 1
	ThreadBlock_JOIN       ThreadBlock_Type = 2
	ThreadBlock_ANNOUNCE   ThreadBlock_Type = 3
	ThreadBlock_LEAVE      ThreadBlock_Type = 4
	ThreadBlock_DATA       ThreadBlock_Type = 5
	ThreadBlock_ANNOTATION ThreadBlock_Type = 6
	ThreadBlock_INVITE     ThreadBlock_Type = 50
)

var ThreadBlock_Type_name = map[int32]string{
	0:  "MERGE",
	1:  "IGNORE",
	2:  "JOIN",
	3:  "ANNOUNCE",
	4:  "LEAVE",
	5:  "DATA",
	6:  "ANNOTATION",
	50: "INVITE",
}
var ThreadBlock_Type_value = map[string]int32{
	"MERGE":      0,
	"IGNORE":     1,
	"JOIN":       2,
	"ANNOUNCE":   3,
	"LEAVE":      4,
	"DATA":       5,
	"ANNOTATION": 6,
	"INVITE":     50,
}

func (x ThreadBlock_Type) String() string {
	return proto.EnumName(ThreadBlock_Type_name, int32(x))
}
func (ThreadBlock_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{1, 0}
}

type ThreadData_Type int32

const (
	ThreadData_PHOTO ThreadData_Type = 0
	ThreadData_TEXT  ThreadData_Type = 1
)

var ThreadData_Type_name = map[int32]string{
	0: "PHOTO",
	1: "TEXT",
}
var ThreadData_Type_value = map[string]int32{
	"PHOTO": 0,
	"TEXT":  1,
}

func (x ThreadData_Type) String() string {
	return proto.EnumName(ThreadData_Type_name, int32(x))
}
func (ThreadData_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{7, 0}
}

type ThreadAnnotation_Type int32

const (
	ThreadAnnotation_COMMENT ThreadAnnotation_Type = 0
	ThreadAnnotation_LIKE    ThreadAnnotation_Type = 1
)

var ThreadAnnotation_Type_name = map[int32]string{
	0: "COMMENT",
	1: "LIKE",
}
var ThreadAnnotation_Type_value = map[string]int32{
	"COMMENT": 0,
	"LIKE":    1,
}

func (x ThreadAnnotation_Type) String() string {
	return proto.EnumName(ThreadAnnotation_Type_name, int32(x))
}
func (ThreadAnnotation_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{8, 0}
}

// for wire transport
type ThreadEnvelope struct {
	Thread               string   `protobuf:"bytes,1,opt,name=thread,proto3" json:"thread,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	CipherBlock          []byte   `protobuf:"bytes,3,opt,name=cipherBlock,proto3" json:"cipherBlock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadEnvelope) Reset()         { *m = ThreadEnvelope{} }
func (m *ThreadEnvelope) String() string { return proto.CompactTextString(m) }
func (*ThreadEnvelope) ProtoMessage()    {}
func (*ThreadEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{0}
}
func (m *ThreadEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadEnvelope.Unmarshal(m, b)
}
func (m *ThreadEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadEnvelope.Marshal(b, m, deterministic)
}
func (dst *ThreadEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadEnvelope.Merge(dst, src)
}
func (m *ThreadEnvelope) XXX_Size() int {
	return xxx_messageInfo_ThreadEnvelope.Size(m)
}
func (m *ThreadEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadEnvelope proto.InternalMessageInfo

func (m *ThreadEnvelope) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

func (m *ThreadEnvelope) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ThreadEnvelope) GetCipherBlock() []byte {
	if m != nil {
		return m.CipherBlock
	}
	return nil
}

type ThreadBlock struct {
	Header               *ThreadBlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Type                 ThreadBlock_Type   `protobuf:"varint,2,opt,name=type,proto3,enum=ThreadBlock_Type" json:"type,omitempty"`
	Payload              *any.Any           `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ThreadBlock) Reset()         { *m = ThreadBlock{} }
func (m *ThreadBlock) String() string { return proto.CompactTextString(m) }
func (*ThreadBlock) ProtoMessage()    {}
func (*ThreadBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{1}
}
func (m *ThreadBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadBlock.Unmarshal(m, b)
}
func (m *ThreadBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadBlock.Marshal(b, m, deterministic)
}
func (dst *ThreadBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadBlock.Merge(dst, src)
}
func (m *ThreadBlock) XXX_Size() int {
	return xxx_messageInfo_ThreadBlock.Size(m)
}
func (m *ThreadBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadBlock proto.InternalMessageInfo

func (m *ThreadBlock) GetHeader() *ThreadBlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ThreadBlock) GetType() ThreadBlock_Type {
	if m != nil {
		return m.Type
	}
	return ThreadBlock_MERGE
}

func (m *ThreadBlock) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ThreadBlockHeader struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Parents              []string             `protobuf:"bytes,2,rep,name=parents,proto3" json:"parents,omitempty"`
	Author               string               `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ThreadBlockHeader) Reset()         { *m = ThreadBlockHeader{} }
func (m *ThreadBlockHeader) String() string { return proto.CompactTextString(m) }
func (*ThreadBlockHeader) ProtoMessage()    {}
func (*ThreadBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{2}
}
func (m *ThreadBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadBlockHeader.Unmarshal(m, b)
}
func (m *ThreadBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadBlockHeader.Marshal(b, m, deterministic)
}
func (dst *ThreadBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadBlockHeader.Merge(dst, src)
}
func (m *ThreadBlockHeader) XXX_Size() int {
	return xxx_messageInfo_ThreadBlockHeader.Size(m)
}
func (m *ThreadBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadBlockHeader proto.InternalMessageInfo

func (m *ThreadBlockHeader) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ThreadBlockHeader) GetParents() []string {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *ThreadBlockHeader) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type ThreadIgnore struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadIgnore) Reset()         { *m = ThreadIgnore{} }
func (m *ThreadIgnore) String() string { return proto.CompactTextString(m) }
func (*ThreadIgnore) ProtoMessage()    {}
func (*ThreadIgnore) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{3}
}
func (m *ThreadIgnore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadIgnore.Unmarshal(m, b)
}
func (m *ThreadIgnore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadIgnore.Marshal(b, m, deterministic)
}
func (dst *ThreadIgnore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadIgnore.Merge(dst, src)
}
func (m *ThreadIgnore) XXX_Size() int {
	return xxx_messageInfo_ThreadIgnore.Size(m)
}
func (m *ThreadIgnore) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadIgnore.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadIgnore proto.InternalMessageInfo

func (m *ThreadIgnore) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ThreadInvite struct {
	Sk                   []byte   `protobuf:"bytes,1,opt,name=sk,proto3" json:"sk,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadInvite) Reset()         { *m = ThreadInvite{} }
func (m *ThreadInvite) String() string { return proto.CompactTextString(m) }
func (*ThreadInvite) ProtoMessage()    {}
func (*ThreadInvite) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{4}
}
func (m *ThreadInvite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadInvite.Unmarshal(m, b)
}
func (m *ThreadInvite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadInvite.Marshal(b, m, deterministic)
}
func (dst *ThreadInvite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadInvite.Merge(dst, src)
}
func (m *ThreadInvite) XXX_Size() int {
	return xxx_messageInfo_ThreadInvite.Size(m)
}
func (m *ThreadInvite) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadInvite.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadInvite proto.InternalMessageInfo

func (m *ThreadInvite) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *ThreadInvite) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ThreadJoin struct {
	Inviter              string   `protobuf:"bytes,1,opt,name=inviter,proto3" json:"inviter,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Inboxes              []string `protobuf:"bytes,3,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadJoin) Reset()         { *m = ThreadJoin{} }
func (m *ThreadJoin) String() string { return proto.CompactTextString(m) }
func (*ThreadJoin) ProtoMessage()    {}
func (*ThreadJoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{5}
}
func (m *ThreadJoin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadJoin.Unmarshal(m, b)
}
func (m *ThreadJoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadJoin.Marshal(b, m, deterministic)
}
func (dst *ThreadJoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadJoin.Merge(dst, src)
}
func (m *ThreadJoin) XXX_Size() int {
	return xxx_messageInfo_ThreadJoin.Size(m)
}
func (m *ThreadJoin) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadJoin.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadJoin proto.InternalMessageInfo

func (m *ThreadJoin) GetInviter() string {
	if m != nil {
		return m.Inviter
	}
	return ""
}

func (m *ThreadJoin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ThreadJoin) GetInboxes() []string {
	if m != nil {
		return m.Inboxes
	}
	return nil
}

type ThreadAnnounce struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Inboxes              []string `protobuf:"bytes,2,rep,name=inboxes,proto3" json:"inboxes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadAnnounce) Reset()         { *m = ThreadAnnounce{} }
func (m *ThreadAnnounce) String() string { return proto.CompactTextString(m) }
func (*ThreadAnnounce) ProtoMessage()    {}
func (*ThreadAnnounce) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{6}
}
func (m *ThreadAnnounce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadAnnounce.Unmarshal(m, b)
}
func (m *ThreadAnnounce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadAnnounce.Marshal(b, m, deterministic)
}
func (dst *ThreadAnnounce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadAnnounce.Merge(dst, src)
}
func (m *ThreadAnnounce) XXX_Size() int {
	return xxx_messageInfo_ThreadAnnounce.Size(m)
}
func (m *ThreadAnnounce) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadAnnounce.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadAnnounce proto.InternalMessageInfo

func (m *ThreadAnnounce) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ThreadAnnounce) GetInboxes() []string {
	if m != nil {
		return m.Inboxes
	}
	return nil
}

type ThreadData struct {
	Type                 ThreadData_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ThreadData_Type" json:"type,omitempty"`
	Data                 string          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Key                  []byte          `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Caption              string          `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ThreadData) Reset()         { *m = ThreadData{} }
func (m *ThreadData) String() string { return proto.CompactTextString(m) }
func (*ThreadData) ProtoMessage()    {}
func (*ThreadData) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{7}
}
func (m *ThreadData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadData.Unmarshal(m, b)
}
func (m *ThreadData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadData.Marshal(b, m, deterministic)
}
func (dst *ThreadData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadData.Merge(dst, src)
}
func (m *ThreadData) XXX_Size() int {
	return xxx_messageInfo_ThreadData.Size(m)
}
func (m *ThreadData) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadData.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadData proto.InternalMessageInfo

func (m *ThreadData) GetType() ThreadData_Type {
	if m != nil {
		return m.Type
	}
	return ThreadData_PHOTO
}

func (m *ThreadData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ThreadData) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ThreadData) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

type ThreadAnnotation struct {
	Type                 ThreadAnnotation_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ThreadAnnotation_Type" json:"type,omitempty"`
	Data                 string                `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Caption              string                `protobuf:"bytes,3,opt,name=caption,proto3" json:"caption,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ThreadAnnotation) Reset()         { *m = ThreadAnnotation{} }
func (m *ThreadAnnotation) String() string { return proto.CompactTextString(m) }
func (*ThreadAnnotation) ProtoMessage()    {}
func (*ThreadAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_thread_32e6702dba96c39a, []int{8}
}
func (m *ThreadAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadAnnotation.Unmarshal(m, b)
}
func (m *ThreadAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadAnnotation.Marshal(b, m, deterministic)
}
func (dst *ThreadAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadAnnotation.Merge(dst, src)
}
func (m *ThreadAnnotation) XXX_Size() int {
	return xxx_messageInfo_ThreadAnnotation.Size(m)
}
func (m *ThreadAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadAnnotation proto.InternalMessageInfo

func (m *ThreadAnnotation) GetType() ThreadAnnotation_Type {
	if m != nil {
		return m.Type
	}
	return ThreadAnnotation_COMMENT
}

func (m *ThreadAnnotation) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ThreadAnnotation) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func init() {
	proto.RegisterType((*ThreadEnvelope)(nil), "ThreadEnvelope")
	proto.RegisterType((*ThreadBlock)(nil), "ThreadBlock")
	proto.RegisterType((*ThreadBlockHeader)(nil), "ThreadBlockHeader")
	proto.RegisterType((*ThreadIgnore)(nil), "ThreadIgnore")
	proto.RegisterType((*ThreadInvite)(nil), "ThreadInvite")
	proto.RegisterType((*ThreadJoin)(nil), "ThreadJoin")
	proto.RegisterType((*ThreadAnnounce)(nil), "ThreadAnnounce")
	proto.RegisterType((*ThreadData)(nil), "ThreadData")
	proto.RegisterType((*ThreadAnnotation)(nil), "ThreadAnnotation")
	proto.RegisterEnum("ThreadBlock_Type", ThreadBlock_Type_name, ThreadBlock_Type_value)
	proto.RegisterEnum("ThreadData_Type", ThreadData_Type_name, ThreadData_Type_value)
	proto.RegisterEnum("ThreadAnnotation_Type", ThreadAnnotation_Type_name, ThreadAnnotation_Type_value)
}

func init() { proto.RegisterFile("thread.proto", fileDescriptor_thread_32e6702dba96c39a) }

var fileDescriptor_thread_32e6702dba96c39a = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0x8d, 0x7e, 0xe2, 0xc4, 0x63, 0x63, 0x36, 0xcb, 0x47, 0xd0, 0x97, 0x52, 0x6a, 0x44, 0x0b,
	0x21, 0x17, 0x0a, 0xb8, 0x4f, 0xa0, 0x24, 0x6a, 0xa2, 0x34, 0x91, 0x8a, 0x50, 0x43, 0x29, 0xa5,
	0xb0, 0xb6, 0x37, 0x91, 0xb0, 0xbd, 0x2b, 0xa4, 0x75, 0xa8, 0x9e, 0xa1, 0x17, 0x7d, 0xdd, 0x5e,
	0x96, 0x5d, 0xad, 0x62, 0x39, 0xa1, 0xf4, 0x6e, 0x8f, 0xe6, 0xcc, 0xcc, 0x39, 0x9a, 0x03, 0x43,
	0x91, 0x95, 0x94, 0xcc, 0xbd, 0xa2, 0xe4, 0x82, 0x1f, 0xfd, 0xff, 0xc0, 0xf9, 0xc3, 0x92, 0x9e,
	0x2a, 0x34, 0x5d, 0xdf, 0x9f, 0x12, 0x56, 0xeb, 0xd2, 0x9b, 0xe7, 0x25, 0x91, 0xaf, 0x68, 0x25,
	0xc8, 0xaa, 0x68, 0x08, 0xee, 0x77, 0x18, 0xa5, 0x6a, 0x56, 0xc0, 0x1e, 0xe9, 0x92, 0x17, 0x14,
	0x1f, 0x42, 0xaf, 0x99, 0xee, 0x18, 0x63, 0xe3, 0xb8, 0x9f, 0x68, 0x84, 0x31, 0xd8, 0x19, 0xa9,
	0x32, 0xc7, 0x54, 0x5f, 0xd5, 0x1b, 0x8f, 0x61, 0x30, 0xcb, 0x8b, 0x8c, 0x96, 0x67, 0x4b, 0x3e,
	0x5b, 0x38, 0xd6, 0xd8, 0x38, 0x1e, 0x26, 0xdd, 0x4f, 0xee, 0x6f, 0x03, 0x06, 0xcd, 0x02, 0x85,
	0xf1, 0x09, 0xf4, 0x32, 0x4a, 0xe6, 0xb4, 0x54, 0xd3, 0x07, 0x13, 0xec, 0x75, 0xaa, 0x57, 0xaa,
	0x92, 0x68, 0x06, 0x7e, 0x07, 0xb6, 0xa8, 0x0b, 0xaa, 0x36, 0x8e, 0x26, 0x07, 0x5d, 0xa6, 0x97,
	0xd6, 0x05, 0x4d, 0x54, 0x19, 0x7b, 0xb0, 0x57, 0x90, 0x7a, 0xc9, 0xc9, 0x5c, 0x09, 0x18, 0x4c,
	0xfe, 0xf3, 0x1a, 0xd7, 0x5e, 0xeb, 0xda, 0xf3, 0x59, 0x9d, 0xb4, 0x24, 0xf7, 0x1e, 0x6c, 0xd9,
	0x8d, 0xfb, 0xb0, 0x7b, 0x1b, 0x24, 0x97, 0x01, 0xda, 0xc1, 0x00, 0xbd, 0xf0, 0x32, 0x8a, 0x93,
	0x00, 0x19, 0x78, 0x1f, 0xec, 0xeb, 0x38, 0x8c, 0x90, 0x89, 0x87, 0xb0, 0xef, 0x47, 0x51, 0xfc,
	0x39, 0x3a, 0x0f, 0x90, 0x25, 0xe9, 0x37, 0x81, 0x7f, 0x17, 0x20, 0x5b, 0x52, 0x2e, 0xfc, 0xd4,
	0x47, 0xbb, 0x78, 0x04, 0x20, 0x29, 0xa9, 0x9f, 0x86, 0x71, 0x84, 0x7a, 0x6a, 0x50, 0x74, 0x17,
	0xa6, 0x01, 0x9a, 0xb8, 0x6b, 0x38, 0x78, 0xe1, 0x0d, 0x7b, 0x60, 0xcf, 0x89, 0xa0, 0xda, 0xfd,
	0xd1, 0x0b, 0xa5, 0x69, 0x7b, 0x9f, 0x44, 0xf1, 0xb0, 0x23, 0xcd, 0x95, 0x94, 0x89, 0xca, 0x31,
	0xc7, 0xd6, 0x71, 0x3f, 0x69, 0xa1, 0xbc, 0x13, 0x59, 0x8b, 0x8c, 0x97, 0xca, 0x75, 0x3f, 0xd1,
	0xc8, 0x75, 0x61, 0xd8, 0xac, 0x0d, 0x1f, 0x18, 0x2f, 0xa9, 0xbc, 0xdb, 0x9c, 0x08, 0xa2, 0xaf,
	0xa9, 0xde, 0xee, 0xe4, 0x89, 0xc3, 0x1e, 0x73, 0x41, 0xf1, 0x08, 0xcc, 0x6a, 0xa1, 0x18, 0xc3,
	0xc4, 0xac, 0x16, 0xb2, 0x87, 0x91, 0x15, 0x6d, 0x6f, 0x2d, 0xdf, 0xee, 0x37, 0x80, 0xa6, 0xe7,
	0x9a, 0xe7, 0x4c, 0xea, 0xca, 0x55, 0x6f, 0xa9, 0x07, 0xb7, 0x10, 0x1f, 0xc1, 0xfe, 0xba, 0xa2,
	0x65, 0xa7, 0xff, 0x09, 0x37, 0x5d, 0x53, 0xfe, 0x83, 0x56, 0x8e, 0xd5, 0xb8, 0xd1, 0xd0, 0xfd,
	0xd0, 0xe6, 0xd0, 0x67, 0x8c, 0xaf, 0xd9, 0x8c, 0x6e, 0xcd, 0x31, 0xfe, 0x3e, 0xc7, 0xdc, 0x9e,
	0xf3, 0xcb, 0x68, 0x65, 0x5e, 0x10, 0x41, 0xf0, 0x5b, 0x1d, 0x21, 0x43, 0x45, 0x08, 0x79, 0x9b,
	0x52, 0x37, 0x41, 0xed, 0x2f, 0x32, 0x37, 0xbf, 0x08, 0x23, 0xb0, 0x16, 0xb4, 0xd6, 0x91, 0x96,
	0x4f, 0xb9, 0x74, 0x46, 0x0a, 0x91, 0x73, 0xe6, 0xd8, 0x8d, 0x65, 0x0d, 0xdd, 0x57, 0x9b, 0x44,
	0x7d, 0xba, 0x8a, 0xd3, 0x18, 0xed, 0xc8, 0x88, 0xa4, 0xc1, 0x97, 0x14, 0x19, 0xee, 0x4f, 0x03,
	0xd0, 0xc6, 0x9a, 0x20, 0xb2, 0x03, 0x9f, 0x6c, 0xe9, 0x3a, 0xf4, 0x9e, 0x13, 0xfe, 0xa5, 0xae,
	0xa3, 0xc5, 0xda, 0xd6, 0xf2, 0x5a, 0x6b, 0x19, 0xc0, 0xde, 0x79, 0x7c, 0x7b, 0x1b, 0x44, 0x69,
	0xa3, 0xe6, 0x26, 0xfc, 0x18, 0x20, 0xe3, 0xcc, 0xfe, 0x6a, 0x16, 0xd3, 0x69, 0x4f, 0xe5, 0xed,
	0xfd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xab, 0x29, 0xd6, 0x48, 0x04, 0x00, 0x00,
}
