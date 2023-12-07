package day6

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var task1Cmd = &cobra.Command{
	Use:   "task1 <inputFile>",
	Short: "Solves task 1 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		races, err := parseRaces(args[0])
		if err != nil {
			return err
		}

		product := 1
		for _, race := range races {
			product *= race.WaysToWin()
		}

		fmt.Println(product)

		return nil
	},
}

func parseRaces(fileName string) ([]Race, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan() // Time:

	var races []Race
	for scanner.Scan() && scanner.Text() != "Distance:" {
		time, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		races = append(races, Race{Time: time})
	}

	var i int
	for scanner.Scan() {
		distance, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		races[i].Distance = distance
		i++
	}

	return races, nil
}
