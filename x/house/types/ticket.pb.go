// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sgenetwork/sge/house/ticket.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/sge-network/sge/types"
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

// DepositTicketPayload indicates data of the deposit ticket.
type DepositTicketPayload struct {
	// kyc_data contains the details of user kyc.
	KycData types.KycDataPayload `protobuf:"bytes,1,opt,name=kyc_data,json=kycData,proto3" json:"kyc_data"`
	// depositor_address is the account who makes a deposit
	DepositorAddress string `protobuf:"bytes,2,opt,name=depositor_address,json=depositorAddress,proto3" json:"depositor_address,omitempty" yaml:"depositor_address"`
}

func (m *DepositTicketPayload) Reset()         { *m = DepositTicketPayload{} }
func (m *DepositTicketPayload) String() string { return proto.CompactTextString(m) }
func (*DepositTicketPayload) ProtoMessage()    {}
func (*DepositTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_240713aaceffb197, []int{0}
}
func (m *DepositTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositTicketPayload.Merge(m, src)
}
func (m *DepositTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *DepositTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_DepositTicketPayload proto.InternalMessageInfo

func (m *DepositTicketPayload) GetKycData() types.KycDataPayload {
	if m != nil {
		return m.KycData
	}
	return types.KycDataPayload{}
}

func (m *DepositTicketPayload) GetDepositorAddress() string {
	if m != nil {
		return m.DepositorAddress
	}
	return ""
}

// WithdrawTicketPayload indicates data of the withdrawal ticket.
type WithdrawTicketPayload struct {
	// kyc_data contains the details of user kyc.
	KycData types.KycDataPayload `protobuf:"bytes,1,opt,name=kyc_data,json=kycData,proto3" json:"kyc_data"`
	// depositor_address is the account who makes a deposit
	DepositorAddress string `protobuf:"bytes,2,opt,name=depositor_address,json=depositorAddress,proto3" json:"depositor_address,omitempty" yaml:"depositor_address"`
}

func (m *WithdrawTicketPayload) Reset()         { *m = WithdrawTicketPayload{} }
func (m *WithdrawTicketPayload) String() string { return proto.CompactTextString(m) }
func (*WithdrawTicketPayload) ProtoMessage()    {}
func (*WithdrawTicketPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_240713aaceffb197, []int{1}
}
func (m *WithdrawTicketPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WithdrawTicketPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WithdrawTicketPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WithdrawTicketPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawTicketPayload.Merge(m, src)
}
func (m *WithdrawTicketPayload) XXX_Size() int {
	return m.Size()
}
func (m *WithdrawTicketPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawTicketPayload.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawTicketPayload proto.InternalMessageInfo

func (m *WithdrawTicketPayload) GetKycData() types.KycDataPayload {
	if m != nil {
		return m.KycData
	}
	return types.KycDataPayload{}
}

func (m *WithdrawTicketPayload) GetDepositorAddress() string {
	if m != nil {
		return m.DepositorAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*DepositTicketPayload)(nil), "sgenetwork.sge.house.DepositTicketPayload")
	proto.RegisterType((*WithdrawTicketPayload)(nil), "sgenetwork.sge.house.WithdrawTicketPayload")
}

func init() { proto.RegisterFile("sgenetwork/sge/house/ticket.proto", fileDescriptor_240713aaceffb197) }

var fileDescriptor_240713aaceffb197 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4e, 0x4f, 0xcd,
	0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x4e, 0x4f, 0xd5, 0xcf, 0xc8, 0x2f, 0x2d, 0x4e,
	0xd5, 0x2f, 0xc9, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x41,
	0x28, 0xd1, 0x2b, 0x4e, 0x4f, 0xd5, 0x03, 0x2b, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b,
	0xd0, 0x07, 0xb1, 0x20, 0x6a, 0xa5, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0x21, 0x12,
	0x10, 0x0e, 0x54, 0x4a, 0x16, 0xcd, 0xa6, 0x92, 0xca, 0x82, 0x54, 0xfd, 0xec, 0xca, 0x64, 0x88,
	0xb4, 0xd2, 0x7e, 0x46, 0x2e, 0x11, 0x97, 0xd4, 0x82, 0xfc, 0xe2, 0xcc, 0x92, 0x10, 0xb0, 0xed,
	0x01, 0x89, 0x95, 0x39, 0xf9, 0x89, 0x29, 0x42, 0x2e, 0x5c, 0x1c, 0xd9, 0x95, 0xc9, 0xf1, 0x29,
	0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xca, 0x7a, 0x68, 0x2e, 0x02, 0x19,
	0xa5, 0xe7, 0x5d, 0x99, 0xec, 0x92, 0x58, 0x92, 0x08, 0xd5, 0xe6, 0xc4, 0x72, 0xe2, 0x9e, 0x3c,
	0x43, 0x10, 0x7b, 0x36, 0x44, 0x54, 0x28, 0x91, 0x4b, 0x30, 0x05, 0x62, 0x7a, 0x7e, 0x51, 0x7c,
	0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xc9, 0xa7,
	0x7b, 0xf2, 0x12, 0x95, 0x89, 0xb9, 0x39, 0x56, 0x4a, 0x18, 0x4a, 0x94, 0x2e, 0x6d, 0xd1, 0x15,
	0x81, 0x7a, 0xc3, 0x11, 0x22, 0x14, 0x5c, 0x52, 0x94, 0x99, 0x97, 0x1e, 0x24, 0x00, 0x57, 0x0b,
	0x15, 0x57, 0x3a, 0xc0, 0xc8, 0x25, 0x1a, 0x9e, 0x59, 0x92, 0x91, 0x52, 0x94, 0x58, 0x3e, 0x34,
	0xbd, 0xe0, 0xe4, 0x74, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x1a, 0xe9,
	0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xa0, 0xd8, 0xd3, 0x45, 0x8e, 0xc9, 0x0a,
	0x58, 0xaa, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xc7, 0xa7, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0x8c, 0xb3, 0x00, 0x5a, 0x02, 0x00, 0x00,
}

func (m *DepositTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DepositorAddress) > 0 {
		i -= len(m.DepositorAddress)
		copy(dAtA[i:], m.DepositorAddress)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.DepositorAddress)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.KycData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *WithdrawTicketPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WithdrawTicketPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WithdrawTicketPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DepositorAddress) > 0 {
		i -= len(m.DepositorAddress)
		copy(dAtA[i:], m.DepositorAddress)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.DepositorAddress)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.KycData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTicket(dAtA []byte, offset int, v uint64) int {
	offset -= sovTicket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.KycData.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = len(m.DepositorAddress)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	return n
}

func (m *WithdrawTicketPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.KycData.Size()
	n += 1 + l + sovTicket(uint64(l))
	l = len(m.DepositorAddress)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: DepositTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.KycData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func (m *WithdrawTicketPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTicket
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
			return fmt.Errorf("proto: WithdrawTicketPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WithdrawTicketPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KycData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.KycData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
				return ErrInvalidLengthTicket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTicket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTicket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTicket
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
func skipTicket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
					return 0, ErrIntOverflowTicket
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
				return 0, ErrInvalidLengthTicket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTicket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTicket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTicket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTicket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTicket = fmt.Errorf("proto: unexpected end of group")
)
