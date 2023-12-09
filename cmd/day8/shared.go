package day8

import (
	"bufio"
	"fmt"
	"strings"
)

type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

type NodeMap map[string]*Node

func (m NodeMap) GetOrCreate(name string) *Node {
	if existing, ok := m[name]; ok {
		return existing
	}

	node := &Node{Name: name}
	m[name] = node

	return node
}

func ParseDirections(scanner *bufio.Scanner) []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

func ParseNodes(scanner *bufio.Scanner) NodeMap {
	nodes := NodeMap{}

	for scanner.Scan() {
		name := scanner.Text()

		node := nodes.GetOrCreate(name)

		scanner.Scan() // =
		scanner.Scan()

		leftName := strings.Trim(scanner.Text(), "(,")
		node.Left = nodes.GetOrCreate(leftName)

		scanner.Scan()
		rightName := strings.TrimRight(scanner.Text(), ")")
		node.Right = nodes.GetOrCreate(rightName)
	}

	return nodes
}

func WalkPath(directions []rune, startingNode *Node, isEndNode func(node *Node) bool) (int, error) {
	currentNode := startingNode

	var steps int
	for !isEndNode(currentNode) {
		for _, direction := range directions {
			switch direction {
			case 'R':
				currentNode = currentNode.Right
			case 'L':
				currentNode = currentNode.Left
			default:
				return 0, fmt.Errorf("unknown direction: %c", direction)
			}
			steps++
		}
	}

	return steps, nil
}
