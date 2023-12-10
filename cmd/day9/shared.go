package day9

import (
	"bufio"
	"strconv"
	"strings"
)

func ParseSequence(line string) ([]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	var seq []int
	for scanner.Scan() {
		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		seq = append(seq, item)
	}

	return seq, nil
}

func Differentiate(seq []int) []int {
	var result []int
	for i := 1; i < len(seq); i++ {
		result = append(result, seq[i]-seq[i-1])
	}

	return result
}

func AllZero(seq []int) bool {
	for _, val := range seq {
		if val != 0 {
			return false
		}
	}

	return true
}
