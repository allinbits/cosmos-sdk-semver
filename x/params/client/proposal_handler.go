package client

import (
	govclient "github.com/cosmos/cosmos-sdk/v42/x/gov/client"
	"github.com/cosmos/cosmos-sdk/v42/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
