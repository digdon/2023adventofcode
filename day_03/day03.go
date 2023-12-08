package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
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

	partRE := regexp.MustCompile(`\d+`)
	possibleParts := []Part{}

	for lineNum, line := range inputLines {
		// Find all of the parts on this line
		partMatches := partRE.FindAllStringIndex(line, -1)

		for _, match := range partMatches {
			partNumStr := line[match[0]:match[1]]
			partNum, _ := strconv.Atoi(partNumStr)
			possibleParts = append(possibleParts, Part{id: partNum, x1: match[0], x2: match[1] - 1, y: lineNum})
		}
	}

	minX, minY, maxX, maxY := 0, 0, len(inputLines[0])-1, len(inputLines)-1

	actualParts := []int{}
	possibleGears := map[Point][]Part{}

	for _, part := range possibleParts {
		startX, endX := max(part.x1-1, minX), min(part.x2+1, maxX)
		startY, endY := max(part.y-1, minY), min(part.y+1, maxY)

		foundPart := false

		for y := startY; y <= endY; y++ {
			for x := startX; x <= endX; x++ {
				char := rune(inputLines[y][x])

				if char != '.' && !unicode.IsDigit(char) {
					foundPart = true

					if char == '*' {
						// Possible gear
						gearPoint := Point{x: x, y: y}
						possibleGears[gearPoint] = append(possibleGears[gearPoint], part)
					}
				}
			}
		}

		if foundPart {
			actualParts = append(actualParts, part.id)
		}
	}

	// fmt.Println(actualParts)
	// fmt.Println(possibleGears)

	part1Sum := 0

	for _, partNum := range actualParts {
		part1Sum += partNum
	}

	fmt.Println("Part 1:", part1Sum)

	part2Sum := 0

	for _, partList := range possibleGears {
		if len(partList) == 2 {
			ratio := partList[0].id * partList[1].id
			part2Sum += ratio
		}
	}

	fmt.Println("Part 2:", part2Sum)
}

type Part struct {
	id     int
	x1, x2 int
	y      int
}

type Point struct {
	x, y int
}
