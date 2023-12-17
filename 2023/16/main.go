package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

type dir uint

var empty [110][110]dir

const (
	north dir = 1 << iota
	south
	east
	west
)

type beam struct {
	index
	dir
}

type index struct {
	i, j int
}

func main() {
	start := time.Now()
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid := bytes.Fields(b)
	fmt.Println("part 1: ", startAt(beam{dir: east}, grid, new([110][110]dir)))
	fmt.Println("part 2: ", part2(grid), time.Since(start))
}

func startAt(b beam, r [][]byte, hits *[110][110]dir) int {
	current := []beam{b}
	pool := [2][]beam{make([]beam, 0, 64), make([]beam, 0, 64)}
	var count int
	var i int
	for len(current) > 0 {
		next := pool[i%2]
		for _, c := range current {
			if hits[c.i][c.j] == 0 {
				count++
			}
			if !visit(c, hits) {
				continue
			}
			nextDirs := nextDir(r[c.i][c.j], c.dir)
			for j := range nextDirs {
				nextBeam := step(c.index, nextDirs[j])
				if nextBeam.i >= 0 && nextBeam.j >= 0 && nextBeam.i < len(r) && nextBeam.j < len(r[0]) {
					next = append(next, nextBeam)
				}
			}
		}
		pool[(i+1)%2] = pool[(i+1)%2][:0]
		current = next
		i++
	}
	return count
}

func part2(grid [][]byte) int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var h [110][110]dir
		for j := range grid[0] {
			out <- startAt(beam{index{0, j}, south}, grid, &h)
			copy(h[:], empty[:])
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		var h [110][110]dir
		for j := range grid[0] {
			out <- startAt(beam{index{len(grid) - 1, j}, north}, grid, &h)
			copy(h[:], empty[:])
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		var h [110][110]dir
		for i := range grid {
			out <- startAt(beam{index{i, 0}, east}, grid, &h)
			copy(h[:], empty[:])
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		var h [110][110]dir
		for i := range grid {
			out <- startAt(beam{index{i, len(grid[0]) - 1}, west}, grid, &h)
			copy(h[:], empty[:])
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

// Returns true if the index has not been visited yet by an identical beam.
func visit(b beam, h *[110][110]dir) bool {
	if h[b.i][b.j]&b.dir != 0 {
		return false
	}
	h[b.i][b.j] |= b.dir
	return true
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
