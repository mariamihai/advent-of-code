package day9

import (
	"bufio"
	"github.com/samber/lo"
	"strings"

	"github.com/mariamihai/advent-of-code/util"
)

func Problem1(filename string) int {
	var result int

	lines := getInput(filename)

	for _, line := range lines {
		//result += resultForLinePart1(line)
		result += resultForLine(line, func(line []int) int { return line[len(line)-1] }, func(a, b int) int { return a + b })
	}

	return result
}

func Problem2(filename string) int {
	var result int

	lines := getInput(filename)

	for _, line := range lines {
		// Can reverse the line and calculate the last value instead
		//result += resultForLinePart1(lo.Reverse(line))

		//result += resultForLinePart2(line)

		result += resultForLine(line, func(line []int) int { return line[0] }, func(a, b int) int { return a - b })
	}

	return result
}

func getInput(filename string) [][]int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	lines := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := util.StringSliceToIntSlice(strings.Fields(scanner.Text()))
		lines = append(lines, line)
	}

	return lines
}

func areAllValuesTheSame(line []int) bool {
	return len(lo.Uniq(line)) == 1
}

// resultForLine generic implementation for both problems
func resultForLine(line []int, getCurrentValue func(line []int) int, result func(a, b int) int) int {
	if areAllValuesTheSame(line) {
		return line[0]
	}

	currentValue := getCurrentValue(line)
	for i := 0; i < len(line)-1; i++ {
		line[i] = line[i+1] - line[i]
	}

	return result(currentValue, resultForLine(line[:len(line)-1], getCurrentValue, result))
}

// resultForLinePart1 first implementation for problem 1
func resultForLinePart1(line []int) int {
	if areAllValuesTheSame(line) {
		return line[0]
	}

	currentValue := line[len(line)-1]
	for i := 0; i < len(line)-1; i++ {
		line[i] = line[i+1] - line[i]
	}

	return currentValue + resultForLinePart1(line[:len(line)-1])
}

// resultForLinePart2 first implementation for problem 2
func resultForLinePart2(line []int) int {
	if areAllValuesTheSame(line) {
		return line[0]
	}

	currentValue := line[0]
	for i := 0; i < len(line)-1; i++ {
		line[i] = line[i+1] - line[i]
	}

	return currentValue - resultForLinePart2(line[:len(line)-1])
}
