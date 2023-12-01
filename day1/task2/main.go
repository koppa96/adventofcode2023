package main

import (
	"errors"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"github.com/samber/lo"
	"strings"
)

func main() {
	sum, err := lib.ProcessFile("input.txt", firstDigit, lastDigit)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

var numbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func firstDigit(str string) (int, error) {
	runes := []rune(str)
	patterns := lo.Keys(numbers)

	for i := 0; i < len(runes); i++ {
		if digit, ok := lib.AsDigit(runes[i]); ok {
			return digit, nil
		}

		for _, pattern := range patterns {
			if strings.HasPrefix(string(runes[i:]), pattern) {
				return numbers[pattern], nil
			}
		}
	}

	return 0, errors.New("no digits found")
}

func lastDigit(str string) (int, error) {
	runes := []rune(str)
	patterns := lo.Keys(numbers)

	for i := len(runes) - 1; i >= 0; i-- {
		if digit, ok := lib.AsDigit(runes[i]); ok {
			return digit, nil
		}

		for _, pattern := range patterns {
			if strings.HasSuffix(string(runes[:i+1]), pattern) {
				return numbers[pattern], nil
			}
		}
	}

	return 0, errors.New("no digits found")
}
