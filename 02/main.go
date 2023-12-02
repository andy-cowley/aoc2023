package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Index  int
	Rounds []Round
}

func isRoundPossible(round Round) bool {
	var inTheBag Round
	inTheBag.Red = 12
	inTheBag.Green = 13
	inTheBag.Blue = 14

	switch {
	case round.Red > inTheBag.Red:
		return false
	case round.Green > inTheBag.Green:
		return false
	case round.Blue > inTheBag.Blue:
		return false
	}
	return true
}

func partA(gameData []Game) int {
	possibleGames := 0

	for _, game := range gameData {
		possible := true
		rounds := game.Rounds

		for _, round := range rounds {
			if !isRoundPossible(round) {
				possible = false
			}
		}
		if possible {
			possibleGames += game.Index
		}
	}
	return possibleGames
}

func partB(gameData []Game) int {
	sumOfPowers := 0
	for _, game := range gameData {
		var minimumSet Round
		for _, round := range game.Rounds {
			if round.Red > minimumSet.Red {
				minimumSet.Red = round.Red
			}
			if round.Green > minimumSet.Green {
				minimumSet.Green = round.Green
			}
			if round.Blue > minimumSet.Blue {
				minimumSet.Blue = round.Blue
			}
		}
		power := minimumSet.Red * minimumSet.Green * minimumSet.Blue
		sumOfPowers += power
	}
	return sumOfPowers
}

func main() {
	filePath := "data.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("failed to read file: %v", err)
	}

	// Close file at the end of the program
	defer file.Close()

	// Read file linewise
	scanner := bufio.NewScanner(file)

	var gameData []Game

	for scanner.Scan() {
		// do something with a line

		var game Game

		lineSlice := strings.Split(scanner.Text(), ":")
		gameIndex, _ := strconv.Atoi(strings.Split(lineSlice[0], " ")[1])
		game.Index = gameIndex
		rounds := strings.Split(lineSlice[1], ";")

		for _, round := range rounds {
			var roundCount Round

			ballCounts := strings.Split(round, ",")
			for _, c := range ballCounts {
				colourCount := strings.Split(c, " ")

				switch {
				case colourCount[2] == "red":
					count, _ := strconv.Atoi(colourCount[1])
					roundCount.Red = count

				case colourCount[2] == "green":
					count, _ := strconv.Atoi(colourCount[1])
					roundCount.Green = count

				case colourCount[2] == "blue":
					count, _ := strconv.Atoi(colourCount[1])
					roundCount.Blue = count
				}
			}

			game.Rounds = append(game.Rounds, roundCount)
		}
		gameData = append(gameData, game)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("failed to read file: %v", err)
	}

	scoreA := partA(gameData)
	fmt.Println("Score A:", scoreA)

	scoreB := partB(gameData)
	fmt.Println("Score B:", scoreB)

}
