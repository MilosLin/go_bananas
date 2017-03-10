package cmd

import (
	"github.com/MilosLin/go_bananas/api"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start api server",
	Long:  `Start api server`,
	Run: func(cmd *cobra.Command, args []string) {
		api.NewAPIService().Start()
	},
}
