package day10

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day10",
	Short: "Solves the tasks of the tenth day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
