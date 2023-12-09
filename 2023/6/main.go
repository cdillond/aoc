package main

import (
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
	m, err := ParseInput1(b)
	if err != nil {
		panic(err)
	}
	t, d, err := ParseInput2(b)
	if err != nil {
		panic(err)
	}

	fmt.Println("part 1: ", Part1(m))
	fmt.Println("part 2: ", Part2(t, d))
}

func Part1(races map[float64]float64) int {
	product := 1
	for time, distance := range races {
		product *= Solve(time, distance)
	}
	return product
}

func Part2(t, d float64) int {
	return Solve(t, d)
}

func Solve(t, d float64) int {
	/*
		t = total time
		x = time pressed
		y = total travelled
		y = x * (t - x)
		0 =  -x^2 +tx - y

		we can solve for x using the quadratic formula
		and then find the number of integers between solutions
	*/
	x1 := math.Floor((-t + math.Sqrt(t*t-4*-1*-d)) / (2 * -1))
	x2 := math.Floor((-t - math.Sqrt(t*t-4*-1*-d)) / (2 * -1))
	return int((math.Abs(x2 - x1)))

}

func ParseInput1(b []byte) (map[float64]float64, error) {
	lines := bytes.Split(b, []byte("\n"))
	if len(lines) != 2 {
		return nil, fmt.Errorf("invalid input 1")
	}
	if len(lines[0]) != len(lines[1]) {
		return nil, fmt.Errorf("invalid input 2")
	}
	times := []int{}
	distances := []int{}
	var num, distance int
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] >= '0' && lines[0][i] <= '9' {
			num *= 10
			num += int(lines[0][i] - '0')

			if i == len(lines[0])-1 || !(lines[0][i+1] >= '0' && lines[0][i+1] <= '9') {
				times = append(times, num)
			}
		} else {
			num = 0
		}

		if lines[1][i] >= '0' && lines[1][i] <= '9' {
			distance *= 10
			distance += int(lines[1][i] - '0')

			if i == len(lines[1])-1 || !(lines[1][i+1] >= '0' && lines[1][i+1] <= '9') {
				distances = append(distances, distance)
			}
		} else {
			distance = 0
		}
	}
	if len(distances) != len(times) {
		return nil, fmt.Errorf("invalid input 3")
	}
	m := make(map[float64]float64, len(times))
	for i, t := range times {
		m[float64(t)] = float64(distances[i])
	}
	return m, nil
}

func ParseInput2(b []byte) (float64, float64, error) {
	lines := bytes.Split(b, []byte("\n"))
	if len(lines) != 2 {
		return 0, 0, fmt.Errorf("invalid input 4")
	}
	if len(lines[0]) != len(lines[1]) {
		return 0, 0, fmt.Errorf("invalid input 5")
	}
	times, ok := bytes.CutPrefix(lines[0], []byte("Time:"))
	if !ok {
		return 0, 0, fmt.Errorf("invalid input 6")
	}
	distances, ok := bytes.CutPrefix(lines[1], []byte("Distance:"))
	if !ok {
		return 0, 0, fmt.Errorf("invalid input 7")
	}
	times = bytes.ReplaceAll(times, []byte{' '}, []byte{})
	distances = bytes.ReplaceAll(distances, []byte{' '}, []byte{})

	time, err := strconv.Atoi(string(times))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input 8")
	}
	distance, err := strconv.Atoi(string(distances))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid input 9")
	}
	return float64(time), float64(distance), nil
}
