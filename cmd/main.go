package main

import (
	"aoc/cmd/client"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// flag variables
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
	dayStr := strconv.Itoa(day)
	path := inputDir + dayStr + ".txt"
	var err error
	if mk {
		if err = mkTemplate(dayStr); err != nil {
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
	if res, err = solve(day, part, path); err != nil {
		log.Fatalln(err)
	}

	if submit {
		log.Println("solution: ", res)
		if err := submitResult(day, year, part, res); err != nil {
			log.Fatalln(err)
		}
		return
	}

	fmt.Println(res)
}

func newClient(day, year int) (cli client.Client, err error) {
	token := os.Getenv("AOC_SESSION")
	if token == "" {
		return cli, errors.New("unable to obtain session cookie")
	}
	return client.New(day, year, token), nil
}

func loadInput(day, year int, path string) error {
	cli, err := newClient(day, year)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return cli.GetInput(f)
}

func submitResult(day, year, part int, res string) error {
	cli, err := newClient(day, year)
	if err != nil {
		return err
	}
	return cli.Submit(part, res, os.Stdout)
}
