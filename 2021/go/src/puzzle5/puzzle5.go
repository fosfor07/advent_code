package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
	"utils"
)

const gridSize int = 1000

func addHorizontalOrVertical(grid2 *[gridSize][gridSize]int, x1 int, x2 int, y1 int, y2 int) {
	if y1 == y2 {
		if x2 > x1 {
			for x := x1; x <= x2; x++ {
				grid2[y1][x]++
			}
		} else {
			for x := x2; x <= x1; x++ {
				grid2[y1][x]++
			}
		}
	} else if x1 == x2 {
		if y2 > y1 {
			for y := y1; y <= y2; y++ {
				grid2[y][x1]++
			}
		} else {
			for y := y2; y <= y1; y++ {
				grid2[y][x1]++
			}
		}
	}
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

	var grid1 [gridSize][gridSize]int
	var grid2 [gridSize][gridSize]int

	for _, line := range lines {
		points := strings.Split(line, "->")
		start := strings.TrimSpace(points[0])
		end := strings.TrimSpace(points[1])

		startArr := strings.Split(start, ",")
		endArr := strings.Split(end, ",")

		x1, _ := strconv.Atoi(startArr[0])
		y1, _ := strconv.Atoi(startArr[1])
		x2, _ := strconv.Atoi(endArr[0])
		y2, _ := strconv.Atoi(endArr[1])

		// part 1
		addHorizontalOrVertical(&grid1, x1, x2, y1, y2)

		// part 2
		// check if it is a diagonal
		if utils.Abs(x1-x2) == utils.Abs(y1-y2) {
			if x2 > x1 && y2 > y1 {
				y := y1
				for x := x1; x <= x2; x++ {
					grid2[y][x]++
					y++
				}
			} else if x1 > x2 && y1 > y2 {
				y := y1
				for x := x1; x >= x2; x-- {
					grid2[y][x]++
					y--
				}
			} else if x2 > x1 && y1 > y2 {
				y := y1
				for x := x1; x <= x2; x++ {
					grid2[y][x]++
					y--
				}
			} else if x1 > x2 && y2 > y1 {
				y := y1
				for x := x1; x >= x2; x-- {
					grid2[y][x]++
					y++
				}
			}
		} else {
			addHorizontalOrVertical(&grid2, x1, x2, y1, y2)
		}
	}

	part1, part2 := 0, 0

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if grid1[x][y] >= 2 {
				part1++
			}
			if grid2[x][y] >= 2 {
				part2++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
