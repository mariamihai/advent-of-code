package day3

import (
	"bufio"
	"fmt"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/samber/lo"
	"unicode"
)

// Problem1 sum priorities for misplaced item types based on value associated with letter
func Problem1() func() {
	return func() {
		file := util.ReadFile("./day3/input2.txt")
		defer util.CloseFile()(file)

		sum := 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			firstHalf := line[:len(line)/2]
			secondHalf := line[len(line)/2:]

			// Get the common character between the 2 halves
			// There is only one character common to both
			commonCharacters := lo.Uniq[rune](lo.Intersect[rune]([]rune(firstHalf), []rune(secondHalf)))

			sum += convertRuneToInt(commonCharacters[0])
		}

		fmt.Printf("Part 1 - the sum of the priorities of the item types that appear in both compartments of each rucksack: %d\n", sum)

		err := scanner.Err()
		util.Boom(err)
	}
}

// Problem2 sum priorities for the groups of three elves
func Problem2() func() {
	return func() {
		file := util.ReadFile("./day3/input2.txt")
		defer util.CloseFile()(file)

		sum := 0
		var elvesValues []string

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			elvesValues = append(elvesValues, line)

			if len(elvesValues) == 3 {
				// Find the common characters between first and second backpacks
				// and between the second and the third one
				firstTwoCommonCharactes := lo.Intersect[rune]([]rune(elvesValues[0]), []rune(elvesValues[1]))
				lastTwoCommonCharactes := lo.Intersect[rune]([]rune(elvesValues[1]), []rune(elvesValues[2]))

				// Get the common character against all three backpacks
				commonAgainstAllThree := lo.Uniq[rune](lo.Intersect[rune](firstTwoCommonCharactes, lastTwoCommonCharactes))

				sum += convertRuneToInt(commonAgainstAllThree[0])

				// Keep the underlying array
				elvesValues = elvesValues[:0]
			}
		}

		fmt.Printf("Part 2 - the sum of the priorities for 3 elves: %d\n", sum)

		err := scanner.Err()
		util.Boom(err)
	}
}

func convertRuneToInt(ch rune) int {
	if unicode.IsLower(ch) {
		// lower case between 1 and 26
		return int(ch) - 96
	}

	// uppercase between 27 and 52
	return int(ch) - 38
}
