package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

type slopeType struct {
	x, y int
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

	xmax := len(lines[0]) - 1
	ymax := len(lines) - 1

	treeMap := make([][]rune, ymax+1)

	for y, line := range lines {
		treeMap[y] = make([]rune, xmax+1)
		for x, ch := range line {
			treeMap[y][x] = ch
		}
	}

	var slopes []slopeType
	slope := slopeType{x: 1, y: 1}
	slopes = append(slopes, slope)
	slope = slopeType{x: 3, y: 1}
	slopes = append(slopes, slope)
	slope = slopeType{x: 5, y: 1}
	slopes = append(slopes, slope)
	slope = slopeType{x: 7, y: 1}
	slopes = append(slopes, slope)
	slope = slopeType{x: 1, y: 2}
	slopes = append(slopes, slope)

	part1, part2 := 0, 1

	for _, slope := range slopes {
		x, y := 0, 0
		trees := 0

		for y <= ymax {
			if treeMap[y][x] == '#' {
				trees++
			}
			x += slope.x
			y += slope.y

			if x > xmax {
				x = x - xmax - 1
			}
		}

		if (slope.x == 3) && (slope.y == 1) {
			part1 = trees
		}
		part2 *= trees
	}

	fmt.Printf("Part 1 = %d\n", part1)
	fmt.Printf("Part 2 = %d\n", part2)
}
