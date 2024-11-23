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

// Set this var (as well as the correct import) prior to building.
var year int = 2021

func main() {
	var (
		day, part       int
		submit, get, mk bool
	)
	flag.IntVar(&day, "d", 0, "day")
	flag.IntVar(&part, "p", 0, "part")
	flag.BoolVar(&submit, "submit", false, "submit answer to advent of code")
	flag.BoolVar(&get, "get", false, "download puzzle input")
	flag.BoolVar(&mk, "mk", false, "create a new puzzle directory")
	flag.Parse()

	if day < 1 || day > 25 {
		log.Fatalln("invalid day")
	}
	path := "../inputs/" + strconv.Itoa(year) + "/" + strconv.Itoa(day) + ".txt"
	var err error
	if mk {
		if err = mkTemplate(day); err != nil {
			log.Fatalln(err)
		}
		return
	}

	if get {
		if err = loadInput(day, year, path); err != nil {
			log.Fatalln(err)
		}
		return
	}

	if part != 1 && part != 2 {
		log.Fatalln("invalid part")
	}

	var res string
	if res, err = y21.Solve(day, part, path); err != nil {
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
