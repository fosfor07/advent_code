package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"readers"
	"strconv"
	"strings"
)

type tileType struct {
	x, y, z int
	flips   int
	color   rune
}

func getAdj(x, y, z int, tiles map[string]tileType) []string {
	var keys []string
	keys = append(keys, strconv.Itoa(x+1)+"_"+strconv.Itoa(y-1)+"_"+strconv.Itoa(z))
	keys = append(keys, strconv.Itoa(x-1)+"_"+strconv.Itoa(y+1)+"_"+strconv.Itoa(z))
	keys = append(keys, strconv.Itoa(x+1)+"_"+strconv.Itoa(y)+"_"+strconv.Itoa(z-1))
	keys = append(keys, strconv.Itoa(x)+"_"+strconv.Itoa(y+1)+"_"+strconv.Itoa(z-1))
	keys = append(keys, strconv.Itoa(x)+"_"+strconv.Itoa(y-1)+"_"+strconv.Itoa(z+1))
	keys = append(keys, strconv.Itoa(x-1)+"_"+strconv.Itoa(y)+"_"+strconv.Itoa(z+1))

	return keys
}

// Returns number of adjacent black tiles and keys of all adjacent tiles.
func cntAdjBlack(key string, tiles map[string]tileType) (int, []string) {
	cnt := 0
	fields := strings.Split(key, "_")
	x, _ := strconv.Atoi(fields[0])
	y, _ := strconv.Atoi(fields[1])
	z, _ := strconv.Atoi(fields[2])

	keys := getAdj(x, y, z, tiles)

	for _, key := range keys {
		if _, ok := tiles[key]; ok {
			if tiles[key].color == 'b' {
				cnt++
			}
		}
	}
	return cnt, keys
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
	tiles := make(map[string]tileType)
	uTiles := make(map[string]tileType)

	// part 1
	for _, line := range lines {
		x, y, z := 0, 0, 0

		for len(line) > 0 {
			dir := ""
			if line[0] == 's' || line[0] == 'n' {
				dir = string(line[0:2])
				line = line[2:]
			} else {
				dir = string(line[0])
				line = line[1:]
			}

			if dir == "e" {
				x++
				y--
			} else if dir == "w" {
				x--
				y++
			} else if dir == "ne" {
				x++
				z--
			} else if dir == "nw" {
				y++
				z--
			} else if dir == "se" {
				y--
				z++
			} else if dir == "sw" {
				x--
				z++
			}
		}

		flips := 0
		key := strconv.Itoa(x) + "_" + strconv.Itoa(y) + "_" + strconv.Itoa(z)
		if _, ok := tiles[key]; ok {
			flips = tiles[key].flips + 1
		} else {
			flips = 1
		}

		var tile tileType
		tile.x, tile.y, tile.z = x, y, z
		tile.flips = flips
		tile.color = 'w'
		tiles[key] = tile
	}

	for k, tile := range tiles {
		var uTile tileType
		uTile.x, uTile.y, uTile.z = tile.x, tile.y, tile.z

		if tile.flips%2 == 1 {
			part1++
			uTile.color = 'b'
		} else {
			uTile.color = 'w'
		}
		tiles[k] = uTile
	}

	// part 2
	for i := 0; i < 100; i++ {
		for k, tile := range tiles {
			b, keys := cntAdjBlack(k, tiles)

			var uTile tileType
			uTile.x, uTile.y, uTile.z = tile.x, tile.y, tile.z

			if tile.color == 'b' {
				if b == 0 || b > 2 {
					uTile.color = 'w'
					uTiles[k] = uTile
				}
			} else if tile.color == 'w' {
				if b == 2 {
					uTile.color = 'b'
					uTiles[k] = uTile
				}
			}

			// Update also adjacent tiles.
			for _, key := range keys {
				b, _ = cntAdjBlack(key, tiles)

				fields := strings.Split(key, "_")
				kx, _ := strconv.Atoi(fields[0])
				ky, _ := strconv.Atoi(fields[1])
				kz, _ := strconv.Atoi(fields[2])

				uTile.x, uTile.y, uTile.z = kx, ky, kz
				if b == 2 {
					uTile.color = 'b'
					uTiles[key] = uTile
				}
			}
		}

		for k, tile := range uTiles {
			tiles[k] = tile
		}
	}

	for _, tile := range tiles {
		if tile.color == 'b' {
			part2++
		}
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
