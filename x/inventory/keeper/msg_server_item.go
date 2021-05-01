package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jack139/artchain/x/inventory/types"
)

func (k msgServer) CreateItem(goCtx context.Context, msg *types.MsgCreateItem) (*types.MsgCreateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendItem(
		ctx,
		msg.Creator,
		msg.RecType,
		msg.ItemDesc,
		msg.ItemDetail,
		msg.ItemDate,
		msg.ItemType,
		msg.ItemSubject,
		msg.ItemMedia,
		msg.ItemSize,
		msg.ItemImage,
		msg.AESKey,
		msg.ItemBasePrice,
		msg.CurrentOwnerId,
		msg.Status,
	)

	return &types.MsgCreateItemResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateItem(goCtx context.Context, msg *types.MsgUpdateItem) (*types.MsgUpdateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var item = types.Item{
		Creator:        msg.Creator,
		Id:             msg.Id,
		RecType:        msg.RecType,
		ItemDesc:       msg.ItemDesc,
		ItemDetail:     msg.ItemDetail,
		ItemDate:       msg.ItemDate,
		ItemType:       msg.ItemType,
		ItemSubject:    msg.ItemSubject,
		ItemMedia:      msg.ItemMedia,
		ItemSize:       msg.ItemSize,
		ItemImage:      msg.ItemImage,
		AESKey:         msg.AESKey,
		ItemBasePrice:  msg.ItemBasePrice,
		CurrentOwnerId: msg.CurrentOwnerId,
		Status:         msg.Status,
		LastDate:  time.Now().Format("2006-01-02 15:04:05"),
	}

	// Checks that the element exists
	if !k.HasItem(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetItemOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetItem(ctx, item)

	return &types.MsgUpdateItemResponse{}, nil
}

func (k msgServer) DeleteItem(goCtx context.Context, msg *types.MsgDeleteItem) (*types.MsgDeleteItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasItem(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetItemOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveItem(ctx, msg.Id)

	return &types.MsgDeleteItemResponse{}, nil
}
