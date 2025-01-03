package keeper

import (
	"fmt"

	dbm "github.com/cosmos/cosmos-db"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v3/pkg/tss"
	"github.com/bandprotocol/chain/v3/x/bandtss/types"
)

// DeactivateMember flags is_active to false. This function will return error if the given address
// isn't the member of the given group.
func (k Keeper) DeactivateMember(ctx sdk.Context, address sdk.AccAddress, groupID tss.GroupID) error {
	member, err := k.GetMember(ctx, address, groupID)
	if err != nil {
		return err
	}

	if !member.IsActive {
		return nil
	}

	member.IsActive = false
	member.Since = ctx.BlockTime()
	k.SetMember(ctx, member)

	if err := k.tssKeeper.DeactivateMember(ctx, groupID, address); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeInactiveStatus,
		sdk.NewAttribute(types.AttributeKeyAddress, address.String()),
		sdk.NewAttribute(types.AttributeKeyGroupID, fmt.Sprintf("%d", member.GroupID)),
	))

	return nil
}

// ActivateMember activates the member. This function returns an error if the given member is too
// soon to activate or the member is not in the given group.
func (k Keeper) ActivateMember(ctx sdk.Context, address sdk.AccAddress, groupID tss.GroupID) error {
	member, err := k.GetMember(ctx, address, groupID)
	if err != nil {
		return err
	}

	if member.IsActive {
		return types.ErrMemberAlreadyActive.Wrapf("address %s is already active", address.String())
	}

	penaltyEndTime := member.Since.Add(k.GetParams(ctx).InactivePenaltyDuration)
	if penaltyEndTime.After(ctx.BlockTime()) {
		return types.ErrPenaltyDurationNotElapsed.Wrapf("penalty end time: %s", penaltyEndTime)
	}

	member.IsActive = true
	member.Since = ctx.BlockTime()
	k.SetMember(ctx, member)

	if err := k.tssKeeper.ActivateMember(ctx, groupID, address); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeActivate,
		sdk.NewAttribute(types.AttributeKeyAddress, address.String()),
		sdk.NewAttribute(types.AttributeKeyGroupID, fmt.Sprintf("%d", member.GroupID)),
	))

	return nil
}

// DeleteMembers removes all members of the group.
func (k Keeper) DeleteMembers(ctx sdk.Context, groupID tss.GroupID) {
	members := k.tssKeeper.MustGetMembers(ctx, groupID)
	for _, m := range members {
		k.DeleteMember(ctx, sdk.MustAccAddressFromBech32(m.Address), groupID)

		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeMemberDeleted,
			sdk.NewAttribute(types.AttributeKeyAddress, m.Address),
			sdk.NewAttribute(types.AttributeKeyGroupID, fmt.Sprintf("%d", groupID)),
		))
	}
}

// AddMembers adds all members of the group.
func (k Keeper) AddMembers(ctx sdk.Context, groupID tss.GroupID) error {
	members := k.tssKeeper.MustGetMembers(ctx, groupID)
	for _, m := range members {
		addr := sdk.MustAccAddressFromBech32(m.Address)
		if err := k.AddMember(ctx, addr, groupID); err != nil {
			return err
		}
	}

	return nil
}

// AddMember adds a new member to the group and return error if already exists
func (k Keeper) AddMember(ctx sdk.Context, address sdk.AccAddress, groupID tss.GroupID) error {
	if k.HasMember(ctx, address, groupID) {
		return types.ErrMemberAlreadyExists.Wrapf("address %s already exists", address.String())
	}

	member := types.NewMember(address, groupID, true, ctx.BlockTime())
	k.SetMember(ctx, member)

	return nil
}

// =====================================
// Member store
// =====================================

// SetMember sets a member information in the store.
func (k Keeper) SetMember(ctx sdk.Context, member types.Member) {
	address := sdk.MustAccAddressFromBech32(member.Address)
	ctx.KVStore(k.storeKey).Set(types.MemberStoreKey(address, member.GroupID), k.cdc.MustMarshal(&member))
}

// HasMember checks that address is in the store or not.
func (k Keeper) HasMember(ctx sdk.Context, address sdk.AccAddress, groupID tss.GroupID) bool {
	return ctx.KVStore(k.storeKey).Has(types.MemberStoreKey(address, groupID))
}

// GetMembersIterator gets an iterator all statuses of address.
func (k Keeper) GetMembersIterator(ctx sdk.Context) dbm.Iterator {
	return storetypes.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.MemberStoreKeyPrefix)
}

// GetMember retrieves a member by address.
func (k Keeper) GetMember(ctx sdk.Context, address sdk.AccAddress, groupID tss.GroupID) (types.Member, error) {
	if !k.HasMember(ctx, address, groupID) {
		return types.Member{}, types.ErrMemberNotFound.Wrapf("address: %s", address)
	}
	bz := ctx.KVStore(k.storeKey).Get(types.MemberStoreKey(address, groupID))

	var member types.Member
	k.cdc.MustUnmarshal(bz, &member)
	return member, nil
}

// GetMembers retrieves all statuses of the store.
func (k Keeper) GetMembers(ctx sdk.Context) []types.Member {
	iterator := k.GetMembersIterator(ctx)
	defer iterator.Close()

	var members []types.Member
	for ; iterator.Valid(); iterator.Next() {
		var status types.Member
		k.cdc.MustUnmarshal(iterator.Value(), &status)
		members = append(members, status)
	}

	return members
}

// DeleteMember removes the status of the address of the group
func (k Keeper) DeleteMember(ctx sdk.Context, address sdk.AccAddress, groupID tss.GroupID) {
	ctx.KVStore(k.storeKey).Delete(types.MemberStoreKey(address, groupID))
}
