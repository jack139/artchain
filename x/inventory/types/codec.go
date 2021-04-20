package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateReview{}, "inventory/CreateReview", nil)
	cdc.RegisterConcrete(&MsgUpdateReview{}, "inventory/UpdateReview", nil)
	cdc.RegisterConcrete(&MsgDeleteReview{}, "inventory/DeleteReview", nil)

	cdc.RegisterConcrete(&MsgCreateItem{}, "inventory/CreateItem", nil)
	cdc.RegisterConcrete(&MsgUpdateItem{}, "inventory/UpdateItem", nil)
	cdc.RegisterConcrete(&MsgDeleteItem{}, "inventory/DeleteItem", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateReview{},
		&MsgUpdateReview{},
		&MsgDeleteReview{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateItem{},
		&MsgUpdateItem{},
		&MsgDeleteItem{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
