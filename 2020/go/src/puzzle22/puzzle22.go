package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
	"utils"
)

func checkHistory(history [][]int, deck []int) bool {
	for _, hDeck := range history {
		eq := utils.EqualIntSlices(deck, hDeck)
		if eq {
			return true
		}
	}

	return false
}

func subGame(deck1, deck2 []int, card1, card2 int) int {
	winner := 1
	sDeck1 := make([]int, card1)
	sDeck2 := make([]int, card2)
	copy(sDeck1, deck1)
	copy(sDeck2, deck2)

	var hist1, hist2 [][]int

	for len(sDeck1) > 0 && len(sDeck2) > 0 {
		if checkHistory(hist1, sDeck1) {
			return winner
		}
		if checkHistory(hist2, sDeck2) {
			return winner
		}

		hist1 = append(hist1, sDeck1)
		hist2 = append(hist2, sDeck2)

		sCard1, sCard2 := sDeck1[0], sDeck2[0]
		sDeck1, sDeck2 = sDeck1[1:], sDeck2[1:]

		if len(sDeck1) >= sCard1 && len(sDeck2) >= sCard2 {
			sWinner := subGame(sDeck1, sDeck2, sCard1, sCard2)
			if sWinner == 1 {
				sDeck1 = append(sDeck1, sCard1, sCard2)
			} else {
				sDeck2 = append(sDeck2, sCard2, sCard1)
			}
		} else {
			if sCard1 > sCard2 {
				sDeck1 = append(sDeck1, sCard1, sCard2)
			} else {
				sDeck2 = append(sDeck2, sCard2, sCard1)
			}
		}
	}

	if len(sDeck2) > 0 {
		winner = 2
	}

	return winner
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

	part1, part2 := 0, 0
	var deck1, deck2 []int
	deckNum := 0

	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}
		if line[0] == 'P' {
			deckNum++
			continue
		}

		if deckNum == 1 {
			card, _ := strconv.Atoi(line)
			deck1 = append(deck1, card)
		} else {
			card, _ := strconv.Atoi(line)
			deck2 = append(deck2, card)
		}
	}

	// part 1
	deck1p1 := make([]int, len(deck1))
	deck2p1 := make([]int, len(deck2))
	copy(deck1p1, deck1)
	copy(deck2p1, deck2)

	for len(deck1p1) > 0 && len(deck2p1) > 0 {
		card1, card2 := deck1p1[0], deck2p1[0]
		deck1p1, deck2p1 = deck1p1[1:], deck2p1[1:]

		if card1 > card2 {
			deck1p1 = append(deck1p1, card1, card2)
		} else {
			deck2p1 = append(deck2p1, card2, card1)
		}
	}

	winDeck := deck2p1
	if len(deck1p1) > 0 {
		winDeck = deck1p1
	}

	value := 1
	for i := len(winDeck) - 1; i >= 0; i-- {
		part1 += value * winDeck[i]
		value++
	}

	// part 2
	var hist1, hist2 [][]int

	for len(deck1) > 0 && len(deck2) > 0 {
		if checkHistory(hist1, deck1) {
			break
		}
		if checkHistory(hist2, deck2) {
			break
		}

		hist1 = append(hist1, deck1)
		hist2 = append(hist2, deck2)

		card1, card2 := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]

		if len(deck1) >= card1 && len(deck2) >= card2 {
			winner := subGame(deck1, deck2, card1, card2)
			if winner == 1 {
				deck1 = append(deck1, card1, card2)
			} else {
				deck2 = append(deck2, card2, card1)
			}
		} else {
			if card1 > card2 {
				deck1 = append(deck1, card1, card2)
			} else {
				deck2 = append(deck2, card2, card1)
			}
		}
	}

	winDeck = deck2
	if len(deck1) > 0 {
		winDeck = deck1
	}

	value = 1
	for i := len(winDeck) - 1; i >= 0; i-- {
		part2 += value * winDeck[i]
		value++
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
