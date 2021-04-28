package trans

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/trans/keeper"
	"github.com/jack139/artchain/x/trans/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the transaction
	for _, elem := range genState.TransactionList {
		k.SetTransaction(ctx, *elem)
	}

	// Set transaction count
	k.SetTransactionCount(ctx, uint64(len(genState.TransactionList)))

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all transaction
	transactionList := k.GetAllTransaction(ctx)
	for _, elem := range transactionList {
		elem := elem
		genesis.TransactionList = append(genesis.TransactionList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
