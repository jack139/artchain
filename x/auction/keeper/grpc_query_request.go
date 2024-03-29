package keeper

import (
	"context"
	"strings"

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

// 按 seller_addr 查询 拍卖请求清单,  FilteredPaginate 版本
func (k Keeper) RequestByChainAddr(c context.Context, req *types.QueryGetRequestByChainAddrRequest) (*types.QueryGetRequestByChainAddrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var requests []*types.Request
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	requestStore := prefix.NewStore(store, types.KeyPrefix(types.RequestKey))

	pageRes, err := query.FilteredPaginate(requestStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var request types.Request
		if err := k.cdc.UnmarshalBinaryBare(value, &request); err != nil {
			return false, err
		}

		// filter 
		if request.SellerId == req.ChainAddr {
			if accumulate {
				requests = append(requests, &request)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetRequestByChainAddrResponse{Request: requests, Pagination: pageRes}, nil
}

// 使用 FilteredPaginate 版本
func (k Keeper) RequestByStatus(c context.Context, req *types.QueryGetRequestByStatusRequest) (*types.QueryGetRequestByStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var requests []*types.Request
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	requestStore := prefix.NewStore(store, types.KeyPrefix(types.RequestKey))

	pageRes, err := query.FilteredPaginate(requestStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var request types.Request
		if err := k.cdc.UnmarshalBinaryBare(value, &request); err != nil {
			return false, err
		}

		// filter 
		if strings.Contains(req.Status, request.Status){ // 状态可以多个
			if accumulate {
				requests = append(requests, &request)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetRequestByStatusResponse{Request: requests, Pagination: pageRes}, nil
}
