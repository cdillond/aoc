package aoc

import (
	"strconv"
	"unsafe"
)

func Atoi(b []byte) int {
	n, err := strconv.ParseInt(unsafe.String(unsafe.SliceData(b), len(b)), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}
