package main

import (
	"fmt"
	"sort"
	"strconv"
)

//const serial = 18
const serial = 7672

func main() {
	var grid [300][300]int

	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			var power int
			var rackID int
			var pS string
			rackID = x + 10
			power = rackID * y
			power += serial
			power *= rackID
			pS = strconv.Itoa(power)
			power, _ = strconv.Atoi(string(pS[len(pS)-3]))
			power -= 5

			grid[x][y] = power
		}
	}

	var regions []int
	var maxReg, maxX, maxY int

	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			if x+2 < 300 && y+2 < 300 {
				var region int
				for n := y; n < (y + 3); n++ {
					for m := x; m < (x + 3); m++ {
						region += grid[m][n]
					}
				}
				if region > maxReg {
					maxReg = region
					maxX = x
					maxY = y
				}
				regions = append(regions, region)
			}
		}
	}

	sort.Ints(regions)
	fmt.Printf("The smallest total power: %d\n", regions[0])
	fmt.Printf("The largest total power: %d\n", regions[len(regions)-1])
	fmt.Printf("ID: %d,%d\n", maxX, maxY)
}
