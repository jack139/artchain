package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/artchain/x/inventory/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ReviewAll(c context.Context, req *types.QueryAllReviewRequest) (*types.QueryAllReviewResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var reviews []*types.Review
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	reviewStore := prefix.NewStore(store, types.KeyPrefix(types.ReviewKey+req.ItemId))

	pageRes, err := query.FilteredPaginate(reviewStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var review types.Review
		if err := k.cdc.UnmarshalBinaryBare(value, &review); err != nil {
			return false, err
		}

		// filter 
		if strings.Contains(req.Status, review.Status){ // 状态可以多个
			if accumulate {
				reviews = append(reviews, &review)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllReviewResponse{Review: reviews, Pagination: pageRes}, nil
}

func (k Keeper) Review(c context.Context, req *types.QueryGetReviewRequest) (*types.QueryGetReviewResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var review types.Review
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasReview(ctx, req.Id, req.ItemId) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey+req.ItemId))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetReviewIDBytes(req.Id)), &review)

	return &types.QueryGetReviewResponse{Review: &review}, nil
}


// 使用 FilteredPaginate 版本
func (k Keeper) ReviewByStatus(c context.Context, req *types.QueryGetReviewByStatusRequest) (*types.QueryGetReviewByStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var reviews []*types.Review
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	reviewStore := prefix.NewStore(store, types.KeyPrefix(types.ReviewKey))

	pageRes, err := query.FilteredPaginate(reviewStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var review types.Review
		if err := k.cdc.UnmarshalBinaryBare(value, &review); err != nil {
			return false, err
		}

		// filter 
		if strings.Contains(req.Status, review.Status){ // 状态可以多个
			if accumulate {
				reviews = append(reviews, &review)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetReviewByStatusResponse{Review: reviews, Pagination: pageRes}, nil
}
