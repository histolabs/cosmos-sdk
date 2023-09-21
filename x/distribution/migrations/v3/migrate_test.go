package v3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/v2/runtime"
	"github.com/cosmos/cosmos-sdk/v2/testutil"
	sdk "github.com/cosmos/cosmos-sdk/v2/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/v2/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/v2/x/distribution"
	"github.com/cosmos/cosmos-sdk/v2/x/distribution/exported"
	v3 "github.com/cosmos/cosmos-sdk/v2/x/distribution/migrations/v3"
	"github.com/cosmos/cosmos-sdk/v2/x/distribution/types"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps exported.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	cdc := moduletestutil.MakeTestEncodingConfig(distribution.AppModuleBasic{}).Codec
	storeKey := storetypes.NewKVStoreKey(v3.ModuleName)
	storeService := runtime.NewKVStoreService(storeKey)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	store := ctx.KVStore(storeKey)

	legacySubspace := newMockSubspace(types.DefaultParams())
	require.NoError(t, v3.MigrateStore(ctx, storeService, legacySubspace, cdc))

	var res types.Params
	bz := store.Get(v3.ParamsKey)
	require.NoError(t, cdc.Unmarshal(bz, &res))
	require.Equal(t, legacySubspace.ps, res)
}
