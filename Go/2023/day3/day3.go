package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	array := getArrayFromScanner(scanner)

	arrayOfNumberIndexes := getNumberIndexes(array)
	arrayOfSymbolIndexes := getSymbolIndexes(array)
	numbersNearSymbols := checkNumberNearSymbols(arrayOfNumberIndexes, arrayOfSymbolIndexes, array)

	sumOfNumbers := 0
	for _, number := range numbersNearSymbols {
		sumOfNumbers += number
	}
	fmt.Println(sumOfNumbers)
	// too low 536262

}

func checkNumberNearSymbols(arrayOfNumberIndexes [][][]int, arrayOfSymbolIndexes [][2]int, array [][]rune) []int {
	var finalArr []int
	for _, numberIndexes := range arrayOfNumberIndexes {
		found := false
		for _, digitXY := range numberIndexes {
			// could rewrite this part to get all the surrounding indexes for the whole number???
			x := digitXY[0]
			y := digitXY[1]
			// this probably causes duplicates or no it does not ????
			nearXYCoordinates := getNearXYCoordinates(x, y)
			found = checkSymbolArray(nearXYCoordinates, arrayOfSymbolIndexes)
			if found {
				break
			}
		}
		if found {
			// number indexes is an array of three x,y coordinates
			int_ := constructInt(numberIndexes, array)
			finalArr = append(finalArr, int_)
			found = false
		}

	}

	return finalArr
}

func constructInt(numberIndexes [][]int, array [][]rune) int {
	lineIndex := numberIndexes[0][0]
	yOfStart := numberIndexes[0][1]
	yOfEnd := numberIndexes[len(numberIndexes)-1][1]
	var runes_ []rune
	for y := yOfStart; y <= yOfEnd; y++ {
		runes_ = append(runes_, array[lineIndex][y])
	}

	intAsString := string(runes_)
	int_, _ := strconv.Atoi(intAsString)
	return int_
}

func checkSymbolArray(nearXY [][2]int, arrayOfSymbolIndexes [][2]int) bool {
	for _, entry := range nearXY {
		for _, secondEntry := range arrayOfSymbolIndexes {
			if entry == secondEntry {
				return true
			}
		}
	}
	return false
}

func getNearXYCoordinates(x int, y int) [][2]int {
	possibleX := []int{x - 1, x, x + 1}
	possibleY := []int{y - 1, y, y + 1}
	combinations := getCombinations(possibleX, possibleY, [2]int{x, y})
	return combinations
}

func getCombinations(possibleX []int, possibleY []int, checkArr [2]int) [][2]int {
	var finalArr [][2]int

	for _, xEntry := range possibleX {
		for _, yEntry := range possibleY {
			possibleArr := [2]int{xEntry, yEntry}

			if possibleArr != checkArr {
				finalArr = append(finalArr, possibleArr)
			}
		}
	}
	return finalArr
}

func getSymbolIndexes(array [][]rune) [][2]int {
	var finalArray [][2]int

	for lineIndex, line := range array {
		for columnIndex, entry := range line {
			if entry != 46 && !unicode.IsDigit(entry) {
				finalArray = append(finalArray, [2]int{lineIndex, columnIndex})
			}
		}
	}
	return finalArray
}

func getNumberIndexes(array [][]rune) [][][]int {
	var finalArray [][][]int
	for lineIndex, line := range array {
		var tempIndexArray [][]int
		previous := false
		for columnIndex, entry := range line {
			xyOfRune := []int{lineIndex, columnIndex}
			if unicode.IsDigit(entry) && previous {
				tempIndexArray = append(tempIndexArray, xyOfRune)
				previous = true
				continue
			}
			if !unicode.IsDigit(entry) && previous {
				finalArray = append(finalArray, tempIndexArray)
				tempIndexArray = nil
				previous = false
				continue

			}
			if unicode.IsDigit(entry) && !previous {
				tempIndexArray = append(tempIndexArray, xyOfRune)
				previous = true
				continue
			}
			if !unicode.IsDigit(entry) && !previous {
				previous = false
				continue
			}
		}
	}
	return finalArray
}

func getArrayFromScanner(scanner *bufio.Scanner) [][]rune {
	var array [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		var runeArray []rune
		runeArray = append(runeArray, []rune(line)...)
		array = append(array, runeArray)
	}

	return array
}
