package day12

import (
	"bufio"
	"fmt"
	"github.com/mariamihai/advent-of-code/util"
	"strings"
)

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		input := line[0]
		arrangements := util.StringSliceToIntSlice(strings.Split(line[1], ","))

		result += evaluate(input, arrangements)
	}

	return result
}

// Problem2 - not resolved
func Problem2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input, arrangements := unfoldRecords(scanner.Text())

		result += evaluate(input, arrangements)
	}

	return result
}

func unfoldRecords(line string) (string, []int) {
	input := strings.Fields(line)
	unfoldedInput := fmt.Sprintf("%s?%s?%s?%s?%s", input[0], input[0], input[0], input[0], input[0])
	initialArrangements := util.StringSliceToIntSlice(strings.Split(input[1], ","))

	arrangements := append(initialArrangements, initialArrangements...)
	arrangements = append(arrangements, initialArrangements...)
	arrangements = append(arrangements, initialArrangements...)
	arrangements = append(arrangements, initialArrangements...)

	return unfoldedInput, arrangements
}

func evaluate(input string, arrangements []int) int {
	noDamagedSpringsRemaining := func() bool { return -1 == strings.Index(input, "#") }

	// ..? && []
	if noDamagedSpringsRemaining() && len(arrangements) == 0 {
		return 1
	}

	// ..? || []
	if len(input) == 0 || len(arrangements) == 0 {
		return 0
	}

	if input[0] == '.' {
		// .??? => evaluate(???)
		return evaluate(input[1:], arrangements)
	}

	if input[0] == '?' {
		// ?.?? => evaluate(.??) + evaluate(#.??)
		return evaluate("#"+input[1:], arrangements) + evaluate(input[1:], arrangements)
	}

	firstDotIndex := strings.Index(input, ".")
	firstQIndex := strings.Index(input, "?")

	if firstDotIndex == arrangements[0] {
		// ##.??? 2, 1 => evaluate(???, 1)
		return evaluate(input[firstDotIndex:], arrangements[1:])
	}

	if firstQIndex == arrangements[0] {
		// ##.? 3
		if firstDotIndex < firstQIndex && firstDotIndex > -1 {
			return 0
		}

		// ##?.?? 2, 1 => evaluate(.??, 1)
		return evaluate(input[firstQIndex+1:], arrangements[1:])
	}

	// ### 1
	if firstDotIndex == -1 && firstQIndex == -1 && len(input) == arrangements[0] && len(arrangements) == 1 {
		return 1
	}

	// ##?? 4 => evaluate(###?, 4)
	if firstQIndex != -1 && firstQIndex < arrangements[0] {
		return evaluate(input[:firstQIndex]+"#"+input[firstQIndex+1:], arrangements)
	}

	return 0
}
