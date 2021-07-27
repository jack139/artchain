package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.NewLegacyAmino()
//var (
//	amino     = codec.NewLegacyAmino()
//	ModuleCdc = codec.NewAminoCodec(amino)
//)

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgMint{}, "faucet/Mint", nil)
	cdc.RegisterConcrete(MsgFaucetKey{}, "faucet/FaucetKey", nil)
}
