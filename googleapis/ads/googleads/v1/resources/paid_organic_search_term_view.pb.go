// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/paid_organic_search_term_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A paid organic search term view providing a view of search stats across
// ads and organic listings aggregated by search term at the ad group level.
type PaidOrganicSearchTermView struct {
	// Output only. The resource name of the search term view.
	// Search term view resource names have the form:
	//
	// `customers/{customer_id}/paidOrganicSearchTermViews/{campaign_id}~
	// {ad_group_id}~{URL-base64 search term}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The search term.
	SearchTerm           *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PaidOrganicSearchTermView) Reset()         { *m = PaidOrganicSearchTermView{} }
func (m *PaidOrganicSearchTermView) String() string { return proto.CompactTextString(m) }
func (*PaidOrganicSearchTermView) ProtoMessage()    {}
func (*PaidOrganicSearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e209ceb6c96a605, []int{0}
}

func (m *PaidOrganicSearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaidOrganicSearchTermView.Unmarshal(m, b)
}
func (m *PaidOrganicSearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaidOrganicSearchTermView.Marshal(b, m, deterministic)
}
func (m *PaidOrganicSearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaidOrganicSearchTermView.Merge(m, src)
}
func (m *PaidOrganicSearchTermView) XXX_Size() int {
	return xxx_messageInfo_PaidOrganicSearchTermView.Size(m)
}
func (m *PaidOrganicSearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_PaidOrganicSearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_PaidOrganicSearchTermView proto.InternalMessageInfo

func (m *PaidOrganicSearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *PaidOrganicSearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func init() {
	proto.RegisterType((*PaidOrganicSearchTermView)(nil), "google.ads.googleads.v1.resources.PaidOrganicSearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/paid_organic_search_term_view.proto", fileDescriptor_8e209ceb6c96a605)
}

var fileDescriptor_8e209ceb6c96a605 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x6a, 0xdd, 0x30,
	0x18, 0x85, 0xb1, 0x0d, 0x85, 0x3a, 0xed, 0xe2, 0x29, 0x09, 0x21, 0x4d, 0x0a, 0x81, 0x4c, 0x12,
	0x4e, 0x3b, 0xa9, 0x93, 0x0c, 0x25, 0xd0, 0xa1, 0xb9, 0xdc, 0x14, 0x0f, 0xc5, 0x60, 0x74, 0xad,
	0x3f, 0x8a, 0xc0, 0x96, 0x8c, 0x64, 0xfb, 0x0e, 0x21, 0x63, 0x87, 0xd2, 0xb7, 0xe8, 0xd8, 0x47,
	0x29, 0x7d, 0x88, 0xcc, 0x79, 0x84, 0x4e, 0xe5, 0xda, 0x96, 0x72, 0x87, 0x26, 0x85, 0x6e, 0xc7,
	0x9c, 0xa3, 0x4f, 0xff, 0xd1, 0xef, 0xf8, 0xbd, 0xd0, 0x5a, 0xd4, 0x80, 0x19, 0xb7, 0x78, 0x92,
	0x1b, 0x35, 0xa4, 0xd8, 0x80, 0xd5, 0xbd, 0xa9, 0xc0, 0xe2, 0x96, 0x49, 0x5e, 0x6a, 0x23, 0x98,
	0x92, 0x55, 0x69, 0x81, 0x99, 0xea, 0xba, 0xec, 0xc0, 0x34, 0xe5, 0x20, 0x61, 0x8d, 0x5a, 0xa3,
	0x3b, 0x9d, 0x1c, 0x4f, 0x67, 0x11, 0xe3, 0x16, 0x79, 0x0c, 0x1a, 0x52, 0xe4, 0x31, 0xfb, 0xaf,
	0xdc, 0x4d, 0xad, 0xc4, 0x57, 0x12, 0x6a, 0x5e, 0xae, 0xe0, 0x9a, 0x0d, 0x52, 0x9b, 0x89, 0xb1,
	0xbf, 0xb7, 0x15, 0x70, 0xc7, 0x66, 0xeb, 0x70, 0xb6, 0xc6, 0xaf, 0x55, 0x7f, 0x85, 0xd7, 0x86,
	0xb5, 0x2d, 0x18, 0x3b, 0xfb, 0x07, 0x5b, 0x47, 0x99, 0x52, 0xba, 0x63, 0x9d, 0xd4, 0x6a, 0x76,
	0x5f, 0xff, 0x0a, 0xe3, 0xbd, 0x05, 0x93, 0xfc, 0x62, 0xea, 0x70, 0x39, 0x56, 0xf8, 0x04, 0xa6,
	0xc9, 0x25, 0xac, 0x93, 0x32, 0x7e, 0xe9, 0x6e, 0x2b, 0x15, 0x6b, 0x60, 0x37, 0x38, 0x0a, 0x4e,
	0x9f, 0x67, 0xe4, 0x8e, 0x46, 0xbf, 0xe9, 0xdb, 0xf8, 0xec, 0xa1, 0xce, 0xac, 0x5a, 0x69, 0x51,
	0xa5, 0x1b, 0xfc, 0x28, 0x72, 0xf9, 0xc2, 0x01, 0x3f, 0xb2, 0x06, 0x92, 0x2c, 0xde, 0xd9, 0x7a,
	0xb5, 0xdd, 0xf0, 0x28, 0x38, 0xdd, 0x39, 0x3b, 0x98, 0x69, 0xc8, 0x55, 0x42, 0x97, 0x9d, 0x91,
	0x4a, 0xe4, 0xac, 0xee, 0x21, 0x8b, 0xee, 0x68, 0xb4, 0x8c, 0xad, 0xa7, 0x92, 0xaf, 0xc1, 0x3d,
	0xfd, 0x12, 0xfc, 0xcf, 0x30, 0xc9, 0x45, 0xd5, 0xdb, 0x4e, 0x37, 0x60, 0x2c, 0xbe, 0x71, 0xf2,
	0x76, 0x5c, 0xea, 0x5f, 0xf3, 0x16, 0xdf, 0x3c, 0xb9, 0xf0, 0xdb, 0xec, 0x5b, 0x18, 0x9f, 0x54,
	0xba, 0x41, 0xff, 0x5c, 0x79, 0x76, 0xf8, 0xe8, 0x54, 0x8b, 0x4d, 0xeb, 0x45, 0xf0, 0xf9, 0xc3,
	0x0c, 0x11, 0xba, 0x66, 0x4a, 0x20, 0x6d, 0x04, 0x16, 0xa0, 0xc6, 0x37, 0xc1, 0x0f, 0x0d, 0x9f,
	0xf8, 0x3b, 0xdf, 0x79, 0xf5, 0x3d, 0x8c, 0xce, 0x29, 0xfd, 0x11, 0x1e, 0x9f, 0x4f, 0x48, 0xca,
	0x2d, 0x9a, 0xe4, 0x46, 0xe5, 0x29, 0x5a, 0xba, 0xe4, 0x4f, 0x97, 0x29, 0x28, 0xb7, 0x85, 0xcf,
	0x14, 0x79, 0x5a, 0xf8, 0xcc, 0x7d, 0x78, 0x32, 0x19, 0x84, 0x50, 0x6e, 0x09, 0xf1, 0x29, 0x42,
	0xf2, 0x94, 0x10, 0x9f, 0x5b, 0x3d, 0x1b, 0x87, 0x7d, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0xde,
	0x00, 0x09, 0x67, 0x49, 0x03, 0x00, 0x00,
}
