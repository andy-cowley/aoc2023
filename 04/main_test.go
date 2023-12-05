package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(t *testing.T) {

	var testCard card
	testCard.Id = 1
	testCard.Winning = []int{41, 48, 83, 86, 17}
	testCard.Actual = []int{83, 86, 6, 31, 17, 9, 48, 53}

	testCardScore := cardScoreA(testCard)
	testCardWant := 8

	if testCardScore != testCardWant {
		t.Errorf("Failed! Got %v, wanted %v", testCardScore, testCardWant)
	} else {
		t.Logf("Success!")
	}

	cardList, err := parseData("test-data-site.txt")
	if err != nil {
		fmt.Printf("Error! %v", err)
		os.Exit(1)
	}

	testA1 := partA(cardList)
	wantA1 := 13

	if testA1 != wantA1 {
		t.Errorf("Failed! Got %v, wanted %v", testA1, wantA1)
	} else {
		t.Logf("Success!")
	}

	cardMap := listToMap(cardList)
	newCardMap, _ := partB(cardMap, cardList)
	testB1 := 0
	for _, card := range newCardMap {
		testB1 += card.Quantity
	}

	wantB1 := 30

	if testB1 != wantB1 {
		t.Errorf("Failed! Got %v, wanted %v", testB1, wantB1)
	} else {
		t.Logf("Success!")
	}
}
