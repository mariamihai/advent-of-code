package day5

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/samber/lo"
	"strings"
)

type seedWithRange struct {
	min         int
	rangeLength int
}

func Problem2Messy(filename string) int {
	// TODO rewrite this mess
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

	_ = calculateForProblem2(seeds, mappers)
	return 0
}

func getSeedsForProblem2(line string) []seedWithRange {
	line = strings.Replace(line, "seeds:", "", 1)
	ints := util.StringToIntSlice(line)

	seeds := []seedWithRange{}
	for i := 0; i < len(ints); i = i + 2 {
		seeds = append(seeds, seedWithRange{
			min:         ints[i],
			rangeLength: ints[i+1],
		})
	}

	return seeds
}

func calculateForProblem2(seeds []seedWithRange, mappers map[string][]almanacLine) []seedWithRange {
	var mappedSeeds []seedWithRange
	// -- seed-to-soil map:
	mappedSeeds = mapSeedsForProblem2(seeds, mappers["seed-to-soil"])
	mappedSeeds = mapSeedsForProblem2(mappedSeeds, mappers["soil-to-fertilizer"])

	mappedSeeds = mapSeedsForProblem2(mappedSeeds, mappers["fertilizer-to-water"])
	mappedSeeds = mapSeedsForProblem2(mappedSeeds, mappers["water-to-light"])
	mappedSeeds = mapSeedsForProblem2(mappedSeeds, mappers["light-to-temperature"])
	mappedSeeds = mapSeedsForProblem2(mappedSeeds, mappers["temperature-to-humidity"])
	mappedSeeds = mapSeedsForProblem2(mappedSeeds, mappers["humidity-to-location"])

	return mappedSeeds
}

func mapSeedsForProblem2(seeds []seedWithRange, converter []almanacLine) []seedWithRange {
	result := seeds

	for _, c := range converter {
		converterResult := []seedWithRange{}

		for _, seed := range result {
			converterResult = append(converterResult, mapSeedForProblem2(seed, c)...)
		}

		result = restrictedResult(converterResult)
	}

	return result
}

func restrictedResult(converterResult []seedWithRange) []seedWithRange {
	result := map[int]seedWithRange{}

	for _, seed := range converterResult {
		value, exists := result[seed.min]
		if exists {
			result[seed.min] = seedWithRange{
				min:         result[seed.min].min,
				rangeLength: lo.Max([]int{value.rangeLength, seed.rangeLength}),
			}
			continue
		}

		result[seed.min] = seedWithRange{
			min:         seed.min,
			rangeLength: seed.rangeLength,
		}
	}

	return lo.MapToSlice(result, func(k int, v seedWithRange) seedWithRange {
		return v
	})
}

func mapSeedForProblem2(seed seedWithRange, c almanacLine) []seedWithRange {
	result := []seedWithRange{}
	seedMin := seed.min
	seedMax := seed.min + seed.rangeLength

	convertMin := c.source
	convertMax := c.source + c.rangeLength

	//    [ A, B]          seed
	//            [ C, D]  convert
	if seedMax < convertMin {
		return []seedWithRange{seed}
	}

	//             [ C, D]  seed
	//    [ A, B]           convert
	if convertMax < seedMin {
		return []seedWithRange{seed}
	}

	//leftMax := lo.Max([]int{seed.min, c.source})
	//rightMin := lo.Min([]int{seed.min + seed.rangeLength, c.source + c.rangeLength})

	//        [ B, C]      seed
	//    [ A,         D]  convert
	if convertMin <= seedMin && seedMax <= convertMax {
		result = append(result, seedWithRange{
			min:         c.destination,
			rangeLength: seed.rangeLength,
		})

		return result
	}

	//    [ A,    C]     seed
	//        [ B,   D]  convert
	if seedMin <= convertMin && convertMin <= seedMax && seedMax <= convertMax {
		result = append(result, seedWithRange{
			min:         c.destination,
			rangeLength: seedMax - convertMin,
		})

		seed.rangeLength = seed.rangeLength - (seedMax - convertMin)
		result = append(result, seed)
		return result
	}

	//    	[ B,    D]  seed
	//    [ A,   C]     convert
	if convertMin <= seedMin && seedMin <= convertMax {
		result = append(result, seedWithRange{
			min:         c.destination,
			rangeLength: convertMax - seedMin,
		})

		seed.min = convertMax
		seed.rangeLength = seedMax - convertMax
		result = append(result, seed)
		return result
	}

	//    [ A,         D]  seed
	//        [ B, C]      convert
	if seedMin <= convertMin && convertMin <= seedMax {
		result = append(result, seedWithRange{
			min:         c.destination,
			rangeLength: c.rangeLength,
		})
		// Before
		result = append(result, seedWithRange{
			min:         seedMin,
			rangeLength: convertMin - seedMin,
		})

		// After
		result = append(result, seedWithRange{
			min:         convertMax,
			rangeLength: seedMax - convertMax,
		})
	}

	return result
}
