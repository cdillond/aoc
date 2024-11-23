package main

import (
	"io/fs"
	"os"
	"path"
	"strconv"
	"strings"
)

func mkTemplate(day int) error {
	dirPath := "../" + strconv.Itoa(year) + "/d" + strconv.Itoa(day)
	err := os.MkdirAll(dirPath, fs.ModePerm)
	if err != nil {
		return err
	}
	var f, ftest *os.File
	if f, err = os.Create(path.Join(dirPath, strconv.Itoa(day)+".go")); err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(srcTemplate(day)); err != nil {
		return err
	}
	if ftest, err = os.Create(path.Join(dirPath, strconv.Itoa(day)+"_test.go")); err != nil {
		return err
	}
	defer ftest.Close()
	if _, err = ftest.WriteString(srcExampleTemplate(day)); err != nil {
		return err
	}
	return nil
}

func srcTemplate(day int) string {
	bldr := new(strings.Builder)
	bldr.WriteString("package d" + strconv.Itoa(day) + "\n\n")
	bldr.WriteString("func Part1(path string) (res string, err error) { return }\n\n")
	bldr.WriteString("func Part2(path string) (res string, err error) { return }\n\n")
	return bldr.String()
}

func srcExampleTemplate(day int) string {
	bldr := new(strings.Builder)
	bldr.WriteString("package d" + strconv.Itoa(day) + "\n\n")
	bldr.WriteString("import (\n\t\"aoc\"\n\t\"testing\"\n)\n\n")
	bldr.WriteString(`func TestPart1(t *testing.T) {
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
}`)
	bldr.WriteString("\n\n")
	bldr.WriteString(`func TestPart2(t *testing.T) {
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
}`)
	bldr.WriteString("\n\n")
	return bldr.String()
}
