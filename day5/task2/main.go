package main

import (
	"bufio"
	"fmt"
	"github.com/koppa96/adventofcode2023/lib"
	"math"
	"os"
	"time"
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

	start := time.Now()
	min := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		fmt.Printf("Processing seed range %d/%d\n", i/2, len(seeds)/2)

		start := seeds[i]
		count := seeds[i+1]

		for j := 0; j < count; j++ {
			location := seedToLocation(start + j)
			if location < min {
				min = location
			}
		}
	}

	fmt.Println(min)
	fmt.Printf("Completed in %s", time.Since(start))
}
