// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

package util

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SubFixture struct {
	A                    string   `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    string   `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubFixture) Reset()         { *m = SubFixture{} }
func (m *SubFixture) String() string { return proto.CompactTextString(m) }
func (*SubFixture) ProtoMessage()    {}
func (*SubFixture) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}
func (m *SubFixture) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubFixture.Unmarshal(m, b)
}
func (m *SubFixture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubFixture.Marshal(b, m, deterministic)
}
func (m *SubFixture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubFixture.Merge(m, src)
}
func (m *SubFixture) XXX_Size() int {
	return xxx_messageInfo_SubFixture.Size(m)
}
func (m *SubFixture) XXX_DiscardUnknown() {
	xxx_messageInfo_SubFixture.DiscardUnknown(m)
}

var xxx_messageInfo_SubFixture proto.InternalMessageInfo

func (m *SubFixture) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *SubFixture) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

type Fixture struct {
	Boolean              bool                   `protobuf:"varint,1,opt,name=Boolean,proto3" json:"Boolean,omitempty"`
	UInt                 uint32                 `protobuf:"varint,2,opt,name=UInt,proto3" json:"UInt,omitempty"`
	Slice                []string               `protobuf:"bytes,3,rep,name=Slice,proto3" json:"Slice,omitempty"`
	Map                  map[string]uint32      `protobuf:"bytes,4,rep,name=Map,proto3" json:"Map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Sub                  *SubFixture            `protobuf:"bytes,5,opt,name=Sub,proto3" json:"Sub,omitempty"`
	SliceSub             []*SubFixture          `protobuf:"bytes,6,rep,name=SliceSub,proto3" json:"SliceSub,omitempty"`
	MapSub               map[uint32]*SubFixture `protobuf:"bytes,7,rep,name=MapSub,proto3" json:"MapSub,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Fixture) Reset()         { *m = Fixture{} }
func (m *Fixture) String() string { return proto.CompactTextString(m) }
func (*Fixture) ProtoMessage()    {}
func (*Fixture) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}
func (m *Fixture) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fixture.Unmarshal(m, b)
}
func (m *Fixture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fixture.Marshal(b, m, deterministic)
}
func (m *Fixture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fixture.Merge(m, src)
}
func (m *Fixture) XXX_Size() int {
	return xxx_messageInfo_Fixture.Size(m)
}
func (m *Fixture) XXX_DiscardUnknown() {
	xxx_messageInfo_Fixture.DiscardUnknown(m)
}

var xxx_messageInfo_Fixture proto.InternalMessageInfo

func (m *Fixture) GetBoolean() bool {
	if m != nil {
		return m.Boolean
	}
	return false
}

func (m *Fixture) GetUInt() uint32 {
	if m != nil {
		return m.UInt
	}
	return 0
}

func (m *Fixture) GetSlice() []string {
	if m != nil {
		return m.Slice
	}
	return nil
}

func (m *Fixture) GetMap() map[string]uint32 {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *Fixture) GetSub() *SubFixture {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *Fixture) GetSliceSub() []*SubFixture {
	if m != nil {
		return m.SliceSub
	}
	return nil
}

func (m *Fixture) GetMapSub() map[uint32]*SubFixture {
	if m != nil {
		return m.MapSub
	}
	return nil
}

func init() {
	proto.RegisterType((*SubFixture)(nil), "Test.SubFixture")
	proto.RegisterType((*Fixture)(nil), "Test.Fixture")
	proto.RegisterMapType((map[string]uint32)(nil), "Test.Fixture.MapEntry")
	proto.RegisterMapType((map[uint32]*SubFixture)(nil), "Test.Fixture.MapSubEntry")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0xc9, 0xd2, 0xb5, 0xdd, 0x57, 0x07, 0x23, 0x88, 0xc4, 0x9d, 0x4a, 0x0f, 0x92, 0x83,
	0x14, 0x9c, 0x20, 0xa2, 0xa7, 0x15, 0x14, 0x44, 0x7a, 0x49, 0xf5, 0xe2, 0x2d, 0x95, 0x1c, 0x8a,
	0xa5, 0x2d, 0x5d, 0x22, 0xee, 0xe8, 0x7f, 0x2e, 0x49, 0x5a, 0x57, 0xd8, 0x6e, 0xdf, 0x23, 0xbf,
	0xf7, 0xf2, 0xbe, 0x04, 0x40, 0xc9, 0x9d, 0x4a, 0xbb, 0xbe, 0x55, 0x2d, 0xf1, 0xde, 0xe4, 0x4e,
	0x25, 0x0c, 0xa0, 0xd0, 0xe5, 0x73, 0xf5, 0xa3, 0x74, 0x2f, 0xc9, 0x19, 0xa0, 0x2d, 0x45, 0x31,
	0x62, 0x0b, 0x8e, 0xb6, 0x46, 0x65, 0x74, 0xe6, 0x54, 0x96, 0xfc, 0x62, 0x08, 0x46, 0x8e, 0x42,
	0x90, 0xb5, 0x6d, 0x2d, 0x45, 0x63, 0xe9, 0x90, 0x8f, 0x92, 0x10, 0xf0, 0xde, 0x5f, 0x1a, 0x65,
	0x6d, 0x4b, 0x6e, 0x67, 0x72, 0x0e, 0xf3, 0xa2, 0xae, 0x3e, 0x25, 0xc5, 0x31, 0x66, 0x0b, 0xee,
	0x04, 0x61, 0x80, 0x73, 0xd1, 0x51, 0x2f, 0xc6, 0x2c, 0xda, 0x5c, 0xa4, 0xa6, 0x4d, 0x3a, 0xe4,
	0xa7, 0xb9, 0xe8, 0x9e, 0x1a, 0xd5, 0xef, 0xb9, 0x41, 0x48, 0x02, 0xb8, 0xd0, 0x25, 0x9d, 0xc7,
	0x88, 0x45, 0x9b, 0x95, 0x23, 0x0f, 0xa5, 0xb9, 0x39, 0x24, 0xd7, 0x10, 0xda, 0x58, 0x03, 0xfa,
	0x36, 0xf2, 0x18, 0xfc, 0x27, 0xc8, 0x0d, 0xf8, 0xb9, 0xe8, 0x0c, 0x1b, 0x58, 0xf6, 0xf2, 0xe8,
	0xfa, 0x42, 0x97, 0xae, 0xc1, 0x00, 0xae, 0xef, 0x20, 0x1c, 0x5b, 0x91, 0x15, 0xe0, 0x2f, 0xb9,
	0x1f, 0x1e, 0xca, 0x8c, 0x66, 0xc5, 0x6f, 0x51, 0x6b, 0x39, 0xec, 0xed, 0xc4, 0xc3, 0xec, 0x1e,
	0xad, 0x5f, 0x21, 0x9a, 0xc4, 0x4d, 0xad, 0x4b, 0x67, 0xbd, 0x9a, 0x5a, 0x4f, 0xd5, 0x3e, 0x84,
	0x65, 0xe1, 0x87, 0x9f, 0x3e, 0x6a, 0x55, 0xd5, 0xa5, 0x6f, 0x3f, 0xf1, 0xf6, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x48, 0x84, 0x85, 0x75, 0xd2, 0x01, 0x00, 0x00,
}
