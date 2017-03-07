package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
	date    string
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of this binary file",
	Long:  `Print the version number and the create date of this binary file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version=%s, date=%s", version, date)
	},
}
