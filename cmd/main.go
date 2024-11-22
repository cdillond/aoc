package main

import (
	"aoc"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	y21 "aoc/2021"
)

var year int = 2021

func main() {
	var (
		day, part   int
		submit, get bool
	)
	flag.IntVar(&day, "d", 0, "day")
	flag.IntVar(&part, "p", 0, "part")
	flag.BoolVar(&submit, "submit", false, "submit answer to advent of code")
	flag.BoolVar(&get, "get", false, "download puzzle input")
	flag.Parse()

	if day < 1 || day > 25 {
		log.Fatalln("invalid day")
	}
	if part != 1 && part != 2 {
		log.Fatalln("invalid part")
	}

	var (
		res string
		err error
	)

	if get {
		path := "../" + strconv.Itoa(year) + "/d" + strconv.Itoa(day) + "/input.txt"
		if err = loadInput(day, year, path); err != nil {
			log.Fatalln(err)
		}
		return
	}

	if res, err = y21.Solve(day, part); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)

	if submit {
		if err := submitResult(day, year, part, res); err != nil {
			log.Fatalln(err)
		}
	}

}

func loadInput(day, year int, path string) error {
	token := os.Getenv("AOC_TOKEN")
	if token == "" {
		return errors.New("unable to obtain session cookie")
	}
	cli := aoc.NewClient(day, year, token)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return cli.GetInput(f)
}

func submitResult(day, year, part int, res string) error {
	token := os.Getenv("AOC_TOKEN")
	if token == "" {
		return errors.New("unable to obtain session cookie")
	}
	cli := aoc.NewClient(day, year, token)
	return cli.Submit(part, res, os.Stdout)
}
