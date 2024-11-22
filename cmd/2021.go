//go:build 2021

package main

import (
	"aoc"
	"aoc/2021/d1"
	"aoc/2021/d2"
	"aoc/2021/d3"
)

func init() {
	base = "../2021/"
	year = 2021

	aoc.Problems[0][0] = d1.Part1
	aoc.Problems[0][1] = d1.Part2

	aoc.Problems[1][0] = d2.Part1
	aoc.Problems[1][1] = d2.Part2

	aoc.Problems[2][0] = d3.Part1
	aoc.Problems[2][1] = d3.Part2

	/*
		aoc.Problems[3][0] = d4.Part1
		aoc.Problems[3][1] = d4.Part2

		aoc.Problems[4][0] = d5.Part1
		aoc.Problems[4][1] = d5.Part2

		aoc.Problems[5][0] = d6.Part1
		aoc.Problems[5][1] = d6.Part2

		aoc.Problems[6][0] = d7.Part1
		aoc.Problems[6][1] = d7.Part2

		aoc.Problems[7][0] = d8.Part1
		aoc.Problems[7][1] = d8.Part2

		aoc.Problems[8][0] = d9.Part1
		aoc.Problems[8][1] = d9.Part2

		aoc.Problems[9][0] = d10.Part1
		aoc.Problems[9][1] = d10.Part2

		aoc.Problems[10][0] = d11.Part1
		aoc.Problems[10][1] = d11.Part2

		aoc.Problems[11][0] = d12.Part1
		aoc.Problems[11][1] = d12.Part2

		aoc.Problems[12][0] = d13.Part1
		aoc.Problems[12][1] = d13.Part2

		aoc.Problems[13][0] = d14.Part1
		aoc.Problems[13][1] = d14.Part2

		aoc.Problems[14][0] = d15.Part1
		aoc.Problems[14][1] = d15.Part2

		aoc.Problems[15][0] = d16.Part1
		aoc.Problems[15][1] = d16.Part2

		aoc.Problems[16][0] = d17.Part1
		aoc.Problems[16][1] = d17.Part2

		aoc.Problems[17][0] = d18.Part1
		aoc.Problems[17][1] = d18.Part2

		aoc.Problems[18][0] = d19.Part1
		aoc.Problems[18][1] = d19.Part2

		aoc.Problems[19][0] = d20.Part1
		aoc.Problems[19][1] = d20.Part2

		aoc.Problems[20][0] = d21.Part1
		aoc.Problems[20][1] = d21.Part2

		aoc.Problems[21][0] = d22.Part1
		aoc.Problems[21][1] = d22.Part2

		aoc.Problems[22][0] = d23.Part1
		aoc.Problems[22][1] = d23.Part2

		aoc.Problems[23][0] = d24.Part1
		aoc.Problems[23][1] = d24.Part2

		aoc.Problems[24][0] = d25.Part1
		aoc.Problems[24][1] = d25.Part2
	*/
}
