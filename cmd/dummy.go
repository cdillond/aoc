//go:build !2024 && !2021

package main

const (
	_        uint8 = 1 << 8 // MUST NOT COMPILE
	year           = 0
	yearStr        = ""
	inputDir       = ""
)

func solve(day, part int, path string) (string, error) { return "", nil }
