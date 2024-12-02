package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Records struct {
	races []Race
}

type Race struct {
	fullTimeLengthOfRace int
	recordDistance       int
}

func (r *Race) getNumberOfWinningMoves() int {
	var winningHoldings []int
	for hold := 0; hold <= r.fullTimeLengthOfRace; hold++ {
		acceleration := hold * 1
		remainingTime := r.fullTimeLengthOfRace - hold
		travaledDistance := acceleration * remainingTime

		if travaledDistance > r.recordDistance {
			winningHoldings = append(winningHoldings, hold)
		}
	}
	return len(winningHoldings)
}

func (r *Records) getNumberOfWaysToBeatRecord() int64 {
	var winningMovesEachGame []int
	for _, race := range r.races {
		winningMovesEachGame = append(winningMovesEachGame, race.getNumberOfWinningMoves())
	}
	return getMul(winningMovesEachGame)
}

func getMul(arr []int) int64 {
	var allMoves int64
	allMoves = 1
	for _, entry := range arr {
		allMoves = allMoves * int64(entry)
	}
	return allMoves
}

func getRecordsFromLines(lines []string) Records {
	var times []int
	var distances []int
	for _, line := range lines {
		if strings.HasPrefix(line, "Time:") {
			trimmed := strings.Trim(line, "Time:")
			split := strings.Fields(trimmed)
			ints := trimSpacesAndConvertToInt(&split)
			times = append(times, ints[:]...)
		}
		if strings.HasPrefix(line, "Distance:") {
			trimmed := strings.Trim(line, "Distance:")
			split := strings.Fields(trimmed)
			ints := trimSpacesAndConvertToInt(&split)
			distances = append(distances, ints[:]...)
		}
	}
	fmt.Println(times, distances)
	var races []Race

	for x := 0; x <= len(times)-1; x++ {
		races = append(races, Race{fullTimeLengthOfRace: times[x], recordDistance: distances[x]})
	}

	return Records{races: races}
}

func getRecordFromLinesPart2(lines []string) Records {
	var time int
	var distance int
	for _, line := range lines {
		if strings.HasPrefix(line, "Time:") {
			trimmed := strings.ReplaceAll(line, "Time:", "")
			trimmed = strings.ReplaceAll(trimmed, " ", "")
			time, _ = strconv.Atoi(trimmed)
		}
		if strings.HasPrefix(line, "Distance:") {
			trimmed := strings.ReplaceAll(line, "Distance:", "")
			trimmed = strings.ReplaceAll(trimmed, " ", "")
			fmt.Println(trimmed)
			distance, _ = strconv.Atoi(trimmed)
		}
	}

	races := []Race{{fullTimeLengthOfRace: time, recordDistance: distance}}
	fmt.Println(races)
	return Records{races: races}
}

func trimSpacesAndConvertToInt(s *[]string) []int {
	ss := *s
	var intsToReturn []int
	for _, entry := range ss {
		trimmedEntry := strings.Trim(entry, " ")
		for strings.Contains(trimmedEntry, " ") {
			trimmedEntry = strings.Trim(trimmedEntry, " ")
		}
		converted, _ := strconv.Atoi(trimmedEntry)
		intsToReturn = append(intsToReturn, converted)
	}
	return intsToReturn
}

func main() {

	linesTest := []string{"Time:      7  15   30", "Distance:  9  40  200"}
	lines := []string{"Time:        47     84     74     67", "Distance:   207   1394   1209   1014"}
	fmt.Println(lines, linesTest)
	fmt.Println()
	fmt.Println()

	records := getRecordFromLinesPart2(lines)

	num := records.getNumberOfWaysToBeatRecord()
	fmt.Println(num)

}
