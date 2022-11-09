package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	crypto string
	in     string
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch the price of a given cryptocurrency returned in a given fiat currency.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(crypto, in)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVarP(&crypto, "crypto", "c", "BTC", "Cryptocurrency to fetch")
	fetchCmd.Flags().StringVarP(&in, "in", "i", "USD", "Fiat price will be returned in")
}
