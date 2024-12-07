package d6

import (
	"bytes"
	"os"

	"github.com/cdillond/aoc"
)

const (
	Day  = "6"
	Year = "2024"
)

const (
	south = 1 << iota
	west
	north
	east
	_
	hash
	_
	visited
)

func b2b(b bool) byte {
	var x byte
	if b {
		x = 1
	}
	return x
}

func turn(dir byte) byte {
	return ((dir << 1) & 0b1111) | ((dir & east) >> 3)
}

func didj(dir byte) (int, int) {
	return int(b2b(dir == south)) - int(b2b(dir == north)), int(b2b(dir == east)) - int(b2b(dir == west))
}

// assumes no cycle is encountered
func walk(i, j int, dir byte, grid [][]byte) {
	var di, dj int
	var ti, tj int
	for {
		di, dj = didj(dir)
		ti, tj = i+di, j+dj
		if ti < 0 || ti >= len(grid) || tj < 0 || tj >= len(grid[0]) {
			return
		}
		if grid[ti][tj] == '#' {
			dir = turn(dir)
			continue
		}
		i = ti
		j = tj
		grid[i][j] |= visited
	}
}

func clean(b []byte) {
	for i := range b {
		if b[i]&visited > 0 {
			b[i] = visited
		} else if b[i] == '#' {
			b[i] = hash
		} else {
			b[i] = 0
		}
	}
}

func patrol(i, j int, dir byte, grid [][]byte) bool {
	var di, dj int
	var ti, tj int
	for {
		di, dj = didj(dir)
		ti, tj = i+di, j+dj
		if ti < 0 || ti >= len(grid) || tj < 0 || tj >= len(grid[0]) {
			return false
		}
		if grid[ti][tj]&hash > 0 {
			if grid[ti][tj]&dir > 0 {
				return true
			} else {
				grid[ti][tj] |= dir
				dir = turn(dir)
				continue
			}
		}
		i = ti
		j = tj
	}
}

func Part1(path string) (res string, err error) {
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return res, err
	}
	if b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	n := bytes.IndexByte(b, '^')
	b[n] = visited
	grid := bytes.Split(b, []byte{'\n'})
	startI, startJ := n/len(grid[0]), n%(len(grid[0])+1)
	walk(startI, startJ, north, grid)

	var c byte
	var count int
	for _, c = range b {
		if c&visited > 0 {
			count++
		}
	}
	return aoc.Itoa(count), nil
}

func Part2(path string) (res string, err error) {
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return res, err
	}
	if b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	n := bytes.IndexByte(b, '^')
	b[n] = visited
	grid := bytes.Split(b, []byte{'\n'})
	startI, startJ := n/len(grid[0]), n%(len(grid[0])+1)

	walk(startI, startJ, north, grid)
	clean(b)
	clone := make([]byte, len(b))
	copy(clone, b)
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] != visited {
				continue
			}

			// try placing a hash in the path
			grid[i][j] = hash
			if patrol(startI, startJ, north, grid) {
				count++
			}
			// restore original grid values
			copy(b, clone)
		}
	}

	return aoc.Itoa(count), nil
}
