package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/inventory/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	// this line is used by starport scaffolding # ibc/keeper/import
	//"github.com/cosmos/modules/incubator/nft"
	//nft "github.com/irisnet/irismod/modules/nft/keeper"
	nft "github.com/cosmos/cosmos-sdk/x/nft/keeper"
)

type (
	Keeper struct {
		cdc      codec.Codec
		storeKey store.StoreKey
		memKey   store.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute
		NFTKeeper nft.Keeper
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey,
	memKey store.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	nftKeeper nft.Keeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		NFTKeeper: nftKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
