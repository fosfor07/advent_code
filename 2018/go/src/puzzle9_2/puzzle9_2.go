package main

import (
	"container/list"
	"fmt"
)

const playersNum = 400
const marbles = 7186400

func main() {
	var player int
	var playersL []int
	playersL = make([]int, playersNum)
	for k := 0; k < playersNum; k++ {
		playersL[k] = 0
	}

	circle := list.New()
	idx := circle.PushBack(0)

	for x := 1; x <= marbles; x++ {
		if circle.Len() < 2 {
			idx = circle.PushBack(x)
		} else if circle.Len() >= 2 {
			if x%23 == 0 {
				playersL[player] += x
				for k := 0; k < 7; k++ {
					idx = idx.Prev()
					if idx == nil {
						idx = circle.Back()
					}
				}
				playersL[player] += idx.Value.(int)
				idxTmp := idx
				idx = idx.Next()
				circle.Remove(idxTmp)
			} else {
				idx = idx.Next()
				if idx == nil {
					idx = circle.Front()
				}
				idx.Next()
				if idx == nil {
					idx = circle.Front()
				}
				idx = circle.InsertAfter(x, idx)
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

	for p, k := range playersL {
		if maxScore < k {
			maxScore = k
			fmt.Printf("Player: %d Score: %d\n", p+1, maxScore)
		}
	}
	fmt.Printf("Max score: %d\n", maxScore)
}
