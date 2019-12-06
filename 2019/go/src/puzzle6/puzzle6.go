package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

func count(orbit []string, orbMap map[string][]string) int {
	numOrbs := 0

	numOrbs += len(orbit)

	for _, obj := range orbit {
		if _, ok := orbMap[obj]; ok {
			numOrbs += count(orbMap[obj], orbMap)
		}
	}

	return numOrbs
}

func findPath(destObj string, orbMap map[string][]string, path map[string]int) int {
	jumps := 0

	for key, orbit := range orbMap {
		for _, obj := range orbit {
			if obj == destObj {
				jumps++
				jumps += findPath(key, orbMap, path)
				path[destObj] = jumps
			}
		}
	}

	return jumps
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

	var orbMap map[string][]string
	orbMap = make(map[string][]string)

	// PART 1
	// total number of orbits
	numOrbs := 0

	for _, line := range lines {
		objects := strings.Split(line, ")")
		orbMap[objects[0]] = append(orbMap[objects[0]], objects[1])
	}

	for _, orbit := range orbMap {
		numOrbs += count(orbit, orbMap)
	}

	fmt.Printf("PART 1\n")
	fmt.Printf("Total number of direct and indirect orbits: %d\n", numOrbs)

	// PART 2
	// path to YOU object
	var youPath map[string]int
	youPath = make(map[string]int)
	// path to SAN object
	var sanPath map[string]int
	sanPath = make(map[string]int)

	findPath("YOU", orbMap, youPath)
	findPath("SAN", orbMap, sanPath)

	// number of jumps from COM to the closest common object of YOU and SAN
	cmnJumps := 0

	for obj, jumps := range youPath {
		if _, ok := sanPath[obj]; ok {
			if cmnJumps < jumps {
				cmnJumps = jumps
			}
		}
	}

	minTransfers := (youPath["YOU"] - cmnJumps - 1) + (sanPath["SAN"] - cmnJumps - 1)

	fmt.Printf("PART 2\n")
	fmt.Printf("Minimum number of orbital transfers: %d\n", minTransfers)
}
