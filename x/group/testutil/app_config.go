package testutil

import (
	"github.com/cosmos/cosmos-sdk/v2/testutil/configurator"
	_ "github.com/cosmos/cosmos-sdk/v2/x/auth"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/authz"          // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/bank"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/consensus"      // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/genutil"        // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/group/module"   // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/mint"           // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/params"         // import as blank for app wiring
	_ "github.com/cosmos/cosmos-sdk/v2/x/staking"        // import as blank for app wiring
)

var AppConfig = configurator.NewAppConfig(
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.StakingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.ParamsModule(),
	configurator.GenutilModule(),
	configurator.GroupModule(),
)
