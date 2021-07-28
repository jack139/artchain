package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jack139/artchain/x/faucet2/types"
)

func (k msgServer) CreateMining(goCtx context.Context, msg *types.MsgCreateMining) (*types.MsgCreateMiningResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendMining(
		ctx,
		msg.Creator,
		msg.Minter,
		msg.LastTime,
		msg.Total,
	)

	return &types.MsgCreateMiningResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMining(goCtx context.Context, msg *types.MsgUpdateMining) (*types.MsgUpdateMiningResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var Mining = types.Mining{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Minter:   msg.Minter,
		LastTime: msg.LastTime,
		Total:    msg.Total,
	}

	// Checks that the element exists
	if !k.HasMining(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetMiningOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMining(ctx, Mining)

	return &types.MsgUpdateMiningResponse{}, nil
}

func (k msgServer) DeleteMining(goCtx context.Context, msg *types.MsgDeleteMining) (*types.MsgDeleteMiningResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasMining(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetMiningOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMining(ctx, msg.Id)

	return &types.MsgDeleteMiningResponse{}, nil
}

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Error("Mint in keeper")

	return &types.MsgMintResponse{}, nil
}
