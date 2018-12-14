package main

import (
	"container/list"
	"fmt"
	"strconv"
)

const trainRecNum = 846601
const resultRecNum = 10

func main() {
	scores := list.New()
	score1 := scores.PushBack(3)
	score2 := scores.PushBack(7)

	var one, two, newR int
	var newRS string
	cnt := 2

	for true {
		one = score1.Value.(int)
		two = score2.Value.(int)
		newR = one + two

		newRS = strconv.Itoa(newR)
		for k := 0; k < len(newRS); k++ {
			newInt := int(newRS[k] - '0')
			scores.PushBack(newInt)
			cnt++
		}
		if cnt >= (trainRecNum)+(resultRecNum) {
			break
		}

		steps1 := one + 1
		steps2 := two + 1
		for k := 0; k < steps1; k++ {
			score1 = score1.Next()
			if score1 == nil {
				score1 = scores.Front()
			}
		}
		for k := 0; k < steps2; k++ {
			score2 = score2.Next()
			if score2 == nil {
				score2 = scores.Front()
			}
		}
	}

	resCnt := 1
	recNum := 0
	skip := true

	fmt.Print("Result: ")
	for e := scores.Front(); e != nil; e = e.Next() {
		if recNum == trainRecNum {
			skip = false
		}
		recNum++

		if skip == false {
			fmt.Printf("%d", e.Value.(int))
			resCnt++
			if resCnt > resultRecNum {
				break
			}
		}
	}
}
