package main

import (
	"bytes"
	"fmt"
	"os"
)

type dir uint

const (
	north dir = 1 << iota
	south
	east
	west
)

type packet struct {
	tile
	dir
}

type tile struct {
	i, j int
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Fields(b)
	fmt.Println("part 1: ", startAt(tile{0, 0}, east, rows))
	fmt.Println("part 2: ", part2(rows))
}

func startAt(t tile, d dir, r [][]byte) int {
	packets := []packet{{t, d}}
	energized := map[tile]struct{}{
		t: {},
	}
	cache := map[packet]struct{}{
		{t, d}: {},
	}
	for len(packets) > 0 {
		next := []packet{}
		for _, p := range packets {
			nextDirs := nextDir(r[p.i][p.j], p.dir)
			for _, d := range nextDirs {
				nextP := step(p.tile, d)
				if inBounds(nextP, r) {
					if _, seen := cache[nextP]; !seen {
						energized[nextP.tile] = struct{}{}
						cache[nextP] = struct{}{}
						next = append(next, nextP)
					}
				}
			}
		}
		packets = next
	}
	return len(energized)
}

func part2(rows [][]byte) int {
	outcomes := make([]int, 0, 4*len(rows))
	for j := range rows[0] {
		top := startAt(tile{0, j}, south, rows)
		bottom := startAt(tile{len(rows) - 1, j}, north, rows)
		outcomes = append(outcomes, top, bottom)
	}
	for i := range rows {
		l := startAt(tile{i, 0}, east, rows)
		r := startAt(tile{i, len(rows[0]) - 1}, west, rows)
		outcomes = append(outcomes, l, r)
	}
	var m int
	for i := range outcomes {
		m = max(m, outcomes[i])
	}
	return m
}

// Returns the packet resulting from taking one step from t in the direction d.
func step(t tile, d dir) packet {
	return packet{tile{t.i - int(d&north) + int((d&south)>>1), t.j - int((d&west)>>3) + int((d&east))>>2}, d}
}

// Returns the slice of directions resulting from the application of the operand b to the direction d.
func nextDir(b byte, d dir) []dir {
	switch b {
	case '.':
		return []dir{d}
	case '|':
		if d >= east {
			return []dir{north, south}
		}
		return []dir{d}
	case '-':
		if d <= south {
			return []dir{west, east}
		}
		return []dir{d}
	case '/':
		if d <= south {
			return []dir{d << 2} // north -> east; south -> west
		}
		return []dir{d >> 2} // east -> north; west -> south
	case '\\':
		if d == north || d == west {
			return []dir{(d&north)<<3 + (d&west)>>3} // north -> west; west -> north
		}
		return []dir{(d&south)<<1 + (d&east)>>1} // south -> east; east -> south
	}
	return nil
}

func inBounds(p packet, r [][]byte) bool {
	return p.i >= 0 && p.j >= 0 && p.i < len(r) && p.j < len(r[0])
}
