package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type entry struct {
	s string
	n int
}

type hmap [256][]entry

func (h *hmap) delete(s string, i uint8) {
	n := -1
	for j := range h[i] {
		if h[i][j].s == s {
			n = j
			break
		}
	}
	if n >= 0 {
		h[i] = append(h[i][:n], h[i][n+1:]...)
	}
}

func (h *hmap) add(e entry, i uint8) {
	n := -1
	for j := range h[i] {
		if h[i][j].s == e.s {
			n = j
			break
		}
	}
	if n >= 0 {
		h[i][n] = e
	} else {
		h[i] = append(h[i], e)
	}
}

func (h *hmap) score() uint64 {
	var sum int
	for i := range h {
		for j := range h[i] {
			sum += (1 + i) * (1 + j) * h[i][j].n
		}
	}
	return uint64(sum)
}

func hash(b []byte) uint8 {
	var cur uint64
	for i := range b {
		t := b[i]
		if t == '\n' || t == '\r' {
			continue
		}
		cur += uint64(t)
		cur *= 17
		cur %= 256
	}
	return uint8(cur)
}

func part1(fields [][]byte) uint64 {
	var sum uint64
	for i := range fields {
		sum += uint64(hash(fields[i]))
	}
	return sum
}

func part2(fields [][]byte) uint64 {
	var m hmap
	for i := range fields {
	loop:
		for j := range fields[i] {
			switch fields[i][j] {
			case '=':
				a, _ := strconv.Atoi(string(fields[i][j+1:]))
				m.add(entry{s: string(fields[i][:j]), n: a}, hash(fields[i][:j]))
				break loop
			case '-':
				m.delete(string(fields[i][:j]), hash(fields[i][:j]))
			}
		}
	}
	return m.score()
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fields := bytes.Split(b, []byte(","))
	fmt.Println("part 1: ", part1(fields))
	fmt.Println("part 2: ", part2(fields))
}
