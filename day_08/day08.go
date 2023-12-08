package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	dirList := inputLines[0]
	nodeMap := map[string]Node{}
	mapRE := regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)

	for i := 2; i < len(inputLines); i++ {
		matches := mapRE.FindStringSubmatch(inputLines[i])
		source, left, right := matches[1], matches[2], matches[3]
		nodeMap[source] = Node{left: left, right: right}
	}

	// Part 1
	fmt.Println("Steps:", part1(dirList, nodeMap))

	// Part 2
	fmt.Println("Part 2:", part2(dirList, nodeMap))
}

func part1(dirList string, nodeMap map[string]Node) int {
	var steps int
	currNodeID := "AAA"
	dirPos := 0

	for currNodeID != "ZZZ" {
		node := nodeMap[currNodeID]
		lr := dirList[dirPos%len(dirList)]

		if lr == 'L' {
			currNodeID = node.left
		} else {
			currNodeID = node.right
		}

		steps++
		dirPos++
	}

	return steps
}

func part2(dirList string, nodeMap map[string]Node) int {
	// Start by finding all of the xxA nodes and putting them in the node list
	currIDList := []string{}

	for nodeID := range nodeMap {
		if nodeID[2] == 'A' {
			currIDList = append(currIDList, nodeID)
		}
	}

	// Calculate the number of steps each starting node would take
	stepsList := []int{}

	for _, id := range currIDList {
		steps, dirPos := 0, 0
		currID := id
		for currID[2] != 'Z' {
			node := nodeMap[currID]
			lr := dirList[dirPos%len(dirList)]

			if lr == 'L' {
				currID = node.left
			} else {
				currID = node.right
			}

			steps++
			dirPos++
		}

		stepsList = append(stepsList, steps)
	}

	// Now for the interesting part. The solution is the GCD of all of the step counts for each start node
	// prime with the first step count
	steps := stepsList[0]

	// Now calculate the GCD for all of the items
	for i := 1; i < len(stepsList); i++ {
		steps = lcm(steps, stepsList[i])
	}

	return steps
}

func lcm(x, y int) int {
	if y == 0 {
		return 0
	}

	return (x * y) / gcd(x, y)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

type Node struct {
	left  string
	right string
}
