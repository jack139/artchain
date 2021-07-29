package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMining{}

func NewMsgCreateMining(creator string, Minter string, LastTime int64, Total string) *MsgCreateMining {
	return &MsgCreateMining{
		Creator:  creator,
		Minter:   Minter,
		LastTime: LastTime,
		Total:    Total,
	}
}

func (msg *MsgCreateMining) Route() string {
	return RouterKey
}

func (msg *MsgCreateMining) Type() string {
	return "CreateMining"
}

func (msg *MsgCreateMining) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMining) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMining) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMining{}

func NewMsgUpdateMining(creator string, id uint64, Minter string, LastTime int64, Total string) *MsgUpdateMining {
	return &MsgUpdateMining{
		Id:       id,
		Creator:  creator,
		Minter:   Minter,
		LastTime: LastTime,
		Total:    Total,
	}
}

func (msg *MsgUpdateMining) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMining) Type() string {
	return "UpdateMining"
}

func (msg *MsgUpdateMining) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMining) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMining) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMining{}

func NewMsgDeleteMining(creator string, id uint64) *MsgDeleteMining {
	return &MsgDeleteMining{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMining) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMining) Type() string {
	return "DeleteMining"
}

func (msg *MsgDeleteMining) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMining) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMining) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}


var _ sdk.Msg = &MsgMint{}

func NewMsgMint(sender string, minter string, time int64) *MsgMint {
	return &MsgMint{
		Sender:  sender,
		Minter:  minter,
		Time: time,
	}
}
func (msg *MsgMint) Route() string {
	return RouterKey
}

func (msg *MsgMint) Type() string {
	return "Mint"
}

func (msg *MsgMint) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
