package y21

// code generated by ../gen.go

import (
	"aoc"

	"aoc/2021/d1"
	"aoc/2021/d2"
	"aoc/2021/d3"
)

type solution func(string) (string, error)

func Solve(day, part int, path string) (string, error) {
	if part != 1 && part != 2 {
		return "", aoc.ErrUndefined
	}

	day--
	part--

	solutions := [...]solution{
		0:	d1.Part1,
		1:	d1.Part2,
		2:	d2.Part1,
		3:	d2.Part2,
		4:	d3.Part1,
		5:	d3.Part2,
	}

	i := (2 * day) + part
	if i < 0 || i >= len(solutions) {
		return "", aoc.ErrUndefined
	}
	if solve := solutions[i]; solve != nil {
		return solve(path)
	}
	return "", aoc.ErrUndefined
}
