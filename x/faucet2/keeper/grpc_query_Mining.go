package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/artchain/x/faucet2/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MiningAll(c context.Context, req *types.QueryAllMiningRequest) (*types.QueryAllMiningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var Minings []*types.Mining
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	MiningStore := prefix.NewStore(store, types.KeyPrefix(types.MiningKey))

	pageRes, err := query.Paginate(MiningStore, req.Pagination, func(key []byte, value []byte) error {
		var Mining types.Mining
		if err := k.cdc.UnmarshalBinaryBare(value, &Mining); err != nil {
			return err
		}

		Minings = append(Minings, &Mining)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMiningResponse{Mining: Minings, Pagination: pageRes}, nil
}

func (k Keeper) Mining(c context.Context, req *types.QueryGetMiningRequest) (*types.QueryGetMiningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var Mining types.Mining
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasMining(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetMiningIDBytes(req.Id)), &Mining)

	return &types.QueryGetMiningResponse{Mining: &Mining}, nil
}
