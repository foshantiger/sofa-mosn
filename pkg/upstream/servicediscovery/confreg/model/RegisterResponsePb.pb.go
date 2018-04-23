// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RegisterResponsePb.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegisterResponsePb struct {
	Success  bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	RegistId string `protobuf:"bytes,2,opt,name=registId" json:"registId,omitempty"`
	Version  int64  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Refused  bool   `protobuf:"varint,4,opt,name=refused" json:"refused,omitempty"`
	Message  string `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
}

func (m *RegisterResponsePb) Reset()                    { *m = RegisterResponsePb{} }
func (m *RegisterResponsePb) String() string            { return proto.CompactTextString(m) }
func (*RegisterResponsePb) ProtoMessage()               {}
func (*RegisterResponsePb) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *RegisterResponsePb) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RegisterResponsePb) GetRegistId() string {
	if m != nil {
		return m.RegistId
	}
	return ""
}

func (m *RegisterResponsePb) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RegisterResponsePb) GetRefused() bool {
	if m != nil {
		return m.Refused
	}
	return false
}

func (m *RegisterResponsePb) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterResponsePb)(nil), "RegisterResponsePb")
}

func init() { proto.RegisterFile("RegisterResponsePb.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xcf, 0xb1, 0xca, 0x83, 0x30,
	0x10, 0x07, 0x70, 0xa2, 0xdf, 0xd7, 0xda, 0x8c, 0x4e, 0xa1, 0x93, 0x74, 0xca, 0x14, 0x0a, 0x7d,
	0x83, 0x6e, 0xdd, 0x24, 0x63, 0xb7, 0x44, 0x4f, 0x11, 0xd4, 0x0b, 0x77, 0x2a, 0xf8, 0x26, 0x7d,
	0xdc, 0xa2, 0xd6, 0x2e, 0x1d, 0x7f, 0xfc, 0xef, 0xfe, 0xc7, 0x49, 0x65, 0xa1, 0x6e, 0x78, 0x00,
	0xb2, 0xc0, 0x01, 0x7b, 0x86, 0xdc, 0x9b, 0x40, 0x38, 0xe0, 0xe5, 0x25, 0x64, 0xfa, 0x1b, 0xa6,
	0x4a, 0x1e, 0x79, 0x2c, 0x0a, 0x60, 0x56, 0x22, 0x13, 0x3a, 0xb1, 0x3b, 0xd3, 0xb3, 0x4c, 0x68,
	0x9d, 0x7f, 0x94, 0x2a, 0xca, 0x84, 0x3e, 0xd9, 0xaf, 0x97, 0xad, 0x09, 0x88, 0x1b, 0xec, 0x55,
	0x9c, 0x09, 0x1d, 0xdb, 0x9d, 0x4b, 0x42, 0x50, 0x8d, 0x0c, 0xa5, 0xfa, 0xdb, 0xfa, 0x3e, 0x5c,
	0x92, 0x0e, 0x98, 0x5d, 0x0d, 0xea, 0x7f, 0xad, 0xdb, 0x79, 0xbf, 0x4a, 0x5d, 0x60, 0x67, 0x5c,
	0xdb, 0x04, 0x37, 0x1b, 0xc6, 0xca, 0x99, 0xed, 0x12, 0xcd, 0x86, 0x81, 0x26, 0x20, 0xd3, 0x61,
	0x09, 0xad, 0x09, 0x3e, 0x17, 0xcf, 0x28, 0x78, 0x7f, 0x58, 0x7f, 0xba, 0xbd, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xf2, 0x8f, 0xbc, 0xee, 0xef, 0x00, 0x00, 0x00,
}
