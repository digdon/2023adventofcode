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
	source uint64
	dest   uint64
	count  uint64
}

type Range struct {
	start uint64
	end   uint64
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
			source, _ := strconv.ParseUint(sourceStr, 10, 64)
			dest, _ := strconv.ParseUint(destStr, 10, 64)
			count, _ := strconv.ParseUint(countStr, 10, 64)
			(*currentList) = append((*currentList), Mapping{source: source, dest: dest, count: count})
		}
	}

	// Part 1
	var nearestLoc uint64 = math.MaxUint64
	var nearestSeed uint64

	for _, seedStr := range seedsList {
		seed, _ := strconv.ParseUint(seedStr, 10, 64)
		soil := getValue(seedSoilList, seed)
		fert := getValue(soilFertilizerList, soil)
		water := getValue(fertilizerWaterList, fert)
		light := getValue(waterLightList, water)
		temp := getValue(lightTempList, light)
		humidity := getValue(tempHumidityList, temp)
		loc := getValue(humidityLocationList, humidity)

		if loc < nearestLoc {
			nearestLoc = loc
			nearestSeed = seed
		}
	}

	fmt.Printf("Part 1: nearest location %d, nearest seed %d\n", nearestLoc, nearestSeed)

	// Part 2
	nearestLoc, nearestSeed = math.MaxUint64, 0

	for i := 0; i < len(seedsList); i += 2 {
		start, _ := strconv.ParseUint(seedsList[i], 10, 64)
		count, _ := strconv.ParseUint(seedsList[i+1], 10, 64)
		end := start + count

		for seed := start; seed <= end; seed++ {
			soil := getValue(seedSoilList, seed)
			fert := getValue(soilFertilizerList, soil)
			water := getValue(fertilizerWaterList, fert)
			light := getValue(waterLightList, water)
			temp := getValue(lightTempList, light)
			humidity := getValue(tempHumidityList, temp)
			loc := getValue(humidityLocationList, humidity)

			if loc < nearestLoc {
				nearestLoc = loc
				nearestSeed = seed
			}
		}
	}

	fmt.Printf("Part 2: nearest location %d, nearest seed %d\n", nearestLoc, nearestSeed)
}

func getValue(list []Mapping, value uint64) uint64 {
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
