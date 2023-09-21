package testutil

import (
	_ "cosmossdk.io/x/evidence" // import as blank for app wiring

	"github.com/cosmos/cosmos-sdk/v2/testutil/configurator"
	_ "github.com/cosmos/cosmos-sdk/v2/x/auth"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/bank"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/consensus"      // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/genutil"        // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/params"         // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/slashing"       // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/staking"        // import as blank for app wiring
)

var AppConfig = configurator.NewAppConfig(
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.StakingModule(),
	configurator.SlashingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.ParamsModule(),
	configurator.EvidenceModule(),
	configurator.GenutilModule(),
)
