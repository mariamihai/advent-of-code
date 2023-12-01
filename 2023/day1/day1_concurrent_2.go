package day1

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"regexp"
)

func Problem1Concurrent2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var calibrationValuesSum int

	threads := 4
	results := make(chan int, threads)
	sem := make(chan bool, threads)

	var cnt int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sem <- true

		line := scanner.Text()
		cnt++

		go func(line string) {
			result := calibratedValueForLineConcurrent2(line)
			defer func() {
				<-sem
				results <- result
			}()
		}(line)
	}

	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	for i := 0; i < cnt; i++ {
		select {
		case lineResult := <-results:
			calibrationValuesSum += lineResult
		}
	}

	return calibrationValuesSum
}

func calibratedValueForLineConcurrent2(line string) int {
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
