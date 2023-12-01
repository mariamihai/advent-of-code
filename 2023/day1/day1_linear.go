package day1

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/mariamihai/advent-of-code/util"
)

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var calibrationValuesSum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		calibrationValuesSum += calibratedValueForLine(line)
	}

	return calibrationValuesSum
}

func Problem2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var calibrationValuesSum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = replacedLettersToDigits(line)
		calibrationValuesSum += calibratedValueForLine(line)
	}

	return calibrationValuesSum
}

func calibratedValueForLine(line string) int {
	emptyLine := len(line) == 0
	if emptyLine {
		return 0
	}

	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(line, -1)

	noNumbersInLine := len(numbers) == 0
	if noNumbersInLine {
		return 0
	}

	lastCnt := len(numbers) - 1

	calibration := numbers[0][0:1] + numbers[lastCnt][len(numbers[lastCnt])-1:]

	return util.StringToInt(calibration)
}

func replacedLettersToDigits(line string) string {
	line = strings.ReplaceAll(line, "one", "o1ne")
	line = strings.ReplaceAll(line, "two", "t2wo")
	line = strings.ReplaceAll(line, "three", "thr3ee")
	line = strings.ReplaceAll(line, "four", "fo4ur")
	line = strings.ReplaceAll(line, "five", "fi5ve")
	line = strings.ReplaceAll(line, "six", "s6ix")
	line = strings.ReplaceAll(line, "seven", "se7ven")
	line = strings.ReplaceAll(line, "eight", "ei8ght")
	line = strings.ReplaceAll(line, "nine", "ni9ne")

	return line
}
