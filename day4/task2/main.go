package main

import (
	"bufio"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"github.com/samber/lo"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	instances := make(map[int]int)
	for scanner.Scan() {
		card, err := lib.ParseCard(scanner.Text())
		if err != nil {
			panic(err)
		}

		instances[card.Id]++
		matches := card.Matches()
		for i := 0; i < matches; i++ {
			instances[card.Id+i+1] += instances[card.Id]
		}
	}

	fmt.Println(lo.Sum(lo.Values(instances)))
}
