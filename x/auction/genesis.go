package auction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/auction/keeper"
	"github.com/jack139/artchain/x/auction/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the bid
	for _, elem := range genState.BidList {
		k.SetBid(ctx, *elem)
	}

	// Set bid count
	k.SetBidCount(ctx, uint64(len(genState.BidList)))

	// Set all the request
	for _, elem := range genState.RequestList {
		k.SetRequest(ctx, *elem)
	}

	// Set request count
	k.SetRequestCount(ctx, uint64(len(genState.RequestList)))

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all bid
	bidList := k.GetAllBid(ctx)
	for _, elem := range bidList {
		elem := elem
		genesis.BidList = append(genesis.BidList, &elem)
	}

	// Get all request
	requestList := k.GetAllRequest(ctx)
	for _, elem := range requestList {
		elem := elem
		genesis.RequestList = append(genesis.RequestList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
