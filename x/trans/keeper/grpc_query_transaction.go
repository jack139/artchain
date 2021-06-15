package keeper

import (
	"context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/artchain/x/trans/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TransactionAll(c context.Context, req *types.QueryAllTransactionRequest) (*types.QueryAllTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var transactions []*types.Transaction
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	transactionStore := prefix.NewStore(store, types.KeyPrefix(types.TransactionKey))

	pageRes, err := query.Paginate(transactionStore, req.Pagination, func(key []byte, value []byte) error {
		var transaction types.Transaction
		if err := k.cdc.UnmarshalBinaryBare(value, &transaction); err != nil {
			return err
		}

		transactions = append(transactions, &transaction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTransactionResponse{Transaction: transactions, Pagination: pageRes}, nil
}

func (k Keeper) Transaction(c context.Context, req *types.QueryGetTransactionRequest) (*types.QueryGetTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var transaction types.Transaction
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasTransaction(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetTransactionIDBytes(req.Id)), &transaction)

	return &types.QueryGetTransactionResponse{Transaction: &transaction}, nil
}

// 使用 FilteredPaginate 版本
// cate: seller, buyer, item, status
func (k Keeper) TransactionSome(c context.Context, req *types.QuerySomeTransactionRequest) (*types.QuerySomeTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var transactions []*types.Transaction
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	transactionStore := prefix.NewStore(store, types.KeyPrefix(types.TransactionKey))

	pageRes, err := query.FilteredPaginate(transactionStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var transaction types.Transaction
		if err := k.cdc.UnmarshalBinaryBare(value, &transaction); err != nil {
			return false, err
		}

		// filter 
		switch req.Cate {
		case "status":
			if strings.Contains(req.Condition, transaction.Status){ // 状态可以多个
				if accumulate {
					transactions = append(transactions, &transaction)
				}
				return true, nil
			}
		case "seller":
			if transaction.SellerId == req.Condition {
				if accumulate {
					transactions = append(transactions, &transaction)
				}
				return true, nil
			}
		case "buyer":
			if transaction.BuyerId == req.Condition {
				if accumulate {
					transactions = append(transactions, &transaction)
				}
				return true, nil
			}
		case "item":
			if transaction.ItemId == req.Condition {
				if accumulate {
					transactions = append(transactions, &transaction)
				}
				return true, nil
			}
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySomeTransactionResponse{Transaction: transactions, Pagination: pageRes}, nil
}
