package day4

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day4",
	Short: "Solves the tasks of the fourth day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
