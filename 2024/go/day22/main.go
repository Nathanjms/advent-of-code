package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
	"os"
	"path/filepath"
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

func partOne(contents []string) {
	buyerSecrets := parseInput(contents)

	// Cache the evolutions... could this get a bit too big?
	cachedEvolutions := make(map[int64]int64, 0)

	for n, secret := range buyerSecrets {
		newSecret := secret
		for i := 0; i < 2000; i++ {
			newSecret = evolve(newSecret, &cachedEvolutions)
		}
		buyerSecrets[n] = newSecret
	}

	sum := int64(0)
	for _, sec := range buyerSecrets {
		sum += sec
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   22,
		Part:  1,
		Value: sum,
	})
}

func partTwo(contents []string) {
	buyerSecrets := parseInput(contents)

	// Cache the evolutions... could this get a bit too big?
	cachedEvolutions := make(map[int64]int64, 0)
	sequencesToBananasPerBuyer := make([]map[[4]int64]int64, len(buyerSecrets))

	for n, secret := range buyerSecrets {
		sequencesToBananasPerBuyer[n] = make(map[[4]int64]int64, 0)
		newSecret := secret
		allSecrets := make([]int64, 2001)
		allSecrets[0] = secret
		for i := 1; i < 2001; i++ {
			newSecret = evolve(allSecrets[i-1], &cachedEvolutions)
			allSecrets[i] = newSecret
			if i > 3 {
				key := [4]int64{
					(allSecrets[i-3] % 10) - (allSecrets[i-4] % 10),
					(allSecrets[i-2] % 10) - (allSecrets[i-3] % 10),
					(allSecrets[i-1] % 10) - (allSecrets[i-2] % 10),
					(allSecrets[i] % 10) - (allSecrets[i-1] % 10),
				}
				if _, ok := sequencesToBananasPerBuyer[n][key]; !ok {
					// Already been set - we always take lowest so must skip now!
					sequencesToBananasPerBuyer[n][key] = (allSecrets[i] % 10) // Last digit is bananas
				}
			}
		}
		buyerSecrets[n] = newSecret
	}

	// Now go through each option and store the sum of bananas we get. THis will be prettyh verbose in Go...
	bananasPerKey := make(map[[4]int64]int64, 0)
	for _, bananasBySequence := range sequencesToBananasPerBuyer {
		for seq, bananaCount := range bananasBySequence {
			bananasPerKey[seq] += bananaCount
		}
	}

	currentBest := int64(0)
	for _, bananas := range bananasPerKey {
		if bananas > currentBest {
			currentBest = bananas
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   22,
		Part:  2,
		Value: currentBest,
	})
}

func evolve(secret int64, cachedEvolutions *map[int64]int64) int64 {
	if val, ok := (*cachedEvolutions)[secret]; ok {
		return val
	}

	originalSecret := secret

	// Otherwise, go through the steps:

	// Step 1:
	res := secret * 64
	secret = (res ^ secret) % 16777216

	// Step 2:
	res = int64(math.Floor(float64(secret) / float64(32)))
	secret = (res ^ secret) % 16777216

	// Step 4:
	res = secret * 2048
	secret = (res ^ secret) % 16777216

	(*cachedEvolutions)[originalSecret] = secret

	return secret
}

func parseInput(contents []string) []int64 {
	ints := make([]int64, 0)
	for _, line := range contents {
		val, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		ints = append(ints, val)
	}
	return ints
}
