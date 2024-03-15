package tss

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/tss/keeper"
	"github.com/bandprotocol/chain/v2/x/tss/types"
)

// InitGenesis performs genesis initialization for the tss module.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, data *types.GenesisState) {
	if err := k.SetParams(ctx, data.Params); err != nil {
		panic(err)
	}

	k.SetGroupCount(ctx, data.GroupCount)
	for _, group := range data.Groups {
		k.SetGroup(ctx, group)
	}

	for _, member := range data.Members {
		k.SetMember(ctx, member)
	}

	k.SetSigningCount(ctx, data.SigningCount)
	for _, signing := range data.Signings {
		k.SetSigning(ctx, signing)
	}

	k.SetReplacementCount(ctx, data.ReplacementCount)
	for _, rep := range data.Replacements {
		k.SetReplacement(ctx, rep)
	}

	for _, deq := range data.DEQueues {
		k.SetDEQueue(ctx, deq)
	}

	for _, de := range data.DEsGenesis {
		address := sdk.MustAccAddressFromBech32(de.Address)
		k.SetDE(ctx, address, de.Index, de.DE)
	}

	for _, status := range data.IsActivesGenesis {
		k.SetMemberIsActive(ctx, sdk.AccAddress(status.Address), status.IsActive)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:           k.GetParams(ctx),
		GroupCount:       k.GetGroupCount(ctx),
		Groups:           k.GetGroups(ctx),
		Members:          k.GetMembers(ctx),
		SigningCount:     k.GetSigningCount(ctx),
		Signings:         k.GetSignings(ctx),
		ReplacementCount: k.GetReplacementCount(ctx),
		Replacements:     k.GetReplacements(ctx),
		DEQueues:         k.GetDEQueues(ctx),
		DEsGenesis:       k.GetDEsGenesis(ctx),
		IsActivesGenesis: k.GetMemberIsActivesGenesis(ctx),
	}
}
