package main

import (
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
	first, second := values()
	log.Println()
	log.Printf("First: %d Second: %d", first, second)
}

func values() (firstResult, secondResult int) {
	input, _ := os.ReadFile("./data/source.txt")
	splitForest := strings.Split(string(input), "\n")
	eastBorder := len(splitForest[0]) - 1
	southBorder := len(splitForest) - 1

	// Initialize the forest
	forest := make([][]int, eastBorder+1)
	for i := 0; i <= eastBorder; i++ {
		forest[i] = make([]int, southBorder+1)
	}
	treeSeen := make([][]bool, eastBorder+1)
	for i := 0; i <= eastBorder; i++ {
		treeSeen[i] = make([]bool, southBorder+1)
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

	// Populate visible trees
	for i := 0; i <= eastBorder; i++ {
		for j := 0; j <= southBorder; j++ {
			if i == 0 || j == 0 || i == eastBorder || j == southBorder {
				treeSeen[i][j] = true
				scenicScore[i][j] = 0
				continue
			}
			currentTree := forest[i][j]
			biggestUp := true
			biggestLeft := true
			biggestDown := true
			biggestRight := true
			viewedUp := 1
			viewedLeft := 1
			viewedDown := 1
			viewedRight := 1

			// Am I bigger then everyone above me
			for k := i - 1; k >= 0; k-- {
				checkingTree := forest[k][j]
				if currentTree <= checkingTree {
					biggestUp = false
				}

				if k == i-1 && currentTree <= checkingTree {
					break
				} else if k == i-1 {
					continue
				}

				if currentTree <= checkingTree {
					viewedUp++
					break
				} else {
					viewedUp++
				}
			}
			if biggestUp {
				treeSeen[i][j] = true
			}

			// Am I bigger then overyone left of me
			for k := j - 1; k >= 0; k-- {
				checkingTree := forest[i][k]
				if currentTree <= checkingTree {
					biggestLeft = false
				}

				if k == j-1 && currentTree <= checkingTree {
					break
				} else if k == j-1 {
					continue
				}

				if currentTree <= checkingTree {
					viewedLeft++
					break
				} else {
					viewedLeft++
				}
			}
			if biggestLeft {
				treeSeen[i][j] = true
			}

			// Am I bigger then everyone below me
			for k := i + 1; k <= southBorder; k++ {
				checkingTree := forest[k][j]
				if currentTree <= checkingTree {
					biggestDown = false
				}

				if k == i+1 && currentTree <= checkingTree {
					break
				} else if k == i+1 {
					continue
				}

				if currentTree <= checkingTree {
					viewedDown++
					break
				} else {
					viewedDown++
				}
			}
			if biggestDown {
				treeSeen[i][j] = true
			}

			// Am I bigger then everyone right of me
			for k := j + 1; k <= eastBorder; k++ {
				checkingTree := forest[i][k]
				if currentTree <= checkingTree {
					biggestRight = false
				}

				if k == j+1 && currentTree <= checkingTree {
					break
				} else if k == j+1 {
					continue
				}

				if currentTree <= checkingTree {
					viewedRight++
					break
				} else {
					viewedRight++
				}
			}
			if biggestRight {
				treeSeen[i][j] = true
			}

			score := viewedUp * viewedLeft * viewedDown * viewedRight
			scenicScore[i][j] = score
		}
	}

	// Scenic Score
	for i := 0; i <= eastBorder; i++ {
		for j := 0; j <= southBorder; j++ {
			if scenicScore[i][j] > secondResult {
				secondResult = scenicScore[i][j]
			}
			if treeSeen[i][j] {
				firstResult += 1
			}
		}
	}

	return
}
