package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"sync"
)

type dir uint

const (
	north dir = 1 << iota
	south
	east
	west
)

type tile struct {
	val   byte
	state dir
}

type beam struct {
	index
	dir
}

type index struct {
	i, j int
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid := parse(bytes.Fields(b))
	fmt.Println("part 1: ", startAt(beam{dir: east}, clone(grid)))
	fmt.Println("part 2: ", part2(grid))
}

func startAt(b beam, r [][]tile) int {
	current := []beam{b}
	var count int
	for len(current) > 0 {
		var next []beam
		for i := range current {
			if r[current[i].i][current[i].j].state == 0 {
				count++
			}
			nextDirs := visit(current[i], r)
			for j := range nextDirs {
				nextBeam := step(current[i].index, nextDirs[j])
				if nextBeam.i >= 0 && nextBeam.j >= 0 && nextBeam.i < len(r) && nextBeam.j < len(r[0]) {
					next = append(next, nextBeam)
				}
			}
		}
		current = next
	}
	return count
}

func part2(grid [][]tile) int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for j := range grid[0] {
			out <- startAt(beam{index{0, j}, south}, clone(grid))
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for j := range grid[0] {
			out <- startAt(beam{index{len(grid) - 1, j}, north}, clone(grid))
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := range grid {
			out <- startAt(beam{index{i, 0}, east}, clone(grid))
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := range grid {
			out <- startAt(beam{index{i, len(grid[0]) - 1}, west}, clone(grid))
		}
		wg.Done()
		wg.Wait()
		close(out)
	}()
	var m int
	for i := range out {
		m = max(m, i)
	}
	return m
}

func parse(grid [][]byte) [][]tile {
	out := make([][]tile, len(grid))
	for i, row := range grid {
		r := make([]tile, len(row))
		for j := range row {
			r[j] = tile{row[j], 0}
		}
		out[i] = r
	}
	return out
}

// Returns the next directions for the beam to take; returns nil if the beam is out of bounds or the
// tile has already been visited by a beam from the same direction as b.
func visit(b beam, r [][]tile) []dir {
	if b.i < 0 || b.j < 0 || b.i >= len(r) || b.j >= len(r[0]) {
		return nil
	}
	if r[b.i][b.j].state&b.dir != 0 {
		return nil
	}
	r[b.i][b.j].state |= b.dir
	return nextDir(r[b.i][b.j].val, b.dir)
}

// Returns a copy of r and its subslices.
func clone(r [][]tile) [][]tile {
	out := make([][]tile, len(r))
	for i := range r {
		out[i] = slices.Clone(r[i])
	}
	return out
}

// Returns the beam resulting from taking one step from t in the direction d.
func step(t index, d dir) beam {
	return beam{index{t.i - int(d&north) + int((d&south)>>1), t.j - int((d&west)>>3) + int((d&east))>>2}, d}
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
