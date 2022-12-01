package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Problem1() {
	file, err := os.Open("./day1/input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	var count = 0
	var previousValue = 10000

	for scanner.Scan() {
		currentDepth := readLine(scanner.Text())

		if hasDepthIncreased(currentDepth, previousValue) {
			count++
		}

		previousValue = currentDepth
	}

	fmt.Printf("Part 1 - Number of larger measurements: %d\n", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readLine(line string) int {
	currentDepth, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal("Error during conversion")
	}

	return currentDepth
}

func hasDepthIncreased(currentDepth, previousValue int) bool {
	return currentDepth > previousValue
}
