package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ovrclk/akash/x/market/query"
	"github.com/ovrclk/akash/x/market/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(key string, cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Market query commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(client.GetCommands(
		cmdGetOrders(key, cdc),
		cmdGetBids(key, cdc),
		cmdGetLeases(key, cdc),
	)...)

	return cmd
}

func cmdGetOrders(key string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "orders",
		RunE: func(cmd *cobra.Command, args []string) error {
			var obj query.Orders
			ctx := context.NewCLIContext().WithCodec(cdc)
			buf, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s", key, query.OrdersPath()), nil)
			if err != nil {
				return err
			}
			if err := cdc.UnmarshalJSON(buf, &obj); err != nil {
				return err
			}
			return ctx.PrintOutput(obj)
		},
	}
}

func cmdGetBids(key string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "bids",
		RunE: func(cmd *cobra.Command, args []string) error {
			var obj query.Bids
			ctx := context.NewCLIContext().WithCodec(cdc)
			buf, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s", key, query.BidsPath()), nil)
			if err != nil {
				return err
			}
			if err := cdc.UnmarshalJSON(buf, &obj); err != nil {
				return err
			}
			return ctx.PrintOutput(obj)
		},
	}
}

func cmdGetLeases(key string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "leases",
		RunE: func(cmd *cobra.Command, args []string) error {
			var obj query.Leases
			ctx := context.NewCLIContext().WithCodec(cdc)
			buf, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s", key, query.LeasesPath()), nil)
			if err != nil {
				return err
			}
			if err := cdc.UnmarshalJSON(buf, &obj); err != nil {
				return err
			}
			return ctx.PrintOutput(obj)
		},
	}
}