package day6

import (
	"bufio"
	"github.com/samber/lo"
	"strconv"
	"strings"

	"github.com/mariamihai/advent-of-code/util"
)

func Problem1(filename string) int {
	input1, input2 := readInput(filename)
	times := util.StringToIntSlice(input1)
	distances := util.StringToIntSlice(input2)

	result := 1

	for i := 0; i < len(times); i++ {
		min := lo.FromPtr(getMin(times[i], distances[i]))
		max := lo.FromPtr(getMax(times[i], distances[i]))

		result *= max - min + 1
	}
	return result
}

func Problem2(filename string) int {
	input1, input2 := readInput(filename)

	time, _ := strconv.Atoi(strings.Replace(input1, " ", "", -1))
	distance, _ := strconv.Atoi(strings.Replace(input2, " ", "", -1))

	min := lo.FromPtr(getMin(time, distance))
	max := lo.FromPtr(getMax(time, distance))

	return max - min + 1
}

func readInput(filename string) (times string, distances string) {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	times = strings.Replace(scanner.Text(), "Time:", "", 1)
	scanner.Scan()
	distances = strings.Replace(scanner.Text(), "Distance:", "", 1)

	return
}

func getMin(time, distance int) *int {
	for i := 1; i <= time; i++ {
		if i*(time-i) > distance {
			return lo.ToPtr(i)
		}
	}
	return nil
}

func getMax(time, distance int) *int {
	for i := time; i >= 1; i-- {
		if i*(time-i) > distance {
			return lo.ToPtr(i)
		}
	}
	return nil
}
