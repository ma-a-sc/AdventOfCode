package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type CategoryMap struct {
	source      int
	destination int
	length      int
}

type Category struct {
	categoryMappings []CategoryMap
}

type Categories struct {
	seedToSoil          Category
	soilToFertilizer    Category
	fertilizerToWater   Category
	waterToLight        Category
	lightToTemperature  Category
	tempatureToHumidity Category
	humidityToLocation  Category
}

func (c *Category) getDestinationForSource(sourceId int) int {
	for _, categoryMapping := range c.categoryMappings {
		min := categoryMapping.source
		max := categoryMapping.source + categoryMapping.length
		if sourceId >= min && sourceId <= max {
			return categoryMapping.destination + sourceId - min
		}
	}
	return sourceId
}

func (c *Categories) getLocationForSeed(seedId int) int {
	soilId := c.seedToSoil.getDestinationForSource(seedId)
	fertilizerId := c.soilToFertilizer.getDestinationForSource(soilId)
	waterId := c.fertilizerToWater.getDestinationForSource(fertilizerId)
	lightId := c.waterToLight.getDestinationForSource(waterId)
	tempId := c.lightToTemperature.getDestinationForSource(lightId)
	humidId := c.tempatureToHumidity.getDestinationForSource(tempId)
	return c.humidityToLocation.getDestinationForSource(humidId)
}

func getCategoryForLines(lines []string) Category {
	// Pop Header
	_, lines = lines[0], lines[1:]
	mappings := []CategoryMap{}
	for _, line := range lines {
		if line == "" {
			break
		}
		map_values := strings.Fields(line)

		destination, err := strconv.Atoi(map_values[0])
		if err != nil {
			panic(err)
		}

		source, err := strconv.Atoi(map_values[1])
		if err != nil {
			panic(err)
		}

		length, err := strconv.Atoi(map_values[2])
		if err != nil {
			panic(err)
		}

		categoryMap := CategoryMap{
			source:      source,
			destination: destination,
			length:      length,
		}
		mappings = append(mappings, categoryMap)
	}
	return Category{categoryMappings: mappings}
}

func parseInput(fileContents string) ([]int, Categories, error) {
	lines := strings.Split(fileContents, "\n")

	first_line, lines := lines[0], lines[1:]

	seedsStrings := strings.Fields(strings.Split(first_line, ": ")[1])
	seeds := []int{}
	for _, seedString := range seedsStrings {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			return seeds, Categories{}, err
		}
		seeds = append(seeds, seed)
	}

	// Pop empty line
	_, lines = lines[0], lines[2:]

	seedsToSoilMappingStrings := []string{}
	for idx, line := range lines {
		if line == "" {
			lines = lines[idx+1:]
			break
		}
		seedsToSoilMappingStrings = append(seedsToSoilMappingStrings, line)
	}
	soilToFertilizerMappingStrings := []string{}
	for idx, line := range lines {
		if line == "" {
			lines = lines[idx+1:]
			break
		}
		soilToFertilizerMappingStrings = append(soilToFertilizerMappingStrings, line)
	}
	fertilizerToWaterMappingStrings := []string{}
	for idx, line := range lines {
		if line == "" {
			lines = lines[idx+1:]
			break
		}
		fertilizerToWaterMappingStrings = append(fertilizerToWaterMappingStrings, line)
	}
	waterToLightMappingStrings := []string{}
	for idx, line := range lines {
		if line == "" {
			lines = lines[idx+1:]
			break
		}
		waterToLightMappingStrings = append(waterToLightMappingStrings, line)
	}
	lightToTemperatureMappingStrings := []string{}
	for idx, line := range lines {
		if line == "" {
			lines = lines[idx+1:]
			break
		}
		lightToTemperatureMappingStrings = append(lightToTemperatureMappingStrings, line)
	}
	temperatureToHumidityMappingStrings := []string{}
	for idx, line := range lines {
		if line == "" {
			lines = lines[idx+1:]
			break
		}
		temperatureToHumidityMappingStrings = append(temperatureToHumidityMappingStrings, line)
	}
	humidityToLocationMappingStrings := []string{}
	for _, line := range lines {
		if line == "" {
			break
		}
		humidityToLocationMappingStrings = append(humidityToLocationMappingStrings, line)
	}

	categories := Categories{
		seedToSoil:          getCategoryForLines(seedsToSoilMappingStrings),
		soilToFertilizer:    getCategoryForLines(soilToFertilizerMappingStrings),
		fertilizerToWater:   getCategoryForLines(fertilizerToWaterMappingStrings),
		waterToLight:        getCategoryForLines(waterToLightMappingStrings),
		lightToTemperature:  getCategoryForLines(lightToTemperatureMappingStrings),
		tempatureToHumidity: getCategoryForLines(temperatureToHumidityMappingStrings),
		humidityToLocation:  getCategoryForLines(humidityToLocationMappingStrings),
	}

	return seeds, categories, nil
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable info:", err)
	}
	exeDir := filepath.Dir(exePath)
	filePath := filepath.Join(exeDir, "input.txt")
	//filePath := filepath.Join(exeDir, "test_input.txt"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	seeds, categories, err := parseInput(string(content))
	if err != nil {
		panic(err)
	}

	minLocation := math.MaxInt
	results := make(chan int)
	var wg sync.WaitGroup

	for len(seeds) > 0 {
		startSeed, numSeeds := seeds[0], seeds[1]
		endSeed := startSeed + numSeeds - 1
		seeds = seeds[2:]

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localMin := math.MaxInt
			for seed := start; seed <= end; seed++ {
				location := categories.getLocationForSeed(seed)
				if location < localMin {
					localMin = location
				}
			}
			results <- localMin
		}(startSeed, endSeed)
	}

	// Close results channel once all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for location := range results {
		if location < minLocation {
			minLocation = location
		}
	}

	fmt.Println(minLocation)
}
