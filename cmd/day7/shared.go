package day7

import (
	"bufio"
	"cmp"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Cards []uint8
	Value uint32
	Bid   int
}

type HandType uint8

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func ComputeWinnings(fileName string, cardValues map[string]uint8, typeComputer func([]uint8) HandType) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	var hands []Hand
	for scanner.Scan() {
		hand, err := ParseHand(scanner.Text(), cardValues, typeComputer)
		if err != nil {
			return 0, err
		}

		hands = append(hands, hand)
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return cmp.Compare(a.Value, b.Value)
	})

	var winnings int
	for i, hand := range hands {
		winnings += (i + 1) * hand.Bid
	}

	return winnings, nil
}

func ParseHand(hand string, cardValues map[string]uint8, typeComputer func([]uint8) HandType) (Hand, error) {
	scanner := bufio.NewScanner(strings.NewReader(hand))
	scanner.Split(bufio.ScanRunes)

	var cards []uint8
	for i := 0; i < 5; i++ {
		if !scanner.Scan() {
			return Hand{}, errors.New("failed to parse hand: not enough cards present")
		}

		card, ok := cardValues[scanner.Text()]
		if !ok {
			return Hand{}, fmt.Errorf("unknown card type: %s", scanner.Text())
		}

		cards = append(cards, card)
	}

	scanner.Scan() // Scan the whitespace

	var bidStr string
	for scanner.Scan() {
		bidStr += scanner.Text()
	}

	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		return Hand{}, err
	}

	handType := typeComputer(cards)

	value := uint32(handType)
	value <<= 4
	value |= uint32(cards[0])
	value <<= 4
	value |= uint32(cards[1])
	value <<= 4
	value |= uint32(cards[2])
	value <<= 4
	value |= uint32(cards[3])
	value <<= 4
	value |= uint32(cards[4])

	return Hand{
		Cards: cards,
		Value: value,
		Bid:   bid,
	}, nil
}
