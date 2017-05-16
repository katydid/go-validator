// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

/*
Package prototests is a generated protocol buffer package.

It is generated from these files:
	msg.proto

It has these top-level messages:
	BigMsg
	SmallMsg
	Packed
*/
package prototests

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_gogo_protobuf_protoc_gen_gogo_descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import compress_gzip "compress/gzip"
import bytes "bytes"
import io_ioutil "io/ioutil"

import strings "strings"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// BigMsg contains a field and a message field.
type BigMsg struct {
	Field            *int64    `protobuf:"varint,1,opt,name=Field" json:"Field,omitempty"`
	Msg              *SmallMsg `protobuf:"bytes,3,opt,name=Msg" json:"Msg,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BigMsg) Reset()                    { *m = BigMsg{} }
func (m *BigMsg) String() string            { return proto.CompactTextString(m) }
func (*BigMsg) ProtoMessage()               {}
func (*BigMsg) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{0} }

func (m *BigMsg) GetField() int64 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *BigMsg) GetMsg() *SmallMsg {
	if m != nil {
		return m.Msg
	}
	return nil
}

// SmallMsg only contains some native fields.
type SmallMsg struct {
	ScarBusStop      *string  `protobuf:"bytes,1,opt,name=ScarBusStop" json:"ScarBusStop,omitempty"`
	FlightParachute  []uint32 `protobuf:"fixed32,12,rep,name=FlightParachute" json:"FlightParachute,omitempty"`
	MapShark         *string  `protobuf:"bytes,18,opt,name=MapShark" json:"MapShark,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SmallMsg) Reset()                    { *m = SmallMsg{} }
func (m *SmallMsg) String() string            { return proto.CompactTextString(m) }
func (*SmallMsg) ProtoMessage()               {}
func (*SmallMsg) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{1} }

func (m *SmallMsg) GetScarBusStop() string {
	if m != nil && m.ScarBusStop != nil {
		return *m.ScarBusStop
	}
	return ""
}

func (m *SmallMsg) GetFlightParachute() []uint32 {
	if m != nil {
		return m.FlightParachute
	}
	return nil
}

func (m *SmallMsg) GetMapShark() string {
	if m != nil && m.MapShark != nil {
		return *m.MapShark
	}
	return ""
}

// Packed contains some repeated packed fields.
type Packed struct {
	Ints             []int64   `protobuf:"varint,4,rep,packed,name=Ints" json:"Ints,omitempty"`
	Floats           []float64 `protobuf:"fixed64,5,rep,packed,name=Floats" json:"Floats,omitempty"`
	Uints            []uint32  `protobuf:"varint,6,rep,packed,name=Uints" json:"Uints,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Packed) Reset()                    { *m = Packed{} }
func (m *Packed) String() string            { return proto.CompactTextString(m) }
func (*Packed) ProtoMessage()               {}
func (*Packed) Descriptor() ([]byte, []int) { return fileDescriptorMsg, []int{2} }

func (m *Packed) GetInts() []int64 {
	if m != nil {
		return m.Ints
	}
	return nil
}

func (m *Packed) GetFloats() []float64 {
	if m != nil {
		return m.Floats
	}
	return nil
}

func (m *Packed) GetUints() []uint32 {
	if m != nil {
		return m.Uints
	}
	return nil
}

func init() {
	proto.RegisterType((*BigMsg)(nil), "prototests.BigMsg")
	proto.RegisterType((*SmallMsg)(nil), "prototests.SmallMsg")
	proto.RegisterType((*Packed)(nil), "prototests.Packed")
}
func (this *BigMsg) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *SmallMsg) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func (this *Packed) Description() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	return MsgDescription()
}
func MsgDescription() (desc *github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet) {
	d := &github_com_gogo_protobuf_protoc_gen_gogo_descriptor.FileDescriptorSet{}
	var gzipped = []byte{
		// 3714 bytes of a gzipped FileDescriptorSet
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x5a, 0x5d, 0x6c, 0x23, 0xd7,
		0x75, 0x36, 0x7f, 0x45, 0x1e, 0x52, 0xd4, 0xe8, 0x4a, 0xd6, 0xce, 0xca, 0xf1, 0xae, 0x96, 0xf1,
		0x8f, 0x6c, 0x37, 0xdc, 0x60, 0xed, 0x5d, 0xaf, 0xb9, 0x4d, 0x0c, 0x8a, 0xa2, 0x14, 0x6e, 0x45,
		0x91, 0x19, 0x4a, 0xf6, 0x3a, 0x7d, 0x18, 0x5c, 0x0d, 0x2f, 0xa9, 0xd9, 0x1d, 0xce, 0x30, 0x33,
		0xc3, 0x5d, 0xcb, 0x4f, 0x2e, 0xdc, 0x1f, 0x04, 0x45, 0xff, 0x0b, 0x34, 0x71, 0x1d, 0xb7, 0x0d,
		0xd0, 0xba, 0x4d, 0xfa, 0x93, 0xf4, 0x27, 0x2d, 0xfa, 0xd4, 0x97, 0xb4, 0x7d, 0x2a, 0x90, 0xf7,
		0x3e, 0x34, 0xad, 0x81, 0xa6, 0xad, 0xdb, 0xa6, 0x8d, 0x81, 0x16, 0xf0, 0x4b, 0x70, 0xff, 0x86,
		0xc3, 0x1f, 0x69, 0xa8, 0x00, 0x8e, 0x9f, 0xa4, 0x7b, 0xee, 0xf9, 0xbe, 0x39, 0x73, 0xee, 0xb9,
		0xe7, 0x9c, 0x7b, 0x39, 0xf0, 0x2f, 0xd7, 0x61, 0xa3, 0xe7, 0x38, 0x3d, 0x8b, 0x5c, 0x1d, 0xb8,
		0x8e, 0xef, 0x1c, 0x0d, 0xbb, 0x57, 0x3b, 0xc4, 0x33, 0x5c, 0x73, 0xe0, 0x3b, 0x6e, 0x89, 0xc9,
		0xd0, 0x12, 0xd7, 0x28, 0x49, 0x8d, 0x62, 0x03, 0x96, 0x77, 0x4c, 0x8b, 0x6c, 0x07, 0x8a, 0x6d,
		0xe2, 0xa3, 0x9b, 0x90, 0xec, 0x9a, 0x16, 0x51, 0x63, 0x1b, 0x89, 0xcd, 0xdc, 0xb5, 0xc7, 0x4a,
		0x13, 0xa0, 0xd2, 0x38, 0xa2, 0x45, 0xc5, 0x1a, 0x43, 0x14, 0xdf, 0x4d, 0xc2, 0xca, 0x8c, 0x59,
		0x84, 0x20, 0x69, 0xe3, 0x3e, 0x65, 0x8c, 0x6d, 0x66, 0x35, 0xf6, 0x3f, 0x52, 0x61, 0x61, 0x80,
		0x8d, 0x7b, 0xb8, 0x47, 0xd4, 0x38, 0x13, 0xcb, 0x21, 0xba, 0x04, 0xd0, 0x21, 0x03, 0x62, 0x77,
		0x88, 0x6d, 0x9c, 0xa8, 0x89, 0x8d, 0xc4, 0x66, 0x56, 0x0b, 0x49, 0xd0, 0x33, 0xb0, 0x3c, 0x18,
		0x1e, 0x59, 0xa6, 0xa1, 0x87, 0xd4, 0x60, 0x23, 0xb1, 0x99, 0xd2, 0x14, 0x3e, 0xb1, 0x3d, 0x52,
		0x7e, 0x12, 0x96, 0x1e, 0x10, 0x7c, 0x2f, 0xac, 0x9a, 0x63, 0xaa, 0x05, 0x2a, 0x0e, 0x29, 0x56,
		0x21, 0xdf, 0x27, 0x9e, 0x87, 0x7b, 0x44, 0xf7, 0x4f, 0x06, 0x44, 0x4d, 0xb2, 0xb7, 0xdf, 0x98,
		0x7a, 0xfb, 0xc9, 0x37, 0xcf, 0x09, 0xd4, 0xc1, 0xc9, 0x80, 0xa0, 0x0a, 0x64, 0x89, 0x3d, 0xec,
		0x73, 0x86, 0xd4, 0x29, 0xfe, 0xab, 0xd9, 0xc3, 0xfe, 0x24, 0x4b, 0x86, 0xc2, 0x04, 0xc5, 0x82,
		0x47, 0xdc, 0xfb, 0xa6, 0x41, 0xd4, 0x34, 0x23, 0x78, 0x72, 0x8a, 0xa0, 0xcd, 0xe7, 0x27, 0x39,
		0x24, 0x0e, 0x55, 0x21, 0x4b, 0x5e, 0xf5, 0x89, 0xed, 0x99, 0x8e, 0xad, 0x2e, 0x30, 0x92, 0xc7,
		0x67, 0xac, 0x22, 0xb1, 0x3a, 0x93, 0x14, 0x23, 0x1c, 0xba, 0x01, 0x0b, 0xce, 0xc0, 0x37, 0x1d,
		0xdb, 0x53, 0x33, 0x1b, 0xb1, 0xcd, 0xdc, 0xb5, 0x8f, 0xcd, 0x0c, 0x84, 0x26, 0xd7, 0xd1, 0xa4,
		0x32, 0xaa, 0x83, 0xe2, 0x39, 0x43, 0xd7, 0x20, 0xba, 0xe1, 0x74, 0x88, 0x6e, 0xda, 0x5d, 0x47,
		0xcd, 0x32, 0x82, 0xcb, 0xd3, 0x2f, 0xc2, 0x14, 0xab, 0x4e, 0x87, 0xd4, 0xed, 0xae, 0xa3, 0x15,
		0xbc, 0xb1, 0x31, 0x5a, 0x83, 0xb4, 0x77, 0x62, 0xfb, 0xf8, 0x55, 0x35, 0xcf, 0x22, 0x44, 0x8c,
		0x8a, 0xff, 0x97, 0x82, 0xa5, 0x79, 0x42, 0xec, 0x16, 0xa4, 0xba, 0xf4, 0x2d, 0xd5, 0xf8, 0x79,
		0x7c, 0xc0, 0x31, 0xe3, 0x4e, 0x4c, 0xff, 0x90, 0x4e, 0xac, 0x40, 0xce, 0x26, 0x9e, 0x4f, 0x3a,
		0x3c, 0x22, 0x12, 0x73, 0xc6, 0x14, 0x70, 0xd0, 0x74, 0x48, 0x25, 0x7f, 0xa8, 0x90, 0xba, 0x03,
		0x4b, 0x81, 0x49, 0xba, 0x8b, 0xed, 0x9e, 0x8c, 0xcd, 0xab, 0x51, 0x96, 0x94, 0x6a, 0x12, 0xa7,
		0x51, 0x98, 0x56, 0x20, 0x63, 0x63, 0xb4, 0x0d, 0xe0, 0xd8, 0xc4, 0xe9, 0xea, 0x1d, 0x62, 0x58,
		0x6a, 0xe6, 0x14, 0x2f, 0x35, 0xa9, 0xca, 0x94, 0x97, 0x1c, 0x2e, 0x35, 0x2c, 0xf4, 0xc2, 0x28,
		0xd4, 0x16, 0x4e, 0x89, 0x94, 0x06, 0xdf, 0x64, 0x53, 0xd1, 0x76, 0x08, 0x05, 0x97, 0xd0, 0xb8,
		0x27, 0x1d, 0xf1, 0x66, 0x59, 0x66, 0x44, 0x29, 0xf2, 0xcd, 0x34, 0x01, 0xe3, 0x2f, 0xb6, 0xe8,
		0x86, 0x87, 0xe8, 0xe3, 0x10, 0x08, 0x74, 0x16, 0x56, 0xc0, 0xb2, 0x50, 0x5e, 0x0a, 0xf7, 0x71,
		0x9f, 0xac, 0xdf, 0x84, 0xc2, 0xb8, 0x7b, 0xd0, 0x2a, 0xa4, 0x3c, 0x1f, 0xbb, 0x3e, 0x8b, 0xc2,
		0x94, 0xc6, 0x07, 0x48, 0x81, 0x04, 0xb1, 0x3b, 0x2c, 0xcb, 0xa5, 0x34, 0xfa, 0xef, 0xfa, 0xf3,
		0xb0, 0x38, 0xf6, 0xf8, 0x79, 0x81, 0xc5, 0x2f, 0xa6, 0x61, 0x75, 0x56, 0xcc, 0xcd, 0x0c, 0xff,
		0x35, 0x48, 0xdb, 0xc3, 0xfe, 0x11, 0x71, 0xd5, 0x04, 0x63, 0x10, 0x23, 0x54, 0x81, 0x94, 0x85,
		0x8f, 0x88, 0xa5, 0x26, 0x37, 0x62, 0x9b, 0x85, 0x6b, 0xcf, 0xcc, 0x15, 0xd5, 0xa5, 0x3d, 0x0a,
		0xd1, 0x38, 0x12, 0x7d, 0x1a, 0x92, 0x22, 0xc5, 0x51, 0x86, 0xa7, 0xe7, 0x63, 0xa0, 0xb1, 0xa8,
		0x31, 0x1c, 0x7a, 0x04, 0xb2, 0xf4, 0x2f, 0xf7, 0x6d, 0x9a, 0xd9, 0x9c, 0xa1, 0x02, 0xea, 0x57,
		0xb4, 0x0e, 0x19, 0x16, 0x66, 0x1d, 0x22, 0x4b, 0x43, 0x30, 0xa6, 0x0b, 0xd3, 0x21, 0x5d, 0x3c,
		0xb4, 0x7c, 0xfd, 0x3e, 0xb6, 0x86, 0x84, 0x05, 0x4c, 0x56, 0xcb, 0x0b, 0xe1, 0x4b, 0x54, 0x86,
		0x2e, 0x43, 0x8e, 0x47, 0xa5, 0x69, 0x77, 0xc8, 0xab, 0x2c, 0xfb, 0xa4, 0x34, 0x1e, 0xa8, 0x75,
		0x2a, 0xa1, 0x8f, 0xbf, 0xeb, 0x39, 0xb6, 0x5c, 0x5a, 0xf6, 0x08, 0x2a, 0x60, 0x8f, 0x7f, 0x7e,
		0x32, 0xf1, 0x3d, 0x3a, 0xfb, 0xf5, 0x26, 0x63, 0xb1, 0xf8, 0xcd, 0x38, 0x24, 0xd9, 0x7e, 0x5b,
		0x82, 0xdc, 0xc1, 0x2b, 0xad, 0x9a, 0xbe, 0xdd, 0x3c, 0xdc, 0xda, 0xab, 0x29, 0x31, 0x54, 0x00,
		0x60, 0x82, 0x9d, 0xbd, 0x66, 0xe5, 0x40, 0x89, 0x07, 0xe3, 0xfa, 0xfe, 0xc1, 0x8d, 0xe7, 0x94,
		0x44, 0x00, 0x38, 0xe4, 0x82, 0x64, 0x58, 0xe1, 0xd9, 0x6b, 0x4a, 0x0a, 0x29, 0x90, 0xe7, 0x04,
		0xf5, 0x3b, 0xb5, 0xed, 0x1b, 0xcf, 0x29, 0xe9, 0x71, 0xc9, 0xb3, 0xd7, 0x94, 0x05, 0xb4, 0x08,
		0x59, 0x26, 0xd9, 0x6a, 0x36, 0xf7, 0x94, 0x4c, 0xc0, 0xd9, 0x3e, 0xd0, 0xea, 0xfb, 0xbb, 0x4a,
		0x36, 0xe0, 0xdc, 0xd5, 0x9a, 0x87, 0x2d, 0x05, 0x02, 0x86, 0x46, 0xad, 0xdd, 0xae, 0xec, 0xd6,
		0x94, 0x5c, 0xa0, 0xb1, 0xf5, 0xca, 0x41, 0xad, 0xad, 0xe4, 0xc7, 0xcc, 0x7a, 0xf6, 0x9a, 0xb2,
		0x18, 0x3c, 0xa2, 0xb6, 0x7f, 0xd8, 0x50, 0x0a, 0x68, 0x19, 0x16, 0xf9, 0x23, 0xa4, 0x11, 0x4b,
		0x13, 0xa2, 0x1b, 0xcf, 0x29, 0xca, 0xc8, 0x10, 0xce, 0xb2, 0x3c, 0x26, 0xb8, 0xf1, 0x9c, 0x82,
		0x8a, 0x55, 0x48, 0xb1, 0xe8, 0x42, 0x08, 0x0a, 0x7b, 0x95, 0xad, 0xda, 0x9e, 0xde, 0x6c, 0x1d,
		0xd4, 0x9b, 0xfb, 0x95, 0x3d, 0x25, 0x36, 0x92, 0x69, 0xb5, 0xcf, 0x1e, 0xd6, 0xb5, 0xda, 0xb6,
		0x12, 0x0f, 0xcb, 0x5a, 0xb5, 0xca, 0x41, 0x6d, 0x5b, 0x49, 0x14, 0x0d, 0x58, 0x9d, 0x95, 0x67,
		0x66, 0xee, 0x8c, 0xd0, 0x12, 0xc7, 0x4f, 0x59, 0x62, 0xc6, 0x35, 0xb5, 0xc4, 0x5f, 0x89, 0xc1,
		0xca, 0x8c, 0x5c, 0x3b, 0xf3, 0x21, 0x2f, 0x42, 0x8a, 0x87, 0x28, 0xaf, 0x3e, 0x4f, 0xcd, 0x4c,
		0xda, 0x2c, 0x60, 0xa7, 0x2a, 0x10, 0xc3, 0x85, 0x2b, 0x70, 0xe2, 0x94, 0x0a, 0x4c, 0x29, 0xa6,
		0x8c, 0x7c, 0x23, 0x06, 0xea, 0x69, 0xdc, 0x11, 0x89, 0x22, 0x3e, 0x96, 0x28, 0x6e, 0x4d, 0x1a,
		0x70, 0xe5, 0xf4, 0x77, 0x98, 0xb2, 0xe2, 0x9d, 0x18, 0xac, 0xcd, 0x6e, 0x54, 0x66, 0xda, 0xf0,
		0x69, 0x48, 0xf7, 0x89, 0x7f, 0xec, 0xc8, 0x62, 0xfd, 0xc4, 0x8c, 0x12, 0x40, 0xa7, 0x27, 0x7d,
		0x25, 0x50, 0xe1, 0x1a, 0x92, 0x38, 0xad, 0xdb, 0xe0, 0xd6, 0x4c, 0x59, 0xfa, 0x85, 0x38, 0x3c,
		0x3c, 0x93, 0x7c, 0xa6, 0xa1, 0x8f, 0x02, 0x98, 0xf6, 0x60, 0xe8, 0xf3, 0x82, 0xcc, 0xf3, 0x53,
		0x96, 0x49, 0xd8, 0xde, 0xa7, 0xb9, 0x67, 0xe8, 0x07, 0xf3, 0x09, 0x36, 0x0f, 0x5c, 0xc4, 0x14,
		0x6e, 0x8e, 0x0c, 0x4d, 0x32, 0x43, 0x2f, 0x9d, 0xf2, 0xa6, 0x53, 0xb5, 0xee, 0x93, 0xa0, 0x18,
		0x96, 0x49, 0x6c, 0x5f, 0xf7, 0x7c, 0x97, 0xe0, 0xbe, 0x69, 0xf7, 0x58, 0x02, 0xce, 0x94, 0x53,
		0x5d, 0x6c, 0x79, 0x44, 0x5b, 0xe2, 0xd3, 0x6d, 0x39, 0x4b, 0x11, 0xac, 0xca, 0xb8, 0x21, 0x44,
		0x7a, 0x0c, 0xc1, 0xa7, 0x03, 0x44, 0xf1, 0x6b, 0x0b, 0x90, 0x0b, 0xb5, 0x75, 0xe8, 0x0a, 0xe4,
		0xef, 0xe2, 0xfb, 0x58, 0x97, 0xad, 0x3a, 0xf7, 0x44, 0x8e, 0xca, 0x5a, 0xa2, 0x5d, 0xff, 0x24,
		0xac, 0x32, 0x15, 0x67, 0xe8, 0x13, 0x57, 0x37, 0x2c, 0xec, 0x79, 0xcc, 0x69, 0x19, 0xa6, 0x8a,
		0xe8, 0x5c, 0x93, 0x4e, 0x55, 0xe5, 0x0c, 0xba, 0x0e, 0x2b, 0x0c, 0xd1, 0x1f, 0x5a, 0xbe, 0x39,
		0xb0, 0x88, 0x4e, 0x0f, 0x0f, 0x1e, 0x4b, 0xc4, 0x81, 0x65, 0xcb, 0x54, 0xa3, 0x21, 0x14, 0xa8,
		0x45, 0x1e, 0xda, 0x86, 0x47, 0x19, 0xac, 0x47, 0x6c, 0xe2, 0x62, 0x9f, 0xe8, 0xe4, 0xf3, 0x43,
		0x6c, 0x79, 0x3a, 0xb6, 0x3b, 0xfa, 0x31, 0xf6, 0x8e, 0xd5, 0x55, 0x4a, 0xb0, 0x15, 0x57, 0x63,
		0xda, 0x45, 0xaa, 0xb8, 0x2b, 0xf4, 0x6a, 0x4c, 0xad, 0x62, 0x77, 0x3e, 0x83, 0xbd, 0x63, 0x54,
		0x86, 0x35, 0xc6, 0xe2, 0xf9, 0xae, 0x69, 0xf7, 0x74, 0xe3, 0x98, 0x18, 0xf7, 0xf4, 0xa1, 0xdf,
		0xbd, 0xa9, 0x3e, 0x12, 0x7e, 0x3e, 0xb3, 0xb0, 0xcd, 0x74, 0xaa, 0x54, 0xe5, 0xd0, 0xef, 0xde,
		0x44, 0x6d, 0xc8, 0xd3, 0xc5, 0xe8, 0x9b, 0xaf, 0x11, 0xbd, 0xeb, 0xb8, 0xac, 0xb2, 0x14, 0x66,
		0xec, 0xec, 0x90, 0x07, 0x4b, 0x4d, 0x01, 0x68, 0x38, 0x1d, 0x52, 0x4e, 0xb5, 0x5b, 0xb5, 0xda,
		0xb6, 0x96, 0x93, 0x2c, 0x3b, 0x8e, 0x4b, 0x03, 0xaa, 0xe7, 0x04, 0x0e, 0xce, 0xf1, 0x80, 0xea,
		0x39, 0xd2, 0xbd, 0xd7, 0x61, 0xc5, 0x30, 0xf8, 0x3b, 0x9b, 0x86, 0x2e, 0x5a, 0x7c, 0x4f, 0x55,
		0xc6, 0x9c, 0x65, 0x18, 0xbb, 0x5c, 0x41, 0xc4, 0xb8, 0x87, 0x5e, 0x80, 0x87, 0x47, 0xce, 0x0a,
		0x03, 0x97, 0xa7, 0xde, 0x72, 0x12, 0x7a, 0x1d, 0x56, 0x06, 0x27, 0xd3, 0x40, 0x34, 0xf6, 0xc4,
		0xc1, 0xc9, 0x24, 0xec, 0x71, 0x76, 0x6c, 0x73, 0x89, 0x81, 0x7d, 0xd2, 0x51, 0x2f, 0x84, 0xb5,
		0x43, 0x13, 0xe8, 0x2a, 0x28, 0x86, 0xa1, 0x13, 0x1b, 0x1f, 0x59, 0x44, 0xc7, 0x2e, 0xb1, 0xb1,
		0xa7, 0x5e, 0x0e, 0x2b, 0x17, 0x0c, 0xa3, 0xc6, 0x66, 0x2b, 0x6c, 0x12, 0x3d, 0x0d, 0xcb, 0xce,
		0xd1, 0x5d, 0x83, 0x47, 0x96, 0x3e, 0x70, 0x49, 0xd7, 0x7c, 0x55, 0x7d, 0x8c, 0xb9, 0x69, 0x89,
		0x4e, 0xb0, 0xb8, 0x6a, 0x31, 0x31, 0x7a, 0x0a, 0x14, 0xc3, 0x3b, 0xc6, 0xee, 0x80, 0x95, 0x76,
		0x6f, 0x80, 0x0d, 0xa2, 0x3e, 0xce, 0x55, 0xb9, 0x7c, 0x5f, 0x8a, 0x69, 0x64, 0x7b, 0x0f, 0xcc,
		0xae, 0x2f, 0x19, 0x9f, 0xe4, 0x91, 0xcd, 0x64, 0x82, 0xed, 0x0e, 0xac, 0x0e, 0x6d, 0xd3, 0xf6,
		0x89, 0x3b, 0x70, 0x09, 0x6d, 0xe2, 0xf9, 0x4e, 0x54, 0xff, 0x75, 0xe1, 0x94, 0x36, 0xfc, 0x30,
		0xac, 0xcd, 0x03, 0x40, 0x5b, 0x19, 0x4e, 0x0b, 0x8b, 0x65, 0xc8, 0x87, 0xe3, 0x02, 0x65, 0x81,
		0x47, 0x86, 0x12, 0xa3, 0x35, 0xb6, 0xda, 0xdc, 0xa6, 0xd5, 0xf1, 0x73, 0x35, 0x25, 0x4e, 0xab,
		0xf4, 0x5e, 0xfd, 0xa0, 0xa6, 0x6b, 0x87, 0xfb, 0x07, 0xf5, 0x46, 0x4d, 0x49, 0x3c, 0x9d, 0xcd,
		0x7c, 0x77, 0x41, 0x79, 0xfd, 0xf5, 0xd7, 0x5f, 0x8f, 0x17, 0xbf, 0x15, 0x87, 0xc2, 0x78, 0x67,
		0x8c, 0x7e, 0x1c, 0x2e, 0xc8, 0x63, 0xac, 0x47, 0x7c, 0xfd, 0x81, 0xe9, 0xb2, 0x50, 0xed, 0x63,
		0xde, 0x5b, 0x06, 0x5e, 0x5e, 0x15, 0x5a, 0x6d, 0xe2, 0xbf, 0x6c, 0xba, 0x34, 0x10, 0xfb, 0xd8,
		0x47, 0x7b, 0x70, 0xd9, 0x76, 0x74, 0xcf, 0xc7, 0x76, 0x07, 0xbb, 0x1d, 0x7d, 0x74, 0x81, 0xa0,
		0x63, 0xc3, 0x20, 0x9e, 0xe7, 0xf0, 0x12, 0x11, 0xb0, 0x7c, 0xcc, 0x76, 0xda, 0x42, 0x79, 0x94,
		0x3b, 0x2b, 0x42, 0x75, 0x22, 0x22, 0x12, 0xa7, 0x45, 0xc4, 0x23, 0x90, 0xed, 0xe3, 0x81, 0x4e,
		0x6c, 0xdf, 0x3d, 0x61, 0xfd, 0x5c, 0x46, 0xcb, 0xf4, 0xf1, 0xa0, 0x46, 0xc7, 0x1f, 0xde, 0x1a,
		0x84, 0xfd, 0xf8, 0x8f, 0x09, 0xc8, 0x87, 0x7b, 0x3a, 0xda, 0x22, 0x1b, 0x2c, 0x7f, 0xc7, 0xd8,
		0x0e, 0xff, 0xf8, 0x99, 0x1d, 0x60, 0xa9, 0x4a, 0x13, 0x7b, 0x39, 0xcd, 0x3b, 0x2d, 0x8d, 0x23,
		0x69, 0x51, 0xa5, 0x7b, 0x9a, 0xf0, 0xfe, 0x3d, 0xa3, 0x89, 0x11, 0xda, 0x85, 0xf4, 0x5d, 0x8f,
		0x71, 0xa7, 0x19, 0xf7, 0x63, 0x67, 0x73, 0xdf, 0x6e, 0x33, 0xf2, 0xec, 0xed, 0xb6, 0xbe, 0xdf,
		0xd4, 0x1a, 0x95, 0x3d, 0x4d, 0xc0, 0xd1, 0x45, 0x48, 0x5a, 0xf8, 0xb5, 0x93, 0xf1, 0x12, 0xc0,
		0x44, 0xf3, 0x3a, 0xfe, 0x22, 0x24, 0x1f, 0x10, 0x7c, 0x6f, 0x3c, 0xf1, 0x32, 0xd1, 0x87, 0x18,
		0xfa, 0x57, 0x21, 0xc5, 0xfc, 0x85, 0x00, 0x84, 0xc7, 0x94, 0x87, 0x50, 0x06, 0x92, 0xd5, 0xa6,
		0x46, 0xc3, 0x5f, 0x81, 0x3c, 0x97, 0xea, 0xad, 0x7a, 0xad, 0x5a, 0x53, 0xe2, 0xc5, 0xeb, 0x90,
		0xe6, 0x4e, 0xa0, 0x5b, 0x23, 0x70, 0x83, 0xf2, 0x90, 0x18, 0x0a, 0x8e, 0x98, 0x9c, 0x3d, 0x6c,
		0x6c, 0xd5, 0x34, 0x25, 0x1e, 0x5e, 0x5e, 0x0f, 0xf2, 0xe1, 0x76, 0xee, 0x47, 0x13, 0x53, 0x7f,
		0x1d, 0x83, 0x5c, 0xa8, 0x3d, 0xa3, 0x8d, 0x01, 0xb6, 0x2c, 0xe7, 0x81, 0x8e, 0x2d, 0x13, 0x7b,
		0x22, 0x28, 0x80, 0x89, 0x2a, 0x54, 0x32, 0xef, 0xa2, 0xfd, 0x48, 0x8c, 0x7f, 0x3b, 0x06, 0xca,
		0x64, 0x6b, 0x37, 0x61, 0x60, 0xec, 0x23, 0x35, 0xf0, 0xad, 0x18, 0x14, 0xc6, 0xfb, 0xb9, 0x09,
		0xf3, 0xae, 0x7c, 0xa4, 0xe6, 0xfd, 0x53, 0x1c, 0x16, 0xc7, 0xba, 0xb8, 0x79, 0xad, 0xfb, 0x3c,
		0x2c, 0x9b, 0x1d, 0xd2, 0x1f, 0x38, 0x3e, 0xb1, 0x8d, 0x13, 0xdd, 0x22, 0xf7, 0x89, 0xa5, 0x16,
		0x59, 0xa2, 0xb8, 0x7a, 0x76, 0x9f, 0x58, 0xaa, 0x8f, 0x70, 0x7b, 0x14, 0x56, 0x5e, 0xa9, 0x6f,
		0xd7, 0x1a, 0xad, 0xe6, 0x41, 0x6d, 0xbf, 0xfa, 0x8a, 0x7e, 0xb8, 0xff, 0x13, 0xfb, 0xcd, 0x97,
		0xf7, 0x35, 0xc5, 0x9c, 0x50, 0xfb, 0x10, 0xb7, 0x7a, 0x0b, 0x94, 0x49, 0xa3, 0xd0, 0x05, 0x98,
		0x65, 0x96, 0xf2, 0x10, 0x5a, 0x81, 0xa5, 0xfd, 0xa6, 0xde, 0xae, 0x6f, 0xd7, 0xf4, 0xda, 0xce,
		0x4e, 0xad, 0x7a, 0xd0, 0xe6, 0x07, 0xe7, 0x40, 0xfb, 0x60, 0x7c, 0x53, 0xbf, 0x99, 0x80, 0x95,
		0x19, 0x96, 0xa0, 0x8a, 0xe8, 0xd9, 0xf9, 0x31, 0xe2, 0x13, 0xf3, 0x58, 0x5f, 0xa2, 0x5d, 0x41,
		0x0b, 0xbb, 0xbe, 0x68, 0xf1, 0x9f, 0x02, 0xea, 0x25, 0xdb, 0x37, 0xbb, 0x26, 0x71, 0xc5, 0x3d,
		0x03, 0x6f, 0xe4, 0x97, 0x46, 0x72, 0x7e, 0xd5, 0xf0, 0x63, 0x80, 0x06, 0x8e, 0x67, 0xfa, 0xe6,
		0x7d, 0xa2, 0x9b, 0xb6, 0xbc, 0x94, 0xa0, 0x8d, 0x7d, 0x52, 0x53, 0xe4, 0x4c, 0xdd, 0xf6, 0x03,
		0x6d, 0x9b, 0xf4, 0xf0, 0x84, 0x36, 0x4d, 0xe0, 0x09, 0x4d, 0x91, 0x33, 0x81, 0xf6, 0x15, 0xc8,
		0x77, 0x9c, 0x21, 0x6d, 0x93, 0xb8, 0x1e, 0xad, 0x17, 0x31, 0x2d, 0xc7, 0x65, 0x81, 0x8a, 0xe8,
		0x63, 0x47, 0xb7, 0x21, 0x79, 0x2d, 0xc7, 0x65, 0x5c, 0xe5, 0x49, 0x58, 0xc2, 0xbd, 0x9e, 0x4b,
		0xc9, 0x25, 0x11, 0xef, 0xcc, 0x0b, 0x81, 0x98, 0x29, 0xae, 0xdf, 0x86, 0x8c, 0xf4, 0x03, 0x2d,
		0xc9, 0xd4, 0x13, 0xfa, 0x80, 0xdf, 0x49, 0xc5, 0x37, 0xb3, 0x5a, 0xc6, 0x96, 0x93, 0x57, 0x20,
		0x6f, 0x7a, 0xfa, 0xe8, 0x72, 0x34, 0xbe, 0x11, 0xdf, 0xcc, 0x68, 0x39, 0xd3, 0x0b, 0x6e, 0xc3,
		0x8a, 0xef, 0xc4, 0xa1, 0x30, 0x7e, 0xb9, 0x8b, 0xb6, 0x21, 0x63, 0x39, 0x06, 0x66, 0xa1, 0xc5,
		0x7f, 0x59, 0xd8, 0x8c, 0xb8, 0x0f, 0x2e, 0xed, 0x09, 0x7d, 0x2d, 0x40, 0xae, 0xff, 0x43, 0x0c,
		0x32, 0x52, 0x8c, 0xd6, 0x20, 0x39, 0xc0, 0xfe, 0x31, 0xa3, 0x4b, 0x6d, 0xc5, 0x95, 0x98, 0xc6,
		0xc6, 0x54, 0xee, 0x0d, 0xb0, 0xcd, 0x42, 0x40, 0xc8, 0xe9, 0x98, 0xae, 0xab, 0x45, 0x70, 0x87,
		0xb5, 0xfd, 0x4e, 0xbf, 0x4f, 0x6c, 0xdf, 0x93, 0xeb, 0x2a, 0xe4, 0x55, 0x21, 0x46, 0xcf, 0xc0,
		0xb2, 0xef, 0x62, 0xd3, 0x1a, 0xd3, 0x4d, 0x32, 0x5d, 0x45, 0x4e, 0x04, 0xca, 0x65, 0xb8, 0x28,
		0x79, 0x3b, 0xc4, 0xc7, 0xc6, 0x31, 0xe9, 0x8c, 0x40, 0x69, 0x76, 0x73, 0x78, 0x41, 0x28, 0x6c,
		0x8b, 0x79, 0x89, 0x2d, 0x7e, 0x3b, 0x06, 0xcb, 0xf2, 0xa0, 0xd2, 0x09, 0x9c, 0xd5, 0x00, 0xc0,
		0xb6, 0xed, 0xf8, 0x61, 0x77, 0x4d, 0x87, 0xf2, 0x14, 0xae, 0x54, 0x09, 0x40, 0x5a, 0x88, 0x60,
		0xbd, 0x0f, 0x30, 0x9a, 0x39, 0xd5, 0x6d, 0x97, 0x21, 0x27, 0x6e, 0xee, 0xd9, 0xcf, 0x3f, 0xfc,
		0x68, 0x0b, 0x5c, 0x44, 0x4f, 0x34, 0x68, 0x15, 0x52, 0x47, 0xa4, 0x67, 0xda, 0xe2, 0x3e, 0x91,
		0x0f, 0xe4, 0x2d, 0x65, 0x32, 0xb8, 0xa5, 0xdc, 0xba, 0x03, 0x2b, 0x86, 0xd3, 0x9f, 0x34, 0x77,
		0x4b, 0x99, 0x38, 0x5e, 0x7b, 0x9f, 0x89, 0x7d, 0x0e, 0x46, 0x2d, 0xe6, 0x57, 0xe2, 0x89, 0xdd,
		0xd6, 0xd6, 0x57, 0xe3, 0xeb, 0xbb, 0x1c, 0xd7, 0x92, 0xaf, 0xa9, 0x91, 0xae, 0x45, 0x0c, 0x6a,
		0x3a, 0x7c, 0xff, 0x09, 0xf8, 0x44, 0xcf, 0xf4, 0x8f, 0x87, 0x47, 0x25, 0xc3, 0xe9, 0x5f, 0xed,
		0x39, 0x3d, 0x67, 0xf4, 0x73, 0x17, 0x1d, 0xb1, 0x01, 0xfb, 0x4f, 0xfc, 0xe4, 0x95, 0x0d, 0xa4,
		0xeb, 0x91, 0xbf, 0x8f, 0x95, 0xf7, 0x61, 0x45, 0x28, 0xeb, 0xec, 0xce, 0x9d, 0x1f, 0x0d, 0xd0,
		0x99, 0xf7, 0x2e, 0xea, 0x37, 0xde, 0x65, 0xb5, 0x5a, 0x5b, 0x16, 0x50, 0x3a, 0xc7, 0x0f, 0x10,
		0x65, 0x0d, 0x1e, 0x1e, 0xe3, 0xe3, 0xfb, 0x92, 0xb8, 0x11, 0x8c, 0xdf, 0x12, 0x8c, 0x2b, 0x21,
		0xc6, 0xb6, 0x80, 0x96, 0xab, 0xb0, 0x78, 0x1e, 0xae, 0xbf, 0x15, 0x5c, 0x79, 0x12, 0x26, 0xd9,
		0x85, 0x25, 0x46, 0x62, 0x0c, 0x3d, 0xdf, 0xe9, 0xb3, 0xa4, 0x77, 0x36, 0xcd, 0xdf, 0xbd, 0xcb,
		0x37, 0x4a, 0x81, 0xc2, 0xaa, 0x01, 0xaa, 0x5c, 0x06, 0xf6, 0x33, 0x43, 0x87, 0x18, 0x56, 0x04,
		0xc3, 0xdf, 0x0b, 0x43, 0x02, 0xfd, 0xf2, 0x4b, 0xb0, 0x4a, 0xff, 0x67, 0x39, 0x29, 0x6c, 0x49,
		0xf4, 0x2d, 0x93, 0xfa, 0xed, 0x37, 0xf8, 0x5e, 0x5c, 0x09, 0x08, 0x42, 0x36, 0x85, 0x56, 0xb1,
		0x47, 0x7c, 0x9f, 0xb8, 0x9e, 0x8e, 0xad, 0x59, 0xe6, 0x85, 0x8e, 0xe9, 0xea, 0x97, 0xde, 0x1b,
		0x5f, 0xc5, 0x5d, 0x8e, 0xac, 0x58, 0x56, 0xf9, 0x10, 0x2e, 0xcc, 0x88, 0x8a, 0x39, 0x38, 0xdf,
		0x14, 0x9c, 0xab, 0x53, 0x91, 0x41, 0x69, 0x5b, 0x20, 0xe5, 0xc1, 0x5a, 0xce, 0xc1, 0xf9, 0x9b,
		0x82, 0x13, 0x09, 0xac, 0x5c, 0x52, 0xca, 0x78, 0x1b, 0x96, 0xef, 0x13, 0xf7, 0xc8, 0xf1, 0xc4,
		0xd5, 0xc8, 0x1c, 0x74, 0x6f, 0x09, 0xba, 0x25, 0x01, 0x64, 0x77, 0x25, 0x94, 0xeb, 0x05, 0xc8,
		0x74, 0xb1, 0x41, 0xe6, 0xa0, 0xf8, 0xb2, 0xa0, 0x58, 0xa0, 0xfa, 0x14, 0x5a, 0x81, 0x7c, 0xcf,
		0x11, 0x65, 0x29, 0x1a, 0xfe, 0xb6, 0x80, 0xe7, 0x24, 0x46, 0x50, 0x0c, 0x9c, 0xc1, 0xd0, 0xa2,
		0x35, 0x2b, 0x9a, 0xe2, 0xb7, 0x24, 0x85, 0xc4, 0x08, 0x8a, 0x73, 0xb8, 0xf5, 0xb7, 0x25, 0x85,
		0x17, 0xf2, 0xe7, 0x8b, 0x90, 0x73, 0x6c, 0xeb, 0xc4, 0xb1, 0xe7, 0x31, 0xe2, 0x77, 0x04, 0x03,
		0x08, 0x08, 0x25, 0xb8, 0x05, 0xd9, 0x79, 0x17, 0xe2, 0x77, 0xdf, 0x93, 0xdb, 0x43, 0xae, 0xc0,
		0x2e, 0x2c, 0xc9, 0x04, 0x65, 0x3a, 0xf6, 0x1c, 0x14, 0xbf, 0x27, 0x28, 0x0a, 0x21, 0x98, 0x78,
		0x0d, 0x9f, 0x78, 0x7e, 0x8f, 0xcc, 0x43, 0xf2, 0x8e, 0x7c, 0x0d, 0x01, 0x11, 0xae, 0x3c, 0x22,
		0xb6, 0x71, 0x3c, 0x1f, 0xc3, 0xef, 0x4b, 0x57, 0x4a, 0x0c, 0xa5, 0xa8, 0xc2, 0x62, 0x1f, 0xbb,
		0xde, 0x31, 0xb6, 0xe6, 0x5a, 0x8e, 0x3f, 0x10, 0x1c, 0xf9, 0x00, 0x24, 0x3c, 0x32, 0xb4, 0xcf,
		0x43, 0xf3, 0x55, 0xe9, 0x91, 0x10, 0x4c, 0x6c, 0x3d, 0xcf, 0x67, 0x17, 0x50, 0xe7, 0x61, 0xfb,
		0x9a, 0xdc, 0x7a, 0x1c, 0xdb, 0x08, 0x33, 0xde, 0x82, 0xac, 0x67, 0xbe, 0x36, 0x17, 0xcd, 0x1f,
		0xca, 0x95, 0x66, 0x00, 0x0a, 0x7e, 0x05, 0x2e, 0xce, 0x2c, 0x13, 0x73, 0x90, 0xfd, 0x91, 0x20,
		0x5b, 0x9b, 0x51, 0x2a, 0x44, 0x4a, 0x38, 0x2f, 0xe5, 0x1f, 0xcb, 0x94, 0x40, 0x26, 0xb8, 0x5a,
		0xf4, 0xa0, 0xe0, 0xe1, 0xee, 0xf9, 0xbc, 0xf6, 0x27, 0xd2, 0x6b, 0x1c, 0x3b, 0xe6, 0xb5, 0x03,
		0x58, 0x13, 0x8c, 0xe7, 0x5b, 0xd7, 0xaf, 0xcb, 0xc4, 0xca, 0xd1, 0x87, 0xe3, 0xab, 0xfb, 0x93,
		0xb0, 0x1e, 0xb8, 0x53, 0x76, 0xa4, 0x9e, 0xde, 0xc7, 0x83, 0x39, 0x98, 0xbf, 0x21, 0x98, 0x65,
		0xc6, 0x0f, 0x5a, 0x5a, 0xaf, 0x81, 0x07, 0x94, 0xfc, 0x0e, 0xa8, 0x92, 0x7c, 0x68, 0xbb, 0xc4,
		0x70, 0x7a, 0xb6, 0xf9, 0x1a, 0xe9, 0xcc, 0x41, 0xfd, 0xa7, 0x13, 0x4b, 0x75, 0x18, 0x82, 0x53,
		0xe6, 0x3a, 0x28, 0x41, 0xaf, 0xa2, 0x9b, 0xfd, 0x81, 0xe3, 0xfa, 0x11, 0x8c, 0x7f, 0x26, 0x57,
		0x2a, 0xc0, 0xd5, 0x19, 0xac, 0x5c, 0x83, 0x02, 0x1b, 0xce, 0x1b, 0x92, 0x7f, 0x2e, 0x88, 0x16,
		0x47, 0x28, 0x91, 0x38, 0x0c, 0xa7, 0x3f, 0xc0, 0xee, 0x3c, 0xf9, 0xef, 0x2f, 0x64, 0xe2, 0x10,
		0x10, 0x91, 0x38, 0xfc, 0x93, 0x01, 0xa1, 0xd5, 0x7e, 0x0e, 0x86, 0x6f, 0xca, 0xc4, 0x21, 0x31,
		0x82, 0x42, 0x36, 0x0c, 0x73, 0x50, 0xfc, 0xa5, 0xa4, 0x90, 0x18, 0x4a, 0xf1, 0xd9, 0x51, 0xa1,
		0x75, 0x49, 0xcf, 0xf4, 0x7c, 0x97, 0xf7, 0xc1, 0x67, 0x53, 0xfd, 0xd5, 0x7b, 0xe3, 0x4d, 0x98,
		0x16, 0x82, 0x96, 0x6f, 0xc3, 0xd2, 0x44, 0x8b, 0x81, 0xa2, 0xbe, 0x59, 0x50, 0x7f, 0xea, 0x7d,
		0x91, 0x8c, 0xc6, 0x3b, 0x8c, 0xf2, 0x1e, 0x5d, 0xf7, 0xf1, 0x3e, 0x20, 0x9a, 0xec, 0x8d, 0xf7,
		0x83, 0xa5, 0x1f, 0x6b, 0x03, 0xca, 0x3b, 0xb0, 0x38, 0xd6, 0x03, 0x44, 0x53, 0xfd, 0xb4, 0xa0,
		0xca, 0x87, 0x5b, 0x80, 0xf2, 0x75, 0x48, 0xd2, 0x7a, 0x1e, 0x0d, 0xff, 0x19, 0x01, 0x67, 0xea,
		0xe5, 0x4f, 0x41, 0x46, 0xd6, 0xf1, 0x68, 0xe8, 0xcf, 0x0a, 0x68, 0x00, 0xa1, 0x70, 0x59, 0xc3,
		0xa3, 0xe1, 0x3f, 0x27, 0xe1, 0x12, 0x42, 0xe1, 0xf3, 0xbb, 0xf0, 0x6f, 0x7e, 0x3e, 0x29, 0xf2,
		0xb0, 0xf4, 0xdd, 0x2d, 0x58, 0x10, 0xc5, 0x3b, 0x1a, 0xfd, 0x05, 0xf1, 0x70, 0x89, 0x28, 0x3f,
		0x0f, 0xa9, 0x39, 0x1d, 0xfe, 0x0b, 0x02, 0xca, 0xf5, 0xcb, 0x55, 0xc8, 0x85, 0x0a, 0x76, 0x34,
		0xfc, 0x17, 0x05, 0x3c, 0x8c, 0xa2, 0xa6, 0x8b, 0x82, 0x1d, 0x4d, 0xf0, 0x4b, 0xd2, 0x74, 0x81,
		0xa0, 0x6e, 0x93, 0xb5, 0x3a, 0x1a, 0xfd, 0xcb, 0xd2, 0xeb, 0x12, 0x52, 0x7e, 0x11, 0xb2, 0x41,
		0xfe, 0x8d, 0xc6, 0xff, 0x8a, 0xc0, 0x8f, 0x30, 0xd4, 0x03, 0xa1, 0xfc, 0x1f, 0x4d, 0xf1, 0xab,
		0xd2, 0x03, 0x21, 0x14, 0xdd, 0x46, 0x93, 0x35, 0x3d, 0x9a, 0xe9, 0xd7, 0xe4, 0x36, 0x9a, 0x28,
		0xe9, 0x74, 0x35, 0x59, 0x1a, 0x8c, 0xa6, 0xf8, 0x75, 0xb9, 0x9a, 0x4c, 0x9f, 0x9a, 0x31, 0x59,
		0x24, 0xa3, 0x39, 0x7e, 0x43, 0x9a, 0x31, 0x51, 0x23, 0xcb, 0x2d, 0x40, 0xd3, 0x05, 0x32, 0x9a,
		0xef, 0x8b, 0x82, 0x6f, 0x79, 0xaa, 0x3e, 0x96, 0x5f, 0x86, 0xb5, 0xd9, 0xc5, 0x31, 0x9a, 0xf5,
		0x4b, 0xef, 0x4f, 0x1c, 0x67, 0xc2, 0xb5, 0xb1, 0x7c, 0x30, 0xca, 0xb2, 0xe1, 0xc2, 0x18, 0x4d,
		0xfb, 0xe6, 0xfb, 0xe3, 0x89, 0x36, 0x5c, 0x17, 0xcb, 0x15, 0x80, 0x51, 0x4d, 0x8a, 0xe6, 0x7a,
		0x4b, 0x70, 0x85, 0x40, 0x74, 0x6b, 0x88, 0x92, 0x14, 0x8d, 0xff, 0xb2, 0xdc, 0x1a, 0x02, 0x41,
		0xb7, 0x86, 0xac, 0x46, 0xd1, 0xe8, 0xb7, 0xe5, 0xd6, 0x90, 0x90, 0xf2, 0x2d, 0xc8, 0xd8, 0x43,
		0xcb, 0xa2, 0xb1, 0x85, 0xce, 0xfe, 0x8c, 0x48, 0xfd, 0xb7, 0x0f, 0x04, 0x58, 0x02, 0xca, 0xd7,
		0x21, 0x45, 0xfa, 0x47, 0xa4, 0x13, 0x85, 0xfc, 0xf7, 0x0f, 0x64, 0x3e, 0xa1, 0xda, 0xe5, 0x17,
		0x01, 0xf8, 0x61, 0x9a, 0xfd, 0x4a, 0x14, 0x81, 0xfd, 0x8f, 0x0f, 0xc4, 0x17, 0x0a, 0x23, 0xc8,
		0x88, 0x80, 0x7f, 0xef, 0x70, 0x36, 0xc1, 0x7b, 0xe3, 0x04, 0xec, 0x00, 0xfe, 0x02, 0x2c, 0xdc,
		0xf5, 0x1c, 0xdb, 0xc7, 0xbd, 0x28, 0xf4, 0x7f, 0x0a, 0xb4, 0xd4, 0xa7, 0x0e, 0xeb, 0x3b, 0x2e,
		0xf1, 0x71, 0xcf, 0x8b, 0xc2, 0xfe, 0x97, 0xc0, 0x06, 0x00, 0x0a, 0x36, 0xb0, 0xe7, 0xcf, 0xf3,
		0xde, 0xff, 0x2d, 0xc1, 0x12, 0x40, 0x8d, 0xa6, 0xff, 0xdf, 0x23, 0x27, 0x51, 0xd8, 0xef, 0x49,
		0xa3, 0x85, 0x7e, 0xf9, 0x53, 0x90, 0xa5, 0xff, 0xf2, 0xaf, 0x76, 0x22, 0xc0, 0xff, 0x23, 0xc0,
		0x23, 0x04, 0x7d, 0xb2, 0xe7, 0x77, 0x7c, 0x33, 0xda, 0xd9, 0xff, 0x2b, 0x56, 0x5a, 0xea, 0x97,
		0x2b, 0x90, 0xf3, 0xfc, 0x4e, 0x67, 0x28, 0x3a, 0x9a, 0x08, 0xf8, 0xf7, 0x3f, 0x08, 0x0e, 0xb9,
		0x01, 0x66, 0xeb, 0xca, 0xec, 0xcb, 0x3a, 0xd8, 0x75, 0x76, 0x1d, 0x7e, 0x4d, 0x07, 0xdf, 0x8d,
		0x43, 0xb6, 0xef, 0xf5, 0xc4, 0x8d, 0x1a, 0xdf, 0x5b, 0xb4, 0x76, 0x78, 0xeb, 0xe7, 0xbb, 0x8c,
		0x2b, 0xee, 0x40, 0x7a, 0xcb, 0xec, 0x35, 0xbc, 0x1e, 0x5a, 0x85, 0x14, 0x33, 0x8d, 0xfd, 0x8c,
		0x94, 0xd0, 0xf8, 0x00, 0x3d, 0x01, 0x89, 0x86, 0xd7, 0x13, 0x1f, 0xf0, 0xac, 0x96, 0x46, 0x0f,
		0x2a, 0xb5, 0xfb, 0xd8, 0xb2, 0x1a, 0x5e, 0x4f, 0xa3, 0x0a, 0x45, 0x17, 0x32, 0x52, 0x80, 0x36,
		0x20, 0xd7, 0x36, 0xb0, 0xbb, 0x35, 0xf4, 0xda, 0xbe, 0x33, 0x90, 0x1f, 0xa8, 0x84, 0x44, 0x68,
		0x13, 0x96, 0x76, 0x2c, 0xb3, 0x77, 0xec, 0xb7, 0xb0, 0x8b, 0x8d, 0xe3, 0xa1, 0x4f, 0xd4, 0xfc,
		0x46, 0x62, 0x73, 0x41, 0x9b, 0x14, 0xa3, 0x75, 0xc8, 0x34, 0xf0, 0xa0, 0x7d, 0x8c, 0xdd, 0x7b,
		0xec, 0x73, 0x87, 0xac, 0x16, 0x8c, 0x8b, 0x2f, 0x41, 0xba, 0xc5, 0x7f, 0xc1, 0x5d, 0x83, 0x64,
		0x9d, 0x5f, 0x07, 0x27, 0x36, 0x13, 0xfc, 0xfe, 0x94, 0x8e, 0xd1, 0x3a, 0xa4, 0x77, 0x2c, 0x07,
		0xfb, 0x1e, 0xfb, 0xba, 0x36, 0xc6, 0x66, 0x84, 0x04, 0xa9, 0x90, 0x3a, 0x34, 0xe5, 0x75, 0xf0,
		0x22, 0x9b, 0xe2, 0x82, 0xad, 0xfc, 0xf7, 0xbe, 0x73, 0x29, 0xf6, 0xff, 0xdf, 0xb9, 0x14, 0xfb,
		0xfa, 0x3f, 0x5f, 0x8a, 0xfd, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xf0, 0x76, 0x09, 0xc7, 0x2f,
		0x00, 0x00,
	}
	r := bytes.NewReader(gzipped)
	gzipr, err := compress_gzip.NewReader(r)
	if err != nil {
		panic(err)
	}
	ungzipped, err := io_ioutil.ReadAll(gzipr)
	if err != nil {
		panic(err)
	}
	if err := github_com_gogo_protobuf_proto.Unmarshal(ungzipped, d); err != nil {
		panic(err)
	}
	return d
}
func (this *BigMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&prototests.BigMsg{")
	if this.Field != nil {
		s = append(s, "Field: "+valueToGoStringMsg(this.Field, "int64")+",\n")
	}
	if this.Msg != nil {
		s = append(s, "Msg: "+fmt.Sprintf("%#v", this.Msg)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SmallMsg) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.SmallMsg{")
	if this.ScarBusStop != nil {
		s = append(s, "ScarBusStop: "+valueToGoStringMsg(this.ScarBusStop, "string")+",\n")
	}
	if this.FlightParachute != nil {
		s = append(s, "FlightParachute: "+fmt.Sprintf("%#v", this.FlightParachute)+",\n")
	}
	if this.MapShark != nil {
		s = append(s, "MapShark: "+valueToGoStringMsg(this.MapShark, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Packed) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&prototests.Packed{")
	if this.Ints != nil {
		s = append(s, "Ints: "+fmt.Sprintf("%#v", this.Ints)+",\n")
	}
	if this.Floats != nil {
		s = append(s, "Floats: "+fmt.Sprintf("%#v", this.Floats)+",\n")
	}
	if this.Uints != nil {
		s = append(s, "Uints: "+fmt.Sprintf("%#v", this.Uints)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMsg(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func NewPopulatedBigMsg(r randyMsg, easy bool) *BigMsg {
	this := &BigMsg{}
	if r.Intn(10) != 0 {
		v1 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v1 *= -1
		}
		this.Field = &v1
	}
	if r.Intn(10) != 0 {
		this.Msg = NewPopulatedSmallMsg(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 4)
	}
	return this
}

func NewPopulatedSmallMsg(r randyMsg, easy bool) *SmallMsg {
	this := &SmallMsg{}
	if r.Intn(10) != 0 {
		v2 := string(randStringMsg(r))
		this.ScarBusStop = &v2
	}
	if r.Intn(10) != 0 {
		v3 := r.Intn(10)
		this.FlightParachute = make([]uint32, v3)
		for i := 0; i < v3; i++ {
			this.FlightParachute[i] = uint32(r.Uint32())
		}
	}
	if r.Intn(10) != 0 {
		v4 := string(randStringMsg(r))
		this.MapShark = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 19)
	}
	return this
}

func NewPopulatedPacked(r randyMsg, easy bool) *Packed {
	this := &Packed{}
	if r.Intn(10) != 0 {
		v5 := r.Intn(10)
		this.Ints = make([]int64, v5)
		for i := 0; i < v5; i++ {
			this.Ints[i] = int64(r.Int63())
			if r.Intn(2) == 0 {
				this.Ints[i] *= -1
			}
		}
	}
	if r.Intn(10) != 0 {
		v6 := r.Intn(10)
		this.Floats = make([]float64, v6)
		for i := 0; i < v6; i++ {
			this.Floats[i] = float64(r.Float64())
			if r.Intn(2) == 0 {
				this.Floats[i] *= -1
			}
		}
	}
	if r.Intn(10) != 0 {
		v7 := r.Intn(10)
		this.Uints = make([]uint32, v7)
		for i := 0; i < v7; i++ {
			this.Uints[i] = uint32(r.Uint32())
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMsg(r, 7)
	}
	return this
}

type randyMsg interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMsg(r randyMsg) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMsg(r randyMsg) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneMsg(r)
	}
	return string(tmps)
}
func randUnrecognizedMsg(r randyMsg, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMsg(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMsg(dAtA []byte, r randyMsg, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(v9))
	case 1:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMsg(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMsg(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("msg.proto", fileDescriptorMsg) }

var fileDescriptorMsg = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8d, 0x4f, 0x6b, 0x83, 0x30,
	0x18, 0xc6, 0xc9, 0x52, 0x5d, 0xfb, 0xb6, 0x63, 0x23, 0x94, 0x11, 0x3c, 0x48, 0xf0, 0x30, 0x72,
	0x99, 0x85, 0x7d, 0x04, 0x0f, 0xc2, 0x0e, 0x42, 0x51, 0xb6, 0x7b, 0x6a, 0x5d, 0x94, 0xea, 0x22,
	0x26, 0x7e, 0xaf, 0x7d, 0xa5, 0xee, 0x13, 0xec, 0xb8, 0xe3, 0xe8, 0xeb, 0xfe, 0xd1, 0x53, 0xf2,
	0xfb, 0x3d, 0x79, 0x9e, 0xc0, 0xa2, 0xb3, 0x3a, 0xee, 0x07, 0xe3, 0x0c, 0x03, 0x3c, 0x5c, 0x65,
	0x9d, 0x0d, 0xee, 0x75, 0xe3, 0xea, 0x71, 0x17, 0x97, 0xa6, 0xdb, 0x68, 0xa3, 0xcd, 0x06, 0xb3,
	0xdd, 0xf8, 0x82, 0x84, 0x80, 0xb7, 0xa9, 0x1a, 0xa5, 0xe0, 0x27, 0x8d, 0xce, 0xac, 0x66, 0x6b,
	0xf0, 0xd2, 0xa6, 0x6a, 0xf7, 0x9c, 0x08, 0x22, 0x69, 0x3e, 0x01, 0xbb, 0x03, 0x9a, 0x59, 0xcd,
	0xa9, 0x20, 0x72, 0xf9, 0xb0, 0x8e, 0xff, 0x3e, 0x8a, 0x8b, 0x4e, 0xb5, 0x6d, 0x66, 0x75, 0x7e,
	0x7a, 0x10, 0x0d, 0x30, 0xff, 0x11, 0x4c, 0xc0, 0xb2, 0x28, 0xd5, 0x90, 0x8c, 0xb6, 0x70, 0xa6,
	0xc7, 0xbd, 0x45, 0xfe, 0x5f, 0x31, 0x09, 0xd7, 0x69, 0xdb, 0xe8, 0xda, 0x6d, 0xd5, 0xa0, 0xca,
	0x7a, 0x74, 0x15, 0x5f, 0x09, 0x2a, 0x2f, 0xf3, 0x73, 0xcd, 0x02, 0x98, 0x67, 0xaa, 0x2f, 0x6a,
	0x35, 0x1c, 0x38, 0xc3, 0xa1, 0x5f, 0x8e, 0x9e, 0xc1, 0xdf, 0xaa, 0xf2, 0x50, 0xed, 0xd9, 0x2d,
	0xcc, 0x1e, 0x5f, 0x9d, 0xe5, 0x33, 0x41, 0x25, 0x4d, 0x2e, 0x6e, 0x48, 0x8e, 0xcc, 0x02, 0xf0,
	0xd3, 0xd6, 0x28, 0x67, 0xb9, 0x27, 0xa8, 0x24, 0x98, 0x7c, 0x1b, 0xc6, 0xc1, 0x7b, 0x6a, 0x4e,
	0x25, 0x5f, 0x50, 0x79, 0x85, 0xd1, 0x24, 0x92, 0xd5, 0xc7, 0x31, 0x24, 0x9f, 0xc7, 0x90, 0xbc,
	0xbd, 0x87, 0xe4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xcd, 0x1a, 0x2c, 0x68, 0x01, 0x00, 0x00,
}
