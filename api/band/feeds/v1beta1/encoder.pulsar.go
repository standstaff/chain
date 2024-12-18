// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package feedsv1beta1

import (
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: band/feeds/v1beta1/encoder.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Encoder is an enumerator that defines the mode of encoding message in tss module.
type Encoder int32

const (
	// ENCODER_UNSPECIFIED is an unspecified encoder mode.
	Encoder_ENCODER_UNSPECIFIED Encoder = 0
	// ENCODER_FIXED_POINT_ABI is a fixed-point price abi encoder (price * 10^9).
	Encoder_ENCODER_FIXED_POINT_ABI Encoder = 1
	// ENCODER_TICK_ABI is a tick abi encoder.
	Encoder_ENCODER_TICK_ABI Encoder = 2
)

// Enum value maps for Encoder.
var (
	Encoder_name = map[int32]string{
		0: "ENCODER_UNSPECIFIED",
		1: "ENCODER_FIXED_POINT_ABI",
		2: "ENCODER_TICK_ABI",
	}
	Encoder_value = map[string]int32{
		"ENCODER_UNSPECIFIED":     0,
		"ENCODER_FIXED_POINT_ABI": 1,
		"ENCODER_TICK_ABI":        2,
	}
)

func (x Encoder) Enum() *Encoder {
	p := new(Encoder)
	*p = x
	return p
}

func (x Encoder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Encoder) Descriptor() protoreflect.EnumDescriptor {
	return file_band_feeds_v1beta1_encoder_proto_enumTypes[0].Descriptor()
}

func (Encoder) Type() protoreflect.EnumType {
	return &file_band_feeds_v1beta1_encoder_proto_enumTypes[0]
}

func (x Encoder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Encoder.Descriptor instead.
func (Encoder) EnumDescriptor() ([]byte, []int) {
	return file_band_feeds_v1beta1_encoder_proto_rawDescGZIP(), []int{0}
}

var File_band_feeds_v1beta1_encoder_proto protoreflect.FileDescriptor

var file_band_feeds_v1beta1_encoder_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x61, 0x6e, 0x64, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x5b, 0x0a, 0x07,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4e, 0x43, 0x4f, 0x44,
	0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1b, 0x0a, 0x17, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x58, 0x45,
	0x44, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x41, 0x42, 0x49, 0x10, 0x01, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x5f, 0x41, 0x42,
	0x49, 0x10, 0x02, 0x1a, 0x04, 0x88, 0xa3, 0x1e, 0x00, 0x42, 0xd6, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x62, 0x61, 0x6e, 0x64, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x0c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x61, 0x6e, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x6e, 0x64, 0x2f,
	0x66, 0x65, 0x65, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x66, 0x65,
	0x65, 0x64, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x46, 0x58,
	0xaa, 0x02, 0x12, 0x42, 0x61, 0x6e, 0x64, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x12, 0x42, 0x61, 0x6e, 0x64, 0x5c, 0x46, 0x65, 0x65,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1e, 0x42, 0x61, 0x6e,
	0x64, 0x5c, 0x46, 0x65, 0x65, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x42, 0x61,
	0x6e, 0x64, 0x3a, 0x3a, 0x46, 0x65, 0x65, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_band_feeds_v1beta1_encoder_proto_rawDescOnce sync.Once
	file_band_feeds_v1beta1_encoder_proto_rawDescData = file_band_feeds_v1beta1_encoder_proto_rawDesc
)

func file_band_feeds_v1beta1_encoder_proto_rawDescGZIP() []byte {
	file_band_feeds_v1beta1_encoder_proto_rawDescOnce.Do(func() {
		file_band_feeds_v1beta1_encoder_proto_rawDescData = protoimpl.X.CompressGZIP(file_band_feeds_v1beta1_encoder_proto_rawDescData)
	})
	return file_band_feeds_v1beta1_encoder_proto_rawDescData
}

var file_band_feeds_v1beta1_encoder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_band_feeds_v1beta1_encoder_proto_goTypes = []interface{}{
	(Encoder)(0), // 0: band.feeds.v1beta1.Encoder
}
var file_band_feeds_v1beta1_encoder_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_band_feeds_v1beta1_encoder_proto_init() }
func file_band_feeds_v1beta1_encoder_proto_init() {
	if File_band_feeds_v1beta1_encoder_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_band_feeds_v1beta1_encoder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_band_feeds_v1beta1_encoder_proto_goTypes,
		DependencyIndexes: file_band_feeds_v1beta1_encoder_proto_depIdxs,
		EnumInfos:         file_band_feeds_v1beta1_encoder_proto_enumTypes,
	}.Build()
	File_band_feeds_v1beta1_encoder_proto = out.File
	file_band_feeds_v1beta1_encoder_proto_rawDesc = nil
	file_band_feeds_v1beta1_encoder_proto_goTypes = nil
	file_band_feeds_v1beta1_encoder_proto_depIdxs = nil
}