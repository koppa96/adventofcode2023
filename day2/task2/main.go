package main

import (
	"bufio"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		game, err := lib.ParseGame(line)
		if err != nil {
			panic(err)
		}

		sum += getPower(game)
	}

	fmt.Println(sum)
}

func getPower(game lib.Game) int {
	minByColor := make(map[string]int)

	for _, draw := range game.Draws {
		for color, amount := range draw {
			if minByColor[color] < amount {
				minByColor[color] = amount
			}
		}
	}

	power := 1
	for _, min := range minByColor {
		power *= min
	}

	return power
}
