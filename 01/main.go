package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func getLineSum(textLine string) (int, error) {
	var lineDigits []int
	for _, c := range textLine {
		i, err := strconv.Atoi(string(c))
		if err == nil {
			lineDigits = append(lineDigits, i)
		} else {
		}
	}
	first := 10 * lineDigits[0]
	last := lineDigits[len(lineDigits)-1]
	lineSum := first + last

	return lineSum, nil
}

type numberIndex struct {
	// Holds the first and last slice index of each value it finds
	First int
	Last  int
	Value int
}

// Implements a sort.Interface for []numberIndex based on First
type ByFirst []numberIndex

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i].First < a[j].First }

func getLineSumWithWords(textLine string) int {
	numberWords := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	var lineDigits []numberIndex

	fmt.Printf("Checking this line:\n%v\n", textLine)
	for key, value := range numberWords {
		pattern := regexp.MustCompile(key)
		loc := pattern.FindAllIndex([]byte(textLine), -1)

		if loc != nil {
			// fmt.Printf("Found: %v\n", pattern)
			for _, i := range loc {
				var numberIndex numberIndex
				numberIndex.First = i[0]
				numberIndex.Last = i[1]
				numberIndex.Value = value
				lineDigits = append(lineDigits, numberIndex)
			}
		}
	}
	// fmt.Printf("slice: %v\n", lineDigits)

	// We can now sort the slice by the First field
	sort.Sort(ByFirst(lineDigits))

	first := 10 * lineDigits[0].Value
	last := lineDigits[len(lineDigits)-1].Value
	lineSum := first + last

	fmt.Println(first, last)
	return lineSum
}

func getCalibrationValuePartA(filePath string) (int, error) {
	fileSum := 0
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %v", err)
	}

	// Close file at the end of the program
	defer file.Close()

	// Read file linewise
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// do something with a line
		lineSum, err := getLineSum(scanner.Text())

		if err != nil {
			return 0, fmt.Errorf("failed to sum line: %v", err)
		}
		fileSum += lineSum
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to read file: %v", err)
	}
	return fileSum, nil
}

func getCalibrationValuePartB(filePath string) int {
	fileSum := 0
	file, err := os.Open(filePath)
	if err != nil {
		return 0
	}

	// Close file at the end of the program
	defer file.Close()

	// Read file linewise
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// do something with a line
		lineSum := getLineSumWithWords(scanner.Text())

		if err != nil {
			return 0
		}
		fileSum += lineSum
	}

	if err := scanner.Err(); err != nil {
		return 0
	}
	return fileSum
}

func main() {
	fileSumPartA, err := getCalibrationValuePartA("data-1.txt")
	fileSumPartB := getCalibrationValuePartB("data-1.txt")

	if err != nil {
		fmt.Printf("failed to read file: %v", err)
	}

	fmt.Printf("Total for part A is: %v\n", fileSumPartA)
	fmt.Printf("Total for part B is: %v\n", fileSumPartB)
}
