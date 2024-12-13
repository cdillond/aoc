package d13

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"

	"github.com/cdillond/aoc"
)

//go:embed input.txt
var input []byte

const (
	Day  = "13"
	Year = "2024"
)

type system struct {
	a1, a2 int64
	b1, b2 int64
	c1, c2 int64
}

var ErrParse = errors.New("error parsing input")

func (s *system) unmarshal(scanner *bufio.Scanner) error {
	if !scanner.Scan() {
		return ErrParse
	}
	b := scanner.Bytes()
	b = b[len("Button A: X+"):]
	n := bytes.IndexByte(b, ',')
	s.a1 = int64(aoc.Atoi(b[:n]))
	n += 1 + len(" Y+")
	s.a2 = int64(aoc.Atoi(b[n:]))

	if !scanner.Scan() {
		return ErrParse
	}
	b = scanner.Bytes()
	b = b[len("Button B: X+"):]
	n = bytes.IndexByte(b, ',')
	s.b1 = int64(aoc.Atoi(b[:n]))
	n += 1 + len(" Y+")
	s.b2 = int64(aoc.Atoi(b[n:]))

	if !scanner.Scan() {
		return ErrParse
	}
	b = scanner.Bytes()
	b = b[len("Prize: X="):]
	n = bytes.IndexByte(b, ',')
	s.c1 = int64(aoc.Atoi(b[:n]))
	n += 1 + len(" Y=")
	s.c2 = int64(aoc.Atoi(b[n:]))
	return nil
}

func (s system) solve() (a, b int64) {
	// if the determinant is 0, the program will panic later on.
	// luckily, with the given input, that never happens. this
	// also means that none of the button vectors are colinear,
	// so the solutions are unique and there's no need to find
	// the 'minimum' cost.
	det := s.a1*s.b2 - s.a2*s.b1

	aNum := s.c1*s.b2 - s.b1*s.c2
	bNum := s.a1*s.c2 - s.c1*s.a2

	if a = aNum / det; a < 0 || a*det != aNum {
		return 0, 0
	}
	if b = bNum / det; b < 0 || b*det != bNum {
		return 0, 0
	}
	return a, b
}

func Part1(_ string) (res string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var s system
	var a, b, total int64
	for {
		if err = s.unmarshal(scanner); err != nil {
			return res, err
		}
		a, b = s.solve()
		total += 3*a + b
		if !scanner.Scan() {
			break
		}
	}
	return aoc.Itoa(int(total)), nil
}

func Part2(_ string) (res string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var s system
	var a, b, total int64
	for {
		if err = s.unmarshal(scanner); err != nil {
			return res, err
		}
		s.c1 += 10000000000000
		s.c2 += 10000000000000
		a, b = s.solve()
		total += 3*a + b
		if !scanner.Scan() {
			break
		}
	}
	return aoc.Itoa(int(total)), nil
}
