package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
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
	ingredients := parseInput(contents)

	bestScore := 0

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			fmt.Println(i, j)
		}
	}

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				for l := 0; l <= 100-i-j-k; l++ {
					// Check that the total is 100:
					if i+j+k+l != 100 {
						continue
					}

					// Calculate the score, 0 if negative:
					capacity := max(ingredients[0].capacity*i+ingredients[1].capacity*j+ingredients[2].capacity*k+ingredients[3].capacity*l, 0)
					durability := max(ingredients[0].durability*i+ingredients[1].durability*j+ingredients[2].durability*k+ingredients[3].durability*l, 0)
					flavor := max(ingredients[0].flavor*i+ingredients[1].flavor*j+ingredients[2].flavor*k+ingredients[3].flavor*l, 0)
					texture := max(ingredients[0].texture*i+ingredients[1].texture*j+ingredients[2].texture*k+ingredients[3].texture*l, 0)

					// Calculate the total score:
					score := capacity * durability * flavor * texture

					bestScore = max(bestScore, score)
				}
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   15,
		Part:  1,
		Value: bestScore,
	})
}

func partTwo(contents []string) {
	ingredients := parseInput(contents)

	bestScore := 0

	// Lots of for loops, but it's only 4 ingredients, so it's not too bad. 2 on example, 4 on input. Handle both:
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				for l := 0; l <= 100-i-j-k; l++ {
					// Check that the total is 100:
					if i+j+k+l != 100 {
						continue
					}

					// Calories:
					calories := ingredients[0].calories*i + ingredients[1].calories*j + ingredients[2].calories*k + ingredients[3].calories*l
					if calories != 500 {
						continue
					}

					// Calculate the score, 0 if negative:
					capacity := max(ingredients[0].capacity*i+ingredients[1].capacity*j+ingredients[2].capacity*k+ingredients[3].capacity*l, 0)
					durability := max(ingredients[0].durability*i+ingredients[1].durability*j+ingredients[2].durability*k+ingredients[3].durability*l, 0)
					flavor := max(ingredients[0].flavor*i+ingredients[1].flavor*j+ingredients[2].flavor*k+ingredients[3].flavor*l, 0)
					texture := max(ingredients[0].texture*i+ingredients[1].texture*j+ingredients[2].texture*k+ingredients[3].texture*l, 0)

					// Calculate the total score:
					score := capacity * durability * flavor * texture

					bestScore = max(bestScore, score)
				}
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   15,
		Part:  2,
		Value: bestScore,
	})
}

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func parseInput(contents []string) []ingredient {
	ingredients := make([]ingredient, 0)

	for _, line := range contents {
		var name string
		var capacity, durability, flavor, texture, calories int
		_, err := fmt.Sscanf(
			line,
			"%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&name,
			&capacity,
			&durability,
			&flavor,
			&texture,
			&calories,
		)
		if err != nil {
			panic(err)
		}
		name = name[:len(name)-1] // Remove the colon

		ingredients = append(ingredients, ingredient{name, capacity, durability, flavor, texture, calories})
	}

	return ingredients
}
