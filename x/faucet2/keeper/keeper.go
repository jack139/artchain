package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/faucet2/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		SupplyKeeper  types.SupplyKeeper
		StakingKeeper types.StakingKeeper
		amount        int64              // set default amount for each mint.
		Limit         time.Duration      // rate limiting for mint, etc 24 * time.Hours
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	supplyKeeper types.SupplyKeeper,
	stakingKeeper types.StakingKeeper,
	amount int64,
	rateLimit time.Duration,
	// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		SupplyKeeper:  supplyKeeper,
		StakingKeeper: stakingKeeper,
		amount:        amount,
		Limit:         rateLimit,
		// this line is used by starport scaffolding # ibc/keeper/return
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
