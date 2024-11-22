package d1

import (
	"aoc"
	"bufio"
	"os"
	"strconv"
)

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cur, prev, count int
	if !scanner.Scan() {
		return res, scanner.Err()
	}
	if cur, err = aoc.A2i(scanner.Bytes()); err != nil {
		return res, err
	}
	for scanner.Scan() {
		prev = cur
		if cur, err = aoc.A2i(scanner.Bytes()); err != nil {
			return "", err
		}
		if cur > prev {
			count++
		}
	}

	return strconv.Itoa(count), nil
}
func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// there's no need to track more than the 4 most recent values.
	rows := make([]int, 4)
	var (
		i, n, sum int
		count     int
	)

	for ; i < 3; i++ {
		scanner.Scan()
		if n, err = aoc.A2i(scanner.Bytes()); err != nil {
			return res, err
		}
		rows[i] = n
		sum += n
	}

	for scanner.Scan() {
		if n, err = aoc.A2i(scanner.Bytes()); err != nil {
			return res, err
		}
		rows[i&3] = n // i&3 == i%4

		// n can be reused to represent the difference between window sums
		n = n - rows[(i-3)&3]
		if n > 0 {
			count++
		}

		sum += n
		i++
	}
	return strconv.Itoa(count), nil
}
