package main

import (
	"errors"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
)

func main() {
	sum, err := lib.ProcessFile("input.txt", firstDigit, lastDigit)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func firstDigit(str string) (int, error) {
	for _, r := range str {
		if digit, ok := lib.AsDigit(r); ok {
			return digit, nil
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
	}

	return 0, errors.New("no digits found")
}
