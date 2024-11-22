package d2

import (
	"aoc"
	"bufio"
	"bytes"
	"embed"
	"io/fs"
)

//go:embed *.txt
var dir embed.FS

func Part1(path string) (res string, err error) {
	var f fs.File
	if f, err = dir.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var h, d int
	for scanner.Scan() {
		before, after, found := bytes.Cut(scanner.Bytes(), []byte{' '})
		if !found {
			break
		}

		switch {
		case bytes.Equal(before, []byte("forward")):
			h += aoc.Atoi(after)
		case bytes.Equal(before, []byte("down")):
			d += aoc.Atoi(after)
		case bytes.Equal(before, []byte("up")):
			d -= aoc.Atoi(after)
		}
	}

	return aoc.Itoa(h * d), nil
}
func Part2(path string) (res string, err error) {
	var f fs.File
	if f, err = dir.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var h, d, aim int
	for scanner.Scan() {
		before, after, found := bytes.Cut(scanner.Bytes(), []byte{' '})
		if !found {
			break
		}

		switch {
		case bytes.Equal(before, []byte("forward")):
			t := aoc.Atoi(after)
			h += t
			d += (aim * t)
		case bytes.Equal(before, []byte("down")):
			aim += aoc.Atoi(after)
		case bytes.Equal(before, []byte("up")):
			aim -= aoc.Atoi(after)
		}
	}

	return aoc.Itoa(h * d), nil
}
