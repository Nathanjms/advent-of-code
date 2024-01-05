package main

import (
	"2023/go/day01"
	"flag"
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
