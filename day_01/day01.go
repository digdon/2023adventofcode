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
	// var runningTotal int
	// part1RE := regexp.MustCompile(`\d`)

	// for _, str := range inputLines {
	// 	digits := part1RE.FindAllString(str, -1)
	// 	numberString := digits[0] + digits[len(digits)-1]
	// 	number, _ := strconv.Atoi(numberString)
	// 	runningTotal += number
	// }

	// fmt.Println("Part 1:", runningTotal)

	digitStrings := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	part1Result := generateCalibration(inputLines, digitStrings)
	fmt.Println("Part 1:", part1Result)

	// Part 2
	wordDigitStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	part2Result := generateCalibration(inputLines, wordDigitStrings)
	fmt.Println("Part 2:", part2Result)
}

var numberWordToDigitMap = map[string]int{
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

func generateCalibration(input []string, numbers []string) int {
	var runningTotal int

	for _, str := range input {
		var values []string

		for pos := 0; pos < len(str); pos++ {
			temp := str[pos:]

			for _, numString := range numbers {
				if strings.HasPrefix(temp, numString) {
					values = append(values, numString)
					break
				}
			}
		}

		first := determineNumber(values[0])
		second := determineNumber(values[len(values)-1])
		number := (first * 10) + second
		runningTotal += number
	}

	return runningTotal
}

func determineNumber(str string) int {
	if value, err := strconv.Atoi(str); err == nil {
		return value
	}

	return numberWordToDigitMap[str]
}
