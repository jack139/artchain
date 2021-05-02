package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTransaction{}

func NewMsgCreateTransaction(creator string, recType string, auctionId string, itemId string, transType string, 
	userId string, transDate string, hammerTime string, hammerPrice string, details string, status string) *MsgCreateTransaction {
	return &MsgCreateTransaction{
		Creator:     creator,
		RecType:     recType,
		AuctionId:   auctionId,
		ItemId:      itemId,
		TransType:   transType,
		UserId:      userId,
		TransDate:   transDate,
		HammerTime:  hammerTime,
		HammerPrice: hammerPrice,
		Details:     details,
		Status:      status,
	}
}

func (msg *MsgCreateTransaction) Route() string {
	return RouterKey
}

func (msg *MsgCreateTransaction) Type() string {
	return "CreateTransaction"
}

func (msg *MsgCreateTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTransaction{}

func NewMsgUpdateTransaction(creator string, id uint64, recType string, auctionId string, itemId string, transType string, userId string, transDate string, hammerTime string, hammerPrice string, details string, status string) *MsgUpdateTransaction {
	return &MsgUpdateTransaction{
		Id:          id,
		Creator:     creator,
		RecType:     recType,
		AuctionId:   auctionId,
		ItemId:      itemId,
		TransType:   transType,
		UserId:      userId,
		TransDate:   transDate,
		HammerTime:  hammerTime,
		HammerPrice: hammerPrice,
		Details:     details,
		Status:      status,
	}
}

func (msg *MsgUpdateTransaction) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTransaction) Type() string {
	return "UpdateTransaction"
}

func (msg *MsgUpdateTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateTransaction{}

func NewMsgDeleteTransaction(creator string, id uint64) *MsgDeleteTransaction {
	return &MsgDeleteTransaction{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteTransaction) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTransaction) Type() string {
	return "DeleteTransaction"
}

func (msg *MsgDeleteTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
