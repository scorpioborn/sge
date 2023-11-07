// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/reward/ticket.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
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

// CreateCampaignPayload is the type for campaign creation payload.
type CreateCampaignPayload struct {
	// promoter is the address of campaign promoter.
	// Funds for the campaign would be deducted from this account.
	Promoter string `protobuf:"bytes,1,opt,name=promoter,proto3" json:"promoter,omitempty"`
	// start_ts is the start timestamp of the campaign.
	StartTs uint64 `protobuf:"varint,2,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	// end_ts is the end timestamp of the campaign.
	EndTs uint64 `protobuf:"varint,3,opt,name=end_ts,json=endTs,proto3" json:"end_ts,omitempty"`
	// category is the category of the campaign.
	Category RewardCategory `protobuf:"varint,4,opt,name=category,proto3,enum=sgenetwork.sge.reward.RewardCategory" json:"category,omitempty"`
	// reward_type is the type of reward.
	RewardType RewardType `protobuf:"varint,5,opt,name=reward_type,json=rewardType,proto3,enum=sgenetwork.sge.reward.RewardType" json:"reward_type,omitempty"`
	// Reward amount
	RewardAmountType RewardAmountType `protobuf:"varint,6,opt,name=reward_amount_type,json=rewardAmountType,proto3,enum=sgenetwork.sge.reward.RewardAmountType" json:"reward_amount_type,omitempty"`
	// reward_amount is the amount of reward.
	RewardAmount *RewardAmount `protobuf:"bytes,7,opt,name=reward_amount,json=rewardAmount,proto3" json:"reward_amount,omitempty"`
	// total_funds is the total funds allocated to the campaign.
	TotalFunds cosmossdk_io_math.Int `protobuf:"bytes,8,opt,name=total_funds,json=totalFunds,proto3,customtype=cosmossdk.io/math.Int" json:"total_funds" yaml:"total_funds"`
	// is_active is the flag to check if the campaign is active or not.
	IsActive bool `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// meta is the metadata of the campaign.
	// It is a stringified base64 encoded json.
	Meta string `protobuf:"bytes,10,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (m *CreateCampaignPayload) Reset()         { *m = CreateCampaignPayload{} }
func (m *CreateCampaignPayload) String() string { return proto.CompactTextString(m) }
func (*CreateCampaignPayload) ProtoMessage()    {}
func (*CreateCampaignPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d710bc1249ca8ae, []int{0}
}
func (m *CreateCampaignPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateCampaignPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateCampaignPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateCampaignPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCampaignPayload.Merge(m, src)
}
func (m *CreateCampaignPayload) XXX_Size() int {
	return m.Size()
}
func (m *CreateCampaignPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCampaignPayload.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCampaignPayload proto.InternalMessageInfo

func (m *CreateCampaignPayload) GetPromoter() string {
	if m != nil {
		return m.Promoter
	}
	return ""
}

func (m *CreateCampaignPayload) GetStartTs() uint64 {
	if m != nil {
		return m.StartTs
	}
	return 0
}

func (m *CreateCampaignPayload) GetEndTs() uint64 {
	if m != nil {
		return m.EndTs
	}
	return 0
}

func (m *CreateCampaignPayload) GetCategory() RewardCategory {
	if m != nil {
		return m.Category
	}
	return RewardCategory_REWARD_CATEGORY_UNSPECIFIED
}

func (m *CreateCampaignPayload) GetRewardType() RewardType {
	if m != nil {
		return m.RewardType
	}
	return RewardType_REWARD_TYPE_UNSPECIFIED
}

func (m *CreateCampaignPayload) GetRewardAmountType() RewardAmountType {
	if m != nil {
		return m.RewardAmountType
	}
	return RewardAmountType_REWARD_AMOUNT_TYPE_UNSPECIFIED
}

func (m *CreateCampaignPayload) GetRewardAmount() *RewardAmount {
	if m != nil {
		return m.RewardAmount
	}
	return nil
}

func (m *CreateCampaignPayload) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *CreateCampaignPayload) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

// UpdateCampaignPayload is the type for campaign update payload.
type UpdateCampaignPayload struct {
	// end_ts is the end timestamp of the campaign.
	EndTs uint64 `protobuf:"varint,1,opt,name=end_ts,json=endTs,proto3" json:"end_ts,omitempty"`
	// funds is the additional amount added to the campaign.
	Funds string `protobuf:"bytes,2,opt,name=funds,proto3" json:"funds,omitempty"`
	// is_active is the flag to check if the campaign is active or not.
	IsActive bool `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (m *UpdateCampaignPayload) Reset()         { *m = UpdateCampaignPayload{} }
func (m *UpdateCampaignPayload) String() string { return proto.CompactTextString(m) }
func (*UpdateCampaignPayload) ProtoMessage()    {}
func (*UpdateCampaignPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d710bc1249ca8ae, []int{1}
}
func (m *UpdateCampaignPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateCampaignPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateCampaignPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateCampaignPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCampaignPayload.Merge(m, src)
}
func (m *UpdateCampaignPayload) XXX_Size() int {
	return m.Size()
}
func (m *UpdateCampaignPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCampaignPayload.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCampaignPayload proto.InternalMessageInfo

func (m *UpdateCampaignPayload) GetEndTs() uint64 {
	if m != nil {
		return m.EndTs
	}
	return 0
}

func (m *UpdateCampaignPayload) GetFunds() string {
	if m != nil {
		return m.Funds
	}
	return ""
}

func (m *UpdateCampaignPayload) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

type RewardPayloadCommon struct {
	// receiver is the address of the account that receives the reward.
	Receiver string `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// source is the source of the reward.
	// It is used to identify the source of the reward.
	// For example, the source of a referral signup is Type - referral.
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source"`
	// source_code is the source code of the reward.
	// It is used to identify the source of the reward.
	// For example, the source code of a referral signup is referral code of
	// referer.
	SourceCode string `protobuf:"bytes,3,opt,name=source_code,proto3" json:"source_code"`
	// source_uid is the address of the source.
	// It is used to identify the source of the reward.
	// For example, the source uid of a referral signup reward is the address of
	// the referer.
	SourceUID string `protobuf:"bytes,4,opt,name=source_uid,proto3" json:"source_uid"`
}

func (m *RewardPayloadCommon) Reset()         { *m = RewardPayloadCommon{} }
func (m *RewardPayloadCommon) String() string { return proto.CompactTextString(m) }
func (*RewardPayloadCommon) ProtoMessage()    {}
func (*RewardPayloadCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d710bc1249ca8ae, []int{2}
}
func (m *RewardPayloadCommon) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardPayloadCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardPayloadCommon.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardPayloadCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardPayloadCommon.Merge(m, src)
}
func (m *RewardPayloadCommon) XXX_Size() int {
	return m.Size()
}
func (m *RewardPayloadCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardPayloadCommon.DiscardUnknown(m)
}

var xxx_messageInfo_RewardPayloadCommon proto.InternalMessageInfo

func (m *RewardPayloadCommon) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *RewardPayloadCommon) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RewardPayloadCommon) GetSourceCode() string {
	if m != nil {
		return m.SourceCode
	}
	return ""
}

func (m *RewardPayloadCommon) GetSourceUID() string {
	if m != nil {
		return m.SourceUID
	}
	return ""
}

// GrantSignupRewardPayload is the type for signup reward grant payload.
type GrantSignupRewardPayload struct {
	// common is the common properties of a reward
	Common RewardPayloadCommon `protobuf:"bytes,1,opt,name=common,proto3" json:"common"`
}

func (m *GrantSignupRewardPayload) Reset()         { *m = GrantSignupRewardPayload{} }
func (m *GrantSignupRewardPayload) String() string { return proto.CompactTextString(m) }
func (*GrantSignupRewardPayload) ProtoMessage()    {}
func (*GrantSignupRewardPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d710bc1249ca8ae, []int{3}
}
func (m *GrantSignupRewardPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrantSignupRewardPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrantSignupRewardPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrantSignupRewardPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantSignupRewardPayload.Merge(m, src)
}
func (m *GrantSignupRewardPayload) XXX_Size() int {
	return m.Size()
}
func (m *GrantSignupRewardPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantSignupRewardPayload.DiscardUnknown(m)
}

var xxx_messageInfo_GrantSignupRewardPayload proto.InternalMessageInfo

func (m *GrantSignupRewardPayload) GetCommon() RewardPayloadCommon {
	if m != nil {
		return m.Common
	}
	return RewardPayloadCommon{}
}

func init() {
	proto.RegisterType((*CreateCampaignPayload)(nil), "sgenetwork.sge.reward.CreateCampaignPayload")
	proto.RegisterType((*UpdateCampaignPayload)(nil), "sgenetwork.sge.reward.UpdateCampaignPayload")
	proto.RegisterType((*RewardPayloadCommon)(nil), "sgenetwork.sge.reward.RewardPayloadCommon")
	proto.RegisterType((*GrantSignupRewardPayload)(nil), "sgenetwork.sge.reward.GrantSignupRewardPayload")
}

func init() { proto.RegisterFile("sge/reward/ticket.proto", fileDescriptor_5d710bc1249ca8ae) }

var fileDescriptor_5d710bc1249ca8ae = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xad, 0xeb, 0xda, 0x57, 0x40, 0xc8, 0xac, 0x22, 0xdb, 0x44, 0x53, 0x8a, 0x10,
	0x05, 0x41, 0x22, 0x6d, 0x37, 0x38, 0xa0, 0x25, 0x08, 0xb6, 0x1b, 0xf2, 0xba, 0x0b, 0x97, 0xca,
	0x4b, 0x4c, 0x16, 0x6d, 0x89, 0x23, 0xdb, 0xd9, 0xc8, 0xb7, 0xe0, 0xc0, 0x87, 0xda, 0x71, 0x47,
	0x84, 0x44, 0x84, 0xb2, 0xdb, 0x8e, 0x7c, 0x02, 0x14, 0x3b, 0x74, 0x29, 0x9a, 0x26, 0x4e, 0x7e,
	0xff, 0xbc, 0xff, 0xfb, 0x39, 0xcf, 0x7e, 0x32, 0x3c, 0x14, 0x21, 0x75, 0x38, 0x3d, 0x23, 0x3c,
	0x70, 0x64, 0xe4, 0x1f, 0x53, 0x69, 0xa7, 0x9c, 0x49, 0x86, 0x06, 0x22, 0xa4, 0x09, 0x95, 0x67,
	0x8c, 0x1f, 0xdb, 0x22, 0xa4, 0xb6, 0xf6, 0x6c, 0xac, 0x85, 0x2c, 0x64, 0xca, 0xe1, 0x54, 0x91,
	0x36, 0x6f, 0x34, 0x29, 0x7a, 0xa9, 0x13, 0xeb, 0x8d, 0x84, 0x4f, 0xe2, 0x94, 0x44, 0x61, 0xa2,
	0x53, 0xe3, 0x6f, 0x6d, 0x18, 0x78, 0x9c, 0x12, 0x49, 0xbd, 0x3a, 0xf1, 0x91, 0xe4, 0x27, 0x8c,
	0x04, 0x68, 0x03, 0xba, 0x29, 0x67, 0x31, 0x93, 0x94, 0x9b, 0xc6, 0xc8, 0x98, 0xf4, 0xf0, 0x5c,
	0xa3, 0x75, 0xe8, 0x0a, 0x49, 0xb8, 0x9c, 0x49, 0x61, 0x2e, 0x8d, 0x8c, 0x49, 0x1b, 0xaf, 0x2a,
	0x3d, 0x15, 0x68, 0x00, 0x1d, 0x9a, 0x04, 0x55, 0x62, 0x59, 0x25, 0x56, 0x68, 0x12, 0x4c, 0x05,
	0xda, 0x81, 0xae, 0x4f, 0x24, 0x0d, 0x19, 0xcf, 0xcd, 0xf6, 0xc8, 0x98, 0xdc, 0xdb, 0x7a, 0x6a,
	0xdf, 0xd8, 0x9b, 0x8d, 0xd5, 0xe2, 0xd5, 0x66, 0x3c, 0x2f, 0x43, 0x2e, 0xf4, 0xb5, 0x65, 0x26,
	0xf3, 0x94, 0x9a, 0x2b, 0x8a, 0xf2, 0xf8, 0x56, 0xca, 0x34, 0x4f, 0x29, 0x06, 0x3e, 0x8f, 0xd1,
	0x01, 0xa0, 0x9a, 0x41, 0x62, 0x96, 0x25, 0x52, 0xa3, 0x3a, 0x0a, 0xf5, 0xec, 0x56, 0xd4, 0x8e,
	0xf2, 0x2b, 0xe0, 0x7d, 0xfe, 0xcf, 0x17, 0xb4, 0x0b, 0x77, 0x17, 0xb0, 0xe6, 0xea, 0xc8, 0x98,
	0xf4, 0xb7, 0x9e, 0xfc, 0x07, 0x11, 0xdf, 0x69, 0xd2, 0xd0, 0x14, 0xfa, 0x92, 0x49, 0x72, 0x32,
	0xfb, 0x9c, 0x25, 0x81, 0x30, 0xbb, 0xd5, 0xc1, 0xbb, 0xdb, 0xe7, 0x85, 0xd5, 0xfa, 0x51, 0x58,
	0x03, 0x9f, 0x89, 0x98, 0x09, 0x11, 0x1c, 0xdb, 0x11, 0x73, 0x62, 0x22, 0x8f, 0xec, 0xbd, 0x44,
	0xfe, 0x2e, 0x2c, 0x94, 0x93, 0xf8, 0xe4, 0xf5, 0xb8, 0x51, 0x39, 0xc6, 0xa0, 0xd4, 0xfb, 0x4a,
	0xa0, 0x4d, 0xe8, 0x45, 0x62, 0x46, 0x7c, 0x19, 0x9d, 0x52, 0xb3, 0x37, 0x32, 0x26, 0x5d, 0xdc,
	0x8d, 0xc4, 0x8e, 0xd2, 0x08, 0x41, 0x3b, 0xa6, 0x92, 0x98, 0xa0, 0x2e, 0x59, 0xc5, 0x63, 0x02,
	0x83, 0x83, 0x34, 0xb8, 0x61, 0x2a, 0xae, 0xaf, 0xd7, 0x68, 0x5e, 0xef, 0x1a, 0xac, 0xe8, 0x1f,
	0x5e, 0x52, 0x10, 0x2d, 0x16, 0xb7, 0x5d, 0x5e, 0xdc, 0x76, 0xfc, 0xd3, 0x80, 0x07, 0xfa, 0x20,
	0x6a, 0xb6, 0xc7, 0xe2, 0x98, 0x25, 0xd5, 0xdc, 0x71, 0xea, 0xd3, 0xe8, 0xf4, 0x7a, 0xee, 0xfe,
	0x6a, 0xf4, 0x12, 0x3a, 0x82, 0x65, 0xdc, 0xa7, 0x7a, 0x1f, 0x77, 0xad, 0x2c, 0xac, 0xce, 0xbe,
	0xfa, 0x72, 0x55, 0x58, 0x75, 0x0e, 0xd7, 0x2b, 0x7a, 0x0b, 0x7d, 0x1d, 0xcd, 0x7c, 0x16, 0xe8,
	0x1f, 0xe8, 0xb9, 0x8f, 0xca, 0xc2, 0x02, 0x5d, 0xe2, 0xb1, 0xa0, 0x2a, 0x6b, 0x9a, 0x70, 0x53,
	0xa0, 0x37, 0x00, 0xb5, 0xcc, 0xa2, 0x40, 0x8d, 0x6d, 0xcf, 0xdd, 0x2c, 0x0b, 0xab, 0xa7, 0xeb,
	0x0f, 0xf6, 0xde, 0x5d, 0x15, 0x56, 0xc3, 0x82, 0x1b, 0xf1, 0x38, 0x00, 0xf3, 0x03, 0x27, 0x89,
	0xdc, 0x8f, 0xc2, 0x24, 0x4b, 0x17, 0x3a, 0x45, 0xbb, 0xd0, 0xf1, 0x55, 0xb7, 0xaa, 0xc3, 0xfe,
	0xd6, 0x8b, 0x5b, 0x07, 0x65, 0xe1, 0x7c, 0xdc, 0x76, 0x35, 0x0c, 0xb8, 0xae, 0x77, 0xbd, 0xf3,
	0x72, 0x68, 0x5c, 0x94, 0x43, 0xe3, 0x57, 0x39, 0x34, 0xbe, 0x5e, 0x0e, 0x5b, 0x17, 0x97, 0xc3,
	0xd6, 0xf7, 0xcb, 0x61, 0xeb, 0xd3, 0xf3, 0x30, 0x92, 0x47, 0xd9, 0xa1, 0xed, 0xb3, 0xd8, 0x11,
	0x21, 0x7d, 0x55, 0xe3, 0xab, 0xd8, 0xf9, 0x32, 0x7f, 0x6c, 0xf2, 0x94, 0x8a, 0xc3, 0x8e, 0x7a,
	0x0b, 0xb6, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xf1, 0xda, 0xf1, 0x87, 0x04, 0x00, 0x00,
}

func (m *CreateCampaignPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateCampaignPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateCampaignPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x52
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.TotalFunds.Size()
		i -= size
		if _, err := m.TotalFunds.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTicket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.RewardAmount != nil {
		{
			size, err := m.RewardAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTicket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.RewardAmountType != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.RewardAmountType))
		i--
		dAtA[i] = 0x30
	}
	if m.RewardType != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.RewardType))
		i--
		dAtA[i] = 0x28
	}
	if m.Category != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.Category))
		i--
		dAtA[i] = 0x20
	}
	if m.EndTs != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTs))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTs != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.StartTs))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Promoter) > 0 {
		i -= len(m.Promoter)
		copy(dAtA[i:], m.Promoter)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Promoter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateCampaignPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateCampaignPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateCampaignPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Funds) > 0 {
		i -= len(m.Funds)
		copy(dAtA[i:], m.Funds)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Funds)))
		i--
		dAtA[i] = 0x12
	}
	if m.EndTs != 0 {
		i = encodeVarintTicket(dAtA, i, uint64(m.EndTs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RewardPayloadCommon) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardPayloadCommon) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardPayloadCommon) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceUID) > 0 {
		i -= len(m.SourceUID)
		copy(dAtA[i:], m.SourceUID)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.SourceUID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceCode) > 0 {
		i -= len(m.SourceCode)
		copy(dAtA[i:], m.SourceCode)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.SourceCode)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTicket(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GrantSignupRewardPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrantSignupRewardPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrantSignupRewardPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Common.MarshalToSizedBuffer(dAtA[:i])
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
func (m *CreateCampaignPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Promoter)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.StartTs != 0 {
		n += 1 + sovTicket(uint64(m.StartTs))
	}
	if m.EndTs != 0 {
		n += 1 + sovTicket(uint64(m.EndTs))
	}
	if m.Category != 0 {
		n += 1 + sovTicket(uint64(m.Category))
	}
	if m.RewardType != 0 {
		n += 1 + sovTicket(uint64(m.RewardType))
	}
	if m.RewardAmountType != 0 {
		n += 1 + sovTicket(uint64(m.RewardAmountType))
	}
	if m.RewardAmount != nil {
		l = m.RewardAmount.Size()
		n += 1 + l + sovTicket(uint64(l))
	}
	l = m.TotalFunds.Size()
	n += 1 + l + sovTicket(uint64(l))
	if m.IsActive {
		n += 2
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	return n
}

func (m *UpdateCampaignPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EndTs != 0 {
		n += 1 + sovTicket(uint64(m.EndTs))
	}
	l = len(m.Funds)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	if m.IsActive {
		n += 2
	}
	return n
}

func (m *RewardPayloadCommon) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = len(m.SourceCode)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	l = len(m.SourceUID)
	if l > 0 {
		n += 1 + l + sovTicket(uint64(l))
	}
	return n
}

func (m *GrantSignupRewardPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Common.Size()
	n += 1 + l + sovTicket(uint64(l))
	return n
}

func sovTicket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTicket(x uint64) (n int) {
	return sovTicket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateCampaignPayload) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreateCampaignPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateCampaignPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Promoter", wireType)
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
			m.Promoter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTs", wireType)
			}
			m.StartTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTs", wireType)
			}
			m.EndTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			m.Category = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Category |= RewardCategory(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardType", wireType)
			}
			m.RewardType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardType |= RewardType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAmountType", wireType)
			}
			m.RewardAmountType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardAmountType |= RewardAmountType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAmount", wireType)
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
			if m.RewardAmount == nil {
				m.RewardAmount = &RewardAmount{}
			}
			if err := m.RewardAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFunds", wireType)
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
			if err := m.TotalFunds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
			m.IsActive = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
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
			m.Meta = string(dAtA[iNdEx:postIndex])
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
func (m *UpdateCampaignPayload) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UpdateCampaignPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateCampaignPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTs", wireType)
			}
			m.EndTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
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
			m.Funds = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTicket
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
			m.IsActive = bool(v != 0)
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
func (m *RewardPayloadCommon) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RewardPayloadCommon: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardPayloadCommon: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
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
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
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
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCode", wireType)
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
			m.SourceCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceUID", wireType)
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
			m.SourceUID = string(dAtA[iNdEx:postIndex])
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
func (m *GrantSignupRewardPayload) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GrantSignupRewardPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrantSignupRewardPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Common", wireType)
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
			if err := m.Common.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
