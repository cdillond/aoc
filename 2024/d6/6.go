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

	visited = 1 << 7
)

type point struct{ i, j int }

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

func walkCycle(m map[point]byte, i, j int, dir byte, grid [][]byte) bool {
	clear(m)
	var p point
	var h byte
	for {
		switch dir {
		case north:
			if i == 0 {
				return false
			}
			if grid[i-1][j] == '#' {
				p.i, p.j = i-1, j
				h = m[p]
				if h&north == north {
					return true
				}
				m[p] = h | north
				dir = east
				continue
			}
			i--
		case south:
			if i == len(grid)-1 {
				return false
			}
			if grid[i+1][j] == '#' {
				p.i, p.j = i+1, j
				h = m[p]
				if h&south == south {
					return true
				}
				m[p] = h | south
				dir = west
				continue
			}
			i++
		case east:
			if j == len(grid[i])-1 {
				return false
			}
			if grid[i][j+1] == '#' {
				p.i, p.j = i, j+1
				h = m[p]
				if h&east == east {
					return true
				}
				m[p] = h | east
				dir = south
				continue
			}
			j++
		case west:
			if j == 0 {
				return false
			}
			if grid[i][j-1] == '#' {
				p.i, p.j = i, j-1
				h = m[p]
				if h&west == west {
					return true
				}
				m[p] = h | west
				dir = north
				continue
			}
			j--
		}
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
	b[n] |= visited
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
	m := make(map[point]byte)
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			c := grid[i][j]
			if c == '#' {
				continue
			}
			// no need to obstruct squares not visited in the initial walk
			if c&visited == 0 {
				continue
			}

			// try placing a hash in the path
			grid[i][j] = '#'

			if walkCycle(m, startI, startJ, north, grid) {
				count++
			}
			// restore original grid value
			grid[i][j] = c
		}
	}

	return aoc.Itoa(count), nil
}
