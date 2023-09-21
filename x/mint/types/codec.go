package types

import (
	"github.com/cosmos/cosmos-sdk/v2/codec"
	"github.com/cosmos/cosmos-sdk/v2/codec/legacy"
	"github.com/cosmos/cosmos-sdk/v2/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/v2/types"
	"github.com/cosmos/cosmos-sdk/v2/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(Params{}, "cosmos-sdk/x/mint/Params", nil)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "cosmos-sdk/x/mint/MsgUpdateParams")
}

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
