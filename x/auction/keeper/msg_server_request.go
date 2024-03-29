package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jack139/artchain/x/auction/types"
)

func (k msgServer) CreateRequest(goCtx context.Context, msg *types.MsgCreateRequest) (*types.MsgCreateRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendRequest(
		ctx,
		msg.Creator,
		msg.RecType,
		msg.ItemId,
		msg.AuctionHouseId,
		msg.SellerId,
		msg.RequestDate,
		msg.ReservePrice,
		msg.Status,
		msg.OpenDate,
		msg.CloseDate,
		msg.LastDate,
	)

	return &types.MsgCreateRequestResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateRequest(goCtx context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var request = types.Request{
		Creator:        msg.Creator,
		Id:             msg.Id,
		RecType:        msg.RecType,
		ItemId:         msg.ItemId,
		AuctionHouseId: msg.AuctionHouseId,
		SellerId:       msg.SellerId,
		RequestDate:    msg.RequestDate,
		ReservePrice:   msg.ReservePrice,
		Status:         msg.Status,
		OpenDate:       msg.OpenDate,
		CloseDate:      msg.CloseDate,
		LastDate:       msg.LastDate,
	}

	// Checks that the element exists
	if !k.HasRequest(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetRequestOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetRequest(ctx, request)

	return &types.MsgUpdateRequestResponse{}, nil
}

func (k msgServer) DeleteRequest(goCtx context.Context, msg *types.MsgDeleteRequest) (*types.MsgDeleteRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasRequest(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetRequestOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRequest(ctx, msg.Id)

	return &types.MsgDeleteRequestResponse{}, nil
}
