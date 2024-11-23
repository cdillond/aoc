package d1

import (
	"aoc"
	"testing"
)

func TestPart1(t *testing.T) {
	examples := []aoc.Example{
		{Path: "example.txt", Want: ""},
	}

	for _, ex := range examples {
		have, err := Part1(ex.Path)
		if have != ex.Want {
			t.Logf("have: %s, want: %s\n", have, ex.Want)
			t.Fail()
		}
		if err != nil {
			t.Log(err)
		}
	}
}

func TestPart2(t *testing.T) {
	examples := []aoc.Example{
		{Path: "example.txt", Want: ""},
	}

	for _, ex := range examples {
		have, err := Part2(ex.Path)
		if have != ex.Want {
			t.Logf("have: %s, want: %s\n", have, ex.Want)
			t.Fail()
		}
		if err != nil {
			t.Log(err)
		}
	}
}
