package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	numRE := regexp.MustCompile(`\s+`)
	timeList := numRE.Split(inputLines[0], -1)[1:]
	distList := numRE.Split(inputLines[1], -1)[1:]

	// Part 1
	score := 1

	for i := 0; i < len(timeList); i++ {
		time, _ := strconv.Atoi(timeList[i])
		dist, _ := strconv.Atoi(distList[i])
		winCount := calculateWinsBS(time, dist)
		// fmt.Printf("time: %d, dist: %d -> wins %d\n", time, dist, winCount)
		score *= winCount
	}

	fmt.Println("Part 1:", score)

	// Part 2
	timeLimit, _ := strconv.Atoi(strings.Join(timeList, ""))
	dist, _ := strconv.Atoi(strings.Join(distList, ""))
	winCount := calculateWinsBS(timeLimit, dist)
	fmt.Println("Part 2:", winCount)
}

// func calculateWins(time, distance int) int {
// 	winCount := 0

// 	for i := 0; i <= time; i++ {
// 		if i*(time-i) > distance {
// 			winCount++
// 		}
// 	}

// 	return winCount
// }

// func calculateWins(time, distance int) int {
// 	wins := 0

// 	for i := 0; i <= time; i++ {
// 		if i*(time-i) > distance {
// 			wins = time - ((i - 1) * 2) - 1
// 			break
// 		}
// 	}

// 	return wins
// }

func calculateWinsBS(time, distance int) int {
	low, high := 0, (time/2)+1
	var mid int
	found := false

	for {
		mid = low + ((high - low) / 2)

		if mid == low || mid == high {
			break
		}

		value := mid * (time - mid)

		if value <= distance {
			low = mid
		} else if value > distance {
			high = mid
			found = true
		}
	}

	if found {
		return (time - (mid * 2) - 1)
	} else {
		return 0
	}
}
