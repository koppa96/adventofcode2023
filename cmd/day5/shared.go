package day5

import (
	"bufio"
	"strconv"
	"unicode"
)

func parseAndCombineMaps(scanner *bufio.Scanner) func(value int) int {
	seedToSoilMap, err := parseMap(scanner)
	if err != nil {
		panic(err)
	}

	soilToFertilizerMap, err := parseMap(scanner)
	if err != nil {
		panic(err)
	}

	fertilizerToWaterMap, err := parseMap(scanner)
	if err != nil {
		panic(err)
	}

	waterToLightMap, err := parseMap(scanner)
	if err != nil {
		panic(err)
	}

	lightToTemperatureMap, err := parseMap(scanner)
	if err != nil {
		panic(err)
	}

	temperatureToHumidityMap, err := parseMap(scanner)
	if err != nil {
		panic(err)
	}

	humidityToLocationMap, err := parseMap(scanner)
	if err != nil {
		panic(err)
	}

	return combineMaps(
		seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWaterMap,
		waterToLightMap,
		lightToTemperatureMap,
		temperatureToHumidityMap,
		humidityToLocationMap,
	)
}

func parseSeeds(scanner *bufio.Scanner) ([]int, error) {
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

func parseMap(scanner *bufio.Scanner) (RangeMap, error) {
	scanner.Scan() // map name

	var rangeMap RangeMap
	for scanner.Scan() && unicode.IsDigit([]rune(scanner.Text())[0]) {
		dest, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return RangeMap{}, err
		}

		scanner.Scan()
		source, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return RangeMap{}, err
		}

		scanner.Scan()
		length, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return RangeMap{}, err
		}

		rangeMap.Ranges = append(
			rangeMap.Ranges, Range{
				SourceStart:      source,
				DestinationStart: dest,
				Length:           length,
			},
		)
	}

	return rangeMap, nil
}

func combineMaps(maps ...RangeMap) func(value int) int {
	return func(value int) int {
		for _, m := range maps {
			value = m.Map(value)
		}

		return value
	}
}
