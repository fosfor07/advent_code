package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"sort"
	"strconv"
	"strings"
)

type pointType struct {
	r, c, h int
}

func basin(grid [][]int, r int, c int, visited map[string]int, bSize int) int {
	if grid[r][c-1] < 9 {
		key := strconv.Itoa(r) + "_" + strconv.Itoa(c-1)
		if _, ok := visited[key]; !ok {
			visited[key] = 1
			bSize += 1
			bSize = basin(grid, r, c-1, visited, bSize)
		}
	}
	if grid[r][c+1] < 9 {
		key := strconv.Itoa(r) + "_" + strconv.Itoa(c+1)
		if _, ok := visited[key]; !ok {
			visited[key] = 1
			bSize += 1
			bSize = basin(grid, r, c+1, visited, bSize)
		}
	}
	if grid[r-1][c] < 9 {
		key := strconv.Itoa(r-1) + "_" + strconv.Itoa(c)
		if _, ok := visited[key]; !ok {
			visited[key] = 1
			bSize += 1
			bSize = basin(grid, r-1, c, visited, bSize)
		}
	}
	if grid[r+1][c] < 9 {
		key := strconv.Itoa(r+1) + "_" + strconv.Itoa(c)
		if _, ok := visited[key]; !ok {
			visited[key] = 1
			bSize += 1
			bSize = basin(grid, r+1, c, visited, bSize)
		}
	}

	return bSize
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

	part1, part2 := 0, 0

	numRows := len(lines) + 2
	numCols := len(lines[0]) + 2

	grid := make([][]int, numRows)
	for r := 0; r < numRows; r++ {
		grid[r] = make([]int, numCols)
	}

	for r, line := range lines {
		for c, ch := range line {
			chInt, _ := strconv.Atoi(string(ch))
			grid[r+1][c+1] = chInt
		}
	}
	for r := 0; r < numRows; r++ {
		grid[r][0] = 9
		grid[r][numCols-1] = 9
	}
	for c := 0; c < numCols; c++ {
		grid[0][c] = 9
		grid[numRows-1][c] = 9
	}

	var lowPoints []pointType

	for r := 1; r < numRows-1; r++ {
		for c := 1; c < numCols-1; c++ {
			if grid[r][c] < grid[r][c-1] && grid[r][c] < grid[r][c+1] &&
				grid[r][c] < grid[r-1][c] && grid[r][c] < grid[r+1][c] {
				p := pointType{r: r, c: c, h: grid[r][c]}
				lowPoints = append(lowPoints, p)
			}
		}
	}

	visited := make(map[string]int)
	var basins []int

	for _, p := range lowPoints {
		part1 += p.h + 1

		key := strconv.Itoa(p.r) + "_" + strconv.Itoa(p.c)
		visited[key] = 1
		bSize := basin(grid, p.r, p.c, visited, 1)

		basins = append(basins, bSize)
	}

	var max3 []int
	for i := 0; i < 3; i++ {
		max3 = append(max3, basins[i])
	}
	sort.Ints(max3)

	for i := 3; i < len(basins); i++ {
		if max3[0] < basins[i] {
			max3[0] = basins[i]
			sort.Ints(max3)
		}
	}
	part2 = max3[0] * max3[1] * max3[2]

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
