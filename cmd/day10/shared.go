package day10

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case West:
		return East
	case East:
		return West
	}

	panic("invalid direction")
}

func (d Direction) IsHorizontal() bool {
	return d == West || d == East
}

type Position struct {
	Row, Col int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.Row, p.Col)
}

func (p Position) Next(d Direction) Position {
	switch d {
	case North:
		return Position{Row: p.Row - 1, Col: p.Col}
	case East:
		return Position{Row: p.Row, Col: p.Col + 1}
	case South:
		return Position{Row: p.Row + 1, Col: p.Col}
	case West:
		return Position{Row: p.Row, Col: p.Col - 1}
	}

	panic("invalid direction")
}

var DirectionsByTile = map[rune][]Direction{
	'|': {North, South},
	'-': {West, East},
	'L': {North, East},
	'J': {North, West},
	'7': {South, West},
	'F': {South, East},
	'.': {},
	'S': {},
}
