package day1

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var task1Cmd = &cobra.Command{
	Use:   "task1 <inputFile>",
	Short: "Solves task 1 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sum, err := processFile(args[0], firstDigit, lastDigit)
		if err != nil {
			return err
		}

		fmt.Println(sum)
		return nil
	},
}

func firstDigit(str string) (int, error) {
	for _, r := range str {
		if digit, ok := asDigit(r); ok {
			return digit, nil
		}
	}

	return 0, errors.New("no digits found")
}

func lastDigit(str string) (int, error) {
	runes := []rune(str)

	for i := len(runes) - 1; i >= 0; i-- {
		if digit, ok := asDigit(runes[i]); ok {
			return digit, nil
		}
	}

	return 0, errors.New("no digits found")
}
