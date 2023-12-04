package main

import (
	"bufio"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		card, err := lib.ParseCard(scanner.Text())
		if err != nil {
			panic(err)
		}

		sum += card.Value()
	}

	fmt.Println(sum)
}
