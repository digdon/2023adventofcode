package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var grid []string
var maxX, maxY int
var visited = map[Point]bool{}

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

	grid = inputLines
	maxX, maxY = len(grid[0])-1, len(grid)-1

	// Start by finding the startPoint S
	var startPoint Point
	for y := 0; y < len(inputLines); y++ {
		for x := 0; x < len(inputLines[y]); x++ {
			if inputLines[y][x] == 'S' {
				startPoint = Point{x: x, y: y}
			}
		}
	}

	fmt.Println(startPoint)

	// Using S, generate a loop
	pointNodeMap := map[Point]Node{}
	startNode := generateNode(startPoint, 0)
	queue := []Node{startNode}
	maxDistance := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		pointNodeMap[node.position] = node
		visited[node.position] = true

		if node.dist > maxDistance {
			maxDistance = node.dist
		}

		if len(node.neighbours) > 0 {
			for _, nPoint := range node.neighbours {
				_, ok := visited[nPoint]

				if !ok {
					nNode := generateNode(nPoint, node.dist+1)
					queue = append(queue, nNode)
				}
			}
		}
	}

	fmt.Println("Part 1:", maxDistance)

}

func findACorner(startNode Node, pointNodeMap map[Point]*Node) *Node {
	corner := &startNode

	for {
		if strings.IndexByte("F7LJ", corner.pipeType) != -1 {
			break
		} else {
			corner = pointNodeMap[corner.neighbours[0]]
		}
	}

	return corner
}

func generateNode(point Point, dist int) Node {
	neighbours := []Point{}

	// Look in each direction
	for i := 0; i < len(directions); i++ {
		dir := directions[i]
		x := point.x + dir.nextX
		y := point.y + dir.nextY

		if x < 0 || x > maxX || y < 0 || y > maxY {
			// out of bounds
			continue
		}

		pipe := grid[y][x]

		if !isPipeAllowed(pipe, dir.allowedPipes) {
			// Pipe is not an allowable type, so it's not connected
			continue
		}

		newPoint := Point{x: x, y: y}
		neighbours = append(neighbours, newPoint)
		//		visited[newPoint] = true
	}

	pipe := grid[point.y][point.x]
	return Node{pipeType: pipe, position: point, dist: dist, neighbours: neighbours}
}

func isPipeAllowed(char byte, allowed []byte) bool {
	if char == '.' {
		return false
	}

	if char == 'S' {
		return true
	}

	for _, val := range allowed {
		if val == char {
			return true
		}
	}

	return false
}

var directions = []Direction{
	{nextX: 0, nextY: -1, allowedPipes: []byte{'|', '7', 'F'}}, // up
	{nextX: 1, nextY: 0, allowedPipes: []byte{'-', 'J', '7'}},  // right
	{nextX: 0, nextY: 1, allowedPipes: []byte{'|', 'L', 'J'}},  // down
	{nextX: -1, nextY: 0, allowedPipes: []byte{'-', 'L', 'F'}}, // left
}

type Direction struct {
	nextX, nextY int
	allowedPipes []byte
}

type Point struct {
	x, y int
}

type Node struct {
	pipeType   byte
	position   Point
	dist       int
	neighbours []Point
	inside     Point
}
