package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const version = "0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of this tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
