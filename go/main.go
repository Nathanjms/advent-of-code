package main

import (
	"flag"

	"github.com/nathanjms/advent-of-code-2023/go/day01"
)

func main() {
	dayPtr := flag.Int("day", 0, "Day to run")
	flag.Parse()

	switch *dayPtr {
	case 1:
		day01.PartOne()
	default:
		panic("Day not implemented")
	}
}
