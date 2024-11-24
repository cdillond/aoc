package main

import (
	"io/fs"
	"os"
	"path"
	"strings"
)

func mkTemplate(day string) error {
	dirPath := "../" + yearStr + "/d" + day
	err := os.MkdirAll(dirPath, fs.ModePerm)
	if err != nil {
		return err
	}
	var f, ftest *os.File
	if f, err = os.Create(path.Join(dirPath, day+".go")); err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(srcTemplate(day)); err != nil {
		return err
	}
	if ftest, err = os.Create(path.Join(dirPath, day+"_test.go")); err != nil {
		return err
	}
	defer ftest.Close()
	if _, err = ftest.WriteString(srcExampleTemplate(day)); err != nil {
		return err
	}
	return nil
}

func srcTemplate(day string) string {
	bldr := new(strings.Builder)
	bldr.WriteString("package d" + day + "\n\n")
	bldr.WriteString("func Part1(path string) (res string, err error) { return }\n\n")
	bldr.WriteString("func Part2(path string) (res string, err error) { return }\n\n")
	return bldr.String()
}

func srcExampleTemplate(day string) string {
	bldr := new(strings.Builder)
	bldr.WriteString("package d" + day + "\n\n")
	bldr.WriteString("import (\n\t\"testing\"\n)\n\n")
	bldr.WriteString(`func TestPart1(t *testing.T) {
	want := ""
	have, err := Part1("example.txt")
	if have != want {
		t.Logf("have: %s, want: %s\n", have, want)
		t.Fail()
	}
	if err != nil {
		t.Log(err)
	}
}

func TestPart2(t *testing.T) {
	want := ""
	have, err := Part2("example.txt")
	if have != want {
		t.Logf("have: %s, want: %s\n", have, want)
		t.Fail()
	}
	if err != nil {
		t.Log(err)
	}
}`)
	bldr.WriteString("\n\n")
	return bldr.String()
}
