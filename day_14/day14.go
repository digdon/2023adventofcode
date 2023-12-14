package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"sort"
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

	allRocks := map[Point]bool{}
	for row := 0; row < len(inputLines); row++ {
		for col := 0; col < len(inputLines[row]); col++ {
			char := inputLines[row][col]

			if char != '.' {
				allRocks[Point{row, col}] = char == 'O'
			}
		}
	}

	maxRows := len(inputLines)
	maxCols := len(inputLines[0])

	// Part 1 stuff
	tiltNorth(allRocks)
	part1Load := calculateLoad(allRocks, maxRows)
	fmt.Println("Part 1:", part1Load)

	// Finish the cycle in preparation for part 2
	tiltWest(allRocks)
	tiltSouth(allRocks, maxRows)
	tiltEast(allRocks, maxCols)

	// Part 2 stuff
	cache := map[string]int{}
	target := 1000000000
	cycle := 0

	key := generateHashKey(allRocks)
	cache[key] = cycle

	for cycle = 1; cycle < target; cycle++ {
		tiltNorth(allRocks)
		tiltWest(allRocks)
		tiltSouth(allRocks, maxRows)
		tiltEast(allRocks, maxCols)

		key = generateHashKey(allRocks)
		val, ok := cache[key]

		if ok {
			cycleLength := cycle - val
			reps := (target - cycle) / cycleLength
			cycle += reps * cycleLength
		} else {
			cache[key] = cycle
		}
	}

	part2Load := calculateLoad(allRocks, maxRows)
	fmt.Println("Part 2:", part2Load)

	// fmt.Println("\nModified grid")

	// for row := 0; row < maxRows; row++ {
	// 	for col := 0; col < maxRows; col++ {
	// 		char := '.'
	// 		val, ok := allRocks[Point{row, col}]

	// 		if ok {
	// 			if val {
	// 				char = 'O'
	// 			} else {
	// 				char = '#'
	// 			}
	// 		}

	// 		fmt.Printf("%c", char)
	// 	}

	// 	fmt.Println()
	// }
}

type Point struct {
	row, col int
}

func (p Point) String() string {
	return fmt.Sprintf("{%d,%d}", p.row, p.col)
}

func calculateLoad(allRocks map[Point]bool, maxRows int) int {
	var load int

	for key, value := range allRocks {
		if value {
			load += maxRows - key.row
		}
	}

	return load
}

func generateHashKey(allRocks map[Point]bool) string {
	rocks := getRocksFromMap(allRocks)
	sort.Slice(rocks, func(i, j int) bool {
		if rocks[i].row < rocks[j].row {
			return true
		} else if rocks[i].row == rocks[j].row {
			return rocks[i].col < rocks[j].col
		} else {
			return false
		}
	})

	hash := sha256.New()

	for _, rock := range rocks {
		hash.Write([]byte(rock.String()))
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func tiltNorth(allRocks map[Point]bool) {
	rocks := getRocksFromMap(allRocks)
	sort.Slice(rocks, func(i, j int) bool { return rocks[i].row < rocks[j].row })

	for _, rock := range rocks {
		if allRocks[rock] {
			// Moveable rock, so try to move it
			newRow := rock.row

			for newRow > 0 {
				tempPoint := Point{newRow - 1, rock.col}
				_, ok := allRocks[tempPoint]

				if ok {
					// Rock in the way - done
					break
				} else {
					newRow--
				}
			}

			if newRow != rock.row {
				// Move the rock
				delete(allRocks, rock)
				allRocks[Point{newRow, rock.col}] = true
			}
		}
	}
}

func tiltWest(allRocks map[Point]bool) {
	rocks := getRocksFromMap(allRocks)
	sort.Slice(rocks, func(i, j int) bool { return rocks[i].col < rocks[j].col })

	for _, rock := range rocks {
		if allRocks[rock] {
			// Moveable rock, so try to move it
			newCol := rock.col

			for newCol > 0 {
				tempPoint := Point{rock.row, newCol - 1}
				_, ok := allRocks[tempPoint]

				if ok {
					// Rock in the way - done
					break
				} else {
					newCol--
				}
			}

			if newCol != rock.col {
				// Move the rock
				delete(allRocks, rock)
				allRocks[Point{rock.row, newCol}] = true
			}
		}
	}
}

func tiltSouth(allRocks map[Point]bool, maxRows int) {
	rocks := getRocksFromMap(allRocks)
	sort.Slice(rocks, func(i, j int) bool { return rocks[i].row > rocks[j].row })

	for _, rock := range rocks {
		if allRocks[rock] {
			// Moveable rock, so try to move it
			newRow := rock.row

			for newRow < maxRows-1 {
				tempPoint := Point{newRow + 1, rock.col}
				_, ok := allRocks[tempPoint]

				if ok {
					// Rock in the way - done
					break
				} else {
					newRow++
				}
			}

			if newRow != rock.row {
				// Move the rock
				delete(allRocks, rock)
				allRocks[Point{newRow, rock.col}] = true
			}
		}
	}
}

func tiltEast(allRocks map[Point]bool, maxCols int) {
	rocks := getRocksFromMap(allRocks)
	sort.Slice(rocks, func(i, j int) bool { return rocks[i].col > rocks[j].col })

	for _, rock := range rocks {
		if allRocks[rock] {
			// Moveable rock, so try to move it
			newCol := rock.col

			for newCol < maxCols-1 {
				tempPoint := Point{rock.row, newCol + 1}
				_, ok := allRocks[tempPoint]

				if ok {
					// Rock in the way - done
					break
				} else {
					newCol++
				}
			}

			if newCol != rock.col {
				// Move the rock
				delete(allRocks, rock)
				allRocks[Point{rock.row, newCol}] = true
			}
		}
	}
}

func getRocksFromMap(rockMap map[Point]bool) []Point {
	rocks := []Point{}
	for key := range rockMap {
		rocks = append(rocks, key)
	}

	return rocks
}
