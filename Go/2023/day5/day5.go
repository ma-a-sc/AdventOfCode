package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var seeds []int
	minLocation := math.MaxInt
	var wg sync.WaitGroup
	results := make(chan int)

	scanner := bufio.NewScanner(file)
	// use the string as the next in line if there is no next in line you got the location
	var lines []string
	var splitlines [][][]int

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	splitlines = prepBlocks(lines[1:])

	seeds = append(seeds, getSeeds(lines[0])...)
	for len(seeds) > 2 {
		startSeed, numSeeds := seeds[0], seeds[1]
		endSeed := startSeed + numSeeds
		seeds = seeds[2:]

		wg.Add(1)
		go func(start, end int, splitlines *[][][]int) {
			defer wg.Done()
			var location int
			for y := start; y <= end; y++ {
				// now in here we implemenet the thing to get the location
				location = getlocation(y, splitlines)
				results <- location
			}
		}(startSeed, endSeed, &splitlines)
	}

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

	//fmt.Println(splitlines)
	//var location []int
	//for _, seed := range seeds {
	//	location = append(location, getlocation(seed, splitlines))
	//}
	//fmt.Println(minFind(location))
}

func prepBlocks(lines []string) [][][]int {
	var currentBlock [][]int
	var splitlines [][][]int
	for index, line := range lines[2:] {
		if len(line) == 0 {
			splitlines = append(splitlines, currentBlock)
			currentBlock = nil
			continue
		}
		if strings.HasSuffix(line, "map:") {
			continue
		}
		if index == len(lines[2:])-1 {
			intsAsStrings := strings.Split(line, " ")
			intsAsInts := parseSliceOfStringsToInt(&intsAsStrings)
			currentBlock = append(currentBlock, intsAsInts)
			splitlines = append(splitlines, currentBlock)
			currentBlock = nil
			continue
		}
		intsAsStrings := strings.Split(line, " ")
		intsAsInts := parseSliceOfStringsToInt(&intsAsStrings)
		currentBlock = append(currentBlock, intsAsInts)
	}
	return splitlines
}

// lets get rid of the string parsing in here
func getlocation(seed int, arrayOfBlocks *[][][]int) int {
	currentInt := seed
	arrayOfBlocks_ := *arrayOfBlocks
	for _, block := range arrayOfBlocks_ {
		currentInt = checkBlock(block, currentInt)
	}
	return currentInt
}

func checkBlock(block [][]int, currentInt int) int {
	cInt := currentInt
	for _, line := range block {
		var found bool
		// have to special case this
		cInt, found = checkIfIntIsInRangeAndGetMappedInt(cInt, line)
		if found {
			break
		}

	}
	return cInt
}

// that thing gets the offset even if the int was not found?
func checkIfIntIsInRangeAndGetMappedInt(number int, parsedInts []int) (int, bool) {
	destination, source, range_ := parsedInts[0], parsedInts[1], parsedInts[2]
	if number >= source && number <= source+range_ {
		offset := getOffset(number, source)
		mappedInt := getMappedInt(destination, offset)
		return mappedInt, true
	}
	return number, false
}

func getMappedInt(destination int, offset int) int {
	return destination + offset
}

func getOffset(number int, source int) int {
	// can this be negative?
	return (number - source) * -1
}

func getSeeds(line string) []int {
	seeedsAsStrings := strings.Split(line, " ")[1:]
	intSlice := parseSliceOfStringsToInt(&seeedsAsStrings)
	return intSlice
}

func getSeedsRangeVersion(line string) []int {
	seeedsAsStrings := strings.Split(line, " ")[1:]
	intSlice := parseSliceOfStringsToInt(&seeedsAsStrings)
	var allInts []int
	for x := 0; x <= len(intSlice)-1; x = x + 2 {
		start := intSlice[x]
		counter := intSlice[x+1]

		for y := 0; y <= counter; y++ {
			allInts = append(allInts, start+y)
		}
	}
	return allInts
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

// anwser too high 662197086
// anwser too high 52510810

// 52510809
