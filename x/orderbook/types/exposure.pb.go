// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/orderbook/exposure.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// BookOddsExposure represents the exposures taken on odds
type BookOddsExposure struct {
	// book_id is id corresponding to the book
	BookID string `protobuf:"bytes,1,opt,name=book_id,proto3" json:"book_id"`
	// odds_id is odd'd uid
	OddsID           string   `protobuf:"bytes,2,opt,name=odds_id,proto3" json:"odds_id"`
	FulfillmentQueue []uint64 `protobuf:"varint,3,rep,packed,name=fulfillment_queue,json=fulfillmentQueue,proto3" json:"fulfillment_queue,omitempty" yaml:"fulfillment_queue"`
}

func (m *BookOddsExposure) Reset()      { *m = BookOddsExposure{} }
func (*BookOddsExposure) ProtoMessage() {}
func (*BookOddsExposure) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa5f8fec7488c62, []int{0}
}
func (m *BookOddsExposure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BookOddsExposure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BookOddsExposure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BookOddsExposure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookOddsExposure.Merge(m, src)
}
func (m *BookOddsExposure) XXX_Size() int {
	return m.Size()
}
func (m *BookOddsExposure) XXX_DiscardUnknown() {
	xxx_messageInfo_BookOddsExposure.DiscardUnknown(m)
}

var xxx_messageInfo_BookOddsExposure proto.InternalMessageInfo

// ParticipationExposure represents the exposures taken on odds by
// participations
type ParticipationExposure struct {
	// book_id is id corresponding to the book
	BookID string `protobuf:"bytes,1,opt,name=book_id,proto3" json:"book_id"`
	// odds_id is odd's uid
	OddsID string `protobuf:"bytes,2,opt,name=odds_id,proto3" json:"odds_id"`
	// participation_index is the id of initial participation queue
	ParticipationIndex uint64 `protobuf:"varint,3,opt,name=participation_index,json=participationIndex,proto3" json:"participation_index,omitempty" yaml:"participation_index"`
	// exposure is the total exposure taken on given odd
	Exposure github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=exposure,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"exposure" yaml:"exposure"`
	// bet_amount is the total bet amount corresponding to the exposure
	BetAmount   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=bet_amount,json=betAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"bet_amount" yaml:"bet_amount"`
	IsFulfilled bool                                   `protobuf:"varint,6,opt,name=is_fulfilled,json=isFulfilled,proto3" json:"is_fulfilled,omitempty" yaml:"is_fulfilled"`
	// number of current round in queue
	Round uint64 `protobuf:"varint,7,opt,name=round,proto3" json:"round,omitempty" yaml:"round"`
}

func (m *ParticipationExposure) Reset()      { *m = ParticipationExposure{} }
func (*ParticipationExposure) ProtoMessage() {}
func (*ParticipationExposure) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa5f8fec7488c62, []int{1}
}
func (m *ParticipationExposure) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParticipationExposure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParticipationExposure.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParticipationExposure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParticipationExposure.Merge(m, src)
}
func (m *ParticipationExposure) XXX_Size() int {
	return m.Size()
}
func (m *ParticipationExposure) XXX_DiscardUnknown() {
	xxx_messageInfo_ParticipationExposure.DiscardUnknown(m)
}

var xxx_messageInfo_ParticipationExposure proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BookOddsExposure)(nil), "sgenetwork.sge.orderbook.BookOddsExposure")
	proto.RegisterType((*ParticipationExposure)(nil), "sgenetwork.sge.orderbook.ParticipationExposure")
}

func init() { proto.RegisterFile("sge/orderbook/exposure.proto", fileDescriptor_3aa5f8fec7488c62) }

var fileDescriptor_3aa5f8fec7488c62 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xb7, 0x49, 0x9a, 0xb6, 0x47, 0x24, 0x52, 0x97, 0x0a, 0xab, 0xaa, 0x7c, 0x91, 0x87, 0x2a,
	0x03, 0xb5, 0x07, 0xb6, 0x6c, 0x35, 0x50, 0x29, 0x53, 0xc1, 0x23, 0x12, 0xb2, 0xe2, 0xdc, 0xd5,
	0x9c, 0x12, 0xfb, 0x19, 0xdf, 0x59, 0xa4, 0xdf, 0x80, 0x91, 0x91, 0x31, 0x1f, 0xa7, 0x0b, 0x52,
	0x47, 0xc4, 0x70, 0x42, 0xc9, 0x82, 0x18, 0xfd, 0x05, 0x40, 0xe7, 0x3f, 0xad, 0x23, 0x58, 0x58,
	0x98, 0xfc, 0xfc, 0x7b, 0xbf, 0x3f, 0x7e, 0xbe, 0x7b, 0xe8, 0x84, 0x47, 0xd4, 0x85, 0x8c, 0xd0,
	0x2c, 0x04, 0x98, 0xbb, 0x74, 0x99, 0x02, 0xcf, 0x33, 0xea, 0xa4, 0x19, 0x08, 0x30, 0x4c, 0x1e,
	0xd1, 0x84, 0x8a, 0x0f, 0x90, 0xcd, 0x1d, 0x1e, 0x51, 0xe7, 0x8e, 0x78, 0xfc, 0x38, 0x82, 0x08,
	0x4a, 0x92, 0xab, 0xaa, 0x8a, 0x6f, 0x7f, 0xd1, 0xd1, 0xc0, 0x03, 0x98, 0x5f, 0x12, 0xc2, 0x5f,
	0xd6, 0x56, 0x86, 0x8b, 0x76, 0x95, 0x24, 0x60, 0xc4, 0xd4, 0x87, 0xfa, 0x68, 0xdf, 0x3b, 0x5a,
	0x4b, 0xdc, 0x53, 0xb4, 0xc9, 0x8b, 0x9f, 0x12, 0x37, 0x4d, 0xbf, 0x29, 0x94, 0x00, 0x08, 0xe1,
	0x4a, 0xf0, 0xe0, 0x5e, 0xa0, 0x3c, 0x2b, 0x41, 0xdd, 0xf4, 0x9b, 0xc2, 0x98, 0xa0, 0x83, 0xab,
	0x7c, 0x71, 0xc5, 0x16, 0x8b, 0x98, 0x26, 0x22, 0x78, 0x9f, 0xd3, 0x9c, 0x9a, 0x9d, 0x61, 0x67,
	0xd4, 0xf5, 0x4e, 0x0a, 0x89, 0xcd, 0xeb, 0x69, 0xbc, 0x18, 0xdb, 0x7f, 0x50, 0x6c, 0x7f, 0xd0,
	0xc2, 0x5e, 0x2b, 0x68, 0xdc, 0xff, 0xb8, 0xc2, 0xda, 0xe7, 0x15, 0xd6, 0x7e, 0xac, 0xb0, 0x66,
	0xff, 0xea, 0xa0, 0xa3, 0x57, 0xd3, 0x4c, 0xb0, 0x19, 0x4b, 0xa7, 0x82, 0x41, 0xf2, 0x1f, 0x87,
	0xba, 0x44, 0x87, 0x69, 0x3b, 0x3a, 0x60, 0x09, 0xa1, 0x4b, 0xb3, 0x33, 0xd4, 0x47, 0x5d, 0xcf,
	0x2a, 0x24, 0x3e, 0xae, 0xc6, 0xfa, 0x0b, 0xc9, 0xf6, 0x8d, 0x2d, 0x74, 0xa2, 0x40, 0xe3, 0x2d,
	0xda, 0x6b, 0x8e, 0xd7, 0xec, 0x96, 0x9f, 0x70, 0x7e, 0x23, 0xb1, 0xf6, 0x4d, 0xe2, 0xd3, 0x88,
	0x89, 0x77, 0x79, 0xe8, 0xcc, 0x20, 0x76, 0x67, 0xc0, 0x63, 0xe0, 0xf5, 0xe3, 0x8c, 0x93, 0xb9,
	0x2b, 0xae, 0x53, 0xca, 0x9d, 0x49, 0x22, 0x0a, 0x89, 0x1f, 0x55, 0x99, 0x8d, 0x8f, 0xed, 0xdf,
	0x59, 0x1a, 0x21, 0x42, 0x21, 0x15, 0xc1, 0x34, 0x86, 0x3c, 0x11, 0xe6, 0x4e, 0x19, 0xf0, 0xfc,
	0x9f, 0x03, 0x0e, 0xaa, 0x80, 0x7b, 0x27, 0xdb, 0xdf, 0x0f, 0xa9, 0x38, 0x2f, 0x6b, 0x63, 0x8c,
	0xfa, 0x8c, 0x07, 0xf5, 0xa1, 0x51, 0x62, 0xf6, 0x86, 0xfa, 0x68, 0xcf, 0x7b, 0x52, 0x48, 0x7c,
	0x58, 0xe9, 0xda, 0x5d, 0xdb, 0x7f, 0xc8, 0xf8, 0x45, 0xf3, 0x66, 0x9c, 0xa2, 0x9d, 0x0c, 0xf2,
	0x84, 0x98, 0xbb, 0xe5, 0x1f, 0x1c, 0x14, 0x12, 0xf7, 0x2b, 0x51, 0x09, 0xdb, 0x7e, 0xd5, 0xde,
	0xbe, 0x01, 0xde, 0xc5, 0xcd, 0xda, 0xd2, 0x6f, 0xd7, 0x96, 0xfe, 0x7d, 0x6d, 0xe9, 0x9f, 0x36,
	0x96, 0x76, 0xbb, 0xb1, 0xb4, 0xaf, 0x1b, 0x4b, 0x7b, 0xf3, 0xb4, 0x35, 0x13, 0x8f, 0xe8, 0x59,
	0xbd, 0x27, 0xaa, 0x76, 0x97, 0xad, 0x95, 0x2a, 0xa7, 0x0b, 0x7b, 0xe5, 0x82, 0x3c, 0xfb, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x21, 0x21, 0x81, 0x28, 0x70, 0x03, 0x00, 0x00,
}

func (m *BookOddsExposure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BookOddsExposure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BookOddsExposure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FulfillmentQueue) > 0 {
		dAtA2 := make([]byte, len(m.FulfillmentQueue)*10)
		var j1 int
		for _, num := range m.FulfillmentQueue {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintExposure(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OddsID) > 0 {
		i -= len(m.OddsID)
		copy(dAtA[i:], m.OddsID)
		i = encodeVarintExposure(dAtA, i, uint64(len(m.OddsID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BookID) > 0 {
		i -= len(m.BookID)
		copy(dAtA[i:], m.BookID)
		i = encodeVarintExposure(dAtA, i, uint64(len(m.BookID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ParticipationExposure) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParticipationExposure) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParticipationExposure) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Round != 0 {
		i = encodeVarintExposure(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x38
	}
	if m.IsFulfilled {
		i--
		if m.IsFulfilled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.BetAmount.Size()
		i -= size
		if _, err := m.BetAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExposure(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Exposure.Size()
		i -= size
		if _, err := m.Exposure.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExposure(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.ParticipationIndex != 0 {
		i = encodeVarintExposure(dAtA, i, uint64(m.ParticipationIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OddsID) > 0 {
		i -= len(m.OddsID)
		copy(dAtA[i:], m.OddsID)
		i = encodeVarintExposure(dAtA, i, uint64(len(m.OddsID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BookID) > 0 {
		i -= len(m.BookID)
		copy(dAtA[i:], m.BookID)
		i = encodeVarintExposure(dAtA, i, uint64(len(m.BookID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExposure(dAtA []byte, offset int, v uint64) int {
	offset -= sovExposure(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BookOddsExposure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BookID)
	if l > 0 {
		n += 1 + l + sovExposure(uint64(l))
	}
	l = len(m.OddsID)
	if l > 0 {
		n += 1 + l + sovExposure(uint64(l))
	}
	if len(m.FulfillmentQueue) > 0 {
		l = 0
		for _, e := range m.FulfillmentQueue {
			l += sovExposure(uint64(e))
		}
		n += 1 + sovExposure(uint64(l)) + l
	}
	return n
}

func (m *ParticipationExposure) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BookID)
	if l > 0 {
		n += 1 + l + sovExposure(uint64(l))
	}
	l = len(m.OddsID)
	if l > 0 {
		n += 1 + l + sovExposure(uint64(l))
	}
	if m.ParticipationIndex != 0 {
		n += 1 + sovExposure(uint64(m.ParticipationIndex))
	}
	l = m.Exposure.Size()
	n += 1 + l + sovExposure(uint64(l))
	l = m.BetAmount.Size()
	n += 1 + l + sovExposure(uint64(l))
	if m.IsFulfilled {
		n += 2
	}
	if m.Round != 0 {
		n += 1 + sovExposure(uint64(m.Round))
	}
	return n
}

func sovExposure(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExposure(x uint64) (n int) {
	return sovExposure(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BookOddsExposure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExposure
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
			return fmt.Errorf("proto: BookOddsExposure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BookOddsExposure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BookID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OddsID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExposure
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
				m.FulfillmentQueue = append(m.FulfillmentQueue, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExposure
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
					return ErrInvalidLengthExposure
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthExposure
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
				if elementCount != 0 && len(m.FulfillmentQueue) == 0 {
					m.FulfillmentQueue = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExposure
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
					m.FulfillmentQueue = append(m.FulfillmentQueue, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field FulfillmentQueue", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExposure(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExposure
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
func (m *ParticipationExposure) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExposure
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
			return fmt.Errorf("proto: ParticipationExposure: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParticipationExposure: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BookID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OddsID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OddsID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParticipationIndex", wireType)
			}
			m.ParticipationIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParticipationIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exposure", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Exposure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BetAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
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
				return ErrInvalidLengthExposure
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExposure
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BetAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsFulfilled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsFulfilled = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExposure
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExposure(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExposure
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
func skipExposure(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExposure
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
					return 0, ErrIntOverflowExposure
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
					return 0, ErrIntOverflowExposure
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
				return 0, ErrInvalidLengthExposure
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExposure
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExposure
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExposure        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExposure          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExposure = fmt.Errorf("proto: unexpected end of group")
)