package d14

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"os"

	"github.com/cdillond/aoc"
)

//go:embed input.txt
var input []byte

const (
	Day  = "14"
	Year = "2024"
)

type robot struct {
	x, y   int
	vx, vy int
}

const (
	height = 103
	width  = 101
)

var ErrParse = errors.New("error parsing input")

func (r *robot) parse(b []byte) error {
	b = b[len("p="):]
	var n int
	if n = bytes.IndexByte(b, ','); n < 0 {
		return ErrParse
	}
	r.x = aoc.Atoi(b[:n])
	b = b[n+1:]
	if n = bytes.IndexByte(b, ' '); n < 0 {
		return ErrParse
	}
	r.y = aoc.Atoi(b[:n])
	b = b[n+1+len("v="):]
	if n = bytes.IndexByte(b, ','); n < 0 {
		return ErrParse
	}
	r.vx = aoc.Atoi(b[:n])
	b = b[n+1:]
	r.vy = aoc.Atoi(b)
	return nil
}

func (r *robot) advance(n int) {
	r.x += n * r.vx
	r.x %= width
	r.y += n * r.vy
	r.y %= height
	if r.x < 0 {
		r.x += width
	}
	if r.y < 0 {
		r.y += height
	}
}

func (r robot) quad() int {
	switch {
	case r.x < width/2 && r.y < height/2:
		return 1
	case r.x > width/2 && r.y < height/2:
		return 2
	case r.x < width/2 && r.y > height/2:
		return 3
	case r.x > width/2 && r.y > height/2:
		return 4
	}
	return 0
}

func Part1(_ string) (res string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var counts [5]int
	var r robot
	for scanner.Scan() {
		if err = r.parse(scanner.Bytes()); err != nil {
			return res, err
		}
		r.advance(100)
		counts[r.quad()]++
	}

	total := 1
	for i := 1; i < len(counts); i++ {
		total *= counts[i]
	}
	return aoc.Itoa(total), nil
}

func printGrid(robots []robot, h, w int) {
	buf := make([]byte, w+1)
	buf[w] = '\n'
	for i := range h {
		for j := range w {
			var count int
			for _, r := range robots {
				if r.x == j && r.y == i {
					count++
				}
			}
			if count == 0 {
				buf[j] = '.'
			} else {
				buf[j] = byte(count) + '0'
			}
		}
		os.Stdout.Write(buf)
	}
}

type point struct{ x, y int }

func Part2(_ string) (res string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var robots []robot
	for scanner.Scan() {
		var r robot
		if err = r.parse(scanner.Bytes()); err != nil {
			return res, err
		}
		robots = append(robots, r)
	}

	cache := make(map[point]struct{})
	var i int
	for ; len(cache) != len(robots); i++ {
		clear(cache)
		for j := range robots {
			robots[j].advance(1)
			cache[point{robots[j].x, robots[j].y}] = struct{}{}
		}
	}
	printGrid(robots, height, width)
	return aoc.Itoa(i), nil
}
