package d7

import (
	"bufio"
	"bytes"
	"math"
	"os"

	"github.com/cdillond/aoc"
)

const (
	Day  = "7"
	Year = "2024"
)

func cat(a, b int) int {
	fa := float64(a)
	fb := float64(b)
	fa *= math.Pow10(1 + int(math.Log10(fb)))
	fa += fb
	return int(fa)
}

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var nums []int
	var count int
	cur, next := make([]int, 0, 256), make([]int, 0, 256)
	for scanner.Scan() {
		b := scanner.Bytes()
		split := bytes.Split(b, []byte{':'})
		target := aoc.Atoi(split[0])

		nums, _ = aoc.ParseInts(split[1], nums[:0])

		cur = append(cur[:0], nums[0])
		for i := 1; i < len(nums) && len(cur) > 0; i++ {
			n := nums[i]
			for _, c := range cur {
				if add := c + n; add <= target {
					next = append(next, add)
				}
				if mul := c * n; mul <= target {
					next = append(next, mul)
				}
			}
			cur, next = next, cur
			next = next[:0]
		}

		for _, n := range cur {
			if n == target {
				count += target
				break
			}
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

	scanner := bufio.NewScanner(f)
	var nums []int
	var count int
	next, cur := make([]int, 0, 512), make([]int, 1, 512)
	for scanner.Scan() {
		b := scanner.Bytes()
		split := bytes.Split(b, []byte{':'})
		target := aoc.Atoi(split[0])

		nums, _ = aoc.ParseInts(split[1], nums[:0])

		cur = append(cur[:0], nums[0])
		var add, mul, catn int
		for i := 1; i < len(nums) && len(cur) > 0; i++ {
			n := nums[i]
			for _, c := range cur {
				if add = c + n; add <= target {
					next = append(next, add)
				}
				if mul = c * n; mul <= target {
					next = append(next, mul)
				}
				if catn = cat(c, n); catn <= target {
					next = append(next, catn)
				}
			}
			cur, next = next, cur
			next = next[:0]
		}

		for _, n := range cur {
			if n == target {
				count += target
				break
			}
		}

	}

	return aoc.Itoa(count), nil
}
