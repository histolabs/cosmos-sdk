package nft

import (
	types "github.com/cosmos/cosmos-sdk/v2/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/v2/types"
	"github.com/cosmos/cosmos-sdk/v2/types/msgservice"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
