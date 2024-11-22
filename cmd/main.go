package main

import (
	"aoc"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// these values will depend on the build tags
var (
	base string
	year int
)

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

	path := base + "d" + strconv.Itoa(day) + "/input.txt"

	var (
		res string
		err error
	)

	if get {
		if err = loadInput(day, year, path); err != nil {
			log.Fatalln(err)
		}
		return
	}

	solution := aoc.Problems[day-1][part-1]
	if solution != nil {
		res, err = solution(path)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("the requested function has not yet been added to aoc.Problems.")
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
		log.Fatalln("unable to obtain session cookie")
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
