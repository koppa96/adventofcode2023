package day8

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day8",
	Short: "Solves the tasks of the eighth day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
