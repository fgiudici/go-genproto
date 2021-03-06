// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/offline_user_data_job_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for
// [OfflineUserDataJobService.CreateOfflineUserDataJobRequest][]
type CreateOfflineUserDataJobRequest struct {
	// Required. The ID of the customer for which to create an offline user data job.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The offline user data job to be created.
	Job                  *resources.OfflineUserDataJob `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CreateOfflineUserDataJobRequest) Reset()         { *m = CreateOfflineUserDataJobRequest{} }
func (m *CreateOfflineUserDataJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOfflineUserDataJobRequest) ProtoMessage()    {}
func (*CreateOfflineUserDataJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51833c2667e4159e, []int{0}
}

func (m *CreateOfflineUserDataJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOfflineUserDataJobRequest.Unmarshal(m, b)
}
func (m *CreateOfflineUserDataJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOfflineUserDataJobRequest.Marshal(b, m, deterministic)
}
func (m *CreateOfflineUserDataJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOfflineUserDataJobRequest.Merge(m, src)
}
func (m *CreateOfflineUserDataJobRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOfflineUserDataJobRequest.Size(m)
}
func (m *CreateOfflineUserDataJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOfflineUserDataJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOfflineUserDataJobRequest proto.InternalMessageInfo

func (m *CreateOfflineUserDataJobRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *CreateOfflineUserDataJobRequest) GetJob() *resources.OfflineUserDataJob {
	if m != nil {
		return m.Job
	}
	return nil
}

// Response message for
// [OfflineUserDataJobService.CreateOfflineUserDataJobResponse][]
type CreateOfflineUserDataJobResponse struct {
	// The resource name of the OfflineUserDataJob.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOfflineUserDataJobResponse) Reset()         { *m = CreateOfflineUserDataJobResponse{} }
func (m *CreateOfflineUserDataJobResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOfflineUserDataJobResponse) ProtoMessage()    {}
func (*CreateOfflineUserDataJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51833c2667e4159e, []int{1}
}

func (m *CreateOfflineUserDataJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOfflineUserDataJobResponse.Unmarshal(m, b)
}
func (m *CreateOfflineUserDataJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOfflineUserDataJobResponse.Marshal(b, m, deterministic)
}
func (m *CreateOfflineUserDataJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOfflineUserDataJobResponse.Merge(m, src)
}
func (m *CreateOfflineUserDataJobResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOfflineUserDataJobResponse.Size(m)
}
func (m *CreateOfflineUserDataJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOfflineUserDataJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOfflineUserDataJobResponse proto.InternalMessageInfo

func (m *CreateOfflineUserDataJobResponse) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [OfflineUserDataJobService.GetOfflineUserDataJob][google.ads.googleads.v3.services.OfflineUserDataJobService.GetOfflineUserDataJob]
type GetOfflineUserDataJobRequest struct {
	// Required. The resource name of the OfflineUserDataJob to get.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOfflineUserDataJobRequest) Reset()         { *m = GetOfflineUserDataJobRequest{} }
func (m *GetOfflineUserDataJobRequest) String() string { return proto.CompactTextString(m) }
func (*GetOfflineUserDataJobRequest) ProtoMessage()    {}
func (*GetOfflineUserDataJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51833c2667e4159e, []int{2}
}

func (m *GetOfflineUserDataJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOfflineUserDataJobRequest.Unmarshal(m, b)
}
func (m *GetOfflineUserDataJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOfflineUserDataJobRequest.Marshal(b, m, deterministic)
}
func (m *GetOfflineUserDataJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOfflineUserDataJobRequest.Merge(m, src)
}
func (m *GetOfflineUserDataJobRequest) XXX_Size() int {
	return xxx_messageInfo_GetOfflineUserDataJobRequest.Size(m)
}
func (m *GetOfflineUserDataJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOfflineUserDataJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOfflineUserDataJobRequest proto.InternalMessageInfo

func (m *GetOfflineUserDataJobRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [OfflineUserDataJobService.RunOfflineUserDataJob][google.ads.googleads.v3.services.OfflineUserDataJobService.RunOfflineUserDataJob]
type RunOfflineUserDataJobRequest struct {
	// Required. The resource name of the OfflineUserDataJob to run.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunOfflineUserDataJobRequest) Reset()         { *m = RunOfflineUserDataJobRequest{} }
func (m *RunOfflineUserDataJobRequest) String() string { return proto.CompactTextString(m) }
func (*RunOfflineUserDataJobRequest) ProtoMessage()    {}
func (*RunOfflineUserDataJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51833c2667e4159e, []int{3}
}

func (m *RunOfflineUserDataJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunOfflineUserDataJobRequest.Unmarshal(m, b)
}
func (m *RunOfflineUserDataJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunOfflineUserDataJobRequest.Marshal(b, m, deterministic)
}
func (m *RunOfflineUserDataJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunOfflineUserDataJobRequest.Merge(m, src)
}
func (m *RunOfflineUserDataJobRequest) XXX_Size() int {
	return xxx_messageInfo_RunOfflineUserDataJobRequest.Size(m)
}
func (m *RunOfflineUserDataJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunOfflineUserDataJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunOfflineUserDataJobRequest proto.InternalMessageInfo

func (m *RunOfflineUserDataJobRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [OfflineUserDataJobService.AddOfflineUserDataJobOperations][google.ads.googleads.v3.services.OfflineUserDataJobService.AddOfflineUserDataJobOperations]
type AddOfflineUserDataJobOperationsRequest struct {
	// Required. The resource name of the OfflineUserDataJob.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// True to enable partial failure for the offline user data job.
	EnablePartialFailure *wrappers.BoolValue `protobuf:"bytes,2,opt,name=enable_partial_failure,json=enablePartialFailure,proto3" json:"enable_partial_failure,omitempty"`
	// Required. The list of operations to be done.
	Operations           []*OfflineUserDataJobOperation `protobuf:"bytes,3,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AddOfflineUserDataJobOperationsRequest) Reset() {
	*m = AddOfflineUserDataJobOperationsRequest{}
}
func (m *AddOfflineUserDataJobOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*AddOfflineUserDataJobOperationsRequest) ProtoMessage()    {}
func (*AddOfflineUserDataJobOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_51833c2667e4159e, []int{4}
}

func (m *AddOfflineUserDataJobOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddOfflineUserDataJobOperationsRequest.Unmarshal(m, b)
}
func (m *AddOfflineUserDataJobOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddOfflineUserDataJobOperationsRequest.Marshal(b, m, deterministic)
}
func (m *AddOfflineUserDataJobOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddOfflineUserDataJobOperationsRequest.Merge(m, src)
}
func (m *AddOfflineUserDataJobOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_AddOfflineUserDataJobOperationsRequest.Size(m)
}
func (m *AddOfflineUserDataJobOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddOfflineUserDataJobOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddOfflineUserDataJobOperationsRequest proto.InternalMessageInfo

func (m *AddOfflineUserDataJobOperationsRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AddOfflineUserDataJobOperationsRequest) GetEnablePartialFailure() *wrappers.BoolValue {
	if m != nil {
		return m.EnablePartialFailure
	}
	return nil
}

func (m *AddOfflineUserDataJobOperationsRequest) GetOperations() []*OfflineUserDataJobOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// Operation to be made for the AddOfflineUserDataJobOperationsRequest.
type OfflineUserDataJobOperation struct {
	// Operation to be made for the AddOfflineUserDataJobOperationsRequest.
	//
	// Types that are valid to be assigned to Operation:
	//	*OfflineUserDataJobOperation_Create
	//	*OfflineUserDataJobOperation_Remove
	//	*OfflineUserDataJobOperation_RemoveAll
	Operation            isOfflineUserDataJobOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *OfflineUserDataJobOperation) Reset()         { *m = OfflineUserDataJobOperation{} }
func (m *OfflineUserDataJobOperation) String() string { return proto.CompactTextString(m) }
func (*OfflineUserDataJobOperation) ProtoMessage()    {}
func (*OfflineUserDataJobOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_51833c2667e4159e, []int{5}
}

func (m *OfflineUserDataJobOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfflineUserDataJobOperation.Unmarshal(m, b)
}
func (m *OfflineUserDataJobOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfflineUserDataJobOperation.Marshal(b, m, deterministic)
}
func (m *OfflineUserDataJobOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfflineUserDataJobOperation.Merge(m, src)
}
func (m *OfflineUserDataJobOperation) XXX_Size() int {
	return xxx_messageInfo_OfflineUserDataJobOperation.Size(m)
}
func (m *OfflineUserDataJobOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_OfflineUserDataJobOperation.DiscardUnknown(m)
}

var xxx_messageInfo_OfflineUserDataJobOperation proto.InternalMessageInfo

type isOfflineUserDataJobOperation_Operation interface {
	isOfflineUserDataJobOperation_Operation()
}

type OfflineUserDataJobOperation_Create struct {
	Create *common.UserData `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type OfflineUserDataJobOperation_Remove struct {
	Remove *common.UserData `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

type OfflineUserDataJobOperation_RemoveAll struct {
	RemoveAll bool `protobuf:"varint,3,opt,name=remove_all,json=removeAll,proto3,oneof"`
}

func (*OfflineUserDataJobOperation_Create) isOfflineUserDataJobOperation_Operation() {}

func (*OfflineUserDataJobOperation_Remove) isOfflineUserDataJobOperation_Operation() {}

func (*OfflineUserDataJobOperation_RemoveAll) isOfflineUserDataJobOperation_Operation() {}

func (m *OfflineUserDataJobOperation) GetOperation() isOfflineUserDataJobOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *OfflineUserDataJobOperation) GetCreate() *common.UserData {
	if x, ok := m.GetOperation().(*OfflineUserDataJobOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *OfflineUserDataJobOperation) GetRemove() *common.UserData {
	if x, ok := m.GetOperation().(*OfflineUserDataJobOperation_Remove); ok {
		return x.Remove
	}
	return nil
}

func (m *OfflineUserDataJobOperation) GetRemoveAll() bool {
	if x, ok := m.GetOperation().(*OfflineUserDataJobOperation_RemoveAll); ok {
		return x.RemoveAll
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OfflineUserDataJobOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OfflineUserDataJobOperation_Create)(nil),
		(*OfflineUserDataJobOperation_Remove)(nil),
		(*OfflineUserDataJobOperation_RemoveAll)(nil),
	}
}

// Response message for
// [OfflineUserDataJobService.AddOfflineUserDataJobOperations][google.ads.googleads.v3.services.OfflineUserDataJobService.AddOfflineUserDataJobOperations]
type AddOfflineUserDataJobOperationsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError  *status.Status `protobuf:"bytes,1,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddOfflineUserDataJobOperationsResponse) Reset() {
	*m = AddOfflineUserDataJobOperationsResponse{}
}
func (m *AddOfflineUserDataJobOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*AddOfflineUserDataJobOperationsResponse) ProtoMessage()    {}
func (*AddOfflineUserDataJobOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_51833c2667e4159e, []int{6}
}

func (m *AddOfflineUserDataJobOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddOfflineUserDataJobOperationsResponse.Unmarshal(m, b)
}
func (m *AddOfflineUserDataJobOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddOfflineUserDataJobOperationsResponse.Marshal(b, m, deterministic)
}
func (m *AddOfflineUserDataJobOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddOfflineUserDataJobOperationsResponse.Merge(m, src)
}
func (m *AddOfflineUserDataJobOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_AddOfflineUserDataJobOperationsResponse.Size(m)
}
func (m *AddOfflineUserDataJobOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddOfflineUserDataJobOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddOfflineUserDataJobOperationsResponse proto.InternalMessageInfo

func (m *AddOfflineUserDataJobOperationsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOfflineUserDataJobRequest)(nil), "google.ads.googleads.v3.services.CreateOfflineUserDataJobRequest")
	proto.RegisterType((*CreateOfflineUserDataJobResponse)(nil), "google.ads.googleads.v3.services.CreateOfflineUserDataJobResponse")
	proto.RegisterType((*GetOfflineUserDataJobRequest)(nil), "google.ads.googleads.v3.services.GetOfflineUserDataJobRequest")
	proto.RegisterType((*RunOfflineUserDataJobRequest)(nil), "google.ads.googleads.v3.services.RunOfflineUserDataJobRequest")
	proto.RegisterType((*AddOfflineUserDataJobOperationsRequest)(nil), "google.ads.googleads.v3.services.AddOfflineUserDataJobOperationsRequest")
	proto.RegisterType((*OfflineUserDataJobOperation)(nil), "google.ads.googleads.v3.services.OfflineUserDataJobOperation")
	proto.RegisterType((*AddOfflineUserDataJobOperationsResponse)(nil), "google.ads.googleads.v3.services.AddOfflineUserDataJobOperationsResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/offline_user_data_job_service.proto", fileDescriptor_51833c2667e4159e)
}

var fileDescriptor_51833c2667e4159e = []byte{
	// 924 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdb, 0x8e, 0xdb, 0x44,
	0x18, 0xae, 0x1d, 0xa9, 0xea, 0x4e, 0x40, 0x48, 0x03, 0x0b, 0x69, 0x5a, 0xba, 0x91, 0xbb, 0x82,
	0x28, 0x94, 0xb1, 0xb4, 0x51, 0x39, 0x18, 0xa5, 0x92, 0xd3, 0xc3, 0xa6, 0x08, 0xe8, 0x2a, 0x55,
	0x57, 0x08, 0xad, 0x64, 0x8d, 0xed, 0x49, 0xf0, 0xca, 0x9e, 0x71, 0x67, 0xec, 0x00, 0xaa, 0x7a,
	0x83, 0xd4, 0x27, 0x40, 0x3c, 0x00, 0x5c, 0xf2, 0x04, 0x7d, 0x86, 0xbd, 0xed, 0xdd, 0x4a, 0x48,
	0xbd, 0xe0, 0x8a, 0x27, 0x40, 0x5c, 0x20, 0x14, 0xcf, 0x4c, 0x0e, 0x9b, 0x78, 0x5d, 0x16, 0x50,
	0xef, 0x26, 0xf3, 0xff, 0xff, 0xf7, 0x9f, 0xbe, 0x7c, 0x63, 0x70, 0x6b, 0xcc, 0xd8, 0x38, 0x26,
	0x36, 0x0e, 0x85, 0x2d, 0x8f, 0xd3, 0xd3, 0xa4, 0x6b, 0x0b, 0xc2, 0x27, 0x51, 0x40, 0x84, 0xcd,
	0x46, 0xa3, 0x38, 0xa2, 0xc4, 0xcb, 0x05, 0xe1, 0x5e, 0x88, 0x33, 0xec, 0x1d, 0x32, 0xdf, 0x53,
	0x66, 0x94, 0x72, 0x96, 0x31, 0xd8, 0x92, 0xa1, 0x08, 0x87, 0x02, 0xcd, 0x50, 0xd0, 0xa4, 0x8b,
	0x34, 0x4a, 0xf3, 0x83, 0xb2, 0x3c, 0x01, 0x4b, 0x12, 0x46, 0x57, 0xb3, 0x48, 0xe4, 0x66, 0xaf,
	0x2c, 0x8e, 0x13, 0xc1, 0x72, 0x5e, 0x5a, 0xa0, 0x0a, 0xbf, 0xac, 0xc3, 0xd3, 0xc8, 0xc6, 0x94,
	0xb2, 0x0c, 0x67, 0x11, 0xa3, 0x42, 0x59, 0xdf, 0x5a, 0xb0, 0x06, 0x71, 0x44, 0x68, 0xa6, 0x0c,
	0x5b, 0x0b, 0x86, 0x51, 0x44, 0xe2, 0xd0, 0xf3, 0xc9, 0xd7, 0x78, 0x12, 0x31, 0xae, 0x1c, 0x2e,
	0x2e, 0x38, 0xe8, 0x4a, 0x94, 0xe9, 0xaa, 0x32, 0xc5, 0x8c, 0x8e, 0x79, 0x4e, 0x69, 0x44, 0xc7,
	0x36, 0x4b, 0x09, 0x5f, 0xca, 0x7c, 0x45, 0x39, 0x15, 0xbf, 0xfc, 0x7c, 0x64, 0x7f, 0xc3, 0x71,
	0x9a, 0x12, 0x7e, 0xb2, 0x32, 0x9e, 0x06, 0xb6, 0xc8, 0x70, 0x96, 0x2b, 0x83, 0xf5, 0xa3, 0x01,
	0xb6, 0x6e, 0x72, 0x82, 0x33, 0x72, 0x4f, 0xb6, 0xfd, 0x40, 0x10, 0x7e, 0x0b, 0x67, 0xf8, 0x53,
	0xe6, 0x0f, 0xc9, 0xc3, 0x9c, 0x88, 0x0c, 0x6e, 0x83, 0x7a, 0x90, 0x8b, 0x8c, 0x25, 0x84, 0x7b,
	0x51, 0xd8, 0x30, 0x5a, 0x46, 0x7b, 0xa3, 0x5f, 0x7b, 0xee, 0x9a, 0x43, 0xa0, 0xef, 0xef, 0x86,
	0xf0, 0x33, 0x50, 0x3b, 0x64, 0x7e, 0xc3, 0x6c, 0x19, 0xed, 0xfa, 0xce, 0x75, 0x54, 0xb6, 0xc1,
	0xd9, 0x9c, 0xd1, 0x6a, 0x42, 0x09, 0x3a, 0x85, 0xb1, 0x76, 0x41, 0xab, 0xbc, 0x2c, 0x91, 0x32,
	0x2a, 0x08, 0xbc, 0x0a, 0x5e, 0xd5, 0x68, 0x1e, 0xc5, 0x09, 0x91, 0x95, 0x0d, 0x5f, 0xd1, 0x97,
	0x5f, 0xe0, 0x84, 0x58, 0xdf, 0x82, 0xcb, 0xbb, 0x24, 0x2b, 0x6f, 0xee, 0xcb, 0xb5, 0x20, 0xfd,
	0xee, 0x73, 0xd7, 0xfc, 0xd3, 0x7d, 0x1f, 0xbc, 0x37, 0x2f, 0x5e, 0x9d, 0xd2, 0x48, 0xa0, 0x80,
	0x25, 0xf6, 0x1a, 0xc8, 0x95, 0xcc, 0xc3, 0x9c, 0xbe, 0x8c, 0xcc, 0x4f, 0x4d, 0xf0, 0x8e, 0x1b,
	0x86, 0xab, 0x7e, 0xf7, 0x66, 0xbc, 0xf9, 0xdf, 0x8b, 0x80, 0x7b, 0xe0, 0x4d, 0x42, 0xb1, 0x1f,
	0x13, 0x2f, 0xc5, 0x3c, 0x8b, 0x70, 0xec, 0x8d, 0x70, 0x14, 0xe7, 0x9c, 0x28, 0x8a, 0x34, 0x35,
	0x45, 0x34, 0x67, 0x51, 0x9f, 0xb1, 0x78, 0x1f, 0xc7, 0x39, 0x19, 0xbe, 0x21, 0x23, 0xf7, 0x64,
	0xe0, 0x1d, 0x19, 0x07, 0x7d, 0x00, 0xe6, 0xc4, 0x6f, 0xd4, 0x5a, 0xb5, 0x76, 0x7d, 0xa7, 0x87,
	0xaa, 0xa4, 0x02, 0x9d, 0x32, 0x06, 0xc5, 0xe2, 0x39, 0xaa, 0xf5, 0xcc, 0x00, 0x97, 0x4e, 0x09,
	0x80, 0x7d, 0x70, 0x3e, 0x28, 0x78, 0x59, 0x0c, 0xaa, 0xbe, 0xd3, 0x2e, 0xcd, 0x2f, 0x85, 0x08,
	0x69, 0x94, 0xc1, 0xb9, 0xa1, 0x8a, 0x9c, 0x62, 0x70, 0x92, 0xb0, 0x89, 0x9e, 0xc4, 0x3f, 0xc2,
	0x90, 0x91, 0x70, 0x0b, 0x00, 0x79, 0xf2, 0x70, 0x1c, 0x37, 0x6a, 0x2d, 0xa3, 0x7d, 0x61, 0x70,
	0x6e, 0xb8, 0x21, 0xef, 0xdc, 0x38, 0xee, 0xd7, 0xc1, 0xc6, 0xac, 0x2d, 0xeb, 0x21, 0x78, 0xb7,
	0x92, 0x0f, 0xea, 0x4f, 0x75, 0x07, 0x6c, 0x9e, 0xd8, 0x97, 0x47, 0x38, 0x67, 0x5c, 0xf5, 0x0b,
	0x75, 0xad, 0x3c, 0x0d, 0xd0, 0xfd, 0x42, 0x49, 0x86, 0xaf, 0xa7, 0x4b, 0x7b, 0xba, 0x3d, 0x75,
	0xdf, 0x79, 0x7a, 0x01, 0x5c, 0x5c, 0x4d, 0x78, 0x5f, 0x2e, 0x05, 0xfe, 0x61, 0x80, 0x46, 0xd9,
	0xff, 0x1b, 0xba, 0xd5, 0x3b, 0xad, 0x90, 0xac, 0x66, 0xff, 0xdf, 0x40, 0xc8, 0x49, 0x58, 0x0f,
	0x8e, 0xdd, 0xd7, 0x16, 0x74, 0xef, 0xda, 0x21, 0xf3, 0xbf, 0x7f, 0xf6, 0xdb, 0x0f, 0x66, 0xcf,
	0xfa, 0xa8, 0x78, 0x68, 0x94, 0x49, 0xd8, 0x8f, 0x16, 0xbc, 0x7a, 0x9d, 0xc7, 0xfa, 0x01, 0x59,
	0xc0, 0x14, 0x8e, 0x5c, 0xbd, 0x63, 0x74, 0xe0, 0xaf, 0x06, 0xd8, 0x5c, 0xab, 0x48, 0xf0, 0x46,
	0x75, 0xd1, 0xa7, 0x49, 0x59, 0xf3, 0x6c, 0xa2, 0x6b, 0x7d, 0x7e, 0xec, 0x2e, 0x6b, 0x40, 0xd1,
	0xe5, 0x87, 0xf0, 0xfa, 0xb4, 0xcb, 0x47, 0x4b, 0x96, 0xde, 0xbc, 0xe9, 0xce, 0xba, 0x36, 0xed,
	0xce, 0x63, 0xf8, 0x93, 0x09, 0xb6, 0x2a, 0xc8, 0x06, 0x07, 0xd5, 0x9d, 0xbe, 0x98, 0x7e, 0x35,
	0xef, 0xfe, 0x07, 0x48, 0x6a, 0xdf, 0xd1, 0xb1, 0xdb, 0x58, 0xea, 0xf6, 0xda, 0x5c, 0x17, 0x8a,
	0x91, 0x0c, 0xac, 0x9b, 0x67, 0x1a, 0x89, 0x83, 0xc3, 0x70, 0x9e, 0x6f, 0xca, 0x81, 0xbf, 0x0c,
	0xb0, 0xb9, 0xf6, 0x6d, 0x78, 0x11, 0x0e, 0x9c, 0xf6, 0xa8, 0x34, 0xdf, 0xd6, 0xf1, 0x0b, 0x9f,
	0x0b, 0x68, 0x96, 0xdc, 0x7a, 0x62, 0x1c, 0xb9, 0x08, 0x6c, 0x9e, 0x54, 0xde, 0xdb, 0x49, 0x9a,
	0x7d, 0x07, 0xd7, 0x5f, 0xaf, 0x65, 0xc7, 0x0d, 0xeb, 0xe3, 0xb3, 0x8d, 0x82, 0xe7, 0xd4, 0x31,
	0x3a, 0xcd, 0x4b, 0x47, 0x6e, 0xa3, 0xec, 0x69, 0xe9, 0x3f, 0x31, 0xc1, 0x76, 0xc0, 0x92, 0xca,
	0x49, 0xf4, 0xaf, 0x94, 0x0a, 0xcc, 0xde, 0xb4, 0x85, 0x3d, 0xe3, 0xab, 0x81, 0xc2, 0x18, 0xb3,
	0x18, 0xd3, 0x31, 0x62, 0x7c, 0x6c, 0x8f, 0x09, 0x2d, 0x1a, 0xb4, 0xe7, 0x59, 0xcb, 0xbf, 0x56,
	0x3f, 0xd1, 0x87, 0x9f, 0xcd, 0xda, 0xae, 0xeb, 0xfe, 0x62, 0xb6, 0x76, 0x25, 0xa0, 0x1b, 0x0a,
	0x24, 0x8f, 0xd3, 0xd3, 0x7e, 0x17, 0xa9, 0xc4, 0xe2, 0x48, 0xbb, 0x1c, 0xb8, 0xa1, 0x38, 0x98,
	0xb9, 0x1c, 0xec, 0x77, 0x0f, 0xb4, 0xcb, 0xef, 0xe6, 0xb6, 0xbc, 0x77, 0x1c, 0x37, 0x14, 0x8e,
	0x33, 0x73, 0x72, 0x9c, 0xfd, 0xae, 0xe3, 0x68, 0x37, 0xff, 0x7c, 0x51, 0x67, 0xf7, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x32, 0xc0, 0xf9, 0xbb, 0x54, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OfflineUserDataJobServiceClient is the client API for OfflineUserDataJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OfflineUserDataJobServiceClient interface {
	// Creates an offline user data job.
	CreateOfflineUserDataJob(ctx context.Context, in *CreateOfflineUserDataJobRequest, opts ...grpc.CallOption) (*CreateOfflineUserDataJobResponse, error)
	// Returns the offline user data job.
	GetOfflineUserDataJob(ctx context.Context, in *GetOfflineUserDataJobRequest, opts ...grpc.CallOption) (*resources.OfflineUserDataJob, error)
	// Adds operations to the offline user data job.
	AddOfflineUserDataJobOperations(ctx context.Context, in *AddOfflineUserDataJobOperationsRequest, opts ...grpc.CallOption) (*AddOfflineUserDataJobOperationsResponse, error)
	// Runs the offline user data job.
	//
	// When finished, the long running operation will contain the processing
	// result or failure information, if any.
	RunOfflineUserDataJob(ctx context.Context, in *RunOfflineUserDataJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type offlineUserDataJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOfflineUserDataJobServiceClient(cc grpc.ClientConnInterface) OfflineUserDataJobServiceClient {
	return &offlineUserDataJobServiceClient{cc}
}

func (c *offlineUserDataJobServiceClient) CreateOfflineUserDataJob(ctx context.Context, in *CreateOfflineUserDataJobRequest, opts ...grpc.CallOption) (*CreateOfflineUserDataJobResponse, error) {
	out := new(CreateOfflineUserDataJobResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.OfflineUserDataJobService/CreateOfflineUserDataJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offlineUserDataJobServiceClient) GetOfflineUserDataJob(ctx context.Context, in *GetOfflineUserDataJobRequest, opts ...grpc.CallOption) (*resources.OfflineUserDataJob, error) {
	out := new(resources.OfflineUserDataJob)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.OfflineUserDataJobService/GetOfflineUserDataJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offlineUserDataJobServiceClient) AddOfflineUserDataJobOperations(ctx context.Context, in *AddOfflineUserDataJobOperationsRequest, opts ...grpc.CallOption) (*AddOfflineUserDataJobOperationsResponse, error) {
	out := new(AddOfflineUserDataJobOperationsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.OfflineUserDataJobService/AddOfflineUserDataJobOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *offlineUserDataJobServiceClient) RunOfflineUserDataJob(ctx context.Context, in *RunOfflineUserDataJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.OfflineUserDataJobService/RunOfflineUserDataJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OfflineUserDataJobServiceServer is the server API for OfflineUserDataJobService service.
type OfflineUserDataJobServiceServer interface {
	// Creates an offline user data job.
	CreateOfflineUserDataJob(context.Context, *CreateOfflineUserDataJobRequest) (*CreateOfflineUserDataJobResponse, error)
	// Returns the offline user data job.
	GetOfflineUserDataJob(context.Context, *GetOfflineUserDataJobRequest) (*resources.OfflineUserDataJob, error)
	// Adds operations to the offline user data job.
	AddOfflineUserDataJobOperations(context.Context, *AddOfflineUserDataJobOperationsRequest) (*AddOfflineUserDataJobOperationsResponse, error)
	// Runs the offline user data job.
	//
	// When finished, the long running operation will contain the processing
	// result or failure information, if any.
	RunOfflineUserDataJob(context.Context, *RunOfflineUserDataJobRequest) (*longrunning.Operation, error)
}

// UnimplementedOfflineUserDataJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOfflineUserDataJobServiceServer struct {
}

func (*UnimplementedOfflineUserDataJobServiceServer) CreateOfflineUserDataJob(ctx context.Context, req *CreateOfflineUserDataJobRequest) (*CreateOfflineUserDataJobResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method CreateOfflineUserDataJob not implemented")
}
func (*UnimplementedOfflineUserDataJobServiceServer) GetOfflineUserDataJob(ctx context.Context, req *GetOfflineUserDataJobRequest) (*resources.OfflineUserDataJob, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetOfflineUserDataJob not implemented")
}
func (*UnimplementedOfflineUserDataJobServiceServer) AddOfflineUserDataJobOperations(ctx context.Context, req *AddOfflineUserDataJobOperationsRequest) (*AddOfflineUserDataJobOperationsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method AddOfflineUserDataJobOperations not implemented")
}
func (*UnimplementedOfflineUserDataJobServiceServer) RunOfflineUserDataJob(ctx context.Context, req *RunOfflineUserDataJobRequest) (*longrunning.Operation, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method RunOfflineUserDataJob not implemented")
}

func RegisterOfflineUserDataJobServiceServer(s *grpc.Server, srv OfflineUserDataJobServiceServer) {
	s.RegisterService(&_OfflineUserDataJobService_serviceDesc, srv)
}

func _OfflineUserDataJobService_CreateOfflineUserDataJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfflineUserDataJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineUserDataJobServiceServer).CreateOfflineUserDataJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.OfflineUserDataJobService/CreateOfflineUserDataJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineUserDataJobServiceServer).CreateOfflineUserDataJob(ctx, req.(*CreateOfflineUserDataJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfflineUserDataJobService_GetOfflineUserDataJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfflineUserDataJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineUserDataJobServiceServer).GetOfflineUserDataJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.OfflineUserDataJobService/GetOfflineUserDataJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineUserDataJobServiceServer).GetOfflineUserDataJob(ctx, req.(*GetOfflineUserDataJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfflineUserDataJobService_AddOfflineUserDataJobOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOfflineUserDataJobOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineUserDataJobServiceServer).AddOfflineUserDataJobOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.OfflineUserDataJobService/AddOfflineUserDataJobOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineUserDataJobServiceServer).AddOfflineUserDataJobOperations(ctx, req.(*AddOfflineUserDataJobOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OfflineUserDataJobService_RunOfflineUserDataJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunOfflineUserDataJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OfflineUserDataJobServiceServer).RunOfflineUserDataJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.OfflineUserDataJobService/RunOfflineUserDataJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OfflineUserDataJobServiceServer).RunOfflineUserDataJob(ctx, req.(*RunOfflineUserDataJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OfflineUserDataJobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.OfflineUserDataJobService",
	HandlerType: (*OfflineUserDataJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOfflineUserDataJob",
			Handler:    _OfflineUserDataJobService_CreateOfflineUserDataJob_Handler,
		},
		{
			MethodName: "GetOfflineUserDataJob",
			Handler:    _OfflineUserDataJobService_GetOfflineUserDataJob_Handler,
		},
		{
			MethodName: "AddOfflineUserDataJobOperations",
			Handler:    _OfflineUserDataJobService_AddOfflineUserDataJobOperations_Handler,
		},
		{
			MethodName: "RunOfflineUserDataJob",
			Handler:    _OfflineUserDataJobService_RunOfflineUserDataJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/offline_user_data_job_service.proto",
}
