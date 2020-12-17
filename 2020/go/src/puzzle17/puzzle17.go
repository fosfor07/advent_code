package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

type pType struct {
	x int
	y int
	z int
	w int
}

func cntActiveNeighbours(p pType, grid [][][][]rune) int {
	active := 0

	for l := p.w - 1; l <= p.w+1; l++ {
		for k := p.z - 1; k <= p.z+1; k++ {
			for j := p.y - 1; j <= p.y+1; j++ {
				for i := p.x - 1; i <= p.x+1; i++ {
					if !(l == p.w && k == p.z && j == p.y && i == p.x) {
						if grid[l][k][j][i] == '#' {
							active++
						}
					}
				}
			}
		}
	}

	return active
}

const numCycles int = 6

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

	wlen := (numCycles * 3) + 1
	zlen := wlen
	ylen := wlen + len(lines)
	xlen := wlen + len(lines[0])

	grid := make([][][][]rune, wlen)
	uGrid := make([][][][]rune, wlen)

	for w := 0; w < wlen; w++ {
		grid[w] = make([][][]rune, zlen)
		uGrid[w] = make([][][]rune, zlen)
		for z := 0; z < zlen; z++ {
			grid[w][z] = make([][]rune, ylen)
			uGrid[w][z] = make([][]rune, ylen)
			for y := 0; y < ylen; y++ {
				grid[w][z][y] = make([]rune, xlen)
				uGrid[w][z][y] = make([]rune, xlen)
				for x := 0; x < xlen; x++ {
					grid[w][z][y][x] = '.'
					uGrid[w][z][y][x] = '.'
				}
			}
		}
	}

	w, z, y := wlen/2, zlen/2, ylen/2
	for _, line := range lines {
		x := xlen / 2
		for _, ch := range line {
			grid[w][z][y][x] = ch
			uGrid[w][z][y][x] = ch
			x++
		}
		y++
	}

	cycle := 0
	for {
		for w := 1; w < wlen-1; w++ {
			for z := 1; z < zlen-1; z++ {
				for y := 1; y < ylen-1; y++ {
					for x := 1; x < xlen-1; x++ {
						var p pType
						p.w, p.z, p.y, p.x = w, z, y, x
						active := cntActiveNeighbours(p, grid)

						if grid[w][z][y][x] == '#' {
							if active != 2 && active != 3 {
								uGrid[w][z][y][x] = '.'
							}
						} else if grid[w][z][y][x] == '.' {
							if active == 3 {
								uGrid[w][z][y][x] = '#'
							}
						}
					}
				}
			}
		}

		for w := 0; w < wlen; w++ {
			for z := 0; z < zlen; z++ {
				for y := 0; y < ylen; y++ {
					for x := 0; x < xlen; x++ {
						if cycle == numCycles-1 {
							if uGrid[w][z][y][x] == '#' {
								part2++
							}
						} else {
							grid[w][z][y][x] = uGrid[w][z][y][x]
						}
					}
				}
			}
		}

		cycle++
		if cycle >= numCycles {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
