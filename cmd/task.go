package cmd

import (
	"github.com/MilosLin/go_bananas/task"
	"github.com/spf13/cobra"
)

var taskName string
var taskArgu string

func init() {
	taskCmd.Flags().StringVarP(&taskName, "name", "n", "", "task's name")
	taskCmd.Flags().StringVarP(&taskArgu, "argu", "a", "", "task's argument")
	RootCmd.AddCommand(taskCmd)
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Fire some task",
	Long:  `Fire some task`,
	Run: func(cmd *cobra.Command, args []string) {
		task.Dispatch(&taskName, &taskArgu)
		//task.corbaDispatch(taskName, args)
	},
}
