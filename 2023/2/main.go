package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type round struct {
	red, green, blue int
}

type game []round

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sum1, sum2 int
	for id := 1; scanner.Scan(); id++ {
		b := scanner.Bytes()
		colon := bytes.IndexByte(b, ':')
		if colon < 0 {
			break
		}
		rounds := bytes.Split(b[colon+1:], []byte{';'})
		g := make(game, len(rounds))
		if len(rounds) == 0 {
			continue
		}
		for i := 0; i < len(rounds); i++ {
			g[i] = ParseRound(rounds[i])
		}

		m := MaxVals(g)
		// perform check for part 1
		if m.red <= 12 && m.green <= 13 && m.blue <= 14 {
			sum1 += id
		}
		// sum the set's "power"
		sum2 += m.red * m.green * m.blue
	}
	fmt.Println("part 1: ", sum1)
	fmt.Println("part 2: ", sum2)
}

func ParseRound(b []byte) round {
	draws := bytes.Split(b, []byte{','})
	var r round
	for i := range draws {
		var n int
		var s string
		fmt.Sscanf(string(draws[i]), "%d %s", &n, &s)
		switch s {
		case "red":
			r.red = n
		case "green":
			r.green = n
		case "blue":
			r.blue = n
		}
	}
	return r
}

func MaxVals(g game) round {
	var red, green, blue int
	for _, r := range g {
		red = max(red, r.red)
		green = max(green, r.green)
		blue = max(blue, r.blue)
	}
	return round{red: red, green: green, blue: blue}
}
