package util

import (
	"log"
	"strconv"
)

func StringToInt(valueAsString string) int {
	valueAsInt, err := strconv.Atoi(valueAsString)

	if err != nil {
		log.Fatalf("Error during conversion for: %s", valueAsString)
	}

	return valueAsInt
}

func StringToInt64(valueAsString string) int64 {
	valueAsInt, err := strconv.Atoi(valueAsString)

	if err != nil {
		log.Fatalf("Error during conversion for: %s", valueAsString)
	}

	return int64(valueAsInt)
}
