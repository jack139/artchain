package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
		msg.LastDate,
	)

	return &types.MsgCreateItemResponse{
		Id: id,
	}, nil
}

/*
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
		LastDate:       msg.LastDate,
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
*/

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

// 更新 item, 只更新有内容的，无内容的填 "\x00"
// lastDate 只填要更新的串，会自动追加到原有的末尾
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
		LastDate:       msg.LastDate,
	}

	// Checks that the element exists
	if !k.HasItem(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetItemOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// 取当前值
	oldItem := k.GetItem(ctx, item.Id)

	// 填充
	if item.Creator == "\x00" {
		item.Creator = oldItem.Creator
	}
	if item.RecType == "\x00" {
		item.RecType = oldItem.RecType
	}
	if item.ItemDesc == "\x00" {
		item.ItemDesc = oldItem.ItemDesc
	}
	if item.ItemDetail == "\x00" {
		item.ItemDetail = oldItem.ItemDetail
	}
	if item.ItemDate == "\x00" {
		item.ItemDate = oldItem.ItemDate
	}
	if item.ItemType == "\x00" {
		item.ItemType = oldItem.ItemType
	}
	if item.ItemSubject == "\x00" {
		item.ItemSubject = oldItem.ItemSubject
	}
	if item.ItemMedia == "\x00" {
		item.ItemMedia = oldItem.ItemMedia
	}
	if item.ItemSize == "\x00" {
		item.ItemSize = oldItem.ItemSize
	}
	if item.ItemImage == "\x00" {
		item.ItemImage = oldItem.ItemImage
	}
	if item.AESKey == "\x00" {
		item.AESKey = oldItem.AESKey
	}
	if item.ItemBasePrice == "\x00" {
		item.ItemBasePrice = oldItem.ItemBasePrice
	}
	if item.CurrentOwnerId == "\x00" {
		item.CurrentOwnerId = oldItem.CurrentOwnerId
	}
	if item.Status == "\x00" {
		item.Status = oldItem.Status
	}

	// lastDate 反序列化
	var lastDateMap []map[string]interface{}
	if err := json.Unmarshal([]byte(oldItem.LastDate), &lastDateMap); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidSequence, err.Error())
	}

	// 构建 lastDate
	s := strings.Split(item.LastDate, "|") // 格式：caller|text
	lastDateMap = append(lastDateMap, map[string]interface{}{
		"caller": s[0],
		"act":    s[1],
		"date":   time.Now().Format("2006-01-02 15:04:05"),
	})
	lastDate, err := json.Marshal(lastDateMap)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidSequence, err.Error())
	}

	item.LastDate = string(lastDate)

	// 保存
	k.SetItem(ctx, item)

	return &types.MsgUpdateItemResponse{}, nil
}
