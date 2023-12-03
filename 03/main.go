package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type partNumber struct {
	Line   int
	Start  int
	End    int
	Value  int
	IsPart bool
}

type symbol struct {
	Line   int
	Index  int
	Value  string
	IsGear bool
	GearA  int
	GearB  int
}

func findNumbers(text string, lineNumber int) ([]partNumber, error) {
	var partNumbers []partNumber
	numberRegex := regexp.MustCompile(`\b\d+\b`)

	foundNumbers := numberRegex.FindAllStringIndex(text, -1)
	for _, foundNumber := range foundNumbers {
		var partNumber partNumber
		partNumber.Line = lineNumber
		partNumber.Start = foundNumber[0]
		partNumber.End = foundNumber[1]
		match := text[partNumber.Start:partNumber.End]
		num, err := strconv.Atoi(match)
		if err == nil {
			partNumber.Value = num
		} else {
			return nil, fmt.Errorf("Error casting string to int: %v", err)
		}
		partNumbers = append(partNumbers, partNumber)
	}

	return partNumbers, nil
}

func findSymbols(text string, lineNumber int) ([]symbol, error) {
	var symbols []symbol
	symbolRegex := regexp.MustCompile(`[^\w\d.\n\r]`)

	foundSymbols := symbolRegex.FindAllStringIndex(text, -1)
	for _, foundSymbol := range foundSymbols {
		var symbol symbol
		symbol.Line = lineNumber
		symbol.Index = foundSymbol[0]
		match := text[symbol.Index : symbol.Index+1]
		symbol.Value = match
		symbols = append(symbols, symbol)
	}
	return symbols, nil
}

func updateGearList(gear symbol, partNumber int, gearList []symbol) []symbol {
	var newGearList []symbol
	if len(gearList) == 0 {
		gear.GearA = partNumber
		newGearList = append(newGearList, gear)
	} else {
		for i, existingGear := range gearList {
			if existingGear.Line == gear.Line && existingGear.Index == gear.Index {
				existingGear.GearB = partNumber
				existingGear.IsGear = true
				// newGearList = append(gearList, existingGear)
				newGearList = slices.Replace(gearList, i, i+1, existingGear)
			} else {
				gear.GearA = partNumber
				newGearList = append(gearList, gear)
			}
		}
	}
	return newGearList
}

func testThisPart(
	partNumber partNumber,
	symbols []symbol,
	gearList []symbol,
) (partNumber, []symbol) {
	for _, symbol := range symbols {

		// Test case: return true if:
		if symbol.Index >= partNumber.Start-1 &&
			symbol.Index <= partNumber.End && //The end index is one mor then the character -- it's EXCLUSIVE
			symbol.Line >= partNumber.Line-1 &&
			symbol.Line <= partNumber.Line+1 {
			partNumber.IsPart = true

			if symbol.Value == "*" {
				gearList = updateGearList(symbol, partNumber.Value, gearList)
			}
		}
	}
	return partNumber, gearList
}

func partA(filePath string) (int, []symbol, error) {
	var partNumbersTest []partNumber
	var partNumbersTested []partNumber
	var symbols []symbol
	var gearList []symbol

	partSum := 0
	line := 1
	file, err := os.Open(filePath)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Close file at the end of the program
	defer file.Close()

	// Read file linewise
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return 0, nil, fmt.Errorf("failed to read file: %v", err)
	}

	for scanner.Scan() {
		// do something with a line
		foundNumbers, err := findNumbers(scanner.Text(), line)
		foundSymbols, err := findSymbols(scanner.Text(), line)
		if err != nil {
			return 0, nil, fmt.Errorf("Error: %v", err)
		}

		partNumbersTest = append(partNumbersTest, foundNumbers...) // Three dots means merge them
		symbols = append(symbols, foundSymbols...)
		line++
	}

	// Check each number to see if it is a part
	for _, testPart := range partNumbersTest {
		var testedPart partNumber
		testedPart, gearList = testThisPart(testPart, symbols, gearList)
		partNumbersTested = append(partNumbersTested, testedPart)
	}

	for _, part := range partNumbersTested {
		if part.IsPart {
			partSum += part.Value
		}
	}

	return partSum, gearList, nil
}

func partB(gears []symbol) int {
	var gearSum int

	for _, gear := range gears {
		if gear.IsGear {
			gearSum += gear.GearA * gear.GearB
		}
	}

	return gearSum
}

func main() {
	partA, gearList, _ := partA("data.txt")
	partB := partB(gearList)

	fmt.Printf("Part A: %v\n", partA)
	fmt.Printf("Part B: %v\n", partB)
}
