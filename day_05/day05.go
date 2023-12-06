package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mapping struct {
	source int64
	dest   int64
	count  int64
}

type Range struct {
	start int64
	end   int64
}

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

	seedSoilList := []Mapping{}
	soilFertilizerList := []Mapping{}
	fertilizerWaterList := []Mapping{}
	waterLightList := []Mapping{}
	lightTempList := []Mapping{}
	tempHumidityList := []Mapping{}
	humidityLocationList := []Mapping{}

	// mode := nothing
	var currentList *[]Mapping
	dataRE := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)
	seedsRE := regexp.MustCompile(`\d+`)
	var seedsList []string

	for _, line := range inputLines {
		if len(line) == 0 {
			continue
		} else if strings.HasPrefix(line, "seeds:") {
			seedsList = seedsRE.FindAllString(line, -1)
		} else if strings.HasPrefix(line, "seed-to-soil") {
			currentList = &seedSoilList
		} else if strings.HasPrefix(line, "soil-to-fertilizer") {
			currentList = &soilFertilizerList
		} else if strings.HasPrefix(line, "fertilizer-to-water") {
			currentList = &fertilizerWaterList
		} else if strings.HasPrefix(line, "water-to-light") {
			currentList = &waterLightList
		} else if strings.HasPrefix(line, "light-to-temperature") {
			currentList = &lightTempList
		} else if strings.HasPrefix(line, "temperature-to-humidity") {
			currentList = &tempHumidityList
		} else if strings.HasPrefix(line, "humidity-to-location") {
			currentList = &humidityLocationList
		} else {
			matches := dataRE.FindStringSubmatch(line)
			destStr, sourceStr, countStr := matches[1], matches[2], matches[3]
			source, _ := strconv.ParseInt(sourceStr, 10, 64)
			dest, _ := strconv.ParseInt(destStr, 10, 64)
			count, _ := strconv.ParseInt(countStr, 10, 64)
			(*currentList) = append((*currentList), Mapping{source: source, dest: dest, count: count})
		}
	}

	// Part 1
	var nearestLoc int64 = math.MaxInt64

	for _, seedStr := range seedsList {
		seed, _ := strconv.ParseInt(seedStr, 10, 64)
		soil := getValue(seedSoilList, seed)
		fert := getValue(soilFertilizerList, soil)
		water := getValue(fertilizerWaterList, fert)
		light := getValue(waterLightList, water)
		temp := getValue(lightTempList, light)
		humidity := getValue(tempHumidityList, temp)
		loc := getValue(humidityLocationList, humidity)

		if loc < nearestLoc {
			nearestLoc = loc
		}
	}

	fmt.Println("Part 1:", nearestLoc)

	// Part 2
	seeds := []Range{}
	for i := 0; i < len(seedsList); i += 2 {
		start, _ := strconv.ParseInt(seedsList[i], 10, 64)
		count, _ := strconv.ParseInt(seedsList[i+1], 10, 64)
		end := start + count - 1
		seeds = append(seeds, Range{start, end})
	}

	seeds = convert(seeds, seedSoilList)
	seeds = convert(seeds, soilFertilizerList)
	seeds = convert(seeds, fertilizerWaterList)
	seeds = convert(seeds, waterLightList)
	seeds = convert(seeds, lightTempList)
	seeds = convert(seeds, tempHumidityList)
	seeds = convert(seeds, humidityLocationList)

	lowest := int64(math.MaxInt64)

	for _, item := range seeds {
		if item.start < lowest {
			lowest = item.start
		}
	}

	fmt.Println("Part 2:", lowest)
}

func convert(source []Range, dest []Mapping) []Range {
	newRanges := []Range{}
	remaining := append([]Range(nil), source...)

	for len(remaining) > 0 {
		rangeItem := remaining[0]
		remaining = remaining[1:]
		sourceStart := rangeItem.start
		sourceEnd := rangeItem.end
		intersectionFound := false

		for _, mapItem := range dest {
			diff := mapItem.dest - mapItem.source
			destStart := mapItem.source
			destEnd := destStart + mapItem.count - 1

			// Work out intersection
			if sourceEnd < destStart || sourceStart > destEnd {
				// No intersection
			} else {
				interStart := max(sourceStart, destStart)
				interEnd := min(sourceEnd, destEnd)
				newRanges = append(newRanges, Range{interStart + diff, interEnd + diff})

				if sourceStart < destStart {
					// Leftovers at the start of the range
					remaining = append(remaining, Range{sourceStart, destStart - 1})
				}

				if sourceEnd > destEnd {
					// Leftovers at the end of the range
					remaining = append(remaining, Range{destEnd + 1, sourceEnd})
				}

				intersectionFound = true
				break
			}
		}

		if !intersectionFound {
			newRanges = append(newRanges, rangeItem)
		}
	}

	return newRanges
}

func min(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int64) int64 {
	if x > y {
		return x
	} else {
		return y
	}
}

func getValue(list []Mapping, value int64) int64 {
	foundValue := value

	for _, item := range list {
		start := item.source
		end := start + item.count - 1

		if value >= start && value <= end {
			foundValue = item.dest + (value - start)
			break
		}
	}

	return foundValue
}
