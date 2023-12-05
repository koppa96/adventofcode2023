package lib

import (
	"bufio"
	"github.com/koppa96/adventofcode2023/lib/day5"
	"strconv"
	"unicode"
)

func ParseSeeds(scanner *bufio.Scanner) ([]int, error) {
	scanner.Scan() // seeds:

	var seeds []int
	for scanner.Scan() && scanner.Text() != "seed-to-soil" {
		seed, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		seeds = append(seeds, seed)
	}

	return seeds, nil
}

func ParseMap(scanner *bufio.Scanner) (day5.RangeMap, error) {
	scanner.Scan() // map name

	var rangeMap day5.RangeMap
	for scanner.Scan() && unicode.IsDigit([]rune(scanner.Text())[0]) {
		dest, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return day5.RangeMap{}, err
		}

		scanner.Scan()
		source, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return day5.RangeMap{}, err
		}

		scanner.Scan()
		length, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return day5.RangeMap{}, err
		}

		rangeMap.Ranges = append(
			rangeMap.Ranges, day5.Range{
				SourceStart:      source,
				DestinationStart: dest,
				Length:           length,
			},
		)
	}

	return rangeMap, nil
}

func CombineMaps(maps ...day5.RangeMap) func(value int) int {
	return func(value int) int {
		for _, m := range maps {
			value = m.Map(value)
		}

		return value
	}
}
