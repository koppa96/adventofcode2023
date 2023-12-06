package day4

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
		file, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(file)
		sum := 0
		for scanner.Scan() {
			card, err := parseCard(scanner.Text())
			if err != nil {
				panic(err)
			}

			sum += card.Value()
		}

		fmt.Println(sum)

		return nil
	},
}
