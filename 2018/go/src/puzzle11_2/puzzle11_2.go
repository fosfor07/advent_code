package main

import (
	"fmt"
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

	var maxReg, maxX, maxY, maxSize int

	for size := 1; size < 301; size++ {
		for y := 0; y < 300; y++ {
			for x := 0; x < 300; x++ {
				if x+size-1 < 300 && y+size-1 < 300 {
					var region int
					for n := y; n < (y + size); n++ {
						for m := x; m < (x + size); m++ {
							region += grid[m][n]
						}
					}
					if region > maxReg {
						maxReg = region
						maxX = x
						maxY = y
						maxSize = size
					}
				}
			}
		}
		fmt.Printf("Size: %d\n", size)
	}

	fmt.Printf("ID: %d,%d,%d\n", maxX, maxY, maxSize)
}
