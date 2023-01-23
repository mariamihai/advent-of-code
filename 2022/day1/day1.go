package day1

import (
	"advent-of-code-2022/util"
	"bufio"
	"fmt"
	"sort"
)

// Problem1 find the maximum calories carried by an elf
func Problem1() func() {
	return func() {
		file := util.ReadFile("./day1/input2.txt")
		defer util.CloseFile()(file)

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
				// Sum another line to current elf (sum the line calories)
				currentCalories += util.StringToInt(line)
			}
		}

		fmt.Printf("Part 1 - Maximum calories carried: %d\n", maxCalories)

		err := scanner.Err()
		util.Boom(err)
	}
}

// Problem2 find the sum of calories carried by the three most carrying elves
func Problem2() func() {
	return func() {
		file := util.ReadFile("./day1/input2.txt")
		defer util.CloseFile()(file)

		var allCalories []int
		var currentCalories = 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			if hasCurrentElfFinishedCalculation(line) {
				allCalories = append(allCalories, currentCalories)
				currentCalories = 0
			} else {
				// Sum another line to current elf (sum the line calories)
				currentCalories += util.StringToInt(line)
			}
		}

		sort.Ints(allCalories)
		var total = 0
		for _, calorie := range allCalories[len(allCalories)-3:] {
			total += calorie
		}

		fmt.Printf("Part 2 - Sum of calories for first 3 elves: %d\n", total)

		err := scanner.Err()
		util.Boom(err)
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
