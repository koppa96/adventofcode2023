package lib

import (
	"bufio"
	"github.com/samber/lo"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers []int
	Numbers        []int
}

func ParseCard(str string) (Card, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	scanner.Scan() // Game
	scanner.Scan() // Id with colon
	id, err := strconv.Atoi(strings.TrimRight(scanner.Text(), ":"))
	if err != nil {
		return Card{}, err
	}

	winningNumbers, err := ParseNumbers(scanner)
	if err != nil {
		return Card{}, err
	}

	numbers, err := ParseNumbers(scanner)
	if err != nil {
		return Card{}, err
	}

	return Card{
		Id:             id,
		WinningNumbers: winningNumbers,
		Numbers:        numbers,
	}, nil
}

func ParseNumbers(scanner *bufio.Scanner) ([]int, error) {
	var numbers []int
	for scanner.Scan() && scanner.Text() != "|" {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}

func (card *Card) Value() int {
	count := card.Matches()
	if count == 0 {
		return 0
	}

	return int(math.Pow(2, float64(count-1)))
}

func (card *Card) Matches() int {
	count := 0
	for _, number := range card.Numbers {
		if lo.Contains(card.WinningNumbers, number) {
			count++
		}
	}

	return count
}
