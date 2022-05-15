package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fosfor07/advent_code/pkg/readers"
)

const boardSize int = 5

type fieldType struct {
	r, c int
}

func calculateFinalScore(boards [][boardSize][boardSize]int, markedNums [][boardSize * boardSize]fieldType, boardId int, lastIndex int) int {
	sumMarked, sumBoard := 0, 0

	for f := 0; f <= lastIndex; f++ {
		sumMarked += boards[boardId][markedNums[boardId][f].r][markedNums[boardId][f].c]
	}
	for r := 0; r < boardSize; r++ {
		for c := 0; c < boardSize; c++ {
			sumBoard += boards[boardId][r][c]
		}
	}
	score := (sumBoard - sumMarked) * boards[boardId][markedNums[boardId][lastIndex].r][markedNums[boardId][lastIndex].c]

	return score
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

	var drawStr string
	var boards [][boardSize][boardSize]int
	var board [boardSize][boardSize]int

	row := 0

	for lineNum, line := range lines {
		if lineNum == 0 {
			drawStr = line
			continue
		} else if line == "" {
			row = 0
			boards = append(boards, board)
			continue
		}

		nums := strings.Fields(line)

		for col, n := range nums {
			nInt, _ := strconv.Atoi(n)
			board[row][col] = nInt
		}
		row++
	}

	var markedNums [][boardSize * boardSize]fieldType
	var fields [boardSize * boardSize]fieldType
	var markedCount []int
	for b := 0; b < len(boards); b++ {
		markedNums = append(markedNums, fields)
		markedCount = append(markedCount, 0)
	}

	b1, lastIndexOnb1, b2, lastIndexOnb2 := 0, 0, 0, 0
	winners := make(map[int]int)

	drawnNumbers := strings.Split(drawStr, ",")
	for _, d := range drawnNumbers {
		dInt, _ := strconv.Atoi(d)

		for b := 0; b < len(boards); b++ {
			for r := 0; r < boardSize; r++ {
				for c := 0; c < boardSize; c++ {
					if boards[b][r][c] == dInt {
						fieldIdx := markedCount[b]
						markedNums[b][fieldIdx].r = r
						markedNums[b][fieldIdx].c = c
						markedCount[b]++
					}
				}
			}
		}

		win := false
		var winBoards []int

		for b := 0; b < len(boards); b++ {
			var rows [boardSize]int
			var cols [boardSize]int
			for f := 0; f < markedCount[b]; f++ {
				rows[markedNums[b][f].r]++
				cols[markedNums[b][f].c]++
			}

			for _, r := range rows {
				if r == boardSize {
					win = true
					winBoards = append(winBoards, b)
				}
			}

			for _, c := range cols {
				if c == boardSize {
					win = true
					winBoards = append(winBoards, b)
				}
			}
		}

		if win {
			for _, winB := range winBoards {
				if _, ok := winners[winB]; !ok {
					winners[winB] = 1
					b2 = winB
				}
			}
			if len(winners) == 1 {
				b1 = b2
				lastIndexOnb1 = markedCount[b1] - 1
			}
			if len(winners) >= 100 {
				lastIndexOnb2 = markedCount[b2] - 1
				break
			}
		}
	}

	part1 := calculateFinalScore(boards, markedNums, b1, lastIndexOnb1)
	part2 := calculateFinalScore(boards, markedNums, b2, lastIndexOnb2)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
