package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("part 1: %d\n", Part1(b))
	fmt.Printf("part 2: %d\n", Part2(b))
}

func Part1(b []byte) int {
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	var sum float64
	for scanner.Scan() {
		line := scanner.Bytes()
		_, line, found := bytes.Cut(line, []byte(": "))
		if !found {
			continue
		}
		// parse the line
		winning := make(map[int]struct{})
		tmp := 0
		i := 0
		// all digits are white-space padded to two bytes
		for ; i < bytes.IndexByte(line, '|'); i++ {
			// these bytes are either whitespace or digits
			if line[i] == ' ' {
				continue
			}

			tmp *= 10
			tmp += int(line[i] - '0')

			// the number is either 1 or two digits long
			// peek to see if the number continues or not
			if !(line[i+1] >= '0' && line[i+1] <= '9') {
				// if not, record the final number
				winning[tmp] = struct{}{}
				tmp = 0
			}
		}

		var count int
		i += 2 // skip over | and the space directly after |
		for ; i < len(line); i++ {
			if line[i] == ' ' {
				continue
			}
			tmp *= 10
			tmp += int(line[i] - '0')
			// the number is either 1 or two digits long
			// peek to see if the number continues or not
			if i == len(line)-1 || !(line[i+1] >= '0' && line[i+1] <= '9') {
				if _, ok := winning[tmp]; ok {
					count++
				}
				tmp = 0
			}
		}
		if count > 0 {
			// starts at 1 point for 1 match, and then doubles
			sum += math.Pow(2, float64(count-1))
		}
	}
	return int(sum)
}

func Part2(b []byte) int {
	cards := []Card{}
	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	for scanner.Scan() {
		line := scanner.Bytes()
		c, err := ParseCard(line)
		if err != nil {
			continue
		}
		cards = append(cards, c)
	}

	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < i+1+cards[i].Count && j < len(cards); j++ {
			cards[j].Multiplier += cards[i].Multiplier
		}
	}
	var sum int
	for i := 0; i < len(cards); i++ {
		sum += cards[i].Multiplier
	}
	return sum
}

type Card struct {
	Winning    map[int]struct{}
	Have       map[int]struct{}
	Count      int
	Multiplier int
}

func ParseCard(b []byte) (Card, error) {
	_, line, found := bytes.Cut(b, []byte(": "))
	if !found {
		return Card{}, fmt.Errorf("problem parsing card")
	}
	split := bytes.Split(line, []byte("|"))
	if len(split) != 2 {
		return Card{}, fmt.Errorf("problem parsing card")
	}
	winning := split[0]
	have := split[1]

	winslice := bytes.Split(winning, []byte(" "))
	haveslice := bytes.Split(have, []byte(" "))

	winmap := make(map[int]struct{})
	havemap := make(map[int]struct{})

	for _, val := range winslice {
		if len(val) > 0 {
			i, err := strconv.Atoi(string(val))
			if err != nil {
				continue
			}
			winmap[i] = struct{}{}
		}
	}
	var count int
	for _, val := range haveslice {
		if len(val) > 0 {
			i, err := strconv.Atoi(string(val))
			if err != nil {
				continue
			}
			havemap[i] = struct{}{}
			if _, ok := winmap[i]; ok {
				count++
			}
		}
	}
	return Card{
		Winning:    winmap,
		Have:       havemap,
		Count:      count,
		Multiplier: 1,
	}, nil

}
