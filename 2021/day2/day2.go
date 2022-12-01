package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Problem1() {
	file, err := os.Open("./day2/input2.txt")
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

	var horizontalPosition = 0
	var verticalPosition = 0 // depth

	for scanner.Scan() {
		lineStr := scanner.Text()

		line := strings.Split(lineStr, " ")

		number := readLine(line[1])
		if line[0] == "forward" {
			horizontalPosition += number
		}
		if line[0] == "down" {
			verticalPosition += number
		}
		if line[0] == "up" {
			verticalPosition -= number
		}
	}

	fmt.Printf("Part 1 - horizontal x depth: %d\n", horizontalPosition*verticalPosition)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Problem2() {
	file, err := os.Open("./day2/input2.txt")
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

	var horizontalPosition = 0
	var verticalPosition = 0 // depth
	var aim = 0

	for scanner.Scan() {
		lineStr := scanner.Text()

		line := strings.Split(lineStr, " ")

		number := readLine(line[1])
		if line[0] == "forward" {
			horizontalPosition += number
			verticalPosition += aim * number
		}
		if line[0] == "down" {
			aim += number
		}
		if line[0] == "up" {
			aim -= number
		}
	}

	fmt.Printf("Part 2 - horizontal x depth: %d\n", horizontalPosition*verticalPosition)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readLine(str string) int {
	currentDepth, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Error during conversion")
	}

	return currentDepth
}
