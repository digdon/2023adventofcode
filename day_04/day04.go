package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
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

	var part1Total int
	cardMatchMap := make(map[int]int)
	part2Queue := []int{}

	for i := 0; i < len(inputLines); i++ {
		cardNum := i + 1
		count := matchCount(inputLines[i])
		part1Total += int(math.Exp2(float64(count - 1))) // one match = 1, two matches = 2, three = 4, four = 8, etc -> 2^(n-1)
		part2Queue = append(part2Queue, cardNum)
		cardMatchMap[cardNum] = count
	}

	fmt.Println("Part 1:", part1Total)

	cardCountMap := make(map[int]int)

	for len(part2Queue) > 0 {
		card := part2Queue[0]            // Grab first item from queue
		part2Queue = part2Queue[1:]      // Remove first item from queue
		cardCountMap[card]++             // update the number of times we've seen this card (based on the card number)
		matchCount := cardMatchMap[card] // get the match count for this card

		// Based on the number of matches, append the next n cards to the queue (these are copies we've won)
		for i := 1; i <= matchCount; i++ {
			part2Queue = append(part2Queue, card+i)
		}
	}

	//	fmt.Println(cardCountMap)

	var part2Total int

	for _, value := range cardCountMap {
		part2Total += value
	}

	fmt.Println("Part 2:", part2Total)
}

var cardRE = regexp.MustCompile(`.*:((\s*\d+)+)\s*\|((\s*\d+)+)$`)
var sepRE = regexp.MustCompile(`\s+`)

func matchCount(line string) int {
	matches := cardRE.FindStringSubmatch(line)
	theirCards := sepRE.Split(strings.TrimSpace(matches[1]), -1)
	myCards := sepRE.Split(strings.TrimSpace(matches[3]), -1)
	// fmt.Println(len(theirCards), theirCards)
	// fmt.Println(len(myCards), myCards)

	unionMap := make(map[string]bool)
	for _, item := range theirCards {
		unionMap[item] = true
	}

	var count int
	for _, item := range myCards {
		if unionMap[item] {
			count++
		}
	}

	return count
}
