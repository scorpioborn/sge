// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package bet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: sgenetwork/sge/bet/odds_type.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OddsType is the representation of the type of the odds.
type OddsType int32

const (
	// invalid odds type
	OddsType_ODDS_TYPE_UNSPECIFIED OddsType = 0
	// decimal odds type (european)
	OddsType_ODDS_TYPE_DECIMAL OddsType = 1
	// fractional odds type (british)
	OddsType_ODDS_TYPE_FRACTIONAL OddsType = 2
	// moneyline odds type (american)
	OddsType_ODDS_TYPE_MONEYLINE OddsType = 3
)

// Enum value maps for OddsType.
var (
	OddsType_name = map[int32]string{
		0: "ODDS_TYPE_UNSPECIFIED",
		1: "ODDS_TYPE_DECIMAL",
		2: "ODDS_TYPE_FRACTIONAL",
		3: "ODDS_TYPE_MONEYLINE",
	}
	OddsType_value = map[string]int32{
		"ODDS_TYPE_UNSPECIFIED": 0,
		"ODDS_TYPE_DECIMAL":     1,
		"ODDS_TYPE_FRACTIONAL":  2,
		"ODDS_TYPE_MONEYLINE":   3,
	}
)

func (x OddsType) Enum() *OddsType {
	p := new(OddsType)
	*p = x
	return p
}

func (x OddsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OddsType) Descriptor() protoreflect.EnumDescriptor {
	return file_sgenetwork_sge_bet_odds_type_proto_enumTypes[0].Descriptor()
}

func (OddsType) Type() protoreflect.EnumType {
	return &file_sgenetwork_sge_bet_odds_type_proto_enumTypes[0]
}

func (x OddsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OddsType.Descriptor instead.
func (OddsType) EnumDescriptor() ([]byte, []int) {
	return file_sgenetwork_sge_bet_odds_type_proto_rawDescGZIP(), []int{0}
}

var File_sgenetwork_sge_bet_odds_type_proto protoreflect.FileDescriptor

var file_sgenetwork_sge_bet_odds_type_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x67, 0x65,
	0x2f, 0x62, 0x65, 0x74, 0x2f, 0x6f, 0x64, 0x64, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x73, 0x67, 0x65, 0x2e, 0x62, 0x65, 0x74, 0x2a, 0x6f, 0x0a, 0x08, 0x4f, 0x64, 0x64, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x44, 0x44, 0x53, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x4f, 0x44, 0x44, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x43,
	0x49, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x44, 0x44, 0x53, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x02,
	0x12, 0x17, 0x0a, 0x13, 0x4f, 0x44, 0x44, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f,
	0x4e, 0x45, 0x59, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x42, 0xc4, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x73, 0x67, 0x65,
	0x2e, 0x62, 0x65, 0x74, 0x42, 0x0d, 0x4f, 0x64, 0x64, 0x73, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x67, 0x65, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x67,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x73, 0x67, 0x65, 0x2f, 0x62, 0x65, 0x74, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x42, 0xaa, 0x02,
	0x12, 0x53, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x53, 0x67, 0x65, 0x2e,
	0x42, 0x65, 0x74, 0xca, 0x02, 0x12, 0x53, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5c, 0x53, 0x67, 0x65, 0x5c, 0x42, 0x65, 0x74, 0xe2, 0x02, 0x1e, 0x53, 0x67, 0x65, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5c, 0x53, 0x67, 0x65, 0x5c, 0x42, 0x65, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x67, 0x65, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x3a, 0x3a, 0x53, 0x67, 0x65, 0x3a, 0x3a, 0x42, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sgenetwork_sge_bet_odds_type_proto_rawDescOnce sync.Once
	file_sgenetwork_sge_bet_odds_type_proto_rawDescData = file_sgenetwork_sge_bet_odds_type_proto_rawDesc
)

func file_sgenetwork_sge_bet_odds_type_proto_rawDescGZIP() []byte {
	file_sgenetwork_sge_bet_odds_type_proto_rawDescOnce.Do(func() {
		file_sgenetwork_sge_bet_odds_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_sgenetwork_sge_bet_odds_type_proto_rawDescData)
	})
	return file_sgenetwork_sge_bet_odds_type_proto_rawDescData
}

var file_sgenetwork_sge_bet_odds_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sgenetwork_sge_bet_odds_type_proto_goTypes = []interface{}{
	(OddsType)(0), // 0: sgenetwork.sge.bet.OddsType
}
var file_sgenetwork_sge_bet_odds_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sgenetwork_sge_bet_odds_type_proto_init() }
func file_sgenetwork_sge_bet_odds_type_proto_init() {
	if File_sgenetwork_sge_bet_odds_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sgenetwork_sge_bet_odds_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sgenetwork_sge_bet_odds_type_proto_goTypes,
		DependencyIndexes: file_sgenetwork_sge_bet_odds_type_proto_depIdxs,
		EnumInfos:         file_sgenetwork_sge_bet_odds_type_proto_enumTypes,
	}.Build()
	File_sgenetwork_sge_bet_odds_type_proto = out.File
	file_sgenetwork_sge_bet_odds_type_proto_rawDesc = nil
	file_sgenetwork_sge_bet_odds_type_proto_goTypes = nil
	file_sgenetwork_sge_bet_odds_type_proto_depIdxs = nil
}
