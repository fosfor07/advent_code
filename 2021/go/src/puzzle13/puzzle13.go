package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type dotType struct {
	r, c int
}

type foldType struct {
	foldDir string
	foldVal int
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

	var dots []dotType
	var folds []foldType
	numRows, numCols := 0, 0

	for _, line := range lines {
		if line == "" {
			continue
		} else if line[0] == 'f' {
			tokens := strings.Fields(line)
			foldDef := strings.Split(tokens[2], "=")
			foldVal, _ := strconv.Atoi(foldDef[1])

			fold := foldType{foldDir: foldDef[0], foldVal: foldVal}
			folds = append(folds, fold)
		} else {
			tokens := strings.Split(line, ",")
			c, _ := strconv.Atoi(tokens[0])
			r, _ := strconv.Atoi(tokens[1])
			if c > numCols {
				numCols = c
			}
			if r > numRows {
				numRows = r
			}
			dot := dotType{r: r, c: c}
			dots = append(dots, dot)
		}
	}

	// Increase by 1 to get the number of rows and the number of columns
	// insted of the highest indices.
	numRows++
	numCols++

	// Make sure that numbers are odd.
	if numRows%2 == 0 {
		numRows++
	}
	if numCols%2 == 0 {
		numCols++
	}

	grid := make([][]rune, numRows)
	for r := 0; r < numRows; r++ {
		grid[r] = make([]rune, numCols)
	}

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			grid[r][c] = ' '
		}
	}

	for _, d := range dots {
		grid[d.r][d.c] = '#'
	}

	part1 := 0
	for i, f := range folds {
		if f.foldDir == "x" {
			for r := 0; r < numRows; r++ {
				for c := 0; c < numCols; c++ {
					if c > f.foldVal && grid[r][c] == '#' {
						fCol := f.foldVal - (c - f.foldVal)
						grid[r][fCol] = grid[r][c]
					}
				}
			}
			numCols = f.foldVal
		} else if f.foldDir == "y" {
			for r := 0; r < numRows; r++ {
				for c := 0; c < numCols; c++ {
					if r > f.foldVal && grid[r][c] == '#' {
						fRow := f.foldVal - (r - f.foldVal)
						grid[fRow][c] = grid[r][c]
					}
				}
			}
			numRows = f.foldVal
		}

		if i == 0 {
			for r := 0; r < numRows; r++ {
				for c := 0; c < numCols; c++ {
					if grid[r][c] == '#' {
						part1++
					}
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2:\n")
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			fmt.Printf("%c", grid[r][c])
		}
		fmt.Printf("\n")
	}
}
