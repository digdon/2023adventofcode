package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	part1(inputLines)
	part2(inputLines)

}

var DIRECTIONS = [][]int64{
	{0, 1},  // R
	{1, 0},  // D
	{0, -1}, // L
	{-1, 0}, // U
}

func part2(inputLines []string) {
	var px, py int64
	var area, perimeter int64

	for _, line := range inputLines {
		hex := line[strings.Index(line, "(#")+2 : strings.LastIndex(line, ")")]
		count, _ := strconv.ParseInt(hex[:5], 16, 64)
		dirIndex, _ := strconv.Atoi(hex[5:])
		dir := DIRECTIONS[dirIndex]
		perimeter += count
		x, y := px+dir[1]*count, py+dir[0]*count
		// area += (py + y) * (px - x) // Trapezoidal formula
		area += (px * y) - (py * x) // Shoelace formula
		px, py = x, y
	}

	// Shoelace formula would then take the total area and divide by 2. Pick's theorem divides the perimeter by 2
	// Based on basic math, we can combine them
	area = (perimeter+area)/2 + 1 // Pick's theorem (more or less - we're solving for i + b)

	fmt.Println("Part 2:", area)
}
