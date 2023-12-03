package lib

import (
	"bufio"
	"github.com/samber/lo"
	"os"
	"strconv"
)

func ParseEngineSchematics(fileName string) ([][]rune, error) {
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

func ParseNumber(runes []rune, pos int) (int, error) {
	startPos := pos
	for startPos > 0 && IsDigit(runes[startPos-1]) {
		startPos -= 1
	}

	endPos := pos
	for endPos < len(runes) && IsDigit(runes[endPos]) {
		endPos += 1
	}

	return strconv.Atoi(string(runes[startPos:endPos]))
}

func IsDigit(r rune) bool {
	return r >= 48 && r < 58
}

func NeighboringPartNumbers(runes [][]rune, row, col int) ([]int, error) {
	nums := make(map[int]struct{})
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if IsDigit(runes[row+i][col+j]) {
				num, err := ParseNumber(runes[row+i], col+j)
				if err != nil {
					return nil, err
				}

				nums[num] = struct{}{}
			}
		}
	}

	return lo.Keys(nums), nil
}
