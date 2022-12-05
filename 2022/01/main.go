package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	log.Printf("%s", "Seasons Greetings!")
	log.Println("")
	first, second := values()
	log.Printf("First: %d Second: %d", first, second)
}

func values() (int, int) {
	input, _ := os.ReadFile("./data/source.txt")
	firstSplit := strings.Split(string(input), "\n\n")

	weights := make([]int, len(firstSplit))
	for i, first := range firstSplit {
		for _, second := range strings.Split(first, "\n") {
			c, _ := strconv.Atoi(strings.TrimSpace(second))
			weights[i] += c
		}
	}

	sort.Ints(weights)
	return weights[len(weights)-1], weights[len(weights)-1] + weights[len(weights)-2] + weights[len(weights)-3]
}
