package d4

import (
	"aoc"
	"bufio"
	"bytes"
	"os"
)

const (
	Day  = "4"
	Year = "2021"
)

func parseFirstLine(b []byte) []int {
	sp := bytes.Split(b, []byte{','})
	out := make([]int, len(sp))
	for i := range sp {
		out[i] = aoc.Atoi(sp[i])
	}
	return out
}

type bufr struct {
	data []byte
	n    int
}

func (b *bufr) next() []byte {
	// get start
	for b.n < len(b.data) && b.data[b.n] == ' ' {
		b.n++
	}
	start := b.n
	// get end
	for b.n < len(b.data) && b.data[b.n] != ' ' {
		b.n++
	}
	end := b.n
	return b.data[start:end]
}

func parseGrid(scanner *bufio.Scanner) [][]int {
	out := make([][]int, 5)
	for i := 0; i < 5 && scanner.Scan(); i++ {
		buf := bufr{data: scanner.Bytes()}

		for buf.n < len(buf.data) {
			out[i] = append(out[i], aoc.Atoi(buf.next()))
		}

	}

	return out
}

func unmarkedSum(nums []int, g grid) int {
	var sum int
	for _, row := range g {
		for _, num := range row {
			if !contains(nums, num) {
				sum += num
			}
		}
	}

	return sum
}

func contains(nums []int, n int) bool {
	for _, num := range nums {
		if n == num {
			return true
		}
	}
	return false
}

type grid [][]int
type gc [2][5]int

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	nums := parseFirstLine(scanner.Bytes())
	var grids []grid
	for scanner.Scan() {
		grids = append(grids, parseGrid(scanner))
	}
	gcs := make([]gc, len(grids))

	var count int

	for _, num := range nums {
		count++
		for i, grid := range grids {
			for row := range grid {
				for col := range grid[row] {
					if grid[row][col] == num {
						gcs[i][0][row]++
						gcs[i][1][col]++
					}
				}
			}
			for j := 0; j < 5; j++ {
				if gcs[i][0][j] == 5 || gcs[i][1][j] == 5 {
					// get the sum of all unmarked nums
					return aoc.Itoa(num * unmarkedSum(nums[:count], grids[i])), nil
				}
			}
		}
	}

	return "", nil
}

func Part2(path string) (res string, err error) {
	// copy part 1...
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	nums := parseFirstLine(scanner.Bytes())
	var grids []grid
	for scanner.Scan() {
		grids = append(grids, parseGrid(scanner))
	}
	gcs := make([]gc, len(grids))

	var count int
	set := make(map[int]struct{})

	for _, num := range nums {
		count++
		for i, grid := range grids {
			for row := range grid {
				for col := range grid[row] {
					if grid[row][col] == num {
						gcs[i][0][row]++
						gcs[i][1][col]++
					}
				}
			}
			for j := 0; j < 5; j++ {
				if gcs[i][0][j] == 5 || gcs[i][1][j] == 5 {
					// add this to the set
					set[i] = struct{}{}

					if len(set) == len(grids) {
						return aoc.Itoa(num * unmarkedSum(nums[:count], grids[i])), nil
					}

				}
			}
		}
	}
	return "", nil

}
