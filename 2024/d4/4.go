package d4

import (
	"bufio"
	"os"

	"github.com/cdillond/aoc"
)

const (
	Day  = "4"
	Year = "2024"
)

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := make([]byte, len(bytes))
		copy(line, bytes)
		lines = append(lines, line)
	}
	var count int
	for i, line := range lines {
		for j, c := range line {
			if c == 'X' {
				// forward
				if j < (len(lines)-3) && line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S' {
					count++
				}

				// reverse
				if j > 2 && line[j-1] == 'M' && line[j-2] == 'A' && line[j-3] == 'S' {
					count++
				}

				// down
				if i < (len(lines)-3) && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
					count++
				}
				// up
				if i > 2 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
					count++
				}

				// forward - down
				if i < (len(lines)-3) && j < (len(lines)-3) && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					count++
				}

				// forward - up
				if i > 2 && j < (len(lines)-3) && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
					count++
				}

				// reverse - down
				if i < (len(lines)-3) && j > 2 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					count++
				}

				// reverse - up
				if i > 2 && j > 2 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
					count++
				}

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

	var lines [][]byte
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		line := make([]byte, len(bytes))
		copy(line, bytes)
		lines = append(lines, line)
	}
	var count int
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == 'A' {
				diag0 := lines[i-1][j-1]|lines[i+1][j+1] == 'S'|'M'
				diag1 := lines[i+1][j-1]|lines[i-1][j+1] == 'S'|'M'
				if diag0 && diag1 {
					count++
				}
			}
		}
	}

	return aoc.Itoa(count), nil
}
