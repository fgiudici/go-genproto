// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/display_keyword_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [DisplayKeywordViewService.GetDisplayKeywordView][google.ads.googleads.v3.services.DisplayKeywordViewService.GetDisplayKeywordView].
type GetDisplayKeywordViewRequest struct {
	// Required. The resource name of the display keyword view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDisplayKeywordViewRequest) Reset()         { *m = GetDisplayKeywordViewRequest{} }
func (m *GetDisplayKeywordViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetDisplayKeywordViewRequest) ProtoMessage()    {}
func (*GetDisplayKeywordViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_157dba0f7495fb59, []int{0}
}

func (m *GetDisplayKeywordViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDisplayKeywordViewRequest.Unmarshal(m, b)
}
func (m *GetDisplayKeywordViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDisplayKeywordViewRequest.Marshal(b, m, deterministic)
}
func (m *GetDisplayKeywordViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDisplayKeywordViewRequest.Merge(m, src)
}
func (m *GetDisplayKeywordViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetDisplayKeywordViewRequest.Size(m)
}
func (m *GetDisplayKeywordViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDisplayKeywordViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDisplayKeywordViewRequest proto.InternalMessageInfo

func (m *GetDisplayKeywordViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDisplayKeywordViewRequest)(nil), "google.ads.googleads.v3.services.GetDisplayKeywordViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/display_keyword_view_service.proto", fileDescriptor_157dba0f7495fb59)
}

var fileDescriptor_157dba0f7495fb59 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcf, 0x6b, 0xd4, 0x40,
	0x18, 0x65, 0x23, 0x08, 0x06, 0xbd, 0x04, 0xc4, 0x76, 0x2d, 0xba, 0x94, 0x1e, 0xa4, 0xe2, 0x0c,
	0x18, 0x8a, 0x30, 0xfe, 0x80, 0x59, 0x85, 0x15, 0x44, 0x29, 0x0a, 0x8b, 0x48, 0x20, 0x4c, 0x33,
	0x9f, 0x71, 0x30, 0xc9, 0xc4, 0xf9, 0xb2, 0x59, 0x8b, 0x78, 0x11, 0xfc, 0x0b, 0xbc, 0x78, 0xf6,
	0xe8, 0x9f, 0xd2, 0xab, 0x37, 0x41, 0xf0, 0xe0, 0xc9, 0x3f, 0xc1, 0x93, 0xa4, 0x93, 0x49, 0x53,
	0xda, 0x74, 0x6f, 0x8f, 0xbc, 0xf7, 0xbd, 0xef, 0x9b, 0xf7, 0x88, 0xff, 0x30, 0xd5, 0x3a, 0xcd,
	0x80, 0x0a, 0x89, 0xd4, 0xc2, 0x06, 0xd5, 0x21, 0x45, 0x30, 0xb5, 0x4a, 0x00, 0xa9, 0x54, 0x58,
	0x66, 0x62, 0x3f, 0x7e, 0x0b, 0xfb, 0x4b, 0x6d, 0x64, 0x5c, 0x2b, 0x58, 0xc6, 0x2d, 0x4b, 0x4a,
	0xa3, 0x2b, 0x1d, 0x4c, 0xec, 0x24, 0x11, 0x12, 0x49, 0x67, 0x42, 0xea, 0x90, 0x38, 0x93, 0xf1,
	0xbd, 0xa1, 0x35, 0x06, 0x50, 0x2f, 0xcc, 0xd0, 0x1e, 0xeb, 0x3f, 0xde, 0x70, 0xd3, 0xa5, 0xa2,
	0xa2, 0x28, 0x74, 0x25, 0x2a, 0xa5, 0x0b, 0x6c, 0xd9, 0x2b, 0x3d, 0x36, 0xc9, 0x14, 0x14, 0x55,
	0x4b, 0x5c, 0xef, 0x11, 0xaf, 0x15, 0x64, 0x32, 0xde, 0x83, 0x37, 0xa2, 0x56, 0xda, 0xb4, 0x82,
	0xf5, 0x9e, 0xc0, 0x1d, 0x62, 0xa9, 0xcd, 0xf7, 0xfe, 0xc6, 0x0c, 0xaa, 0x47, 0xf6, 0xa6, 0x27,
	0xf6, 0xa4, 0xb9, 0x82, 0xe5, 0x73, 0x78, 0xb7, 0x00, 0xac, 0x82, 0x97, 0xfe, 0x25, 0x37, 0x11,
	0x17, 0x22, 0x87, 0xb5, 0xd1, 0x64, 0x74, 0xe3, 0xc2, 0x34, 0xfc, 0xcd, 0xbd, 0x7f, 0xfc, 0x96,
	0x7f, 0xf3, 0x28, 0x86, 0x16, 0x95, 0x0a, 0x49, 0xa2, 0x73, 0x7a, 0x8a, 0xe5, 0x45, 0xe7, 0xf4,
	0x4c, 0xe4, 0x70, 0xfb, 0xab, 0xe7, 0xaf, 0x9f, 0x14, 0xbd, 0xb0, 0x49, 0x06, 0xbf, 0x46, 0xfe,
	0xe5, 0x53, 0x0f, 0x0b, 0x1e, 0x90, 0x55, 0x2d, 0x90, 0xb3, 0x5e, 0x34, 0xde, 0x19, 0x9c, 0xef,
	0x3a, 0x22, 0x27, 0xa7, 0x37, 0x9f, 0xfe, 0xe4, 0xc7, 0x93, 0xf8, 0xf4, 0xe3, 0xcf, 0x17, 0xef,
	0x4e, 0xb0, 0xd3, 0xb4, 0xfb, 0xe1, 0x18, 0x73, 0x3f, 0x59, 0x60, 0xa5, 0x73, 0x30, 0x48, 0xb7,
	0x5d, 0xdd, 0x3d, 0x2b, 0xa4, 0xdb, 0x1f, 0xc7, 0x57, 0x0f, 0xf8, 0xda, 0x50, 0x76, 0xd3, 0xcf,
	0x9e, 0xbf, 0x95, 0xe8, 0x7c, 0xe5, 0x43, 0xa7, 0xd7, 0x06, 0x03, 0xdc, 0x6d, 0xda, 0xdd, 0x1d,
	0xbd, 0x7a, 0xdc, 0x7a, 0xa4, 0x3a, 0x13, 0x45, 0x4a, 0xb4, 0x49, 0x69, 0x0a, 0xc5, 0x61, 0xf7,
	0xf4, 0x68, 0xeb, 0xf0, 0x6f, 0x71, 0xd7, 0x81, 0x6f, 0xde, 0xb9, 0x19, 0xe7, 0xdf, 0xbd, 0xc9,
	0xcc, 0x1a, 0x72, 0x89, 0xc4, 0xc2, 0x06, 0xcd, 0x43, 0xd2, 0x2e, 0xc6, 0x03, 0x27, 0x89, 0xb8,
	0xc4, 0xa8, 0x93, 0x44, 0xf3, 0x30, 0x72, 0x92, 0xbf, 0xde, 0x96, 0xfd, 0xce, 0x18, 0x97, 0xc8,
	0x58, 0x27, 0x62, 0x6c, 0x1e, 0x32, 0xe6, 0x64, 0x7b, 0xe7, 0x0f, 0xef, 0x0c, 0xff, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x44, 0xdc, 0xd5, 0xd2, 0xbd, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DisplayKeywordViewServiceClient is the client API for DisplayKeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DisplayKeywordViewServiceClient interface {
	// Returns the requested display keyword view in full detail.
	GetDisplayKeywordView(ctx context.Context, in *GetDisplayKeywordViewRequest, opts ...grpc.CallOption) (*resources.DisplayKeywordView, error)
}

type displayKeywordViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDisplayKeywordViewServiceClient(cc grpc.ClientConnInterface) DisplayKeywordViewServiceClient {
	return &displayKeywordViewServiceClient{cc}
}

func (c *displayKeywordViewServiceClient) GetDisplayKeywordView(ctx context.Context, in *GetDisplayKeywordViewRequest, opts ...grpc.CallOption) (*resources.DisplayKeywordView, error) {
	out := new(resources.DisplayKeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.DisplayKeywordViewService/GetDisplayKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisplayKeywordViewServiceServer is the server API for DisplayKeywordViewService service.
type DisplayKeywordViewServiceServer interface {
	// Returns the requested display keyword view in full detail.
	GetDisplayKeywordView(context.Context, *GetDisplayKeywordViewRequest) (*resources.DisplayKeywordView, error)
}

// UnimplementedDisplayKeywordViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDisplayKeywordViewServiceServer struct {
}

func (*UnimplementedDisplayKeywordViewServiceServer) GetDisplayKeywordView(ctx context.Context, req *GetDisplayKeywordViewRequest) (*resources.DisplayKeywordView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDisplayKeywordView not implemented")
}

func RegisterDisplayKeywordViewServiceServer(s *grpc.Server, srv DisplayKeywordViewServiceServer) {
	s.RegisterService(&_DisplayKeywordViewService_serviceDesc, srv)
}

func _DisplayKeywordViewService_GetDisplayKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDisplayKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayKeywordViewServiceServer).GetDisplayKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.DisplayKeywordViewService/GetDisplayKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayKeywordViewServiceServer).GetDisplayKeywordView(ctx, req.(*GetDisplayKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DisplayKeywordViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.DisplayKeywordViewService",
	HandlerType: (*DisplayKeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDisplayKeywordView",
			Handler:    _DisplayKeywordViewService_GetDisplayKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/display_keyword_view_service.proto",
}
