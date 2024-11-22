package d1

import (
	"aoc"
	"bufio"
	"os"
	"strconv"
)

func init() {
	aoc.Problems[0][0] = Part1
	aoc.Problems[0][1] = Part2
}

func Part1(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cur, prev, i, count int
	for ; scanner.Scan(); i++ {
		prev = cur
		if cur, err = strconv.Atoi(scanner.Text()); err != nil {
			return "", err
		}
		if i > 0 && cur > prev {
			count++
		}
	}

	return strconv.Itoa(count), nil
}
func Part2(path string) (string, error) {
	var (
		f   *os.File
		err error
		res string
	)
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var (
		rows      []int
		i, n, sum int
		count     int
	)

	for ; i < 3; i++ {
		scanner.Scan()
		if n, err = strconv.Atoi(scanner.Text()); err != nil {
			return res, err
		}
		rows = append(rows, n)
		sum += n
	}

	for scanner.Scan() {
		if n, err = strconv.Atoi(scanner.Text()); err != nil {
			return res, err
		}
		rows = append(rows, n)

		n = n - rows[i-3]
		if n > 0 {
			count++
		}

		sum += n
		i++
	}

	return strconv.Itoa(count), nil
}
