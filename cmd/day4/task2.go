package day4

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
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
		instances := make(map[int]int)
		for scanner.Scan() {
			card, err := parseCard(scanner.Text())
			if err != nil {
				panic(err)
			}

			instances[card.Id]++
			matches := card.Matches()
			for i := 0; i < matches; i++ {
				instances[card.Id+i+1] += instances[card.Id]
			}
		}

		fmt.Println(lo.Sum(lo.Values(instances)))

		return nil
	},
}
