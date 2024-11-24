//go:build 2024

package main

import (
	y24 "aoc/2024"
)

const (
	year     = 2024
	yearStr  = "2024"
	inputDir = "../inputs/2024/"
)

func solve(day, part int, path string) (string, error) { return y24.Solve(day, part, path) }
