package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Crypt-Get-P.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crypt-Get-P v0.2")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
