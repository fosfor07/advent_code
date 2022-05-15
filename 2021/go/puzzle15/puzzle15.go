package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
	"github.com/fosfor07/advent_code/pkg/readers"
)

func buildGraph(g *dijkstra.Graph, riskLevels [][]int64) (int, int) {
	vID, cID := 0, 0
	sID, eID := 0, 0
	exists := make(map[string]int)

	for r := 0; r < len(riskLevels); r++ {
		for c := 0; c < len(riskLevels[0]); c++ {
			key := fmt.Sprintf("r%dc%d", r, c)
			if _, ok := exists[key]; !ok {
				g.AddVertex(vID)
				exists[key] = vID

				if r == 0 && c == 0 {
					sID = vID
				}
				if r == len(riskLevels)-1 && c == len(riskLevels[0])-1 {
					eID = vID
				}
				cID = vID
				vID++
			}

			key = fmt.Sprintf("r%dc%d", r-1, c)
			if _, ok := exists[key]; ok {
				g.AddArc(cID, exists[key], riskLevels[r-1][c])
				g.AddArc(exists[key], cID, riskLevels[r][c])
			}

			key = fmt.Sprintf("r%dc%d", r+1, c)
			if _, ok := exists[key]; ok {
				g.AddArc(cID, exists[key], riskLevels[r+1][c])
				g.AddArc(exists[key], cID, riskLevels[r][c])
			}

			key = fmt.Sprintf("r%dc%d", r, c-1)
			if _, ok := exists[key]; ok {
				g.AddArc(cID, exists[key], riskLevels[r][c-1])
				g.AddArc(exists[key], cID, riskLevels[r][c])
			}

			key = fmt.Sprintf("r%dc%d", r, c+1)
			if _, ok := exists[key]; ok {
				g.AddArc(cID, exists[key], riskLevels[r][c+1])
				g.AddArc(exists[key], cID, riskLevels[r][c])
			}
		}
	}

	return sID, eID
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

	numRows := len(lines)
	numCols := len(lines[0])

	risks1 := make([][]int64, numRows)
	for r := 0; r < numRows; r++ {
		risks1[r] = make([]int64, numCols)
	}

	for r, line := range lines {
		for c, ch := range line {
			risks1[r][c], _ = strconv.ParseInt(string(ch), 10, 64)
		}
	}

	risks2 := make([][]int64, numRows*5)
	for r := 0; r < numRows*5; r++ {
		risks2[r] = make([]int64, numCols*5)
	}

	for r2 := 0; r2 < numRows*5; r2++ {
		r1 := r2 % len(risks1)
		rDelta := r2 / len(risks1)

		for c2 := 0; c2 < numCols*5; c2++ {
			c1 := c2 % len(risks1[0])
			cDelta := c2 / len(risks1[0])

			risk := risks1[r1][c1] + int64(rDelta) + int64(cDelta)
			for risk > 9 {
				risk = risk - 9
			}
			risks2[r2][c2] = risk
		}
	}

	graph := dijkstra.NewGraph()
	sID, eID := buildGraph(graph, risks1)

	best, err := graph.Shortest(sID, eID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", best.Distance)

	graph = dijkstra.NewGraph()
	sID, eID = buildGraph(graph, risks2)

	best, err = graph.Shortest(sID, eID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 2: %d\n", best.Distance)
}
