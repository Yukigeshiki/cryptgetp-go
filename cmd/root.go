package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cryptgetp",
	Short: "A just for fun CLI tool to fetch cryptocurrency prices written in Go",
	Long: `To fetch the price of your preferred cryptocurrency, returned in your preferred fiat currency, just run with the relevant flags. For example:
$ ./cryptgetp fetch --crypto BTC --in USD --key <your-key-here>`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
