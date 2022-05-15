package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type pointType struct {
	r, c, h int
}

func findBasinSize(grid [][]int, r int, c int, visited map[string]int, bSize int) int {
	if grid[r][c-1] < 9 {
		key := strconv.Itoa(r) + "_" + strconv.Itoa(c-1)
		if _, ok := visited[key]; !ok {
			visited[key] = 1
			bSize += 1
			bSize = findBasinSize(grid, r, c-1, visited, bSize)
		}
	}
	if grid[r][c+1] < 9 {
		key := strconv.Itoa(r) + "_" + strconv.Itoa(c+1)
		if _, ok := visited[key]; !ok {
			visited[key] = 1
			bSize += 1
			bSize = findBasinSize(grid, r, c+1, visited, bSize)
		}
	}
	if grid[r-1][c] < 9 {
		key := strconv.Itoa(r-1) + "_" + strconv.Itoa(c)
		if _, ok := visited[key]; !ok {
			visited[key] = 1
			bSize += 1
			bSize = findBasinSize(grid, r-1, c, visited, bSize)
		}
	}
	if grid[r+1][c] < 9 {
		key := strconv.Itoa(r+1) + "_" + strconv.Itoa(c)
		if _, ok := visited[key]; !ok {
			visited[key] = 1
			bSize += 1
			bSize = findBasinSize(grid, r+1, c, visited, bSize)
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

	var basins []int
	visited := make(map[string]int)

	for _, p := range lowPoints {
		part1 += p.h + 1

		key := strconv.Itoa(p.r) + "_" + strconv.Itoa(p.c)
		visited[key] = 1
		bSize := findBasinSize(grid, p.r, p.c, visited, 1)

		basins = append(basins, bSize)
	}

	sort.Ints(basins)
	part2 = basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
