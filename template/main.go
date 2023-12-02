package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Close file at the end of the program
	defer file.Close()

	// Read file linewise
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	return nil
}

func main() {
}
