package day4

import (
	"advent-of-code-2022/util"
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type ZoneLimits struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// Problem - find the number of zones contained in another, for each pair
// AND find the number of zones contained or overlapping in another, for each pair
func Problem() func() {
	return func() {
		file := util.ReadFile("./day4/input2.txt")
		defer util.CloseFile()(file)

		countProblem1 := 0
		countProblem2 := 0

		scanner := bufio.NewScanner(file)
		scanner.Split(util.CustomSplit(createCustomData()))

		for scanner.Scan() {
			line := scanner.Bytes()

			var zones []ZoneLimits
			err := json.Unmarshal(line, &zones)
			util.Boom(err)

			// Problem 1 requirements
			if isZoneIncluded(zones[0], zones[1]) {
				countProblem1++
			}

			// Problem 2 requirements (keeping the 2 separate)
			if isZoneIncluded(zones[0], zones[1]) || isZoneOverlapping(zones[0], zones[1]) {
				countProblem2++
			}
		}

		fmt.Printf("Part 1 - number of pairs with zones included one in another: %d\n", countProblem1)
		fmt.Printf("Part 2 - number of pairs with zones included one in another or overlapping: %d\n", countProblem2)

		err := scanner.Err()
		util.Boom(err)
	}
}

func createCustomData() func(data string) []byte {
	return func(data string) []byte {
		var lineZones []ZoneLimits

		pairZonesAsStrings := strings.Split(data, ",")
		zone1 := generateZoneFromString(pairZonesAsStrings[0])
		zone2 := generateZoneFromString(pairZonesAsStrings[1])

		lineZones = append(lineZones, zone1, zone2)

		bytes, err := json.Marshal(lineZones)
		util.Boom(err)

		return bytes
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
	util.Boom(err)

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
