package main

import (
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"github.com/samber/lo"
)

func main() {
	runes, err := lib.ParseEngineSchematics("../input.txt")
	if err != nil {
		panic(err)
	}

	sum, err := sumPartNumbers(runes)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func sumPartNumbers(runes [][]rune) (int, error) {
	sum := 0
	for row, rowRunes := range runes {
		for col, r := range rowRunes {
			if isSymbol(r) {
				neighborSum, err := sumNeighboringPartNumbers(runes, row, col)
				if err != nil {
					return 0, err
				}

				sum += neighborSum
			}
		}
	}

	return sum, nil
}

func sumNeighboringPartNumbers(runes [][]rune, row, col int) (int, error) {
	nums, err := lib.NeighboringPartNumbers(runes, row, col)
	if err != nil {
		return 0, err
	}

	return lo.Sum(nums), nil
}

func isSymbol(r rune) bool {
	return r != '.' && (r < 48 || r > 57)
}
