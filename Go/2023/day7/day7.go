package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var typeError = errors.New("Failed to determine type of Hand.")
var stregthMap = map[rune]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type Hand struct {
	cards      []rune
	typeOfHand int
	rank       int
	bid        int
	Winning    int
}

func (h *Hand) assignRank(index int) {
	h.rank = index
}

func (h *Hand) calcWinning() {
	h.Winning = h.rank * h.bid
}

func constructHandFromLine(line string) Hand {
	fmt.Println(line)
	split := strings.Split(line, " ")
	cards, bid := split[0], split[1]
	bid_, _ := strconv.Atoi(bid)
	var cardsAsRunes []rune
	for _, rune_ := range cards {
		cardsAsRunes = append(cardsAsRunes, rune_)
	}
	typeOfHand, err := determineTypeOfHand(&cardsAsRunes)
	if err != nil {
		panic(err)
	}
	return Hand{cards: cardsAsRunes, typeOfHand: typeOfHand, bid: bid_}
}

func determineTypeOfHand(cards *[]rune) (int, error) {
	// have to refactor this
	cardMap := make(map[rune]int)

	cards_ := *cards
	for _, rune_ := range cards_ {
		cardMap[rune_]++
	}
	js := cardMap['J']
	fmt.Println(js)

	switch len(cardMap) {
	// only unique cards
	case 5:
		switch js {
		case 1:
			// pair
			return onePair, nil
		case 0:
			// highest card
			return highCard, nil
		}
	// one pair
	case 4:
		switch js {
		case 2:
			// three of a kind
			return threeOfAKind, nil
		case 1:
			// three of a kind
			return threeOfAKind, nil
		case 0:
			// pair
			return onePair, nil
		}

	case 3:
		var switch_ bool
		for _, value := range cardMap {
			if value == 3 {
				switch_ = true
			}
		}
		switch switch_ {
		case true:
			// three of a kind
			// 3 J -> four of a kind
			// 1 J -> four of a kind
			// 0 J -> three of a kind
			switch js {
			case 3:
				return fourOfAKind, nil
			case 1:
				return fourOfAKind, nil
			case 0:
				return threeOfAKind, nil
			}
		// two pair
		// 2 J -> four of a kind
		// 1 J -> full house
		// 0 J -> two pair
		case false:
			switch js {
			case 2:
				return fourOfAKind, nil
			case 1:
				return fullHouse, nil
			case 0:
				return twoPair, nil
			}
		}
	case 2:
		var switch_ bool
		for _, value := range cardMap {
			if value == 3 {
				switch_ = true
			}
		}
		switch switch_ {
		// full house
		// 3J -> five of a kind
		// 2J -> five of a kind
		case true:
			switch js {
			case 3:
				return fiveOfAKind, nil
			case 2:
				return fiveOfAKind, nil
			case 0:
				return fullHouse, nil
			}
			// four of a kind
		// 4J -> five of a kind
		// 1J -> five of a kind
		case false:
			switch js {
			case 4:
				return fiveOfAKind, nil
			case 1:
				return fiveOfAKind, nil
			case 0:
				return fourOfAKind, nil
			}

		}
	case 1:
		return fiveOfAKind, nil
	}
	return 0, typeError
}

type byTypeOfHand []Hand

func (b byTypeOfHand) Len() int      { return len(b) }
func (b byTypeOfHand) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byTypeOfHand) Less(i, j int) bool {
	if b[i].typeOfHand != b[j].typeOfHand {
		return b[i].typeOfHand < b[j].typeOfHand
	}
	leftLesser := compareHands(&b[i], &b[j])
	return leftLesser
}

func compareHands(hand1 *Hand, hand2 *Hand) bool {
	cardsHand1 := hand1.cards
	cardsHand2 := hand2.cards

	var leftLesser bool
	for x := 0; x <= len(cardsHand1)-1; x++ {
		left := stregthMap[cardsHand1[x]]
		right := stregthMap[cardsHand2[x]]
		if left == right {
			continue
		}
		leftLesser = left < right
		break
	}
	return leftLesser
}

func assignRanksAndWinning(hands *[]Hand) {
	Hands := *hands
	for index, hand := range Hands {
		hand.assignRank(index + 1)
		hand.calcWinning()
		Hands[index] = hand
	}

	*hands = Hands
}

func calcTotalWinnings(hands []Hand) int {
	var sumWinnings int
	for _, hand := range hands {
		sumWinnings += hand.Winning
	}
	return sumWinnings
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	var lines []string
	var Hands []Hand

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for _, line := range lines {
		Hands = append(Hands, constructHandFromLine(line))
	}

	sort.Sort(byTypeOfHand(Hands))
	fmt.Println(Hands)
	assignRanksAndWinning(&Hands)
	for _, hand := range Hands {
		fmt.Println(string(hand.cards), hand.typeOfHand)
	}
	fmt.Println(Hands)
	totalWinnings := calcTotalWinnings(Hands)
	fmt.Println(totalWinnings)
}

// 251261032 too high
// 251355437 will also be too high
// 250757288
