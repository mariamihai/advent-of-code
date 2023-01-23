package util

import (
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
