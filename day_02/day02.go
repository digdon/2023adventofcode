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

	gameNumRE := regexp.MustCompile(`Game (\d+)`)

	var gameSum, gamePower int

	for _, line := range inputLines {
		matches := gameNumRE.FindStringSubmatch(line)
		gameNum, _ := strconv.Atoi(matches[1])
		sub := line[strings.Index(line, ":")+1:]
		valid, power := isGameValid(sub)
		fmt.Printf("Game %d: %t, %d\n", gameNum, valid, power)
		gamePower += power
		if valid {
			gameSum += gameNum
		}
	}

	fmt.Println("Part 1:", gameSum)
	fmt.Println("Part 2:", gamePower)
}

var setsRE = regexp.MustCompile(`((\d+\s+[a-z]+)(,\s*)?)+`)
var cubesRE = regexp.MustCompile(`(\d+) ([a-z]+)`)
var colourAvailableMap = map[string]int{"red": 12, "green": 13, "blue": 14}

func isGameValid(line string) (bool, int) {
	colourMaxMap := map[string]int{"red": 0, "green": 0, "blue": 0}
	valid := true

	setMatches := setsRE.FindAllString(line, -1)

	for _, set := range setMatches {
		cubeMatches := cubesRE.FindAllStringSubmatch(set, -1)

		for _, item := range cubeMatches {
			colour := item[2]
			count, _ := strconv.Atoi(item[1])

			if count > colourAvailableMap[colour] {
				valid = false
			}

			if count > colourMaxMap[colour] {
				colourMaxMap[colour] = count
			}
		}
	}

	power := 1
	for _, value := range colourMaxMap {
		power *= value
	}

	return valid, power
}
