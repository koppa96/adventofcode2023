package main

import (
	"bufio"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	seeds, err := lib.ParseSeeds(scanner)
	if err != nil {
		panic(err)
	}

	seedToSoilMap, err := lib.ParseMap(scanner)
	if err != nil {
		panic(err)
	}

	soilToFertilizerMap, err := lib.ParseMap(scanner)
	if err != nil {
		panic(err)
	}

	fertilizerToWaterMap, err := lib.ParseMap(scanner)
	if err != nil {
		panic(err)
	}

	waterToLightMap, err := lib.ParseMap(scanner)
	if err != nil {
		panic(err)
	}

	lightToTemperatureMap, err := lib.ParseMap(scanner)
	if err != nil {
		panic(err)
	}

	temperatureToHumidityMap, err := lib.ParseMap(scanner)
	if err != nil {
		panic(err)
	}

	humidityToLocationMap, err := lib.ParseMap(scanner)
	if err != nil {
		panic(err)
	}

	seedToLocation := lib.CombineMaps(
		seedToSoilMap,
		soilToFertilizerMap,
		fertilizerToWaterMap,
		waterToLightMap,
		lightToTemperatureMap,
		temperatureToHumidityMap,
		humidityToLocationMap,
	)

	min := seedToLocation(seeds[0])
	for i := 1; i < len(seeds); i++ {
		location := seedToLocation(seeds[i])
		if location < min {
			min = location
		}
	}

	fmt.Println(min)
}
