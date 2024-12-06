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
	north = 1 << iota
	east
	south
	west

	visited = 1 << 7
)

// replace all periods with '0' bytes
func clean(b []byte) {
	for i, c := range b {
		if c == '.' {
			b[i] = '\x00'
		}
	}
}

// remove direction information from all visited bytes
func reset(b []byte) {
	for i, c := range b {
		if c != '#' && c != '\n' {
			b[i] &= visited
		}
	}
}

type guard struct {
	dir  byte
	i, j int
}

func (g *guard) next(grid [][]byte) (more, cycle bool) {
start:
	switch g.dir {
	case north:
		if g.i == 0 {
			return false, false
		}
		if grid[g.i-1][g.j] == '#' {
			g.dir = east
			goto start
		}
		g.i--
	case south:
		if g.i == len(grid)-1 {
			return false, false
		}
		if grid[g.i+1][g.j] == '#' {
			g.dir = west
			goto start
		}
		g.i++
	case east:
		if g.j == len(grid[g.i])-1 {
			return false, false
		}
		if grid[g.i][g.j+1] == '#' {
			g.dir = south
			goto start
		}
		g.j++
	case west:
		if g.j == 0 {
			return false, false
		}
		if grid[g.i][g.j-1] == '#' {
			g.dir = north
			goto start
		}
		g.j--
	}
	c := visited | byte(g.dir)
	if grid[g.i][g.j]&c == c {
		return false, true
	}
	grid[g.i][g.j] |= c
	return true, false
}

func Part1(path string) (res string, err error) {
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return res, err
	}
	if b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	clean(b)
	n := bytes.IndexByte(b, '^')
	b[n] = visited
	grid := bytes.Split(b, []byte{'\n'})
	pos := guard{dir: north, i: n / len(grid[0]), j: n % (len(grid[0]) + 1)}
	for more, _ := pos.next(grid); more; more, _ = pos.next(grid) {
	}

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
	clean(b)
	n := bytes.IndexByte(b, '^')
	b[n] = visited
	grid := bytes.Split(b, []byte{'\n'})
	startI, startJ := n/len(grid[0]), n%(len(grid[0])+1)
	pos := guard{dir: north, i: startI, j: startJ}

	// visit everywhere possible
	for more, _ := pos.next(grid); more; more, _ = pos.next(grid) {
	}
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
			reset(b)
			// now walk the path again, and report if it ends in a cycle
			pos = guard{dir: north, i: startI, j: startJ}
			var more, cycle bool
			for more, cycle = pos.next(grid); more; more, cycle = pos.next(grid) {
			}
			if cycle {
				count++
			}
			// restore the former value
			grid[i][j] = c
		}
	}

	return aoc.Itoa(count), nil
}
