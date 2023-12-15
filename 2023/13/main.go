package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fields := bytes.Split(b, []byte("\n\n"))
	var sum1, sum2 int
	for _, matrix := range fields {
		rows := bytes.Fields(matrix)
		sum1 += part1(rows)
		sum2 += part2(rows)
	}
	fmt.Println("part 1: ", sum1)
	fmt.Println("part 2: ", sum2)
}

func part1(rows [][]byte) int {
	var res int
	// check horizontal
	for i := 0; i < len(rows)-1; i++ {
		if bytes.Equal(rows[i], rows[i+1]) {
			// there is a potential reflection line
			if check(i, rows) {
				res += 100 * (i + 1)
				break
			}
		}
	}
	// check vertical
	if res == 0 {
		cols := toCols(rows)
		for i := 0; i < len(cols)-1; i++ {
			if bytes.Equal(cols[i], cols[i+1]) {
				// there is a potential reflection line
				if check(i, cols) {
					return res + i + 1
				}
			}
		}
	}
	return res
}

func part2(rows [][]byte) int {
	var res int
	// check horizontal
	for i := 0; i < len(rows)-1; i++ {
		if ok, corrected := equal(rows[i], rows[i+1], false); ok {
			// there is a potential reflection line
			if done, f := check2(i, rows, corrected); done && f {
				res += 100 * (i + 1)
				break
			}
		}
	}

	// check vertical
	if res == 0 {
		cols := toCols(rows)
		for i := 0; i < len(cols)-1; i++ {
			if ok, found := equal(cols[i], cols[i+1], false); ok {
				// there is a potential reflection line
				if done, f := check2(i, cols, found); done && f {
					res += (i + 1)
					break
				}
			}
		}
	}
	return res
}

// The first returned bool indicates whether a reflection line was found. The second bool indicates whether a single smudge has been corrected.
func check2(i int, rows [][]byte, corrected bool) (bool, bool) {
	found := corrected
	k := i - 1 // if corrected is true, then it won't do to start comparisons from i and i+1
	for j := i + 2; k >= 0 && j < len(rows); k, j = k-1, j+1 {
		ok, f := equal(rows[k], rows[j], found)
		if !ok {
			return false, corrected
		}
		if f {
			found = true
		}
	}
	return true, found
}

// The first returned bool indicates whether b1==b2. The second bool indicates whether a single smudge has been corrected.
func equal(b1, b2 []byte, corrected bool) (bool, bool) {
	if corrected {
		return bytes.Equal(b1, b2), corrected
	}
	switch dif(b1, b2) {
	case 0:
		return true, false
	case 1:
		return true, true
	default:
		return false, false
	}
}

// Returns the number of differences between the elements of b1 and b2. Retruns -1 if the two slices are of different length.
func dif(b1, b2 []byte) int {
	if len(b1) != len(b2) {
		return -1
	}
	var n int
	for i := range b1 {
		if b1[i] != b2[i] {
			n++
		}
	}
	return n
}

// Returns true if i is the start index of the reflection line.
func check(i int, rows [][]byte) bool {
	for j := i + 1; i >= 0 && j < len(rows); i, j = i-1, j+1 {
		if !bytes.Equal(rows[i], rows[j]) {
			return false
		}
	}
	return true
}

// Returns a slice of the columns of the matrix represented by rows.
func toCols(rows [][]byte) [][]byte {
	cols := make([][]byte, len(rows[0]))
	for i := range cols {
		cols[i] = make([]byte, len(rows))
	}
	for i := range rows {
		for j := range rows[i] {
			cols[j][i] = rows[i][j]
		}
	}
	return cols
}
