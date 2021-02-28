package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
	"utils"
)

type dirType struct {
	degrees int
	sign    int
}

type shipType struct {
	x, y int
	dir  dirType
}

type waypointType struct {
	x, y int
}

var dirsMap = map[byte]dirType{
	'N': {degrees: 0, sign: -1},
	'E': {degrees: 90, sign: 1},
	'S': {degrees: 180, sign: 1},
	'W': {degrees: 270, sign: -1},
}

func rotateShip(ship *shipType, dir byte, degrees int) {
	if dir == 'R' {
		ship.dir.degrees = (ship.dir.degrees + degrees) % 360
	} else if dir == 'L' {
		ship.dir.degrees = (ship.dir.degrees - degrees + 360) % 360
	}

	if ship.dir.degrees == 90 || ship.dir.degrees == 180 {
		ship.dir.sign = 1
	} else if ship.dir.degrees == 0 || ship.dir.degrees == 270 {
		ship.dir.sign = -1
	}
}

func rotateWaypoint(wp *waypointType, dir byte, degrees int) {
	x, y := wp.x, wp.y

	if degrees == 180 {
		wp.x = -x
		wp.y = -y
	} else if (dir == 'R' && degrees == 90) || (dir == 'L' && degrees == 270) {
		wp.x = -y
		wp.y = x
	} else if (dir == 'L' && degrees == 90) || (dir == 'R' && degrees == 270) {
		wp.x = y
		wp.y = -x
	}
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

	// part 1
	ship1 := shipType{x: 0, y: 0, dir: dirType{degrees: 90, sign: 1}}

	// part 2
	wp := waypointType{x: 10, y: -1}
	ship2 := waypointType{x: 0, y: 0}

	for _, line := range lines {
		action, valueStr := line[0], line[1:]
		value, _ := strconv.Atoi(valueStr)

		// part 1
		if action == 'R' || action == 'L' {
			rotateShip(&ship1, action, value)
		} else {
			var dirDef dirType
			if action == 'F' {
				dirDef.degrees = ship1.dir.degrees
				dirDef.sign = ship1.dir.sign
			} else {
				dirDef = dirsMap[action]
			}

			if dirDef.degrees == 0 || dirDef.degrees == 180 {
				ship1.y += value * dirDef.sign
			} else if dirDef.degrees == 90 || dirDef.degrees == 270 {
				ship1.x += value * dirDef.sign
			}
		}

		// part 2
		if action == 'R' || action == 'L' {
			rotateWaypoint(&wp, action, value)
		} else if action == 'F' {
			ship2.y += wp.y * value
			ship2.x += wp.x * value
		} else {
			dirDef := dirsMap[action]

			if dirDef.degrees == 0 || dirDef.degrees == 180 {
				wp.y += value * dirDef.sign
			} else if dirDef.degrees == 90 || dirDef.degrees == 270 {
				wp.x += value * dirDef.sign
			}
		}
	}

	part1 := utils.Abs(ship1.x) + utils.Abs(ship1.y)
	part2 := utils.Abs(ship2.x) + utils.Abs(ship2.y)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
