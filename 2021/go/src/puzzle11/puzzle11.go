package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
)

const gridSize int = 10
const frameValue int = 999
const part1StepsNum int = 100

var dirs = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func flash(grid *[gridSize + 2][gridSize + 2]int, r int, c int, visited map[string]int, fNum int) int {
	key := strconv.Itoa(r) + "_" + strconv.Itoa(c)
	if _, ok := visited[key]; !ok {
		visited[key] = 1
		fNum += 1

		for _, d := range dirs {
			nextR, nextC := r+d[0], c+d[1]
			if nextR > 0 && nextR < gridSize+1 && nextC > 0 && nextC < gridSize+1 {
				grid[nextR][nextC]++
				if grid[nextR][nextC] > 9 && grid[nextR][nextC] != frameValue {
					fNum = flash(grid, nextR, nextC, visited, fNum)
				}
			}
		}
	}

	return fNum
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

	var grid [gridSize + 2][gridSize + 2]int

	for r := 0; r < gridSize+2; r++ {
		for c := 0; c < gridSize+2; c++ {
			if r == 0 || r == gridSize+1 || c == 0 || c == gridSize+1 {
				grid[r][c] = frameValue
			}
		}
	}

	for r, line := range lines {
		for c, ch := range line {
			d, _ := strconv.Atoi(string(ch))
			grid[r+1][c+1] = d
		}
	}

	part1, part2 := 0, 1

	for {
		fNum, total := 0, 0
		visited := make(map[string]int)

		for r := 1; r < gridSize+1; r++ {
			for c := 1; c < gridSize+1; c++ {
				grid[r][c]++
			}
		}

		for r := 1; r < gridSize+1; r++ {
			for c := 1; c < gridSize+1; c++ {
				if grid[r][c] > 9 {
					fNum = flash(&grid, r, c, visited, 0)
					if part2 <= part1StepsNum {
						part1 += fNum
					}
					total += fNum
				}
			}
		}

		for r := 1; r < gridSize+1; r++ {
			for c := 1; c < gridSize+1; c++ {
				if grid[r][c] > 9 {
					grid[r][c] = 0
				}
			}
		}

		if total == gridSize*gridSize {
			break
		}
		part2++
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
