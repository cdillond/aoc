package main

import (
	"bytes"
	"fmt"
	"os"
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

	maxPath, maxSet := WalkPipes(start, rows)

	fmt.Println("part 1: ", (len(maxPath)+1)/2)

	// replace the 'S' with its actual, inferred type
	rows[start.i][start.j] = findStartType(maxPath)

	fmt.Println("part 2: ", countInterior(rows, maxSet))

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

// For each direction, WalkPipes walks along the pipes and stops when a cycle is detected or the pipes lead out of bounds or to a '.' tile.
// Returns the path of the longest loop and the set of all offsets in the path.
func WalkPipes(start offset, pipes [][]byte) ([]offset, map[offset]struct{}) {
	var maxPath []offset
	var maxSet map[offset]struct{}
	var maxLen int
	for i := north; i <= west; i++ {
		index := step(i, start)
		d := i
		var path []offset
		seen := make(map[offset]struct{})
		for {
			if index.i < 0 || index.j < 0 || index.i >= len(pipes) || index.j >= len(pipes[index.i]) {
				return path, nil
			}
			if _, ok := seen[index]; ok {
				break
			}
			seen[index] = struct{}{}
			path = append(path, index)
			p := pipes[index.i][index.j]
			nextDir := next(p, d)
			if nextDir == stop {
				break
			}
			d = nextDir
			index = step(nextDir, index)
		}
		if len(path) > maxLen {
			maxLen = len(path)
			maxPath = path
			maxSet = seen
		}
	}
	return maxPath, maxSet
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
