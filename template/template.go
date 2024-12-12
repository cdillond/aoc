package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path"
	"strconv"
	"text/template"
	"time"
)

func main() {
	now := time.Now()
	maxYear := now.Year()
	if now.Month() < time.December {
		maxYear--
	}
	year := flag.Int("y", maxYear, "puzzle year [2015, "+strconv.Itoa(maxYear)+"]")
	day := flag.Int("d", 0, "puzzle day [1, 25]")
	flag.Parse()
	if *day < 1 || *day > 25 {
		flag.Usage()
		os.Exit(1)
	}
	if *year < 2015 || *year > maxYear {
		flag.Usage()
		os.Exit(1)
	}

	if err := MkTemplate(strconv.Itoa(*day), strconv.Itoa(*year)); err != nil {
		log.Fatalln(err)
	}
}

func MkTemplate(day, year string) error {
	const (
		f1 = "day.go.tmpl"
		f2 = "day_test.go.tmpl"
		f3 = "input.txt"
	)
	t, err := template.ParseFiles(f1, f2)
	if err != nil {
		return err
	}

	dirPath := path.Join("..", year, "d"+day)
	if err = os.MkdirAll(dirPath, fs.ModePerm); err != nil {
		return err
	}

	var f, ftest, fin *os.File

	if f, err = os.Create(path.Join(dirPath, day+".go")); err != nil {
		return err
	}
	defer f.Close()
	data := struct {
		PackageName string
		Year        string
		Day         string
	}{
		PackageName: "d" + day,
		Year:        year,
		Day:         day,
	}
	if err = t.Lookup(f1).Execute(f, data); err != nil {
		return err
	}

	if ftest, err = os.Create(path.Join(dirPath, day+"_test.go")); err != nil {
		return err
	}
	defer ftest.Close()

	if err = t.Lookup(f2).Execute(ftest, data); err != nil {
		return err
	}

	if fin, err = os.Create(path.Join(dirPath, "input.txt")); err != nil {
		return err
	}
	defer fin.Close()

	return nil
}
