package client

import (
	govclient "github.com/cosmos/cosmos-sdk/v2/x/gov/client"
	"github.com/cosmos/cosmos-sdk/v2/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
