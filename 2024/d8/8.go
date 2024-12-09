package d8

import (
	"bufio"
	"os"

	"github.com/cdillond/aoc"
)

const (
	Day  = "8"
	Year = "2024"
)

type point struct{ i, j int }

func dydx(p1, p2 point) (int, int) { return p1.i - p2.i, p1.j - p2.j }

func antinodes(p0, p1 point) (point, point) {
	dy0, dx0 := dydx(p0, p1)
	dy1, dx1 := dydx(p1, p0)
	p0.i += dy0
	p0.j += dx0
	p1.i += dy1
	p1.j += dx1
	return p0, p1
}

func (p point) isInBounds(iMax, jMax int) bool {
	return p.i > -1 && p.i < iMax && p.j > -1 && p.j < jMax
}

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	var (
		scanner = bufio.NewScanner(f)
		m       = make(map[byte][]point)
		i, jMax int
		p       point
		seen    struct{}
	)

	for scanner.Scan() {
		b := scanner.Bytes()
		jMax = max(jMax, len(b))
		for j, c := range b {
			if c != '.' {
				p.i, p.j = i, j
				s := append(m[c], p)
				m[c] = s
			}
		}
		i++
	}

	set := make(map[point]struct{})
	for _, val := range m {
		for j := 0; j < len(val)-1; j++ {
			for k := j + 1; k < len(val); k++ {
				a1, a2 := antinodes(val[j], val[k])
				if a1.isInBounds(i, jMax) {
					set[a1] = seen
				}
				if a2.isInBounds(i, jMax) {
					set[a2] = seen
				}
			}
		}
	}

	return aoc.Itoa(len(set)), nil
}

func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	var (
		scanner = bufio.NewScanner(f)
		m       = make(map[byte][]point)
		i, jMax int
		p       point
		seen    struct{}
	)

	for scanner.Scan() {
		b := scanner.Bytes()
		jMax = max(jMax, len(b))
		for j, c := range b {
			if c != '.' {
				p.i, p.j = i, j
				s := append(m[c], p)
				m[c] = s
			}
		}
		i++
	}

	set := make(map[point]struct{})
	for _, antennae := range m {
		for j := 0; j < len(antennae)-1; j++ {
			for k := j + 1; k < len(antennae); k++ {
				dy, dx := dydx(antennae[j], antennae[k])

				for p = antennae[j]; p.isInBounds(i, jMax); p.i, p.j = p.i+dy, p.j+dx {
					set[p] = seen
				}

				dy, dx = dydx(antennae[k], antennae[j])
				for p = antennae[k]; p.isInBounds(i, jMax); p.i, p.j = p.i+dy, p.j+dx {
					set[p] = seen
				}
			}
		}
	}

	return aoc.Itoa(len(set)), nil
}
