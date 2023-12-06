package day5

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var task1Cmd = &cobra.Command{
	Use:   "task1 <inputFile>",
	Short: "Solves task 1 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open("../input.txt")
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		seeds, err := parseSeeds(scanner)
		if err != nil {
			panic(err)
		}

		mapFn := parseAndCombineMaps(scanner)

		min := mapFn(seeds[0])
		for i := 1; i < len(seeds); i++ {
			location := mapFn(seeds[i])
			if location < min {
				min = location
			}
		}

		fmt.Println(min)

		return nil
	},
}
