package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"

	"github.com/fosfor07/advent_code/pkg/readers"
)

func extend(intCodes []int, newLen int) []int {
	for i := len(intCodes); i <= newLen; i++ {
		intCodes = append(intCodes, 0)
	}
	return intCodes
}

func getPositions(opcode int, intCodes []int, idx int) (int, int, int) {
	pos1 := 0
	pos2 := 0
	pos3 := 0

	if opcode == 1 || opcode == 2 {
		pos1 = intCodes[idx+1]
		pos2 = intCodes[idx+2]
		pos3 = intCodes[idx+3]
	} else if opcode == 3 || opcode == 4 || opcode == 9 {
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

func getArguments(intCodes []int, idx int,
	pos1 int, pos2 int, pos3 int,
	mode1 int, mode2 int, mode3 int,
	relBase int) (int, int, int) {

	arg1 := 0
	arg2 := 0
	arg3 := 0

	if mode1 == 0 {
		arg1 = intCodes[pos1]
	} else if mode1 == 1 {
		arg1 = intCodes[idx+1]
	} else if mode1 == 2 {
		arg1 = intCodes[pos1+relBase]
	}

	if mode2 == 0 {
		arg2 = intCodes[pos2]
	} else if mode2 == 1 {
		arg2 = intCodes[idx+2]
	} else if mode2 == 2 {
		arg2 = intCodes[pos2+relBase]
	}

	if mode3 == 0 {
		arg3 = pos3
	} else if mode3 == 1 {
		arg3 = idx + 3
	} else if mode3 == 2 {
		arg3 = pos3 + relBase
	}

	return arg1, arg2, arg3
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
	idx := 0
	opcode := 0
	relBase := 0
	inputs := []int{1, 2}

	intCodes = extend(intCodes, 5000)

	program := ""
	sIdx := 0
	output := ""

	for _, input := range inputs {
		fmt.Printf("\nPART %d\n", input)
		program = ""
		sIdx = 0

		if input == 1 {
			program += "NOT C T\n"
			program += "AND D T\n"
			program += "NOT B J\n"
			program += "AND D J\n"
			program += "OR T J\n"
			program += "NOT A T\n"
			program += "OR T J\n"
			program += "WALK\n"
		} else if input == 2 {
			program += "NOT C T\n"
			program += "AND D T\n"
			program += "AND H T\n"
			program += "NOT B J\n"
			program += "AND D J\n"
			program += "OR T J\n"
			program += "NOT A T\n"
			program += "OR T J\n"
			program += "RUN\n"
		}

		idx = 0
		relBase = 0

		// reset the computer's memory
		numCodes := copy(intCodes, intCodesInit)
		if numCodes != len(intCodesInit) {
			panic(fmt.Sprintf("Error: Copy failed! Copied elements: %d\n", numCodes))
		}

		// execute the program
		for {
			// determine opcode
			strCode := strconv.Itoa(intCodes[idx])
			if len(strCode) < 5 {
				strCode = fmt.Sprintf("%05s", strCode)
			}

			opcode, _ = strconv.Atoi(string(strCode[3:]))
			mode1, _ := strconv.Atoi(string(strCode[2]))
			mode2, _ := strconv.Atoi(string(strCode[1]))
			mode3, _ := strconv.Atoi(string(strCode[0]))

			if opcode == 1 {
				pos1, pos2, pos3 := getPositions(opcode, intCodes, idx)
				arg1, arg2, arg3 := getArguments(intCodes, idx, pos1, pos2, pos3,
					mode1, mode2, mode3, relBase)

				intCodes[arg3] = arg1 + arg2
				idx += 4
			} else if opcode == 2 {
				pos1, pos2, pos3 := getPositions(opcode, intCodes, idx)
				arg1, arg2, arg3 := getArguments(intCodes, idx, pos1, pos2, pos3,
					mode1, mode2, mode3, relBase)

				intCodes[arg3] = arg1 * arg2
				idx += 4
			} else if opcode == 3 {
				pos1, _, _ := getPositions(opcode, intCodes, idx)

				if mode1 == 0 {
					intCodes[pos1] = int(program[sIdx])
				} else if mode1 == 2 {
					intCodes[pos1+relBase] = int(program[sIdx])
				}
				sIdx++
				idx += 2
			} else if opcode == 4 {
				pos1, _, _ := getPositions(opcode, intCodes, idx)

				asciiCode := 0
				if mode1 == 0 {
					asciiCode = intCodes[pos1]
				} else if mode1 == 1 {
					asciiCode = intCodes[idx+1]
				} else if mode1 == 2 {
					asciiCode = intCodes[pos1+relBase]
				}

				if asciiCode <= unicode.MaxASCII {
					if asciiCode == 10 {
						fmt.Printf("%s\n", output)
						output = ""
					} else {
						output += string(asciiCode)
					}
				} else {
					fmt.Printf("Damage: %d\n", asciiCode)
				}
				idx += 2
			} else if opcode == 5 {
				pos1, pos2, _ := getPositions(opcode, intCodes, idx)
				arg1, arg2, _ := getArguments(intCodes, idx, pos1, pos2, 0,
					mode1, mode2, mode3, relBase)

				if arg1 != 0 {
					idx = arg2
				} else {
					idx += 3
				}
			} else if opcode == 6 {
				pos1, pos2, _ := getPositions(opcode, intCodes, idx)
				arg1, arg2, _ := getArguments(intCodes, idx, pos1, pos2, 0,
					mode1, mode2, mode3, relBase)

				if arg1 == 0 {
					idx = arg2
				} else {
					idx += 3
				}
			} else if opcode == 7 {
				pos1, pos2, pos3 := getPositions(opcode, intCodes, idx)
				arg1, arg2, arg3 := getArguments(intCodes, idx, pos1, pos2, pos3,
					mode1, mode2, mode3, relBase)

				if arg1 < arg2 {
					intCodes[arg3] = 1
				} else {
					intCodes[arg3] = 0
				}
				idx += 4
			} else if opcode == 8 {
				pos1, pos2, pos3 := getPositions(opcode, intCodes, idx)
				arg1, arg2, arg3 := getArguments(intCodes, idx, pos1, pos2, pos3,
					mode1, mode2, mode3, relBase)

				if arg1 == arg2 {
					intCodes[arg3] = 1
				} else {
					intCodes[arg3] = 0
				}
				idx += 4
			} else if opcode == 9 {
				pos1, _, _ := getPositions(opcode, intCodes, idx)
				arg1, _, _ := getArguments(intCodes, idx, pos1, 0, 0,
					mode1, mode2, mode3, relBase)

				relBase += arg1
				idx += 2
			} else if opcode == 99 {
				break
			}
		}
	}
}
