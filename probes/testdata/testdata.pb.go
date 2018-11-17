// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/testdata/testdata.proto

package testdata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/google/cloudprober/probes/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FancyProbe struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FancyProbe) Reset()         { *m = FancyProbe{} }
func (m *FancyProbe) String() string { return proto.CompactTextString(m) }
func (*FancyProbe) ProtoMessage()    {}
func (*FancyProbe) Descriptor() ([]byte, []int) {
	return fileDescriptor_testdata_cc9d738d6f49fdf4, []int{0}
}
func (m *FancyProbe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FancyProbe.Unmarshal(m, b)
}
func (m *FancyProbe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FancyProbe.Marshal(b, m, deterministic)
}
func (dst *FancyProbe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FancyProbe.Merge(dst, src)
}
func (m *FancyProbe) XXX_Size() int {
	return xxx_messageInfo_FancyProbe.Size(m)
}
func (m *FancyProbe) XXX_DiscardUnknown() {
	xxx_messageInfo_FancyProbe.DiscardUnknown(m)
}

var xxx_messageInfo_FancyProbe proto.InternalMessageInfo

func (m *FancyProbe) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

var E_FancyProbe = &proto.ExtensionDesc{
	ExtendedType:  (*proto1.ProbeDef)(nil),
	ExtensionType: (*FancyProbe)(nil),
	Field:         200,
	Name:          "cloudprober.probes.testdata.fancy_probe",
	Tag:           "bytes,200,opt,name=fancy_probe,json=fancyProbe",
	Filename:      "github.com/google/cloudprober/probes/testdata/testdata.proto",
}

func init() {
	proto.RegisterType((*FancyProbe)(nil), "cloudprober.probes.testdata.FancyProbe")
	proto.RegisterExtension(E_FancyProbe)
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/testdata/testdata.proto", fileDescriptor_testdata_cc9d738d6f49fdf4)
}

var fileDescriptor_testdata_cc9d738d6f49fdf4 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0x29, 0x28, 0xca, 0x4f, 0x4a, 0x2d, 0xd2, 0x07, 0x53, 0xc5, 0xfa, 0x25, 0xa9,
	0xc5, 0x25, 0x29, 0x89, 0x25, 0x89, 0x70, 0x86, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x34,
	0x92, 0x5a, 0x3d, 0x88, 0x5a, 0x3d, 0x98, 0x12, 0x29, 0x73, 0xa2, 0x8c, 0x06, 0x1b, 0xa4, 0x9f,
	0x9c, 0x9f, 0x97, 0x96, 0x99, 0x0e, 0x31, 0x55, 0x49, 0x81, 0x8b, 0xcb, 0x2d, 0x31, 0x2f, 0xb9,
	0x32, 0x00, 0xa4, 0x42, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x49,
	0x83, 0x33, 0x08, 0xcc, 0xb6, 0x4a, 0xe7, 0xe2, 0x4e, 0x03, 0xa9, 0x88, 0x07, 0x1b, 0x22, 0x24,
	0xa3, 0x87, 0xc5, 0x1d, 0x60, 0xdd, 0x2e, 0xa9, 0x69, 0x12, 0x27, 0x18, 0x15, 0x18, 0x35, 0xb8,
	0x8d, 0xd4, 0xf5, 0xf0, 0x38, 0x56, 0x0f, 0x61, 0x61, 0x10, 0x57, 0x1a, 0x9c, 0x0d, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x5b, 0x33, 0x00, 0xb5, 0x1f, 0x01, 0x00, 0x00,
}