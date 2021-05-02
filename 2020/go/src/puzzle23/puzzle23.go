package main

import (
	"fmt"
	"strconv"
)

type cupType struct {
	label int
	next  *cupType
}

type listType struct {
	head, tail, curCup *cupType
	cups               map[int]*cupType
	highLabel          int
}

func createCupsList(cupList *listType) {
	for _, ch := range labelsString {
		labelInt, _ := strconv.Atoi(string(ch))

		c := cupType{label: labelInt, next: nil}
		if cupList.head == nil {
			cupList.head = &c
		} else {
			cupList.tail.next = &c
		}
		cupList.tail = &c

		cupList.cups[labelInt] = &c
	}
	cupList.tail.next = cupList.head
	cupList.curCup = cupList.head
}

func selectDestinationLabel(cupList *listType, pickedCups [3]*cupType) int {
	destLabel := cupList.curCup.label - 1
	for {
		if destLabel < lowLabel {
			destLabel = cupList.highLabel
		}
		fnd := false
		for _, c := range pickedCups {
			if c.label == destLabel {
				destLabel--
				fnd = true
				break
			}
		}
		if !fnd {
			break
		}
	}

	return destLabel
}

func playGame(cupList *listType, iterations int) {
	for i := 0; i < iterations; i++ {
		var pickedCups [3]*cupType

		// Pick up the three cups that are immediately clockwise of the current cup.
		tmpCup := cupList.curCup
		for i := 0; i < 3; i++ {
			pickedCups[i] = tmpCup.next
			tmpCup = tmpCup.next
		}
		cupList.curCup.next = tmpCup.next

		tmpCup = cupList.cups[selectDestinationLabel(cupList, pickedCups)]

		// Place the picked up cups so that they are immediately clockwise of the destination cup.
		pickedCups[2].next = tmpCup.next
		tmpCup.next = pickedCups[0]

		// Select a new current cup.
		cupList.curCup = cupList.curCup.next
	}
}

const labelsString string = "167248359"
const lowLabel int = 1
const highLabel1 int = 9
const numIterations1 int = 100
const highLabel2 int = 1000000
const numIterations2 int = 10000000

func main() {
	var cupList1 listType
	cupList1.highLabel = highLabel1
	cupList1.cups = make(map[int]*cupType)
	createCupsList(&cupList1)

	var cupList2 listType
	cupList2.highLabel = highLabel2
	cupList2.cups = make(map[int]*cupType)
	createCupsList(&cupList2)

	for i := highLabel1 + 1; i <= highLabel2; i++ {
		c := cupType{label: i, next: nil}
		cupList2.tail.next = &c
		cupList2.tail = &c

		cupList2.cups[i] = &c
	}
	cupList2.tail.next = cupList2.head

	playGame(&cupList1, numIterations1)

	part1 := ""
	nextCup := cupList1.cups[lowLabel].next
	for {
		if nextCup.label == lowLabel {
			break
		}
		labelStr := strconv.Itoa(nextCup.label)
		part1 += labelStr
		nextCup = nextCup.next
	}

	playGame(&cupList2, numIterations2)
	label1Cup := cupList2.cups[lowLabel]
	firstCup := label1Cup.next
	secondCup := firstCup.next

	fmt.Printf("Part 1: %s\n", part1)
	fmt.Printf("Part 2: %d\n", firstCup.label*secondCup.label)
}
