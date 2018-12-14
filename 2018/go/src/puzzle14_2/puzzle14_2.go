package main

import (
	"container/list"
	"fmt"
	"strconv"
)

const pattern = "846601"

func main() {
	scores := list.New()
	score1 := scores.PushBack(3)
	score2 := scores.PushBack(7)

	var one, two, newR int
	var newRStr, matchStr string
	cnt := 2
	end := false

	for true {
		one = score1.Value.(int)
		two = score2.Value.(int)
		newR = one + two

		newRStr = strconv.Itoa(newR)
		for k := 0; k < len(newRStr); k++ {
			matchStr += string(newRStr[k])
			if len(matchStr) > len(pattern) {
				matchStr = matchStr[1:]
			}

			newElem := int(newRStr[k] - '0')
			scores.PushBack(newElem)
			cnt++

			if matchStr == pattern {
				fmt.Printf("Result: %d\n", cnt-len(pattern))
				end = true
				break
			}
		}

		if end {
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
}
