package main

import (
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
)

func main() {
	runes, err := lib.ParseEngineSchematics("../input.txt")
	if err != nil {
		panic(err)
	}

	sum, err := sumGearRatios(runes)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func sumGearRatios(runes [][]rune) (int, error) {
	sum := 0
	for row, rowRunes := range runes {
		for col, r := range rowRunes {
			if r == '*' {
				partNums, err := lib.NeighboringPartNumbers(runes, row, col)
				if err != nil {
					return 0, err
				}

				if len(partNums) == 2 {
					sum += partNums[0] * partNums[1]
				}
			}
		}
	}

	return sum, nil
}
