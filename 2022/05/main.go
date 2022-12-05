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
	first, second := values()
	log.Printf("First: %s Second: %s", first, second)
}

func values() (firstResult, secondResult string) {
	input, _ := os.ReadFile("./data/source.txt")

	indexMap := stackIndexes()

	stackOne := make(map[int]string)
	stackTwo := make(map[int]string)
	// Build the stacks
	for _, row := range strings.Split(string(input), "\n") {
		for i, char := range row {
			if char == '[' {
				stackOne[indexMap[i+1]] = stackOne[indexMap[i+1]] + string(row[i+1])
				stackTwo[indexMap[i+1]] = stackTwo[indexMap[i+1]] + string(row[i+1])
			}
		}

		if strings.HasPrefix(row, "move") {
			step := strings.Split(row, " ")
			count, _ := strconv.Atoi(step[1])
			from, _ := strconv.Atoi(step[3])
			to, _ := strconv.Atoi(step[5])

			stackTwo[to] = stackTwo[from][0:count] + stackTwo[to]
			stackTwo[from] = stackTwo[from][count:]

			for count > 0 {
				stackOne[to] = stackOne[from][0:1] + stackOne[to]
				stackOne[from] = stackOne[from][1:]
				count--
			}
		} else {
			log.Println(row)
		}
	}

	for _, value := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if len(stackOne[value]) > 0 {
			firstResult += string(stackOne[value][0])
		}

		if len(stackTwo[value]) > 0 {
			secondResult += string(stackTwo[value][0])
		}
	}

	return
}

func stackIndexes() map[int]int {
	out := make(map[int]int)
	for _, value := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if value == 1 {
			out[1] = 1
		} else {
			out[4*(value-1)+1] = value
		}
	}
	return out
}
