package day7

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day7",
	Short: "Solves the tasks of the seventh day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
