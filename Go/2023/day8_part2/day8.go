package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	label string
	left  string
	right string
}

// now we copy in the Node it is bad for memory performance but it wont break the code cause we do not alter the node itself in the code
func (n Node) getNextNodeKey(direction rune) (string, bool) {
	if direction == 'L' {
		return n.left, true
	}
	if direction == 'R' {
		return n.right, true
	}
	return "", false
}

type lisfOfNodes struct {
	instructions        []rune
	startingNodesLabels []string
	nodes               map[string]Node
}

func (l *lisfOfNodes) walkNodes() int {
	var stepsCounterArr []int
	fromToMap := make(map[string]string)
	fmt.Println(l.startingNodesLabels)

	for _, startNode := range l.startingNodesLabels {
		fromToMap[startNode] = startNode
	}
	for _, node := range l.startingNodesLabels {
		var stepsCounter int
		var nextNodeKey string
		for x := 0; x <= len(l.instructions)-1; x++ {
			// get all the next Nodes for the given instructions for every entry in the fromToMap
			currentInstruction := l.instructions[x]
			if stepsCounter == 0 {
				nextNode, err := l.nodes[node].getNextNodeKey(currentInstruction)
				if err == false {
					fmt.Println("Got no direction.")
				}
				nextNodeKey = nextNode
			} else {
				// here is still initial node
				nextNode, err := l.nodes[nextNodeKey].getNextNodeKey(currentInstruction)
				if err == false {
					fmt.Println("Got no direction.")
				}
				nextNodeKey = nextNode
			}

			stepsCounter++
			if nextNodeKey[2] == 'Z' {
				break
			}
			if x == len(l.instructions)-1 {
				// has to be minues one cause after the loop round finishes it will be incremented by 1 so if you want
				// it to be 0 at the start of the next iteration you gotta put it to -1
				x = -1
			}
			fmt.Println(stepsCounter)
		}
		stepsCounterArr = append(stepsCounterArr, stepsCounter)
	}
	fmt.Println(stepsCounterArr)
	leastCommonDenominator := LCM(stepsCounterArr[0], stepsCounterArr[1], stepsCounterArr[2:]...)
	return leastCommonDenominator
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

func allEndOnZ(nodeMap *map[string]string) bool {
	var all bool
	for _, nextNode := range *nodeMap {
		if nextNode[2] != 'Z' {
			fmt.Println(nextNode, nextNode[2], nextNode[2] != 'Z')
			all = false
			break
		}
		all = true
	}
	return all
}

func constructListOfNodesFromLines(lines []string) lisfOfNodes {
	var instructions []string
	var instructionRunes []rune
	var aEndingNodeKeys []string
	mapOfNodes := make(map[string]Node)
	for _, line := range lines {
		if !strings.Contains(line, "=") && len(line) != 0 {
			instructions = append(instructions, line)
			continue
		}
		if len(line) == 0 {
			continue
		}
		leftRight := strings.Split(line, "=")
		label := strings.ReplaceAll(leftRight[0], " ", "")
		directionsInBrackets := strings.ReplaceAll(leftRight[1], " ", "")
		directions := strings.ReplaceAll(strings.ReplaceAll(directionsInBrackets, ")", ""), "(", "")
		left := strings.Split(directions, ",")[0]
		right := strings.Split(directions, ",")[1]

		mapOfNodes[label] = Node{label: label, left: left, right: right}

		if label[2] == 'A' {
			aEndingNodeKeys = append(aEndingNodeKeys, label)
		}
	}

	for _, line := range instructions {
		for _, rune_ := range line {
			instructionRunes = append(instructionRunes, rune_)
		}
	}

	return lisfOfNodes{instructions: instructionRunes, startingNodesLabels: aEndingNodeKeys, nodes: mapOfNodes}
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	listOFNodes := constructListOfNodesFromLines(lines)
	fmt.Printf("%+v", listOFNodes)

	steps := listOFNodes.walkNodes()
	fmt.Println(steps)

}

// too low 14257
