package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type hand struct {
	cards  [5]int
	counts [13]int
	bid    int
	score
}

type score uint

const (
	hi score = iota
	pair
	two
	three
	full
	four
	five
)

var order1 = [13]byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var order2 = [13]byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(b, []byte("\n"))
	fmt.Println("part 1: ", Solve(rows, 1))
	fmt.Println("part 2: ", Solve(rows, 2))
}

func Solve(rows [][]byte, version int) int {
	hands := make([]hand, 0, len(rows))
	for _, row := range rows {
		if len(row) > 0 {
			hand, err := ParseRow(row, version)
			if err != nil {
				continue
			}
			hands = append(hands, hand)
		}
	}
	// sort in ascending order, first by score and then by the values of each successive card
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score != hands[j].score {
			return hands[i].score < hands[j].score
		} else {
			for n, c := range hands[i].cards {
				if c < hands[j].cards[n] {
					return true
				}
				if c > hands[j].cards[n] {
					return false
				}
			}
			return false
		}
	})
	var total int
	for i := range hands {
		total += hands[i].bid * (i + 1)
	}
	return total
}

func ParseRow(b []byte, version int) (hand, error) {
	h := hand{}
	i := 0
	switch version {
	case 1:
		for ; i < 5; i++ {
			card := GetRank1(b[i])
			h.cards[i] = card
			h.counts[card]++
		}
		h.score = CalcScore1(h)
	case 2:
		for ; i < 5; i++ {
			card := GetRank2(b[i])
			h.cards[i] = card
			h.counts[card]++
		}
		h.score = CalcScore2(h)
	}
	bid, err := strconv.Atoi(string(b[i+1:]))
	if err != nil {
		return hand{}, err
	}
	h.bid = bid
	return h, nil
}

func GetRank1(c byte) int {
	for i := range order1 {
		if c == order1[i] {
			return i
		}
	}
	return -1
}
func GetRank2(c byte) int {
	for i := range order2 {
		if c == order2[i] {
			return i
		}
	}
	return -1
}

func CalcScore1(h hand) score {
	var out score
	for i := range h.counts {
		switch h.counts[i] {
		case 5:
			out = five
		case 4:
			out = four
		case 3:
			if out == pair {
				out = full
			} else {
				out = three
			}
		case 2:
			if out == three {
				out = full
			} else if out == pair {
				out = two
			} else {
				out = pair
			}
		}
	}
	return out
}

func CalcScore2(h hand) score {
	score := CalcScore1(h)
	jcount := h.counts[0]
	if jcount == 0 {
		return score
	}
	switch score {
	case five, four, full:
		return five
	case three:
		// if jcount == 2, then it's a full house, which has already been handled
		// if jcount == 3 or jcount == 1, the result is the same (four of a kind)
		return four
	case two:
		if jcount == 2 {
			return four
		}
		if jcount == 1 {
			return full
		}
	case pair:
		// three of a kind beats two-pair
		return three
	case hi:
		return pair
	}
	return score
}
