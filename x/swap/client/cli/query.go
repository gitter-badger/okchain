package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/okchain/x/swap/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group swap queries under a subcommand
	swapQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	swapQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdSwapTokenPair(queryRoute, cdc),
		)...,
	)

	return swapQueryCmd
}

func GetCmdSwapTokenPair(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "tokenpair [token]",
		Short: "TokenPair token name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			tokenName := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QuerySwapTokenPair, tokenName), nil)
			if err != nil {
				fmt.Printf("token piar - %s doesn't exist. error:%s \n", tokenName, err.Error())
				return nil
			}

			fmt.Println(string(res))
			return nil
		},
	}
}
