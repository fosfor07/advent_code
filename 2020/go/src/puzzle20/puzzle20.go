package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"matrix"
	"readers"
	"strconv"
	"strings"
	"utils"
)

const tileSize int = 10
const tileSizeWoBorders int = tileSize - 2

type tileType struct {
	tile      [][]rune
	borders   []borderType
	mBorders  int
	r, c      int
	processed bool
}

type borderType struct {
	border string
	dir    rune
}

func updateBorders(t *tileType) {
	M := len(t.tile)
	N := len(t.tile[0])
	b0, b1, b2, b3 := "", "", "", ""

	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			if c == 0 {
				b0 += string(t.tile[r][c])
			} else if c == N-1 {
				b1 += string(t.tile[r][c])
			}
			if r == 0 {
				b2 += string(t.tile[r][c])
			} else if r == M-1 {
				b3 += string(t.tile[r][c])
			}
		}
	}

	t.borders[0] = borderType{border: b0, dir: 'l'}
	t.borders[1] = borderType{border: b1, dir: 'r'}
	t.borders[2] = borderType{border: b2, dir: 'u'}
	t.borders[3] = borderType{border: b3, dir: 'd'}
}

func rotateTile(t *tileType) {
	matrix.RotateCwRune(t.tile)

	for bID, b := range t.borders {
		if b.dir == 'd' {
			t.borders[bID].dir = 'l'
		} else if b.dir == 'u' {
			t.borders[bID].dir = 'r'
		} else if b.dir == 'l' {
			t.borders[bID].dir = 'u'
		} else if b.dir == 'r' {
			t.borders[bID].dir = 'd'
		}
	}
}

func rotateTillBordersMatch(t *tileType, b1ID int, adjTile *tileType, b2ID int) {
	for {
		if t.borders[b1ID].dir == 'u' && adjTile.borders[b2ID].dir == 'd' {
			t.c = adjTile.c
			t.r = adjTile.r + 1
			break
		} else if t.borders[b1ID].dir == 'd' && adjTile.borders[b2ID].dir == 'u' {
			t.c = adjTile.c
			t.r = adjTile.r - 1
			break
		} else if t.borders[b1ID].dir == 'l' && adjTile.borders[b2ID].dir == 'r' {
			t.c = adjTile.c + 1
			t.r = adjTile.r
			break
		} else if t.borders[b1ID].dir == 'r' && adjTile.borders[b2ID].dir == 'l' {
			t.c = adjTile.c - 1
			t.r = adjTile.r
			break
		}
		rotateTile(t)
	}
	updateBorders(t)

	// Get current orientation of the border and check if we need to flip the tile.
	newbID := 0
	if t.borders[b2ID].dir == 'l' {
		newbID = 1
	} else if t.borders[b2ID].dir == 'r' {
		newbID = 0
	} else if t.borders[b2ID].dir == 'u' {
		newbID = 3
	} else if t.borders[b2ID].dir == 'd' {
		newbID = 2
	}

	if t.borders[newbID].border == utils.ReverseString(adjTile.borders[b2ID].border) {
		if t.borders[newbID].dir == 'd' || t.borders[newbID].dir == 'u' {
			matrix.FlipHRune(t.tile)
		} else {
			matrix.FlipVRune(t.tile)
		}
		updateBorders(t)
	}
}

func calcTileCoordinates(tiles map[int]*tileType, tID int, mBordersMap map[int][]int) {
	tiles[tID].processed = true

	for _, mTileId := range mBordersMap[tID] {
		if !tiles[mTileId].processed {
			for b1ID := range tiles[mTileId].borders {
				fnd := false
				for b2ID := range tiles[tID].borders {
					if tiles[mTileId].borders[b1ID].border == tiles[tID].borders[b2ID].border ||
						tiles[mTileId].borders[b1ID].border == utils.ReverseString(tiles[tID].borders[b2ID].border) {
						fnd = true
						rotateTillBordersMatch(tiles[mTileId], b1ID, tiles[tID], b2ID)
						break
					}
				}
				if fnd {
					break
				}
			}
			calcTileCoordinates(tiles, mTileId, mBordersMap)
		}
	}
}

func writeTile(image [][]rune, t *tileType) {
	rs := t.r * tileSizeWoBorders
	cs := t.c * tileSizeWoBorders

	tRow := 1
	for r := rs; r < rs+tileSizeWoBorders; r++ {
		tCol := 1
		for c := cs; c < cs+tileSizeWoBorders; c++ {
			image[r][c] = t.tile[tRow][tCol]
			tCol++
		}
		tRow++
	}
}

func countMonsters(image [][]rune) int {
	seaMonsters := 0

	for r := 0; r < len(image); r++ {
		for c := 0; c < len(image); c++ {
			if c > 0 && c < len(image)-19 && r > 1 {
				if image[r][c] == '#' && image[r][c+3] == '#' && image[r][c+6] == '#' &&
					image[r][c+9] == '#' && image[r][c+12] == '#' && image[r][c+15] == '#' &&
					image[r-1][c-1] == '#' && image[r-1][c+4] == '#' && image[r-1][c+5] == '#' &&
					image[r-1][c+10] == '#' && image[r-1][c+11] == '#' && image[r-1][c+16] == '#' &&
					image[r-1][c+17] == '#' && image[r-1][c+18] == '#' && image[r-2][c+17] == '#' {
					seaMonsters++
				}
			}
		}
	}

	return seaMonsters
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

	tiles := make(map[int]*tileType)
	tileID, row := 0, 0

	brdr1, brdr2, brdr3, brdr4 := "", "", "", ""

	for _, line := range lines {
		if len(line) <= 0 {
			continue
		}

		if line[0] == 'T' {
			fields := strings.Split(line, " ")
			fields[1] = strings.Trim(fields[1], ":")
			tileID, _ = strconv.Atoi(fields[1])

			t := tileType{tile: make([][]rune, tileSize), mBorders: 0, c: 0, r: 0, processed: false}
			for r := 0; r < tileSize; r++ {
				t.tile[r] = make([]rune, tileSize)
			}
			tiles[tileID] = &t
			row = 0
		} else {
			for col, ch := range line {
				tiles[tileID].tile[row][col] = ch

				if col == 0 {
					brdr1 += string(ch)
				} else if col == tileSize-1 {
					brdr2 += string(ch)
				}
				if row == 0 {
					brdr3 += string(ch)
				} else if row == tileSize-1 {
					brdr4 += string(ch)
				}
			}
			row++
		}

		if len(brdr4) == tileSize {
			b1 := borderType{border: brdr1, dir: 'l'}
			b2 := borderType{border: brdr2, dir: 'r'}
			b3 := borderType{border: brdr3, dir: 'u'}
			b4 := borderType{border: brdr4, dir: 'd'}
			tiles[tileID].borders = append(tiles[tileID].borders, b1, b2, b3, b4)
			brdr1, brdr2, brdr3, brdr4 = "", "", "", ""
		}
	}

	imageSize := int(math.Sqrt(float64(len(tiles)))) * (tileSize - 2)
	image := make([][]rune, imageSize)
	for r := 0; r < imageSize; r++ {
		image[r] = make([]rune, imageSize)
	}

	mBordersMap := make(map[int][]int)
	for id1, t1 := range tiles {
		for id2, t2 := range tiles {
			if id1 == id2 {
				continue
			}

			for _, b1 := range t1.borders {
				fnd := false
				for _, b2 := range t2.borders {
					if b1.border == b2.border || b1.border == utils.ReverseString(b2.border) {
						tiles[id1].mBorders++
						mBordersMap[id1] = append(mBordersMap[id1], id2)
						fnd = true
					}
				}
				if fnd {
					break
				}
			}
		}
	}

	calcTileCoordinates(tiles, tileID, mBordersMap)

	part1, part2 := 1, 0
	minC, minR := 0, 0
	for tID, t := range tiles {
		if t.mBorders == 2 {
			part1 *= tID
		}

		if t.c < minC {
			minC = t.c
		}
		if t.r < minR {
			minR = t.r
		}
	}
	for tID := range tiles {
		// Column and row numbers must be >= 0.
		tiles[tID].c -= minC
		tiles[tID].r -= minR
		writeTile(image, tiles[tID])
	}

	hashNum := 0
	for r := 0; r < imageSize; r++ {
		for c := 0; c < imageSize; c++ {
			if image[r][c] == '#' {
				hashNum++
			}
		}
	}

	seaMonsters := 0
	for {
		for i := 0; i < 4; i++ {
			seaMonsters = countMonsters(image)
			if seaMonsters != 0 {
				break
			}
			matrix.RotateCwRune(image)
		}

		if seaMonsters == 0 {
			matrix.FlipHRune(image)
		} else {
			break
		}
	}

	part2 = hashNum - (seaMonsters * 15)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n(Number of sea monsters: %d)\n", part2, seaMonsters)
}
