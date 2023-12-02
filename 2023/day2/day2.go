package day2

import (
	"bufio"
	"encoding/json"
	"strings"

	"github.com/mariamihai/advent-of-code/util"
)

type Game struct {
	Number  int
	Subsets []Subset
}

type Subset struct {
	Red   int
	Green int
	Blue  int
}

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var sum int

	scanner := bufio.NewScanner(file)
	scanner.Split(util.CustomSplit(createCustomData()))

	for scanner.Scan() {
		line := scanner.Bytes()

		var gameEntry Game
		err := json.Unmarshal(line, &gameEntry)
		util.Boom(err)

		validGame := true

		for _, inputSubset := range gameEntry.Subsets {
			if isGamePossible(inputSubset.Red, inputSubset.Green, inputSubset.Blue) {
				validGame = false
			}
		}

		if validGame {
			sum += gameEntry.Number
		}
	}

	return sum
}

// isGamePossible checks if subset is valid against maximum values accepted
// (max red cubes = 12, max green cubes = 13 and max blue cubes = 14)
func isGamePossible(red, green, blue int) bool {
	return red > 12 || green > 13 || blue > 14
}

func Problem2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var sum int

	scanner := bufio.NewScanner(file)
	scanner.Split(util.CustomSplit(createCustomData()))

	for scanner.Scan() {
		line := scanner.Bytes()

		var gameEntry Game
		err := json.Unmarshal(line, &gameEntry)
		util.Boom(err)

		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, inputSubset := range gameEntry.Subsets {
			if minRed < inputSubset.Red {
				minRed = inputSubset.Red
			}
			if minGreen < inputSubset.Green {
				minGreen = inputSubset.Green
			}
			if minBlue < inputSubset.Blue {
				minBlue = inputSubset.Blue
			}
		}

		sum += minRed * minGreen * minBlue
	}

	return sum
}

// createCustomData maps a line to a structure containing a slice of RGB values + the number of the day
// Eg.: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func createCustomData() func(data string) []byte {
	return func(data string) []byte {
		splitLine := strings.Split(data, ": ")

		var gameEntry Game
		gameEntry.Number = util.StringToInt(strings.Replace(splitLine[0], "Game ", "", 1))
		gameEntry.Subsets = []Subset{}

		inputSubsets := strings.Split(splitLine[1], "; ")

		calculateColorValue := func(value, color string) int { return util.StringToInt(strings.Replace(value, " "+color, "", 1)) }

		for _, inputSubset := range inputSubsets {
			values := strings.Split(inputSubset, ", ")

			var redInput, greenInput, blueInput int

			for _, value := range values {
				if strings.Contains(value, "red") {
					redInput = calculateColorValue(value, "red")
				}
				if strings.Contains(value, "green") {
					greenInput = calculateColorValue(value, "green")
				}
				if strings.Contains(value, "blue") {
					blueInput = calculateColorValue(value, "blue")
				}
			}

			gameEntry.Subsets = append(gameEntry.Subsets, Subset{Red: redInput, Green: greenInput, Blue: blueInput})
		}

		bytes, err := json.Marshal(gameEntry)
		util.Boom(err)

		return bytes
	}
}
