// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/strategicreserve/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// MsgInvokeFeeGrant is the message type for granting a feegrant from strategic
// reserve module account into the state.
type MsgInvokeFeeGrant struct {
	// creator is the address of the creator account of the feegrant.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// ticket is the jwt ticket data.
	Ticket string `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (m *MsgInvokeFeeGrant) Reset()         { *m = MsgInvokeFeeGrant{} }
func (m *MsgInvokeFeeGrant) String() string { return proto.CompactTextString(m) }
func (*MsgInvokeFeeGrant) ProtoMessage()    {}
func (*MsgInvokeFeeGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbf4aff262295ba9, []int{0}
}
func (m *MsgInvokeFeeGrant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInvokeFeeGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInvokeFeeGrant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInvokeFeeGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInvokeFeeGrant.Merge(m, src)
}
func (m *MsgInvokeFeeGrant) XXX_Size() int {
	return m.Size()
}
func (m *MsgInvokeFeeGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInvokeFeeGrant.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInvokeFeeGrant proto.InternalMessageInfo

func (m *MsgInvokeFeeGrant) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgInvokeFeeGrant) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

// MsgInvokeFeeGrantResponse is the returning value in the response
// of MsgInvokeFeeGrant request.
type MsgInvokeFeeGrantResponse struct {
	FeeGrant *FeeGrant `protobuf:"bytes,1,opt,name=fee_grant,json=feeGrant,proto3" json:"fee_grant,omitempty"`
}

func (m *MsgInvokeFeeGrantResponse) Reset()         { *m = MsgInvokeFeeGrantResponse{} }
func (m *MsgInvokeFeeGrantResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInvokeFeeGrantResponse) ProtoMessage()    {}
func (*MsgInvokeFeeGrantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbf4aff262295ba9, []int{1}
}
func (m *MsgInvokeFeeGrantResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInvokeFeeGrantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInvokeFeeGrantResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInvokeFeeGrantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInvokeFeeGrantResponse.Merge(m, src)
}
func (m *MsgInvokeFeeGrantResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInvokeFeeGrantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInvokeFeeGrantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInvokeFeeGrantResponse proto.InternalMessageInfo

func (m *MsgInvokeFeeGrantResponse) GetFeeGrant() *FeeGrant {
	if m != nil {
		return m.FeeGrant
	}
	return nil
}

// MsgRevokeFeeGrant is the message type for revoking a feegrant from strategic
// reserve module account from the state.
type MsgRevokeFeeGrant struct {
	// creator is the address of the creator account of the feegrant.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// ticket is the jwt ticket data.
	Ticket string `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (m *MsgRevokeFeeGrant) Reset()         { *m = MsgRevokeFeeGrant{} }
func (m *MsgRevokeFeeGrant) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeFeeGrant) ProtoMessage()    {}
func (*MsgRevokeFeeGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbf4aff262295ba9, []int{2}
}
func (m *MsgRevokeFeeGrant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeFeeGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeFeeGrant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeFeeGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeFeeGrant.Merge(m, src)
}
func (m *MsgRevokeFeeGrant) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeFeeGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeFeeGrant.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeFeeGrant proto.InternalMessageInfo

func (m *MsgRevokeFeeGrant) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRevokeFeeGrant) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

// MsgRevokeFeeGrantResponse is the returning value in the response
// of MsgRevokeFeeGrant request.
type MsgRevokeFeeGrantResponse struct {
	FeeGrant *FeeGrant `protobuf:"bytes,1,opt,name=fee_grant,json=feeGrant,proto3" json:"fee_grant,omitempty"`
}

func (m *MsgRevokeFeeGrantResponse) Reset()         { *m = MsgRevokeFeeGrantResponse{} }
func (m *MsgRevokeFeeGrantResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeFeeGrantResponse) ProtoMessage()    {}
func (*MsgRevokeFeeGrantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbf4aff262295ba9, []int{3}
}
func (m *MsgRevokeFeeGrantResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeFeeGrantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeFeeGrantResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeFeeGrantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeFeeGrantResponse.Merge(m, src)
}
func (m *MsgRevokeFeeGrantResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeFeeGrantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeFeeGrantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeFeeGrantResponse proto.InternalMessageInfo

func (m *MsgRevokeFeeGrantResponse) GetFeeGrant() *FeeGrant {
	if m != nil {
		return m.FeeGrant
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInvokeFeeGrant)(nil), "sgenetwork.sge.strategicreserve.MsgInvokeFeeGrant")
	proto.RegisterType((*MsgInvokeFeeGrantResponse)(nil), "sgenetwork.sge.strategicreserve.MsgInvokeFeeGrantResponse")
	proto.RegisterType((*MsgRevokeFeeGrant)(nil), "sgenetwork.sge.strategicreserve.MsgRevokeFeeGrant")
	proto.RegisterType((*MsgRevokeFeeGrantResponse)(nil), "sgenetwork.sge.strategicreserve.MsgRevokeFeeGrantResponse")
}

func init() { proto.RegisterFile("sge/strategicreserve/tx.proto", fileDescriptor_bbf4aff262295ba9) }

var fileDescriptor_bbf4aff262295ba9 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4b, 0xc3, 0x40,
	0x14, 0xc6, 0x73, 0x0a, 0xd5, 0x9e, 0x20, 0x18, 0x44, 0x62, 0xc0, 0x53, 0xea, 0xa2, 0x83, 0x17,
	0x48, 0x37, 0x47, 0xc1, 0x8a, 0x43, 0x96, 0x8c, 0x2e, 0x92, 0x86, 0x2f, 0x67, 0x08, 0xe6, 0xc2,
	0xdd, 0x59, 0xeb, 0xe6, 0xe8, 0xe8, 0x9f, 0xe5, 0xd8, 0xd1, 0x51, 0x92, 0x7f, 0x44, 0x92, 0xa6,
	0x82, 0x6d, 0xa1, 0x14, 0xba, 0xbd, 0xc7, 0xfb, 0xde, 0xef, 0x7d, 0xef, 0xf1, 0xe8, 0x89, 0x16,
	0xf0, 0xb4, 0x51, 0x91, 0x81, 0x48, 0x63, 0x05, 0x0d, 0x35, 0x82, 0x67, 0xc6, 0xbc, 0x50, 0xd2,
	0x48, 0xfb, 0x54, 0x0b, 0xe4, 0x30, 0xaf, 0x52, 0x65, 0x5c, 0x0b, 0xf0, 0x79, 0xa5, 0x7b, 0x28,
	0xa4, 0x90, 0x8d, 0xd6, 0xab, 0xa3, 0x69, 0x9b, 0x7b, 0xbe, 0x94, 0x9a, 0x00, 0x42, 0x45, 0xb9,
	0x99, 0x8a, 0x7a, 0xb7, 0xf4, 0x20, 0xd0, 0xe2, 0x3e, 0x1f, 0xc9, 0x0c, 0x03, 0xe0, 0xae, 0x2e,
	0xd9, 0x0e, 0xdd, 0x89, 0x15, 0x22, 0x23, 0x95, 0x43, 0xce, 0xc8, 0x45, 0x37, 0x9c, 0xa5, 0xf6,
	0x11, 0xed, 0x98, 0x34, 0xce, 0x60, 0x9c, 0xad, 0xa6, 0xd0, 0x66, 0xbd, 0x98, 0x1e, 0x2f, 0x60,
	0x42, 0xe8, 0x42, 0xe6, 0x1a, 0xf6, 0x80, 0x76, 0x13, 0xe0, 0xb1, 0x19, 0xdb, 0x00, 0xf7, 0xfc,
	0x4b, 0xbe, 0x62, 0x27, 0xfe, 0x47, 0xd9, 0x4d, 0xda, 0xa8, 0xf5, 0x1a, 0x62, 0x23, 0x5e, 0xff,
	0x63, 0x36, 0xed, 0xd5, 0xff, 0x20, 0x74, 0x3b, 0xd0, 0xc2, 0x7e, 0x27, 0x74, 0x7f, 0xee, 0xba,
	0xfe, 0x4a, 0xde, 0xc2, 0x29, 0xdd, 0xeb, 0xf5, 0x7b, 0x66, 0x2b, 0xdd, 0x04, 0x5f, 0x25, 0x23,
	0x93, 0x92, 0x91, 0x9f, 0x92, 0x91, 0xcf, 0x8a, 0x59, 0x93, 0x8a, 0x59, 0xdf, 0x15, 0xb3, 0x1e,
	0xfa, 0x22, 0x35, 0x4f, 0x2f, 0x43, 0x1e, 0xcb, 0x67, 0x4f, 0x0b, 0x5c, 0xb5, 0x03, 0xea, 0xd8,
	0x1b, 0x2f, 0x79, 0xc8, 0xb7, 0x02, 0x7a, 0xd8, 0x69, 0x1e, 0xa7, 0xff, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x60, 0x0b, 0x51, 0xa3, 0xb5, 0x02, 0x00, 0x00,
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
	// InvokeFeeGrant defines a method to add a feegrant with the given data.
	InvokeFeeGrant(ctx context.Context, in *MsgInvokeFeeGrant, opts ...grpc.CallOption) (*MsgInvokeFeeGrantResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InvokeFeeGrant(ctx context.Context, in *MsgInvokeFeeGrant, opts ...grpc.CallOption) (*MsgInvokeFeeGrantResponse, error) {
	out := new(MsgInvokeFeeGrantResponse)
	err := c.cc.Invoke(ctx, "/sgenetwork.sge.strategicreserve.Msg/InvokeFeeGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// InvokeFeeGrant defines a method to add a feegrant with the given data.
	InvokeFeeGrant(context.Context, *MsgInvokeFeeGrant) (*MsgInvokeFeeGrantResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InvokeFeeGrant(ctx context.Context, req *MsgInvokeFeeGrant) (*MsgInvokeFeeGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvokeFeeGrant not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InvokeFeeGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInvokeFeeGrant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InvokeFeeGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgenetwork.sge.strategicreserve.Msg/InvokeFeeGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InvokeFeeGrant(ctx, req.(*MsgInvokeFeeGrant))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sgenetwork.sge.strategicreserve.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InvokeFeeGrant",
			Handler:    _Msg_InvokeFeeGrant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sge/strategicreserve/tx.proto",
}

func (m *MsgInvokeFeeGrant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInvokeFeeGrant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInvokeFeeGrant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ticket) > 0 {
		i -= len(m.Ticket)
		copy(dAtA[i:], m.Ticket)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ticket)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInvokeFeeGrantResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInvokeFeeGrantResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInvokeFeeGrantResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeGrant != nil {
		{
			size, err := m.FeeGrant.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeFeeGrant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeFeeGrant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeFeeGrant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ticket) > 0 {
		i -= len(m.Ticket)
		copy(dAtA[i:], m.Ticket)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ticket)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeFeeGrantResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeFeeGrantResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeFeeGrantResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeGrant != nil {
		{
			size, err := m.FeeGrant.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgInvokeFeeGrant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Ticket)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInvokeFeeGrantResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FeeGrant != nil {
		l = m.FeeGrant.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRevokeFeeGrant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Ticket)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRevokeFeeGrantResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FeeGrant != nil {
		l = m.FeeGrant.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInvokeFeeGrant) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgInvokeFeeGrant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInvokeFeeGrant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticket", wireType)
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
			m.Ticket = string(dAtA[iNdEx:postIndex])
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
func (m *MsgInvokeFeeGrantResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgInvokeFeeGrantResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInvokeFeeGrantResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeGrant", wireType)
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
			if m.FeeGrant == nil {
				m.FeeGrant = &FeeGrant{}
			}
			if err := m.FeeGrant.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgRevokeFeeGrant) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRevokeFeeGrant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeFeeGrant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticket", wireType)
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
			m.Ticket = string(dAtA[iNdEx:postIndex])
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
func (m *MsgRevokeFeeGrantResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRevokeFeeGrantResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeFeeGrantResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeGrant", wireType)
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
			if m.FeeGrant == nil {
				m.FeeGrant = &FeeGrant{}
			}
			if err := m.FeeGrant.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
