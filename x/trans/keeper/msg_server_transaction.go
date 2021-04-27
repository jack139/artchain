package keeper

import (
    "fmt"
	"context"

    "github.com/jack139/artchain/x/trans/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k msgServer) CreateTransaction(goCtx context.Context,  msg *types.MsgCreateTransaction) (*types.MsgCreateTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    id := k.AppendTransaction(
        ctx,
        msg.Creator,
        msg.RecType,
        msg.AuctionId,
        msg.ItemId,
        msg.TransType,
        msg.UserId,
        msg.TransDate,
        msg.HammerTime,
        msg.HammerPrice,
        msg.Details,
        msg.Status,
    )

	return &types.MsgCreateTransactionResponse{
	    Id: id,
	}, nil
}

func (k msgServer) UpdateTransaction(goCtx context.Context,  msg *types.MsgUpdateTransaction) (*types.MsgUpdateTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    var transaction = types.Transaction{
		Creator: msg.Creator,
		Id:      msg.Id,
    	RecType: msg.RecType,
    	AuctionId: msg.AuctionId,
    	ItemId: msg.ItemId,
    	TransType: msg.TransType,
    	UserId: msg.UserId,
    	TransDate: msg.TransDate,
    	HammerTime: msg.HammerTime,
    	HammerPrice: msg.HammerPrice,
    	Details: msg.Details,
    	Status: msg.Status,
	}

    // Checks that the element exists
    if !k.HasTransaction(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }

    // Checks if the the msg sender is the same as the current owner
    if msg.Creator != k.GetTransactionOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.SetTransaction(ctx, transaction)

	return &types.MsgUpdateTransactionResponse{}, nil
}

func (k msgServer) DeleteTransaction(goCtx context.Context,  msg *types.MsgDeleteTransaction) (*types.MsgDeleteTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    if !k.HasTransaction(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }
    if msg.Creator != k.GetTransactionOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.RemoveTransaction(ctx, msg.Id)

	return &types.MsgDeleteTransactionResponse{}, nil
}
