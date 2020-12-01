package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	ints, err := readers.ReadInts(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	modFuel, fuel1, fuel2 := 0, 0, 0

	for _, mass := range ints {
		modFuel = mass / 3
		modFuel = modFuel - 2
		fuel1 += modFuel
		fuel2 += modFuel

		for modFuel > 0 {
			modFuel = modFuel / 3
			modFuel = modFuel - 2
			if modFuel > 0 {
				fuel2 += modFuel
			}
		}
	}

	fmt.Printf("Part 1: %d\n", fuel1)
	fmt.Printf("Part 2: %d\n", fuel2)
}
