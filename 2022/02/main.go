package main

import (
	"log"
	"os"
	"strings"
)

const (
	Lose     = 1
	Rock     = 1
	Draw     = 2
	Paper    = 2
	Win      = 3
	Scissors = 3
)

func main() {
	log.Printf("%s", "Seasons Greetings!")
	log.Println("")
	first, second := scores()
	log.Printf("First: %d Second: %d", first, second)
}

func scores() (first, second int) {
	input, _ := os.ReadFile("./data/source.txt")
	rows := strings.Split(string(input), "\n")
	for _, row := range rows {
		myThrow, theirThrow := throws(row)
		first += firstScores(myThrow, theirThrow)
		second += secondScores(theirThrow, myThrow)
	}

	return
}

func throws(input string) (mine, theirs int) {
	throws := strings.Split(input, " ")
	// Theirs
	switch throws[0] {
	case "A":
		theirs = Rock
	case "B":
		theirs = Paper
	case "C":
		theirs = Scissors
	}

	// Mine
	switch throws[1] {
	case "X":
		mine = Rock
	case "Y":
		mine = Paper
	case "Z":
		mine = Scissors
	}
	return
}

func firstScores(myThrow, theirThrow int) (score int) {
	if myThrow == theirThrow {
		score = myThrow + 3
	} else if (myThrow == Rock) && (theirThrow == Scissors) ||
		(myThrow == Scissors) && (theirThrow == Paper) ||
		(myThrow == Paper) && (theirThrow == Rock) {
		score = myThrow + 6
	} else {
		score = myThrow + 0
	}

	return
}

func secondScores(theirThrow, neededResult int) (score int) {
	if neededResult == Draw {
		score = theirThrow + 3
	} else if neededResult == Lose {
		if theirThrow == Rock {
			score = Scissors
		} else if theirThrow == Paper {
			score = Rock
		} else {
			score = Paper
		}
	} else {
		if theirThrow == Rock {
			score = 6 + Paper
		} else if theirThrow == Paper {
			score = 6 + Scissors
		} else {
			score = 6 + Rock
		}
	}

	return
}
