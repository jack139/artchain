package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateBid{}, "auction/CreateBid", nil)
	cdc.RegisterConcrete(&MsgUpdateBid{}, "auction/UpdateBid", nil)
	cdc.RegisterConcrete(&MsgDeleteBid{}, "auction/DeleteBid", nil)

	cdc.RegisterConcrete(&MsgCreateRequest{}, "auction/CreateRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateRequest{}, "auction/UpdateRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteRequest{}, "auction/DeleteRequest", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBid{},
		&MsgUpdateBid{},
		&MsgDeleteBid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRequest{},
		&MsgUpdateRequest{},
		&MsgDeleteRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
