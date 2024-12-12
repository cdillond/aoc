package aoc

import (
	"errors"
	"os"
	"strconv"
	"unsafe"
)

func ReadOrDie(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if len(b) < 1 {
		panic("empty input file")
	}
	if b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	return b
}

func ReadTrimmed(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(b) < 1 {
		return nil, errors.New("empty input file")
	}
	if b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}
	return b, nil
}

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

func ParseInts(b []byte, nums []int) ([]int, error) {
	var (
		a, z, n int
		c       byte
		err     error
	)

loop:
	for ; a < len(b); a++ {
		c = b[a]
		if c != 0x20 && c != '\t' {
			break
		}
	}
	if a == len(b) {
		return nums, nil
	}
	for z = a + 1; z < len(b); z++ {
		c = b[z]
		if c == 0x20 || c == '\t' {
			break
		}
	}

	if n, err = A2i(b[a:z]); err != nil {
		return nums, err
	}
	nums = append(nums, n)
	if z < len(b) {
		a = z + 1
		goto loop
	}

	return nums, nil

}
