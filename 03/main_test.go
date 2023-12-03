package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	testResult1, _ := partA("test-data-site.txt")
	correctResult1 := 4361

	if testResult1 != correctResult1 {
		t.Errorf("Failed! got %v, want %v", testResult1, correctResult1)
	} else {
		t.Logf("Success!")
	}

	testResult2, _ := partA("test-data-case-1.txt")
	correctResult2 := 4419

	if testResult2 != correctResult2 {
		t.Errorf("Failed! got %v, want %v", testResult2, correctResult2)
	} else {
		t.Logf("Success!")
	}

	testResult3, _ := partA("test-data-case-2.txt")
	correctResult3 := 7006
	// Line1: 467
	// Line2: 0
	// Line3: 1285
	// Line4: 617
	// Line5: 1372
	// Line6: 58
	// Line7: 1190
	// Line8: 755
	// Line9: 0
	// Line10: 1262

	if testResult3 != correctResult3 {
		t.Errorf("Failed! got %v, want %v", testResult3, correctResult3)
	} else {
		t.Logf("Success!")
	}

	var initGear symbol
	Gears = append(Gears, initGear)
	testResult1B := partB(Gears)
	correctResult1B := 467835
	if testResult1B != correctResult1B {
		t.Errorf("Failed! got %v, want %v", testResult1B, correctResult1B)
	} else {
		t.Logf("Success!")
	}

}
