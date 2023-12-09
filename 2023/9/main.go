package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var pt1, pt2 int64
	for scanner.Scan() {
		b := scanner.Bytes()
		row := parseRow(b)
		d := differences(row)
		pt1 += row[len(row)-1] + d[len(d)-1]
		starts := []int64{row[0], d[0]}
		for !zeros(d) {
			d = differences(d)
			pt1 += d[len(d)-1]
			starts = append(starts, d[0])
		}
		pt2 += negs(starts...)
	}
	fmt.Println("part 1: ", pt1)
	fmt.Println("part 2: ", pt2)
}

func negs(in ...int64) int64 {
	var out int64
	for i := len(in) - 1; i >= 0; i-- {
		out = in[i] - out
	}
	return out
}

// Returns a slice consisting of the differences between consecutive values of the input row.
func differences(row []int64) []int64 {
	out := make([]int64, 0, len(row)/2)
	for i := 0; i < len(row)-1; i++ {
		out = append(out, row[i+1]-row[i])
	}
	return out
}

// Returns true if row consists of only zeros.
func zeros(row []int64) bool {
	for i := range row {
		if row[i] != 0 {
			return false
		}
	}
	return true
}

// Parses b into an int64 slice.
func parseRow(b []byte) []int64 {
	split := bytes.Split(b, []byte(" "))
	out := make([]int64, len(split))
	for i := range split {
		n, err := strconv.ParseInt(string(split[i]), 10, 64)
		if err == nil {
			out[i] = n
		}
	}
	return out
}
