// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sgenetwork/sge/subaccount/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the subaccount module's genesis state.
type GenesisState struct {
	// params contains the subaccount parameters.
	Params       Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	SubaccountId uint64 `protobuf:"varint,2,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id,omitempty"`
	// subaccounts contains all the subaccounts.
	Subaccounts []GenesisSubaccount `protobuf:"bytes,3,rep,name=subaccounts,proto3" json:"subaccounts"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e56b8808b8d551c, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSubaccountId() uint64 {
	if m != nil {
		return m.SubaccountId
	}
	return 0
}

func (m *GenesisState) GetSubaccounts() []GenesisSubaccount {
	if m != nil {
		return m.Subaccounts
	}
	return nil
}

// GenesisSubaccount defines the genesis subaccount containing owner, address
// and balance information.
type GenesisSubaccount struct {
	// address is the address of the subaccount.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// owner is the owner of the subaccount.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// balance defines the balance status of a subaccount
	Balance AccountSummary `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance"`
	// locked_balances defines the lockup of balances history of a subaccount
	LockedBalances []LockedBalance `protobuf:"bytes,4,rep,name=locked_balances,json=lockedBalances,proto3" json:"locked_balances"`
}

func (m *GenesisSubaccount) Reset()         { *m = GenesisSubaccount{} }
func (m *GenesisSubaccount) String() string { return proto.CompactTextString(m) }
func (*GenesisSubaccount) ProtoMessage()    {}
func (*GenesisSubaccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e56b8808b8d551c, []int{1}
}
func (m *GenesisSubaccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisSubaccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisSubaccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisSubaccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisSubaccount.Merge(m, src)
}
func (m *GenesisSubaccount) XXX_Size() int {
	return m.Size()
}
func (m *GenesisSubaccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisSubaccount.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisSubaccount proto.InternalMessageInfo

func (m *GenesisSubaccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GenesisSubaccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GenesisSubaccount) GetBalance() AccountSummary {
	if m != nil {
		return m.Balance
	}
	return AccountSummary{}
}

func (m *GenesisSubaccount) GetLockedBalances() []LockedBalance {
	if m != nil {
		return m.LockedBalances
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sgenetwork.sge.subaccount.GenesisState")
	proto.RegisterType((*GenesisSubaccount)(nil), "sgenetwork.sge.subaccount.GenesisSubaccount")
}

func init() {
	proto.RegisterFile("sgenetwork/sge/subaccount/genesis.proto", fileDescriptor_4e56b8808b8d551c)
}

var fileDescriptor_4e56b8808b8d551c = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6b, 0xe2, 0x40,
	0x18, 0xc6, 0x13, 0x75, 0x95, 0x1d, 0xdd, 0x5d, 0x76, 0xf0, 0x10, 0x3d, 0x64, 0x5d, 0x0b, 0x6d,
	0x0a, 0x35, 0x01, 0xdb, 0x7b, 0x31, 0x97, 0x22, 0x94, 0x52, 0xb4, 0x50, 0xe8, 0x45, 0x26, 0xc9,
	0x30, 0x0d, 0x26, 0x19, 0xc9, 0x4c, 0xb0, 0x7e, 0x8b, 0xde, 0xfb, 0x35, 0xfa, 0x21, 0xa4, 0x27,
	0xe9, 0xa9, 0xa7, 0x52, 0xf4, 0x8b, 0x14, 0x33, 0x63, 0x13, 0x28, 0xe6, 0x94, 0xcc, 0xf3, 0xfe,
	0x9e, 0xf7, 0x1f, 0x2f, 0x38, 0x62, 0x04, 0x47, 0x98, 0xcf, 0x69, 0x3c, 0xb5, 0x18, 0xc1, 0x16,
	0x4b, 0x1c, 0xe4, 0xba, 0x34, 0x89, 0xb8, 0xb5, 0x0d, 0x30, 0x9f, 0x99, 0xb3, 0x98, 0x72, 0x0a,
	0x5b, 0x19, 0x68, 0x32, 0x82, 0xcd, 0x0c, 0x6c, 0xb7, 0x5c, 0xca, 0x42, 0xca, 0x26, 0x29, 0x68,
	0x89, 0x87, 0x70, 0xb5, 0x9b, 0x84, 0x12, 0x2a, 0xf4, 0xed, 0x9f, 0x54, 0x0b, 0x8a, 0x3a, 0x28,
	0x40, 0x91, 0x8b, 0x25, 0x78, 0xb8, 0x1f, 0x9c, 0xa1, 0x18, 0x85, 0xb2, 0x4c, 0xf7, 0x45, 0x05,
	0x8d, 0x0b, 0xd1, 0xee, 0x98, 0x23, 0x8e, 0xe1, 0x39, 0xa8, 0x0a, 0x40, 0x53, 0x3b, 0xaa, 0x51,
	0xef, 0xff, 0x37, 0xf7, 0xb6, 0x6f, 0x5e, 0xa7, 0xa0, 0x5d, 0x59, 0xbe, 0xff, 0x53, 0x46, 0xd2,
	0x06, 0x0f, 0xc0, 0xaf, 0x0c, 0x99, 0xf8, 0x9e, 0x56, 0xea, 0xa8, 0x46, 0x65, 0xd4, 0xc8, 0xc4,
	0xa1, 0x07, 0x6f, 0x40, 0x3d, 0x7b, 0x33, 0xad, 0xdc, 0x29, 0x1b, 0xf5, 0xfe, 0x49, 0x41, 0xa9,
	0x5d, 0x8f, 0x5f, 0x8a, 0xac, 0x9a, 0x4f, 0xd3, 0x7d, 0x2a, 0x81, 0xbf, 0xdf, 0x40, 0xd8, 0x07,
	0x35, 0xe4, 0x79, 0x31, 0x66, 0x62, 0xa4, 0x9f, 0xb6, 0xf6, 0xfa, 0xdc, 0x6b, 0xca, 0x65, 0x0f,
	0x44, 0x64, 0xcc, 0x63, 0x3f, 0x22, 0xa3, 0x1d, 0x08, 0x4d, 0xf0, 0x83, 0xce, 0x23, 0x1c, 0xa7,
	0xcd, 0x17, 0x39, 0x04, 0x06, 0x87, 0xa0, 0x26, 0xf7, 0xaf, 0x95, 0xd3, 0xb5, 0x1d, 0x17, 0xcc,
	0x32, 0x10, 0xdf, 0x71, 0x12, 0x86, 0x28, 0x5e, 0xc8, 0x41, 0x76, 0x7e, 0x78, 0x0b, 0xfe, 0x04,
	0xd4, 0x9d, 0x62, 0x6f, 0x22, 0x15, 0xa6, 0x55, 0xd2, 0xf5, 0x18, 0x05, 0x29, 0x2f, 0x53, 0x87,
	0x2d, 0x0c, 0x32, 0xe3, 0xef, 0x20, 0x2f, 0x32, 0xfb, 0x6a, 0xb9, 0xd6, 0xd5, 0xd5, 0x5a, 0x57,
	0x3f, 0xd6, 0xba, 0xfa, 0xb8, 0xd1, 0x95, 0xd5, 0x46, 0x57, 0xde, 0x36, 0xba, 0x72, 0x77, 0x46,
	0x7c, 0x7e, 0x9f, 0x38, 0xa6, 0x4b, 0xc3, 0xed, 0xb1, 0xf4, 0xf2, 0x87, 0xf3, 0x60, 0x05, 0x98,
	0x20, 0x77, 0x91, 0xbf, 0x20, 0xbe, 0x98, 0x61, 0xe6, 0x54, 0xd3, 0x0b, 0x3a, 0xfd, 0x0c, 0x00,
	0x00, 0xff, 0xff, 0x31, 0xc9, 0x12, 0xdf, 0x09, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subaccounts) > 0 {
		for iNdEx := len(m.Subaccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subaccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.SubaccountId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SubaccountId))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisSubaccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisSubaccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisSubaccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LockedBalances) > 0 {
		for iNdEx := len(m.LockedBalances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockedBalances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.SubaccountId != 0 {
		n += 1 + sovGenesis(uint64(m.SubaccountId))
	}
	if len(m.Subaccounts) > 0 {
		for _, e := range m.Subaccounts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisSubaccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LockedBalances) > 0 {
		for _, e := range m.LockedBalances {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountId", wireType)
			}
			m.SubaccountId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubaccountId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subaccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subaccounts = append(m.Subaccounts, GenesisSubaccount{})
			if err := m.Subaccounts[len(m.Subaccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisSubaccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisSubaccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisSubaccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedBalances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockedBalances = append(m.LockedBalances, LockedBalance{})
			if err := m.LockedBalances[len(m.LockedBalances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
