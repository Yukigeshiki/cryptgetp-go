package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cryptgetp",
	Short: "A just for fun CLI application to fetch cryptocurrency prices",
	Long: `To fetch the price of your preferred cryptocurrency, returned in your preferred fiat currency, just run with relevant flags. For example:
$ ./cryptgetp fetch --crypto BTC --in USD`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
