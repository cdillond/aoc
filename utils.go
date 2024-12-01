package aoc

import (
	"strconv"
	"unsafe"
)

func Abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func Atoi(b []byte) int {
	n, err := strconv.ParseInt(unsafe.String(unsafe.SliceData(b), len(b)), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func A2i(b []byte) (int, error) {
	n, err := strconv.ParseInt(unsafe.String(unsafe.SliceData(b), len(b)), 10, 64)
	return int(n), err
}

func Stoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func Itoa(n int) string { return strconv.FormatInt(int64(n), 10) }
