package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	lines, err := readers.ReadLines(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	part1, part2 := 0, 0

	numRows := len(lines) + 6
	numCols := len(lines[0]) + 6

	grid := make([][]rune, numRows)
	for r := 0; r < numRows; r++ {
		grid[r] = make([]rune, numCols)
	}

	for r, line := range lines {
		for c, ch := range line {
			grid[r+3][c+3] = ch
		}
	}

	for r := 3; r < numRows-3; r++ {
		for c := 3; c < numCols-3; c++ {
			if grid[r][c] == 'X' {
				if grid[r][c+1] == 'M' && grid[r][c+2] == 'A' && grid[r][c+3] == 'S' {
					part1++
				}
				if grid[r][c-1] == 'M' && grid[r][c-2] == 'A' && grid[r][c-3] == 'S' {
					part1++
				}
				if grid[r+1][c] == 'M' && grid[r+2][c] == 'A' && grid[r+3][c] == 'S' {
					part1++
				}
				if grid[r-1][c] == 'M' && grid[r-2][c] == 'A' && grid[r-3][c] == 'S' {
					part1++
				}
				if grid[r+1][c+1] == 'M' && grid[r+2][c+2] == 'A' && grid[r+3][c+3] == 'S' {
					part1++
				}
				if grid[r+1][c-1] == 'M' && grid[r+2][c-2] == 'A' && grid[r+3][c-3] == 'S' {
					part1++
				}
				if grid[r-1][c-1] == 'M' && grid[r-2][c-2] == 'A' && grid[r-3][c-3] == 'S' {
					part1++
				}
				if grid[r-1][c+1] == 'M' && grid[r-2][c+2] == 'A' && grid[r-3][c+3] == 'S' {
					part1++
				}
			}

			if grid[r][c] == 'A' {
				if grid[r-1][c-1] == 'M' && grid[r-1][c+1] == 'M' && grid[r+1][c+1] == 'S' && grid[r+1][c-1] == 'S' ||
					grid[r-1][c-1] == 'S' && grid[r-1][c+1] == 'S' && grid[r+1][c+1] == 'M' && grid[r+1][c-1] == 'M' ||
					grid[r-1][c-1] == 'M' && grid[r-1][c+1] == 'S' && grid[r+1][c+1] == 'S' && grid[r+1][c-1] == 'M' ||
					grid[r-1][c-1] == 'S' && grid[r-1][c+1] == 'M' && grid[r+1][c+1] == 'M' && grid[r+1][c-1] == 'S' {
					part2++
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
