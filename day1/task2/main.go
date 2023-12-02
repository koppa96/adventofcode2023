package main

import (
	"errors"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"strings"
)

func main() {
	sum, err := lib.ProcessFile("input.txt", firstDigit, lastDigit)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum)
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

func firstDigit(str string) (int, error) {
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if digit, ok := lib.AsDigit(runes[i]); ok {
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

func lastDigit(str string) (int, error) {
	runes := []rune(str)

	for i := len(runes) - 1; i >= 0; i-- {
		if digit, ok := lib.AsDigit(runes[i]); ok {
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
