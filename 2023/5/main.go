package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type entry struct {
	src, dst, len int64
}

type seed struct {
	start, end int64 // [start, end)
}

func MapTo(n int64, m []entry) int64 {
	for i := range m {
		if n >= m[i].src && n < m[i].src+m[i].len {
			return n + m[i].dst - m[i].src
		}
	}
	return n
}

// The inverse of MapTo.
func MapFrom(n int64, m []entry) int64 {
	for i := range m {
		if n >= m[i].dst && n < m[i].dst+m[i].len {
			return n + m[i].src - m[i].dst
		}
	}
	return n
}

func main() {
	// prepare maps
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	split := bytes.Split(b, []byte("\n\n"))
	var maps [][]entry
	for i := range split {
		scanner := bufio.NewScanner(bytes.NewBuffer(split[i]))
		scanner.Scan() // discard first row
		var m []entry
		for scanner.Scan() {
			var dst, src, len int64
			n, err := fmt.Sscanf(scanner.Text(), "%d %d %d", &dst, &src, &len)
			if n != 3 || err != nil {
				continue
			}
			m = append(m, entry{dst: dst, src: src, len: len})
		}
		// sort by src range
		sort.Slice(m, func(i, j int) bool {
			return m[i].src < m[j].src
		})

		// fill in the gaps
		var gaps []entry
		var end int64
		for j := 0; j < len(m); j++ {
			gap := m[j].src - end
			if gap != 0 {
				gaps = append(gaps, entry{src: end, dst: end, len: gap})
			}
			end = m[j].src + m[j].len
		}
		m = append(m, gaps...)

		// sort by src range again, though, not really needed
		sort.Slice(m, func(i, j int) bool {
			return m[i].src < m[j].src
		})
		maps = append(maps, m)
	}

	// parse seeds
	seedBytes, err := os.ReadFile("seeds.txt")
	if err != nil {
		panic(err)
	}
	splitSeeds := bytes.Split(seedBytes, []byte(" "))
	seeds := make([]int64, len(splitSeeds))
	for i := range splitSeeds {
		n, err := strconv.ParseInt(string(splitSeeds[i]), 10, 64)
		if err != nil {
			continue
		}
		seeds[i] = n
	}

	fmt.Println("part 1: ", Part1(maps, seeds))
	fmt.Println("part 2: ", Part2(maps, ToRange(seeds)))

}

func Part1(maps [][]entry, seeds []int64) int64 {
	outcomes := make([]int64, len(seeds))
	// transform each seed by each successive map
	for i, s := range seeds {
		for j := range maps {
			s = MapTo(s, maps[j])
		}
		outcomes[i] = s
	}
	// return the smallest value in outcomes
	return sliceMin(outcomes)
}

func Part2(maps [][]entry, seeds []seed) int64 {
	// sort the final map by dst, and then work backwards
	last := maps[len(maps)-1]
	sort.Slice(last, func(i, j int) bool { return last[i].dst < last[j].dst })

	// find the lowest dst that maps to a value in a seed range
	for _, e := range last {
		for i := e.dst; i < e.dst+e.len; i++ {
			val := int64(i)
			for j := len(maps) - 1; j >= 0; j-- {
				val = MapFrom(val, maps[j])
			}
			for _, r := range seeds {
				if InRange(val, r) {
					return i // return the dst, not the val
				}
			}
		}

	}
	return -1
}

func sliceMin(s []int64) int64 {
	res := int64(1<<63 - 1)
	for i := range s {
		res = min(res, s[i])
	}
	return res
}

func ToRange(seeds []int64) []seed {
	res := make([]seed, len(seeds)/2)
	for i := 0; i < len(seeds)-2; i += 2 {
		res[i/2] = seed{start: seeds[i], end: seeds[i] + seeds[i+1]}
	}
	return res
}

func InRange(n int64, s seed) bool {
	return n >= s.start && n < s.end
}
