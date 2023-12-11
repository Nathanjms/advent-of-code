package sharedstruct

import "fmt"

type Output struct {
	Day   int
	Part  int
	Value any
}

func PrintOutput(output Output) {
	fmt.Printf("% +v\n", output)
}
