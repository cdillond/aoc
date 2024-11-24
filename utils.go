package aoc

//go:generate go run gen.go

import (
	"errors"
	"strconv"
	"unsafe"
)

var ErrUndefined = errors.New("no solution found")

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

func Itoa(n int) string { return strconv.FormatInt(int64(n), 10) }
