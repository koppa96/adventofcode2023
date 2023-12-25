package day11

import (
	"github.com/koppa96/adventofcode2023/cmd/day10"
)

type Universe struct {
	emptyRows []int
	emptyCols []int
	Galaxies  []day10.Position
}

func ParseUniverse(runes [][]rune) Universe {
	var galaxies []day10.Position

	var emptyRows []int
	for i := range runes {
		empty := true
		for j := range runes[i] {
			if !isEmpty(runes[i][j]) {
				empty = false
				galaxies = append(galaxies, day10.Position{Row: i, Col: j})
			}
		}

		if empty {
			emptyRows = append(emptyRows, i)
		}
	}

	var emptyCols []int
	for i := range runes[0] {
		if isColEmpty(runes, i) {
			emptyCols = append(emptyCols, i)
		}
	}

	return Universe{
		emptyRows: emptyRows,
		emptyCols: emptyCols,
		Galaxies:  galaxies,
	}
}

func (u Universe) DistanceBetween(start day10.Position, end day10.Position, expansionCoefficient int) int {
	distance := abs(start.Row-end.Row) + abs(start.Col-end.Col)
	for _, row := range u.emptyRows {
		if isBetween(row, start.Row, end.Row) {
			distance += expansionCoefficient - 1
		}
	}

	for _, col := range u.emptyCols {
		if isBetween(col, start.Col, end.Col) {
			distance += expansionCoefficient - 1
		}
	}

	return distance
}

func isBetween(value, bound1, bound2 int) bool {
	return value-bound1 >= 0 && value-bound2 <= 0 || value-bound1 <= 0 && value-bound2 >= 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func isColEmpty(runes [][]rune, col int) bool {
	for i := range runes {
		if !isEmpty(runes[i][col]) {
			return false
		}
	}

	return true
}

func isEmpty(r rune) bool {
	return r == '.'
}
