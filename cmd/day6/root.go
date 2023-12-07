package day6

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day6",
	Short: "Solves the tasks of the sixth day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
