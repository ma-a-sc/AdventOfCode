package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	from string
	to   string
	// in the source text file the destination is the first number and the source number is the second
	// in the struct it is the opposite
	internalMap map[uint16]uint16
}

func (m Map) constructMapFromListOfStrings(s []string) Map {
	internalMap := make(map[uint16]uint16)
	for _, entry := range s {
		if strings.Contains(entry, "map:") {
			cutEntry := strings.Trim(entry, " map:")
			fromAndTOInfo := strings.Split(cutEntry, "-to-")
			m.from, m.to = fromAndTOInfo[0], fromAndTOInfo[1]
			continue
		}
		intsAsStrings := strings.Split(entry, " ")
		intsOfEntry := parseSliceOfStringsToInt(&intsAsStrings)

		destination, source, range_ := intsOfEntry[0], intsOfEntry[1], intsOfEntry[2]

		for x := 0; x < range_; x++ {
			internalMap[uint16(source+x)] = uint16(destination + x)
		}
	}

	m.internalMap = internalMap
	return m
}

var seeds []int
var listOfMaps []Map

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	// use the string as the next in line if there is no next in line you got the location

	var currentMapDetails []string
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for index, line := range lines {
		if len(line) == 0 {
			if len(currentMapDetails) != 0 {
				m := Map.constructMapFromListOfStrings(Map{}, currentMapDetails)
				listOfMaps = append(listOfMaps, m)
				currentMapDetails = nil
			}
			continue
		}
		// we need something like index line equal the last and then append it
		if index+1 == len(lines) {
			if len(currentMapDetails) != 0 {
				m := Map.constructMapFromListOfStrings(Map{}, currentMapDetails)
				listOfMaps = append(listOfMaps, m)
				currentMapDetails = nil
			}
			continue
		}
		if strings.HasPrefix(line, "seeds:") {
			seeds = append(seeds, getSeeds(line)...)
			continue
		}
		currentMapDetails = append(currentMapDetails, line)

	}
	var lowestLocation uint16
	for _, seed := range seeds {
		seedsLocation := getLowestLocationForSeed(&seed)
		if lowestLocation == 0 {
			lowestLocation = seedsLocation
		}
		if seedsLocation < lowestLocation {
			lowestLocation = seedsLocation
		}
	}
	fmt.Println(lowestLocation)

}

func getLowestLocationForSeed(seed *int) uint16 {
	key := "seed"
	currentNumber := uint16(*seed)
	for {
		relevantMap, found := getRelevantMap(key)
		if !found {
			panic("Not found any map. Programming error likely.")
		}
		// update the lookup
		var nextNumber uint16
		var ok bool
		nextNumber, ok = relevantMap.internalMap[currentNumber]

		fmt.Println(relevantMap.from, relevantMap.to, nextNumber)
		if ok {
			currentNumber = nextNumber
		}
		fmt.Println(currentNumber)
		key = relevantMap.to
		if key == "location" {
			break
		}
	}
	return currentNumber
}

func getRelevantMap(key string) (Map, bool) {
	for _, map_ := range listOfMaps {
		if map_.from == key {
			return map_, true
		}
	}
	return Map{}, false
}

func getSeeds(line string) []int {
	seeedsAsStrings := strings.Split(line, " ")[1:]
	intSlice := parseSliceOfStringsToInt(&seeedsAsStrings)
	return intSlice
}

func parseSliceOfStringsToInt(strings *[]string) []int {
	stringsToParse := *strings
	var toIntParsedStrings []int

	for _, entry := range stringsToParse {
		num, err := strconv.Atoi(entry)
		if err != nil {
			fmt.Println("Failed to convert", entry, "to int: ", err)
		}
		toIntParsedStrings = append(toIntParsedStrings, num)
	}

	return toIntParsedStrings
}
