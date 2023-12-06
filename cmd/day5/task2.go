package day5

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"math"
	"os"
	"time"
)

var task2Cmd = &cobra.Command{
	Use:   "task2 <inputFile>",
	Short: "Solves task 2 for the provided input file.",
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

		start := time.Now()
		min := math.MaxInt
		for i := 0; i < len(seeds); i += 2 {
			fmt.Printf("Processing seed range %d/%d\n", i/2, len(seeds)/2)

			start := seeds[i]
			count := seeds[i+1]

			for j := 0; j < count; j++ {
				location := mapFn(start + j)
				if location < min {
					min = location
				}
			}
		}

		fmt.Println(min)
		fmt.Printf("Completed in %s", time.Since(start))

		return nil
	},
}
