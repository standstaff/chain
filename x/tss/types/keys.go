package types

import (
	"github.com/bandprotocol/chain/v2/pkg/tss"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// module name
	ModuleName = "tss"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the tss module
	QuerierRoute = ModuleName
)

var (
	// GlobalStoreKeyPrefix is the prefix for global primitive state variables.
	GlobalStoreKeyPrefix = []byte{0x00}

	// GroupCountStoreKey is the key that keeps the total request count.
	GroupCountStoreKey = append(GlobalStoreKeyPrefix, []byte("GroupCount")...)

	// GroupStoreKeyPrefix is the prefix for group store.
	GroupStoreKeyPrefix = []byte{0x01}

	// DKGContextStoreKeyPrefix is the prefix for dkg context store.
	DKGContextStoreKeyPrefix = []byte{0x02}

	// MemberStoreKeyPrefix is the prefix for member store.
	MemberStoreKeyPrefix = []byte{0x03}

	// Round1Commitment is the key that keeps the member commitments on round 1.
	Round1CommitmentStoreKeyPrefix = []byte{0x04}
)

func GroupStoreKey(groupID tss.GroupID) []byte {
	return append(GroupStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(groupID))...)
}

func DKGContextStoreKey(groupID tss.GroupID) []byte {
	return append(DKGContextStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(groupID))...)
}

func MembersStoreKey(groupID tss.GroupID) []byte {
	return append(MemberStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(groupID))...)
}

func MemberOfGroupKey(groupID tss.GroupID, memberID tss.MemberID) []byte {
	buf := append(MemberStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(groupID))...)
	buf = append(buf, sdk.Uint64ToBigEndian(uint64(memberID))...)
	return buf
}

func Round1CommitmentStoreKey(groupID tss.GroupID) []byte {
	return append(Round1CommitmentStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(groupID))...)
}

func Round1CommitmentMemberStoreKey(groupID tss.GroupID, memberID tss.MemberID) []byte {
	buf := append(Round1CommitmentStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(groupID))...)
	buf = append(buf, sdk.Uint64ToBigEndian(uint64(memberID))...)
	return buf
}
