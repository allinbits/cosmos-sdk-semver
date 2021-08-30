package keys

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/v42/client"
	"github.com/cosmos/cosmos-sdk/v42/client/flags"
	"github.com/cosmos/cosmos-sdk/v42/crypto/hd"
	"github.com/cosmos/cosmos-sdk/v42/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/v42/testutil"
	sdk "github.com/cosmos/cosmos-sdk/v42/types"
)

func Test_runListCmd(t *testing.T) {
	cmd := ListKeysCmd()
	cmd.Flags().AddFlagSet(Commands("home").PersistentFlags())

	kbHome1 := t.TempDir()
	kbHome2 := t.TempDir()

	mockIn := testutil.ApplyMockIODiscardOutErr(cmd)
	kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, kbHome2, mockIn)
	require.NoError(t, err)

	clientCtx := client.Context{}.WithKeyring(kb)
	ctx := context.WithValue(context.Background(), client.ClientContextKey, &clientCtx)

	path := "" //sdk.GetConfig().GetFullBIP44Path()
	_, err = kb.NewAccount("something", testutil.TestMnemonic, "", path, hd.Secp256k1)
	require.NoError(t, err)

	t.Cleanup(func() {
		kb.Delete("something") // nolint:errcheck
	})

	type args struct {
		cmd  *cobra.Command
		args []string
	}

	testData := []struct {
		name    string
		kbDir   string
		wantErr bool
	}{
		{"keybase: empty", kbHome1, false},
		{"keybase: w/key", kbHome2, false},
	}
	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cmd.SetArgs([]string{
				fmt.Sprintf("--%s=%s", flags.FlagHome, tt.kbDir),
				fmt.Sprintf("--%s=false", flagListNames),
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
			})

			if err := cmd.ExecuteContext(ctx); (err != nil) != tt.wantErr {
				t.Errorf("runListCmd() error = %v, wantErr %v", err, tt.wantErr)
			}

			cmd.SetArgs([]string{
				fmt.Sprintf("--%s=%s", flags.FlagHome, tt.kbDir),
				fmt.Sprintf("--%s=true", flagListNames),
				fmt.Sprintf("--%s=%s", flags.FlagKeyringBackend, keyring.BackendTest),
			})

			if err := cmd.ExecuteContext(ctx); (err != nil) != tt.wantErr {
				t.Errorf("runListCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
