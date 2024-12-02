package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type card struct {
	id             int
	instances      int
	winningNumbers int
	left           []int
	right          []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	lines = append(lines, getLines(scanner)...)

	cards := make(map[int]card)

	getCards(lines, &cards)
	attachWinningNumbers(&cards)

	cards = processWinningNumbersIncreaseInstances(cards)
	fmt.Println(cards)
	sumOfInstances := getSumOfInstances(cards)
	fmt.Println(sumOfInstances)
}

func getSumOfInstances(cards map[int]card) int {
	var sumOfCardInstances int
	for _, currentCard := range cards {
		sumOfCardInstances += currentCard.instances
	}

	return sumOfCardInstances
}

func processWinningNumbersIncreaseInstances(cards map[int]card) map[int]card {
	for key := 1; key <= len(cards); key++ {
		currentCard := cards[key]
		numberOfCardsToIncrement := key + currentCard.winningNumbers
		for count := 1; count <= currentCard.instances; count++ {
			for ID := key + 1; ID <= numberOfCardsToIncrement; ID++ {
				if ID <= len(cards) {
					furtherCard := cards[ID]
					furtherCard.instances++
					cards[ID] = furtherCard
				}
			}
		}
	}
	return cards
}

func attachWinningNumbers(cards *map[int]card) {
	cardsMap := *cards
	for key, currentCard := range cardsMap {
		for _, number := range currentCard.left {
			if slices.Contains(currentCard.right, number) {
				currentCard.winningNumbers += 1
				cardsMap[key] = currentCard
			}
		}
	}
	*cards = cardsMap
}

func getCards(lines []string, cards *map[int]card) {
	cardsMap := *cards
	for _, line := range lines {
		lineSplit := strings.Split(line, ":")
		id := getID(lineSplit[0])
		lineCardStripped := lineSplit[1]
		fmt.Println(lineSplit[0])
		fmt.Println(id)
		linesSeparated := strings.Split(lineCardStripped, "|")

		left := processNumbers(strings.Trim(linesSeparated[0], " "))
		right := processNumbers(strings.Trim(linesSeparated[1], " "))

		cardsMap[id] = card{id: id, instances: 1, winningNumbers: 0, left: left, right: right}
	}
	*cards = cardsMap
}

func getID(s string) int {
	id := strings.TrimPrefix(s, "Card ")
	for strings.Contains(id, " ") {
		id = strings.Trim(id, " ")
	}
	var ints []int
	for _, rune_ := range id {
		if unicode.IsDigit(rune_) {
			ints = append(ints, int(rune_))
		}
	}
	intID, _ := strconv.Atoi(id)
	return intID
}

// not needed anymore was part of part one
func getWinningNumberCount(strippedLines []map[string][]int) int {
	var correctNumberCounter int
	for _, entry := range strippedLines {
		for _, number := range entry["left"] {
			if slices.Contains(entry["right"], number) {
				correctNumberCounter += 1
			}
		}
	}
	return correctNumberCounter
}

// not needed anymore was part of part one
func getPoints(strippedLines []map[string][]int) int {
	var points int
	for _, entry := range strippedLines {
		var correctNumberCounter int
		for _, number := range entry["left"] {
			if slices.Contains(entry["right"], number) {
				correctNumberCounter += 1
			}
		}
		points += getPointsGame(correctNumberCounter)
	}
	return points
}

// not needed anymore was part of part one
func getPointsGame(correctNumberCounter int) int {
	var points int
	for x := 0; x < correctNumberCounter; x++ {
		if x == 0 {
			points += 1
		} else {
			points *= 2
		}
	}
	return points
}

// part of part one
func getStrippedLines(lines []string) []map[string][]int {
	var processedLines []map[string][]int

	for _, line := range lines {
		currrentLine := make(map[string][]int)
		lineCardStripped := strings.Split(line, ":")[1]
		linesSeparated := strings.Split(lineCardStripped, "|")

		currrentLine["left"] = processNumbers(strings.Trim(linesSeparated[0], " "))
		currrentLine["right"] = processNumbers(strings.Trim(linesSeparated[1], " "))

		processedLines = append(processedLines, currrentLine)
	}

	return processedLines
}

func processNumbers(numbers string) []int {
	splitNumbers := strings.Split(numbers, " ")
	var processedNumbers []int
	for _, number := range splitNumbers {
		if len(number) <= 0 {
			continue
		}
		num, _ := strconv.Atoi(number)
		processedNumbers = append(processedNumbers, num)
	}

	return processedNumbers
}

func getLines(scanner *bufio.Scanner) []string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
