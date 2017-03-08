package cmd

import (
	"github.com/MilosLin/go_bananas/task"
	"github.com/spf13/cobra"
)

var t_name string
var t_argu string

func init() {
	taskCmd.Flags().StringVarP(&t_name, "name", "n", "", "task's name")
	taskCmd.Flags().StringVarP(&t_argu, "argu", "a", "", "task's argument")
	RootCmd.AddCommand(taskCmd)
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Fire some task",
	Long:  `Fire some task`,
	Run: func(cmd *cobra.Command, args []string) {
		task.Dispatch(t_name, t_argu)
		//task.IncognitoDispatch(t_name, args)
	},
}
