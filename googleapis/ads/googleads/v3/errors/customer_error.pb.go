// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/customer_error.proto

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

// Set of errors that are related to requests dealing with Customer.
type CustomerErrorEnum_CustomerError int32

const (
	// Enum unspecified.
	CustomerErrorEnum_UNSPECIFIED CustomerErrorEnum_CustomerError = 0
	// The received error code is not known in this version.
	CustomerErrorEnum_UNKNOWN CustomerErrorEnum_CustomerError = 1
	// Customer status is not allowed to be changed from DRAFT and CLOSED.
	// Currency code and at least one of country code and time zone needs to be
	// set when status is changed to ENABLED.
	CustomerErrorEnum_STATUS_CHANGE_DISALLOWED CustomerErrorEnum_CustomerError = 2
	// CustomerService cannot get a customer that has not been fully set up.
	CustomerErrorEnum_ACCOUNT_NOT_SET_UP CustomerErrorEnum_CustomerError = 3
)

var CustomerErrorEnum_CustomerError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "STATUS_CHANGE_DISALLOWED",
	3: "ACCOUNT_NOT_SET_UP",
}

var CustomerErrorEnum_CustomerError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"STATUS_CHANGE_DISALLOWED": 2,
	"ACCOUNT_NOT_SET_UP":       3,
}

func (x CustomerErrorEnum_CustomerError) String() string {
	return proto.EnumName(CustomerErrorEnum_CustomerError_name, int32(x))
}

func (CustomerErrorEnum_CustomerError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_355f6e19d610cfcf, []int{0, 0}
}

// Container for enum describing possible customer errors.
type CustomerErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerErrorEnum) Reset()         { *m = CustomerErrorEnum{} }
func (m *CustomerErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerErrorEnum) ProtoMessage()    {}
func (*CustomerErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_355f6e19d610cfcf, []int{0}
}

func (m *CustomerErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerErrorEnum.Unmarshal(m, b)
}
func (m *CustomerErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerErrorEnum.Marshal(b, m, deterministic)
}
func (m *CustomerErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerErrorEnum.Merge(m, src)
}
func (m *CustomerErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerErrorEnum.Size(m)
}
func (m *CustomerErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.CustomerErrorEnum_CustomerError", CustomerErrorEnum_CustomerError_name, CustomerErrorEnum_CustomerError_value)
	proto.RegisterType((*CustomerErrorEnum)(nil), "google.ads.googleads.v3.errors.CustomerErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/customer_error.proto", fileDescriptor_355f6e19d610cfcf)
}

var fileDescriptor_355f6e19d610cfcf = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x5d, 0x07, 0x0a, 0x19, 0x62, 0xcd, 0x41, 0x44, 0xc6, 0x0e, 0x7d, 0x80, 0xf4, 0x90,
	0x5b, 0x3c, 0x65, 0x6d, 0x9c, 0xc3, 0x91, 0x16, 0xda, 0x6e, 0x20, 0x85, 0x52, 0xd7, 0x12, 0x06,
	0x5b, 0x32, 0x92, 0x6e, 0xf8, 0x3c, 0x1e, 0x7d, 0x14, 0x1f, 0x45, 0xf0, 0x1d, 0xa4, 0xcd, 0x36,
	0xd8, 0x41, 0x4f, 0xf9, 0xf2, 0xe7, 0xf7, 0x7d, 0xf9, 0xf2, 0x07, 0x58, 0x28, 0x25, 0xd6, 0xb5,
	0x5f, 0x56, 0xc6, 0xb7, 0xb2, 0x55, 0x7b, 0xec, 0xd7, 0x5a, 0x2b, 0x6d, 0xfc, 0xe5, 0xce, 0x34,
	0x6a, 0x53, 0xeb, 0xa2, 0xbb, 0xa3, 0xad, 0x56, 0x8d, 0x82, 0x23, 0x4b, 0xa2, 0xb2, 0x32, 0xe8,
	0x64, 0x42, 0x7b, 0x8c, 0xac, 0xe9, 0x61, 0x78, 0x0c, 0xdd, 0xae, 0xfc, 0x52, 0x4a, 0xd5, 0x94,
	0xcd, 0x4a, 0x49, 0x63, 0xdd, 0xde, 0x3b, 0xb8, 0x0d, 0x0e, 0xa9, 0xac, 0xe5, 0x99, 0xdc, 0x6d,
	0xbc, 0x25, 0xb8, 0x3e, 0x1b, 0xc2, 0x1b, 0x30, 0xc8, 0x78, 0x12, 0xb3, 0x60, 0xfa, 0x34, 0x65,
	0xa1, 0x7b, 0x01, 0x07, 0xe0, 0x2a, 0xe3, 0x2f, 0x3c, 0x5a, 0x70, 0xb7, 0x07, 0x87, 0xe0, 0x3e,
	0x49, 0x69, 0x9a, 0x25, 0x45, 0xf0, 0x4c, 0xf9, 0x84, 0x15, 0xe1, 0x34, 0xa1, 0xb3, 0x59, 0xb4,
	0x60, 0xa1, 0xeb, 0xc0, 0x3b, 0x00, 0x69, 0x10, 0x44, 0x19, 0x4f, 0x0b, 0x1e, 0xa5, 0x45, 0xc2,
	0xd2, 0x22, 0x8b, 0xdd, 0xfe, 0xf8, 0xa7, 0x07, 0xbc, 0xa5, 0xda, 0xa0, 0xff, 0xeb, 0x8f, 0xe1,
	0x59, 0x93, 0xb8, 0x2d, 0x1d, 0xf7, 0x5e, 0xc3, 0x83, 0x4b, 0xa8, 0x75, 0x29, 0x05, 0x52, 0x5a,
	0xf8, 0xa2, 0x96, 0xdd, 0x97, 0x8e, 0x9b, 0xdb, 0xae, 0xcc, 0x5f, 0x8b, 0x7c, 0xb4, 0xc7, 0x87,
	0xd3, 0x9f, 0x50, 0xfa, 0xe9, 0x8c, 0x26, 0x36, 0x8c, 0x56, 0x06, 0x59, 0xd9, 0xaa, 0x39, 0x46,
	0xdd, 0x93, 0xe6, 0xeb, 0x08, 0xe4, 0xb4, 0x32, 0xf9, 0x09, 0xc8, 0xe7, 0x38, 0xb7, 0xc0, 0xb7,
	0xe3, 0xd9, 0x29, 0x21, 0xb4, 0x32, 0x84, 0x9c, 0x10, 0x42, 0xe6, 0x98, 0x10, 0x0b, 0xbd, 0x5d,
	0x76, 0xed, 0xf0, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0xca, 0x0f, 0x94, 0xe5, 0x01, 0x00,
	0x00,
}
