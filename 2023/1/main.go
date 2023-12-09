package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum1, sum2 int
	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) == 0 {
			continue
		}
		sum1 += Part1(b)
		sum2 += Part2(b)
	}
	fmt.Println("part 1: ", sum1)
	fmt.Println("part 2: ", sum2)
}

func Part1(row []byte) int {
	first, last := -1, -1
	for i := 0; i < len(row) && (first == -1 || last == -1); i++ {
		if first == -1 && IsNum(row[i]) {
			first = int(row[i] - '0')
		}
		if last == -1 && IsNum(row[len(row)-1-i]) {
			last = int(row[len(row)-1-i] - '0')
		}
	}
	return first*10 + last
}

func Part2(row []byte) int {
	first, last := -1, -1
	for i := 0; i < len(row) && (first == -1 || last == -1); i++ {
		if first == -1 {
			if IsNum(row[i]) {
				first = int(row[i] - '0')
			} else {
				first = ToDigit(row[i:])
			}
		}
		if last == -1 {
			if IsNum(row[len(row)-1-i]) {
				last = int(row[len(row)-1-i] - '0')
			} else {
				last = ToDigit(row[len(row)-1-i:])
			}
		}
	}
	return first*10 + last
}

func IsNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func ToDigit(b []byte) int {
	if bytes.HasPrefix(b, []byte("one")) {
		return 1
	} else if bytes.HasPrefix(b, []byte("two")) {
		return 2
	} else if bytes.HasPrefix(b, []byte("three")) {
		return 3
	} else if bytes.HasPrefix(b, []byte("four")) {
		return 4
	} else if bytes.HasPrefix(b, []byte("five")) {
		return 5
	} else if bytes.HasPrefix(b, []byte("six")) {
		return 6
	} else if bytes.HasPrefix(b, []byte("seven")) {
		return 7
	} else if bytes.HasPrefix(b, []byte("eight")) {
		return 8
	} else if bytes.HasPrefix(b, []byte("nine")) {
		return 9
	} else if bytes.HasPrefix(b, []byte("zero")) {
		return 0
	}
	return -1
}
