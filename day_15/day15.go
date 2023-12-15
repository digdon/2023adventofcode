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

func generateHash(input string) int {
	value := 0

	for i := 0; i < len(input); i++ {
		char := int(input[i])
		value += char
		value *= 17
		value %= 256
	}

	return value
}

func part1(inputLines []string) {
	hashSum := 0

	for _, line := range inputLines {
		steps := strings.Split(line, ",")
		for _, step := range steps {
			hashSum += generateHash(step)
		}
	}

	fmt.Println("Part 1:", hashSum)
}

type LensSlot struct {
	label string
	value int
}

func part2(inputLines []string) {
	boxes := make([][]LensSlot, 256)

	// Process all of the steps
	for _, line := range inputLines {
		for _, step := range strings.Split(line, ",") {
			pos := strings.IndexAny(step, "=-")
			label, op := step[:pos], step[pos]
			targetBox := generateHash(label)

			if op == '-' {
				boxes[targetBox] = removeLens(boxes[targetBox], label)
			} else {
				lensValue, _ := strconv.Atoi(step[pos+1:])
				boxes[targetBox] = addLens(boxes[targetBox], label, lensValue)
			}
		}
	}

	// Calculate focusing power
	power := 0

	for i, box := range boxes {
		for j, slot := range box {
			lensPower := (i + 1) * (j + 1) * slot.value
			power += lensPower
		}
	}

	fmt.Println("Part 2:", power)
}

func removeLens(box []LensSlot, label string) []LensSlot {
	for pos := 0; pos < len(box); pos++ {
		if box[pos].label == label {
			return append(box[:pos], box[pos+1:]...)
		}
	}

	return box
}

func addLens(box []LensSlot, label string, lensValue int) []LensSlot {
	// Look for existing label
	for i := 0; i < len(box); i++ {
		if box[i].label == label {
			box[i].value = lensValue
			return box
		}
	}

	// Lable not already found, so a new entry
	return append(box, LensSlot{label, lensValue})
}
