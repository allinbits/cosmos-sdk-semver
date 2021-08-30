package keeper_test

import (
	"fmt"
	"context"
	"bytes"

	"github.com/cosmos/cosmos-sdk/v43/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/v43/types"
	"github.com/cosmos/cosmos-sdk/v43/x/auth/types"
)

const addrStr = "cosmos13c3d4wq2t22dl0dstraf8jc3f902e3fsy9n3wv"
var addrBytes = []byte{0x8e, 0x22, 0xda, 0xb8, 0xa, 0x5a, 0x94, 0xdf, 0xbd, 0xb0, 0x58, 0xfa, 0x93, 0xcb, 0x11, 0x49, 0x5e, 0xac, 0xc5, 0x30}

func (suite *KeeperTestSuite) TestGRPCQueryAccounts() {
	var (
		req *types.QueryAccountsRequest
	)
	_, _, first := testdata.KeyTestPubAddr()
	_, _, second := testdata.KeyTestPubAddr()

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func(res *types.QueryAccountsResponse)
	}{
		{
			"success",
			func() {
				suite.app.AccountKeeper.SetAccount(suite.ctx,
					suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, first))
				suite.app.AccountKeeper.SetAccount(suite.ctx,
					suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, second))
				req = &types.QueryAccountsRequest{}
			},
			true,
			func(res *types.QueryAccountsResponse) {
				for _, acc := range res.Accounts {
					var account types.AccountI
					err := suite.app.InterfaceRegistry().UnpackAny(acc, &account)
					suite.Require().NoError(err)

					suite.Require().True(
						first.Equals(account.GetAddress()) || second.Equals(account.GetAddress()))
				}
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			tc.malleate()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.queryClient.Accounts(ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

			tc.posttests(res)
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryAccount() {
	var (
		req *types.QueryAccountRequest
	)
	_, _, addr := testdata.KeyTestPubAddr()

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func(res *types.QueryAccountResponse)
	}{
		{
			"empty request",
			func() {
				req = &types.QueryAccountRequest{}
			},
			false,
			func(res *types.QueryAccountResponse) {},
		},
		{
			"invalid request",
			func() {
				req = &types.QueryAccountRequest{Address: ""}
			},
			false,
			func(res *types.QueryAccountResponse) {},
		},
		{
			"invalid request with empty byte array",
			func() {
				req = &types.QueryAccountRequest{Address: ""}
			},
			false,
			func(res *types.QueryAccountResponse) {},
		},
		{
			"account not found",
			func() {
				req = &types.QueryAccountRequest{Address: addr.String()}
			},
			false,
			func(res *types.QueryAccountResponse) {},
		},
		{
			"success",
			func() {
				suite.app.AccountKeeper.SetAccount(suite.ctx,
					suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr))
				req = &types.QueryAccountRequest{Address: addr.String()}
			},
			true,
			func(res *types.QueryAccountResponse) {
				var newAccount types.AccountI
				err := suite.app.InterfaceRegistry().UnpackAny(res.Account, &newAccount)
				suite.Require().NoError(err)
				suite.Require().NotNil(newAccount)
				suite.Require().True(addr.Equals(newAccount.GetAddress()))
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			tc.malleate()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.queryClient.Account(ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

			tc.posttests(res)
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryParameters() {
	var (
		req       *types.QueryParamsRequest
		expParams types.Params
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				req = &types.QueryParamsRequest{}
				expParams = suite.app.AccountKeeper.GetParams(suite.ctx)
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			tc.malleate()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.queryClient.Params(ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(expParams, res.Params)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestGRPCQueryModuleAccounts() {
	var (
		req *types.QueryModuleAccountsRequest
	)

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func(res *types.QueryModuleAccountsResponse)
	}{
		{
			"success",
			func() {
				req = &types.QueryModuleAccountsRequest{}
			},
			true,
			func(res *types.QueryModuleAccountsResponse) {
				var mintModuleExists = false
				for _, acc := range res.Accounts {
					var account types.AccountI
					err := suite.app.InterfaceRegistry().UnpackAny(acc, &account)
					suite.Require().NoError(err)

					moduleAccount, ok := account.(types.ModuleAccountI)

					suite.Require().True(ok)
					if moduleAccount.GetName() == "mint" {
						mintModuleExists = true
					}
				}
				suite.Require().True(mintModuleExists)
			},
		},
		{
			"invalid module name",
			func() {
				req = &types.QueryModuleAccountsRequest{}
			},
			true,
			func(res *types.QueryModuleAccountsResponse) {
				var mintModuleExists = false
				for _, acc := range res.Accounts {
					var account types.AccountI
					err := suite.app.InterfaceRegistry().UnpackAny(acc, &account)
					suite.Require().NoError(err)

					moduleAccount, ok := account.(types.ModuleAccountI)

					suite.Require().True(ok)
					if moduleAccount.GetName() == "falseCase" {
						mintModuleExists = true
					}
				}
				suite.Require().False(mintModuleExists)
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			tc.malleate()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.queryClient.ModuleAccounts(ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

			tc.posttests(res)
		})
	}
}

func (suite *KeeperTestSuite) TestBech32Prefix() {
	suite.SetupTest() // reset
	req := &types.Bech32PrefixRequest{}
	res, err := suite.queryClient.Bech32Prefix(context.Background(), req)
	suite.Require().NoError(err)
	suite.Require().NotNil(res)
	suite.Require().Equal(sdk.Bech32MainPrefix, res.Bech32Prefix)
}

func (suite *KeeperTestSuite) TestAddressBytesToString() {
	testCases := []struct {
		msg       string
		req       *types.AddressBytesToStringRequest
		expPass   bool
	}{
		{
			"success",
			&types.AddressBytesToStringRequest{AddressBytes: addrBytes},
			true,
		},
		{
			"request is empty",
			&types.AddressBytesToStringRequest{},
			false,
		},
		{
			"empty account address in request",
			&types.AddressBytesToStringRequest{AddressBytes: []byte{}},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			res, err := suite.queryClient.AddressBytesToString(context.Background(), tc.req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(res.AddressString, addrStr)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestAddressStringToBytes() {
	testCases := []struct {
		msg       string
		req       *types.AddressStringToBytesRequest
		expPass   bool
	}{
		{
			"success",
			&types.AddressStringToBytesRequest{AddressString: addrStr},
			true,
		},
		{
			"request is empty",
			&types.AddressStringToBytesRequest{},
			false,
		},
		{
			"AddressString field in request is empty",
			&types.AddressStringToBytesRequest{AddressString: ""},
			false,
		},
		{
			"address prefix is incorrect",
			&types.AddressStringToBytesRequest{AddressString: "regen13c3d4wq2t22dl0dstraf8jc3f902e3fsy9n3wv" },
			false,
		},

	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			res, err := suite.queryClient.AddressStringToBytes(context.Background(), tc.req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().True(bytes.Equal(res.AddressBytes, addrBytes))
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

		})
	}
}