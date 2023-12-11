package day8

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/mariamihai/advent-of-code/util/formula"
	"strings"
)

func Problem1(filename string) int {
	isEndNodeCondition := func(currentNode Node) bool { return currentNode.value == "ZZZ" }

	instructions, network := getInput(filename)

	return calculationForValue(network, instructions, network["AAA"], isEndNodeCondition)
}

func Problem2(filename string) int {
	isStartNodeCondition := func(node Node) bool { return string(node.value[len(node.value)-1:]) == "A" }
	isEndNodeCondition := func(currentNode Node) bool { return string(currentNode.value[len(currentNode.value)-1:]) == "Z" }

	instructions, network := getInput(filename)
	pathsResults := []int{}

	for _, node := range network {
		if isStartNodeCondition(node) {
			pathsResults = append(pathsResults, calculationForValue(network, instructions, node, isEndNodeCondition))
		}
	}

	return formula.LCM(pathsResults...)
}

func calculationForValue(network map[string]Node, instructions string, node Node, endNodeCondition func(Node) bool) int {
	result := 0
	step := 0

	getNextNode := func(stepInstruction uint8) Node {
		if stepInstruction == 'L' {
			return network[node.left]

		} else {
			return network[node.right]
		}
	}

	for true {
		if endNodeCondition(node) {
			break
		}

		node = getNextNode(instructions[step%len(instructions)])

		result++
		step++
	}

	return result
}

type Node struct {
	value       string
	left, right string
}

func getInput(filename string) (string, map[string]Node) {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	network := make(map[string]Node)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan()
	scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)
		line = strings.Replace(line, "(", "", -1)
		line = strings.Replace(line, ")", "", -1)

		value := strings.Split(line, "=")[0]
		leftRight := strings.Split(line, "=")[1]
		leftNode := strings.Split(leftRight, ",")[0]
		rightNode := strings.Split(leftRight, ",")[1]

		network[value] = Node{
			value: value,
			left:  leftNode,
			right: rightNode,
		}
	}

	return instructions, network
}
