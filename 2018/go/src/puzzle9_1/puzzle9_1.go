package main

import (
	"fmt"
)

const playersNum = 400
const marblesNum = 71864

func main() {
	var idx, player int
	var playersL []int
	playersL = make([]int, playersNum)
	for k := range playersL {
		playersL[k] = 0
	}

	var circle []int
	circle = append(circle, 0)

	for x := 1; x <= marblesNum; x++ {
		if len(circle) < 2 {
			circle = append(circle, x)
			idx = x
		} else if len(circle) >= 2 {
			if x%23 == 0 {
				playersL[player] += x
				idx -= 7
				if idx < 0 {
					idx = len(circle) + idx
				}
				playersL[player] += circle[idx]
				circle = append(circle[:idx], circle[idx+1:]...)
			} else {
				idx += 2
				if idx > len(circle) {
					idx = idx - len(circle)
					circle = append(circle, 0)
					copy(circle[idx+1:], circle[idx:])
					circle[idx] = x
				} else if idx < len(circle) {
					circle = append(circle, 0)
					copy(circle[idx+1:], circle[idx:])
					circle[idx] = x
				} else {
					circle = append(circle, x)
				}
			}
		}

		if player >= playersNum-1 {
			player = 0
		} else {
			player++
		}

		// DEBUG
		/*fmt.Printf("circle: %v", circle)
		buf := bufio.NewReader(os.Stdin)
		buf.ReadBytes('\n')*/
	}

	maxScore := 0

	for _, k := range playersL {
		if maxScore < k {
			maxScore = k
		}
	}
	fmt.Printf("Max score: %d\n", maxScore)
}
