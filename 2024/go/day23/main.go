package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
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

	_, contents := sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	nodeMaps := buildMap(contents)
	var solutions [][]string

	// Go through each and look for 2 levels deep, the same one again/ Recursive solution or BFS probably would work
	for targetString := range nodeMaps {
		// We need to go 2 levels deep and see if we get the original string again. If we do, that one is a match.
		solutions = append(solutions, hasTargetAfterDepth(2, targetString, targetString, []string{}, &nodeMaps)...)
	}

	uniqueSolutionsWithT := make(map[string]bool)

	for i := 0; i < len(solutions); i++ {
		hasTString := false

		for _, key := range solutions[i] {
			if key[0] == 't' {
				hasTString = true
				break
			}
		}
		if hasTString {
			sort.Slice(solutions[i], func(n, m int) bool {
				return solutions[i][n] > solutions[i][m]
			})

			uniqueSolutionsWithT[strings.Join(solutions[i], "-")] = true
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   23,
		Part:  1,
		Value: len(uniqueSolutionsWithT),
	})
}

func partTwo(contents []string) {
	nodeMaps := buildMap(contents)

	searchesByNode := make(map[string][]string, 0)
	for node := range nodeMaps {
		searchesByNode[node] = search(node, []string{node}, &nodeMaps)
	}

	// fmt.Println(searchesByNode)

	// Make them unique by changing each to a map... kinda wish we were using a language with sets now :c
	mapsByNode := make(map[string]map[string]bool, 0)
	for node, values := range searchesByNode {
		for _, val := range values {
			if _, ok := mapsByNode[node]; !ok {
				mapsByNode[node] = make(map[string]bool, 0)
			}
			mapsByNode[node][val] = true
		}
	}

	// Now we un-map them and then sort... probably could do this better but it hopefully will work...
	uniqueSearchesByNode := make(map[string][]string, 0)
	for node, nodeMap := range mapsByNode {
		uniqueSearchesByNode[node] = make([]string, 0)
		for newNode := range nodeMap {
			uniqueSearchesByNode[node] = append(uniqueSearchesByNode[node], newNode)
		}
	}

	// Pick out the longest, any will do (I think...)
	longestIndex := ""
	currentLongestLength := math.MinInt
	for i, v := range uniqueSearchesByNode {
		if len(v) > currentLongestLength {
			longestIndex = i
			currentLongestLength = len(v)
		}
	}

	sort.Slice(uniqueSearchesByNode[longestIndex], func(i, j int) bool {
		return uniqueSearchesByNode[longestIndex][i] < uniqueSearchesByNode[longestIndex][j]
	})

	var builder strings.Builder

	for i, node := range uniqueSearchesByNode[longestIndex] {
		builder.WriteString(node)
		if i < len(uniqueSearchesByNode[longestIndex])-1 {
			builder.WriteString(",")
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   23,
		Part:  2,
		Value: builder.String(),
	})
}

func search(node string, existingConnections []string, nodeMap *map[string][]string) []string {
	// For each connection off the node...
	for _, neighbour := range (*nodeMap)[node] {
		// Skip if neighbour is already in the existing conns:
		isInExisting := false
		for _, conn := range existingConnections {
			if conn == neighbour {
				isInExisting = true
				break
			}
		}

		if isInExisting {
			continue
		}

		//  Neighbour must be connected to everything in the set:
		connectedToEverything := true
		for _, queryNode := range existingConnections {
			// For each of these, it must be connected to the new node
			if !nodeIsAttached(queryNode, neighbour, nodeMap) {
				connectedToEverything = false
				break
			}
		}

		if !connectedToEverything {
			continue
		}

		// Add this node and then trial all it's neighbours
		existingConnections = append(existingConnections, neighbour)

		existingConnections = append(existingConnections, search(neighbour, existingConnections, nodeMap)...)
	}

	return existingConnections
}

func nodeIsAttached(baseNode string, searchNode string, nodeMap *map[string][]string) bool {
	for _, neighbour := range (*nodeMap)[baseNode] {
		if neighbour == searchNode {
			return true
		}
	}

	return false
}

func hasTargetAfterDepth(depth int, currentString string, targetString string, previousStrings []string, nodeMap *map[string][]string) [][]string {
	targets := make([][]string, 0)

	previousStrings = append(previousStrings, currentString)

	if depth == 0 {
		// We need the target string to be inside the map for the current string
		for _, possibleTargetString := range (*nodeMap)[currentString] {
			if possibleTargetString == targetString {
				targets = append(targets, previousStrings)
				return targets
			}
		}
		return targets // not a valid path
	}

	// Otherwise, call this again and drop the depth by 1
	for _, nextString := range (*nodeMap)[currentString] {
		targets = append(targets, hasTargetAfterDepth(depth-1, nextString, targetString, previousStrings, nodeMap)...)
	}

	return targets
}

// func expectOutputAfterDepth(depthRemaining int, currentString string, targetString string, previousStrings []string, nodeMap *map[string][]string) (bool, []string) {
// 	if depthRemaining == 0 {
// 		// We need the target string to be inside the map for the current string
// 		for _, possibleTargetString := range (*nodeMap)[currentString] {
// 			if possibleTargetString == targetString {
// 				return true, previousStrings
// 			}
// 		}
// 		return false, []string{}
// 	}

// 	// Otherwise, call this again and drop the depth by 1
// 	for _, nextString := range (*nodeMap)[currentString] {

// 	}

// 	return false, []string{}
// }

func buildMap(contents []string) map[string][]string {
	nodeMap := make(map[string][]string, 0)
	for _, line := range contents {
		stringsTemp := strings.Split(line, "-")
		if _, ok := nodeMap[stringsTemp[0]]; !ok {
			nodeMap[stringsTemp[0]] = make([]string, 0)
		}
		nodeMap[stringsTemp[0]] = append(nodeMap[stringsTemp[0]], stringsTemp[1])
		if _, ok := nodeMap[stringsTemp[1]]; !ok {
			nodeMap[stringsTemp[1]] = make([]string, 0)
		}
		nodeMap[stringsTemp[1]] = append(nodeMap[stringsTemp[1]], stringsTemp[0])
	}

	return nodeMap
}
