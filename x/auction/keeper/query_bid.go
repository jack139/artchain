package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

func listBid(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino, auctionId string,) ([]byte, error) {
	msgs := keeper.GetAllBid(ctx, auctionId)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, msgs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func getBid(ctx sdk.Context, key string, auctionId string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	id, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return nil, err
	}

	if !keeper.HasBid(ctx, id, auctionId) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	msg := keeper.GetBid(ctx, id, auctionId)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
