package faucet2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/faucet2/keeper"
	"github.com/jack139/artchain/x/faucet2/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the Mining
	for _, elem := range genState.MiningList {
		k.SetMining(ctx, *elem)
	}

	// Set Mining count
	k.SetMiningCount(ctx, uint64(len(genState.MiningList)))

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all Mining
	MiningList := k.GetAllMining(ctx)
	for _, elem := range MiningList {
		elem := elem
		genesis.MiningList = append(genesis.MiningList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
