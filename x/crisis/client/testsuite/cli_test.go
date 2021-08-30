// +build norace

package testutil

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/v42/testutil/network"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewIntegrationTestSuite(cfg))
}
