package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/artchain/x/person/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserAll(c context.Context, req *types.QueryAllUserRequest) (*types.QueryAllUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []*types.User
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))

	pageRes, err := query.Paginate(userStore, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return err
		}

		users = append(users, &user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserResponse{User: users, Pagination: pageRes}, nil
}

func (k Keeper) User(c context.Context, req *types.QueryGetUserRequest) (*types.QueryGetUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var user types.User
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasUser(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	k.cdc.MustUnmarshal(store.Get(GetUserIDBytes(req.Id)), &user)

	return &types.QueryGetUserResponse{User: &user}, nil
}


func (k Keeper) UserByChainAddr(c context.Context, req *types.QueryGetUserByChainAddrRequest) (*types.QueryGetUserByChainAddrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var user types.User
	ctx := sdk.UnwrapSDKContext(c)

	user = k.GetUserByChainAddr(ctx, req.ChainAddr)

	return &types.QueryGetUserByChainAddrResponse{User: &user}, nil
}

/* 按用户类型查询列表 */
func (k Keeper) UserByUserType(c context.Context, req *types.QueryGetUserByUserTypeRequest) (*types.QueryGetUserByUserTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []*types.User
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))

	pageRes, err := query.FilteredPaginate(userStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return false, err
		}

		// filter 
		if strings.HasPrefix(user.UserType, req.UserType) {
			if accumulate {
				users = append(users, &user)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetUserByUserTypeResponse{User: users, Pagination: pageRes}, nil
}

// 使用 FilteredPaginate 版本
func (k Keeper) UserByStatus(c context.Context, req *types.QueryGetUserByStatusRequest) (*types.QueryGetUserByStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []*types.User
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))

	pageRes, err := query.FilteredPaginate(userStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var user types.User
		if err := k.cdc.Unmarshal(value, &user); err != nil {
			return false, err
		}

		// filter 
		if strings.Contains(req.Status, user.Status){ // 状态可以多个
			if accumulate {
				users = append(users, &user)
			}
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetUserByStatusResponse{User: users, Pagination: pageRes}, nil
}
