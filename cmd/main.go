package main

import (
	"aoc/cmd/client"
	"aoc/cmd/html"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	// Update this import path when solving a new problem
	prob "aoc/2021/d5"
)

func main() {
	// flag variables
	var part, submit, get bool
	flag.BoolVar(&part, "t", false, "part two")
	flag.BoolVar(&submit, "submit", false, "submit answer to advent of code")
	flag.BoolVar(&get, "get", false, "download puzzle input")
	flag.Parse()

	var err error

	input := path.Join("..", "inputs", prob.Year, prob.Day+".txt")

	if get {
		if err = loadInput(prob.Day, prob.Year, input); err != nil {
			log.Fatalln(err)
		}
		return
	}

	var res string

	switch part {
	case false:
		if res, err = prob.Part1(input); err != nil {
			log.Fatalln(err)
		}
	case true:
		if res, err = prob.Part2(input); err != nil {
			log.Fatalln(err)
		}
	}

	if submit {
		log.Println("solution: ", res)
		if err := submitResult(partToStr(part), prob.Day, prob.Year, res); err != nil {
			log.Fatalln(err)
		}
		return
	}

	fmt.Println(res)
}

func partToStr(part bool) string {
	if part {
		return "2"
	}
	return "1"
}

func loadInput(day, year, input string) error {
	cli, err := client.New(day, year)
	if err != nil {
		return err
	}
	f, err := os.Create(input)
	if err != nil {
		return err
	}
	defer f.Close()

	return cli.GetInput(f)
}

func submitResult(part, day, year, res string) error {
	cli, err := client.New(day, year)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = cli.Submit(part, res, buf); err != nil {
		return err
	}

	out, err := html.Response(buf)

	fmt.Println(string(out))
	return err
}
