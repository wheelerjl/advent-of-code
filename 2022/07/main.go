package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Node struct {
	Parent string
	Size   int
}

func main() {
	log.Printf("%s", "Seasons Greetings!")
	log.Println("")
	first, second := values()
	log.Printf("First: %d Second: %d", first, second)
}

func values() (firstResult, secondResult int) {
	input, _ := os.ReadFile("./data/source.txt")
	nodeMap := make(map[string]*Node)
	// Keep track of all nodes
	var currentNodeID string

	for _, row := range strings.Split(string(input), "\n") {
		if strings.HasPrefix(row, "dir ") || row == "$ ls" {
			continue
		} else if row == "$ cd .." {
			if currentNode, ok := nodeMap[currentNodeID]; ok {
				currentNodeID = currentNode.Parent
			}
		} else if strings.HasPrefix(row, "$ cd") {
			newNodeID := uuid.NewString()
			nodeMap[newNodeID] = &Node{
				Parent: currentNodeID,
			}
			currentNodeID = newNodeID
		} else {
			sizeData := strings.Split(row, " ")
			currentSize, err := strconv.Atoi(sizeData[0])
			if err == nil {
				if node, ok := nodeMap[currentNodeID]; ok {
					node.Size += currentSize
					currentParentID := node.Parent
					for currentParentID != "" {
						if parentNode, ok := nodeMap[currentParentID]; ok {
							parentNode.Size += currentSize
							currentParentID = parentNode.Parent
						}
					}
				}
			}
		}
	}

	var nodes []Node
	var usedSpace int
	for _, node := range nodeMap {
		if node.Parent == "" {
			usedSpace += node.Size
			continue
		}
		if node.Size <= 100000 {
			firstResult += node.Size
		}
		nodes = append(nodes, *node)
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Size < nodes[j].Size
	})

	for _, node := range nodes {
		if 70000000-usedSpace+node.Size > 30000000 {
			secondResult = node.Size
			break
		}
	}

	return
}
