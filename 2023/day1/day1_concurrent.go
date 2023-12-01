package day1

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"regexp"
	"sync"
)

func Problem1Concurrent(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var calibrationValuesSum int

	ch := make(chan int)
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		wg.Add(1)
		go calibratedValueForLineConcurrent(line, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for lineResult := range ch {
		calibrationValuesSum += lineResult
	}

	return calibrationValuesSum
}

func calibratedValueForLineConcurrent(line string, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	emptyLine := len(line) == 0
	if emptyLine {
		ch <- 0
	}

	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(line, -1)

	noNumbersInLine := len(numbers) == 0
	if noNumbersInLine {
		ch <- 0
	}

	lastCnt := len(numbers) - 1

	calibration := numbers[0][0:1] + numbers[lastCnt][len(numbers[lastCnt])-1:]

	ch <- util.StringToInt(calibration)
}
