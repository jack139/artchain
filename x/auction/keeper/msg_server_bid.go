package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jack139/artchain/x/auction/types"
)

func (k msgServer) CreateBid(goCtx context.Context, msg *types.MsgCreateBid) (*types.MsgCreateBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendBid(
		ctx,
		msg.Creator,
		msg.RecType,
		msg.AuctionId,
		msg.BidNo,
		msg.ItemId,
		msg.BuyerId,
		msg.BidPrice,
		msg.BidTime,
	)

	return &types.MsgCreateBidResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBid(goCtx context.Context, msg *types.MsgUpdateBid) (*types.MsgUpdateBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var bid = types.Bid{
		Creator:   msg.Creator,
		Id:        msg.Id,
		RecType:   msg.RecType,
		AuctionId: msg.AuctionId,
		BidNo:     msg.BidNo,
		ItemId:    msg.ItemId,
		BuyerId:   msg.BuyerId,
		BidPrice:  msg.BidPrice,
		BidTime:   msg.BidTime,
		LastDate:  time.Now().Format("2006-01-02 15:04:05"),
	}

	// Checks that the element exists
	if !k.HasBid(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetBidOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBid(ctx, bid)

	return &types.MsgUpdateBidResponse{}, nil
}

func (k msgServer) DeleteBid(goCtx context.Context, msg *types.MsgDeleteBid) (*types.MsgDeleteBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasBid(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetBidOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBid(ctx, msg.Id)

	return &types.MsgDeleteBidResponse{}, nil
}
