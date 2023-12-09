package cmd

import (
	"github.com/koppa96/adventofcode2023/cmd/day1"
	"github.com/koppa96/adventofcode2023/cmd/day2"
	"github.com/koppa96/adventofcode2023/cmd/day3"
	"github.com/koppa96/adventofcode2023/cmd/day4"
	"github.com/koppa96/adventofcode2023/cmd/day5"
	"github.com/koppa96/adventofcode2023/cmd/day6"
	"github.com/koppa96/adventofcode2023/cmd/day7"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "adventofcode",
	Short: "AdventOfCode is a CLI tool that can be used to calculate the solution to the tasks of 2023's Advent of Code.",
}

func init() {
	RootCmd.AddCommand(day1.RootCmd)
	RootCmd.AddCommand(day2.RootCmd)
	RootCmd.AddCommand(day3.RootCmd)
	RootCmd.AddCommand(day4.RootCmd)
	RootCmd.AddCommand(day5.RootCmd)
	RootCmd.AddCommand(day6.RootCmd)
	RootCmd.AddCommand(day7.RootCmd)
}
