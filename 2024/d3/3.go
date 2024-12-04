package d3

import (
	"bytes"
	"os"
	"regexp"

	"github.com/cdillond/aoc"
)

const (
	Day  = "3"
	Year = "2024"
)

func Part1(path string) (res string, err error) {
	var re *regexp.Regexp
	if re, err = regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)"); err != nil {
		return res, err
	}

	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return res, err
	}

	matches := re.FindAll(b, -1)
	var ans int
	for _, m := range matches {
		m = m[4 : len(m)-1]
		i := bytes.IndexByte(m, ',')
		a, b := aoc.Atoi(m[:i]), aoc.Atoi(m[i+1:])
		ans += a * b
	}
	return aoc.Itoa(ans), nil
}

const (
	Do      = 1
	Dont    = 2
	Product = 3
)

type parser struct {
	n int
	b []byte
}

func (p *parser) next() (c byte, ok bool) {
	if p.n >= len(p.b) {
		return c, false
	}
	c = p.b[p.n]
	p.n++
	return c, true
}

func (p *parser) match() (kind int, val int) {
	c, ok := p.next()

search_m_d:
	if !ok {
		return 0, 0
	}
	switch c {
	case 'm':
		goto search_ul
	case 'd':
		goto search_o
	default:
		c, ok = p.next()
		goto search_m_d
	}

search_o:
	if c, ok = p.next(); c != 'o' {
		goto search_m_d
	}
	c, ok = p.next()

	switch c {
	case '(':
		goto match_do
	case 'n':
		goto match_dont
	default:
		goto search_m_d
	}

match_do:
	if c, ok = p.next(); c != ')' {
		goto search_m_d
	}
	return Do, 0

match_dont:
	if c, ok = p.next(); c != '\'' {
		goto search_m_d
	}
	if c, ok = p.next(); c != 't' {
		goto search_m_d
	}
	if c, ok = p.next(); c != '(' {
		goto search_m_d
	}
	if c, ok = p.next(); c != ')' {
		goto search_m_d
	}
	return Dont, 0

search_ul:
	if c, ok = p.next(); c != 'u' {
		goto search_m_d
	}
	if c, ok = p.next(); c != 'l' {
		goto search_m_d
	}
	if c, ok = p.next(); c != '(' {
		goto search_m_d
	}

	var digitCount, numCount int
	var nums [2]int

match_num:
	c, ok = p.next()
	switch {
	case c == ',' && numCount == 0 && digitCount > 0:
		numCount++
		digitCount = 0
		goto match_num
	case c >= '0' && c <= '9' && digitCount < 3:
		nums[numCount] *= 10
		nums[numCount] += (int(c) - '0')
		digitCount++
		goto match_num
	case c == ')' && numCount == 1 && digitCount > 0:
		break
	default:
		goto search_m_d
	}
	return Product, nums[0] * nums[1]
}

func Part2(path string) (res string, err error) {
	var p parser
	if p.b, err = os.ReadFile(path); err != nil {
		return res, err
	}
	var ans int
	state := Do
	for kind, val := p.match(); kind != 0; kind, val = p.match() {
		if kind == Product {
			if state == Do {
				ans += val
			}
		} else {
			state = kind
		}
	}
	return aoc.Itoa(ans), nil
}
