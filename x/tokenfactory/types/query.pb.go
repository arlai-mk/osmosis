// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/tokenfactory/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f22013ad0f72e3f, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f22013ad0f72e3f, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryDenomAuthorityMetadataRequest defines the request structure for the
// DenomAuthorityMetadata gRPC query.
type QueryDenomAuthorityMetadataRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
}

func (m *QueryDenomAuthorityMetadataRequest) Reset()         { *m = QueryDenomAuthorityMetadataRequest{} }
func (m *QueryDenomAuthorityMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomAuthorityMetadataRequest) ProtoMessage()    {}
func (*QueryDenomAuthorityMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f22013ad0f72e3f, []int{2}
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomAuthorityMetadataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomAuthorityMetadataRequest.Merge(m, src)
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomAuthorityMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomAuthorityMetadataRequest proto.InternalMessageInfo

func (m *QueryDenomAuthorityMetadataRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// QueryDenomAuthorityMetadataResponse defines the response structure for the
// DenomAuthorityMetadata gRPC query.
type QueryDenomAuthorityMetadataResponse struct {
	AuthorityMetadata DenomAuthorityMetadata `protobuf:"bytes,1,opt,name=authority_metadata,json=authorityMetadata,proto3" json:"authority_metadata" yaml:"authority_metadata"`
}

func (m *QueryDenomAuthorityMetadataResponse) Reset()         { *m = QueryDenomAuthorityMetadataResponse{} }
func (m *QueryDenomAuthorityMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomAuthorityMetadataResponse) ProtoMessage()    {}
func (*QueryDenomAuthorityMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f22013ad0f72e3f, []int{3}
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomAuthorityMetadataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomAuthorityMetadataResponse.Merge(m, src)
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomAuthorityMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomAuthorityMetadataResponse proto.InternalMessageInfo

func (m *QueryDenomAuthorityMetadataResponse) GetAuthorityMetadata() DenomAuthorityMetadata {
	if m != nil {
		return m.AuthorityMetadata
	}
	return DenomAuthorityMetadata{}
}

// QueryDenomsFromCreatorRequest defines the request structure for the
// DenomsFromCreator gRPC query.
type QueryDenomsFromCreatorRequest struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *QueryDenomsFromCreatorRequest) Reset()         { *m = QueryDenomsFromCreatorRequest{} }
func (m *QueryDenomsFromCreatorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomsFromCreatorRequest) ProtoMessage()    {}
func (*QueryDenomsFromCreatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f22013ad0f72e3f, []int{4}
}
func (m *QueryDenomsFromCreatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomsFromCreatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomsFromCreatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomsFromCreatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomsFromCreatorRequest.Merge(m, src)
}
func (m *QueryDenomsFromCreatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomsFromCreatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomsFromCreatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomsFromCreatorRequest proto.InternalMessageInfo

func (m *QueryDenomsFromCreatorRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// QueryDenomsFromCreatorRequest defines the response structure for the
// DenomsFromCreator gRPC query.
type QueryDenomsFromCreatorResponse struct {
	Denoms []string `protobuf:"bytes,1,rep,name=denoms,proto3" json:"denoms,omitempty" yaml:"denoms"`
}

func (m *QueryDenomsFromCreatorResponse) Reset()         { *m = QueryDenomsFromCreatorResponse{} }
func (m *QueryDenomsFromCreatorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomsFromCreatorResponse) ProtoMessage()    {}
func (*QueryDenomsFromCreatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f22013ad0f72e3f, []int{5}
}
func (m *QueryDenomsFromCreatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomsFromCreatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomsFromCreatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomsFromCreatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomsFromCreatorResponse.Merge(m, src)
}
func (m *QueryDenomsFromCreatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomsFromCreatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomsFromCreatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomsFromCreatorResponse proto.InternalMessageInfo

func (m *QueryDenomsFromCreatorResponse) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "osmosis.tokenfactory.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "osmosis.tokenfactory.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryDenomAuthorityMetadataRequest)(nil), "osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataRequest")
	proto.RegisterType((*QueryDenomAuthorityMetadataResponse)(nil), "osmosis.tokenfactory.v1beta1.QueryDenomAuthorityMetadataResponse")
	proto.RegisterType((*QueryDenomsFromCreatorRequest)(nil), "osmosis.tokenfactory.v1beta1.QueryDenomsFromCreatorRequest")
	proto.RegisterType((*QueryDenomsFromCreatorResponse)(nil), "osmosis.tokenfactory.v1beta1.QueryDenomsFromCreatorResponse")
}

func init() {
	proto.RegisterFile("osmosis/tokenfactory/v1beta1/query.proto", fileDescriptor_6f22013ad0f72e3f)
}

var fileDescriptor_6f22013ad0f72e3f = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xcf, 0x68, 0x1b, 0xe9, 0xf8, 0x07, 0x33, 0x16, 0xd1, 0x50, 0x37, 0x3a, 0x96, 0x92, 0x4a,
	0xdd, 0x31, 0x6d, 0x0f, 0x62, 0x15, 0xcd, 0x56, 0xf4, 0xa0, 0x05, 0xdd, 0x9b, 0x5e, 0xc2, 0x24,
	0x9d, 0x6e, 0x17, 0xb3, 0x3b, 0xdb, 0x9d, 0x49, 0x31, 0x94, 0x5e, 0x3c, 0x78, 0x16, 0x3c, 0xfa,
	0x1d, 0xfc, 0x1c, 0x3d, 0x16, 0x7a, 0xf1, 0xb4, 0x48, 0x52, 0xfc, 0x00, 0xf9, 0x04, 0xb2, 0x33,
	0x93, 0xd8, 0xba, 0x71, 0x89, 0x7a, 0xca, 0x32, 0xef, 0xf7, 0x7e, 0x7f, 0xde, 0x7b, 0x04, 0x56,
	0xb9, 0x08, 0xb8, 0xf0, 0x05, 0x91, 0xfc, 0x1d, 0x0b, 0xb7, 0x68, 0x4b, 0xf2, 0xb8, 0x4b, 0x76,
	0x6b, 0x4d, 0x26, 0x69, 0x8d, 0xec, 0x74, 0x58, 0xdc, 0xb5, 0xa3, 0x98, 0x4b, 0x8e, 0xe6, 0x0c,
	0xd2, 0x3e, 0x89, 0xb4, 0x0d, 0xb2, 0x3c, 0xeb, 0x71, 0x8f, 0x2b, 0x20, 0x49, 0xbf, 0x74, 0x4f,
	0x79, 0xce, 0xe3, 0xdc, 0x6b, 0x33, 0x42, 0x23, 0x9f, 0xd0, 0x30, 0xe4, 0x92, 0x4a, 0x9f, 0x87,
	0xc2, 0x54, 0xef, 0xb4, 0x14, 0x25, 0x69, 0x52, 0xc1, 0xb4, 0xd4, 0x48, 0x38, 0xa2, 0x9e, 0x1f,
	0x2a, 0xb0, 0xc1, 0xae, 0xe6, 0xfa, 0xa4, 0x1d, 0xb9, 0xcd, 0x63, 0x5f, 0x76, 0x37, 0x98, 0xa4,
	0x9b, 0x54, 0x52, 0xd3, 0xb5, 0x98, 0xdb, 0x15, 0xd1, 0x98, 0x06, 0xc6, 0x0c, 0x9e, 0x85, 0xe8,
	0x75, 0x6a, 0xe1, 0x95, 0x7a, 0x74, 0xd9, 0x4e, 0x87, 0x09, 0x89, 0xdf, 0xc0, 0x2b, 0xa7, 0x5e,
	0x45, 0xc4, 0x43, 0xc1, 0x90, 0x03, 0x8b, 0xba, 0xf9, 0x1a, 0xb8, 0x09, 0xaa, 0xe7, 0x97, 0xe7,
	0xed, 0xbc, 0xe1, 0xd8, 0xba, 0xdb, 0x99, 0x3a, 0x48, 0x2a, 0x05, 0xd7, 0x74, 0xe2, 0x97, 0x10,
	0x2b, 0xea, 0xa7, 0x2c, 0xe4, 0x41, 0xfd, 0xf7, 0x00, 0xc6, 0x00, 0x5a, 0x80, 0xd3, 0x9b, 0x29,
	0x40, 0x09, 0xcd, 0x38, 0x97, 0x07, 0x49, 0xe5, 0x42, 0x97, 0x06, 0xed, 0x07, 0x58, 0x3d, 0x63,
	0x57, 0x97, 0xf1, 0x57, 0x00, 0x6f, 0xe7, 0xd2, 0x19, 0xe7, 0x1f, 0x01, 0x44, 0xa3, 0x69, 0x35,
	0x02, 0x53, 0x36, 0x31, 0x56, 0xf3, 0x63, 0x8c, 0xa7, 0x76, 0x6e, 0xa5, 0xb1, 0x06, 0x49, 0xe5,
	0xba, 0xf6, 0x95, 0x65, 0xc7, 0x6e, 0x29, 0xb3, 0x20, 0xbc, 0x01, 0x6f, 0xfc, 0xf2, 0x2b, 0x9e,
	0xc5, 0x3c, 0x58, 0x8f, 0x19, 0x95, 0x3c, 0x1e, 0x26, 0x5f, 0x82, 0xe7, 0x5a, 0xfa, 0xc5, 0x64,
	0x47, 0x83, 0xa4, 0x72, 0x49, 0x6b, 0x98, 0x02, 0x76, 0x87, 0x10, 0xfc, 0x02, 0x5a, 0x7f, 0xa2,
	0x33, 0xc9, 0x17, 0x61, 0x51, 0x8d, 0x2a, 0xdd, 0xd9, 0xd9, 0xea, 0x8c, 0x53, 0x1a, 0x24, 0x95,
	0x8b, 0x27, 0x46, 0x29, 0xb0, 0x6b, 0x00, 0xcb, 0xc7, 0x53, 0x70, 0x5a, 0xb1, 0xa1, 0x2f, 0x00,
	0x16, 0xf5, 0xf6, 0xd0, 0xbd, 0xfc, 0xe1, 0x64, 0x8f, 0xa7, 0x5c, 0xfb, 0x8b, 0x0e, 0x6d, 0x12,
	0x2f, 0x7d, 0x38, 0x3a, 0xfe, 0x7c, 0x66, 0x01, 0xcd, 0x93, 0x09, 0x2e, 0x17, 0xfd, 0x00, 0xf0,
	0xea, 0xf8, 0xa5, 0xa0, 0x27, 0x13, 0x68, 0xe7, 0x5e, 0x5e, 0xb9, 0xfe, 0x1f, 0x0c, 0x26, 0xcd,
	0x73, 0x95, 0xa6, 0x8e, 0x1e, 0xe7, 0xa7, 0xd1, 0x53, 0x27, 0x7b, 0xea, 0x77, 0x9f, 0x64, 0x0f,
	0x08, 0x1d, 0x01, 0x58, 0xca, 0x6c, 0x16, 0xad, 0x4d, 0xea, 0x70, 0xcc, 0x79, 0x95, 0x1f, 0xfe,
	0x5b, 0xb3, 0x49, 0xb6, 0xae, 0x92, 0x3d, 0x42, 0x6b, 0x93, 0x24, 0x6b, 0x6c, 0xc5, 0x3c, 0x68,
	0x98, 0x4b, 0x25, 0x7b, 0xe6, 0x63, 0xdf, 0x71, 0x0f, 0x7a, 0x16, 0x38, 0xec, 0x59, 0xe0, 0x7b,
	0xcf, 0x02, 0x9f, 0xfa, 0x56, 0xe1, 0xb0, 0x6f, 0x15, 0xbe, 0xf5, 0xad, 0xc2, 0xdb, 0xfb, 0x9e,
	0x2f, 0xb7, 0x3b, 0x4d, 0xbb, 0xc5, 0x83, 0xa1, 0xc0, 0xdd, 0x36, 0x6d, 0x8a, 0x91, 0xda, 0x6e,
	0x6d, 0x85, 0xbc, 0x3f, 0xad, 0x29, 0xbb, 0x11, 0x13, 0xcd, 0xa2, 0xfa, 0x37, 0x5b, 0xf9, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x03, 0xd9, 0x34, 0x93, 0xd8, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params defines a gRPC query method that returns the tokenfactory module's
	// parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// DenomAuthorityMetadata defines a gRPC query method for fetching
	// DenomAuthorityMetadata for a particular denom.
	DenomAuthorityMetadata(ctx context.Context, in *QueryDenomAuthorityMetadataRequest, opts ...grpc.CallOption) (*QueryDenomAuthorityMetadataResponse, error)
	// DenomsFromCreator defines a gRPC query method for fetching all
	// denominations created by a specific admin/creator.
	DenomsFromCreator(ctx context.Context, in *QueryDenomsFromCreatorRequest, opts ...grpc.CallOption) (*QueryDenomsFromCreatorResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.tokenfactory.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomAuthorityMetadata(ctx context.Context, in *QueryDenomAuthorityMetadataRequest, opts ...grpc.CallOption) (*QueryDenomAuthorityMetadataResponse, error) {
	out := new(QueryDenomAuthorityMetadataResponse)
	err := c.cc.Invoke(ctx, "/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomsFromCreator(ctx context.Context, in *QueryDenomsFromCreatorRequest, opts ...grpc.CallOption) (*QueryDenomsFromCreatorResponse, error) {
	out := new(QueryDenomsFromCreatorResponse)
	err := c.cc.Invoke(ctx, "/osmosis.tokenfactory.v1beta1.Query/DenomsFromCreator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params defines a gRPC query method that returns the tokenfactory module's
	// parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// DenomAuthorityMetadata defines a gRPC query method for fetching
	// DenomAuthorityMetadata for a particular denom.
	DenomAuthorityMetadata(context.Context, *QueryDenomAuthorityMetadataRequest) (*QueryDenomAuthorityMetadataResponse, error)
	// DenomsFromCreator defines a gRPC query method for fetching all
	// denominations created by a specific admin/creator.
	DenomsFromCreator(context.Context, *QueryDenomsFromCreatorRequest) (*QueryDenomsFromCreatorResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DenomAuthorityMetadata(ctx context.Context, req *QueryDenomAuthorityMetadataRequest) (*QueryDenomAuthorityMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomAuthorityMetadata not implemented")
}
func (*UnimplementedQueryServer) DenomsFromCreator(ctx context.Context, req *QueryDenomsFromCreatorRequest) (*QueryDenomsFromCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomsFromCreator not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.tokenfactory.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomAuthorityMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomAuthorityMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomAuthorityMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomAuthorityMetadata(ctx, req.(*QueryDenomAuthorityMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomsFromCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomsFromCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomsFromCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.tokenfactory.v1beta1.Query/DenomsFromCreator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomsFromCreator(ctx, req.(*QueryDenomsFromCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.tokenfactory.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DenomAuthorityMetadata",
			Handler:    _Query_DenomAuthorityMetadata_Handler,
		},
		{
			MethodName: "DenomsFromCreator",
			Handler:    _Query_DenomsFromCreator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/tokenfactory/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryDenomAuthorityMetadataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomAuthorityMetadataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomAuthorityMetadataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomAuthorityMetadataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomAuthorityMetadataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomAuthorityMetadataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AuthorityMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryDenomsFromCreatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomsFromCreatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomsFromCreatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomsFromCreatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomsFromCreatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomsFromCreatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDenomAuthorityMetadataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomAuthorityMetadataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AuthorityMetadata.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDenomsFromCreatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomsFromCreatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomAuthorityMetadataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomAuthorityMetadataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomAuthorityMetadataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomAuthorityMetadataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomAuthorityMetadataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomAuthorityMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorityMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AuthorityMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomsFromCreatorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomsFromCreatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomsFromCreatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomsFromCreatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomsFromCreatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomsFromCreatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
