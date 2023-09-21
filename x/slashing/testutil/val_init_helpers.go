package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/v2/types"
)

// InitTokens is the default power validators are initialized to have within tests
var InitTokens = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
