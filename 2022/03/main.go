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
	x := make(map[rune]int)
	runes := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for i, value := range runes {
		x[value] = i + 1
		x[unicode.ToUpper(value)] = i + 27
	}
	return x
}
