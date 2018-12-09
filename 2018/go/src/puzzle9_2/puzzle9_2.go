package main

import (
	"container/list"
	"fmt"
)

const playersNum = 400
const marblesNum = 7186400

func main() {
	var player int
	var playersL []int
	playersL = make([]int, playersNum)
	for k := range playersL {
		playersL[k] = 0
	}

	circle := list.New()
	marble := circle.PushBack(0)

	for x := 1; x <= marblesNum; x++ {
		if circle.Len() < 2 {
			marble = circle.PushBack(x)
		} else if circle.Len() >= 2 {
			if x%23 == 0 {
				playersL[player] += x
				for k := 0; k < 7; k++ {
					marble = marble.Prev()
					if marble == nil {
						marble = circle.Back()
					}
				}
				playersL[player] += marble.Value.(int)
				marbleTmp := marble
				marble = marble.Next()
				circle.Remove(marbleTmp)
			} else {
				marble = marble.Next()
				if marble == nil {
					marble = circle.Front()
				}
				marble.Next()
				if marble == nil {
					marble = circle.Front()
				}
				marble = circle.InsertAfter(x, marble)
			}

			// DEBUG
			/*for e := circle.Front(); e != nil; e = e.Next() {
				fmt.Printf("%d ", e.Value.(int))
			}
			buf := bufio.NewReader(os.Stdin)
			buf.ReadBytes('\n')*/
		}

		if player >= playersNum-1 {
			player = 0
		} else {
			player++
		}
	}

	maxScore := 0

	for _, k := range playersL {
		if maxScore < k {
			maxScore = k
		}
	}
	fmt.Printf("Max score: %d\n", maxScore)
}
