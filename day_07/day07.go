package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputLines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Part 1
	handList := buildHandList(inputLines, false)
	orderHands(handList, cardStrengthPart1)
	fmt.Println("Part 1:", calculateWinnings(handList))

	// Part 2
	handList = buildHandList(inputLines, true)
	orderHands(handList, cardStrengthPart2)
	fmt.Println("Part 2:", calculateWinnings(handList))
}

func buildHandList(inputLines []string, jokers bool) []Hand {
	handList := []Hand{}

	for _, line := range inputLines {
		pos := strings.Index(line, " ")
		cardsStr, bidStr := line[:pos], line[pos+1:]
		cards := strings.Split(cardsStr, "")
		bid, _ := strconv.Atoi(bidStr)
		handType := calculateHandType(cards, jokers)
		hand := Hand{
			cards:    cards,
			handType: handType,
			bid:      bid,
		}
		handList = append(handList, hand)
	}

	return handList
}

type Hand struct {
	cards    []string
	handType int
	bid      int
}

// type HandType int

const (
	unknown int = iota
	HIGH
	ONEPAIR
	TWOPAIR
	THREE
	FULL
	FOUR
	FIVE
)

var jokerMap = [][]int{
	// Organized by [hand type][joker count]
	{0},                                // unknown
	{HIGH, ONEPAIR, THREE, FOUR, FIVE}, // high
	{ONEPAIR, THREE, FOUR, FIVE},       // one pair
	{TWOPAIR, FULL},                    // two pair
	{THREE, FOUR, FIVE},                // three of a kind
	{FULL},                             // full house
	{FOUR, FIVE},                       // four of a kind
}

func calculateHandType(hand []string, jokers bool) int {
	cardCountMap := map[string]int{}

	for _, card := range hand {
		cardCountMap[card]++
	}

	var three, pair bool
	handType := unknown

	for card, count := range cardCountMap {
		if jokers && card == "J" {
			continue
		}

		switch count {
		case 5:
			// return FIVE
			handType = FIVE
		case 4:
			// return FOUR
			handType = FOUR
		case 3:
			three = true
		case 2:
			if pair {
				// return TWOPAIR
				handType = TWOPAIR
			} else {
				pair = true
			}
		}
	}

	if handType == unknown {
		if three {
			if pair {
				// return FULL
				handType = FULL
			} else {
				// return THREE
				handType = THREE
			}
		} else if pair {
			// return ONEPAIR
			handType = ONEPAIR
		} else {
			handType = HIGH
		}
	}

	if jokers {
		jCount := cardCountMap["J"]

		if jCount == 5 {
			handType = FIVE
		} else {
			handType = jokerMap[handType][jCount]
		}
	}

	return handType
}

var cardStrengthPart1 = "23456789TJQKA"
var cardStrengthPart2 = "J23456789TQKA"

func (h Hand) compare(other Hand, cardStrength string) int {
	if h.handType != other.handType {
		return int(h.handType) - int(other.handType)
	}

	for i := 0; i < len(h.cards); i++ {
		h1Strength := strings.Index(cardStrength, h.cards[i])
		h2Strength := strings.Index(cardStrength, other.cards[i])

		if h1Strength != h2Strength {
			return h1Strength - h2Strength
		}
	}

	return 0
}

func orderHands(handList []Hand, cardStrength string) {
	n := len(handList)

	for swapped := true; swapped; {
		swapped = false

		for i := 0; i < n-1; i++ {
			hand1 := handList[i]
			hand2 := handList[i+1]
			value := hand1.compare(hand2, cardStrength)
			// fmt.Println(hand1, hand2, value)
			if value > 0 {
				handList[i+1], handList[i] = handList[i], handList[i+1]
				swapped = true
			}
		}

		n--
	}
}

func calculateWinnings(handList []Hand) int64 {
	var winnings int64

	for i, hand := range handList {
		winnings += int64(hand.bid * (i + 1))
	}

	return winnings
}
