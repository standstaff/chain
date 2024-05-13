package keeper

import (
	"context"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/bandprotocol/chain/v2/x/feeds/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the x/feeds MsgServer interface.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

// SubmitSignals register new signals and update feeds.
func (ms msgServer) SubmitSignals(
	goCtx context.Context,
	req *types.MsgSubmitSignals,
) (*types.MsgSubmitSignalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	delegator, err := sdk.AccAddressFromBech32(req.Delegator)
	if err != nil {
		return nil, err
	}

	// check whether delegator has enough delegation for signals
	err = ms.CheckDelegatorDelegation(ctx, delegator, req.Signals)
	if err != nil {
		return nil, err
	}

	// calculate power different of each signal by decresing signal power with previous signal
	signalIDToPowerDiff := ms.CalculateDelegatorSignalsPowerDiff(ctx, delegator, req.Signals)

	// sort keys to guarantee order of signalIDToPowerDiff iteration
	keys := make([]string, 0, len(signalIDToPowerDiff))
	for k := range signalIDToPowerDiff {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, signalID := range keys {
		powerDiff := signalIDToPowerDiff[signalID]
		feed, err := ms.GetFeed(ctx, signalID)
		if err != nil {
			feed = types.Feed{
				SignalID:                    signalID,
				Power:                       0,
				Interval:                    0,
				LastIntervalUpdateTimestamp: ctx.BlockTime().Unix(),
			}
		}
		feed.Power += powerDiff

		if feed.Power < 0 {
			return nil, types.ErrPowerNegative
		}

		feed.Interval, feed.DeviationInThousandth = calculateIntervalAndDeviation(feed.Power, ms.GetParams(ctx))
		ms.SetFeed(ctx, feed)
	}

	return &types.MsgSubmitSignalsResponse{}, nil
}

// SubmitPrices register new validator-prices.
func (ms msgServer) SubmitPrices(
	goCtx context.Context,
	req *types.MsgSubmitPrices,
) (*types.MsgSubmitPricesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	blockTime := ctx.BlockTime().Unix()

	val, err := sdk.ValAddressFromBech32(req.Validator)
	if err != nil {
		return nil, err
	}

	// check if it's in top bonded validators.
	err = ms.ValidateSubmitPricesRequest(ctx, blockTime, req, val)
	if err != nil {
		return nil, err
	}

	cooldownTime := ms.GetParams(ctx).CooldownTime

	for _, price := range req.Prices {
		valPrice, err := ms.NewValidatorPrice(ctx, blockTime, price, val, cooldownTime)
		if err != nil {
			return nil, err
		}

		err = ms.SetValidatorPrice(ctx, valPrice)
		if err != nil {
			return nil, err
		}

		emitEventSubmitPrice(ctx, valPrice)
	}

	return &types.MsgSubmitPricesResponse{}, nil
}

// UpdatePriceService updates price service.
func (ms msgServer) UpdatePriceService(
	goCtx context.Context,
	req *types.MsgUpdatePriceService,
) (*types.MsgUpdatePriceServiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	admin := ms.GetParams(ctx).Admin
	if admin != req.Admin {
		return nil, types.ErrInvalidSigner.Wrapf(
			"invalid admin; expected %s, got %s",
			admin,
			req.Admin,
		)
	}

	if err := ms.SetPriceService(ctx, req.PriceService); err != nil {
		return nil, err
	}

	emitEventUpdatePriceService(ctx, req.PriceService)

	return &types.MsgUpdatePriceServiceResponse{}, nil
}

// UpdateParams updates the feeds module params.
func (ms msgServer) UpdateParams(
	goCtx context.Context,
	req *types.MsgUpdateParams,
) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if ms.authority != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf(
			"invalid authority; expected %s, got %s",
			ms.authority,
			req.Authority,
		)
	}

	if err := ms.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	emitEventUpdateParams(ctx, req.Params)

	return &types.MsgUpdateParamsResponse{}, nil
}
