package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/faucet2 module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrWithdrawTooOften = sdkerrors.Register(ModuleName, 100, "Each address can withdraw only once")
	ErrFaucetKeyEmpty   = sdkerrors.Register(ModuleName, 101, "Armor should Not be empty.")
	ErrFaucetKeyExisted = sdkerrors.Register(ModuleName, 102, "Faucet key existed")
	// this line is used by starport scaffolding # ibc/errors
)
