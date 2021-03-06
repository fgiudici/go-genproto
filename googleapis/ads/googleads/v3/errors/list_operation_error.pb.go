// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/list_operation_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible list operation errors.
type ListOperationErrorEnum_ListOperationError int32

const (
	// Enum unspecified.
	ListOperationErrorEnum_UNSPECIFIED ListOperationErrorEnum_ListOperationError = 0
	// The received error code is not known in this version.
	ListOperationErrorEnum_UNKNOWN ListOperationErrorEnum_ListOperationError = 1
	// Field required in value is missing.
	ListOperationErrorEnum_REQUIRED_FIELD_MISSING ListOperationErrorEnum_ListOperationError = 7
	// Duplicate or identical value is sent in multiple list operations.
	ListOperationErrorEnum_DUPLICATE_VALUES ListOperationErrorEnum_ListOperationError = 8
)

var ListOperationErrorEnum_ListOperationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	7: "REQUIRED_FIELD_MISSING",
	8: "DUPLICATE_VALUES",
}

var ListOperationErrorEnum_ListOperationError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"REQUIRED_FIELD_MISSING": 7,
	"DUPLICATE_VALUES":       8,
}

func (x ListOperationErrorEnum_ListOperationError) String() string {
	return proto.EnumName(ListOperationErrorEnum_ListOperationError_name, int32(x))
}

func (ListOperationErrorEnum_ListOperationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d4bd5dcfdab7508, []int{0, 0}
}

// Container for enum describing possible list operation errors.
type ListOperationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOperationErrorEnum) Reset()         { *m = ListOperationErrorEnum{} }
func (m *ListOperationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ListOperationErrorEnum) ProtoMessage()    {}
func (*ListOperationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d4bd5dcfdab7508, []int{0}
}

func (m *ListOperationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOperationErrorEnum.Unmarshal(m, b)
}
func (m *ListOperationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOperationErrorEnum.Marshal(b, m, deterministic)
}
func (m *ListOperationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOperationErrorEnum.Merge(m, src)
}
func (m *ListOperationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ListOperationErrorEnum.Size(m)
}
func (m *ListOperationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOperationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ListOperationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.ListOperationErrorEnum_ListOperationError", ListOperationErrorEnum_ListOperationError_name, ListOperationErrorEnum_ListOperationError_value)
	proto.RegisterType((*ListOperationErrorEnum)(nil), "google.ads.googleads.v3.errors.ListOperationErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/list_operation_error.proto", fileDescriptor_4d4bd5dcfdab7508)
}

var fileDescriptor_4d4bd5dcfdab7508 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xf6, 0x07, 0x2b, 0xd3, 0x85, 0x21, 0x48, 0x85, 0x22, 0x5d, 0xe4, 0x01, 0x26,
	0x8b, 0xac, 0x1c, 0x57, 0xd3, 0x66, 0x5a, 0x06, 0x63, 0x1a, 0x1b, 0x13, 0x41, 0x02, 0x21, 0x3a,
	0x61, 0x08, 0xb4, 0x33, 0x61, 0x26, 0x76, 0xe9, 0xc3, 0xb8, 0xf4, 0x51, 0x7c, 0x14, 0x97, 0x3e,
	0x81, 0x24, 0x63, 0xba, 0x29, 0xba, 0xca, 0xe1, 0xe6, 0x3b, 0x67, 0xce, 0xbd, 0xe0, 0x8a, 0x4b,
	0xc9, 0xb7, 0xa5, 0x5b, 0x30, 0xed, 0x1a, 0xd9, 0xaa, 0xbd, 0xe7, 0x96, 0x4a, 0x49, 0xa5, 0xdd,
	0x6d, 0xa5, 0x9b, 0x5c, 0xd6, 0xa5, 0x2a, 0x9a, 0x4a, 0x8a, 0xbc, 0x9b, 0xc2, 0x5a, 0xc9, 0x46,
	0xda, 0x33, 0xc3, 0xc3, 0x82, 0x69, 0x78, 0xb0, 0xc2, 0xbd, 0x07, 0x8d, 0x75, 0x7a, 0xd9, 0x47,
	0xd7, 0x95, 0x5b, 0x08, 0x21, 0x9b, 0x2e, 0x42, 0x1b, 0xb7, 0xf3, 0x0a, 0x26, 0x41, 0xa5, 0x9b,
	0x75, 0x1f, 0x4d, 0x5a, 0x13, 0x11, 0x2f, 0x3b, 0x87, 0x01, 0xfb, 0xf8, 0x8f, 0x7d, 0x06, 0xc6,
	0x49, 0x18, 0x47, 0x64, 0x41, 0x97, 0x94, 0xf8, 0xd6, 0x3f, 0x7b, 0x0c, 0x46, 0x49, 0x78, 0x13,
	0xae, 0x1f, 0x42, 0x6b, 0x60, 0x4f, 0xc1, 0x64, 0x43, 0xee, 0x12, 0xba, 0x21, 0x7e, 0xbe, 0xa4,
	0x24, 0xf0, 0xf3, 0x5b, 0x1a, 0xc7, 0x34, 0x5c, 0x59, 0x23, 0xfb, 0x1c, 0x58, 0x7e, 0x12, 0x05,
	0x74, 0x81, 0xef, 0x49, 0x9e, 0xe2, 0x20, 0x21, 0xb1, 0x75, 0x3a, 0xff, 0x1a, 0x00, 0xe7, 0x59,
	0xee, 0xe0, 0xdf, 0x4b, 0xcc, 0x2f, 0x8e, 0xab, 0x44, 0x6d, 0xff, 0x68, 0xf0, 0xe8, 0xff, 0x58,
	0xb9, 0xdc, 0x16, 0x82, 0x43, 0xa9, 0xb8, 0xcb, 0x4b, 0xd1, 0x6d, 0xd7, 0x9f, 0xb2, 0xae, 0xf4,
	0x6f, 0x97, 0xbd, 0x36, 0x9f, 0xb7, 0xe1, 0xff, 0x15, 0xc6, 0xef, 0xc3, 0xd9, 0xca, 0x84, 0x61,
	0xa6, 0xa1, 0x91, 0xad, 0x4a, 0x3d, 0xd8, 0x3d, 0xa9, 0x3f, 0x7a, 0x20, 0xc3, 0x4c, 0x67, 0x07,
	0x20, 0x4b, 0xbd, 0xcc, 0x00, 0x9f, 0x43, 0xc7, 0x4c, 0x11, 0xc2, 0x4c, 0x23, 0x74, 0x40, 0x10,
	0x4a, 0x3d, 0x84, 0x0c, 0xf4, 0x74, 0xd2, 0xb5, 0xf3, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x04, 0xee, 0x90, 0xf6, 0x01, 0x00, 0x00,
}
