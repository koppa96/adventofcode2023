package day3

import (
	"bufio"
	"github.com/samber/lo"
	"os"
	"strconv"
)

func parseEngineSchematics(fileName string) ([][]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var runes [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		runes = append(runes, []rune(line))
	}

	return runes, nil
}

func parseNumber(runes []rune, pos int) (int, error) {
	startPos := pos
	for startPos > 0 && isDigit(runes[startPos-1]) {
		startPos -= 1
	}

	endPos := pos
	for endPos < len(runes) && isDigit(runes[endPos]) {
		endPos += 1
	}

	return strconv.Atoi(string(runes[startPos:endPos]))
}

func isDigit(r rune) bool {
	return r >= 48 && r < 58
}

func neighboringPartNumbers(runes [][]rune, row, col int) ([]int, error) {
	nums := make(map[int]struct{})
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if isDigit(runes[row+i][col+j]) {
				num, err := parseNumber(runes[row+i], col+j)
				if err != nil {
					return nil, err
				}

				nums[num] = struct{}{}
			}
		}
	}

	return lo.Keys(nums), nil
}
