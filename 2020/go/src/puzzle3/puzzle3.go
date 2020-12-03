package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

type slopeType struct {
	x   int
	y   int
	res int
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

	xlen := len(lines[0])
	ylen := len(lines)
	xmax := xlen - 1
	ymax := ylen - 1

	treeMap := make([][]rune, len(lines))

	for y, line := range lines {
		treeMap[y] = make([]rune, xlen)
		for x, ch := range line {
			treeMap[y][x] = ch
		}
	}

	var slopes []slopeType
	slope := slopeType{x: 1, y: 1, res: 0}
	slopes = append(slopes, slope)
	slope = slopeType{x: 3, y: 1, res: 0}
	slopes = append(slopes, slope)
	slope = slopeType{x: 5, y: 1, res: 0}
	slopes = append(slopes, slope)
	slope = slopeType{x: 7, y: 1, res: 0}
	slopes = append(slopes, slope)
	slope = slopeType{x: 1, y: 2, res: 0}
	slopes = append(slopes, slope)

	xs, ys := 0, 0
	trees := 0
	part1, part2 := 0, 1

	for k := 0; k < 5; k++ {
		for ys <= ymax {
			if treeMap[ys][xs] == '#' {
				trees++
			}
			xs += slopes[k].x
			ys += slopes[k].y

			if xs > xmax {
				xs = xs - xmax - 1
			}
		}
		slopes[k].res = trees

		if (slopes[k].x == 3) && (slopes[k].y == 1) {
			part1 = trees
		}

		trees = 0
		xs = 0
		ys = 0
	}

	for _, slope := range slopes {
		part2 *= slope.res
	}

	fmt.Printf("Part 1 = %d\n", part1)
	fmt.Printf("Part 2 = %d\n", part2)
}
