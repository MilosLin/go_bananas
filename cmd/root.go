package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd : Root Command
var RootCmd = &cobra.Command{
	Use:   "",
	Short: "golang practice project",
	Long:  `A simple golang framework to practice golang.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Use \"help\" command to see more detail")
	},
}

// Execute : Execute Root Command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
}
