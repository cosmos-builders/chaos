package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos-builders/chaos/x/amm/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	k.SetLastPairID(ctx, genState.LastPairId)
	for _, pair := range genState.Pairs {
		k.SetPair(ctx, pair)
		k.SetPairIndex(ctx, pair)
	}
}

// ExportGenesis returns the module's exported genesis
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetParams(ctx),
		k.GetLastPairID(ctx),
		k.GetAllPairs(ctx),
	)
}
