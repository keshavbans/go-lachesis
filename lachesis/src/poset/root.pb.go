// Code generated by protoc-gen-go. DO NOT EDIT.
// source: root.proto

package poset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RootEvent struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=Hash,json=hash,proto3" json:"Hash,omitempty"`
	CreatorID            uint64   `protobuf:"varint,2,opt,name=CreatorID,json=creatorID,proto3" json:"CreatorID,omitempty"`
	Index                int64    `protobuf:"varint,3,opt,name=Index,json=index,proto3" json:"Index,omitempty"`
	LamportTimestamp     int64    `protobuf:"varint,4,opt,name=LamportTimestamp,json=lamportTimestamp,proto3" json:"LamportTimestamp,omitempty"`
	Round                int64    `protobuf:"varint,5,opt,name=Round,json=round,proto3" json:"Round,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RootEvent) Reset()         { *m = RootEvent{} }
func (m *RootEvent) String() string { return proto.CompactTextString(m) }
func (*RootEvent) ProtoMessage()    {}
func (*RootEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{0}
}

func (m *RootEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RootEvent.Unmarshal(m, b)
}
func (m *RootEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RootEvent.Marshal(b, m, deterministic)
}
func (m *RootEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootEvent.Merge(m, src)
}
func (m *RootEvent) XXX_Size() int {
	return xxx_messageInfo_RootEvent.Size(m)
}
func (m *RootEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RootEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RootEvent proto.InternalMessageInfo

func (m *RootEvent) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *RootEvent) GetCreatorID() uint64 {
	if m != nil {
		return m.CreatorID
	}
	return 0
}

func (m *RootEvent) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RootEvent) GetLamportTimestamp() int64 {
	if m != nil {
		return m.LamportTimestamp
	}
	return 0
}

func (m *RootEvent) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

type Root struct {
	NextRound            int64                 `protobuf:"varint,1,opt,name=NextRound,json=nextRound,proto3" json:"NextRound,omitempty"`
	SelfParent           *RootEvent            `protobuf:"bytes,2,opt,name=SelfParent,json=selfParent,proto3" json:"SelfParent,omitempty"`
	Others               map[string]*RootEvent `protobuf:"bytes,3,rep,name=Others,json=others,proto3" json:"Others,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Root) Reset()         { *m = Root{} }
func (m *Root) String() string { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()    {}
func (*Root) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a043f6ee9336a8, []int{1}
}

func (m *Root) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Root.Unmarshal(m, b)
}
func (m *Root) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Root.Marshal(b, m, deterministic)
}
func (m *Root) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Root.Merge(m, src)
}
func (m *Root) XXX_Size() int {
	return xxx_messageInfo_Root.Size(m)
}
func (m *Root) XXX_DiscardUnknown() {
	xxx_messageInfo_Root.DiscardUnknown(m)
}

var xxx_messageInfo_Root proto.InternalMessageInfo

func (m *Root) GetNextRound() int64 {
	if m != nil {
		return m.NextRound
	}
	return 0
}

func (m *Root) GetSelfParent() *RootEvent {
	if m != nil {
		return m.SelfParent
	}
	return nil
}

func (m *Root) GetOthers() map[string]*RootEvent {
	if m != nil {
		return m.Others
	}
	return nil
}

func init() {
	proto.RegisterType((*RootEvent)(nil), "poset.RootEvent")
	proto.RegisterType((*Root)(nil), "poset.Root")
	proto.RegisterMapType((map[string]*RootEvent)(nil), "poset.Root.OthersEntry")
}

func init() { proto.RegisterFile("root.proto", fileDescriptor_08a043f6ee9336a8) }

var fileDescriptor_08a043f6ee9336a8 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x49, 0xf7, 0x0f, 0xec, 0xac, 0x87, 0x25, 0x08, 0x06, 0xf1, 0xb0, 0xf4, 0x20, 0x8b,
	0x87, 0x55, 0xea, 0x45, 0xbc, 0x6a, 0xc1, 0xa2, 0xa8, 0x44, 0x5f, 0x20, 0xda, 0xc8, 0x16, 0x77,
	0x33, 0x4b, 0x32, 0x2d, 0xed, 0x83, 0xf8, 0x66, 0x3e, 0x90, 0x24, 0x4b, 0x4b, 0x11, 0x7a, 0xcb,
	0xf7, 0xcd, 0x8f, 0x99, 0xef, 0x0b, 0x80, 0x45, 0xa4, 0xba, 0xb7, 0x48, 0xc8, 0x93, 0x1e, 0x9d,
	0xa6, 0xf1, 0x0f, 0x83, 0x4c, 0x22, 0xd2, 0x74, 0xa5, 0x0d, 0x71, 0x0e, 0xf1, 0x83, 0x72, 0x8d,
	0x60, 0x25, 0xab, 0x8e, 0x64, 0xdc, 0x28, 0xd7, 0xf0, 0x33, 0xc8, 0xee, 0xac, 0x56, 0x84, 0x76,
	0x76, 0x2f, 0x46, 0x25, 0xab, 0x62, 0x99, 0x7d, 0x6e, 0x0d, 0x7e, 0x0c, 0xc9, 0xcc, 0xcc, 0xf5,
	0x5a, 0x44, 0x25, 0xab, 0x22, 0x99, 0x2c, 0xbc, 0xe0, 0x17, 0x50, 0x3c, 0xa9, 0xae, 0x47, 0x4b,
	0xef, 0x8b, 0x4e, 0x3b, 0x52, 0x5d, 0x2f, 0xe2, 0x00, 0x14, 0xed, 0x3f, 0xdf, 0x6f, 0x90, 0xb8,
	0x34, 0x73, 0x91, 0x0c, 0x1b, 0xac, 0x17, 0xe3, 0x5f, 0x06, 0xb1, 0xcf, 0xe5, 0xcf, 0x3f, 0xeb,
	0x35, 0x0d, 0x08, 0x0b, 0x48, 0x66, 0xb6, 0x06, 0xbf, 0x02, 0x78, 0xd3, 0xed, 0xd7, 0xab, 0xb2,
	0xda, 0x50, 0x48, 0x97, 0x4f, 0x8a, 0x3a, 0x54, 0xab, 0x77, 0xb5, 0x24, 0xb8, 0x1d, 0xc3, 0x2f,
	0x21, 0x7d, 0xa1, 0x46, 0x5b, 0x27, 0xa2, 0x32, 0xaa, 0xf2, 0xc9, 0xc9, 0x1e, 0x5d, 0x0f, 0x93,
	0xa9, 0x21, 0xbb, 0x91, 0x29, 0x06, 0x71, 0xfa, 0x08, 0xf9, 0x9e, 0xcd, 0x0b, 0x88, 0xbe, 0xf5,
	0x26, 0x24, 0xc9, 0xa4, 0x7f, 0xf2, 0x73, 0x48, 0x56, 0xaa, 0x5d, 0xea, 0x83, 0xe7, 0x87, 0xf1,
	0xed, 0xe8, 0x86, 0x7d, 0xa4, 0xe1, 0xf3, 0xaf, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xfd,
	0x96, 0x79, 0x8a, 0x01, 0x00, 0x00,
}