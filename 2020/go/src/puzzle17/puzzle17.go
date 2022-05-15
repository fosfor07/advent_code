package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type cubeType struct {
	x, y, z, w int
}

func cntActiveNeighbours1(c cubeType, grid [][][]rune) int {
	active := 0

	for k := c.z - 1; k <= c.z+1; k++ {
		for j := c.y - 1; j <= c.y+1; j++ {
			for i := c.x - 1; i <= c.x+1; i++ {
				if !(k == c.z && j == c.y && i == c.x) {
					if grid[k][j][i] == '#' {
						active++
					}
				}
			}
		}
	}

	return active
}

func cntActiveNeighbours2(c cubeType, grid [][][][]rune) int {
	active := 0

	for l := c.w - 1; l <= c.w+1; l++ {
		for k := c.z - 1; k <= c.z+1; k++ {
			for j := c.y - 1; j <= c.y+1; j++ {
				for i := c.x - 1; i <= c.x+1; i++ {
					if !(l == c.w && k == c.z && j == c.y && i == c.x) {
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

	wlen := (numCycles * 3) + 2
	zlen := wlen
	ylen := wlen + len(lines)
	xlen := wlen + len(lines[0])

	// part 1
	grid1 := make([][][]rune, zlen)
	uGrid1 := make([][][]rune, zlen)

	for z := 0; z < zlen; z++ {
		grid1[z] = make([][]rune, ylen)
		uGrid1[z] = make([][]rune, ylen)
		for y := 0; y < ylen; y++ {
			grid1[z][y] = make([]rune, xlen)
			uGrid1[z][y] = make([]rune, xlen)
			for x := 0; x < xlen; x++ {
				grid1[z][y][x] = '.'
				uGrid1[z][y][x] = '.'
			}
		}
	}

	// part 2
	grid2 := make([][][][]rune, wlen)
	uGrid2 := make([][][][]rune, wlen)

	for w := 0; w < wlen; w++ {
		grid2[w] = make([][][]rune, zlen)
		uGrid2[w] = make([][][]rune, zlen)
		for z := 0; z < zlen; z++ {
			grid2[w][z] = make([][]rune, ylen)
			uGrid2[w][z] = make([][]rune, ylen)
			for y := 0; y < ylen; y++ {
				grid2[w][z][y] = make([]rune, xlen)
				uGrid2[w][z][y] = make([]rune, xlen)
				for x := 0; x < xlen; x++ {
					grid2[w][z][y][x] = '.'
					uGrid2[w][z][y][x] = '.'
				}
			}
		}
	}

	w, z, y := wlen/2, zlen/2, ylen/2
	for _, line := range lines {
		x := xlen / 2
		for _, ch := range line {
			grid1[z][y][x] = ch
			uGrid1[z][y][x] = ch
			grid2[w][z][y][x] = ch
			uGrid2[w][z][y][x] = ch
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
						if w == 1 {
							c1 := cubeType{w: 1, z: z, y: y, x: x}
							active1 := cntActiveNeighbours1(c1, grid1)

							if grid1[z][y][x] == '#' {
								if active1 != 2 && active1 != 3 {
									uGrid1[z][y][x] = '.'
								}
							} else if grid1[z][y][x] == '.' {
								if active1 == 3 {
									uGrid1[z][y][x] = '#'
								}
							}
						}

						c2 := cubeType{w: w, z: z, y: y, x: x}
						active2 := cntActiveNeighbours2(c2, grid2)

						if grid2[w][z][y][x] == '#' {
							if active2 != 2 && active2 != 3 {
								uGrid2[w][z][y][x] = '.'
							}
						} else if grid2[w][z][y][x] == '.' {
							if active2 == 3 {
								uGrid2[w][z][y][x] = '#'
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
						if w == 1 {
							if cycle == numCycles-1 {
								if uGrid1[z][y][x] == '#' {
									part1++
								}
							} else {
								grid1[z][y][x] = uGrid1[z][y][x]
							}
						}

						if cycle == numCycles-1 {
							if uGrid2[w][z][y][x] == '#' {
								part2++
							}
						} else {
							grid2[w][z][y][x] = uGrid2[w][z][y][x]
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
