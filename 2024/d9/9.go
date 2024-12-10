package d9

import (
	"os"
	"strings"

	"github.com/cdillond/aoc"
)

const (
	Day  = "9"
	Year = "2024"
)

type file struct{ length, value int }

type Range struct {
	length    int
	numEmpty  int
	files     []file
	origEmpty bool
}

type Range2 struct {
	length    int
	numEmpty  int
	files     []file
	origEmpty bool
	del       bool
}

func Transfer(a, b *Range) {
	l := len(b.files)
	for i := 0; a.numEmpty != 0 && i < l; i++ {
		f := b.files[i]
		if f.length != 0 {
			toTransfer := min(a.numEmpty, f.length)
			a.files = append(a.files, file{length: toTransfer, value: f.value})
			a.numEmpty -= toTransfer
			b.files[i].length -= toTransfer
			b.numEmpty += toTransfer
		}
		if b.files[i].length > 0 {
			break
		}
	}
}

func Consume(a, b *Range) {
	a.numEmpty -= b.length
	b.numEmpty = b.length
	a.files = append(a.files, b.files[0])
	b.files[0].length = 0

}

func (r Range) Checksum(start, sum int) (int, int) {
	for _, f := range r.files {
		for range f.length {
			sum += start * f.value
			start++
		}
	}
	return start, sum
}

func (r Range2) Checksum2(start, sum int) (int, int) {
	if r.del {
		start += r.length
	} else {
		for _, f := range r.files {
			for range f.length {
				sum += start * f.value
				start++
			}
		}
		start += r.numEmpty
	}

	return start, sum
}

func (r Range) String() string {
	bldr := new(strings.Builder)
	for _, f := range r.files {
		for range f.length {
			bldr.Write([]byte(aoc.Itoa(f.value)))
		}
	}
	for range r.numEmpty {
		bldr.WriteByte('.')
	}
	return bldr.String()
}

func (r Range2) String() string {
	bldr := new(strings.Builder)
	if r.del {
		for range r.length {
			bldr.WriteByte('.')
		}
	} else {
		for _, f := range r.files {
			for range f.length {
				bldr.Write([]byte(aoc.Itoa(f.value)))
			}
		}
		for range r.numEmpty {
			bldr.WriteByte('.')
		}

	}

	return bldr.String()
}

func Part1(path string) (res string, err error) {
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return res, err
	}
	var disk []Range
	var idx int
	for i, c := range b {
		if i&1 != 0 {
			disk = append(disk, Range{length: int(c - '0'), numEmpty: int(c - '0')})
		} else {
			disk = append(disk, Range{length: int(c - '0'), files: []file{{length: int(c - '0'), value: idx}}})
			idx++
		}
	}

	p1, p2 := 0, len(disk)-1
	for p1 < p2 {
		for disk[p1].numEmpty < 1 {
			p1++
		}
		for disk[p2].numEmpty >= disk[p2].length {
			p2--
		}
		Transfer(&disk[p1], &disk[p2])
	}

	var index, checksum int
	for _, r := range disk {
		index, checksum = r.Checksum(index, checksum)
	}
	return aoc.Itoa(checksum), nil
}

func Part2(path string) (res string, err error) {
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return res, err
	}
	var disk []Range2
	var idx int
	for i, c := range b {
		if i&1 != 0 {
			disk = append(disk, Range2{length: int(c - '0'), numEmpty: int(c - '0'), origEmpty: true})
		} else {
			disk = append(disk, Range2{length: int(c - '0'), files: []file{{length: int(c - '0'), value: idx}}})
			idx++
		}
	}
	//fmt.Println(disk)

	for p2 := len(disk) - 1; p2 > 0; p2-- {
		if disk[p2].origEmpty == false {
			for p1 := 0; p1 < p2; p1++ {
				if disk[p1].origEmpty && disk[p1].numEmpty >= disk[p2].length {
					a0 := disk[p1]
					a1 := disk[p2]
					a1.del = true
					a0.numEmpty -= a1.length
					a0.files = append(a0.files, a1.files...)

					disk[p1] = a0
					disk[p2] = a1
					break
				}
			}
		}
	}

	var index, checksum int
	//fmt.Println(disk)
	for _, r := range disk {
		index, checksum = r.Checksum2(index, checksum)
	}
	return aoc.Itoa(checksum), nil

}
