package day6

import (
	"advent-of-code-2022/util"
	"bufio"
	"fmt"
	"github.com/samber/lo"
)

func Problem(nrOfCharacters int) {
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
