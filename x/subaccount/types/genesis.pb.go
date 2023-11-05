// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/subaccount/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
	return fileDescriptor_021997686c533a34, []int{0}
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
	return fileDescriptor_021997686c533a34, []int{1}
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

func init() { proto.RegisterFile("sge/subaccount/genesis.proto", fileDescriptor_021997686c533a34) }

var fileDescriptor_021997686c533a34 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcf, 0x4e, 0xc2, 0x30,
	0x18, 0x5f, 0x05, 0x21, 0x14, 0xd4, 0xd8, 0x70, 0x98, 0x68, 0x26, 0xe2, 0x65, 0x26, 0xb2, 0x25,
	0xf8, 0x00, 0x46, 0x2e, 0x84, 0xc4, 0x83, 0x19, 0x26, 0x26, 0x5e, 0x48, 0xb7, 0x35, 0x95, 0xc0,
	0x56, 0xb2, 0x76, 0x41, 0xde, 0xc2, 0xc7, 0x22, 0x9e, 0x38, 0x7a, 0x32, 0x06, 0x0e, 0xbe, 0x86,
	0xa1, 0x2d, 0x8e, 0x60, 0xe4, 0xb4, 0x7e, 0xbf, 0xfd, 0xfe, 0xb5, 0xf9, 0xe0, 0x19, 0xa7, 0xc4,
	0xe5, 0xa9, 0x8f, 0x83, 0x80, 0xa5, 0xb1, 0x70, 0x29, 0x89, 0x09, 0x1f, 0x70, 0x67, 0x9c, 0x30,
	0xc1, 0xd0, 0x09, 0x5f, 0xcd, 0x62, 0xc2, 0x92, 0xa1, 0xc3, 0x29, 0x71, 0x32, 0x62, 0xad, 0x4a,
	0x19, 0x65, 0x92, 0xe5, 0xae, 0x4e, 0x4a, 0x50, 0x3b, 0xdd, 0xb2, 0x1b, 0xe3, 0x04, 0x47, 0xda,
	0xad, 0xb6, 0x9d, 0xe5, 0xe3, 0x11, 0x8e, 0x03, 0xa2, 0xfe, 0x36, 0xde, 0x01, 0xac, 0x74, 0x54,
	0x7a, 0x4f, 0x60, 0x41, 0xd0, 0x2d, 0x2c, 0x28, 0xb9, 0x09, 0xea, 0xc0, 0x2e, 0xb7, 0x2e, 0x9c,
	0x7f, 0xdb, 0x38, 0x0f, 0x92, 0xd8, 0xce, 0xcf, 0x3e, 0xcf, 0x0d, 0x4f, 0xcb, 0xd0, 0x25, 0x3c,
	0xc8, 0x28, 0xfd, 0x41, 0x68, 0xee, 0xd5, 0x81, 0x9d, 0xf7, 0x2a, 0x19, 0xd8, 0x0d, 0xd1, 0x23,
	0x2c, 0x67, 0x33, 0x37, 0x73, 0xf5, 0x9c, 0x5d, 0x6e, 0x5d, 0xef, 0x88, 0x5a, 0x77, 0xfc, 0x45,
	0x74, 0xea, 0xa6, 0x4d, 0xe3, 0x1b, 0xc0, 0xe3, 0x3f, 0x44, 0x64, 0xc2, 0x22, 0x0e, 0xc3, 0x84,
	0x70, 0x75, 0xa5, 0x92, 0xb7, 0x1e, 0x51, 0x15, 0xee, 0xb3, 0x49, 0x4c, 0x12, 0x59, 0xb1, 0xe4,
	0xa9, 0x01, 0x75, 0x61, 0x51, 0xbf, 0x91, 0x99, 0x93, 0x4f, 0x70, 0xb5, 0xa3, 0xd7, 0x9d, 0xfa,
	0xf6, 0xd2, 0x28, 0xc2, 0xc9, 0x54, 0x97, 0x5a, 0xeb, 0xd1, 0x13, 0x3c, 0x1a, 0xb1, 0x60, 0x48,
	0xc2, 0xbe, 0x46, 0xb8, 0x99, 0x97, 0x57, 0xb5, 0x77, 0x58, 0xde, 0x4b, 0x45, 0x5b, 0x09, 0xb4,
	0xe3, 0xe1, 0x68, 0x13, 0xe4, 0xed, 0xce, 0x6c, 0x61, 0x81, 0xf9, 0xc2, 0x02, 0x5f, 0x0b, 0x0b,
	0xbc, 0x2d, 0x2d, 0x63, 0xbe, 0xb4, 0x8c, 0x8f, 0xa5, 0x65, 0x3c, 0x37, 0xe9, 0x40, 0xbc, 0xa4,
	0xbe, 0x13, 0xb0, 0xc8, 0xe5, 0x94, 0x34, 0x75, 0xc8, 0xea, 0xec, 0xbe, 0x6e, 0xee, 0x81, 0x98,
	0x8e, 0x09, 0xf7, 0x0b, 0x72, 0x0d, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x12, 0x0e,
	0x76, 0x92, 0x02, 0x00, 0x00,
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
