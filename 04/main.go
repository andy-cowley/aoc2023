package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	Id       int
	Winning  []int
	Actual   []int
	Score    int
	Quantity int
}

func parseData(filePath string) ([]card, error) {
	var cardList []card
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Close file at the end of the program
	defer file.Close()

	// Read file linewise
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// do something with a line
		var newCard card

		// Split the data out
		cardMetaData := strings.Split(scanner.Text(), ":")[0]
		cardData := strings.Split(scanner.Text(), ":")[1]
		winningString := strings.Split(cardData, "|")[0]
		actualString := strings.Split(cardData, "|")[1]
		winningStringList := strings.Split(fixSpaces(winningString), " ")
		actualStringList := strings.Split(fixSpaces(actualString), " ")

		winning := make([]int, len(winningStringList))
		for i, s := range winningStringList {
			n, _ := strconv.Atoi(s)
			winning[i] = n
		}

		actual := make([]int, len(actualStringList))
		for i, s := range actualStringList {
			n, _ := strconv.Atoi(s)
			actual[i] = n
		}

		newCard.Id, _ = strconv.Atoi(strings.Split(fixSpaces(cardMetaData), " ")[1])
		newCard.Winning = winning
		newCard.Actual = actual

		cardList = append(cardList, newCard)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	return cardList, nil
}

func Power(n, m int) int {
	return int(math.Pow(float64(n), float64(m)))
}

func fixSpaces(s string) string {
	s = strings.TrimSpace(s)
	return strings.Join(strings.Fields(s), " ")
}

func cardScoreA(card card) int {
	var winningNumbers int
	for _, n := range card.Actual {
		winner := slices.Contains(card.Winning, n)
		if winner {
			winningNumbers++
		}
	}
	return 1 * Power(2, winningNumbers-1)
}

func cardScoreB(card card) card {
	var winningNumbers int
	for _, n := range card.Actual {
		winner := slices.Contains(card.Winning, n)
		if winner {
			winningNumbers++
		}
	}
	card.Score = winningNumbers
	return card
}

func partA(cardList []card) int {
	var score int

	for _, card := range cardList {
		score += cardScoreA(card)
	}

	return score
}

func listToMap(cardList []card) map[int]card {
	cardMap := make(map[int]card)
	for i := 0; i < len(cardList); i += 1 {
		cardList[i] = cardScoreB(cardList[i])
		cardList[i].Quantity = 1
		cardMap[cardList[i].Id] = cardList[i]
	}
	return cardMap
}

func cardLookUp(cardMap map[int]card, cardNumber int) card {
	card := cardMap[cardNumber]
	return card
}

func partB(cardMap map[int]card, cardList []card) (map[int]card, []card) {
	var newList []card
	for _, card := range cardList {
		for i := 1; i <= card.Score; i++ {
			newCard := cardLookUp(cardMap, card.Id+i)
			newCard.Quantity++
			cardMap[newCard.Id] = newCard
			newList = append(newList, newCard)
		}
	}

	if len(newList) > 0 {
		cardMap, newList = partB(cardMap, newList)
	}

	return cardMap, newList
}

func main() {
	cardList, err := parseData("data.txt")
	if err != nil {
		fmt.Printf("Error! %v", err)
		os.Exit(1)
	}

	cardMap := listToMap(cardList)

	partA := partA(cardList)
	newCardMap, _ := partB(cardMap, cardList)

	partB := 0
	for _, card := range newCardMap {
		partB += card.Quantity
	}

	fmt.Printf("Total Points: %v\n", partA)
	fmt.Printf("Total Card: %v", partB)
}
