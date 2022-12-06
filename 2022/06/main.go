package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("%s", "Seasons Greetings!")
	log.Println("")
	first, second := values()
	log.Printf("First: %d Second: %d", first, second)
}

func values() (firstResult, secondResult int) {
	input, _ := os.ReadFile("./data/source.txt")

	foundFirst := false
	foundSecond := false
	for i := range string(input) {
		if foundFirst && foundSecond {
			break
		}

		if !foundFirst {
			charMap := make(map[string]bool)
			for _, value := range []int{0, 1, 2, 3} {
				charMap[string(input[i+value])] = true
			}

			if len(charMap) == 4 {
				firstResult = i + 4
				foundFirst = true
			}
		}

		if !foundSecond {
			charMap := make(map[string]bool)
			for _, value := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13} {
				charMap[string(input[i+value])] = true
			}

			if len(charMap) == 14 {
				secondResult = i + 14
				foundSecond = true
			}
		}
	}

	return
}
