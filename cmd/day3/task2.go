package day3

import (
	"fmt"
	"github.com/spf13/cobra"
)

var task2Cmd = &cobra.Command{
	Use:   "task2 <inputFile>",
	Short: "Solves task 2 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		runes, err := parseEngineSchematics(args[0])
		if err != nil {
			panic(err)
		}

		sum, err := sumGearRatios(runes)
		if err != nil {
			panic(err)
		}

		fmt.Println(sum)

		return nil
	},
}

func sumGearRatios(runes [][]rune) (int, error) {
	sum := 0
	for row, rowRunes := range runes {
		for col, r := range rowRunes {
			if r == '*' {
				partNums, err := neighboringPartNumbers(runes, row, col)
				if err != nil {
					return 0, err
				}

				if len(partNums) == 2 {
					sum += partNums[0] * partNums[1]
				}
			}
		}
	}

	return sum, nil
}
