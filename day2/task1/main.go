package main

import (
	"bufio"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"os"
)

var cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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

		if isGamePossible(game) {
			sum += game.Id
		}
	}

	fmt.Println(sum)
}

func isGamePossible(game lib.Game) bool {
	for _, draw := range game.Draws {
		for color, amount := range draw {
			if cubes[color] < amount {
				return false
			}
		}
	}

	return true
}
