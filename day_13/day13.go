package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	patterns := [][]string{}
	currPattern := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			patterns = append(patterns, currPattern)
			currPattern = []string{}
		} else {
			currPattern = append(currPattern, line)
		}
	}

	patterns = append(patterns, currPattern)

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var part1Sum int
	var part2Sum int

	for _, pattern := range patterns {
		part1Sum += calcVerticalReflection(pattern, 0) + calcHorizontalReflection(pattern, 0)
		part2Sum += calcVerticalReflection(pattern, 1) + calcHorizontalReflection(pattern, 1)
	}

	fmt.Println("Part 1:", part1Sum)
	fmt.Println("Part 2:", part2Sum)
}

func calcVerticalReflection(pattern []string, targetDiff int) int {
	for col := 0; col < len(pattern[0])-1; col++ {
		// Start by looking at adjacent columns
		diffCount := columnDiff(pattern, col, col+1)

		// Now start trying to match columns to the left and right, until we reach an edge or exceed the target diff
		for prevCol, nextCol := col-1, col+2; diffCount <= targetDiff && prevCol >= 0 && nextCol < len(pattern[0]); prevCol, nextCol = prevCol-1, nextCol+1 {
			diffCount += columnDiff(pattern, prevCol, nextCol)
		}

		if diffCount == targetDiff {
			return col + 1
		}
	}

	return 0
}

func columnDiff(pattern []string, col1, col2 int) int {
	var diffCount int

	for row := 0; row < len(pattern); row++ {
		if pattern[row][col1] != pattern[row][col2] {
			diffCount++
		}
	}

	return diffCount
}

func calcHorizontalReflection(pattern []string, targetDiff int) int {
	for row := 0; row < len(pattern)-1; row++ {
		// Start by looking at adjacent rows
		diffCount := rowDiff(pattern, row, row+1)

		// Now start trying to match rows above and below, until we reach an edge or exceed the target diff
		for prevRow, nextRow := row-1, row+2; diffCount <= targetDiff && prevRow >= 0 && nextRow < len(pattern); prevRow, nextRow = prevRow-1, nextRow+1 {
			diffCount += rowDiff(pattern, prevRow, nextRow)
		}

		if diffCount == targetDiff {
			return (row + 1) * 100
		}
	}

	return 0
}

func rowDiff(pattern []string, row1, row2 int) int {
	var diffCount int

	for col := 0; col < len(pattern[row1]); col++ {
		if pattern[row1][col] != pattern[row2][col] {
			diffCount++
		}
	}

	return diffCount
}
