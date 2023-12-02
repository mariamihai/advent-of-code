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

	maxRedCubes := 12
	maxGreeenCubes := 13
	maxBlueCubes := 14
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
			if inputSubset.Red > maxRedCubes ||
				inputSubset.Green > maxGreeenCubes ||
				inputSubset.Blue > maxBlueCubes {
				validGame = false
			}
		}

		//fmt.Printf("[%t] %+v\n", validGame, gameEntry)

		if validGame {
			sum += gameEntry.Number
		}
	}

	return sum
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

func createCustomData() func(data string) []byte {
	return func(data string) []byte {
		split := strings.Split(data, ": ")

		var gameEntry Game
		gameEntry.Number = util.StringToInt(strings.Replace(split[0], "Game ", "", 1))
		gameEntry.Subsets = []Subset{}

		inputSubsets := strings.Split(split[1], "; ")

		for _, inputSubset := range inputSubsets {
			values := strings.Split(inputSubset, ", ")

			var redInput, greenInput, blueInput int

			// TODO redo this mess
			for _, value := range values {
				if strings.Contains(value, "red") {
					redInput = util.StringToInt(strings.Replace(value, " red", "", 1))
				}
				if strings.Contains(value, "green") {
					greenInput = util.StringToInt(strings.Replace(value, " green", "", 1))
				}
				if strings.Contains(value, "blue") {
					blueInput = util.StringToInt(strings.Replace(value, " blue", "", 1))
				}
			}

			gameEntry.Subsets = append(gameEntry.Subsets, Subset{Red: redInput, Green: greenInput, Blue: blueInput})
		}

		bytes, err := json.Marshal(&gameEntry)
		util.Boom(err)

		return bytes
	}
}
