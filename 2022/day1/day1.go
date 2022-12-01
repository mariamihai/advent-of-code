package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Problem1 find the maximum calories carried by an elf
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

	var maxCalories = 0
	var currentCalories = 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if hasCurrentElfFinishedCalculation(line) {
			if currentElfCarriesMore(currentCalories, maxCalories) {
				maxCalories = currentCalories
			}

			currentCalories = 0
		} else {
			// Sum another line to current elf
			lineCalories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}

			currentCalories += lineCalories
		}
	}

	fmt.Printf("Part 1 - Maximum calories carried: %d\n", maxCalories)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Problem2() {
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
	var allCalories []int
	var currentCalories = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			allCalories = append(allCalories, currentCalories)
			currentCalories = 0
		} else {
			lineCalories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			currentCalories += lineCalories
		}
	}
	sort.Ints(allCalories)
	var total = 0
	for _, calorie := range allCalories[len(allCalories)-3:] {
		total += calorie
	}
	fmt.Printf("Part 2 - %d\n", total)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// hasCurrentElfFinishedCalculation current line is an empty one separating two elves notes
func hasCurrentElfFinishedCalculation(line string) bool {
	return line == ""
}

// currentElfCarriesMore current elf carries more calories than the current maximum holder
func currentElfCarriesMore(currentCalories, maxCalories int) bool {
	return currentCalories > maxCalories
}
