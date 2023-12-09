package day7

import (
	"fmt"
	"github.com/spf13/cobra"
)

var task2Cmd = &cobra.Command{
	Use:   "task2 <inputFile>",
	Short: "Solves task 2 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		winnings, err := ComputeWinnings(args[0], cardValuesWithJoker, computeHandTypeWithJoker)
		if err != nil {
			return err
		}

		fmt.Println(winnings)

		return nil
	},
}

var cardValuesWithJoker = map[string]uint8{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func computeHandTypeWithJoker(hand []uint8) HandType {
	var countsByCardType [13]uint8
	var maxCount, secondMaxCount, jokersCount, maxType uint8
	for _, card := range hand {
		if card == 0 {
			jokersCount++
			continue
		}

		countsByCardType[card]++
		if countsByCardType[card] > maxCount {
			if maxType != card {
				secondMaxCount = maxCount
			}

			maxCount = countsByCardType[card]
			maxType = card
		} else if countsByCardType[card] > secondMaxCount {
			secondMaxCount = countsByCardType[card]
		}
	}

	maxCount += jokersCount
	switch maxCount {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if secondMaxCount == 2 {
			return FullHouse
		}

		return ThreeOfAKind
	case 2:
		if secondMaxCount == 2 {
			return TwoPair
		}

		return OnePair
	case 1:
		return HighCard
	}

	panic("famous last words: reaching this line is not possible")
}
