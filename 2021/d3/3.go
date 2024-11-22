package d3

import (
	"aoc"
	"bufio"
	"bytes"
	"embed"
	"io"
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

func Part2(path string) (res string, err error) {
	var f fs.File
	if f, err = dir.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	txt, err := io.ReadAll(f)
	if err != nil {
		return res, err
	}
	lines := bytes.Fields(txt)
	// first find oxygen rating
	cur := make([][]byte, len(lines))
	next := make([][]byte, 0, len(lines))
	copy(cur, lines)

	for i := 0; i < len(cur[0]); i++ {
		// pass 1, determine the most common bit
		var onesCount int
		for n := range cur {
			if cur[n][i] == '1' {
				onesCount++
			}
		}

		// ones are more common
		if onesCount >= len(cur)-onesCount {
			for _, line := range cur {
				if line[i] == '1' {
					next = append(next, line)
				}
			}
		} else {
			for _, line := range cur {
				if line[i] == '0' {
					next = append(next, line)
				}
			}

		}

		cur = next

		next = nil
		if len(cur) == 1 {
			break
		}

	}

	var oxygen uint64
	for _, char := range cur[0] {
		oxygen <<= 1
		if char == '1' {
			oxygen |= 1
		}
	}

	// first find oxygen rating
	cur = make([][]byte, len(lines))
	next = make([][]byte, 0, len(lines))
	copy(cur, lines)

	for i := 0; i < len(cur[0]); i++ {
		// pass 1, determine the most common bit
		var zeroCount int
		for n := range cur {
			if cur[n][i] == '0' {
				zeroCount++
			}
		}

		// ones are more common
		if zeroCount <= len(cur)-zeroCount {
			for _, line := range cur {
				if line[i] == '0' {
					next = append(next, line)
				}
			}
		} else {
			for _, line := range cur {
				if line[i] == '1' {
					next = append(next, line)
				}
			}

		}

		cur = next

		next = nil

		if len(cur) == 1 {
			break
		}
	}

	var co2 uint64
	for _, char := range cur[0] {
		co2 <<= 1
		if char == '1' {
			co2 |= 1
		}
	}

	return aoc.Itoa(int(oxygen * co2)), nil
}
