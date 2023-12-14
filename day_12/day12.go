package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	var part1Sum int

	for _, line := range inputLines {
		// value := calculateArrangementsNew(line)
		value := calc(line, 1)
		fmt.Println(line, "->", value)
		part1Sum += value
	}

	fmt.Println("Part 1:", part1Sum)

	// inputLine := "???.### 1,1,3"
	// inputLine := ".??..??...?##. 1,1,3"
	// inputLine := "?###???????? 3,2,1"
	// inputLine := "??.?###????????? 2,4,4"
	// fmt.Println(calc(inputLine, 1))
}

func calc(input string, copies int) int {
	splitPos := strings.Index(input, " ")
	springDataString := input[:splitPos]
	groupDataString := input[splitPos+1:]

	if copies > 1 {
		tempSprintData := make([]string, copies)
		tempGroupData := make([]string, copies)

		for i := 0; i < copies; i++ {
			tempSprintData[i] = springDataString
			tempGroupData[i] = groupDataString
		}

		springDataString = strings.Join(tempSprintData, "?")
		groupDataString = strings.Join(tempGroupData, ",")
	}

	// fmt.Println(springDataString)
	// fmt.Println(groupDataString)
	// return 0
	springData := []byte(springDataString)

	// Work out the spring groupings
	groupSepRE := regexp.MustCompile(`,`)
	matches := groupSepRE.Split(groupDataString, -1)
	groupList := []int{}
	for _, match := range matches {
		value, _ := strconv.Atoi(match)
		groupList = append(groupList, value)
	}

	return walkArrangements(springData, groupList, 0, 0)
}

func walkArrangements(springData []byte, groupList []int, pos int, contiguous int) int {
	// fmt.Println(string(springData))

	for ; pos < len(springData); pos++ {
		char := springData[pos]

		if char == '?' {
			tempSpringData := make([]byte, len(springData))
			copy(tempSpringData, springData)
			tempSpringData[pos] = '.'
			count := walkArrangements(tempSpringData, groupList, pos, contiguous)
			tempSpringData[pos] = '#'
			count += walkArrangements(tempSpringData, groupList, pos, contiguous)
			return count
		} else if char == '#' {
			if len(groupList) == 0 {
				// fmt.Println("\tfailed because we ran out of groups")
				return 0
			}

			contiguous++
			if contiguous > groupList[0] {
				// fmt.Printf("\tfailed because cont %d > %v\n", contiguous, groupList)
				return 0
			}
		} else if char == '.' {
			if contiguous > 0 {
				if contiguous != groupList[0] {
					// fmt.Printf("\tfailed because cont %d != %v\n", contiguous, groupList)
					return 0
				}

				groupList = groupList[1:]
				contiguous = 0
			}
		}
	}

	if springData[pos-1] == '#' {
		if len(groupList) != 1 {
			// fmt.Println("\t failed because not all groups matched")
			return 0
		} else if contiguous == groupList[0] {
			// fmt.Println("Success!", string(springData))
			return 1
		} else {
			// fmt.Printf("\tfailed because cont %d > %v\n", contiguous, groupList)
			return 0
		}
	} else {
		if len(groupList) == 0 {
			// fmt.Println("Success!", string(springData))
			return 1
		} else {
			// fmt.Println("\t failed because not all groups matched")
			return 0
		}
	}
}

func calculateArrangementsNew(inputLine string) int {
	var validCount int

	pos := strings.Index(inputLine, " ")
	springData := inputLine[:pos]
	groupData := inputLine[pos+1:]

	// Work out the spring groupings
	groupSepRE := regexp.MustCompile(`,`)
	matches := groupSepRE.Split(groupData, -1)
	groupList := []int{}
	for _, match := range matches {
		value, _ := strconv.Atoi(match)
		groupList = append(groupList, value)
	}

	// fmt.Println(springData)
	// fmt.Println(groupList)

	// Work out the number of unknown springs
	unknownRE := regexp.MustCompile(`\?`)
	unknownCount := len(unknownRE.FindAllStringIndex(springData, -1))

	for arr := 0; arr < (1 << unknownCount); arr++ {
		// fmt.Println("Arr:", arr)
		elements := []byte{}
		contiguous := 0
		groupPos := 0
		arrangementGood := true

		for springNum, i := 0, 0; i < len(springData); i++ {
			char := springData[i]

			if char == '?' {
				springValue := 1 << springNum

				if springValue&arr == springValue {
					char = '#'
				} else {
					char = '.'
				}

				springNum++
			}

			if char == '#' {
				if groupPos >= len(groupList) {
					arrangementGood = false
					break
				}

				contiguous++

				if contiguous > groupList[groupPos] {
					arrangementGood = false
				}
			} else {
				if i > 0 && elements[i-1] == '#' {
					// End of a possible contiguous group - count to be sure
					if contiguous != groupList[groupPos] {
						arrangementGood = false
					} else {
						groupPos++
					}
				}

				contiguous = 0
			}

			elements = append(elements, char)

			if !arrangementGood {
				break
			}
		}

		// fmt.Println(string(elements))

		if len(elements) < len(springData) {
			// Not a full string, so it must be invalid
			continue
		}

		if elements[len(elements)-1] == '#' {
			if contiguous != groupList[groupPos] {
				continue
			} else {
				groupPos++
			}
		}

		if groupPos < len(groupList) {
			// Missed some of the groups, so it must be invalid
			continue
		}

		if arrangementGood {
			// fmt.Println("Good:", string(elements))
			validCount++
		}
	}

	return validCount
}
