package day1

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day1",
	Short: "Solves the tasks of the first day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
