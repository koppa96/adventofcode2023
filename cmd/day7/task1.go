package day7

import (
	"fmt"
	"github.com/spf13/cobra"
)

var task1Cmd = &cobra.Command{
	Use:   "task1 <inputFile>",
	Short: "Solves task 1 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		winnings, err := ComputeWinnings(args[0], cardValuesWithoutJoker, computeHandTypeWithoutJoker)
		if err != nil {
			return err
		}

		fmt.Println(winnings)

		return nil
	},
}

var cardValuesWithoutJoker = map[string]uint8{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func computeHandTypeWithoutJoker(hand []uint8) HandType {
	var countsByCardType [13]uint8
	var maxCount, secondMaxCount, maxType uint8
	for _, card := range hand {
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
