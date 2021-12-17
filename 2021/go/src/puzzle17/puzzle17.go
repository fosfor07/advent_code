package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
)

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

	part1, part2, part1X, part1Y := 0, 0, 0, 0
	velX, velY := 0, 0                   // velocity
	minX, minY, maxX, maxY := 0, 0, 0, 0 // target

	for _, line := range lines {
		tokens := strings.Fields(line)
		h := strings.Split(tokens[2], "=")
		v := strings.Split(tokens[3], "=")

		xs := strings.Split(h[1], "..")
		xs[1] = strings.Trim(xs[1], ",")
		minX, _ = strconv.Atoi(xs[0])
		maxX, _ = strconv.Atoi(xs[1])

		ys := strings.Split(v[1], "..")
		minY, _ = strconv.Atoi(ys[0])
		maxY, _ = strconv.Atoi(ys[1])
	}

	for x := 1; x <= maxX; x++ {
		for y := minY; y < 500; y++ {
			velX, velY = x, y
			posX, posY, highY := 0, 0, 0
			hit := false

			for {
				posX += velX
				posY += velY
				if velX > 0 {
					velX--
				} else if velX < 0 {
					velX++
				}
				velY--

				if posY > highY {
					highY = posY
				}

				if posX >= minX && posX <= maxX && posY >= minY && posY <= maxY {
					part2++
					hit = true
					break
				}

				if posY < minY {
					break
				}
			}

			if hit && highY > part1 {
				part1 = highY
				part1X = x
				part1Y = y
			}
		}
	}

	fmt.Printf("Part 1: %d (%d, %d)\n", part1, part1X, part1Y)
	fmt.Printf("Part 2: %d\n", part2)
}
