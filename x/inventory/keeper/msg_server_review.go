package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jack139/artchain/x/inventory/types"
)

func (k msgServer) CreateReview(goCtx context.Context, msg *types.MsgCreateReview) (*types.MsgCreateReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendReview(
		ctx,
		msg.Creator,
		msg.RecType,
		msg.ItemId,
		msg.ReviewerId,
		msg.ReviewDetail,
		msg.ReviewDate,
		msg.UpCount,
		msg.DownCount,
		msg.Status,
	)

	return &types.MsgCreateReviewResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateReview(goCtx context.Context, msg *types.MsgUpdateReview) (*types.MsgUpdateReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var review = types.Review{
		Creator:      msg.Creator,
		Id:           msg.Id,
		RecType:      msg.RecType,
		ItemId:       msg.ItemId,
		ReviewerId:   msg.ReviewerId,
		ReviewDetail: msg.ReviewDetail,
		ReviewDate:   msg.ReviewDate,
		UpCount:      msg.UpCount,
		DownCount:    msg.DownCount,
		LastDate:  time.Now().Format("2006-01-02 15:04:05"),
	}

	// Checks that the element exists
	if !k.HasReview(ctx, msg.Id, msg.ItemId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetReviewOwner(ctx, msg.Id, msg.ItemId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetReview(ctx, review)

	return &types.MsgUpdateReviewResponse{}, nil
}

func (k msgServer) DeleteReview(goCtx context.Context, msg *types.MsgDeleteReview) (*types.MsgDeleteReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasReview(ctx, msg.Id, msg.ItemId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetReviewOwner(ctx, msg.Id, msg.ItemId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveReview(ctx, msg.Id, msg.ItemId)

	return &types.MsgDeleteReviewResponse{}, nil
}
