package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Printf("%s", "Seasons Greetings!")
	log.Println("")
	first, second := scores()
	log.Printf("First: %d Second: %d", first, second)
}

func scores() (firstResult, secondResult int) {
	input, _ := os.ReadFile("./data/test.txt")

	for _, row := range strings.Split(string(input), "\n") {
		assignments := strings.Split(row, ",")
		firstLow, _ := strconv.Atoi(strings.Split(assignments[0], "-")[0])
		firstHigh, _ := strconv.Atoi(strings.Split(assignments[0], "-")[1])
		secondLow, _ := strconv.Atoi(strings.Split(assignments[1], "-")[0])
		secondHigh, _ := strconv.Atoi(strings.Split(assignments[1], "-")[1])

		if firstLow <= secondLow && firstHigh >= secondHigh || secondLow <= firstLow && secondHigh >= firstHigh {
			firstResult += 1
		}

		if secondLow <= firstHigh && secondHigh >= firstHigh || firstLow <= secondHigh && firstHigh >= secondHigh {
			secondResult += 1
		}
	}

	return
}
