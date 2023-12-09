package main

import (
	"bytes"
	"fmt"
	"os"
)

type tuple struct {
	i, j int
}

func IsSymbol(b byte) bool {
	return (b < '0' || b > '9') && b != '.'
}

func IsNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(b, []byte("\n"))
	fmt.Println("part 1: ", Part1(rows)) // 556057
	fmt.Println("part 2: ", Part2(rows)) // 82824352
}

func Part1(rows [][]byte) int {
	// read the bytes into a slice of slices
	var sum int
	// iterate through each row to find numbers
	for i, row := range rows {
		for j := 0; j < len(row); j++ {
			// check if b[j] is a digit
			var val, k int
			if IsNum(row[j]) {
				// if yes, we can continue to accumulate its value
				val = int(row[j] - '0')
				for k = j + 1; k < len(row) && IsNum(row[k]); k++ {
					val *= 10
					val += int(row[k] - '0')
				}
				// we now have our number, its start index, and its end index
				var symbol bool
				// check for adjacency to a symbol manually, area by area
				// upper row
				if i != 0 {
					for l := max(0, j-1); l < min(k+1, len(rows[i-1])-1); l++ {
						if IsSymbol(rows[i-1][l]) {
							symbol = true
							goto adjacent
						}
					}
				}
				// lower row
				if i < len(rows)-1 {
					for l := max(0, j-1); l < min(k+1, len(rows[i+1])-1); l++ {
						if IsSymbol(rows[i+1][l]) {
							symbol = true
							goto adjacent
						}
					}
				}
				// left-hand
				if j != 0 && IsSymbol(rows[i][j-1]) {
					symbol = true
					goto adjacent
				}
				// right-hand
				if k < len(rows[i]) && IsSymbol(rows[i][k]) {
					symbol = true
				}
			adjacent:
				if symbol {
					sum += val
				}

			}
			// skip to the end of the number that was found
			if k != 0 {
				j = k
			}
		}
	}
	return sum
}

func Part2(rows [][]byte) int {

	var sum int
	// first transform the bytes into a slice of int slices to make it conceptually easier to handle
	// 0 = ignore; -142 = gear; positive integer = number; negative integer > -142 = offset of integer startpoint
	// -142 is used because the max length of each line (including \n is 141); in any case, this would overflow int64
	// and all of the numbers are three digits or fewer
	islice := make([][]int, len(rows))

	for m, row := range rows {
		var num int
		start := -1
		if IsNum(row[0]) {
			num = int(row[0] - '0')
			start = 0
		}
		islice[m] = make([]int, len(row))
		for i := 1; i < len(row); i++ {
			if row[i] == '*' {
				islice[m][i] = -142
			}
			if IsNum(row[i]) && start == -1 { // start a new number
				num = int(row[i] - '0')
				start = i
				continue
			}
			if IsNum(row[i]) { // add to the existing number
				num *= 10
				num += int(row[i] - '0')
				islice[m][i] = start - i // record the offset
			}
			if !IsNum(row[i]) && start != -1 { // record the full number and reset the start
				islice[m][start] = num
				num = 0
				start = -1
			}
		}
	}

	for i, row := range islice {
		for j, val := range row {
			// check if val is a gear
			if val != -142 {
				continue
			}
			factors := make(map[tuple]int) // map the index ("i,j") to the value; easiest just to do this as a string
			// upper
			if i > 0 {
				for k := max(0, j-1); k <= min(j+1, len(row)); k++ {
					if islice[i-1][k] > 0 {
						factors[tuple{i - 1, k}] = islice[i-1][k]
					} else if islice[i-1][k] < 0 && islice[i-1][k] != -142 {
						factors[tuple{i - 1, k + islice[i-1][k]}] = islice[i-1][k+islice[i-1][k]]
					}
				}
			}

			// lower
			if i < len(islice)-1 {
				for k := max(0, j-1); k <= min(j+1, len(row)); k++ {
					if islice[i+1][k] > 0 {
						factors[tuple{i + 1, k}] = islice[i+1][k]
					} else if islice[i+1][k] < 0 && islice[i+1][k] != -142 {
						factors[tuple{i + 1, k + islice[i+1][k]}] = islice[i+1][k+islice[i+1][k]]
					}
				}
			}
			// left
			if j > 0 {
				if row[j-1] > 0 {
					factors[tuple{i, j - 1}] = row[j-1]
				} else if row[j-1] < 0 && row[j-1] != -142 {
					factors[tuple{i, j - 1 + row[j-1]}] = row[j-1+row[j-1]]
				}
			}
			// right
			if j < len(row)-1 && row[j+1] > 0 {
				factors[tuple{i, j + 1}] = row[j+1]
			}
			if len(factors) == 2 {
				tmp := 1
				for _, val := range factors {
					tmp *= val
				}
				sum += tmp
			}
		}
	}
	return sum
}
