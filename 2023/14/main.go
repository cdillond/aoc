package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Fields(b)
	fmt.Println("part 1: ", score(north(rows)))
	fmt.Println("part 2: ", score(cycles(1_000_000_000, rows)))
}

func cycles(n int, rows [][]byte) [][]byte {
	out := make([][]byte, len(rows))
	for i := range out {
		out[i] = bytes.Clone(rows[i])
	}
	cache := make(map[string]int)
	cache[toString(out)] = 0
	var cstart int
	i := 0
	for ; i < n; i++ {
		out = cycle(out)
		s := toString(out)
		if x, ok := cache[s]; ok {
			out = cycle(out)
			cstart = x
			i++
			break
		} else {
			cache[s] = i + 1
		}
	}
	// (n - i - 1) is the number of cycles remaining in the main loop
	// (i - cstart) is the super-cycle's period
	rem := (n - i - 1) % (i - cstart)
	for j := 0; j < rem; j++ {
		out = cycle(out)
	}
	return out
}

func toString(rows [][]byte) string {
	bldr := new(strings.Builder)
	for i := range rows {
		bldr.Write(rows[i])
	}
	return bldr.String()
}

func score(rows [][]byte) int {
	var sum int
	for i, row := range rows {
		for j := range row {
			if row[j] == 'O' {
				sum += (len(rows) - i)
			}
		}
	}
	return sum
}

func cycle(rows [][]byte) [][]byte {
	return east(south(west(north(rows))))
}

func north(rows [][]byte) [][]byte {
	tilted := make([][]byte, len(rows[0]))
	for j := 0; j < len(rows[0]); j++ {
		tmp := make([]byte, 0, len(rows[0]))
		var dcount int
		for i := 0; i < len(rows); i++ {
			switch rows[i][j] {
			case 'O':
				tmp = append(tmp, 'O')
			case '.':
				dcount++
			case '#':
				for d := 0; d < dcount; d++ {
					tmp = append(tmp, '.')
				}
				tmp = append(tmp, '#')
				dcount = 0
			}
		}
		for k := len(tmp); k < len(rows[0]); k++ {
			tmp = append(tmp, '.')
		}
		for i := range tilted {
			tilted[i] = append(tilted[i], tmp[i])
		}
	}
	return tilted
}

func south(rows [][]byte) [][]byte {
	tilted := make([][]byte, len(rows[0]))
	for j := 0; j < len(rows[0]); j++ {
		tmp := make([]byte, 0, len(rows[0]))
		var dcount int
		for i := len(rows) - 1; i >= 0; i-- {
			switch rows[i][j] {
			case 'O':
				tmp = append(tmp, 'O')
			case '.':
				dcount++
			case '#':
				for d := 0; d < dcount; d++ {
					tmp = append(tmp, '.')
				}
				tmp = append(tmp, '#')
				dcount = 0
			}
		}
		for k := len(tmp); k < len(rows[0]); k++ {
			tmp = append(tmp, '.')
		}
		for i := range tilted {
			tilted[len(tilted)-1-i] = append(tilted[len(tilted)-1-i], tmp[i])
		}
	}
	return tilted
}

func east(rows [][]byte) [][]byte {
	tilted := make([][]byte, 0, len(rows))
	for i := range rows {
		tmp := make([]byte, 0, len(rows[i]))
		var ocount int
		for j := range rows[i] {
			switch rows[i][j] {
			case 'O':
				ocount++
			case '.':
				tmp = append(tmp, '.')
			case '#':
				for k := 0; k < ocount; k++ {
					tmp = append(tmp, 'O')
				}
				tmp = append(tmp, '#')
				ocount = 0
			}
		}
		for k := len(tmp); k < len(rows[i]); k++ {
			tmp = append(tmp, 'O')
		}
		tilted = append(tilted, tmp)
	}
	return tilted
}

func west(rows [][]byte) [][]byte {
	tilted := make([][]byte, 0, len(rows))
	for i := range rows {
		tmp := make([]byte, 0, len(rows[i]))
		var dcount int
		for j := range rows[i] {
			switch rows[i][j] {
			case 'O':
				tmp = append(tmp, 'O')
			case '.':
				dcount++
			case '#':
				for k := 0; k < dcount; k++ {
					tmp = append(tmp, '.')
				}
				tmp = append(tmp, '#')
				dcount = 0
			}
		}
		for k := len(tmp); k < len(rows[i]); k++ {
			tmp = append(tmp, '.')
		}
		tilted = append(tilted, tmp)
	}
	return tilted
}
