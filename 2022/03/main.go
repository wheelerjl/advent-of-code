package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	log.Printf("%s", "Seasons Greetings!")
	log.Println("")
	first, second := scores()
	log.Printf("First: %d Second: %d", first, second)
}

func scores() (firstResult, secondResult int) {
	input, _ := os.ReadFile("./data/source.txt")
	rows := strings.Split(string(input), "\n")

	priorities := priorities()
	for i, row := range rows {
		left, right := row[:len(row)/2], row[len(row)/2:]
		for _, value := range left {
			if strings.ContainsRune(right, value) {
				firstResult += priorities[value]
				break
			}
		}

		if i%3 == 0 {
			for _, value := range row {
				if strings.ContainsRune(rows[i+1], value) && strings.ContainsRune(rows[i+2], value) {
					secondResult += priorities[value]
					break
				}
			}
		}
	}

	return
}

func priorities() map[rune]int {
	out := make(map[rune]int)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i, value := range alphabet {
		out[value] = i + 1
		out[unicode.ToUpper(value)] = i + 27
	}
	return out
}
