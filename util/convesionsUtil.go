package util

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
)

func StringToInt(valueAsString string) int {
	valueAsInt, err := strconv.Atoi(valueAsString)
	Boom(err)

	return valueAsInt
}

func StringToInt64(valueAsString string) int64 {
	valueAsInt, err := strconv.Atoi(valueAsString)
	Boom(err)

	return int64(valueAsInt)
}

func NumbersAsStringSlice(str string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(str, -1)
}

func StringToIntSlice(str string) []int {
	return lo.Map(NumbersAsStringSlice(str), func(x string, index int) int {
		result, err := strconv.Atoi(x)
		if err != nil {
			// Ignore error
			return 0
		}
		return result
	})
}

func StringSliceToIntSlice(input []string) []int {
	return lo.Map(input, func(x string, index int) int {
		result, err := strconv.Atoi(x)
		if err != nil {
			return 0
		}
		return result
	})
}

func AbsInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
