package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos-builders/chaos/x/amm/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) GetFeeRate(ctx sdk.Context) (feeRate sdk.Dec) {
	k.paramSpace.Get(ctx, types.KeyFeeRate, &feeRate)
	return
}

func (k Keeper) GetMinInitialLiquidity(ctx sdk.Context) (minInitialLiquidity sdk.Int) {
	k.paramSpace.Get(ctx, types.KeyMinInitialLiquidity, &minInitialLiquidity)
	return
}
