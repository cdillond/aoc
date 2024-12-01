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
	defer f.Close()

	left, right := make([]int, 0, 1024), make([]int, 0, 1024)

	scanner := bufio.NewScanner(f)

	var before, after []byte
	for scanner.Scan() {
		before, after, _ = bytes.Cut(scanner.Bytes(), []byte("   "))
		left = append(left, aoc.Atoi(before))
		right = append(right, aoc.Atoi(after))
	}

	slices.Sort(left)
	slices.Sort(right)

	var distance int
	for j := range left {
		distance += aoc.Abs(left[j] - right[j])
	}

	return aoc.Itoa(distance), nil
}

func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	left, right := make([]int, 0, 1024), make([]int, 0, 1024)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		before, after, _ := bytes.Cut(scanner.Bytes(), []byte("   "))
		left = append(left, aoc.Atoi(before))
		right = append(right, aoc.Atoi(after))
	}

	var total int
	for _, l := range left {
		var count int
		for _, r := range right {
			if l == r {
				count++
			}
		}
		total += count * l
	}

	return aoc.Itoa(total), nil
}
