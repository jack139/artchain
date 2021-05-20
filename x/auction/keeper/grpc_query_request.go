package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/artchain/x/auction/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RequestAll(c context.Context, req *types.QueryAllRequestRequest) (*types.QueryAllRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var requests []*types.Request
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	requestStore := prefix.NewStore(store, types.KeyPrefix(types.RequestKey))

	pageRes, err := query.Paginate(requestStore, req.Pagination, func(key []byte, value []byte) error {
		var request types.Request
		if err := k.cdc.UnmarshalBinaryBare(value, &request); err != nil {
			return err
		}

		requests = append(requests, &request)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRequestResponse{Request: requests, Pagination: pageRes}, nil
}

func (k Keeper) Request(c context.Context, req *types.QueryGetRequestRequest) (*types.QueryGetRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var request types.Request
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasRequest(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetRequestIDBytes(req.Id)), &request)

	return &types.QueryGetRequestResponse{Request: &request}, nil
}

// 按 seller_addr 查询 拍卖请求清单
func (k Keeper) RequestByChainAddr(c context.Context, req *types.QueryGetRequestByChainAddrRequest) (*types.QueryGetRequestByChainAddrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var requests []*types.Request
	ctx := sdk.UnwrapSDKContext(c)

	r := k.GetRequestByChainAddr(ctx, req.ChainAddr)
	for i, _ := range r{
		requests = append(requests, &r[i])
	}

	return &types.QueryGetRequestByChainAddrResponse{Request: requests}, nil
}
