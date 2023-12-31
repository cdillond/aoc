package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type entry struct {
	label string
	n     int
}

type hmap [256][]entry

func (h *hmap) delete(label string, i uint8) {
	for j := range h[i] {
		if h[i][j].label == label {
			h[i] = append(h[i][:j], h[i][j+1:]...)
			return
		}
	}
}

func (h *hmap) add(e entry, i uint8) {
	for j := range h[i] {
		if h[i][j].label == e.label {
			h[i][j] = e
			return
		}
	}
	h[i] = append(h[i], e)
}

func (h *hmap) score() int {
	var sum int
	for i := range h {
		for j := range h[i] {
			sum += (1 + i) * (1 + j) * h[i][j].n
		}
	}
	return sum
}

func hash(b []byte) uint8 {
	var cur uint8
	for i := range b {
		t := b[i]
		if t == '\n' || t == '\r' {
			continue
		}
		cur += t
		cur *= 17
		//cur %= 256 can be omitted since the return value is a uint8
	}
	return cur
}

func part1(fields [][]byte) uint64 {
	var sum uint64
	for i := range fields {
		sum += uint64(hash(fields[i]))
	}
	return sum
}

func part2(fields [][]byte) int {
	var m hmap
	for i := range fields {
	loop:
		for j := range fields[i] {
			switch fields[i][j] {
			case '=':
				a, _ := strconv.Atoi(string(fields[i][j+1:]))
				m.add(entry{label: string(fields[i][:j]), n: a}, hash(fields[i][:j]))
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
