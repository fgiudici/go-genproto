// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/custom_interest_member_type.proto

package enums

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

// Enum containing possible custom interest member types.
type CustomInterestMemberTypeEnum_CustomInterestMemberType int32

const (
	// Not specified.
	CustomInterestMemberTypeEnum_UNSPECIFIED CustomInterestMemberTypeEnum_CustomInterestMemberType = 0
	// Used for return value only. Represents value unknown in this version.
	CustomInterestMemberTypeEnum_UNKNOWN CustomInterestMemberTypeEnum_CustomInterestMemberType = 1
	// Custom interest member type KEYWORD.
	CustomInterestMemberTypeEnum_KEYWORD CustomInterestMemberTypeEnum_CustomInterestMemberType = 2
	// Custom interest member type URL.
	CustomInterestMemberTypeEnum_URL CustomInterestMemberTypeEnum_CustomInterestMemberType = 3
)

var CustomInterestMemberTypeEnum_CustomInterestMemberType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "KEYWORD",
	3: "URL",
}

var CustomInterestMemberTypeEnum_CustomInterestMemberType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"KEYWORD":     2,
	"URL":         3,
}

func (x CustomInterestMemberTypeEnum_CustomInterestMemberType) String() string {
	return proto.EnumName(CustomInterestMemberTypeEnum_CustomInterestMemberType_name, int32(x))
}

func (CustomInterestMemberTypeEnum_CustomInterestMemberType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e2f9e6f3d781c9a, []int{0, 0}
}

// The types of custom interest member, either KEYWORD or URL.
type CustomInterestMemberTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomInterestMemberTypeEnum) Reset()         { *m = CustomInterestMemberTypeEnum{} }
func (m *CustomInterestMemberTypeEnum) String() string { return proto.CompactTextString(m) }
func (*CustomInterestMemberTypeEnum) ProtoMessage()    {}
func (*CustomInterestMemberTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e2f9e6f3d781c9a, []int{0}
}

func (m *CustomInterestMemberTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterestMemberTypeEnum.Unmarshal(m, b)
}
func (m *CustomInterestMemberTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterestMemberTypeEnum.Marshal(b, m, deterministic)
}
func (m *CustomInterestMemberTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterestMemberTypeEnum.Merge(m, src)
}
func (m *CustomInterestMemberTypeEnum) XXX_Size() int {
	return xxx_messageInfo_CustomInterestMemberTypeEnum.Size(m)
}
func (m *CustomInterestMemberTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterestMemberTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterestMemberTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.CustomInterestMemberTypeEnum_CustomInterestMemberType", CustomInterestMemberTypeEnum_CustomInterestMemberType_name, CustomInterestMemberTypeEnum_CustomInterestMemberType_value)
	proto.RegisterType((*CustomInterestMemberTypeEnum)(nil), "google.ads.googleads.v3.enums.CustomInterestMemberTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/custom_interest_member_type.proto", fileDescriptor_1e2f9e6f3d781c9a)
}

var fileDescriptor_1e2f9e6f3d781c9a = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x6b, 0xfa, 0x30,
	0x18, 0xfe, 0x59, 0xe1, 0x27, 0xc4, 0xc3, 0x4a, 0x4f, 0x63, 0xe8, 0x41, 0x3f, 0x40, 0x72, 0xc8,
	0x2d, 0x3b, 0x8c, 0xaa, 0x9d, 0x88, 0x5b, 0x15, 0x37, 0x95, 0x8d, 0x42, 0xa9, 0x36, 0x84, 0x82,
	0x49, 0x4a, 0x93, 0x0a, 0x7e, 0x9d, 0x1d, 0xf7, 0x51, 0xf6, 0x51, 0x76, 0xdc, 0x27, 0x18, 0x4d,
	0x6c, 0x6f, 0xdd, 0xa5, 0x3c, 0xcd, 0xf3, 0x3e, 0x7f, 0xde, 0x17, 0x3c, 0x30, 0x29, 0xd9, 0x89,
	0xa2, 0x24, 0x55, 0xc8, 0xc2, 0x0a, 0x9d, 0x31, 0xa2, 0xa2, 0xe4, 0x0a, 0x1d, 0x4b, 0xa5, 0x25,
	0x8f, 0x33, 0xa1, 0x69, 0x41, 0x95, 0x8e, 0x39, 0xe5, 0x07, 0x5a, 0xc4, 0xfa, 0x92, 0x53, 0x98,
	0x17, 0x52, 0x4b, 0x6f, 0x68, 0x55, 0x30, 0x49, 0x15, 0x6c, 0x0c, 0xe0, 0x19, 0x43, 0x63, 0x70,
	0x37, 0xa8, 0xfd, 0xf3, 0x0c, 0x25, 0x42, 0x48, 0x9d, 0xe8, 0x4c, 0x0a, 0x65, 0xc5, 0x63, 0x01,
	0x06, 0x53, 0x93, 0xb0, 0xb8, 0x06, 0x3c, 0x1b, 0xff, 0xd7, 0x4b, 0x4e, 0x03, 0x51, 0xf2, 0x71,
	0x08, 0x6e, 0xdb, 0x78, 0xef, 0x06, 0xf4, 0xb7, 0xe1, 0xcb, 0x3a, 0x98, 0x2e, 0x1e, 0x17, 0xc1,
	0xcc, 0xfd, 0xe7, 0xf5, 0x41, 0x6f, 0x1b, 0x2e, 0xc3, 0xd5, 0x3e, 0x74, 0x3b, 0xd5, 0xcf, 0x32,
	0x78, 0xdb, 0xaf, 0x36, 0x33, 0xd7, 0xf1, 0x7a, 0xa0, 0xbb, 0xdd, 0x3c, 0xb9, 0xdd, 0xc9, 0x4f,
	0x07, 0x8c, 0x8e, 0x92, 0xc3, 0x3f, 0x3b, 0x4f, 0x86, 0x6d, 0x99, 0xeb, 0xaa, 0xf4, 0xba, 0xf3,
	0x3e, 0xb9, 0xea, 0x99, 0x3c, 0x25, 0x82, 0x41, 0x59, 0x30, 0xc4, 0xa8, 0x30, 0x2b, 0xd5, 0x47,
	0xcc, 0x33, 0xd5, 0x72, 0xd3, 0x7b, 0xf3, 0xfd, 0x70, 0xba, 0x73, 0xdf, 0xff, 0x74, 0x86, 0x73,
	0x6b, 0xe5, 0xa7, 0x0a, 0x5a, 0x58, 0xa1, 0x1d, 0x86, 0xd5, 0xfe, 0xea, 0xab, 0xe6, 0x23, 0x3f,
	0x55, 0x51, 0xc3, 0x47, 0x3b, 0x1c, 0x19, 0xfe, 0xdb, 0x19, 0xd9, 0x47, 0x42, 0xfc, 0x54, 0x11,
	0xd2, 0x4c, 0x10, 0xb2, 0xc3, 0x84, 0x98, 0x99, 0xc3, 0x7f, 0x53, 0x0c, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x87, 0x00, 0x9b, 0xbb, 0xeb, 0x01, 0x00, 0x00,
}
