package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type swap struct {
	i, j int
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	fields := bytes.Split(b, []byte("\n\n"))
	var count int
	var count2 int
	for _, field := range fields {
		rows := strings.Fields(string(field))
		for i, row := range rows {
			if i < len(rows)-1 && row == rows[i+1] {
				if ok := check(rows, i); ok {
					count += 100 * (i + 1)
				}
			}
		}

		cols := rotate(rows)
		for i, col := range cols {
			if i < len(cols)-1 && col == cols[i+1] {
				if ok := check(cols, i); ok {
					count += (i + 1)
				}
			}
		}

		var smudged bool
		//var found bool
		for i, row := range rows {
			if i < len(rows)-1 {
				switch cmp(row, rows[i+1]) {
				case 0:
				case 1:
					if smudged {
						continue
					}
					smudged = true
				default:
					continue
				}
				if ok, _ := check2(rows, i, smudged); ok {
					//if !s {
					//	continue
					//}
					count2 += 100 * (i + 1)

					break
				}
			}
		}

	}
	fmt.Println(count)
	fmt.Println(count2) // too low 32900 too high 38100
}

func idif(s1, s2 string) int {
	for i := range s1 {
		if s1[i] != s2[i] {
			return i
		}
	}
	return -1
}

func cmp(s1, s2 string) int {
	var dif int
	for i := range s1 {
		if s1[i] != s2[i] {
			dif++
		}
	}
	return dif
}

func check2(r []string, i int, smudged bool) (bool, bool) {
	//if smudged {
	//	return check(r, i)
	//}
	for j := i + 1; i >= 0 && j < len(r); i, j = i-1, j+1 {
		if r[i] != r[j] {
			if !smudged {
				if cmp(r[i], r[j]) == 1 {
					smudged = true
					continue
				}
			}
			return false, smudged
		}
	}
	return true, smudged
}

func check(r []string, i int) bool {
	for j := i + 1; i >= 0 && j < len(r); i, j = i-1, j+1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func rotate(r []string) []string {
	out := make([]string, len(r[0]))
	for i := range r[0] {
		for j := range r {
			out[i] = out[i] + string(r[j][i])
		}
	}
	return out
}
