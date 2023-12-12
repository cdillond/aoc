package main

import (
	"bytes"
	"fmt"
	"os"
)

type galaxy struct {
	i, j int
}

type pair struct {
	a, b galaxy
}

func (p pair) distance() int {
	dY := p.b.i - p.a.i
	dX := p.b.j - p.a.j

	if dY < 0 {
		dY = -dY
	}
	if dX < 0 {
		dX = -dX
	}
	return dX + dY
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := bytes.Fields(b)

	emptyR := make([]bool, len(rows))    // records whether the i'th row of rows is empty
	emptyC := make([]bool, len(rows[0])) // records whether j'th col of rows is empty
	counts := make([]int, len(rows[0]))
	for i, row := range rows {
		var full bool
		for j := range row {
			if row[j] == '.' {
				counts[j]++
			} else {
				full = true
			}
		}
		if !full {
			emptyR[i] = true
		}
	}
	for i := range counts {
		emptyC[i] = counts[i] == len(rows)
	}

	fmt.Println("part 1: ", solve(rows, 2, emptyR, emptyC))
	fmt.Println("part 2: ", solve(rows, 1_000_000, emptyR, emptyC))
}

func solve(rows [][]byte, adj int, emptyR, emptyC []bool) int {
	// the adjustment factor is always 1 more than the number of rows or cols that need to be added
	galaxies := findGalaxies(rows, adj-1, emptyR, emptyC)
	pairs := toPairs(galaxies)
	var sum int
	for _, p := range pairs {
		sum += p.distance()
	}
	return sum
}

// Returns a slice of all unique galaxy pairings.
func toPairs(g []galaxy) []pair {
	var out []pair
	for i := 0; i < len(g)-1; i++ {
		for j := i + 1; j < len(g); j++ {
			out = append(out, pair{g[i], g[j]})
		}
	}
	return out
}

// Returns a slice of galaxies, whose indices are adjusted by adj based on the number of empty
// rows and columns that precede them.
func findGalaxies(rows [][]byte, adj int, emptyR, emptyC []bool) []galaxy {
	var out []galaxy
	var iCount int
	for i, row := range rows {
		if emptyR[i] {
			iCount++
		}
		var jCount int
		for j := range row {
			if emptyC[j] {
				jCount++
			}
			if row[j] == '#' {
				out = append(out, galaxy{i + iCount*adj, j + jCount*adj})
			}
		}
	}
	return out
}
