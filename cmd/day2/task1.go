package day2

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var task1Cmd = &cobra.Command{
	Use:   "task1 <inputFile>",
	Short: "Solves task 1 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(file)
		sum := 0
		for scanner.Scan() {
			line := scanner.Text()
			game, err := parseGame(line)
			if err != nil {
				panic(err)
			}

			if isGamePossible(game) {
				sum += game.Id
			}
		}

		fmt.Println(sum)

		return nil
	},
}

func isGamePossible(game Game) bool {
	for _, draw := range game.Draws {
		for color, amount := range draw {
			if cubes[color] < amount {
				return false
			}
		}
	}

	return true
}
