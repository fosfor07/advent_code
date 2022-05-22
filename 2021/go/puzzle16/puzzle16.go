package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type opType struct {
	id, result int64
	pkts, bits int64
	vals       []int64
}

func hexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return "", err
	}

	// %08b indicates base 2, zero padded, with 8 characters
	return fmt.Sprintf("%08b", ui), nil
}

func readLiteralValue(binInput string, idx int64) (int64, int64) {
	prefix := binInput[idx]
	idx += 1

	valueStr := ""
	for string(prefix) == "1" {
		valueStr += binInput[idx : idx+4]
		idx += 4
		prefix = binInput[idx]
		idx += 1
	}

	valueStr += binInput[idx : idx+4]
	valueInt, _ := strconv.ParseInt(valueStr, 2, 64)
	idx += 4

	return idx, valueInt
}

func evaluate(op *opType) {
	if op.id == 0 {
		for _, v := range op.vals {
			op.result += v
		}
	} else if op.id == 1 {
		for i, v := range op.vals {
			if i == 0 {
				op.result = v
			} else {
				op.result *= v
			}
		}
	} else if op.id == 2 {
		for i, v := range op.vals {
			if v < op.result || i == 0 {
				op.result = v
			}
		}
	} else if op.id == 3 {
		for _, v := range op.vals {
			if v > op.result {
				op.result = v
			}
		}
	} else if op.id == 5 {
		if op.vals[0] > op.vals[1] {
			op.result = 1
		} else {
			op.result = 0
		}
	} else if op.id == 6 {
		if op.vals[0] < op.vals[1] {
			op.result = 1
		} else {
			op.result = 0
		}
	} else if op.id == 7 {
		if op.vals[0] == op.vals[1] {
			op.result = 1
		} else {
			op.result = 0
		}
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

	binStr, binInput := "", ""

	for _, line := range lines {
		for i := 0; i < len(line); i += 2 {
			v := line[i : i+2]
			binStr, _ = hexToBin(v)
			binInput += binStr
		}
	}

	var part1, part2 int64 = 0, 0
	var idx int64 = 0
	var ops []opType
	var done bool = false

	for {
		ver := binInput[idx : idx+3]
		idx += 3
		verInt, _ := strconv.ParseInt(ver, 2, 64)
		part1 += verInt
		opId := binInput[idx : idx+3]
		opIdInt, _ := strconv.ParseInt(opId, 2, 64)
		idx += 3

		// Update operators which define length in bits.
		for i := range ops {
			if ops[i].bits > 0 {
				ops[i].bits -= 6
			}
		}

		// Update operators which define length in packets.
		if len(ops) > 0 {
			if ops[len(ops)-1].pkts > 0 {
				ops[len(ops)-1].pkts--
			}
		}

		if opId == "100" {
			newIdx, val := readLiteralValue(binInput, idx)
			for i := range ops {
				if ops[i].bits > 0 {
					ops[i].bits -= (newIdx - idx)
				}
			}
			idx = newIdx

			ops[len(ops)-1].vals = append(ops[len(ops)-1].vals, val)
		} else {
			lenType := binInput[idx]
			idx += 1

			if string(lenType) == "0" {
				length := binInput[idx : idx+15]
				idx += 15
				lenInt, _ := strconv.ParseInt(length, 2, 64)

				for i := range ops {
					if ops[i].bits > 0 {
						ops[i].bits -= 16
					}
				}

				ops = append(ops, opType{id: opIdInt, result: 0, pkts: 0, bits: lenInt})
			} else {
				length := binInput[idx : idx+11]
				idx += 11
				lenInt, _ := strconv.ParseInt(length, 2, 64)

				for i := range ops {
					if ops[i].bits > 0 {
						ops[i].bits -= 12
					}
				}

				ops = append(ops, opType{id: opIdInt, result: 0, pkts: lenInt, bits: 0})
			}
		}

		for len(ops) > 0 && ops[len(ops)-1].pkts == 0 && ops[len(ops)-1].bits == 0 {
			evaluate(&ops[len(ops)-1])

			if len(ops) > 1 {
				ops[len(ops)-2].vals = append(ops[len(ops)-2].vals, ops[len(ops)-1].result)
				ops = ops[:len(ops)-1]
			} else if len(ops) == 1 {
				part2 = ops[len(ops)-1].result
				done = true
				break
			}
		}

		if done {
			break
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
