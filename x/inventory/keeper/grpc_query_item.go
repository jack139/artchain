package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/artchain/x/inventory/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ItemAll(c context.Context, req *types.QueryAllItemRequest) (*types.QueryAllItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var items []*types.Item
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	itemStore := prefix.NewStore(store, types.KeyPrefix(types.ItemKey))

	pageRes, err := query.Paginate(itemStore, req.Pagination, func(key []byte, value []byte) error {
		var item types.Item
		if err := k.cdc.UnmarshalBinaryBare(value, &item); err != nil {
			return err
		}

		items = append(items, &item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllItemResponse{Item: items, Pagination: pageRes}, nil
}

func (k Keeper) Item(c context.Context, req *types.QueryGetItemRequest) (*types.QueryGetItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var item types.Item
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasItem(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetItemIDBytes(req.Id)), &item)

	return &types.QueryGetItemResponse{Item: &item}, nil
}


func (k Keeper) ItemAllByOwner(c context.Context, req *types.QueryAllItemByOwnerRequest) (*types.QueryAllItemByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var items []*types.Item
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	itemStore := prefix.NewStore(store, types.KeyPrefix(types.ItemKey))
	itemOwnerStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemOwnerKey+req.CurrentOwnerId))

	// 从 ItemOwnerKey 索引中取索引数据
	// 再从 ItemKey 中取实际数据
	pageRes, err := query.Paginate(itemOwnerStore, req.Pagination, func(key []byte, value []byte) error {
		var item types.Item
		if err := k.cdc.UnmarshalBinaryBare(itemStore.Get(value), &item); err != nil {
			return err
		}

		items = append(items, &item)
		
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllItemByOwnerResponse{Item: items, Pagination: pageRes}, nil
}
