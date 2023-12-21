package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func part1(inputLines []string) {
	dataRegex := regexp.MustCompile(`([UDLR])\s+(\d+)\s+\(#([A-Fa-f0-9]{6})\)`)
	steps := []step{}

	// Build list of 'commands' and determine the dimensions of the 'grid'
	for _, line := range inputLines {
		matches := dataRegex.FindStringSubmatch(line)
		dir := DIRECTIONS[DIRECTIONS_INDEX[matches[1]]]
		count, _ := strconv.ParseInt(matches[2], 10, 32)
		steps = append(steps, step{dir, count})
	}

	var row, col int64
	var rowMin, colMin int64 = math.MaxInt, math.MaxInt
	var rowMax, colMax int64 = math.MinInt, math.MinInt

	// Work out the dimensions of the grid
	for _, step := range steps {
		row += step.dir[0] * step.count
		col += step.dir[1] * step.count

		if row > rowMax {
			rowMax = row
		}

		if row < rowMin {
			rowMin = row
		}

		if col > colMax {
			colMax = col
		}

		if col < colMin {
			colMin = col
		}
	}

	// Build the grid
	blockCount := 0
	currRow, currCol := max(0, -(rowMin)), max(0, -(colMin))
	maxRows, maxCols := (rowMax-rowMin)+1, (colMax-colMin)+1
	fmt.Println(currRow, ",", currCol)
	grid := make([][]byte, maxRows)
	for i := range grid {
		grid[i] = make([]byte, maxCols)
	}

	for _, step := range steps {
		for i := int64(0); i < step.count; i++ {
			grid[currRow][currCol] = '#'
			blockCount++
			currRow += step.dir[0]
			currCol += step.dir[1]
		}
	}

	fmt.Println(blockCount)

	// Find a starting point for the flood fill
	var floodRow, floodCol int64 = 0, 0
	for ; floodCol < maxCols; floodCol++ {
		if grid[floodRow][floodCol] == '#' {
			break
		}
	}
	floodRow++
	floodCol++

	fmt.Println(floodRow, floodCol)
	floodCount := floodFill(grid, floodRow, floodCol)

	// for i := 0; i < maxRows; i++ {
	// 	for j := 0; j < maxCols; j++ {
	// 		char := grid[i][j]

	// 		if char == 0 {
	// 			char = '.'
	// 		}
	// 		fmt.Printf("%c", char)
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println("Part 1:", floodCount+blockCount)
}

var DIRECTIONS_INDEX = map[string]int{
	"R": 0,
	"D": 1,
	"L": 2,
	"U": 3,
}

type step struct {
	dir   []int64
	count int64
}

func floodFill(grid [][]byte, startRow, startCol int64) int {
	floodCount := 0

	rowQueue := []int64{startRow}
	colQueue := []int64{startCol}

	for len(rowQueue) > 0 {
		row := rowQueue[0]
		rowQueue = rowQueue[1:]
		col := colQueue[0]
		colQueue = colQueue[1:]

		if grid[row][col] == '#' {
			// fmt.Printf("Already seen %d,%d\n", row, col)
			continue
		}

		grid[row][col] = '#'
		floodCount++

		for _, dir := range DIRECTIONS {
			nextRow, nextCol := row+dir[0], col+dir[1]

			if grid[nextRow][nextCol] != '#' {
				rowQueue = append(rowQueue, nextRow)
				colQueue = append(colQueue, nextCol)
			}
		}
	}

	return floodCount
}
