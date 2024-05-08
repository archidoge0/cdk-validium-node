package keeper

import (
	"github.com/RiemaLabs/nubit-validator/utils/tk/sh/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the x/mint store with data from the genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, ak types.AccountKeeper, data *types.GenesisState) {
	minter := types.DefaultMinter()
	minter.BondDenom = data.BondDenom
	k.SetMinter(ctx, minter)
	// override the genesis time with the actual genesis time supplied in `InitChain`
	blockTime := ctx.BlockTime()
	gt := types.GenesisTime{
		GenesisTime: &blockTime,
	}
	k.SetGenesisTime(ctx, gt)
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns a x/mint GenesisState for the given context.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	bondDenom := k.GetMinter(ctx).BondDenom
	return types.NewGenesisState(bondDenom)
}