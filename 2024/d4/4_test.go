package d4

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := "18"
	have, err := Part1("example.txt")
	if have != want {
		t.Logf("have: %s, want: %s\n", have, want)
		t.Fail()
	}
	if err != nil {
		t.Log(err)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Part1("../../inputs/2024/4.txt")
		if err != nil {
			b.Logf(err.Error())
		}
	}
}

func TestPart2(t *testing.T) {
	want := "9"
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
	for i := 0; i < b.N; i++ {
		_, err := Part1("../../inputs/2024/4.txt")
		if err != nil {
			b.Logf(err.Error())
		}
	}
}
