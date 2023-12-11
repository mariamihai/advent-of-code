package day1

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
)

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var count = 0
	var previousValue = 10000

	hasDepthIncreased := func(currentDepth, previousValue int) bool { return currentDepth > previousValue }

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentDepth := util.StringToInt(scanner.Text())

		if hasDepthIncreased(currentDepth, previousValue) {
			count++
		}

		previousValue = currentDepth
	}

	util.Boom(scanner.Err())

	return count
}
