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
	a, b   int64
}

func (s *system) unmarshal(scanner *bufio.Scanner) error {
	var (
		ok       bool
		ErrParse = errors.New("error parsing input")
	)

	if ok = scanner.Scan(); !ok {
		return ErrParse
	}
	b := scanner.Bytes()
	b = b[len("Button A: X+"):]
	n := bytes.IndexByte(b, ',')
	s.a1 = int64(aoc.Atoi(b[:n]))
	n += 1 + len(" Y+")
	s.a2 = int64(aoc.Atoi(b[n:]))

	if ok = scanner.Scan(); !ok {
		return ErrParse
	}
	b = scanner.Bytes()
	b = b[len("Button B: X+"):]
	n = bytes.IndexByte(b, ',')
	s.b1 = int64(aoc.Atoi(b[:n]))
	n += 1 + len(" Y+")
	s.b2 = int64(aoc.Atoi(b[n:]))

	if ok = scanner.Scan(); !ok {
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

func (s *system) solve() {
	aNum, aDen := s.c1*s.b2-s.b1*s.c2, s.a1*s.b2-s.b1*s.a2
	bNum, bDen := s.a1*s.c2-s.c1*s.a2, s.a1*s.b2-s.b1*s.a2

	if s.a = aNum / aDen; s.a*aDen != aNum {
		s.a, s.b = 0, 0
		return
	}
	if s.b = bNum / bDen; s.b*bDen != bNum {
		s.a, s.b = 0, 0
	}
}

func (s system) cost() int64 {
	return 3*s.a + s.b

}

func Part1(_ string) (res string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	s := new(system)
	var total int64
	for {
		if err = s.unmarshal(scanner); err != nil {
			return res, err
		}
		s.solve()
		total += s.cost()
		if !scanner.Scan() {
			break
		}
	}
	return aoc.Itoa(int(total)), nil
}

func Part2(_ string) (res string, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	s := new(system)
	var total int64
	for {
		if err = s.unmarshal(scanner); err != nil {
			return res, err
		}
		s.c1 += 10000000000000
		s.c2 += 10000000000000
		s.solve()
		total += s.cost()
		if !scanner.Scan() {
			break
		}
	}
	return aoc.Itoa(int(total)), nil
}
