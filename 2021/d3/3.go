package d3

import (
	"aoc"
	"bufio"
	"embed"
	"io/fs"
)

//go:embed *.txt
var dir embed.FS

func countOnes(ones []int, line []byte) {
	_ = line[len(ones)-1]
	for i, c := range line {
		if c == '1' {
			ones[i]++
		}
	}
}

func Part1(path string) (res string, err error) {
	var f fs.File
	if f, err = dir.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// process the first line
	scanner.Scan()
	line := scanner.Bytes()
	ones := make([]int, len(line))
	i := 1
	countOnes(ones, line)

	for scanner.Scan() {
		countOnes(ones, scanner.Bytes())
		i++
	}
	half := i / 2
	// assume there is an odd number of lines
	var gamma, epsilon uint64
	for _, count := range ones {
		gamma <<= 1
		epsilon <<= 1

		// ones are most common
		if count > half {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}

	return aoc.Itoa(int(gamma * epsilon)), nil
}

func Part2(path string) (res string, err error) { return }
