package lib

import (
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Draws []map[string]int
}

func ParseGame(game string) (Game, error) {
	gameParts := strings.Split(game, ": ")

	id, err := parseGameId(gameParts[0])
	if err != nil {
		return Game{}, err
	}

	draws, err := parseDraws(gameParts[1])
	if err != nil {
		return Game{}, err
	}

	return Game{
		Id:    id,
		Draws: draws,
	}, nil
}

func parseGameId(linePart string) (int, error) {
	parts := strings.Split(linePart, " ")
	return strconv.Atoi(parts[1])
}

func parseDraws(linePart string) ([]map[string]int, error) {
	var parsedDraws []map[string]int

	draws := strings.Split(linePart, "; ")
	for _, draw := range draws {
		parsedDraw := make(map[string]int)
		colorDraws := strings.Split(draw, ", ")

		for _, colorDraw := range colorDraws {
			colorDrawParts := strings.Split(colorDraw, " ")
			color := colorDrawParts[1]

			amount, err := strconv.Atoi(colorDrawParts[0])
			if err != nil {
				return nil, err
			}

			parsedDraw[color] = amount
		}

		parsedDraws = append(parsedDraws, parsedDraw)
	}

	return parsedDraws, nil
}
