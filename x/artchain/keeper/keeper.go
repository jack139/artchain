package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/artchain/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

type (
	Keeper struct {
		cdc      codec.Codec
		storeKey store.StoreKey
		memKey   store.StoreKey
	}
)

func NewKeeper(cdc codec.Codec, storeKey, memKey store.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
