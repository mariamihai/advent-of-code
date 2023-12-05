package day3

import (
	"bufio"
	"regexp"
	"strings"
	"unicode"

	"github.com/mariamihai/advent-of-code/util"
)

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var allLines []string
	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		allLines = append(allLines, line)
	}

	for i := 0; i < len(allLines); i++ {
		line := allLines[i]

		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(line, -1)
		numbersIndexes := re.FindAllStringIndex(line, -1)

		// Ignore line if there are no numbers on it
		if len(numbers) == 0 {
			continue
		}

		for index, number := range numbers {
			startOfNumber := numbersIndexes[index][0]
			endOfNumber := numbersIndexes[index][1] - 1

			var hasAdjancy bool

			// ----- Check prior line
			if i != 0 {
				if isSymbolOnAdjacentLine(allLines[i-1], startOfNumber-1, endOfNumber+1) {
					hasAdjancy = true
				}
			}

			// ----- Check current line
			if isSymbolBefore(line, startOfNumber) {
				hasAdjancy = true
			}
			if isSymbolAfter(line, endOfNumber) {
				hasAdjancy = true
			}

			// ----- Check next line
			if i != len(allLines)-1 {
				if isSymbolOnAdjacentLine(allLines[i+1], startOfNumber-1, endOfNumber+1) {
					hasAdjancy = true
				}
			}

			if hasAdjancy {
				sum += util.StringToInt(number)
			}
		}
	}

	return sum
}

func isSymbolOnAdjacentLine(line string, startIndex, endIndex int) bool {
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex >= len(line) {
		endIndex = len(line) - 1
	}

	for i := startIndex; i <= endIndex; i++ {
		if !strings.Contains(string(line[i]), ".") && !unicode.IsDigit(rune(line[i])) {
			return true
		}
	}

	return false
}

func isSymbolBefore(line string, index int) bool {
	if index <= 0 {
		return false
	}

	return !strings.Contains(string(line[index-1]), ".") && !unicode.IsDigit(rune(line[index-1]))
}

func isSymbolAfter(line string, index int) bool {
	if index >= len(line)-1 {
		return false
	}

	return !strings.Contains(string(line[index+1]), ".") && !unicode.IsDigit(rune(line[index+1]))
}
