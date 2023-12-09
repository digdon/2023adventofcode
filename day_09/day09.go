package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	part1Sum := 0
	part2Sum := 0

	for _, line := range inputLines {
		sequences := getHistorySequences(line)
		part1Sum += getNextValue(sequences)
		part2Sum += getPrevValue(sequences)
	}

	fmt.Println("Part 1:", part1Sum)
	fmt.Println("Part 2:", part2Sum)
}

func getNextValue(sequences [][]int) int {
	var value int
	for prev, i := 0, len(sequences)-1; i >= 0; i-- {
		value = sequences[i][len(sequences[i])-1] + prev
		prev = value
	}

	return value
}

func getPrevValue(sequences [][]int) int {
	var value int
	for prev, i := 0, len(sequences)-1; i >= 0; i-- {
		value = sequences[i][0] - prev
		prev = value
	}

	return value
}

func getHistorySequences(input string) [][]int {
	sequences := [][]int{getNumbers(input)}

	for allZero := false; !allZero; {
		currSeq := sequences[len(sequences)-1]
		nextSeq := []int{}

		allZero = true
		for i := 1; i < len(currSeq); i++ {
			diff := currSeq[i] - currSeq[i-1]
			nextSeq = append(nextSeq, diff)

			if diff != 0 {
				allZero = false
			}
		}

		sequences = append(sequences, nextSeq)
	}

	return sequences
}

var splitRE = regexp.MustCompile(`\s+`)

func getNumbers(input string) []int {
	values := []int{}
	split := splitRE.Split(input, -1)

	for _, valueStr := range split {
		value, _ := strconv.Atoi(valueStr)
		values = append(values, value)
	}

	return values
}
