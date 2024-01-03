package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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
	grid := inputLines
	maxX, maxY := len(grid[0])-1, len(grid)-1

	// Start by finding the start S
	var start Point
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if grid[y][x] == 'S' {
				start = Point{x: x, y: y}
				break
			}
		}
	}

	// Determine the start direction
	dirList := findStartDirNeighbours(grid, start.x, start.y)
	dir := dirList[0]

	// Walk the loop and record pieces
	pipePath := []Point{start}
	pipeMap := map[Point]bool{start: true}

	for x, y := start.x+directions[dir].nextX, start.y+directions[dir].nextY; !(x == start.x && y == start.y); x, y = x+directions[dir].nextX, y+directions[dir].nextY {
		pipePath = append(pipePath, Point{x, y})
		pipeMap[Point{x, y}] = true
		ch := grid[y][x]

		switch ch {
		case 'F':
			if dir == LEFT {
				dir = DOWN
			} else {
				dir = RIGHT
			}

		case '7':
			if dir == RIGHT {
				dir = DOWN
			} else {
				dir = LEFT
			}

		case 'L':
			if dir == DOWN {
				dir = RIGHT
			} else {
				dir = UP
			}

		case 'J':
			if dir == DOWN {
				dir = LEFT
			} else {
				dir = UP
			}
		}
	}

	fmt.Println("Part 1:", (len(pipePath)+1)/2)

	// Part 2 stuff - shoelace and Pick's theorem
	pipePath = append(pipePath, start)
	area := 0
	for i := 1; i < len(pipePath); i++ {
		px, py := pipePath[i-1].x, pipePath[i-1].y
		x, y := pipePath[i].x, pipePath[i].y
		area += (px * y) - (py * x)
	}

	value := (int(math.Abs(float64(area)))-(len(pipePath)-1))/2 + 1
	fmt.Println("Part 2 (Pick/shoelace):", value)

	// Part 2 - raycasting
	// First we need to replace 'S' with the correct pipe piece, based on what's attached to it
	var startChar byte
	if slices.Contains(dirList, UP) {
		if slices.Contains(dirList, RIGHT) {
			startChar = 'L'
		} else if slices.Contains(dirList, LEFT) {
			startChar = 'J'
		} else if slices.Contains(dirList, DOWN) {
			startChar = '|'
		}
	} else if slices.Contains(dirList, DOWN) {
		if slices.Contains(dirList, RIGHT) {
			startChar = 'F'
		} else if slices.Contains(dirList, LEFT) {
			startChar = '7'
		}
	} else if slices.Contains(dirList, LEFT) && slices.Contains(dirList, RIGHT) {
		startChar = '-'
	} else {
		fmt.Println("Some kind of error?")
		os.Exit(1)
	}

	grid[start.y] = fmt.Sprintf("%s%c%s", grid[start.y][:start.x], startChar, grid[start.y][start.x+1:])

	var tempPoint Point
	count := 0
	for y := 0; y < maxY; y++ {
		in := false
		var prevCorner byte
		for x := 0; x < maxX; x++ {
			tempPoint.x = x
			tempPoint.y = y

			if pipeMap[tempPoint] {
				ch := grid[y][x]

				// if ch == '|' || ch == 'L' || ch == 'J' {
				// 	in = !in
				// }
				if ch == '|' {
					in = !in
				} else if ch == 'L' || ch == 'F' {
					prevCorner = ch
				} else if ch == 'J' {
					if prevCorner == 'F' {
						in = !in
					}
					prevCorner = 0
				} else if ch == '7' {
					if prevCorner == 'L' {
						in = !in
					}
					prevCorner = 0
				}
			} else if in {
				count++
			}
		}
	}

	fmt.Println("Part 2 (raycast):", count)
}

func findStartDirNeighbours(grid []string, x, y int) []direction {
	dirList := []direction{}

	for dir := dirStart + 1; dir < dirEnd; dir++ {
		nextX, nextY := x+directions[dir].nextX, y+directions[dir].nextY

		if nextX < 0 || nextX > len(grid[0])-1 || nextY < 0 || nextY > len(grid)-1 {
			continue
		}

		ch := grid[nextY][nextX]
		if slices.Contains(directions[dir].allowedPipes, ch) {
			dirList = append(dirList, dir)
		}
	}

	return dirList
}

type direction int

const (
	dirStart direction = iota
	RIGHT
	DOWN
	LEFT
	UP
	dirEnd
)

var directions = []Direction{
	{},
	{nextX: 1, nextY: 0, allowedPipes: []byte{'-', 'J', '7'}},  // right
	{nextX: 0, nextY: 1, allowedPipes: []byte{'|', 'L', 'J'}},  // down
	{nextX: -1, nextY: 0, allowedPipes: []byte{'-', 'L', 'F'}}, // left
	{nextX: 0, nextY: -1, allowedPipes: []byte{'|', '7', 'F'}}, // up
}

type Direction struct {
	nextX, nextY int
	allowedPipes []byte
}

type Point struct {
	x, y int
}
