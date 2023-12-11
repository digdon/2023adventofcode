package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
)

func main() {

	grid, emptyColMap := loadGrid()
	grid = expandGrid(grid, emptyColMap)
	fmt.Println(grid)
	galaxies := findGalaxies(grid)
	fmt.Println(galaxies)

	part1Sum := 0

	for i := 0; i < len(galaxies); i++ {
		source := galaxies[i]

		for j := i + 1; j < len(galaxies); j++ {
			target := galaxies[j]
			dist := int(math.Abs(float64(source.x-target.x)) + math.Abs(float64(source.y-target.y)))
			part1Sum += dist
		}
	}

	fmt.Println("Part 1:", part1Sum)
}

func loadGrid() ([]string, map[int]bool) {
	var inputLines []string
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()

	if !ok {
		fmt.Println("No text to read - we're done")
		return nil, nil
	}

	galaxyRE := regexp.MustCompile(`#`)
	line := scanner.Text()

	emptyColMap := map[int]bool{}
	for i := 0; i < len(line); i++ {
		emptyColMap[i] = true
	}

	for {
		inputLines = append(inputLines, line)
		matches := galaxyRE.FindAllStringIndex(line, -1)

		if len(matches) > 0 {
			for _, galaxy := range matches {
				delete(emptyColMap, galaxy[0])
			}
		} else {
			// Empty line - duplicate it
			inputLines = append(inputLines, line)
		}

		ok = scanner.Scan()

		if !ok {
			break
		}

		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return inputLines, emptyColMap
}

func expandGrid(grid []string, emptyColMap map[int]bool) []string {
	if len(emptyColMap) > 0 {
		// empty columns, so we need to expand those
		cols := []int{}
		for k := range emptyColMap {
			cols = append(cols, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(cols)))
		fmt.Println(cols)

		for _, colVal := range cols {
			for i := 0; i < len(grid); i++ {
				grid[i] = grid[i][:colVal] + "." + grid[i][colVal:]
			}
		}
	}

	return grid
}

func findGalaxies(grid []string) []Point {
	galaxies := []Point{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' {
				galaxies = append(galaxies, Point{x: x, y: y})
			}
		}
	}

	return galaxies
}

type Point struct {
	x, y int
}
