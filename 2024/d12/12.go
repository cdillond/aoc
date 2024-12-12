package d12

import (
	"bytes"
	_ "embed"

	"github.com/cdillond/aoc"
)

//go:embed input.txt
var input []byte

const (
	Day  = "12"
	Year = "2024"
)

type position int

const (
	upper position = iota
	lower
	left
	right
)

type edge struct {
	position
	start, end point
}

// if the points are adjacent and the edges have the same orientation, it's possible to merge the edges.
func (e *edge) merge(e1 edge) bool {
	if e.position != e1.position {
		return false
	}
	switch e.position {
	case upper, lower:
		if e.start.i != e1.start.i {
			return false
		}
		if e1.start.j-e.end.j == 1 {
			e.end.j = e1.end.j
			return true
		} else if e.start.j-e1.end.j == 1 {
			e.start.j = e1.start.j
			return true
		}
	case left, right:
		if e.start.j != e1.start.j {
			return false
		}
		if e1.start.i-e.end.i == 1 {
			e.end.i = e1.end.i
			return true
		} else if e.start.i-e1.end.i == 1 {
			e.start.i = e1.start.i
			return true
		}
	}
	return false
}

type point struct{ i, j int }

func (p point) inBounds(grid [][]byte) bool {
	return p.i >= 0 && p.i < len(grid) && p.j >= 0 && p.j < len(grid[p.i])
}

func (p point) perim(grid [][]byte) (out int) {
	c := grid[p.i][p.j]
	dirs := [...][2]int{
		upper: {-1, 0},
		lower: {1, 0},
		left:  {0, -1},
		right: {0, 1},
	}

	for _, m := range dirs {
		t := point{p.i + m[0], p.j + m[1]}
		if !t.inBounds(grid) {
			out++
			continue
		}
		if grid[t.i][t.j] != c {
			out++
		}
	}

	return out

}

func (p point) appendEdges(dst []edge, grid [][]byte) []edge {
	c := grid[p.i][p.j]
	// upper
	if (p.i == 0) || (grid[p.i-1][p.j] != c) {
		dst = append(dst, edge{position: upper, start: p, end: p})
	}
	// lower
	if (p.i == len(grid)-1) || (grid[p.i+1][p.j] != c) {
		dst = append(dst, edge{position: lower, start: p, end: p})
	}
	// left
	if (p.j == 0) || (grid[p.i][p.j-1] != c) {
		dst = append(dst, edge{position: left, start: p, end: p})
	}
	// right
	if (p.j == len(grid[p.i])-1) || (grid[p.i][p.j+1] != c) {
		dst = append(dst, edge{position: right, start: p, end: p})
	}
	return dst
}

func Part1(_ string) (res string, err error) {
	dirs := [...][2]int{
		upper: {-1, 0},
		lower: {1, 0},
		left:  {0, -1},
		right: {0, 1},
	}
	grid := bytes.Split(input, []byte{'\n'})
	var total int
	seen := make(map[point]struct{})
	var cur, next []point
	for i, row := range grid {
		for j, c := range row {
			if _, ok := seen[point{i, j}]; ok {
				continue
			}
			// do a BFS
			region := make(map[point]struct{})
			region[point{i, j}] = struct{}{}
			cur = append(cur[:0], point{i, j})
			var perim int
			for len(cur) > 0 {
				for _, p := range cur {
					perim += p.perim(grid)
					for _, m := range dirs {
						t := point{p.i + m[0], p.j + m[1]}
						if _, ok := region[t]; !ok {
							if !t.inBounds(grid) {
								continue
							}
							if grid[t.i][t.j] == c {
								seen[t] = struct{}{}
								region[t] = struct{}{}
								next = append(next, t)
							}
						}
					}
				}

				cur, next = next, cur
				next = next[:0]
			}
			area := len(region)
			total += area * perim
		}
	}
	return aoc.Itoa(total), nil
}

func Part2(_ string) (res string, err error) {
	grid := bytes.Split(input, []byte{'\n'})
	var total int
	seen := make(map[point]struct{})
	var cur, next []point
	var edges []edge
	for i, row := range grid {
		for j, c := range row {
			if _, ok := seen[point{i, j}]; ok {
				continue
			}
			start := point{i, j}

			// do a BFS
			region := make(map[point]struct{})
			region[start] = struct{}{}
			cur = append(cur[:0], start)
			edges = edges[:0]
			for len(cur) > 0 {
				for _, p := range cur {
					edges = p.appendEdges(edges, grid)

					// up
					if p.i > 0 && grid[p.i-1][p.j] == c {
						t := point{p.i - 1, p.j}
						if _, ok := region[t]; !ok {
							seen[t] = struct{}{}
							region[t] = struct{}{}
							next = append(next, t)
						}
					}

					// down
					if p.i < len(grid)-1 && grid[p.i+1][p.j] == c {
						t := point{p.i + 1, p.j}
						if _, ok := region[t]; !ok {
							seen[t] = struct{}{}
							region[t] = struct{}{}
							next = append(next, t)
						}
					}

					// left
					if p.j > 0 && grid[p.i][p.j-1] == c {
						t := point{p.i, p.j - 1}
						if _, ok := region[t]; !ok {
							seen[t] = struct{}{}
							region[t] = struct{}{}
							next = append(next, t)
						}
					}

					// right
					if p.j < len(row)-1 && grid[p.i][p.j+1] == c {
						t := point{p.i, p.j + 1}
						if _, ok := region[t]; !ok {
							seen[t] = struct{}{}
							region[t] = struct{}{}
							next = append(next, t)
						}
					}
				}

				cur, next = next, cur
				next = next[:0]
			}
			area := len(region)

			queue, final := edges, make([]edge, 0, len(edges))
			for len(queue) > 0 {
				curPt := queue[0]
				var merged bool
				for i := 1; i < len(queue); i++ {
					if merged = queue[i].merge(curPt); merged {
						break
					}
				}
				if !merged {
					final = append(final, curPt)
				}
				queue = queue[1:]
			}
			total += area * len(final)
		}
	}
	return aoc.Itoa(total), nil

}
