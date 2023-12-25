package day11

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var task2Cmd = &cobra.Command{
	Use:   "task2 <inputFile>",
	Short: "Solves task 2 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(args[0])
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)

		var runes [][]rune
		for scanner.Scan() {
			runes = append(runes, []rune(scanner.Text()))
		}

		universe := ParseUniverse(runes)

		var sum int64
		for i := 0; i < len(universe.Galaxies)-1; i++ {
			for j := i + 1; j < len(universe.Galaxies); j++ {
				sum += int64(universe.DistanceBetween(universe.Galaxies[i], universe.Galaxies[j], 1_000_000))
			}
		}

		fmt.Println(sum)

		return nil
	},
}
