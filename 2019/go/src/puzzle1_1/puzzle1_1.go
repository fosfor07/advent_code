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

	fuel := 0

	for _, mass := range ints {
		mass = mass / 3
		mass = mass - 2
		fuel += mass
	}

	fmt.Printf("Sum of the fuel requirements: %d\n", fuel)
}
