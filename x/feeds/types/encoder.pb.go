// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: band/feeds/v1beta1/encoder.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	math "math"
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

// Encoder is an enumerator that defines the mode of encoding message in tss module.
type Encoder int32

const (
	// ENCODER_UNSPECIFIED is an unspecified encoder mode.
	ENCODER_UNSPECIFIED Encoder = 0
	// ENCODER_FIXED_POINT_ABI is a fixed-point price abi encoder (price * 10^9).
	ENCODER_FIXED_POINT_ABI Encoder = 1
	// ENCODER_TICK_ABI is a tick abi encoder.
	ENCODER_TICK_ABI Encoder = 2
)

var Encoder_name = map[int32]string{
	0: "ENCODER_UNSPECIFIED",
	1: "ENCODER_FIXED_POINT_ABI",
	2: "ENCODER_TICK_ABI",
}

var Encoder_value = map[string]int32{
	"ENCODER_UNSPECIFIED":     0,
	"ENCODER_FIXED_POINT_ABI": 1,
	"ENCODER_TICK_ABI":        2,
}

func (x Encoder) String() string {
	return proto.EnumName(Encoder_name, int32(x))
}

func (Encoder) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ac3e992b65436f01, []int{0}
}

func init() {
	proto.RegisterEnum("band.feeds.v1beta1.Encoder", Encoder_name, Encoder_value)
}

func init() { proto.RegisterFile("band/feeds/v1beta1/encoder.proto", fileDescriptor_ac3e992b65436f01) }

var fileDescriptor_ac3e992b65436f01 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x4a, 0xcc, 0x4b,
	0xd1, 0x4f, 0x4b, 0x4d, 0x4d, 0x29, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f,
	0xcd, 0x4b, 0xce, 0x4f, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0xa9,
	0xd0, 0x03, 0xab, 0xd0, 0x83, 0xaa, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xeb, 0x83,
	0x58, 0x10, 0x95, 0x5a, 0xd1, 0x5c, 0xec, 0xae, 0x10, 0xad, 0x42, 0xe2, 0x5c, 0xc2, 0xae, 0x7e,
	0xce, 0xfe, 0x2e, 0xae, 0x41, 0xf1, 0xa1, 0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae,
	0x2e, 0x02, 0x0c, 0x42, 0xd2, 0x5c, 0xe2, 0x30, 0x09, 0x37, 0xcf, 0x08, 0x57, 0x97, 0xf8, 0x00,
	0x7f, 0x4f, 0xbf, 0x90, 0x78, 0x47, 0x27, 0x4f, 0x01, 0x46, 0x21, 0x11, 0x2e, 0x01, 0x98, 0x64,
	0x88, 0xa7, 0xb3, 0x37, 0x58, 0x94, 0x49, 0x8a, 0xa5, 0x63, 0xb1, 0x1c, 0x83, 0x93, 0xc7, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0xdc, 0x0a, 0x76, 0x4c, 0x72, 0x7e, 0x8e, 0x7e, 0x72, 0x46,
	0x62, 0x66, 0x9e, 0x7e, 0x99, 0xb1, 0x7e, 0x05, 0xd4, 0x83, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49,
	0x6c, 0x60, 0x05, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x0d, 0x10, 0xc0, 0xfb, 0x00,
	0x00, 0x00,
}