package day2

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

			sum += getPower(game)
		}

		fmt.Println(sum)

		return nil
	},
}

func getPower(game Game) int {
	minByColor := make(map[string]int)

	for _, draw := range game.Draws {
		for color, amount := range draw {
			if minByColor[color] < amount {
				minByColor[color] = amount
			}
		}
	}

	power := 1
	for _, min := range minByColor {
		power *= min
	}

	return power
}
