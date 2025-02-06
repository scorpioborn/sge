// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/legacy/subaccount/v1beta/query.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryParamsRequest is the request type for the Query/Params RPC method
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_630ad87ee8983ecc, []int{0}
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

// QueryParamsResponse is the response type for the Query/Params RPC method
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_630ad87ee8983ecc, []int{1}
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

// QuerySubaccountRequest is the request type for the Query/Subaccount RPC
type QuerySubaccountRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QuerySubaccountRequest) Reset()         { *m = QuerySubaccountRequest{} }
func (m *QuerySubaccountRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySubaccountRequest) ProtoMessage()    {}
func (*QuerySubaccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_630ad87ee8983ecc, []int{2}
}
func (m *QuerySubaccountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubaccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubaccountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubaccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubaccountRequest.Merge(m, src)
}
func (m *QuerySubaccountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubaccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubaccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubaccountRequest proto.InternalMessageInfo

func (m *QuerySubaccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// QuerySubaccountResponse is the response type for the Query/Subaccount RPC
type QuerySubaccountResponse struct {
	Address                     string                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance                     AccountSummary        `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance"`
	LockedBalance               []LockedBalance       `protobuf:"bytes,3,rep,name=locked_balance,json=lockedBalance,proto3" json:"locked_balance"`
	UnlockedBalance             []LockedBalance       `protobuf:"bytes,4,rep,name=unlocked_balance,json=unlockedBalance,proto3" json:"unlocked_balance"`
	WithdrawableUnlockedBalance cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=withdrawable_unlocked_balance,json=withdrawableUnlockedBalance,proto3,customtype=cosmossdk.io/math.Int" json:"withdrawable_unlocked_balance"`
}

func (m *QuerySubaccountResponse) Reset()         { *m = QuerySubaccountResponse{} }
func (m *QuerySubaccountResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySubaccountResponse) ProtoMessage()    {}
func (*QuerySubaccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_630ad87ee8983ecc, []int{3}
}
func (m *QuerySubaccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubaccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubaccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubaccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubaccountResponse.Merge(m, src)
}
func (m *QuerySubaccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubaccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubaccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubaccountResponse proto.InternalMessageInfo

func (m *QuerySubaccountResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QuerySubaccountResponse) GetBalance() AccountSummary {
	if m != nil {
		return m.Balance
	}
	return AccountSummary{}
}

func (m *QuerySubaccountResponse) GetLockedBalance() []LockedBalance {
	if m != nil {
		return m.LockedBalance
	}
	return nil
}

func (m *QuerySubaccountResponse) GetUnlockedBalance() []LockedBalance {
	if m != nil {
		return m.UnlockedBalance
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "sge.legacy.subaccount.v1beta.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "sge.legacy.subaccount.v1beta.QueryParamsResponse")
	proto.RegisterType((*QuerySubaccountRequest)(nil), "sge.legacy.subaccount.v1beta.QuerySubaccountRequest")
	proto.RegisterType((*QuerySubaccountResponse)(nil), "sge.legacy.subaccount.v1beta.QuerySubaccountResponse")
}

func init() {
	proto.RegisterFile("sge/legacy/subaccount/v1beta/query.proto", fileDescriptor_630ad87ee8983ecc)
}

var fileDescriptor_630ad87ee8983ecc = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0xa4, 0x4d, 0xf5, 0x6d, 0xf5, 0x01, 0x5a, 0x42, 0x09, 0xa1, 0x75, 0x2b, 0xab,
	0x42, 0x01, 0x5a, 0x9b, 0x9a, 0xf2, 0x00, 0xf5, 0x0d, 0xa9, 0x42, 0x90, 0x08, 0x09, 0x10, 0x52,
	0xb4, 0xb6, 0x57, 0x8e, 0x15, 0x7b, 0xd7, 0xf5, 0xae, 0x09, 0x11, 0xe2, 0xc2, 0x1d, 0x09, 0x89,
	0x87, 0xe0, 0x05, 0x10, 0x77, 0x6e, 0x3d, 0x56, 0x70, 0x41, 0x1c, 0x2a, 0x94, 0xf0, 0x20, 0xc8,
	0xbb, 0xeb, 0x36, 0x49, 0x2b, 0x8b, 0x72, 0xf3, 0x8e, 0xff, 0xff, 0xdf, 0xce, 0xcc, 0xce, 0x80,
	0x36, 0x0b, 0xb0, 0x15, 0xe1, 0x00, 0x79, 0x23, 0x8b, 0x65, 0x2e, 0xf2, 0x3c, 0x9a, 0x11, 0x6e,
	0xbd, 0xda, 0x71, 0x31, 0x47, 0xd6, 0x41, 0x86, 0xd3, 0x91, 0x99, 0xa4, 0x94, 0x53, 0xb8, 0xca,
	0x02, 0x6c, 0x4a, 0xa5, 0x79, 0xaa, 0x34, 0xa5, 0xb2, 0x75, 0xc3, 0xa3, 0x2c, 0xa6, 0xac, 0x27,
	0xb4, 0x96, 0x3c, 0x48, 0x63, 0xab, 0x11, 0xd0, 0x80, 0xca, 0x78, 0xfe, 0xa5, 0xa2, 0xab, 0x01,
	0xa5, 0x41, 0x84, 0x2d, 0x94, 0x84, 0x16, 0x22, 0x84, 0x72, 0xc4, 0x43, 0x4a, 0x0a, 0xcf, 0x9d,
	0xd2, 0xb4, 0x5c, 0x14, 0x21, 0xe2, 0x61, 0xa5, 0xbd, 0x5d, 0xaa, 0x4d, 0x50, 0x8a, 0x62, 0x85,
	0x35, 0x1a, 0x00, 0x3e, 0xc9, 0x4b, 0x7a, 0x2c, 0x82, 0x1d, 0x7c, 0x90, 0x61, 0xc6, 0x8d, 0xe7,
	0xe0, 0xea, 0x4c, 0x94, 0x25, 0x94, 0x30, 0x0c, 0x1d, 0x50, 0x97, 0xe6, 0xa6, 0xb6, 0xa1, 0xb5,
	0x97, 0xed, 0x4d, 0xb3, 0xac, 0x03, 0xa6, 0x74, 0x3b, 0x0b, 0x87, 0xc7, 0xeb, 0x95, 0x8e, 0x72,
	0x1a, 0x36, 0x58, 0x11, 0xe8, 0xee, 0x89, 0x5c, 0x5d, 0x0a, 0x9b, 0x60, 0x09, 0xf9, 0x7e, 0x8a,
	0x99, 0xc4, 0xff, 0xd7, 0x29, 0x8e, 0xc6, 0x97, 0x1a, 0xb8, 0x7e, 0xc6, 0xa4, 0x72, 0xb2, 0xe7,
	0x5c, 0x4e, 0xf3, 0xdb, 0xe7, 0xed, 0x86, 0x6a, 0xf7, 0x9e, 0xfc, 0xd3, 0xe5, 0x69, 0x48, 0x82,
	0x13, 0x1e, 0xdc, 0x07, 0x4b, 0xaa, 0x61, 0xcd, 0xaa, 0x28, 0x64, 0xab, 0xbc, 0x90, 0x3d, 0x79,
	0xec, 0x66, 0x71, 0x8c, 0xd2, 0x91, 0x2a, 0xa8, 0x40, 0xc0, 0x67, 0xe0, 0x52, 0x44, 0xbd, 0x01,
	0xf6, 0x7b, 0x05, 0xb4, 0xb6, 0x51, 0x6b, 0x2f, 0xdb, 0x77, 0xcb, 0xa1, 0xfb, 0xc2, 0xe3, 0x48,
	0x8b, 0x62, 0xfe, 0x1f, 0x4d, 0x07, 0xe1, 0x4b, 0x70, 0x25, 0x23, 0x73, 0xec, 0x85, 0x7f, 0x65,
	0x5f, 0x2e, 0x50, 0x05, 0x1d, 0x81, 0xb5, 0x61, 0xc8, 0xfb, 0x7e, 0x8a, 0x86, 0xc8, 0x8d, 0x70,
	0xef, 0xcc, 0x55, 0x8b, 0xa2, 0x9f, 0x6b, 0xb9, 0xfb, 0xe7, 0xf1, 0xfa, 0x35, 0xd9, 0x53, 0xe6,
	0x0f, 0xcc, 0x90, 0x5a, 0x31, 0xe2, 0x7d, 0xf3, 0x21, 0xe1, 0x9d, 0x9b, 0xd3, 0x8c, 0xa7, 0xb3,
	0x57, 0xd8, 0x5f, 0xab, 0x60, 0x51, 0x3c, 0x1c, 0xfc, 0xa4, 0x01, 0x70, 0xfa, 0x7a, 0x70, 0xb7,
	0x3c, 0xff, 0xf3, 0x27, 0xa4, 0xf5, 0xe0, 0x82, 0x2e, 0x39, 0x22, 0xc6, 0xd6, 0xbb, 0xef, 0xbf,
	0x3f, 0x56, 0x6f, 0xc1, 0x4d, 0x2b, 0xdf, 0x8b, 0xa9, 0x85, 0x98, 0xfa, 0x7c, 0xa3, 0x66, 0xe3,
	0x2d, 0x7c, 0xaf, 0x81, 0xba, 0x9c, 0x5c, 0x78, 0xef, 0x2f, 0xee, 0x9b, 0x59, 0x9c, 0xd6, 0xce,
	0x05, 0x1c, 0x2a, 0x3b, 0x5d, 0x64, 0xd7, 0x84, 0x2b, 0xf3, 0xd9, 0xc9, 0x85, 0x71, 0x1e, 0x1d,
	0x8e, 0x75, 0xed, 0x68, 0xac, 0x6b, 0xbf, 0xc6, 0xba, 0xf6, 0x61, 0xa2, 0x57, 0x8e, 0x26, 0x7a,
	0xe5, 0xc7, 0x44, 0xaf, 0xbc, 0xd8, 0x0d, 0x42, 0xde, 0xcf, 0x5c, 0xd3, 0xa3, 0x71, 0xee, 0xdd,
	0x26, 0x98, 0x0f, 0x69, 0x3a, 0x10, 0x9c, 0xd7, 0xe7, 0xec, 0x3f, 0x1f, 0x25, 0x98, 0xb9, 0x75,
	0xb1, 0xf8, 0xf7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xc1, 0xe3, 0x08, 0xe8, 0x04, 0x00,
	0x00,
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
	// Subaccount fetches a subaccount given the owner.
	Subaccount(ctx context.Context, in *QuerySubaccountRequest, opts ...grpc.CallOption) (*QuerySubaccountResponse, error)
	// Params returns the subaccount module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Subaccount(ctx context.Context, in *QuerySubaccountRequest, opts ...grpc.CallOption) (*QuerySubaccountResponse, error) {
	out := new(QuerySubaccountResponse)
	err := c.cc.Invoke(ctx, "/sge.legacy.subaccount.v1beta.Query/Subaccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/sge.legacy.subaccount.v1beta.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Subaccount fetches a subaccount given the owner.
	Subaccount(context.Context, *QuerySubaccountRequest) (*QuerySubaccountResponse, error)
	// Params returns the subaccount module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Subaccount(ctx context.Context, req *QuerySubaccountRequest) (*QuerySubaccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subaccount not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Subaccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubaccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Subaccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sge.legacy.subaccount.v1beta.Query/Subaccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Subaccount(ctx, req.(*QuerySubaccountRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/sge.legacy.subaccount.v1beta.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sge.legacy.subaccount.v1beta.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subaccount",
			Handler:    _Query_Subaccount_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sge/legacy/subaccount/v1beta/query.proto",
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

func (m *QuerySubaccountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubaccountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubaccountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySubaccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubaccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubaccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.WithdrawableUnlockedBalance.Size()
		i -= size
		if _, err := m.WithdrawableUnlockedBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.UnlockedBalance) > 0 {
		for iNdEx := len(m.UnlockedBalance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnlockedBalance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.LockedBalance) > 0 {
		for iNdEx := len(m.LockedBalance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockedBalance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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

func (m *QuerySubaccountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySubaccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovQuery(uint64(l))
	if len(m.LockedBalance) > 0 {
		for _, e := range m.LockedBalance {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if len(m.UnlockedBalance) > 0 {
		for _, e := range m.UnlockedBalance {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	l = m.WithdrawableUnlockedBalance.Size()
	n += 1 + l + sovQuery(uint64(l))
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
func (m *QuerySubaccountRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySubaccountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubaccountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QuerySubaccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySubaccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubaccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
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
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedBalance", wireType)
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
			m.LockedBalance = append(m.LockedBalance, LockedBalance{})
			if err := m.LockedBalance[len(m.LockedBalance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnlockedBalance", wireType)
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
			m.UnlockedBalance = append(m.UnlockedBalance, LockedBalance{})
			if err := m.UnlockedBalance[len(m.UnlockedBalance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawableUnlockedBalance", wireType)
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
			if err := m.WithdrawableUnlockedBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
