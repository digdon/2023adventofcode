package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	obstacles := map[Point]bool{}
	maxRows, maxCols := len(inputLines), len(inputLines[0])
	var startPoint Point

	for row, line := range inputLines {
		for col := 0; col < maxCols; col++ {
			if line[col] == '#' {
				obstacles[Point{row, col}] = true
			} else if line[col] == 'S' {
				startPoint = Point{row, col}
			}
		}
	}

	// part1(obstacles, maxRows, maxCols, startPoint)
	visited := map[Point]int{}
	queue := []data{{startPoint, 0}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if _, ok := visited[item.point]; ok {
			continue
		}

		visited[item.point] = item.dist
		currPoint := item.point

		for i := 0; i < len(DIRECTIONS); i++ {
			nextRow, nextCol := currPoint.row+DIRECTIONS[i][0], currPoint.col+DIRECTIONS[i][1]
			nextPoint := Point{nextRow, nextCol}

			if nextRow < 0 || nextRow >= maxRows || nextCol < 0 || nextCol >= maxCols {
				continue
			} else if _, ok := visited[nextPoint]; !obstacles[nextPoint] && !ok {
				queue = append(queue, data{nextPoint, item.dist + 1})
			}
		}
	}

	part1Tiles := 0
	var evenCornerTiles, oddCornerTiles uint64 = 0, 0
	var evenBlockTiles, oddBlockTiles uint64 = 0, 0

	for _, value := range visited {
		if value%2 == 0 {
			evenBlockTiles++

			if value <= 64 {
				part1Tiles++
			} else {
				evenCornerTiles++
			}
		} else {
			oddBlockTiles++

			if value > 65 {
				oddCornerTiles++
			}
		}
	}

	fmt.Println("Part 1:", part1Tiles)

	// Part 2 foolishness
	var n uint64 = uint64((26501365 - maxCols/2) / maxCols)
	var even, odd uint64 = n * n, (n + 1) * (n + 1)
	var part2 uint64 = (odd * oddBlockTiles) + (even * evenBlockTiles) - (((n + 1) * oddCornerTiles) + (n * evenCornerTiles))
	fmt.Println("Part 2:", part2)
}

type data struct {
	point Point
	dist  int
}

func part1(obstacles map[Point]bool, maxRows, maxCols int, startPoint Point) {
	tiles := map[Point]bool{startPoint: true}

	for step := 0; step < 135; step++ {
		nextTiles := map[Point]bool{}

		for currPoint := range tiles {

			for i := 0; i < len(DIRECTIONS); i++ {
				nextRow, nextCol := currPoint.row+DIRECTIONS[i][0], currPoint.col+DIRECTIONS[i][1]
				nextPoint := Point{nextRow, nextCol}

				if nextRow < 0 || nextRow >= maxRows || nextCol < 0 || nextCol >= maxCols {
					continue
				} else if !obstacles[nextPoint] {
					nextTiles[nextPoint] = true
				}
			}
		}

		tiles = nextTiles

		fmt.Println(step+1, len(tiles))
	}

	fmt.Println("Part 1:", len(tiles))
}

var DIRECTIONS = [][]int{
	{0, 1},  // R
	{1, 0},  // D
	{0, -1}, // L
	{-1, 0}, // U
}

type Point struct {
	row, col int
}
