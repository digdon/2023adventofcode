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

	mirrors := map[Point]byte{}

	for row, line := range inputLines {
		for col := 0; col < len(line); col++ {
			if char := line[col]; char != '.' {
				// Found a mirror - record its location
				mirrors[Point{row, col}] = char
			}
		}
	}

	// fmt.Println(mirrors)

	maxRows, maxCols := len(inputLines), len(inputLines[0])

	// Part 1
	part1Value := calcEnergizedTiles(Beam{point: Point{0, 0}, dir: RIGHT}, mirrors, maxRows, maxCols)
	fmt.Println("Part 1:", part1Value)

	// Part 2
	maxTiles := 0
	for row := 0; row < len(inputLines); row++ {
		start := Beam{Point{row, 0}, RIGHT}
		tiles := calcEnergizedTiles(start, mirrors, maxRows, maxCols)

		if tiles > maxTiles {
			maxTiles = tiles
		}

		start = Beam{Point{row, maxCols - 1}, LEFT}
		tiles = calcEnergizedTiles(start, mirrors, maxRows, maxCols)

		if tiles > maxTiles {
			maxTiles = tiles
		}
	}

	for col := 0; col < maxCols; col++ {
		start := Beam{Point{0, col}, DOWN}
		tiles := calcEnergizedTiles(start, mirrors, maxRows, maxCols)

		if tiles > maxTiles {
			maxTiles = tiles
		}

		start = Beam{Point{maxRows - 1, col}, UP}
		tiles = calcEnergizedTiles(start, mirrors, maxRows, maxCols)

		if tiles > maxTiles {
			maxTiles = tiles
		}
	}

	fmt.Println("Part 2:", maxTiles)
}

func calcEnergizedTiles(start Beam, mirrors map[Point]byte, maxRow, maxCol int) int {
	energized := map[Point]int{}
	queue := []Beam{start}
	cache := map[Beam]bool{}

	for len(queue) > 0 {
		beam := queue[0]
		queue = queue[1:]

		for beam.point.row >= 0 && beam.point.row < maxRow && beam.point.col >= 0 && beam.point.col < maxCol {
			if cache[beam] {
				// We've been at this point, in this direction already - break the beam loop
				// Direction is the key point here, because we could hit this same spot from a different direction, which is still valid
				// fmt.Println(beam, "already seen")
				break
			}

			// fmt.Println(beam.point)

			cache[beam] = true // record that we've hit this spot before, in this direction
			energized[beam.point]++

			// Is the beam at a mirror or splitter?
			if obstacle, found := mirrors[beam.point]; found {
				if mirror, ok := mirrorDirectionMap[obstacle]; ok {
					// A mirror
					beam.dir = mirror[beam.dir]
				} else {
					// A splitter
					if (beam.dir == LEFT || beam.dir == RIGHT) && obstacle == '|' {
						beam.dir = UP                                                    // Change direction for original beam
						newBeam := Beam{Point{beam.point.row + 1, beam.point.col}, DOWN} // create and queue up second (split) beam
						queue = append(queue, newBeam)
					} else if (beam.dir == UP || beam.dir == DOWN) && obstacle == '-' {
						beam.dir = LEFT // Change direction for original beam
						newBeam := Beam{Point{beam.point.row, beam.point.col + 1}, RIGHT}
						queue = append(queue, newBeam)
					}
				}
			}

			// Move the beam
			beam.point.row += DIRECTIONS[beam.dir][0]
			beam.point.col += DIRECTIONS[beam.dir][1]
		}
	}

	return len(energized)
}

type direction int

const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

var DIRECTIONS = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

type Beam struct {
	point Point
	dir   direction
}

type Point struct {
	row, col int
}

var mirrorDirectionMap = map[byte]([]direction){
	'/':  {RIGHT, LEFT, DOWN, UP},
	'\\': {LEFT, RIGHT, UP, DOWN},
}
