package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

type busType struct {
	busID, expMod int
}

// returns x where (a * x) % b == 1
func modMultiplicativeInverse(a int, b int) int {
	b0, t, q := b, 0, 0
	x0, x1 := 0, 1
	if b == 1 {
		return 1
	}

	for a > 1 {
		q = a / b
		t = b
		b = a % b
		a = t
		t = x0
		x0 = x1 - q*x0
		x1 = t
	}
	if x1 < 0 {
		x1 += b0
	}
	return x1
}

func chineseRemainderTheorem(a, n []int) int {
	prod := 1
	for i := 0; i < len(n); i++ {
		prod *= n[i]
	}

	sum := 0
	for i := 0; i < len(n); i++ {
		p := prod / n[i]
		sum += a[i] * modMultiplicativeInverse(p, n[i]) * p
	}
	return sum % prod
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

	dTime, _ := strconv.Atoi(lines[0])
	buses := strings.Split(lines[1], ",")
	minWait, bestBus := 0, 0

	var busList []busType

	for i, bus := range buses {
		if bus != "x" {
			busID, _ := strconv.Atoi(bus)
			waitTime := busID - dTime%busID
			if waitTime < minWait || (minWait == 0 && dTime%busID != 0) {
				minWait = waitTime
				bestBus = busID
			}

			b := busType{busID: busID, expMod: (busID - i) % busID}
			busList = append(busList, b)
		}
	}

	var a, n []int
	for _, bus := range busList {
		a = append(a, bus.expMod)
		n = append(n, bus.busID)
	}
	part2 := chineseRemainderTheorem(a, n)

	fmt.Printf("Part 1: %d\n", bestBus*minWait)
	fmt.Printf("Part 2: %d\n", part2)
}
