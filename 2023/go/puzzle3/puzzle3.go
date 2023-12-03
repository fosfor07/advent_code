package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/fosfor07/advent_code/pkg/readers"
	"github.com/fosfor07/advent_code/pkg/utils"
)

type pointType struct {
	r, c int
}

type gearType struct {
	cnt, n1, n2 int
}

func isAdjacentToSymbol(grid [][]rune, r int, c int) bool {
	if (grid[r-1][c-1] != '.' && !unicode.IsDigit(grid[r-1][c-1])) ||
		(grid[r-1][c] != '.' && !unicode.IsDigit(grid[r-1][c])) ||
		(grid[r-1][c+1] != '.' && !unicode.IsDigit(grid[r-1][c+1])) ||
		(grid[r][c-1] != '.' && !unicode.IsDigit(grid[r][c-1])) ||
		(grid[r][c+1] != '.' && !unicode.IsDigit(grid[r][c+1])) ||
		(grid[r+1][c-1] != '.' && !unicode.IsDigit(grid[r+1][c-1])) ||
		(grid[r+1][c] != '.' && !unicode.IsDigit(grid[r+1][c])) ||
		(grid[r+1][c+1] != '.' && !unicode.IsDigit(grid[r+1][c+1])) {
		return true
	}
	return false
}

func getGearLocations(grid [][]rune, r int, c int) []pointType {
	gears := []pointType{}
	location := pointType{}

	if grid[r-1][c-1] == '*' {
		location.r = r - 1
		location.c = c - 1
		gears = append(gears, location)
	}
	if grid[r-1][c] == '*' {
		location.r = r - 1
		location.c = c
		gears = append(gears, location)
	}
	if grid[r-1][c+1] == '*' {
		location.r = r - 1
		location.c = c + 1
		gears = append(gears, location)
	}

	if grid[r][c-1] == '*' {
		location.r = r
		location.c = c - 1
		gears = append(gears, location)
	}
	if grid[r][c+1] == '*' {
		location.r = r
		location.c = c + 1
		gears = append(gears, location)
	}

	if grid[r+1][c-1] == '*' {
		location.r = r + 1
		location.c = c - 1
		gears = append(gears, location)
	}
	if grid[r+1][c] == '*' {
		location.r = r + 1
		location.c = c
		gears = append(gears, location)
	}
	if grid[r+1][c+1] == '*' {
		location.r = r + 1
		location.c = c + 1
		gears = append(gears, location)
	}

	return gears
}

func updateGears(locations []pointType, gears map[string]gearType, partNum int) {
	for _, loc := range locations {
		key := strconv.Itoa(loc.r) + "_" + strconv.Itoa(loc.c)
		if _, ok := gears[key]; ok {
			if gears[key].cnt == 1 {
				gear := gearType{cnt: 2, n1: gears[key].n1, n2: partNum}
				gears[key] = gear
			}
		} else {
			gear := gearType{cnt: 1, n1: partNum}
			gears[key] = gear
		}
	}
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
	gears := make(map[string]gearType)

	numRows := len(lines) + 2
	numCols := len(lines[0]) + 2

	grid := make([][]rune, numRows)
	for r := 0; r < numRows; r++ {
		grid[r] = make([]rune, numCols)
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			grid[r][c] = '.'
		}
	}

	for r, line := range lines {
		for c, ch := range line {
			grid[r+1][c+1] = ch
		}
	}

	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			if unicode.IsDigit(grid[r][c]) {
				if isAdjacentToSymbol(grid, r, c) {
					digBefore, digAfter := "", ""

					currentCol := c
					for {
						currentCol -= 1
						if unicode.IsDigit(grid[r][currentCol]) {
							digBefore += string(grid[r][currentCol])
						} else {
							break
						}
					}

					currentCol = c
					for {
						currentCol += 1
						if unicode.IsDigit(grid[r][currentCol]) {
							digAfter += string(grid[r][currentCol])
						} else {
							break
						}
					}

					digBeforeRev := utils.ReverseString(digBefore)
					partNumStr := digBeforeRev + string(grid[r][c]) + digAfter
					partNum, _ := strconv.Atoi(partNumStr)
					part1 += partNum

					locations := getGearLocations(grid, r, c)
					if len(locations) > 0 {
						updateGears(locations, gears, partNum)
					}

					// skip the following digits
					c = currentCol
				}
			}
		}
	}

	for _, g := range gears {
		if g.cnt == 2 {
			part2 += g.n1 * g.n2
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
