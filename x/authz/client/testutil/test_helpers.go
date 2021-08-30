package testutil

import (
	"github.com/cosmos/cosmos-sdk/v43/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/v43/testutil/cli"
	"github.com/cosmos/cosmos-sdk/v43/testutil/network"
	"github.com/cosmos/cosmos-sdk/v43/x/authz/client/cli"
)

func ExecGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
