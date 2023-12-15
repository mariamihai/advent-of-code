package day15

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"strings"
)

func Problem1(filename string) int {
	input := readInput(filename)

	var result int
	for i := 0; i < len(input); i++ {
		result += customHash(input[i])
	}

	return result
}

func readInput(filename string) []string {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func customHash(input string) int {
	var currentValue int
	for j := 0; j < len(input); j++ {
		currentValue = (currentValue + int(input[j])) * 17 % 256
	}
	return currentValue
}
