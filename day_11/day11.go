package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
	galaxies := fetchGalaxyPositions(inputLines, 1)
	part1Sum := calculateDistSum(galaxies)
	fmt.Println("Part 1:", part1Sum)

	// Part 2
	galaxies = fetchGalaxyPositions(inputLines, 1000000-1)
	part2Sum := calculateDistSum(galaxies)
	fmt.Println("Part 2:", part2Sum)
}

func calculateDistSum(galaxies []Point) int64 {
	var total int64 = 0

	for i := 0; i < len(galaxies); i++ {
		source := galaxies[i]

		for j := i + 1; j < len(galaxies); j++ {
			target := galaxies[j]
			dist := int64(math.Abs(float64(source.x-target.x)) + math.Abs(float64(source.y-target.y)))
			total += dist
		}
	}

	return total
}

func fetchGalaxyPositions(grid []string, expansion int) []Point {
	galaxies := []Point{}
	emptyRows := []int{}
	emptyColsMap := map[int]bool{}

	for i := 0; i < len(grid[0]); i++ {
		emptyColsMap[i] = true
	}

	for row := 0; row < len(grid); row++ {
		emptyRow := true
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '#' {
				galaxies = append(galaxies, Point{x: col, y: row})
				delete(emptyColsMap, col)
				emptyRow = false
			}
		}

		if emptyRow {
			emptyRows = append(emptyRows, row)
		}
	}

	emptyCols := make([]int, 0, len(emptyColsMap))
	for k := range emptyColsMap {
		emptyCols = append(emptyCols, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(emptyRows)))
	sort.Sort(sort.Reverse(sort.IntSlice(emptyCols)))

	// Expand empty rows
	for _, row := range emptyRows {
		for i := 0; i < len(galaxies); i++ {
			if galaxies[i].y > row {
				galaxies[i].y += expansion
			}
		}
	}

	// Expand empty columns
	for _, col := range emptyCols {
		for i := 0; i < len(galaxies); i++ {
			if galaxies[i].x > col {
				galaxies[i].x += expansion
			}
		}
	}

	return galaxies
}

type Point struct {
	x, y int
}
