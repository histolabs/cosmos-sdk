package module

import (
	sdk "github.com/cosmos/cosmos-sdk/v2/types"
	"github.com/cosmos/cosmos-sdk/v2/x/group/keeper"
)

// EndBlocker called at every block, updates proposal's `FinalTallyResult` and
// prunes expired proposals.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) error {
	if err := k.TallyProposalsAtVPEnd(ctx); err != nil {
		return err
	}

	return k.PruneProposals(ctx)
}
