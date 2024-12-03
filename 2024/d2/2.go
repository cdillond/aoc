package d2

import (
	"bufio"
	"os"

	"github.com/cdillond/aoc"
)

const (
	Day  = "2"
	Year = "2024"
)

type buffer struct {
	b    []byte
	a, z int
}

func (b *buffer) next() (n int, done bool) {
	if b.z >= len(b.b) {
		return n, true
	}
	// skip any leading spaces
	for b.z < len(b.b) && b.b[b.z] == ' ' {
		b.z++
		b.a = b.z
	}
	b.z++
	// consume the chars comprising the next integer
	for b.z < len(b.b) && b.b[b.z] != ' ' {
		b.z++
	}
	n = aoc.Atoi(b.b[b.a:b.z])
	// prepare the buffer for the next call
	b.a = b.z + 1
	b.z = b.a
	return n, false
}

func Part1(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	var count int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var buf buffer
		buf.b = scanner.Bytes()

		last, _ := buf.next()
		cur, _ := buf.next()
		increasing := (cur - last) > 0
		var done, unsafe bool
		for !done {
			dif := cur - last
			if unsafe = (increasing && dif < 0) || (!increasing && dif > 0); unsafe {
				break
			}
			dif = aoc.Abs(dif)
			if unsafe = (dif < 1) || (dif > 3); unsafe {
				break
			}
			last = cur
			cur, done = buf.next()
		}

		if !unsafe {
			count++
		}

	}

	return aoc.Itoa(count), nil
}

func check2(skip int, nums []int) (int, bool) {
	if len(nums) <= 1 {
		return -1, true
	}
	const (
		notSet   = 0
		increase = 1
		decrease = 2
	)
	var direction, dif int
	prev := -1
	for i, cur := range nums {
		if i == skip {
			continue
		}
		// check direction
		switch direction {
		case notSet:
			j := i - 1
			if j == skip {
				j--
			}
			if j >= 0 {
				switch {
				case cur > nums[j]:
					direction = increase
				case cur < nums[j]:
					direction = decrease
				default:
					return i, false
				}
			}
		case increase:
			if cur <= nums[prev] {
				return i, false
			}
		case decrease:
			if cur >= nums[prev] {
				return i, false
			}
		}
		if prev < 0 {
			prev = 0
			continue
		}

		dif = aoc.Abs(cur - nums[prev])
		if (dif < 1) || (dif > 3) {
			return i, false
		}

		prev = i

	}

	return -1, true
}

func check(nums []int) (int, bool) {
	if len(nums) <= 1 {
		return -1, true
	}
	last := nums[0]
	cur := nums[1]

	increasing := (cur - last) > 0
	i := 1
	for {
		dif := cur - last
		if (increasing && dif < 0) || (!increasing && dif > 0) {
			return i, false
		}
		dif = aoc.Abs(dif)
		if (dif < 1) || (dif > 3) {
			return i, false
		}
		i++
		if i >= len(nums) {
			break
		}

		last = cur
		cur = nums[i]
	}

	return -1, true
}

func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	var count int
	nums := make([]int, 0, 32)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		nums = nums[:0]
		buf := buffer{b: scanner.Bytes()}
		for {
			n, done := buf.next()
			if done {
				break
			}
			nums = append(nums, n)
		}
		skip, safe := check2(-1, nums)
		for skip > -1 {
			_, safe = check2(skip, nums)
			skip--
		}
		if safe {
			count++
		}

	}
	return aoc.Itoa(count), nil
}

/*
func Part2(path string) (res string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return res, err
	}
	defer f.Close()

	var count int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		fields := bytes.Fields(scanner.Bytes())
		nums := make([]int, len(fields))
		for i, b := range fields {
			nums[i] = aoc.Atoi(b)
		}
		skip, safe := check(nums)
		if safe {
			count++
			continue
		}
		tmp := make([]int, len(nums)-1)
		n := copy(tmp, nums[:skip])
		copy(tmp[n:], nums[skip+1:])
		_, safe = check(tmp)
		if safe {
			count++
			continue
		}
		skip--
		if skip < 0 {
			continue
		}
		n = copy(tmp, nums[:skip])
		copy(tmp[n:], nums[skip+1:])
		_, safe = check(tmp)
		if safe {
			count++
			continue
		}
		skip--
		if skip < 0 {
			continue
		}
		n = copy(tmp, nums[:skip])
		copy(tmp[n:], nums[skip+1:])
		_, safe = check(tmp)
		if safe {
			count++
		}
	}

	return aoc.Itoa(count), nil
}
*/
