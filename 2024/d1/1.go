package d1

import (
	"bufio"
	"bytes"
	"os"
	"slices"

	"github.com/cdillond/aoc"
)

const (
	Day  = "1"
	Year = "2024"
)

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}

	scanner := bufio.NewScanner(f)
	var left, right []int
	for scanner.Scan() {
		before, after, _ := bytes.Cut(scanner.Bytes(), []byte("   "))
		left = append(left, aoc.Atoi(before))
		right = append(right, aoc.Atoi(after))
	}

	slices.Sort(left)
	slices.Sort(right)

	var distance int
	for i := range left {
		distance += aoc.Abs(left[i] - right[i])
	}

	return aoc.Itoa(distance), nil
}

func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}

	scanner := bufio.NewScanner(f)
	var left, right []int
	for scanner.Scan() {
		before, after, _ := bytes.Cut(scanner.Bytes(), []byte("   "))
		left = append(left, aoc.Atoi(before))
		right = append(right, aoc.Atoi(after))
	}

	seen := make(map[int]int, len(left))

	var (
		total int
		count int
		ok    bool
	)
	for _, n := range left {
		if count, ok = seen[n]; ok {
			total += count * n
			continue
		}
		for _, v := range right {
			if n == v {
				count++
			}
		}
		seen[n] = count
		total += count * n
	}

	return aoc.Itoa(total), nil
}
