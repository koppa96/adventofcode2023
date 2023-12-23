package day10

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
)

var task1Cmd = &cobra.Command{
	Use:   "task1 <inputFile>",
	Short: "Solves task 1 for the provided input file.",
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

		start, err := FindStart(tiles)
		if err != nil {
			return err
		}

		directions := DirectionsFromStart(start, tiles)
		pos := start.Next(directions[0])
		lastDir := directions[0]
		steps := 1
		for pos != start {
			dirs := DirectionsByTile[tiles[pos.Row][pos.Col]]
			nextDir, ok := lo.Find(
				dirs, func(item Direction) bool {
					return item != lastDir.Opposite()
				},
			)

			if !ok {
				return fmt.Errorf("failed to find next direction at (%d, %d)", pos.Row, pos.Col)
			}

			pos = pos.Next(nextDir)
			lastDir = nextDir
			steps++
		}

		fmt.Println(steps / 2)

		return nil
	},
}

func FindStart(tiles [][]rune) (Position, error) {
	for row := range tiles {
		for col := range tiles[row] {
			if tiles[row][col] == 'S' {
				return Position{Row: row, Col: col}, nil
			}
		}
	}

	return Position{}, errors.New("no starting position found")
}

func DirectionsFromStart(position Position, tiles [][]rune) []Direction {
	var directions []Direction

	if position.Row != 0 {
		tileDirections := DirectionsByTile[tiles[position.Row-1][position.Col]]
		if lo.Contains(tileDirections, South) {
			directions = append(directions, South.Opposite())
		}
	}

	if position.Col != len(tiles[0])-1 {
		tileDirections := DirectionsByTile[tiles[position.Row][position.Col+1]]
		if lo.Contains(tileDirections, West) {
			directions = append(directions, West.Opposite())
		}
	}

	if position.Row != len(tiles)-1 {
		tileDirections := DirectionsByTile[tiles[position.Row+1][position.Col]]
		if lo.Contains(tileDirections, North) {
			directions = append(directions, North.Opposite())
		}
	}

	if position.Col != 0 {
		tileDirections := DirectionsByTile[tiles[position.Row][position.Col-1]]
		if lo.Contains(tileDirections, East) {
			directions = append(directions, East.Opposite())
		}
	}

	return directions
}
