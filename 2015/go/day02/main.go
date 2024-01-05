package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	executableDir := filepath.Dir(executablePath)
	var inputPath = filepath.Join(executableDir, "example-input1")
	fmt.Println("Input Path:", inputPath)
	// Use inputPath as needed
}
