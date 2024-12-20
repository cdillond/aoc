package d2

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := "2"
	have, err := Part1("../../inputs/2024/2.txt")
	if have != want {
		t.Logf("have: %s, want: %s\n", have, want)
		t.Fail()
	}
	if err != nil {
		t.Log(err)
	}
}

func BenchmarkPart1(b *testing.B) {
	for _ = range b.N {
		_, err := Part1("../../inputs/2024/2.txt")
		if err != nil {
			b.Logf(err.Error())
		}
	}
}

func TestPart2(t *testing.T) {
	want := "4"
	have, err := Part2("example.txt")
	if have != want {
		t.Logf("have: %s, want: %s\n", have, want)
		t.Fail()
	}
	if err != nil {
		t.Log(err)
	}
}

func BenchmarkPart2(b *testing.B) {
	for _ = range b.N {
		_, err := Part2("../../inputs/2024/2.txt")
		if err != nil {
			b.Logf(err.Error())
		}
	}
}
