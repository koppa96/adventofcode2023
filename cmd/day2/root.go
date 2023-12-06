package day2

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day2",
	Short: "Solves the tasks of the second day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
