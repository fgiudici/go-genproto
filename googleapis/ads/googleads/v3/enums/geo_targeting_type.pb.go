// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/geo_targeting_type.proto

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

// The possible geo targeting types.
type GeoTargetingTypeEnum_GeoTargetingType int32

const (
	// Not specified.
	GeoTargetingTypeEnum_UNSPECIFIED GeoTargetingTypeEnum_GeoTargetingType = 0
	// The value is unknown in this version.
	GeoTargetingTypeEnum_UNKNOWN GeoTargetingTypeEnum_GeoTargetingType = 1
	// Location the user is interested in while making the query.
	GeoTargetingTypeEnum_AREA_OF_INTEREST GeoTargetingTypeEnum_GeoTargetingType = 2
	// Location of the user issuing the query.
	GeoTargetingTypeEnum_LOCATION_OF_PRESENCE GeoTargetingTypeEnum_GeoTargetingType = 3
)

var GeoTargetingTypeEnum_GeoTargetingType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AREA_OF_INTEREST",
	3: "LOCATION_OF_PRESENCE",
}

var GeoTargetingTypeEnum_GeoTargetingType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"AREA_OF_INTEREST":     2,
	"LOCATION_OF_PRESENCE": 3,
}

func (x GeoTargetingTypeEnum_GeoTargetingType) String() string {
	return proto.EnumName(GeoTargetingTypeEnum_GeoTargetingType_name, int32(x))
}

func (GeoTargetingTypeEnum_GeoTargetingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89674afbd6140ba4, []int{0, 0}
}

// Container for enum describing possible geo targeting types.
type GeoTargetingTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoTargetingTypeEnum) Reset()         { *m = GeoTargetingTypeEnum{} }
func (m *GeoTargetingTypeEnum) String() string { return proto.CompactTextString(m) }
func (*GeoTargetingTypeEnum) ProtoMessage()    {}
func (*GeoTargetingTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_89674afbd6140ba4, []int{0}
}

func (m *GeoTargetingTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoTargetingTypeEnum.Unmarshal(m, b)
}
func (m *GeoTargetingTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoTargetingTypeEnum.Marshal(b, m, deterministic)
}
func (m *GeoTargetingTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoTargetingTypeEnum.Merge(m, src)
}
func (m *GeoTargetingTypeEnum) XXX_Size() int {
	return xxx_messageInfo_GeoTargetingTypeEnum.Size(m)
}
func (m *GeoTargetingTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoTargetingTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GeoTargetingTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.GeoTargetingTypeEnum_GeoTargetingType", GeoTargetingTypeEnum_GeoTargetingType_name, GeoTargetingTypeEnum_GeoTargetingType_value)
	proto.RegisterType((*GeoTargetingTypeEnum)(nil), "google.ads.googleads.v3.enums.GeoTargetingTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/geo_targeting_type.proto", fileDescriptor_89674afbd6140ba4)
}

var fileDescriptor_89674afbd6140ba4 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0x76, 0x1d, 0x28, 0x64, 0x07, 0x4b, 0x99, 0x20, 0xe2, 0x0e, 0xdb, 0x03, 0xa4, 0x87, 0x82,
	0x87, 0x78, 0xca, 0x66, 0x36, 0x8a, 0x92, 0x96, 0xad, 0x9b, 0x20, 0x85, 0x5a, 0x6d, 0x08, 0x85,
	0x2d, 0x29, 0x4b, 0x36, 0xdc, 0xeb, 0x78, 0xf4, 0x51, 0x7c, 0x94, 0x3d, 0x85, 0x34, 0xb1, 0x3d,
	0x0c, 0xf4, 0x12, 0x3e, 0x7e, 0xdf, 0x1f, 0xbe, 0x7c, 0xe0, 0x8e, 0x4b, 0xc9, 0xd7, 0xcc, 0xcf,
	0x0b, 0xe5, 0x5b, 0x58, 0xa3, 0x7d, 0xe0, 0x33, 0xb1, 0xdb, 0x28, 0x9f, 0x33, 0x99, 0xe9, 0x7c,
	0xcb, 0x99, 0x2e, 0x05, 0xcf, 0xf4, 0xa1, 0x62, 0xb0, 0xda, 0x4a, 0x2d, 0xbd, 0x81, 0x15, 0xc3,
	0xbc, 0x50, 0xb0, 0xf5, 0xc1, 0x7d, 0x00, 0x8d, 0xef, 0xe6, 0xb6, 0x89, 0xad, 0x4a, 0x3f, 0x17,
	0x42, 0xea, 0x5c, 0x97, 0x52, 0x28, 0x6b, 0x1e, 0x7d, 0x80, 0xfe, 0x8c, 0xc9, 0xa4, 0xc9, 0x4d,
	0x0e, 0x15, 0x23, 0x62, 0xb7, 0x19, 0xbd, 0x02, 0xf7, 0xf4, 0xee, 0x5d, 0x82, 0xde, 0x92, 0x2e,
	0x62, 0x32, 0x09, 0xa7, 0x21, 0x79, 0x70, 0xcf, 0xbc, 0x1e, 0xb8, 0x58, 0xd2, 0x47, 0x1a, 0x3d,
	0x53, 0xb7, 0xe3, 0xf5, 0x81, 0x8b, 0xe7, 0x04, 0x67, 0xd1, 0x34, 0x0b, 0x69, 0x42, 0xe6, 0x64,
	0x91, 0xb8, 0x8e, 0x77, 0x0d, 0xfa, 0x4f, 0xd1, 0x04, 0x27, 0x61, 0x44, 0x6b, 0x26, 0x9e, 0x93,
	0x05, 0xa1, 0x13, 0xe2, 0x76, 0xc7, 0xc7, 0x0e, 0x18, 0xbe, 0xcb, 0x0d, 0xfc, 0xb7, 0xfd, 0xf8,
	0xea, 0xb4, 0x45, 0x5c, 0xd7, 0x8e, 0x3b, 0x2f, 0xe3, 0x5f, 0x1f, 0x97, 0xeb, 0x5c, 0x70, 0x28,
	0xb7, 0xdc, 0xe7, 0x4c, 0x98, 0x4f, 0x35, 0xeb, 0x55, 0xa5, 0xfa, 0x63, 0xcc, 0x7b, 0xf3, 0x7e,
	0x3a, 0xdd, 0x19, 0xc6, 0x5f, 0xce, 0x60, 0x66, 0xa3, 0x70, 0xa1, 0xa0, 0x85, 0x35, 0x5a, 0x05,
	0xb0, 0x5e, 0x42, 0x7d, 0x37, 0x7c, 0x8a, 0x0b, 0x95, 0xb6, 0x7c, 0xba, 0x0a, 0x52, 0xc3, 0x1f,
	0x9d, 0xa1, 0x3d, 0x22, 0x84, 0x0b, 0x85, 0x50, 0xab, 0x40, 0x68, 0x15, 0x20, 0x64, 0x34, 0x6f,
	0xe7, 0xa6, 0x58, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x2a, 0xe0, 0x62, 0xe4, 0x01, 0x00,
	0x00,
}
