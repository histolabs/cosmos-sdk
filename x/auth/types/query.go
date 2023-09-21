package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/v2/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/v2/types"
)

func (m *QueryAccountResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var account sdk.AccountI
	return unpacker.UnpackAny(m.Account, &account)
}

var _ codectypes.UnpackInterfacesMessage = &QueryAccountResponse{}
