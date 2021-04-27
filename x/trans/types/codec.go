package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreateTransaction{}, "trans/CreateTransaction", nil)
cdc.RegisterConcrete(&MsgUpdateTransaction{}, "trans/UpdateTransaction", nil)
cdc.RegisterConcrete(&MsgDeleteTransaction{}, "trans/DeleteTransaction", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateTransaction{},
	&MsgUpdateTransaction{},
	&MsgDeleteTransaction{},
)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
