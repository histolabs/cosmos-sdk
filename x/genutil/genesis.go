package genutil

import (
	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/core/genesis"

	"github.com/cosmos/cosmos-sdk/v2/client"
	sdk "github.com/cosmos/cosmos-sdk/v2/types"
	"github.com/cosmos/cosmos-sdk/v2/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, stakingKeeper types.StakingKeeper,
	deliverTx genesis.TxHandler, genesisState types.GenesisState,
	txEncodingConfig client.TxEncodingConfig,
) (validators []abci.ValidatorUpdate, err error) {
	if len(genesisState.GenTxs) > 0 {
		validators, err = DeliverGenTxs(ctx, genesisState.GenTxs, stakingKeeper, deliverTx, txEncodingConfig)
	}
	return
}
