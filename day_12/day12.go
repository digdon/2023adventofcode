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

	var part1Sum int64
	for _, line := range inputLines {
		sum := calc(line, 1)
		// fmt.Println(line, "->", sum)
		part1Sum += sum
	}
	fmt.Println("Part 1:", part1Sum)

	var part2Sum int64
	for _, line := range inputLines {
		sum := calc(line, 5)
		// fmt.Println(line, "->", sum)
		part2Sum += sum
	}
	fmt.Println("Part 2:", part2Sum)
}

type CacheKey struct {
	springPos, groupPos, contiguous int
}

func calc(input string, copies int) int64 {
	splitPos := strings.Index(input, " ")
	springDataString := input[:splitPos]
	groupDataString := input[splitPos+1:]

	if copies > 1 {
		springDataString += strings.Repeat("?"+springDataString, copies-1)
		groupDataString += strings.Repeat(","+groupDataString, copies-1)
	}

	springs := []byte(springDataString)
	groupStrings := strings.Split(groupDataString, ",")
	groups := make([]int, len(groupStrings))
	for i, str := range groupStrings {
		groups[i], _ = strconv.Atoi(str)
	}

	cache := map[CacheKey]int64{}
	return walkArrangements(cache, springs, groups, 0, 0, 0)
}

func walkArrangements(cache map[CacheKey]int64, springs []byte, groups []int, springPos, groupPos, contiguous int) int64 {
	cacheKey := CacheKey{springPos: springPos, groupPos: groupPos, contiguous: contiguous}
	value, ok := cache[cacheKey]

	if ok {
		// fmt.Printf("Cache hit for %v -> %d\n", cacheKey, value)
		return value
	}

	// fmt.Println(depth, springs, groups, springPos, groupPos, contiguous)
	if springPos == len(springs) {
		// We've reached the end
		// Have all of the groups been matched with no leftover counts?
		if groupPos == len(groups) && contiguous == 0 {
			return 1
		}

		// If we have leftover counts, do we have a single group left and do they match?
		if groupPos == len(groups)-1 && contiguous == groups[groupPos] {
			return 1
		}

		// Unmatched groups and/or leftover counts mean this arrangement failed
		return 0
	}

	var count int64

	char := springs[springPos]

	if char == '?' {
		// First act as a '.'
		if contiguous == 0 {
			// Nothing contiguous, so continue processing
			count += walkArrangements(cache, springs, groups, springPos+1, groupPos, contiguous)
		} else {
			// Do we have a group match?
			if groupPos < len(groups) && contiguous == groups[groupPos] {
				// A match - increment group and reset contiguous
				count += walkArrangements(cache, springs, groups, springPos+1, groupPos+1, 0)
			}
		}

		// Now act as a '#'
		count += walkArrangements(cache, springs, groups, springPos+1, groupPos, contiguous+1)
	} else if char == '.' {
		if contiguous == 0 {
			// Nothing contiguous, so continue processing
			count += walkArrangements(cache, springs, groups, springPos+1, groupPos, contiguous)
		} else {
			// Do we have a group match?
			if groupPos < len(groups) && contiguous == groups[groupPos] {
				// A match - increment group and reset contiguous
				count += walkArrangements(cache, springs, groups, springPos+1, groupPos+1, 0)
			}
		}
	} else { //if char == '#' {
		count += walkArrangements(cache, springs, groups, springPos+1, groupPos, contiguous+1)
	}

	cache[cacheKey] = count
	return count
}
