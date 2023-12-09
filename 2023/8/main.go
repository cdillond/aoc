package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type node struct {
	l, r uint
	end  bool
}

func main() {

	f, err := os.Open("nodes.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var nodes []node
	var names []string
	var start, end uint // for part 1
	var starts []uint   // for part 2
	// this is more complicated than it needs to be; using a map and a different data structure for the node
	// would simplify the code, but I was interested in seeing if I could optimize the program a little...
	for scanner.Scan() {
		b := scanner.Bytes()
		a := string(b[:3])
		l := string(b[7:10])
		r := string(b[12:15])

		i := findIndex(a, names)
		if i == -1 {
			nodes = append(nodes, node{})
			names = append(names, a)
			i = len(nodes) - 1
		}
		li := findIndex(l, names)
		if li == -1 {
			nodes = append(nodes, node{})
			names = append(names, l)
			li = len(nodes) - 1
		}
		ri := findIndex(r, names)
		if ri == -1 {
			nodes = append(nodes, node{})
			names = append(names, r)
			ri = len(nodes) - 1
		}
		nodes[i].l = uint(li)
		nodes[i].r = uint(ri)
		if a[2] == 'Z' {
			nodes[i].end = true
		} else if a[2] == 'A' {
			starts = append(starts, uint(i))
		}
		if a == "AAA" {
			start = uint(i)
		} else if a == "ZZZ" {
			end = uint(i)
		}
	}
	inst, err := os.ReadFile("instructions.txt")
	if err != nil {
		panic(err)
	}
	ns := time.Now()
	fmt.Println("part 1: ", Part1(start, end, nodes, inst), time.Since(ns))
	ns = time.Now()
	fmt.Println("part 2: ", Part2(starts, nodes, inst), time.Since(ns))
}

func Part1(start, end uint, nodes []node, b []byte) int {
	n := nodes[start]
	i := 0
	for ; ; i++ {
		switch b[i%len(b)] {
		case 'R':
			if n.r == end {
				return i + 1
			}
			n = nodes[n.r]
		case 'L':
			if n.l == end {
				return i + 1
			}
			n = nodes[n.l]
		}
	}
}

func Part2(starts []uint, nodes []node, b []byte) int {
	// it just so happens that each cycle in the input has only one exit node, each cycle length is
	// evenly divisible by len(b), and everything repeats cleanly, which means the lcm of all
	// cycle lengths will be the correct answer
	nums := make([]int, len(starts))
	for i := 0; i < len(starts); i++ {
		nums[i] = cycle(starts[i], nodes, b)
	}
	return lcm(nums...)
}

func cycle(start uint, nodes []node, b []byte) int {
	n := nodes[start]
	i := 0
	for ; ; i++ {
		if n.end {
			return i
		}
		switch b[i%len(b)] {
		case 'R':
			n = nodes[n.r]
		case 'L':
			n = nodes[n.l]
		}
	}
}

func findIndex(s string, n []string) int {
	for i := range n {
		if s == n[i] {
			return i
		}

	}
	return -1
}

// find the gcd of two ints using the Euclidean method
func gcd(x, y int) int {
	for y != 0 {
		tmp := y
		y = x % y
		x = tmp
	}
	return x
}

// find the lcm of at least two ints (panics if fewer than two)
func lcm(ints ...int) int {
	res := ints[0] * ints[1] / gcd(ints[0], ints[1])
	for i := 2; i < len(ints); i++ {
		res = res * ints[i] / gcd(res, ints[i])
	}
	return res
}
