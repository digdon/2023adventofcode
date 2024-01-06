package main

import (
	"bufio"
	"log"
	"os"
	"slices"
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

	broadcastDest := []string{}
	flipFlops := map[string]FlipFlop{}
	conjunctions := map[string]Conjunction{}

	for _, line := range inputLines {
		parts := strings.Split(line, " -> ")
		dests := strings.Split(parts[1], ", ")

		if parts[0] == "broadcaster" {
			broadcastDest = dests
		} else {
			name := parts[0][1:]
			if parts[0][0] == '%' {
				// flipflop
				flipFlops[name] = FlipFlop{false, dests}
			} else {
				// conjunction
				conjunctions[name] = Conjunction{map[string]bool{}, dests}
			}
		}
	}

	// Look for conjunction inputs
	for conjName, conjValue := range conjunctions {
		// for _, item := range broadcastDest {
		// 	if item == conjName {
		// 		conjValue.connectedPulses[item] = false
		// 	}
		// }

		for key, value := range flipFlops {
			if slices.Contains(value.destinations, conjName) {
				conjValue.connectedPulses[key] = false
			}
		}
	}
}

type FlipFlop struct {
	flag         bool
	destinations []string
}

type Conjunction struct {
	connectedPulses map[string]bool
	destinations    []string
}
