package distribution_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"cosmossdk.io/depinject"
	"cosmossdk.io/log"

	simtestutil "github.com/cosmos/cosmos-sdk/v2/testutil/sims"
	authkeeper "github.com/cosmos/cosmos-sdk/v2/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/v2/x/auth/types"
	"github.com/cosmos/cosmos-sdk/v2/x/distribution/testutil"
	"github.com/cosmos/cosmos-sdk/v2/x/distribution/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	var accountKeeper authkeeper.AccountKeeper

	app, err := simtestutil.SetupAtGenesis(
		depinject.Configs(
			testutil.AppConfig,
			depinject.Supply(log.NewNopLogger()),
		),
		&accountKeeper)
	assert.NilError(t, err)

	ctx := app.BaseApp.NewContext(false)
	acc := accountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	assert.Assert(t, acc != nil)
}
