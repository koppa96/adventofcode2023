package day6

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var task2Cmd = &cobra.Command{
	Use:   "task2 <inputFile>",
	Short: "Solves task 2 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		race, err := parseSingleRace(args[0])
		if err != nil {
			return err
		}

		fmt.Println(race.WaysToWin())

		return nil
	},
}

func parseSingleRace(fileName string) (Race, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Race{}, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan() // Time:

	timeStr := ""
	for scanner.Scan() && scanner.Text() != "Distance:" {
		timeStr += scanner.Text()
	}

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return Race{}, err
	}

	distanceStr := ""
	for scanner.Scan() {
		distanceStr += scanner.Text()
	}

	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return Race{}, err
	}

	return Race{Time: time, Distance: distance}, nil
}
