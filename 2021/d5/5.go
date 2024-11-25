package d5

import (
	"aoc"
	"bufio"
	"bytes"
	"os"
)

const (
	Day  = "5"
	Year = "2021"
)

type point struct{ x, y int }

type line struct{ a, b point }

func (l line) notDiagonal() bool  { return l.a.x == l.b.x || l.a.y == l.b.y }
func (l line) isHorizontal() bool { return l.a.y == l.b.y }
func (l line) isVertical() bool   { return l.a.x == l.b.x }

func parseLine(b []byte) line {
	before, after, _ := bytes.Cut(b, []byte(" -> "))

	var l line
	n := bytes.IndexByte(before, ',')
	l.a.x = aoc.Atoi(before[:n])
	l.a.y = aoc.Atoi(before[n+1:])

	n = bytes.IndexByte(after, ',')
	l.b.x = aoc.Atoi(after[:n])
	l.b.y = aoc.Atoi(after[n+1:])

	return l
}

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()
	var maxX, maxY int
	scanner := bufio.NewScanner(f)
	var lines []line
	for scanner.Scan() {
		l := parseLine(scanner.Bytes())
		if l.notDiagonal() {
			lines = append(lines, l)
			maxX = max(maxX, max(l.a.x, l.b.x))
			maxY = max(maxY, max(l.a.y, l.b.y))
		}

	}

	grid := make([][]int, maxY+1)
	for i := range grid {
		grid[i] = make([]int, maxX+1)
	}

	for _, l := range lines {
		if l.isHorizontal() {
			row := grid[l.a.y]
			for i := min(l.a.x, l.b.x); i <= max(l.a.x, l.b.x); i++ {
				row[i]++
			}
		} else {
			col := l.a.x
			for i := min(l.a.y, l.b.y); i <= max(l.a.y, l.b.y); i++ {
				grid[i][col]++
			}
		}
	}

	var count int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] >= 2 {
				count++
			}
		}
	}

	return aoc.Itoa(count), nil
}

func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()
	var maxX, maxY int
	scanner := bufio.NewScanner(f)
	var lines []line
	for scanner.Scan() {
		l := parseLine(scanner.Bytes())
		lines = append(lines, l)
		maxX = max(maxX, max(l.a.x, l.b.x))
		maxY = max(maxY, max(l.a.y, l.b.y))
	}

	grid := make([][]int, maxY+1)
	for i := range grid {
		grid[i] = make([]int, maxX+1)
	}

	leq := func(a, b int) bool { return a <= b }
	geq := func(a, b int) bool { return a >= b }

	for _, l := range lines {
		if l.isHorizontal() {
			row := grid[l.a.y]
			for i := min(l.a.x, l.b.x); i <= max(l.a.x, l.b.x); i++ {
				row[i]++
			}
		} else if l.isVertical() {
			col := l.a.x
			for i := min(l.a.y, l.b.y); i <= max(l.a.y, l.b.y); i++ {
				grid[i][col]++
			}
		} else {
			var xDelta, yDelta int
			var cond func(a, b int) bool
			switch {
			// x increases and y increases
			case l.a.x < l.b.x && l.a.y < l.b.y:
				xDelta, yDelta = 1, 1
				cond = leq
			// x increases and y decreases
			case l.a.x < l.b.x && l.a.y > l.b.y:
				xDelta, yDelta = 1, -1
				cond = geq
			// x decreases and y decreases
			case l.a.x > l.b.x && l.a.y > l.b.y:
				xDelta, yDelta = -1, -1
				cond = geq
			// x decreases and y increases
			case l.a.x > l.b.x && l.a.y < l.b.y:
				xDelta, yDelta = -1, 1
				cond = leq
			}

			x, y := l.a.x, l.a.y
			for cond(y, l.b.y) {
				grid[y][x]++
				x += xDelta
				y += yDelta
			}

		}
	}

	var count int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] >= 2 {
				count++
			}
		}
	}

	return aoc.Itoa(count), nil
}
