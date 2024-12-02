package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

var spelledDigits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var spelledDigitsMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	var wg sync.WaitGroup
	cha := make(chan int, 10)

	finalSum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var texts []string

	for scanner.Scan() {
		texts = append(texts, scanner.Text())
	}
	result := make([]int, len(texts))

	fmt.Println(len(texts))
	for x, y := 0, 100; y <= len(texts); x, y = x+100, y+100 {
		if y == 1000 {
			y = 1001
		}
		wg.Add(1)
		go func(texts []string, cha chan int) {
			for _, text := range texts {
				if checkIfSpelledString(text) {
					cha <- checkString(text)
				} else {
					cha <- checkForDigits(text)
				}
			}
			// if I do not defer the wg.Done then it works????????
			wg.Done()
		}(texts[x:y], cha)
	}

	for i, _ := range result {
		result[i] = <-cha
	}

	wg.Wait()
	time.Sleep(1000 * time.Millisecond)

	for _, re := range result {
		finalSum += re
	}
	fmt.Println(finalSum)
}

func checkIfSpelledString(text string) bool {
	found := false
	for _, digit := range spelledDigits {
		if strings.Contains(text, digit) {
			found = true
		}
	}
	return found
}

func checkString(text string) int {
	firstInt := slidingWindowCheck(text, true)
	secondInt := slidingWindowCheck(text, false)
	return (firstInt * 10) + secondInt
}

func slidingWindowCheck(text string, left bool) int {
	if left {

		for x, y := 0, 5; y <= len(text); x, y = x+1, y+1 {
			found, number := checkSubString(text[x:y], left)
			if found {
				return number
			}
		}

	} else {

		for x, y := len(text)-5, len(text); x >= 0; x, y = x-1, y-1 {
			found, number := checkSubString(text[x:y], left)
			if found {
				return number
			}
		}
	}
	return 0
}

func checkSubString(text string, left bool) (bool, int) {
	if left {

		for x := 1; x <= len(text); x++ {
			if x >= 3 {
				found, number := searchSpelledDigits(text[:x])
				if found {
					return found, number
				}
			}

			for _, r := range text[:x] {
				if unicode.IsDigit(r) {
					number, _ := strconv.Atoi(string(r))
					return true, number
				}
			}
		}

	} else {

		for x := len(text) - 1; x >= 0; x-- {
			if len(text[x:]) >= 3 {
				found, number := searchSpelledDigits(text[x:])
				if found {
					return found, number
				}
			}

			for _, r := range reverse(text[x:]) {
				if unicode.IsDigit(r) {
					number, _ := strconv.Atoi(string(r))
					return true, number
				}
			}

		}
	}
	return false, 0
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func searchSpelledDigits(text string) (bool, int) {
	for x := 0; x < len(spelledDigits); x++ {
		if strings.Contains(text, spelledDigits[x]) {
			number := spelledDigitsMap[spelledDigits[x]]

			return true, number
		}
	}
	return false, 0

}

func checkForDigits(text string) int {
	var first rune
	var last rune

	for x := 0; x <= len(text)-1; x++ {
		currentRune := rune(text[x])
		if unicode.IsDigit(currentRune) {
			first = currentRune
			break
		}
	}

	for x := len(text) - 1; x >= 0; x-- {
		currentRune := rune(text[x])
		if unicode.IsDigit(currentRune) {
			last = currentRune
			break
		}
	}

	firstint, _ := strconv.Atoi(string(first))
	secondint, _ := strconv.Atoi(string(last))

	return (firstint * 10) + secondint
}
