package day6

import (
	"bufio"
	"fmt"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/samber/lo"
)

func Problem(nrOfCharacters int) func() {
	return func() {
		file := util.ReadFile("./day6/input.txt")
		defer util.CloseFile()(file)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			for i := nrOfCharacters; i < len(line); i++ {
				testingCharacters := line[i-nrOfCharacters : i]

				runes := []rune(testingCharacters)

				uniqueCharacters := lo.Uniq[rune](runes)

				if len(uniqueCharacters) == nrOfCharacters {
					fmt.Printf(" Found the solution - index %d\n", i)
					break
				}
			}
		}
	}
}
