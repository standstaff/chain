// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: restake/v1beta1/types.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// Vault is used for tracking the status of the vaults.
type Vault struct {
	// key is the key of the vault.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// is_active is the status of the vault
	IsActive bool `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (m *Vault) Reset()         { *m = Vault{} }
func (m *Vault) String() string { return proto.CompactTextString(m) }
func (*Vault) ProtoMessage()    {}
func (*Vault) Descriptor() ([]byte, []int) {
	return fileDescriptor_be4ee20adc0c7118, []int{0}
}
func (m *Vault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vault.Merge(m, src)
}
func (m *Vault) XXX_Size() int {
	return m.Size()
}
func (m *Vault) XXX_DiscardUnknown() {
	xxx_messageInfo_Vault.DiscardUnknown(m)
}

var xxx_messageInfo_Vault proto.InternalMessageInfo

func (m *Vault) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Vault) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

// Lock is used to store lock information of each user on each vault.
type Lock struct {
	// staker_address is the owner's address of the staker.
	StakerAddress string `protobuf:"bytes,1,opt,name=staker_address,json=stakerAddress,proto3" json:"staker_address,omitempty"`
	// key is the key of the vault that this lock is locked to.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// power is the number of locked power.
	Power cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=power,proto3,customtype=cosmossdk.io/math.Int" json:"power"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_be4ee20adc0c7118, []int{1}
}
func (m *Lock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lock.Merge(m, src)
}
func (m *Lock) XXX_Size() int {
	return m.Size()
}
func (m *Lock) XXX_DiscardUnknown() {
	xxx_messageInfo_Lock.DiscardUnknown(m)
}

var xxx_messageInfo_Lock proto.InternalMessageInfo

func (m *Lock) GetStakerAddress() string {
	if m != nil {
		return m.StakerAddress
	}
	return ""
}

func (m *Lock) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// Stake is used to store staked coins of an address.
type Stake struct {
	// staker_address is the address that this stake belongs to.
	StakerAddress string `protobuf:"bytes,1,opt,name=staker_address,json=stakerAddress,proto3" json:"staker_address,omitempty"`
	// coins are the coins that the address has staked.
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Stake) Reset()         { *m = Stake{} }
func (m *Stake) String() string { return proto.CompactTextString(m) }
func (*Stake) ProtoMessage()    {}
func (*Stake) Descriptor() ([]byte, []int) {
	return fileDescriptor_be4ee20adc0c7118, []int{2}
}
func (m *Stake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stake.Merge(m, src)
}
func (m *Stake) XXX_Size() int {
	return m.Size()
}
func (m *Stake) XXX_DiscardUnknown() {
	xxx_messageInfo_Stake.DiscardUnknown(m)
}

var xxx_messageInfo_Stake proto.InternalMessageInfo

func (m *Stake) GetStakerAddress() string {
	if m != nil {
		return m.StakerAddress
	}
	return ""
}

func (m *Stake) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

// LockResponse is used as response of the query to show the power
// that is locked by the vault for the user.
type LockResponse struct {
	// key is the key of the vault that this lock belongs to.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// power is the number of locked power.
	Power cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=power,proto3,customtype=cosmossdk.io/math.Int" json:"power"`
}

func (m *LockResponse) Reset()         { *m = LockResponse{} }
func (m *LockResponse) String() string { return proto.CompactTextString(m) }
func (*LockResponse) ProtoMessage()    {}
func (*LockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be4ee20adc0c7118, []int{3}
}
func (m *LockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockResponse.Merge(m, src)
}
func (m *LockResponse) XXX_Size() int {
	return m.Size()
}
func (m *LockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LockResponse proto.InternalMessageInfo

func (m *LockResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Vault)(nil), "restake.v1beta1.Vault")
	proto.RegisterType((*Lock)(nil), "restake.v1beta1.Lock")
	proto.RegisterType((*Stake)(nil), "restake.v1beta1.Stake")
	proto.RegisterType((*LockResponse)(nil), "restake.v1beta1.LockResponse")
}

func init() { proto.RegisterFile("restake/v1beta1/types.proto", fileDescriptor_be4ee20adc0c7118) }

var fileDescriptor_be4ee20adc0c7118 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xbd, 0x8e, 0xd3, 0x40,
	0x10, 0xf6, 0x3a, 0x67, 0x74, 0xb7, 0xfc, 0x5b, 0x87, 0xe4, 0xbb, 0x93, 0xec, 0xe8, 0x2a, 0x0b,
	0x14, 0x2f, 0xc7, 0x89, 0x06, 0x21, 0xa1, 0x84, 0xea, 0x10, 0x95, 0x4f, 0xa2, 0xa0, 0x89, 0xd6,
	0xf6, 0xe2, 0xac, 0x7c, 0xde, 0xb5, 0xbc, 0x7b, 0x81, 0x7b, 0x0b, 0x1e, 0x01, 0x51, 0x21, 0x2a,
	0x8a, 0x88, 0x67, 0x48, 0x19, 0xa5, 0x42, 0x14, 0x01, 0x25, 0x05, 0x3c, 0x06, 0xda, 0x9f, 0x10,
	0x0a, 0xba, 0x34, 0xf6, 0xce, 0x7e, 0x33, 0xf3, 0x7d, 0xfb, 0xcd, 0xc0, 0xa3, 0x96, 0x08, 0x89,
	0x2b, 0x82, 0xc6, 0x27, 0x19, 0x91, 0xf8, 0x04, 0xc9, 0xab, 0x86, 0x88, 0xa4, 0x69, 0xb9, 0xe4,
	0xfe, 0x6d, 0x0b, 0x26, 0x16, 0x3c, 0xdc, 0x2f, 0x79, 0xc9, 0x35, 0x86, 0xd4, 0xc9, 0xa4, 0x1d,
	0xde, 0xc5, 0x35, 0x65, 0x1c, 0xe9, 0xaf, 0xbd, 0x3a, 0xc8, 0xb9, 0xa8, 0xb9, 0x18, 0x9a, 0x5c,
	0x13, 0x58, 0x28, 0x34, 0x11, 0xca, 0xb0, 0xd8, 0xb0, 0xe6, 0x9c, 0x32, 0x83, 0x1f, 0x3f, 0x85,
	0xde, 0x2b, 0x7c, 0x79, 0x21, 0xfd, 0x3b, 0xb0, 0x53, 0x91, 0xab, 0x00, 0x74, 0x41, 0xbc, 0x97,
	0xaa, 0xa3, 0x7f, 0x04, 0xf7, 0xa8, 0x18, 0xe2, 0x5c, 0xd2, 0x31, 0x09, 0xdc, 0x2e, 0x88, 0x77,
	0xd3, 0x5d, 0x2a, 0xfa, 0x3a, 0x7e, 0xb2, 0xf3, 0xfb, 0x43, 0x04, 0x8e, 0x3f, 0x02, 0xb8, 0xf3,
	0x92, 0xe7, 0x95, 0xff, 0x0c, 0xde, 0xd2, 0xda, 0xdb, 0x21, 0x2e, 0x8a, 0x96, 0x08, 0x61, 0x1a,
	0x0d, 0x82, 0xf9, 0xa4, 0xb7, 0x6f, 0x05, 0xf5, 0x0d, 0x72, 0x2e, 0x5b, 0xca, 0xca, 0xf4, 0xa6,
	0xc9, 0xb7, 0x97, 0x6b, 0x7a, 0x77, 0x43, 0xdf, 0x87, 0x5e, 0xc3, 0xdf, 0x92, 0x36, 0xe8, 0xe8,
	0x4e, 0x0f, 0xa6, 0x8b, 0xc8, 0xf9, 0xbe, 0x88, 0xee, 0x99, 0x6e, 0xa2, 0xa8, 0x12, 0xca, 0x51,
	0x8d, 0xe5, 0x28, 0x39, 0x63, 0x72, 0x3e, 0xe9, 0x41, 0x4b, 0x73, 0xc6, 0x64, 0x6a, 0x2a, 0xad,
	0xc8, 0xaf, 0x00, 0x7a, 0xe7, 0x8a, 0x6c, 0x7b, 0x95, 0x6f, 0xa0, 0xa7, 0xbc, 0x13, 0x81, 0xdb,
	0xed, 0xc4, 0xd7, 0x1f, 0x1d, 0x24, 0xb6, 0x48, 0xb9, 0xbb, 0x1e, 0x5b, 0xf2, 0x9c, 0x53, 0x36,
	0x78, 0xac, 0xe4, 0x7e, 0xfe, 0x11, 0xc5, 0x25, 0x95, 0xa3, 0xcb, 0x2c, 0xc9, 0x79, 0x6d, 0x07,
	0x63, 0x7f, 0x3d, 0x51, 0x54, 0x76, 0xfc, 0xaa, 0x40, 0x7c, 0xfa, 0xf5, 0xe5, 0x3e, 0x48, 0x4d,
	0x7b, 0x2b, 0x9c, 0xc2, 0x1b, 0xca, 0xdc, 0x94, 0x88, 0x86, 0x33, 0x41, 0xfe, 0x33, 0xa2, 0xbf,
	0x1e, 0xb9, 0xdb, 0x79, 0x34, 0x78, 0x31, 0x5d, 0x86, 0x60, 0xb6, 0x0c, 0xc1, 0xcf, 0x65, 0x08,
	0xde, 0xaf, 0x42, 0x67, 0xb6, 0x0a, 0x9d, 0x6f, 0xab, 0xd0, 0x79, 0xfd, 0xf0, 0x9f, 0x07, 0x64,
	0x98, 0x15, 0x7a, 0x6d, 0x72, 0x7e, 0x81, 0xf2, 0x11, 0xa6, 0x0c, 0x8d, 0x4f, 0xd1, 0x3b, 0xb4,
	0xde, 0x6a, 0xfd, 0x9c, 0xec, 0x9a, 0x4e, 0x39, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x5a,
	0xcc, 0x88, 0xed, 0x02, 0x00, 0x00,
}

func (this *Vault) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Vault)
	if !ok {
		that2, ok := that.(Vault)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.IsActive != that1.IsActive {
		return false
	}
	return true
}
func (this *Lock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Lock)
	if !ok {
		that2, ok := that.(Lock)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.StakerAddress != that1.StakerAddress {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !this.Power.Equal(that1.Power) {
		return false
	}
	return true
}
func (this *Stake) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Stake)
	if !ok {
		that2, ok := that.(Stake)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.StakerAddress != that1.StakerAddress {
		return false
	}
	if len(this.Coins) != len(that1.Coins) {
		return false
	}
	for i := range this.Coins {
		if !this.Coins[i].Equal(&that1.Coins[i]) {
			return false
		}
	}
	return true
}
func (this *LockResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LockResponse)
	if !ok {
		that2, ok := that.(LockResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !this.Power.Equal(that1.Power) {
		return false
	}
	return true
}
func (m *Vault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Lock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Power.Size()
		i -= size
		if _, err := m.Power.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.StakerAddress) > 0 {
		i -= len(m.StakerAddress)
		copy(dAtA[i:], m.StakerAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.StakerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Stake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.StakerAddress) > 0 {
		i -= len(m.StakerAddress)
		copy(dAtA[i:], m.StakerAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.StakerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Power.Size()
		i -= size
		if _, err := m.Power.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.IsActive {
		n += 2
	}
	return n
}

func (m *Lock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StakerAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Power.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *Stake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StakerAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *LockResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Power.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Vault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Lock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Lock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Power.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Stake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Stake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *LockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: LockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Power.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)