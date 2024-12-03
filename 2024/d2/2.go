package d2

import (
	"bufio"
	"os"

	"github.com/cdillond/aoc"
)

const (
	Day  = "2"
	Year = "2024"
)

func check(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	const (
		up   = 1
		down = 2
	)
	dir, a, b := 0, 0, 1

	switch {
	case nums[b] > nums[a]:
		dir = up
	case nums[b] < nums[a]:
		dir = down
	default:
		return false
	}
	prev := a
	for i := b; i < len(nums); i++ {
		dif := nums[i] - nums[prev]
		if dif == 0 {
			return false
		} else if dif < 0 && dir == up {
			return false
		} else if dif > 0 && dir == down {
			return false
		}
		dif = aoc.Abs(dif)
		if dif < 1 || dif > 3 {
			return false
		}
		prev = i
	}
	return true
}

func checkSkip(nums []int, skip int) (int, bool) {
	if len(nums) <= 1 {
		return -1, true
	}
	const (
		up   = 1
		down = 2
	)
	var dir, a int
	if a == skip {
		a++
	}
	b := a + 1
	if b == skip {
		b++
	}
	if b >= len(nums) {
		return -1, true
	}
	switch {
	case nums[b] > nums[a]:
		dir = up
	case nums[b] < nums[a]:
		dir = down
	default:
		return b, false
	}
	prev := a
	for i := b; i < len(nums); i++ {
		if i == skip {
			continue
		}
		dif := nums[i] - nums[prev]
		if dif == 0 {
			return i, false
		} else if dir == up && dif < 0 {
			return i, false
		} else if dir == down && dif > 0 {
			return i, false
		}
		dif = aoc.Abs(dif)
		if dif < 1 || dif > 3 {
			return i, false
		}
		prev = i
	}
	return -1, true
}

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	var count int
	nums := make([]int, 0, 32)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums, err = aoc.ParseInts(scanner.Bytes(), nums[:0])
		if err != nil {
			return res, err
		}
		if check(nums) {
			count++
		}
	}

	return aoc.Itoa(count), nil
}

func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	var count int
	nums := make([]int, 0, 32)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums, err = aoc.ParseInts(scanner.Bytes(), nums[:0])
		if err != nil {
			return res, err
		}
		skip, safe := checkSkip(nums, -1)
		for i := 0; !safe && i < 3; i++ {
			_, safe = checkSkip(nums, skip)
			skip--
		}
		if safe {
			count++
		}
	}
	return aoc.Itoa(count), nil
}
