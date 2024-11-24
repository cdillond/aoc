This repository includes solutions to Advent of Code puzzles as well as a client for downloading puzzle inputs.

## Building
The solutions are designed to be independently imported by external go programs, but this repository also includes the source code (in the `cmd` directory) for a program that provides a convenient CLI for running them. The `cmd/cmd` program must be built with a build tag specifying the year, e.g.:
```
git clone git@github.com:cdillond/aoc.git
cd aoc/cmd
go build -tags 2024
```
  
Otherwise, the program will fail to compile. Currently, the only supported tags are 2021 and 2024, but it should be quite simple to add more.

If you modify the existing repository, you must run `go generate` before recompiling `cmd/cmd`.

## Running
Before using the client to download inputs and/or submit answers, you must first set the value of an environment variable named `AOC_SESSION` to the value of the `session` cookie from the Advent of Code website.
```
export AOC_SESSION={your_session_cookie_value}
./cmd -get -d 1
./cmd -submit -d 1 -p 2
```
Puzzle inputs are stored at the path `aoc/inputs/$YEAR/$DAY.txt`. Example inputs for a given day should be stored at `aoc/$YEAR/d$DAY/example.txt`.

Templates for solution packages can be run using the `mk` option. If built with the 2024 tag, the command `./cmd -mk -d 1` will generate files (and any necessary directories) at `aoc/2024/d1/1.go` and `aoc/2024/d1/1_test.go`.