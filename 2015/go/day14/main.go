package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
)

func getCurrentDirectory() string {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	return dirname
}

// Default Input path is current directory + example-input
var inputPath = filepath.Join(getCurrentDirectory(), "example-input")

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	reindeers := parseInput(contents)

	time := 2503

	distanceByReindeer := make(map[string]int)
	for _, reindeer := range reindeers {
		// If can fly x km/h for y seconds, then rest for z seconds, then in y+z time, can fly x*y km
		timeCycles := math.Floor(float64(time) / float64(reindeer.flyTime+reindeer.restTime))

		// Distance is the number of full cycles * distance per cycle, plus any remainder
		distance := int(timeCycles) * reindeer.speed * reindeer.flyTime

		// Handle the remainder:
		remainder := time % (reindeer.flyTime + reindeer.restTime)
		flyTime := min(remainder, reindeer.flyTime) // Either the remainder, or the fly time, whichever is smaller

		distance += flyTime * reindeer.speed

		distanceByReindeer[reindeer.name] = distance
	}

	bestDistance := 0
	for _, distance := range distanceByReindeer {
		if distance > bestDistance {
			bestDistance = distance
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   14,
		Part:  1,
		Value: bestDistance,
	})
}

func partTwo(contents []string) {
	reindeers := parseInput(contents)

	time := 2503

	// Now we log the distance of each reindeer at each second. Probably ways to optimise this, but it's fast enough.
	distanceByReindeerPerSecond := make(map[string][]int)
	for _, reindeer := range reindeers {
		currentReindeerDistance := 0
		for second := 0; second < time; second++ {
			timeCycle := second % (reindeer.flyTime + reindeer.restTime)
			if timeCycle < reindeer.flyTime {
				currentReindeerDistance += reindeer.speed
			}
			distanceByReindeerPerSecond[reindeer.name] = append(distanceByReindeerPerSecond[reindeer.name], currentReindeerDistance)
		}
	}

	pointsPerReindeer := make(map[string]int)
	for second := 0; second < time; second++ {
		// Find the max distance for this second
		maxDistance := 0
		for _, reindeer := range reindeers {
			if distanceByReindeerPerSecond[reindeer.name][second] > maxDistance {
				maxDistance = distanceByReindeerPerSecond[reindeer.name][second]
			}
		}
		// Award a point to each reindeer who has this distance
		for _, reindeer := range reindeers {
			if distanceByReindeerPerSecond[reindeer.name][second] == maxDistance {
				pointsPerReindeer[reindeer.name]++
			}
		}
	}

	bestPoints := 0
	for _, points := range pointsPerReindeer {
		if points > bestPoints {
			bestPoints = points
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   14,
		Part:  2,
		Value: bestPoints,
	})
}

type reindeer struct {
	name     string
	speed    int
	flyTime  int
	restTime int
}

func parseInput(contents []string) []reindeer {
	// Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
	var reindeers []reindeer
	for _, line := range contents {
		var name string
		var speed, flyTime, restTime int
		_, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &flyTime, &restTime)
		if err != nil {
			panic(err)
		}
		reindeers = append(reindeers, reindeer{name, speed, flyTime, restTime})
	}
	return reindeers
}
