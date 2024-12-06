package d5

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/cdillond/aoc"
)

const (
	Day  = "5"
	Year = "2024"
)

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	m := make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		slice := strings.Split(text, "|")
		if len(slice) < 2 {
			panic(slice)
		}
		v := append(m[slice[0]], slice[1])
		m[slice[0]] = v
	}
	var (
		count int
		v     []string
		ok    bool
		str   string
	)
	for scanner.Scan() {
		text := scanner.Text()
		slice := strings.Split(text, ",")

		for i, s := range slice {
			if v, ok = m[s]; ok {
				for _, str = range v {
					if n := slices.Index(slice, str); n != -1 && n < i {
						goto skip
					}
				}
			}
		}
		count += aoc.Stoi(slice[(len(slice)+(^(len(slice)&1)&1))>>1])

	skip:
	}

	return aoc.Itoa(count), nil
}

func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()
	m := make(map[string][]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		slice := strings.Split(text, "|")
		if len(slice) < 2 {
			panic(slice)
		}
		v := append(m[slice[0]], slice[1])
		m[slice[0]] = v
	}
	var count int
	for scanner.Scan() {
		text := scanner.Text()
		slice := strings.Split(text, ",")
		dst := make([]string, len(slice))
		copy(dst, slice)
		sort.Slice(slice, func(i, j int) bool {
			if v, ok := m[slice[i]]; ok {
				for _, str := range v {
					if str == slice[j] {
						return true
					}
				}
			}
			return i < j
		})
		if !slices.Equal(dst, slice) {
			count += aoc.Atoi([]byte(slice[(len(slice)+(^(len(slice)&1)&1))/2]))
		}

	}

	return aoc.Itoa(count), nil
}
