package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(b, []byte("\n"))
	var pt1, pt2 int
	for _, row := range rows {
		split := bytes.Fields(row)
		pt1 += solve(split[0], parse(split[1]))
		pt2 += solve(unfoldSprings(split[0]), unfold(parse(split[1])))
	}
	fmt.Println("part 1: ", pt1)
	fmt.Println("part 2: ", pt2)
}

type hrun struct {
	count int // number of consecutive '#' chars
	n     int // current index of the runlen slice
	mul   int // cache multiplier (number of equivalent hruns)
}

func solve(src []byte, lens []int) int {
	runs := []hrun{{}}
	set := make(map[hrun]int)
	for i := range src {
		var out []hrun
		for _, s := range runs {
			mul := set[s]
			delete(set, s)
			s.mul = mul

			switch src[i] {
			case '#':
				s.count++
				if s.n < len(lens) && s.count <= lens[s.n] {
					out = append(out, s)
				}
			case '.':
				if s.count > 0 {
					if s.n < len(lens) && s.count == lens[s.n] {
						s.n++
						s.count = 0
						out = append(out, s)
					}
				} else {
					out = append(out, s)
				}
			case '?':
				dot := s // copy s
				// treat s as if ? = #
				s.count++
				if s.n < len(lens) && s.count <= lens[s.n] {
					out = append(out, s)
				}
				// treat dot as if ? = .
				if dot.count > 0 {
					if dot.n < len(lens) && dot.count == lens[dot.n] {
						dot.n++
						dot.count = 0
						out = append(out, dot)
					}
				} else {
					out = append(out, dot)
				}
			}
		}

		var pruned []hrun
		for _, s := range out {
			if _, ok := set[s]; !ok {
				pruned = append(pruned, s)
				set[s] = max(1, s.mul)
			} else {
				set[s] += s.mul
			}
		}
		runs = pruned
	}
	var final int
	for _, s := range runs {
		if s.count == 0 && s.n == len(lens) {
			final += set[s]
		} else {
			if s.n == len(lens)-1 && s.count == lens[s.n] {
				final += set[s]
			}
		}
	}

	return final
}

func unfold[T any](s []T) []T {
	dst := make([]T, 0, 5*len(s))
	for i := 0; i < 5; i++ {
		dst = append(dst, s...)
	}
	return dst
}

func unfoldSprings(s []byte) []byte {
	var out []byte
	for i := 0; i < 4; i++ {
		out = append(out, s...)
		out = append(out, '?')
	}
	out = append(out, s...)
	return out
}

func parse(b []byte) []int {
	var out []int
	i, j := 0, 0
	for ; j < len(b); j++ {
		if b[j] == ',' {
			n, _ := strconv.Atoi(string(b[i:j]))
			out = append(out, n)
			i = j + 1
		}
	}
	n, _ := strconv.Atoi(string(b[i:]))
	out = append(out, n)
	return out
}
