package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUser{}

func NewMsgCreateUser(creator string, recType string, name string, userType string, address string, phone string, email string, bank string, accountNo string, status string, regDate string) *MsgCreateUser {
	return &MsgCreateUser{
		Creator:   creator,
		RecType:   recType,
		Name:      name,
		UserType:  userType,
		Address:   address,
		Phone:     phone,
		Email:     email,
		Bank:      bank,
		AccountNo: accountNo,
		Status:    status,
		RegDate:   regDate,
	}
}

func (msg *MsgCreateUser) Route() string {
	return RouterKey
}

func (msg *MsgCreateUser) Type() string {
	return "CreateUser"
}

func (msg *MsgCreateUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUser{}

func NewMsgUpdateUser(creator string, id uint64, recType string, name string, userType string, address string, phone string, email string, bank string, accountNo string, status string, regDate string) *MsgUpdateUser {
	return &MsgUpdateUser{
		Id:        id,
		Creator:   creator,
		RecType:   recType,
		Name:      name,
		UserType:  userType,
		Address:   address,
		Phone:     phone,
		Email:     email,
		Bank:      bank,
		AccountNo: accountNo,
		Status:    status,
		RegDate:   regDate,
	}
}

func (msg *MsgUpdateUser) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUser) Type() string {
	return "UpdateUser"
}

func (msg *MsgUpdateUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateUser{}

func NewMsgDeleteUser(creator string, id uint64) *MsgDeleteUser {
	return &MsgDeleteUser{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteUser) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUser) Type() string {
	return "DeleteUser"
}

func (msg *MsgDeleteUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
