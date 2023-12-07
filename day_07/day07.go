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

	handList := []Hand{}

	for _, line := range inputLines {
		pos := strings.Index(line, " ")
		cards, bidStr := line[:pos], line[pos+1:]
		bid, _ := strconv.Atoi(bidStr)
		handTypePart1 := calculateHandType(cards)
		hand := Hand{
			cards:         strings.Split(cards, ""),
			handTypePart1: handTypePart1,
			bid:           bid,
		}
		handList = append(handList, hand)
	}

	// Part 1
	orderHands(handList, cardStrengthPart1)
	fmt.Println("Part 1:", calculateWinnings(handList))

	// Part 2
	// orderHands(handList, cardStrengthPart2)
}

type Hand struct {
	cards         []string
	handTypePart1 HandType
	handTypePart2 HandType
	bid           int
}

type HandType int

const (
	unknown HandType = iota
	HIGH
	ONEPAIR
	TWOPAIR
	THREE
	FULL
	FOUR
	FIVE
)

func calculateHandType(hand string) HandType {
	cardCountMap := map[rune]int{}

	for _, card := range hand {
		cardCountMap[card]++
	}

	var three, pair bool

	for _, count := range cardCountMap {
		switch count {
		case 5:
			return FIVE
		case 4:
			return FOUR
		case 3:
			three = true
		case 2:
			if pair {
				return TWOPAIR
			} else {
				pair = true
			}
		}
	}

	if three {
		if pair {
			return FULL
		} else {
			return THREE
		}
	} else if pair {
		return ONEPAIR
	}

	return HIGH
}

var cardStrengthPart1 = "23456789TJQKA"
var cardStrengthPart2 = "J23456789TQKA"

func (h Hand) compare(other Hand, cardStrength string) int {
	if h.handTypePart1 != other.handTypePart1 {
		return int(h.handTypePart1) - int(other.handTypePart1)
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
			fmt.Println(hand1, hand2, value)
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
