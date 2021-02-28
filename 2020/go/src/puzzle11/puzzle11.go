package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

var dirs = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func initMaps(seatsMap, updatedMap [][]rune, xmax, ymax int, lines []string) {
	for y := 0; y <= ymax; y++ {
		seatsMap[y] = make([]rune, xmax+1)
		updatedMap[y] = make([]rune, xmax+1)
		for x := 0; x <= xmax; x++ {
			if x == 0 || x == xmax || y == 0 || y == ymax {
				seatsMap[y][x] = '-'
				updatedMap[y][x] = '-'
			} else {
				seatsMap[y][x] = rune(lines[y-1][x-1])
				updatedMap[y][x] = rune(lines[y-1][x-1])
			}
		}
	}
}

func cntAdjacentOccupied(x, y int, seatsMap [][]rune) int {
	occupied := 0

	for _, d := range dirs {
		nextY, nextX := y+d[0], x+d[1]
		if seatsMap[nextY][nextX] == '#' {
			occupied++
		}
	}

	return occupied
}

func cntVisibleOccupied(x, y int, seatsMap [][]rune) int {
	occupied := 0

	for _, d := range dirs {
		nextY, nextX := y, x

		for {
			nextY += d[0]
			nextX += d[1]
			if seatsMap[nextY][nextX] == '#' {
				occupied++
				break
			} else if seatsMap[nextY][nextX] != '.' {
				break
			}
		}
	}

	return occupied
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	lines, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	xmax := len(lines[0]) + 1
	ymax := len(lines) + 1

	seatsMap := make([][]rune, ymax+1)
	updatedMap := make([][]rune, ymax+1)

	// part 1
	initMaps(seatsMap, updatedMap, xmax, ymax, lines)

	part1 := 0
	for {
		part1 = 0

		for y := 1; y < ymax; y++ {
			for x := 1; x < xmax; x++ {
				occupied := cntAdjacentOccupied(x, y, seatsMap)

				if seatsMap[y][x] == 'L' {
					if occupied == 0 {
						updatedMap[y][x] = '#'
					}
				} else if seatsMap[y][x] == '#' {
					if occupied > 3 {
						updatedMap[y][x] = 'L'
					}
				}
			}
		}

		noChange := true
		for y, row := range updatedMap {
			for x := range row {
				if noChange {
					if updatedMap[y][x] == '#' {
						part1++
					}
					if seatsMap[y][x] != updatedMap[y][x] {
						noChange = false
					}
				}
				seatsMap[y][x] = updatedMap[y][x]
			}
		}

		if noChange {
			break
		}
	}

	// part 2
	initMaps(seatsMap, updatedMap, xmax, ymax, lines)

	part2 := 0
	for {
		part2 = 0

		for y := 1; y < ymax; y++ {
			for x := 1; x < xmax; x++ {
				occupied := cntVisibleOccupied(x, y, seatsMap)

				if seatsMap[y][x] == 'L' {
					if occupied == 0 {
						updatedMap[y][x] = '#'
					}
				} else if seatsMap[y][x] == '#' {
					if occupied > 4 {
						updatedMap[y][x] = 'L'
					}
				}
			}
		}

		noChange := true
		for y, row := range updatedMap {
			for x := range row {
				if noChange {
					if updatedMap[y][x] == '#' {
						part2++
					}
					if seatsMap[y][x] != updatedMap[y][x] {
						noChange = false
					}
				}
				seatsMap[y][x] = updatedMap[y][x]
			}
		}

		if noChange {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
