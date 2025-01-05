package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type Direction string

const (
	DirectionUp    Direction = "up"
	DirectionDown  Direction = "down"
	DirectionRight Direction = "right"
	DirectionLeft  Direction = "left"
)

type positionType struct {
	r, c int
	dir  Direction
}

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

	startR, startC := 0, 0
	numRows := len(lines) + 2
	numCols := len(lines[0]) + 2

	baseGrid := make([][]rune, numRows)
	for r := 0; r < numRows; r++ {
		baseGrid[r] = make([]rune, numCols)
	}

	for r, line := range lines {
		for c, ch := range line {
			baseGrid[r+1][c+1] = ch

			if ch == '^' {
				startR = r + 1
				startC = c + 1
			}
		}
	}

	currentGrid := make([][]rune, len(baseGrid))
	for rowIdx, row := range baseGrid {
		currentGrid[rowIdx] = make([]rune, len(row))
		copy(currentGrid[rowIdx], baseGrid[rowIdx])
	}

	// part 1
	position := positionType{r: startR, c: startC, dir: DirectionUp}

	for {
		currentGrid[position.r][position.c] = 'X'
		position = guardMoves(currentGrid, position)

		if position.r == 0 || position.r == numRows-1 || position.c == 0 || position.c == numCols-1 {
			break
		}
	}

	for r := 1; r < numRows-1; r++ {
		for c := 1; c < numCols-1; c++ {
			if currentGrid[r][c] == 'X' {
				part1++
			}
		}
	}

	// part 2
	for r := 1; r < numRows-1; r++ {
		for c := 1; c < numCols-1; c++ {
			if (r == startR && c == startC) || currentGrid[r][c] == '#' {
				continue
			}

			for rowIdx := range baseGrid {
				copy(currentGrid[rowIdx], baseGrid[rowIdx])
			}
			position = positionType{r: startR, c: startC, dir: DirectionUp}

			currentGrid[r][c] = '#'
			positions := make(map[positionType]struct{})

			for {
				positions[position] = struct{}{}
				position = guardMoves(currentGrid, position)

				if position.r == 0 || position.r == numRows-1 || position.c == 0 || position.c == numCols-1 {
					break
				}
				if _, ok := positions[position]; ok {
					part2++
					break
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func guardMoves(grid [][]rune, p positionType) positionType {
	if p.dir == DirectionUp {
		if grid[p.r-1][p.c] == '#' {
			p.dir = DirectionRight
		} else {
			p.r--
		}
	} else if p.dir == DirectionDown {
		if grid[p.r+1][p.c] == '#' {
			p.dir = DirectionLeft
		} else {
			p.r++
		}
	} else if p.dir == DirectionRight {
		if grid[p.r][p.c+1] == '#' {
			p.dir = DirectionDown
		} else {
			p.c++
		}
	} else if p.dir == DirectionLeft {
		if grid[p.r][p.c-1] == '#' {
			p.dir = DirectionUp
		} else {
			p.c--
		}
	}

	return p
}
