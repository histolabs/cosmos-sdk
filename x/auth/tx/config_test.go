package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/v2/codec"
	"github.com/cosmos/cosmos-sdk/v2/codec/testutil"
	"github.com/cosmos/cosmos-sdk/v2/std"
	"github.com/cosmos/cosmos-sdk/v2/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/v2/types"
	txtestutil "github.com/cosmos/cosmos-sdk/v2/x/auth/tx/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := testutil.CodecOptions{}.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, txtestutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
