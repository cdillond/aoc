package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"runtime/pprof"
	"time"

	"github.com/cdillond/aoc/cmd/client"
	"github.com/cdillond/aoc/cmd/html"

	// Update this import path when solving a new problem
	puzzle "github.com/cdillond/aoc/2024/d13"
)

func main() {
	//debug.SetGCPercent(-1)
	// flag variables
	var part, submit, get, clock, prof bool
	var customPath string

	flag.BoolVar(&submit, "submit", false, "submit answer to advent of code")
	flag.BoolVar(&get, "get", false, "download puzzle input")
	flag.BoolVar(&clock, "time", false, "measure and print the time taken to execute the solution function")
	flag.BoolVar(&prof, "prof", false, "run a profile")
	flag.StringVar(&customPath, "path", "", "use input at specified path")
	flag.Parse()

	if prof {
		out, err := os.Create("cpu.prof")
		if err != nil {
			log.Panicln(err)
		}
		defer out.Close()
		if err = pprof.StartCPUProfile(out); err != nil {
			log.Panicln(err)
		}
		defer pprof.StopCPUProfile()
	}

	part = flag.Arg(0) == "2"

	var err error
	input := customPath
	if input == "" {
		input = path.Join("..", puzzle.Year, "d"+puzzle.Day, "input.txt")
		// input = path.Join("..", "inputs", puzzle.Year, puzzle.Day+".txt")
	}

	if get {
		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			log.Panicln(err)
		}
		availAt, err := time.ParseInLocation("02-01-2006", puzzle.Day+"-12-"+puzzle.Year, loc)

		if err != nil {
			log.Panicln(err)
		}
		if availAt.After(time.Now()) {
			log.Println("waiting until ", availAt.String())
			t := time.NewTimer(time.Until(availAt))
			<-t.C
			t.Stop()
			log.Println("fetching input")
		}

		if err = loadInput(puzzle.Day, puzzle.Year, input); err != nil {
			log.Panicln(err)
		}
		return
	}

	var res string
	var start, stop time.Time
	switch part {
	case false:
		if clock {
			start = time.Now()
		}
		if res, err = puzzle.Part1(input); err != nil {
			log.Panicln(err)
		}
		if clock {
			stop = time.Now()
		}
	case true:
		if clock {
			start = time.Now()
		}
		if res, err = puzzle.Part2(input); err != nil {
			log.Panicln(err)
		}
		if clock {
			stop = time.Now()
		}
	}
	if clock && (!time.Now().IsZero()) {
		log.Println("time: ", stop.Sub(start))
	}

	if submit {
		log.Println("solution: ", res)
		if err := submitResult(partToStr(part), puzzle.Day, puzzle.Year, res); err != nil {
			log.Panicln(err)
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

	log.Println(string(out))
	return err
}
