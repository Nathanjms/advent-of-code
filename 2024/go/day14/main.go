package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

func getCurrentDirectory() string {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	return dirname
}

// Default Input path is current directory + example-input
var inputPath = filepath.Join(getCurrentDirectory(), "example-input")
var isUsingExample = true

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
		isUsingExample = false
	}

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

type robotStruct struct {
	position [2]int
	velocity [2]int
}

func partOne(contents []string) {
	robots := parseInput(contents)

	columns := 101
	rows := 103
	if isUsingExample {
		columns = 11
		rows = 7
	}

	// Now do a loop with a cache (as maybe will be cyclic?), of position and velocity vectors
	cacheMap := make(map[robotStruct][2]int)

	seconds := 100

	for i := 1; i <= seconds; i++ {
		for j, robot := range robots {
			// Shift it along!

			// First check if we've computed this before for this pos and vel?
			if val, ok := cacheMap[robot]; ok {
				robots[j].position = val
				continue
			}

			// Else let's do it and add to cache
			newPosition := [2]int{
				((robot.position[0] + robot.velocity[0]) + rows) % rows,
				(robot.position[1] + robot.velocity[1] + columns) % columns,
			}

			cacheMap[robot] = newPosition
			robots[j].position = newPosition

		}
		// displayMap(rows, columns, robots)
	}

	robotsPerQuadrant := [4]int{}
	for _, robot := range robots {
		// Determine quadrant;
		if robot.position[0] == (rows-1)/2 || robot.position[1] == (columns-1)/2 {
			// Skip in in middle - easy case!
			continue
		}

		// Might be a better way to do this but it'll do!
		if robot.position[0] < (rows-1)/2 && robot.position[1] < (columns-1)/2 {
			robotsPerQuadrant[0]++
		} else if robot.position[0] < (rows-1)/2 && robot.position[1] > (columns-1)/2 {
			robotsPerQuadrant[1]++
		} else if robot.position[0] > (rows-1)/2 && robot.position[1] < (columns-1)/2 {
			robotsPerQuadrant[2]++
		} else if robot.position[0] > (rows-1)/2 && robot.position[1] > (columns-1)/2 {
			robotsPerQuadrant[3]++
		}
	}

	multiplier := 1
	for i := 0; i < len(robotsPerQuadrant); i++ {
		multiplier *= robotsPerQuadrant[i]
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   14,
		Part:  1,
		Value: multiplier,
	})
}

func partTwo(contents []string) {
	robots := parseInput(contents)

	columns := 101
	rows := 103
	if isUsingExample {
		columns = 11
		rows = 7
	}

	cacheMap := make(map[robotStruct][2]int)

	seconds := 10_000 // Try 10,000

	// We'll output each second into a file ans then grep for when we have lots of robots in a line, hopefully this works
	fo, err := os.Create("output")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for i := 1; i <= seconds; i++ {
		for j, robot := range robots {
			if val, ok := cacheMap[robot]; ok {
				robots[j].position = val
				continue
			}

			newPosition := [2]int{
				((robot.position[0] + robot.velocity[0]) + rows) % rows,
				(robot.position[1] + robot.velocity[1] + columns) % columns,
			}

			cacheMap[robot] = newPosition
			robots[j].position = newPosition

		}

		grid := buildMap(rows, columns, robots)
		outputMap(&grid, fo, i)
	}
	// Not really sure how to do this one... maybe just output
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   14,
		Part:  2,
		Value: "Run grep on the output and look for the tree! `grep 00000000000000000000 ./output -C 10`",
	})
}

func parseInput(contents []string) []robotStruct {
	robots := make([]robotStruct, 0)
	re := regexp.MustCompile(`-?\d+`)
	for _, line := range contents {
		var posX, posY, velX, velY int

		matches := re.FindAllString(line, -1)

		posX, _ = strconv.Atoi(matches[0])
		posY, _ = strconv.Atoi(matches[1])
		velX, _ = strconv.Atoi(matches[2])
		velY, _ = strconv.Atoi(matches[3])

		// Note X and Y are flipped to match matrix format and not cartesian coords!
		robots = append(robots, robotStruct{
			position: [2]int{posY, posX},
			velocity: [2]int{velY, velX},
		})

	}

	return robots
}

func buildMap(rows int, columns int, robots []robotStruct) [][]byte {
	grid := make([][]byte, 0)

	for i := 0; i < rows; i++ {
		var builder bytes.Buffer
		for j := 0; j < columns; j++ {
			builder.WriteByte('.')
		}
		grid = append(grid, builder.Bytes())
	}

	for _, robot := range robots {
		grid[robot.position[0]][robot.position[1]] = '0'
	}

	return grid

}

func outputMap(grid *[][]byte, fo *os.File, seconds int) {
	fo.WriteString("Seconds: " + strconv.Itoa(seconds))
	fo.WriteString("\n")

	for _, line := range *grid {
		// write a chunk
		if _, err := fo.Write(line); err != nil {
			panic(err)
		}

		fo.WriteString("\n")
	}
}

func displayMap(rows int, columns int, robots []robotStruct) {
	grid := buildMap(rows, columns, robots)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()

}
