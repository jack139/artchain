package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jack139/artchain/x/person/types"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendUser(
		ctx,
		msg.Creator,
		msg.RecType,
		msg.Name,
		msg.UserType,
		msg.UserInfo,
		msg.Status,
		msg.RegDate,
		msg.ChainAddr,
	)

	return &types.MsgCreateUserResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateUser(goCtx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var user = types.User{
		Creator:   msg.Creator,
		Id:        msg.Id,
		RecType:   msg.RecType,
		Name:      msg.Name,
		UserType:  msg.UserType,
		UserInfo:  msg.UserInfo,
		Status:    msg.Status,
		RegDate:   msg.RegDate,
		ChainAddr: msg.ChainAddr,
	}

	// Checks that the element exists
	if !k.HasUser(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetUserOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetUser(ctx, user)

	return &types.MsgUpdateUserResponse{}, nil
}

func (k msgServer) DeleteUser(goCtx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasUser(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetUserOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUser(ctx, msg.Id)

	return &types.MsgDeleteUserResponse{}, nil
}
