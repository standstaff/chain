package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/pkg/tss"
)

const (
	// module name
	ModuleName = "bandtss"

	// StoreKey to be used when creating the KVStore.
	StoreKey = ModuleName

	// RouterKey is the message route for the bandtss module
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the bandtss module
	QuerierRoute = ModuleName
)

var (

	// GlobalStoreKeyPrefix is the prefix for global primitive state variables.
	GlobalStoreKeyPrefix = []byte{0x00}
	// ParamsKeyPrefix is a prefix for keys that store bandtss's parameters
	ParamsKeyPrefix = []byte{0x01}
	// MemberStoreKeyPrefix is the prefix for status store.
	MemberStoreKeyPrefix = []byte{0x02}
	// GroupIDStoreKeyPrefix is the prefix for groupID store.
	GroupIDStoreKeyPrefix = []byte{0x03}
	// SigningFeeStoreKeyPrefix is the prefix for signing fee store.
	SigningFeeStoreKeyPrefix = []byte{0x04}

	// CurrentGroupIDKey is the key for storing the current group ID under GroupIDStoreKeyPrefix.
	CurrentGroupIDStoreKey = append(GroupIDStoreKeyPrefix, []byte{0x00}...)
	// ReplacingGroupIDKey  is the key for storing the replacing group ID under GroupIDStoreKeyPrefix.
	ReplacingGroupIDStoreKey = append(GroupIDStoreKeyPrefix, []byte{0x01}...)
)

func MemberStoreKey(address sdk.AccAddress) []byte {
	return append(MemberStoreKeyPrefix, address...)
}

func SigningFeeStoreKey(signingID tss.SigningID) []byte {
	return append(SigningFeeStoreKeyPrefix, sdk.Uint64ToBigEndian(uint64(signingID))...)
}
