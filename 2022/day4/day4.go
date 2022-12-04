package day4

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ZoneLimits struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// Problem1 - find the number of zones contained in another, for each pair
func Problem1() {
	file, err := os.Open("./day4/input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	count := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(customSplit())

	for scanner.Scan() {
		line := scanner.Bytes()

		var zones []ZoneLimits
		err := json.Unmarshal(line, &zones)
		if err != nil {
			log.Fatal(err)
		}

		if isZoneIncluded(zones[0], zones[1]) {
			count++
		}
	}

	fmt.Printf("Part 1 - number of pairs with zones included one in another: %d\n", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func customSplit() func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Return nothing if at end of file and no data passed
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Find the index of the input of a newline
		if i := strings.Index(string(data), "\n"); i >= 0 {
			return i + 1, createCustomData(string(data[0:i])), nil
		}

		// If at end of file with data return the data
		if atEOF {
			return len(data), createCustomData(string(data)), nil
		}

		return
	}
}

func createCustomData(data string) []byte {
	var lineZones []ZoneLimits

	pairZonesAsStrings := strings.Split(data, ",")
	zone1 := generateZoneFromString(pairZonesAsStrings[0])
	zone2 := generateZoneFromString(pairZonesAsStrings[1])

	lineZones = append(lineZones, zone1, zone2)

	bytes, err := json.Marshal(lineZones)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

// Problem2 - find the number of zones contained or overlapping in another, for each pair
func Problem2() {
	file, err := os.Open("./day4/input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	count := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(customSplit())
	for scanner.Scan() {
		line := scanner.Bytes()

		var zones []ZoneLimits
		err := json.Unmarshal(line, &zones)
		if err != nil {
			log.Fatal(err)
		}

		if isZoneIncluded(zones[0], zones[1]) || isZoneOverlapping(zones[0], zones[1]) {
			count++
		}
	}

	fmt.Printf("Part 2.2 - number of pairs with zones included one in another or overlapping: %d\n", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func generateZoneFromString(zoneAsString string) ZoneLimits {
	valuesAsString := strings.Split(zoneAsString, "-")

	return ZoneLimits{
		Start: stringToInt(valuesAsString[0]),
		End:   stringToInt(valuesAsString[1]),
	}
}

func stringToInt(valueAsString string) int {
	valueAsInt, err := strconv.Atoi(valueAsString)
	if err != nil {
		log.Fatalf("Error during conversion for: %s", valueAsString)
	}

	return valueAsInt
}

func isZoneIncluded(zone1, zone2 ZoneLimits) bool {
	// Either zone1 is included in zone2 or the other way around
	return (zone1.Start >= zone2.Start && zone1.End <= zone2.End) ||
		(zone1.Start <= zone2.Start && zone1.End >= zone2.End)
}

func isZoneOverlapping(zone1, zone2 ZoneLimits) bool {
	return (zone1.Start <= zone2.Start && zone1.End >= zone2.Start) || (zone1.Start <= zone2.End && zone1.End >= zone2.End)
}
