package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateItem{}

func NewMsgCreateItem(creator string, recType string, itemDesc string, itemDetail string, itemDate string, itemType string, itemSubject string, itemMedia string, itemSize string, itemImage string, AESKey string, itemBasePrice string, currentOwnerId string) *MsgCreateItem {
	return &MsgCreateItem{
		Creator:        creator,
		RecType:        recType,
		ItemDesc:       itemDesc,
		ItemDetail:     itemDetail,
		ItemDate:       itemDate,
		ItemType:       itemType,
		ItemSubject:    itemSubject,
		ItemMedia:      itemMedia,
		ItemSize:       itemSize,
		ItemImage:      itemImage,
		AESKey:         AESKey,
		ItemBasePrice:  itemBasePrice,
		CurrentOwnerId: currentOwnerId,
	}
}

func (msg *MsgCreateItem) Route() string {
	return RouterKey
}

func (msg *MsgCreateItem) Type() string {
	return "CreateItem"
}

func (msg *MsgCreateItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateItem{}

func NewMsgUpdateItem(creator string, id uint64, recType string, itemDesc string, itemDetail string, itemDate string, itemType string, itemSubject string, itemMedia string, itemSize string, itemImage string, AESKey string, itemBasePrice string, currentOwnerId string) *MsgUpdateItem {
	return &MsgUpdateItem{
		Id:             id,
		Creator:        creator,
		RecType:        recType,
		ItemDesc:       itemDesc,
		ItemDetail:     itemDetail,
		ItemDate:       itemDate,
		ItemType:       itemType,
		ItemSubject:    itemSubject,
		ItemMedia:      itemMedia,
		ItemSize:       itemSize,
		ItemImage:      itemImage,
		AESKey:         AESKey,
		ItemBasePrice:  itemBasePrice,
		CurrentOwnerId: currentOwnerId,
	}
}

func (msg *MsgUpdateItem) Route() string {
	return RouterKey
}

func (msg *MsgUpdateItem) Type() string {
	return "UpdateItem"
}

func (msg *MsgUpdateItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateItem{}

func NewMsgDeleteItem(creator string, id uint64) *MsgDeleteItem {
	return &MsgDeleteItem{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteItem) Route() string {
	return RouterKey
}

func (msg *MsgDeleteItem) Type() string {
	return "DeleteItem"
}

func (msg *MsgDeleteItem) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
