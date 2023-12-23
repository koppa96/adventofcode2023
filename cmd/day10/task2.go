package day10

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
)

var task2Cmd = &cobra.Command{
	Use:   "task2 <inputFile>",
	Short: "Solves task 2 for the provided input file.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(args[0])
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)

		var tiles [][]rune
		for scanner.Scan() {
			tiles = append(tiles, []rune(scanner.Text()))
		}

		path, err := WalkPath(tiles)
		if err != nil {
			return err
		}

		var inside int
		for row := range tiles {
			for col := range tiles[row] {
				if isInside(Position{Row: row, Col: col}, tiles, path) {
					fmt.Printf("Inside: (%d, %d)\n", row, col)
					inside++
				}
			}
		}

		fmt.Println(inside)

		return nil
	},
}

func isInside(pos Position, tiles [][]rune, path map[Position]struct{}) bool {
	if _, ok := path[pos]; ok {
		return false
	}

	i := pos
	crossings := 0
	var verticalConnectionInDirection *Direction
	for i.Row != 0 {
		i = i.Next(North)
		if _, ok := path[i]; ok {
			// We only need to increase the amount of crossings if the vertically connected path elements are like this shape
			// F     7
			// |* or |*
			// J     L
			// If the entry of the vertically connected path elements is from the same direction as the out direction
			// The path to the map's edge doesn't cross it, just touch it.
			if connectedVertically(i, tiles) {
				if verticalConnectionInDirection == nil {
					horizontalDir, _ := horizontalDirection(i, tiles)
					verticalConnectionInDirection = &horizontalDir
				} else if horizontalDir, ok := horizontalDirection(i, tiles); ok {
					if horizontalDir != *verticalConnectionInDirection {
						crossings++
					}

					verticalConnectionInDirection = nil
				}
			} else {
				crossings++
			}
		}
	}

	return crossings%2 != 0
}

func horizontalDirection(pos Position, tiles [][]rune) (Direction, bool) {
	tile := tiles[pos.Row][pos.Col]
	directions := DirectionsByTile[tile]
	for _, d := range directions {
		if d.IsHorizontal() {
			return d, true
		}
	}

	return 0, false
}

func connectedVertically(pos Position, tiles [][]rune) bool {
	tile := tiles[pos.Row][pos.Col]
	return lo.Some(DirectionsByTile[tile], []Direction{North, South})
}

func WalkPath(tiles [][]rune) (map[Position]struct{}, error) {
	positions := make(map[Position]struct{})
	start, err := FindStart(tiles)
	if err != nil {
		return nil, err
	}

	positions[start] = struct{}{}

	directions := DirectionsFromStart(start, tiles)
	DirectionsByTile['S'] = directions

	pos := start.Next(directions[0])
	lastDir := directions[0]
	for pos != start {
		positions[pos] = struct{}{}
		dirs := DirectionsByTile[tiles[pos.Row][pos.Col]]
		nextDir, ok := lo.Find(
			dirs, func(item Direction) bool {
				return item != lastDir.Opposite()
			},
		)

		if !ok {
			return nil, fmt.Errorf("failed to find next direction at (%d, %d)", pos.Row, pos.Col)
		}

		pos = pos.Next(nextDir)
		lastDir = nextDir
	}

	return positions, nil
}
