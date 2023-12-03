package day3

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"regexp"
)

type numbersInfo struct {
	number string
	x      int
	y      int
}

func Problem2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var allLines []string
	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		allLines = append(allLines, line)
	}

	allNumbersInfo := getAllNumbersInfo(allLines)

	for i := 0; i < len(allLines); i++ {
		allGears := getAllGearsLocationsForLine(allLines[i])

		if len(allGears) == 0 {
			continue
		}

		for _, gearLocation := range allGears {
			var isValid int
			multiplicationForGear := 1

			// Check line before
			if i > 0 {
				contains, number := checkAdjacentLine(gearLocation, allNumbersInfo[i-1])
				isValid += contains
				if contains == 1 {
					multiplicationForGear *= number
				}
			}

			// Check current line
			contains, number := checkSameLine(gearLocation, allNumbersInfo[i])
			isValid += contains
			if contains == 1 {
				multiplicationForGear *= number
			}

			// Check line after
			if i < len(allLines)-1 {
				contains, number := checkAdjacentLine(gearLocation, allNumbersInfo[i+1])
				isValid += contains
				if contains == 1 {
					multiplicationForGear *= number
				}
			}

			if isValid == 2 {
				sum += multiplicationForGear
			}
		}
	}

	return sum
}

func checkSameLine(gearLocation int, numbers []numbersInfo) (int, int) {
	if len(numbers) == 0 {
		return 0, 1
	}

	for _, number := range numbers {
		if number.x-1 == gearLocation || number.y+1 == gearLocation {
			return 1, util.StringToInt(number.number)
		}
	}

	return 0, 1
}

func checkAdjacentLine(gearLocation int, numbers []numbersInfo) (int, int) {
	if len(numbers) == 0 {
		return 0, 1
	}

	for _, number := range numbers {
		if number.x-1 <= gearLocation && gearLocation <= number.y+1 {
			return 1, util.StringToInt(number.number)
		}
	}

	return 0, 1
}

func getAllGearsLocationsForLine(line string) []int {
	var result []int

	re := regexp.MustCompile("[*]")
	allGears := re.FindAllStringIndex(line, -1)

	if len(allGears) == 0 {
		return result
	}

	for _, location := range allGears {
		result = append(result, location[0])
	}

	return result
}

func getAllNumbersInfo(allLines []string) map[int][]numbersInfo {
	allNumbersInfo := make(map[int][]numbersInfo)

	for i := 0; i < len(allLines); i++ {
		var lineNumbersInfo []numbersInfo

		line := allLines[i]

		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(line, -1)
		numbersIndexes := re.FindAllStringIndex(line, -1)

		// Ignore line if there are no numbers on it
		if len(numbers) == 0 {
			continue
		}

		for index, number := range numbers {
			lineNumbersInfo = append(lineNumbersInfo, numbersInfo{
				number: number,
				x:      numbersIndexes[index][0],
				y:      numbersIndexes[index][1] - 1,
			})
		}

		allNumbersInfo[i] = lineNumbersInfo
	}

	return allNumbersInfo
}
