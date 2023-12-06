package day5

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day3",
	Short: "Solves the tasks of the fifth day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
