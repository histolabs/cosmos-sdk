package authz

import (
	"github.com/cosmos/cosmos-sdk/v2/client"
	addresscodec "github.com/cosmos/cosmos-sdk/v2/codec/address"
	"github.com/cosmos/cosmos-sdk/v2/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/v2/testutil/cli"
	"github.com/cosmos/cosmos-sdk/v2/x/authz/client/cli"
)

func CreateGrant(clientCtx client.Context, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization(addresscodec.NewBech32Codec("cosmos"))
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
