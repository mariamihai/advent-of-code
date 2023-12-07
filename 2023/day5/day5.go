package day5

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/samber/lo"
	"math"
	"strings"
)

type almanacLine struct {
	destination int
	source      int
	rangeLength int
}

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var seeds []int
	var changedStep bool
	var step string

	mappers := make(map[string][]almanacLine)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			seeds = getSeedsForProblem1(line)
			continue
		}

		changedStep, step = mapTo(line, step)
		if changedStep {
			continue
		}

		if mappers[step] == nil {
			mappers[step] = []almanacLine{}
		}

		mappers[step] = append(mappers[step], mapToAlmanacLine(line))
	}

	result := calculate(seeds, mappers)

	return lo.Min(result)
}

func Problem2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var seeds []seedWithRange
	var changedStep bool
	var step string

	mappers := make(map[string][]almanacLine)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			seeds = getSeedsForProblem2(line)
			continue
		}

		changedStep, step = mapTo(line, step)
		if changedStep {
			continue
		}

		if mappers[step] == nil {
			mappers[step] = []almanacLine{}
		}

		mappers[step] = append(mappers[step], mapToAlmanacLine(line))
	}

	min := math.MaxInt

	for _, seed := range seeds {
		currentResult := calculate(lo.RangeFrom(seed.min, seed.rangeLength), mappers)
		currentMin := lo.Min(currentResult)

		if currentMin < min {
			min = currentMin
		}
	}

	return min
}

func mapTo(line, originalStep string) (bool, string) {
	if strings.Contains(line, "seed-to-soil map:") {
		return true, "seed-to-soil"
	}
	if strings.Contains(line, "soil-to-fertilizer map:") {
		return true, "soil-to-fertilizer"
	}
	if strings.Contains(line, "fertilizer-to-water map:") {
		return true, "fertilizer-to-water"
	}
	if strings.Contains(line, "water-to-light map:") {
		return true, "water-to-light"
	}
	if strings.Contains(line, "light-to-temperature map:") {
		return true, "light-to-temperature"
	}
	if strings.Contains(line, "temperature-to-humidity map:") {
		return true, "temperature-to-humidity"
	}
	if strings.Contains(line, "humidity-to-location map:") {
		return true, "humidity-to-location"
	}

	if line == "" {
		return true, ""
	}

	return false, originalStep
}

func calculate(seeds []int, mappers map[string][]almanacLine) []int {
	var mappedSeeds []int
	// -- seed-to-soil map:
	mappedSeeds = mapSeeds(seeds, mappers["seed-to-soil"])
	mappedSeeds = mapSeeds(mappedSeeds, mappers["soil-to-fertilizer"])

	mappedSeeds = mapSeeds(mappedSeeds, mappers["fertilizer-to-water"])
	mappedSeeds = mapSeeds(mappedSeeds, mappers["water-to-light"])
	mappedSeeds = mapSeeds(mappedSeeds, mappers["light-to-temperature"])
	mappedSeeds = mapSeeds(mappedSeeds, mappers["temperature-to-humidity"])
	mappedSeeds = mapSeeds(mappedSeeds, mappers["humidity-to-location"])

	return mappedSeeds
}

func getSeedsForProblem1(line string) []int {
	line = strings.Replace(line, "seeds:", "", 1)
	return util.StringToIntSlice(line)
}

func mapToAlmanacLine(line string) almanacLine {
	lineInput := util.StringToIntSlice(line)

	return almanacLine{
		destination: lineInput[0],
		source:      lineInput[1],
		rangeLength: lineInput[2],
	}
}

func mapSeeds(seeds []int, converter []almanacLine) []int {
	result := []int{}

	for _, seed := range seeds {
		result = append(result, mapSeed(seed, converter))
	}

	return result
}

func mapSeed(seed int, converter []almanacLine) int {
	for _, c := range converter {
		if c.source <= seed && seed <= c.source+c.rangeLength {
			return c.destination + seed - c.source
		}
	}

	return seed
}
