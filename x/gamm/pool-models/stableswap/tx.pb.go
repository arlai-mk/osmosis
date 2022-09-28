// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/pool-models/stableswap/tx.proto

package stableswap

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// ===================== MsgCreatePool
type MsgCreateStableswapPool struct {
	Sender               string                                   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	PoolParams           *PoolParams                              `protobuf:"bytes,2,opt,name=pool_params,json=poolParams,proto3" json:"pool_params,omitempty" yaml:"pool_params"`
	InitialPoolLiquidity github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=initial_pool_liquidity,json=initialPoolLiquidity,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_pool_liquidity"`
	FuturePoolGovernor   string                                   `protobuf:"bytes,4,opt,name=future_pool_governor,json=futurePoolGovernor,proto3" json:"future_pool_governor,omitempty" yaml:"future_pool_governor"`
}

func (m *MsgCreateStableswapPool) Reset()         { *m = MsgCreateStableswapPool{} }
func (m *MsgCreateStableswapPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStableswapPool) ProtoMessage()    {}
func (*MsgCreateStableswapPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{0}
}
func (m *MsgCreateStableswapPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStableswapPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStableswapPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStableswapPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStableswapPool.Merge(m, src)
}
func (m *MsgCreateStableswapPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStableswapPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStableswapPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStableswapPool proto.InternalMessageInfo

func (m *MsgCreateStableswapPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgCreateStableswapPool) GetPoolParams() *PoolParams {
	if m != nil {
		return m.PoolParams
	}
	return nil
}

func (m *MsgCreateStableswapPool) GetInitialPoolLiquidity() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialPoolLiquidity
	}
	return nil
}

func (m *MsgCreateStableswapPool) GetFuturePoolGovernor() string {
	if m != nil {
		return m.FuturePoolGovernor
	}
	return ""
}

// Returns a poolID with custom poolName.
type MsgCreateStableswapPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateStableswapPoolResponse) Reset()         { *m = MsgCreateStableswapPoolResponse{} }
func (m *MsgCreateStableswapPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStableswapPoolResponse) ProtoMessage()    {}
func (*MsgCreateStableswapPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{1}
}
func (m *MsgCreateStableswapPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStableswapPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStableswapPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStableswapPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStableswapPoolResponse.Merge(m, src)
}
func (m *MsgCreateStableswapPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStableswapPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStableswapPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStableswapPoolResponse proto.InternalMessageInfo

func (m *MsgCreateStableswapPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

// Sender must be the pool's scaling_factor_governor in order for the tx to
// succeed. Adjusts stableswap scaling factors.
type MsgStableSwapAdjustScalingFactors struct {
	Sender         string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	PoolID         uint64   `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	ScalingFactors []uint64 `protobuf:"varint,3,rep,packed,name=scaling_factors,json=scalingFactors,proto3" json:"scaling_factors,omitempty" yaml:"stableswap_scaling_factor"`
}

func (m *MsgStableSwapAdjustScalingFactors) Reset()         { *m = MsgStableSwapAdjustScalingFactors{} }
func (m *MsgStableSwapAdjustScalingFactors) String() string { return proto.CompactTextString(m) }
func (*MsgStableSwapAdjustScalingFactors) ProtoMessage()    {}
func (*MsgStableSwapAdjustScalingFactors) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{2}
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStableSwapAdjustScalingFactors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStableSwapAdjustScalingFactors.Merge(m, src)
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_Size() int {
	return m.Size()
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStableSwapAdjustScalingFactors.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStableSwapAdjustScalingFactors proto.InternalMessageInfo

func (m *MsgStableSwapAdjustScalingFactors) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgStableSwapAdjustScalingFactors) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *MsgStableSwapAdjustScalingFactors) GetScalingFactors() []uint64 {
	if m != nil {
		return m.ScalingFactors
	}
	return nil
}

type MsgStableSwapAdjustScalingFactorsResponse struct {
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) Reset() {
	*m = MsgStableSwapAdjustScalingFactorsResponse{}
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) String() string {
	return proto.CompactTextString(m)
}
func (*MsgStableSwapAdjustScalingFactorsResponse) ProtoMessage() {}
func (*MsgStableSwapAdjustScalingFactorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{3}
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStableSwapAdjustScalingFactorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStableSwapAdjustScalingFactorsResponse.Merge(m, src)
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStableSwapAdjustScalingFactorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStableSwapAdjustScalingFactorsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateStableswapPool)(nil), "osmosis.gamm.poolmodels.stableswap.v1beta1.MsgCreateStableswapPool")
	proto.RegisterType((*MsgCreateStableswapPoolResponse)(nil), "osmosis.gamm.poolmodels.stableswap.v1beta1.MsgCreateStableswapPoolResponse")
	proto.RegisterType((*MsgStableSwapAdjustScalingFactors)(nil), "osmosis.gamm.poolmodels.stableswap.v1beta1.MsgStableSwapAdjustScalingFactors")
	proto.RegisterType((*MsgStableSwapAdjustScalingFactorsResponse)(nil), "osmosis.gamm.poolmodels.stableswap.v1beta1.MsgStableSwapAdjustScalingFactorsResponse")
}

func init() {
	proto.RegisterFile("osmosis/gamm/pool-models/stableswap/tx.proto", fileDescriptor_46b7c8a0f24de97c)
}

var fileDescriptor_46b7c8a0f24de97c = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4d, 0x6f, 0xd3, 0x3e,
	0x1c, 0xc7, 0x9b, 0xb5, 0xea, 0x5f, 0x7f, 0x57, 0x80, 0x88, 0xaa, 0x51, 0x8a, 0x94, 0x94, 0x70,
	0xe9, 0x80, 0xc6, 0x6b, 0x91, 0x90, 0xe0, 0xb6, 0x14, 0x0d, 0x4d, 0x50, 0x69, 0x4b, 0xc5, 0x05,
	0x0e, 0x95, 0xd3, 0x78, 0xc1, 0x90, 0xc4, 0x21, 0x76, 0xbb, 0xf5, 0xc8, 0x9d, 0x03, 0x2f, 0x03,
	0xf1, 0x42, 0xa6, 0x1d, 0x77, 0xe4, 0x14, 0x50, 0xfb, 0x0e, 0x7a, 0xe3, 0x86, 0x1c, 0xa7, 0x0f,
	0x93, 0xd6, 0x6d, 0x95, 0x76, 0x8a, 0xf3, 0xcb, 0xc7, 0xdf, 0xef, 0xef, 0xc1, 0x31, 0x78, 0x4a,
	0x59, 0x40, 0x19, 0x61, 0xd0, 0x43, 0x41, 0x00, 0x23, 0x4a, 0xfd, 0x46, 0x40, 0x5d, 0xec, 0x33,
	0xc8, 0x38, 0x72, 0x7c, 0xcc, 0x8e, 0x50, 0x04, 0xf9, 0xb1, 0x19, 0xc5, 0x94, 0x53, 0xf5, 0x71,
	0x46, 0x9b, 0x82, 0x36, 0x05, 0x2d, 0x61, 0x73, 0x01, 0x9b, 0xc3, 0xa6, 0x83, 0x39, 0x6a, 0x56,
	0xb5, 0x7e, 0x0a, 0x43, 0x07, 0x31, 0x0c, 0xb3, 0x20, 0xec, 0x53, 0x12, 0x4a, 0xad, 0x6a, 0xd9,
	0xa3, 0x1e, 0x4d, 0x97, 0x50, 0xac, 0xb2, 0xe8, 0x8b, 0xeb, 0xe4, 0xb3, 0x58, 0xf6, 0x04, 0x21,
	0xb7, 0x1a, 0xdf, 0xf2, 0xe0, 0x5e, 0x87, 0x79, 0xed, 0x18, 0x23, 0x8e, 0xbb, 0x73, 0x64, 0x9f,
	0x52, 0x5f, 0xdd, 0x02, 0x45, 0x86, 0x43, 0x17, 0xc7, 0x15, 0xa5, 0xa6, 0xd4, 0xff, 0xb7, 0xee,
	0x4e, 0x13, 0xfd, 0xd6, 0x08, 0x05, 0xfe, 0x4b, 0x43, 0xc6, 0x0d, 0x3b, 0x03, 0x54, 0x0a, 0x4a,
	0x42, 0xb4, 0x17, 0xa1, 0x18, 0x05, 0xac, 0xb2, 0x51, 0x53, 0xea, 0xa5, 0xd6, 0x73, 0xf3, 0xfa,
	0x95, 0x9b, 0xc2, 0x71, 0x3f, 0xdd, 0x6d, 0x6d, 0x4e, 0x13, 0x5d, 0x95, 0x3e, 0x4b, 0xa2, 0x86,
	0x0d, 0xa2, 0x39, 0xa3, 0x7e, 0x55, 0xc0, 0x26, 0x09, 0x09, 0x27, 0xc8, 0x4f, 0xcb, 0xe9, 0xf9,
	0xe4, 0xcb, 0x80, 0xb8, 0x84, 0x8f, 0x2a, 0xf9, 0x5a, 0xbe, 0x5e, 0x6a, 0xdd, 0x37, 0x65, 0x2b,
	0x4d, 0xd1, 0xca, 0xb9, 0x4b, 0x9b, 0x92, 0xd0, 0xda, 0x3e, 0x4d, 0xf4, 0xdc, 0xcf, 0xdf, 0x7a,
	0xdd, 0x23, 0xfc, 0xe3, 0xc0, 0x31, 0xfb, 0x34, 0x80, 0x59, 0xdf, 0xe5, 0xa3, 0xc1, 0xdc, 0xcf,
	0x90, 0x8f, 0x22, 0xcc, 0xd2, 0x0d, 0xcc, 0x2e, 0x67, 0x56, 0x22, 0xc9, 0xb7, 0x33, 0x23, 0xf5,
	0x00, 0x94, 0x0f, 0x07, 0x7c, 0x10, 0x63, 0x99, 0x81, 0x47, 0x87, 0x38, 0x0e, 0x69, 0x5c, 0x29,
	0xa4, 0xdd, 0xd2, 0xa7, 0x89, 0xfe, 0x40, 0x56, 0x71, 0x11, 0x65, 0xd8, 0xaa, 0x0c, 0x0b, 0xcd,
	0xd7, 0xb3, 0xe0, 0x2e, 0xd0, 0x57, 0x4c, 0xc3, 0xc6, 0x2c, 0xa2, 0x21, 0xc3, 0xea, 0x23, 0xf0,
	0x5f, 0x2a, 0x44, 0xdc, 0x74, 0x2c, 0x05, 0x0b, 0x8c, 0x13, 0xbd, 0x28, 0x90, 0xbd, 0x57, 0x76,
	0x51, 0x7c, 0xda, 0x73, 0x8d, 0x13, 0x05, 0x3c, 0xec, 0x30, 0x4f, 0x4a, 0x74, 0x8f, 0x50, 0xb4,
	0xe3, 0x7e, 0x1a, 0x30, 0xde, 0xed, 0x23, 0x9f, 0x84, 0xde, 0x2e, 0xea, 0x73, 0x1a, 0xb3, 0x75,
	0x06, 0xbc, 0xe4, 0xba, 0xb1, 0xca, 0x55, 0x3d, 0x00, 0x77, 0x98, 0x74, 0xe8, 0x1d, 0x4a, 0x8b,
	0x74, 0x18, 0x05, 0xab, 0x2e, 0x3a, 0x3e, 0x4d, 0xf4, 0x5a, 0x26, 0xbe, 0x38, 0x8a, 0xe7, 0x79,
	0xc3, 0xbe, 0xcd, 0xce, 0xa5, 0x68, 0x3c, 0x01, 0x5b, 0x57, 0xd6, 0x31, 0x6b, 0x4d, 0xeb, 0xef,
	0x06, 0xc8, 0x77, 0x98, 0xa7, 0xfe, 0x50, 0x40, 0xf9, 0xc2, 0x13, 0xdd, 0x5e, 0xe7, 0x44, 0xae,
	0x18, 0x44, 0xf5, 0xcd, 0x0d, 0x88, 0xcc, 0xa7, 0x79, 0xa2, 0x00, 0xed, 0x8a, 0x29, 0x75, 0xd6,
	0xf4, 0xbb, 0x5c, 0xae, 0xfa, 0xee, 0x46, 0xe5, 0x66, 0x85, 0x58, 0x1f, 0x4e, 0xc7, 0x9a, 0x72,
	0x36, 0xd6, 0x94, 0x3f, 0x63, 0x4d, 0xf9, 0x3e, 0xd1, 0x72, 0x67, 0x13, 0x2d, 0xf7, 0x6b, 0xa2,
	0xe5, 0xde, 0xef, 0x2c, 0xfd, 0x66, 0x99, 0x75, 0xc3, 0x47, 0x0e, 0x9b, 0xbd, 0xc0, 0x61, 0x73,
	0x1b, 0x1e, 0x5f, 0x76, 0x77, 0x39, 0xc5, 0xf4, 0xb2, 0x7a, 0xf6, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0xdb, 0xb6, 0x17, 0x10, 0x79, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateStableswapPool(ctx context.Context, in *MsgCreateStableswapPool, opts ...grpc.CallOption) (*MsgCreateStableswapPoolResponse, error)
	StableSwapAdjustScalingFactors(ctx context.Context, in *MsgStableSwapAdjustScalingFactors, opts ...grpc.CallOption) (*MsgStableSwapAdjustScalingFactorsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateStableswapPool(ctx context.Context, in *MsgCreateStableswapPool, opts ...grpc.CallOption) (*MsgCreateStableswapPoolResponse, error) {
	out := new(MsgCreateStableswapPoolResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.poolmodels.stableswap.v1beta1.Msg/CreateStableswapPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StableSwapAdjustScalingFactors(ctx context.Context, in *MsgStableSwapAdjustScalingFactors, opts ...grpc.CallOption) (*MsgStableSwapAdjustScalingFactorsResponse, error) {
	out := new(MsgStableSwapAdjustScalingFactorsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.poolmodels.stableswap.v1beta1.Msg/StableSwapAdjustScalingFactors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateStableswapPool(context.Context, *MsgCreateStableswapPool) (*MsgCreateStableswapPoolResponse, error)
	StableSwapAdjustScalingFactors(context.Context, *MsgStableSwapAdjustScalingFactors) (*MsgStableSwapAdjustScalingFactorsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateStableswapPool(ctx context.Context, req *MsgCreateStableswapPool) (*MsgCreateStableswapPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStableswapPool not implemented")
}
func (*UnimplementedMsgServer) StableSwapAdjustScalingFactors(ctx context.Context, req *MsgStableSwapAdjustScalingFactors) (*MsgStableSwapAdjustScalingFactorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StableSwapAdjustScalingFactors not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateStableswapPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateStableswapPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateStableswapPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.poolmodels.stableswap.v1beta1.Msg/CreateStableswapPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateStableswapPool(ctx, req.(*MsgCreateStableswapPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StableSwapAdjustScalingFactors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStableSwapAdjustScalingFactors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StableSwapAdjustScalingFactors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.poolmodels.stableswap.v1beta1.Msg/StableSwapAdjustScalingFactors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StableSwapAdjustScalingFactors(ctx, req.(*MsgStableSwapAdjustScalingFactors))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.gamm.poolmodels.stableswap.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStableswapPool",
			Handler:    _Msg_CreateStableswapPool_Handler,
		},
		{
			MethodName: "StableSwapAdjustScalingFactors",
			Handler:    _Msg_StableSwapAdjustScalingFactors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/gamm/pool-models/stableswap/tx.proto",
}

func (m *MsgCreateStableswapPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStableswapPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStableswapPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FuturePoolGovernor) > 0 {
		i -= len(m.FuturePoolGovernor)
		copy(dAtA[i:], m.FuturePoolGovernor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FuturePoolGovernor)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.InitialPoolLiquidity) > 0 {
		for iNdEx := len(m.InitialPoolLiquidity) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialPoolLiquidity[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolParams != nil {
		{
			size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateStableswapPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStableswapPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStableswapPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgStableSwapAdjustScalingFactors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStableSwapAdjustScalingFactors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStableSwapAdjustScalingFactors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ScalingFactors) > 0 {
		dAtA3 := make([]byte, len(m.ScalingFactors)*10)
		var j2 int
		for _, num := range m.ScalingFactors {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintTx(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x1a
	}
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateStableswapPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolParams != nil {
		l = m.PoolParams.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.InitialPoolLiquidity) > 0 {
		for _, e := range m.InitialPoolLiquidity {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.FuturePoolGovernor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateStableswapPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func (m *MsgStableSwapAdjustScalingFactors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	if len(m.ScalingFactors) > 0 {
		l = 0
		for _, e := range m.ScalingFactors {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	return n
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateStableswapPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateStableswapPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStableswapPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolParams == nil {
				m.PoolParams = &PoolParams{}
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPoolLiquidity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialPoolLiquidity = append(m.InitialPoolLiquidity, types.Coin{})
			if err := m.InitialPoolLiquidity[len(m.InitialPoolLiquidity)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuturePoolGovernor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FuturePoolGovernor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateStableswapPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateStableswapPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStableswapPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStableSwapAdjustScalingFactors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStableSwapAdjustScalingFactors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStableSwapAdjustScalingFactors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ScalingFactors = append(m.ScalingFactors, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ScalingFactors) == 0 {
					m.ScalingFactors = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ScalingFactors = append(m.ScalingFactors, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ScalingFactors", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStableSwapAdjustScalingFactorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStableSwapAdjustScalingFactorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStableSwapAdjustScalingFactorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
