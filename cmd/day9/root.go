package day9

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "day9",
	Short: "Solves the tasks of the ninth day",
}

func init() {
	RootCmd.AddCommand(task1Cmd)
	RootCmd.AddCommand(task2Cmd)
}
