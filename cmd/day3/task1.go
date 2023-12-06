package day3

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var task1Cmd = &cobra.Command{
	Use:   "task1 <inputFile>",
	Short: "Solves task 1 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		runes, err := parseEngineSchematics(args[0])
		if err != nil {
			panic(err)
		}

		sum, err := sumPartNumbers(runes)
		if err != nil {
			panic(err)
		}

		fmt.Println(sum)

		return nil
	},
}

func sumPartNumbers(runes [][]rune) (int, error) {
	sum := 0
	for row, rowRunes := range runes {
		for col, r := range rowRunes {
			if isSymbol(r) {
				neighborSum, err := sumNeighboringPartNumbers(runes, row, col)
				if err != nil {
					return 0, err
				}

				sum += neighborSum
			}
		}
	}

	return sum, nil
}

func sumNeighboringPartNumbers(runes [][]rune, row, col int) (int, error) {
	nums, err := neighboringPartNumbers(runes, row, col)
	if err != nil {
		return 0, err
	}

	return lo.Sum(nums), nil
}

func isSymbol(r rune) bool {
	return r != '.' && (r < 48 || r > 57)
}
