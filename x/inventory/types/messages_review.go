package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateReview{}

func NewMsgCreateReview(creator string, recType string, itemId string, reviewerId string, reviewDetail string, 
	reviewDate string, upCount string, downCount string, status string, lastDate string) *MsgCreateReview {
	return &MsgCreateReview{
		Creator:      creator,
		RecType:      recType,
		ItemId:       itemId,
		ReviewerId:   reviewerId,
		ReviewDetail: reviewDetail,
		ReviewDate:   reviewDate,
		UpCount:      upCount,
		DownCount:    downCount,
		Status:       status,
		LastDate:     lastDate,
	}
}

func (msg *MsgCreateReview) Route() string {
	return RouterKey
}

func (msg *MsgCreateReview) Type() string {
	return "CreateReview"
}

func (msg *MsgCreateReview) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateReview) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateReview) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateReview{}

func NewMsgUpdateReview(creator string, id uint64, recType string, itemId string, reviewerId string, 
	reviewDetail string, reviewDate string, upCount string, downCount string, status string, 
	lastDate string) *MsgUpdateReview {
	return &MsgUpdateReview{
		Id:           id,
		Creator:      creator,
		RecType:      recType,
		ItemId:       itemId,
		ReviewerId:   reviewerId,
		ReviewDetail: reviewDetail,
		ReviewDate:   reviewDate,
		UpCount:      upCount,
		DownCount:    downCount,
		Status:       status,
		LastDate:     lastDate,
	}
}

func (msg *MsgUpdateReview) Route() string {
	return RouterKey
}

func (msg *MsgUpdateReview) Type() string {
	return "UpdateReview"
}

func (msg *MsgUpdateReview) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateReview) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateReview) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateReview{}

func NewMsgDeleteReview(creator string, id uint64) *MsgDeleteReview {
	return &MsgDeleteReview{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteReview) Route() string {
	return RouterKey
}

func (msg *MsgDeleteReview) Type() string {
	return "DeleteReview"
}

func (msg *MsgDeleteReview) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteReview) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteReview) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
