package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	instructions := lines[0]
	nodes := parseLines(lines)
	nodes = addressNodes(nodes)

	current := nodes[findNode(nodes, "AAA")]
	steps := 0

	for current.id != "ZZZ" {
		if instructions[steps%len(instructions)] == 'L' {
			current = nodes[current.leftIndex]
			steps++
		} else {
			current = nodes[current.rightIndex]
			steps++
		}
	}

	fmt.Println("Part 1:", steps)

	beginNodes := findBeginNodes(nodes)
	P := []int{}

	for _, node := range beginNodes {
		numOfStepsUntilZ := stepsUntilZ(node, nodes, instructions)
		P = append(P, numOfStepsUntilZ)
	}

	fmt.Println("Part 2:", LCM(P[0], P[1], P...))

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func stepsUntilZ(node Node, nodes []Node, instructions string) int {
	steps := 0

	for node.id[len(node.id)-1:] != "Z" {
		if instructions[steps%len(instructions)] == 'L' {
			node = nodes[node.leftIndex]
			steps++
		} else {
			node = nodes[node.rightIndex]
			steps++
		}
	}

	return steps
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func findBeginNodes(nodes []Node) []Node {
	beginNodes := []Node{}
	for _, node := range nodes {
		if node.id[len(node.id)-1:] == "A" {
			beginNodes = append(beginNodes, node)
		}
	}

	return beginNodes
}

func addressNodes(nodes []Node) []Node {
	newNodes := []Node{}

	for _, node := range nodes {
		node.leftIndex = findNode(nodes, node.left)
		node.rightIndex = findNode(nodes, node.right)
		newNodes = append(newNodes, node)
	}

	return newNodes
}

func findNode(nodes []Node, nodeId string) int {
	for idx, node := range nodes {
		if node.id == nodeId {
			return idx
		}
	}
	// should never happen
	fmt.Println("I should not run")
	return -1
}

type Node struct {
	left, right           string
	id                    string
	leftIndex, rightIndex int
}

func parseNodeFromLine(line string, lines []string) Node {
	s := strings.Split(line, "=")

	nodeId := strings.Trim(s[0], " ")

	s2 := strings.Split(strings.Trim(s[1], " "), ",")
	left, right := s2[0][1:], s2[1][1:len(s2[1])-1]

	return Node{left, right, nodeId, -1, -1}
}

func parseLines(lines []string) []Node {
	nodes := []Node{}

	for i := 2; i < len(lines); i++ {
		nodes = append(nodes, parseNodeFromLine(lines[i], lines))
	}

	return nodes
}
