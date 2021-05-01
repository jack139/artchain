package inventory

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/inventory/keeper"
	"github.com/jack139/artchain/x/inventory/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the review
	for _, elem := range genState.ReviewList {
		k.SetReview(ctx, *elem)
	}

	// Set review count
	k.SetReviewCount(ctx, uint64(len(genState.ReviewList)))

	// Set all the item
	for _, elem := range genState.ItemList {
		k.SetItem(ctx, *elem)
	}

	// Set item count
	k.SetItemCount(ctx, uint64(len(genState.ItemList)))

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all review
	reviewList := k.GetAllReview(ctx, "") // 其实 不能导出 任何东西，因为没有提供 itemId
	for _, elem := range reviewList {
		elem := elem
		genesis.ReviewList = append(genesis.ReviewList, &elem)
	}

	// Get all item
	itemList := k.GetAllItem(ctx)
	for _, elem := range itemList {
		elem := elem
		genesis.ItemList = append(genesis.ItemList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
