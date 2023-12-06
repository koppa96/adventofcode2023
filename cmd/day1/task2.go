package day1

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var task2Cmd = &cobra.Command{
	Use:   "task2 <inputFile>",
	Short: "Solves task 2 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sum, err := processFile(args[0], firstDigitOrNumberString, lastDigitOrNumberString)
		if err != nil {
			panic(err)
		}

		fmt.Println(sum)
		return nil
	},
}

var numbers = [10]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func firstDigitOrNumberString(str string) (int, error) {
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if digit, ok := asDigit(runes[i]); ok {
			return digit, nil
		}

		for value, pattern := range numbers {
			if strings.HasPrefix(string(runes[i:]), pattern) {
				return value, nil
			}
		}
	}

	return 0, errors.New("no digits found")
}

func lastDigitOrNumberString(str string) (int, error) {
	runes := []rune(str)

	for i := len(runes) - 1; i >= 0; i-- {
		if digit, ok := asDigit(runes[i]); ok {
			return digit, nil
		}

		for value, pattern := range numbers {
			if strings.HasSuffix(string(runes[:i+1]), pattern) {
				return value, nil
			}
		}
	}

	return 0, errors.New("no digits found")
}
