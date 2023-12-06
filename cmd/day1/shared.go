package day1

import (
	"bufio"
	"os"
)

type DigitFinder func(str string) (int, error)

func asDigit(r rune) (int, bool) {
	if r >= 48 && r < 58 {
		return int(r - 48), true
	}

	return 0, false
}

func processFile(filename string, firstDigit DigitFinder, lastDigit DigitFinder) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first, err := firstDigit(line)
		if err != nil {
			panic(err)
		}

		last, err := lastDigit(line)
		if err != nil {
			panic(err)
		}

		num := first*10 + last
		sum += num
	}

	return sum, nil
}
