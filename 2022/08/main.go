package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent string
	Size   int
}

func main() {
	log.Printf("%s", "Seasons Greetings!")
	log.Println("")
	first, second := values()
	log.Println("")
	log.Printf("First: %d Second: %d", first, second)
}

func values() (firstResult, secondResult int) {
	input, _ := os.ReadFile("./data/test.txt")
	splitForest := strings.Split(string(input), "\n")
	eastBorder := len(splitForest[0]) - 1
	southBorder := len(splitForest) - 1

	// Initialize the forest
	forest := make([][]int, eastBorder+1)
	for i := 0; i <= eastBorder; i++ {
		forest[i] = make([]int, southBorder+1)
	}
	scenicScore := make([][]int, eastBorder+1)
	for i := 0; i <= eastBorder; i++ {
		scenicScore[i] = make([]int, southBorder+1)
	}

	// Populate the forest
	for i := 0; i <= eastBorder; i++ {
		for j := 0; j <= southBorder; j++ {
			treeHeight, _ := strconv.Atoi(string(splitForest[j][i]))
			forest[j][i] = treeHeight
		}
	}

	// Populate the forest
	for i := 0; i <= eastBorder; i++ {
		row := ""
		for j := 0; j <= southBorder; j++ {
			row += fmt.Sprintf("%d", forest[i][j])
		}
		log.Print(row)
	}

	// Populate visible trees
	for i := 0; i <= eastBorder; i++ {
		for j := 0; j <= southBorder; j++ {
			currentTree := forest[i][j]

			// Am I bigger then everyone above me
			countedSeen := false
			biggestUp := true
			for k := i - 1; k >= 0; k-- {
				checkingTree := forest[k][j]
				if currentTree <= checkingTree {
					biggestUp = false
					break
				}
			}
			if biggestUp && !countedSeen {
				countedSeen = true
				firstResult += 1
			}

			// Am I bigger then everyone below me
			biggestDown := true
			for k := i + 1; k <= southBorder; k++ {
				checkingTree := forest[k][j]
				if currentTree <= checkingTree {
					biggestDown = false
					break
				}
			}
			if biggestDown && !countedSeen {
				countedSeen = true
				firstResult += 1
			}

			// Am I bigger then overyone left of me
			biggestLeft := true
			for k := j - 1; k >= 0; k-- {
				checkingTree := forest[i][k]
				if currentTree <= checkingTree {
					biggestLeft = false
					break
				}
			}
			if biggestLeft && !countedSeen {
				countedSeen = true
				firstResult += 1
			}

			// Am I bigger then everyone right of me
			biggestRight := true
			for k := j + 1; k <= eastBorder; k++ {
				checkingTree := forest[i][k]
				if currentTree <= checkingTree {
					biggestRight = false
					break
				}
			}
			if biggestRight && !countedSeen {
				countedSeen = true
				firstResult += 1
			}
		}
	}

	return
}
