package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strings"
)

func execute(intCodes []int) {
	idx := 0

	for {
		if intCodes[idx] == 1 {
			pos1 := intCodes[idx+1]
			pos2 := intCodes[idx+2]
			pos3 := intCodes[idx+3]
			intCodes[pos3] = intCodes[pos1] + intCodes[pos2]
			idx += 4
		} else if intCodes[idx] == 2 {
			pos1 := intCodes[idx+1]
			pos2 := intCodes[idx+2]
			pos3 := intCodes[idx+3]
			intCodes[pos3] = intCodes[pos1] * intCodes[pos2]
			idx += 4
		} else if intCodes[idx] == 99 {
			break
		}
	}
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	intCodes, err := readers.ReadCsvInts(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	intCodesInit := make([]int, len(intCodes))
	// store the initial computer's memory in intCodesInit
	numCodes := copy(intCodesInit, intCodes)
	if numCodes != len(intCodes) {
		panic(fmt.Sprintf("Error: Copy failed! Copied elements: %d\n", numCodes))
	}

	// part 1
	// initialize inputs - values at addresses 1 and 2
	intCodes[1] = 12 // noun
	intCodes[2] = 2  // verb

	// execute the program
	execute(intCodes)

	// print output available at address 0
	fmt.Printf("Part 1 = %d\n", intCodes[0])

	// part 2
	const ExpectedOutput int = 19690720
	noun, verb := 0, 0

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			// reset the computer's memory
			numCodes := copy(intCodes, intCodesInit)
			if numCodes != len(intCodesInit) {
				panic(fmt.Sprintf("Error: Copy failed! Copied elements: %d\n", numCodes))
			}

			// initialize inputs - values at addresses 1 and 2
			intCodes[1] = i // noun
			intCodes[2] = j // verb

			// execute the program
			execute(intCodes)

			// verify output available at address 0
			if intCodes[0] == ExpectedOutput {
				noun = i
				verb = j
				break
			}
		}

		if intCodes[0] == ExpectedOutput {
			break
		}
	}

	fmt.Printf("Part 2 = %d (noun = %d, verb = %d)\n", 100*noun+verb, noun, verb)
}
