package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {
	log.Printf("%s", "Seasons Greetings!")
	log.Println("")
	first, second := values()
	log.Printf("First: %d Second: %d", first, second)
}

type Node struct {
	Parent string
	Size   int
}

func values() (firstResult, secondResult int) {
	input, _ := os.ReadFile("./data/source.txt")
	nodeMap := make(map[string]Node)
	// Keep track of all nodes
	var currentNodeID string

	for _, row := range strings.Split(string(input), "\n") {
		if strings.HasPrefix(row, "dir ") || row == "$ ls" {
			continue
		} else if row == "$ cd .." {
			if currentNode, ok := nodeMap[currentNodeID]; ok {
				currentNodeID = currentNode.Parent
			}
			continue
		} else if strings.HasPrefix(row, "$ cd") {
			newNodeID := uuid.NewString()
			nodeMap[newNodeID] = Node{
				Parent: currentNodeID,
			}
			currentNodeID = newNodeID

			continue
		}

		sizeData := strings.Split(row, " ")
		currentSize, err := strconv.Atoi(sizeData[0])
		if err == nil {
			if node, ok := nodeMap[currentNodeID]; ok {
				node.Size += currentSize
				nodeMap[currentNodeID] = node
				currentParentID := node.Parent
				for currentParentID != "" {
					if parentNode, ok := nodeMap[currentParentID]; ok {
						parentNode.Size += currentSize
						nodeMap[currentParentID] = parentNode
						currentParentID = parentNode.Parent
					}
				}
			}
		}
	}

	var nodes []Node
	var usedSpace int
	for _, node := range nodeMap {
		if node.Size <= 100000 {
			firstResult += node.Size
		}
		nodes = append(nodes, node)
		if node.Parent == "" {
			usedSpace += node.Size
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Size < nodes[j].Size
	})

	for i, node := range nodes {
		log.Printf("Index: %d Size: %d", i, node.Size)
		space := 70000000 - usedSpace
		if space+node.Size <= 30000000 {
			continue
		} else {
			secondResult = nodes[i].Size
			break
		}
	}

	return
}
