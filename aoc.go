package aoc

import "errors"

//go:generate go run gen.go

var ErrUndefined = errors.New("no solution found")

type Example struct {
	Path string
	Want string
}
