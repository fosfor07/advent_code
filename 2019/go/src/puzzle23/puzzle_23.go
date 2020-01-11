package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
	"time"
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

const computersNum int = 50
const natAddress int = 255
const aliveCode int = 200

// timeouts in milliseconds
// natTimeout must be greater than computerTimeout
const computerTimeout time.Duration = 5
const natTimeout time.Duration = 10

type msgType struct {
	a int
	x int
	y int
}

func execute(intCodesInit []int, addr int, channels []chan msgType,
	natChan chan msgType) {

	intCodes := make([]int, len(intCodesInit))
	intCodes = extend(intCodes, 5000)

	// initialize the computer's memory
	numCodes := copy(intCodes, intCodesInit)
	if numCodes != len(intCodesInit) {
		panic(fmt.Sprintf("Error: Copy failed! Copied elements: %d\n", numCodes))
	}

	idx := 0
	opcode := 0
	relBase := 0

	readAddr := true
	expInY := false
	expOutX := false
	expOutY := false
	firstNAT := true

	var sndMsg msgType
	var rcvMsg msgType
	var natMsg msgType
	natMsg.a = aliveCode

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

			// computer requests its network address
			// via a single input instruction
			if readAddr {
				if mode1 == 0 {
					intCodes[pos1] = addr
				} else if mode1 == 2 {
					intCodes[pos1+relBase] = addr
				}
				readAddr = false
			} else {
				msgVal := -1
				if expInY {
					msgVal = rcvMsg.y
					expInY = false
				} else {
					select {
					case rcvMsg = <-channels[addr]:
						//fmt.Printf("Computer %d received X: %d, Y: %d addressed to %d\n",
						//	addr, rcvMsg.x, rcvMsg.y, rcvMsg.a)

						msgVal = rcvMsg.x
						expInY = true
					case <-time.After(time.Millisecond * computerTimeout):
						if expInY {
							fmt.Printf("ERROR: expInY is true for computer %d!!!\n", addr)
						}
						msgVal = -1
					}
				}

				if mode1 == 0 {
					intCodes[pos1] = msgVal
				} else if mode1 == 2 {
					intCodes[pos1+relBase] = msgVal
				}
			}
			idx += 2
		} else if opcode == 4 {
			pos1, _, _ := getPositions(opcode, intCodes, idx)

			msgVal := 0
			if mode1 == 0 {
				msgVal = intCodes[pos1]
			} else if mode1 == 1 {
				msgVal = intCodes[idx+1]
			} else if mode1 == 2 {
				msgVal = intCodes[pos1+relBase]
			}

			if expOutX {
				// set X
				sndMsg.x = msgVal
				expOutX = false
				expOutY = true
			} else if expOutY {
				// set Y and send message
				sndMsg.y = msgVal
				expOutY = false
				//fmt.Printf("Computer %d sends X: %d, Y: %d to address %d\n",
				//	addr, sndMsg.x, sndMsg.y, sndMsg.a)
				if sndMsg.a == natAddress {
					if firstNAT {
						fmt.Printf("\nPART 1\nY = %d", sndMsg.y)
						firstNAT = false
					}
					natChan <- sndMsg
				} else {
					// inform NAT that network is not idle
					natChan <- natMsg
					channels[sndMsg.a] <- sndMsg
				}
			} else {
				// set address
				sndMsg.a = msgVal
				expOutX = true
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

func natProcess(natChan chan msgType, respChan chan msgType) {
	var lastRcvd msgType
	var lastSent msgType
	// check if NAT received any message
	rcvdFlag := false

	// NAT always sends messages to computer #0
	lastRcvd.a = 0

	for {
		select {
		case rcvMsg := <-natChan:
			if rcvMsg.a != aliveCode {
				// update the last received message
				lastRcvd.x = rcvMsg.x
				lastRcvd.y = rcvMsg.y
				rcvdFlag = true
				//fmt.Printf("NAT updated last received msg! X: %d, Y: %d\n",
				//	lastRcvd.x, lastRcvd.y)
			}
		case <-time.After(time.Millisecond * natTimeout):
			if (lastSent.y == lastRcvd.y) && rcvdFlag {
				fmt.Printf("\nPART 2\nY = %d", lastSent.y)
				return
			}

			//fmt.Printf("Timeout -> NAT sends msg to computer %d. X: %d, Y: %d\n",
			//	lastRcvd.a, lastRcvd.x, lastRcvd.y)

			// send message to computer #0
			respChan <- lastRcvd

			// update the last sent message
			lastSent.a = lastRcvd.a
			lastSent.x = lastRcvd.x
			lastSent.y = lastRcvd.y
		}
	}
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

	var channels []chan msgType

	for addr := 0; addr < computersNum; addr++ {
		c := make(chan msgType)
		channels = append(channels, c)
	}

	natChan := make(chan msgType)

	for addr := 0; addr < computersNum; addr++ {
		go execute(intCodesInit, addr, channels, natChan)
	}

	natProcess(natChan, channels[0])
}
