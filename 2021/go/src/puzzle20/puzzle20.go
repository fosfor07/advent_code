package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
)

const boardSize int = 400
const startPoint int = 150

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

	part1, part2 := 0, 0
	algorithm := ""
	var grid [boardSize][boardSize]rune

	for r := 0; r < boardSize; r++ {
		for c := 0; c < boardSize; c++ {
			grid[r][c] = '.'
		}
	}
	grid2 := grid

	for i, line := range lines {
		if line == "" {
			continue
		} else if i == 0 {
			for _, ch := range line {
				algorithm += string(ch)
			}
		} else {
			for j, ch := range line {
				grid[startPoint+i][startPoint+j] = ch
			}
		}
	}

	for s := 0; s < 50; s++ {
		for r := 1; r < boardSize-1; r++ {
			for c := 1; c < boardSize-1; c++ {
				idxStr := string(grid[r-1][c-1]) + string(grid[r-1][c]) + string(grid[r-1][c+1]) +
					string(grid[r][c-1]) + string(grid[r][c]) + string(grid[r][c+1]) +
					string(grid[r+1][c-1]) + string(grid[r+1][c]) + string(grid[r+1][c+1])

				idxBin := ""
				for _, ch := range idxStr {
					if ch == '#' {
						idxBin += "1"
					} else {
						idxBin += "0"
					}
				}

				idx, _ := strconv.ParseInt(idxBin, 2, 64)
				grid2[r][c] = rune(algorithm[idx])
			}
		}

		grid = grid2

		if s == 1 {
			for r := 50; r < boardSize-50; r++ {
				for c := 50; c < boardSize-50; c++ {
					if grid[r][c] == '#' {
						part1++
					}
				}
			}
		}
	}

	for r := 50; r < boardSize-50; r++ {
		for c := 50; c < boardSize-50; c++ {
			//fmt.Printf("%c", grid[r][c])
			if grid[r][c] == '#' {
				part2++
			}
		}
		//fmt.Printf("\n")
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
