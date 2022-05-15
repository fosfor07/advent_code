package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

func getPositions(intCodes []int, opcode, idx int) (int, int, int) {
	pos1, pos2, pos3 := 0, 0, 0

	if opcode == 1 || opcode == 2 {
		pos1 = intCodes[idx+1]
		pos2 = intCodes[idx+2]
		pos3 = intCodes[idx+3]
	} else if opcode == 3 || opcode == 4 {
		pos1 = intCodes[idx+1]
	} else if opcode == 5 || opcode == 6 {
		pos1 = intCodes[idx+1]
		pos2 = intCodes[idx+2]
	} else if opcode == 7 || opcode == 8 {
		pos1 = intCodes[idx+1]
		pos2 = intCodes[idx+2]
		pos3 = intCodes[idx+3]
	}

	return pos1, pos2, pos3
}

func getArguments(intCodes []int, idx, pos1, pos2, mode1, mode2 int) (int, int) {
	arg1, arg2 := 0, 0

	if mode1 == 0 {
		arg1 = intCodes[pos1]
	} else if mode1 == 1 {
		arg1 = intCodes[idx+1]
	}

	if mode2 == 0 {
		arg2 = intCodes[pos2]
	} else if mode2 == 1 {
		arg2 = intCodes[idx+2]
	}

	return arg1, arg2
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = bytes.TrimSpace(input)

	intCodesInit, err := readers.ReadCsvInts(strings.NewReader(string(input)))
	if err != nil {
		panic(err)
	}

	intCodes := make([]int, len(intCodesInit))
	inputs := []int{1, 5}

	for part, input := range inputs {
		idx := 0

		// reset the computer's memory
		numCodes := copy(intCodes, intCodesInit)
		if numCodes != len(intCodesInit) {
			panic(fmt.Sprintf("Error: Copy failed! Copied elements: %d\n", numCodes))
		}

		// execute the program
		for {
			strCode := strconv.Itoa(intCodes[idx])
			if len(strCode) < 5 {
				strCode = fmt.Sprintf("%05s", strCode)
			}

			opcode, _ := strconv.Atoi(string(strCode[3:]))
			mode1, _ := strconv.Atoi(string(strCode[2]))
			mode2, _ := strconv.Atoi(string(strCode[1]))

			if opcode == 1 {
				pos1, pos2, pos3 := getPositions(intCodes, opcode, idx)
				arg1, arg2 := getArguments(intCodes, idx, pos1, pos2, mode1, mode2)

				intCodes[pos3] = arg1 + arg2
				idx += 4
			} else if opcode == 2 {
				pos1, pos2, pos3 := getPositions(intCodes, opcode, idx)
				arg1, arg2 := getArguments(intCodes, idx, pos1, pos2, mode1, mode2)

				intCodes[pos3] = arg1 * arg2
				idx += 4
			} else if opcode == 3 {
				pos1, _, _ := getPositions(intCodes, opcode, idx)
				intCodes[pos1] = input
				idx += 2
			} else if opcode == 4 {
				pos1, _, _ := getPositions(intCodes, opcode, idx)
				output := -1

				if mode1 == 0 {
					output = intCodes[pos1]
				} else if mode1 == 1 {
					output = intCodes[idx+1]
				}

				if output != 0 {
					if part == 0 {
						fmt.Printf("Part 1: %d\n", output)
					} else if part == 1 {
						fmt.Printf("Part 2: %d\n", output)
					}
				}
				idx += 2
			} else if opcode == 5 {
				pos1, pos2, _ := getPositions(intCodes, opcode, idx)
				arg1, arg2 := getArguments(intCodes, idx, pos1, pos2, mode1, mode2)

				if arg1 != 0 {
					idx = arg2
				} else {
					idx += 3
				}
			} else if opcode == 6 {
				pos1, pos2, _ := getPositions(intCodes, opcode, idx)
				arg1, arg2 := getArguments(intCodes, idx, pos1, pos2, mode1, mode2)

				if arg1 == 0 {
					idx = arg2
				} else {
					idx += 3
				}
			} else if opcode == 7 {
				pos1, pos2, pos3 := getPositions(intCodes, opcode, idx)
				arg1, arg2 := getArguments(intCodes, idx, pos1, pos2, mode1, mode2)

				if arg1 < arg2 {
					intCodes[pos3] = 1
				} else {
					intCodes[pos3] = 0
				}
				idx += 4
			} else if opcode == 8 {
				pos1, pos2, pos3 := getPositions(intCodes, opcode, idx)
				arg1, arg2 := getArguments(intCodes, idx, pos1, pos2, mode1, mode2)

				if arg1 == arg2 {
					intCodes[pos3] = 1
				} else {
					intCodes[pos3] = 0
				}
				idx += 4
			} else if opcode == 99 {
				break
			}
		}
	}
}
