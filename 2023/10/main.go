package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

type offset struct {
	i, j int
}

type direction uint

const (
	stop direction = iota
	north
	east
	south
	west
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Fields(b)
	start := findStart(rows)

	// walk the pipes, in all four directions from the start point
	n := WalkPipes(step(north, start), north, rows, []offset{}, make(map[offset]struct{}))
	e := WalkPipes(step(east, start), east, rows, []offset{}, make(map[offset]struct{}))
	s := WalkPipes(step(south, start), south, rows, []offset{}, make(map[offset]struct{}))
	w := WalkPipes(step(west, start), west, rows, []offset{}, make(map[offset]struct{}))

	// choose the longest loop; the two longest will be redundant, as will the two shortest
	res := [][]offset{n, e, s, w}
	sort.Slice(res, func(i, j int) bool { return len(res[i]) > len(res[j]) })

	fmt.Println("part 1: ", (len(res[0])+1)/2)

	// re-calculate the set of all (i,j) positions in the loop
	seen := make(map[offset]struct{}, len(res[0]))
	for i := range res[0] {
		seen[res[0][i]] = struct{}{}
	}

	// replace the 'S' with its actual, inferred type
	rows[start.i][start.j] = findStartType(res[0])

	fmt.Println("part 2: ", countInterior(rows, seen))

}

// Returns the direction of the next symbol given the current symbol and direction.
func next(c byte, d direction) direction {
	switch c {
	case '|':
		return [5]direction{0, north, 0, south, 0}[d]
	case '-':
		return [5]direction{0, 0, east, 0, west}[d]
	case 'L':
		return [5]direction{0, 0, 0, east, north}[d]
	case 'J':
		return [5]direction{0, 0, north, west, 0}[d]
	case '7':
		return [5]direction{0, west, south, 0, 0}[d]
	case 'F':
		return [5]direction{0, east, 0, 0, south}[d]
	default:
		return stop
	}
}

// Returns the offset of the next symbol given the target direction and the current offset.
func step(d direction, o offset) offset {
	switch d {
	case north:
		return offset{o.i - 1, o.j}
	case east:
		return offset{o.i, o.j + 1}
	case south:
		return offset{o.i + 1, o.j}
	case west:
		return offset{o.i, o.j - 1}
	default:
		return o
	}
}

// Returns the offset of the start symbol.
func findStart(rows [][]byte) offset {
	for i, row := range rows {
		for j := range row {
			if row[j] == 'S' {
				return offset{i, j}
			}
		}
	}
	return offset{}
}

// Walks along the pipes and returns when a cycle is detected or the pipes lead out of bounds or to a '.' tile.
func WalkPipes(start offset, d direction, pipes [][]byte, path []offset, seen map[offset]struct{}) []offset {
	if start.i < 0 || start.j < 0 || start.i >= len(pipes) || start.j >= len(pipes[start.i]) {
		return path
	}
	if _, ok := seen[start]; ok {
		return path
	}
	seen[start] = struct{}{}
	path = append(path, start)
	p := pipes[start.i][start.j]
	nextDir := next(p, d)
	if nextDir == stop {
		return path
	}
	nextOff := step(nextDir, start)
	return WalkPipes(nextOff, nextDir, pipes, path, seen)
}

// Infers the type of pipe 'S' is from the pipes directly before and after it in the loop.
func findStartType(path []offset) byte {
	first := path[0]
	last := path[len(path)-2] // the final node is the start node, so the second to last node is actually last

	k := offset{first.i - last.i, first.j - last.j}
	switch k {
	case offset{2, 0}, offset{-2, 0}:
		return '|'
	case offset{0, 2}, offset{0, -2}:
		return '-'
	case offset{-1, -1}:
		return 'L'
	case offset{-1, 1}:
		return 'J'
	case offset{1, 1}:
		return '7'
	case offset{-1, 1}:
		return 'F'
	}
	return 0
}

func countInterior(rows [][]byte, loop map[offset]struct{}) int {
	// walk each row and track the number of intersections with an edge.
	// odd numbered intersections are outside the loop, even ones are inside; this is just tracked via a bool.
	// based on the current state, either count or ignore the tiles/pipes not in the loop.
	var count int
	// any indices on the exterior of the graph, by default, cannot be counted as 'interior'
	for i := 1; i < len(rows)-1; i++ {
		var in bool
		// | pipes are always vertical edges
		// L*7 and F*J patterns connect to form, in effect, a single vertical edge
		var lastVertex byte
		for j := 0; j < len(rows[i])-1; j++ {

			if _, ok := loop[offset{i, j}]; ok {
				switch rows[i][j] {
				case '|':
					in = !in
					lastVertex = 0
				case 'L', 'F':
					lastVertex = rows[i][j]
				case 'J':
					if lastVertex == 'F' {
						in = !in
					}
					lastVertex = 0
				case '7':
					if lastVertex == 'L' {
						in = !in

					}
					lastVertex = 0
				}
			} else {
				lastVertex = 0
				if in {
					count++
				}
			}

		}
	}
	return count
}
