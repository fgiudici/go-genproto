// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/click_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// A click view with metrics aggregated at each click level, including both
// valid and invalid clicks. For non-Search campaigns, metrics.clicks
// represents the number of valid and invalid interactions.
// Queries including ClickView must have a filter limiting the results to one
// day and can be requested for dates back to 90 days before the time of the
// request.
type ClickView struct {
	// Output only. The resource name of the click view.
	// Click view resource names have the form:
	//
	// `customers/{customer_id}/clickViews/{date (yyyy-MM-dd)}~{gclid}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The Google Click ID.
	Gclid *wrappers.StringValue `protobuf:"bytes,2,opt,name=gclid,proto3" json:"gclid,omitempty"`
	// Output only. The location criteria matching the area of interest associated with the
	// impression.
	AreaOfInterest *common.ClickLocation `protobuf:"bytes,3,opt,name=area_of_interest,json=areaOfInterest,proto3" json:"area_of_interest,omitempty"`
	// Output only. The location criteria matching the location of presence associated with the
	// impression.
	LocationOfPresence *common.ClickLocation `protobuf:"bytes,4,opt,name=location_of_presence,json=locationOfPresence,proto3" json:"location_of_presence,omitempty"`
	// Output only. Page number in search results where the ad was shown.
	PageNumber           *wrappers.Int64Value `protobuf:"bytes,5,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClickView) Reset()         { *m = ClickView{} }
func (m *ClickView) String() string { return proto.CompactTextString(m) }
func (*ClickView) ProtoMessage()    {}
func (*ClickView) Descriptor() ([]byte, []int) {
	return fileDescriptor_c61fa672950e615e, []int{0}
}

func (m *ClickView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickView.Unmarshal(m, b)
}
func (m *ClickView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickView.Marshal(b, m, deterministic)
}
func (m *ClickView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickView.Merge(m, src)
}
func (m *ClickView) XXX_Size() int {
	return xxx_messageInfo_ClickView.Size(m)
}
func (m *ClickView) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickView.DiscardUnknown(m)
}

var xxx_messageInfo_ClickView proto.InternalMessageInfo

func (m *ClickView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ClickView) GetGclid() *wrappers.StringValue {
	if m != nil {
		return m.Gclid
	}
	return nil
}

func (m *ClickView) GetAreaOfInterest() *common.ClickLocation {
	if m != nil {
		return m.AreaOfInterest
	}
	return nil
}

func (m *ClickView) GetLocationOfPresence() *common.ClickLocation {
	if m != nil {
		return m.LocationOfPresence
	}
	return nil
}

func (m *ClickView) GetPageNumber() *wrappers.Int64Value {
	if m != nil {
		return m.PageNumber
	}
	return nil
}

func init() {
	proto.RegisterType((*ClickView)(nil), "google.ads.googleads.v1.resources.ClickView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/click_view.proto", fileDescriptor_c61fa672950e615e)
}

var fileDescriptor_c61fa672950e615e = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6b, 0x13, 0x4f,
	0x18, 0xc7, 0x49, 0xf6, 0xd7, 0x1f, 0x74, 0xaa, 0x45, 0x16, 0x0f, 0xb1, 0x16, 0x6d, 0x8b, 0x85,
	0x22, 0x3a, 0x43, 0x5a, 0x51, 0x58, 0x4f, 0x13, 0x0f, 0xa5, 0x22, 0x4d, 0x88, 0x98, 0x83, 0x04,
	0x96, 0xc9, 0xec, 0x93, 0x75, 0x70, 0x77, 0x66, 0x99, 0xd9, 0x24, 0x07, 0xe9, 0x9b, 0xf1, 0xe8,
	0x4b, 0xf1, 0x55, 0xf4, 0xdc, 0xab, 0x37, 0xbd, 0xc8, 0xee, 0xfc, 0x69, 0x40, 0xaa, 0xe2, 0xed,
	0x9b, 0x3c, 0xdf, 0xef, 0xe7, 0x79, 0xe6, 0x99, 0x59, 0x74, 0x9c, 0x2b, 0x95, 0x17, 0x40, 0x58,
	0x66, 0x88, 0x95, 0x8d, 0x5a, 0xf6, 0x89, 0x06, 0xa3, 0x16, 0x9a, 0x83, 0x21, 0xbc, 0x10, 0xfc,
	0x63, 0xba, 0x14, 0xb0, 0xc2, 0x95, 0x56, 0xb5, 0x8a, 0xf7, 0xad, 0x11, 0xb3, 0xcc, 0xe0, 0x90,
	0xc1, 0xcb, 0x3e, 0x0e, 0x99, 0x9d, 0x93, 0x9b, 0xb0, 0x5c, 0x95, 0xa5, 0x92, 0x8e, 0x59, 0x28,
	0xce, 0x6a, 0xa1, 0xa4, 0xe5, 0xee, 0x3c, 0xf4, 0xa1, 0x4a, 0x90, 0xb9, 0x80, 0x22, 0x4b, 0x67,
	0xf0, 0x81, 0x2d, 0x85, 0xd2, 0xce, 0x70, 0x6f, 0xcd, 0xe0, 0x7b, 0xb9, 0xd2, 0x03, 0x57, 0x6a,
	0x7f, 0xcd, 0x16, 0x73, 0xb2, 0xd2, 0xac, 0xaa, 0x40, 0x1b, 0x57, 0xdf, 0x5d, 0x8b, 0x32, 0x29,
	0x55, 0xdd, 0x36, 0x76, 0xd5, 0x83, 0x6f, 0x11, 0xda, 0x7c, 0xd5, 0x8c, 0x34, 0x11, 0xb0, 0x8a,
	0x87, 0xe8, 0xb6, 0xa7, 0xa7, 0x92, 0x95, 0xd0, 0xeb, 0xec, 0x75, 0x8e, 0x36, 0x07, 0x8f, 0x2f,
	0x69, 0xf4, 0x9d, 0x3e, 0x42, 0x07, 0xd7, 0x67, 0x76, 0xaa, 0x12, 0x06, 0x73, 0x55, 0x92, 0x80,
	0x18, 0xdf, 0xf2, 0x80, 0x73, 0x56, 0x42, 0xfc, 0x02, 0x6d, 0xe4, 0xbc, 0x10, 0x59, 0xaf, 0xbb,
	0xd7, 0x39, 0xda, 0x3a, 0xde, 0x75, 0x39, 0xec, 0x87, 0xc5, 0x6f, 0x6b, 0x2d, 0x64, 0x3e, 0x61,
	0xc5, 0x02, 0x06, 0xd1, 0x25, 0x8d, 0xc6, 0xd6, 0x1f, 0x4f, 0xd1, 0x1d, 0xa6, 0x81, 0xa5, 0x6a,
	0x9e, 0x0a, 0x59, 0x83, 0x06, 0x53, 0xf7, 0xa2, 0x96, 0xf1, 0x14, 0xdf, 0x74, 0x09, 0x76, 0xc3,
	0xb8, 0x9d, 0xe5, 0x8d, 0x5b, 0xb0, 0x85, 0x6e, 0x37, 0xac, 0xe1, 0xfc, 0xcc, 0x91, 0xe2, 0x0c,
	0xdd, 0xf5, 0x37, 0xd0, 0x74, 0xa8, 0x34, 0x18, 0x90, 0x1c, 0x7a, 0xff, 0xfd, 0x73, 0x87, 0xd8,
	0xf3, 0x86, 0xf3, 0x91, 0xa3, 0xc5, 0x14, 0x6d, 0x55, 0x2c, 0x87, 0x54, 0x2e, 0xca, 0x19, 0xe8,
	0xde, 0x46, 0x0b, 0xbf, 0xff, 0xcb, 0x0a, 0xce, 0x64, 0xfd, 0xfc, 0xd9, 0xda, 0x06, 0x50, 0x13,
	0x3a, 0x6f, 0x33, 0xc9, 0xbb, 0x2b, 0x3a, 0xfe, 0x9b, 0xb5, 0xc7, 0x4f, 0xf8, 0xc2, 0xd4, 0xaa,
	0x04, 0x6d, 0xc8, 0x27, 0x2f, 0x2f, 0xec, 0x63, 0x6b, 0xea, 0xcd, 0xbf, 0xe1, 0x31, 0x5f, 0x0c,
	0x7e, 0x74, 0xd0, 0x21, 0x57, 0x25, 0xfe, 0xe3, 0x73, 0x1e, 0x6c, 0x87, 0x16, 0xa3, 0x66, 0xde,
	0x51, 0xe7, 0xfd, 0x6b, 0x17, 0xca, 0x55, 0xc1, 0x64, 0x8e, 0x95, 0xce, 0x49, 0x0e, 0xb2, 0x3d,
	0x0d, 0xb9, 0x1e, 0xef, 0x37, 0x9f, 0xd5, 0xcb, 0xa0, 0x3e, 0x77, 0xa3, 0x53, 0x4a, 0xbf, 0x74,
	0xf7, 0x4f, 0x2d, 0x92, 0x66, 0x06, 0x5b, 0xd9, 0xa8, 0x49, 0x1f, 0x8f, 0xbd, 0xf3, 0xab, 0xf7,
	0x4c, 0x69, 0x66, 0xa6, 0xc1, 0x33, 0x9d, 0xf4, 0xa7, 0xc1, 0x73, 0xd5, 0x3d, 0xb4, 0x85, 0x24,
	0xa1, 0x99, 0x49, 0x92, 0xe0, 0x4a, 0x92, 0x49, 0x3f, 0x49, 0x82, 0x6f, 0xf6, 0x7f, 0x3b, 0xec,
	0xc9, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xef, 0xfd, 0x96, 0x02, 0x04, 0x00, 0x00,
}
