package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/v43/simapp"
	sdk "github.com/cosmos/cosmos-sdk/v43/types"
	authtypes "github.com/cosmos/cosmos-sdk/v43/x/auth/types"
	"github.com/cosmos/cosmos-sdk/v43/x/distribution/types"
)

var (
	PKS = simapp.CreateTestPubKeys(5)

	valConsPk1 = PKS[0]
	valConsPk2 = PKS[1]
	valConsPk3 = PKS[2]

	valConsAddr1 = sdk.ConsAddress(valConsPk1.Address())
	valConsAddr2 = sdk.ConsAddress(valConsPk2.Address())

	distrAcc = authtypes.NewEmptyModuleAccount(types.ModuleName)
)
