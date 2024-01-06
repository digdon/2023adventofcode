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

	// We've calculated the area, via the shoelace formula (well, it still needs to be halved). We have the perimeter. From there, we can use Pick's theorem to work
	// out the number of interior points. Pick's theorem is A = i + b/2 - 1. Rearranging to solve for i, we end up with i = A - b/2 + 1
	// The area calculated by shoelace needs to be divided by 2 -> i = A/2 - b/2 + 1. That can be simplified to i = (A-b)/2 + 1
	i := (area-perimeter)/2 + 1

	// We already had the number of blocks in the perimeter (b) and now we have the internal blocks. Just add 'em up
	blocks := i + perimeter

	fmt.Println("Part 2:", blocks)
}
