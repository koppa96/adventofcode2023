package day3

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day3",
	Short: "Solves the tasks of the third day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
