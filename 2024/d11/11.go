package d11

import (
	"os"

	"github.com/cdillond/aoc"
)

const (
	Day  = "11"
	Year = "2024"
)

var lut = [...]int{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	1000000000000000000,
}

func p10(n int) int {
	return lut[n]
}

func numDigits(n int) int {
	for i := 1; i < len(lut); i++ {
		if n < lut[i] {
			return i
		}
	}
	return len(lut)
}

func solve(rounds int, line []byte) (int, error) {
	nums, err := aoc.ParseInts(line, nil)
	if err != nil {
		return 0, err
	}

	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}
	next := make(map[int]int, len(cache))
	var digits int
	for range rounds {
		for num, count := range cache {
			if num == 0 {
				next[1] += count
			} else if digits = numDigits(num); digits&1 == 0 {
				digits = p10(digits / 2)
				lowerHalf := num % digits
				upperHalf := num / digits
				next[lowerHalf] += count
				next[upperHalf] += count
			} else {
				next[num*2024] += count
			}
		}
		cache, next = next, cache
		clear(next)
	}

	var sum int
	for _, count := range cache {
		sum += count
	}
	return sum, nil
}

func Part1(path string) (res string, err error) {
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return res, err
	}
	if len(b) > 0 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	var ans int
	if ans, err = solve(25, b); err != nil {
		return res, err
	}
	return aoc.Itoa(ans), nil
}

func Part2(path string) (res string, err error) {
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return res, err
	}
	if len(b) > 0 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	var ans int
	if ans, err = solve(75, b); err != nil {
		return res, err
	}
	return aoc.Itoa(ans), nil
}
