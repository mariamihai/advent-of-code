package day14

import (
	"bufio"
	"github.com/samber/lo"
	"os"
	"regexp"
	"strings"

	"github.com/mariamihai/advent-of-code/util"
)

func Problem1(filename string) int {
	input := readInput(filename)

	var rocksOnColumns [][]int
	for i := 0; i < len(input); i++ {
		rocksOnColumns = append(rocksOnColumns, getRocksForColumn(input, i))
	}

	tiltAll(&input, rocksOnColumns, tiltToN)

	return calculateResult(input)
}

func Problem2(filename string) int {
	f, err := os.Create("./day14/out.txt")
	util.Boom(err)
	defer f.Close()

	input := readInput(filename)

	var rocksOnColumns [][]int
	var rocksOnRows [][]int
	// nrRows = nrColumns
	for i := 0; i < len(input); i++ {
		rocksOnRows = append(rocksOnRows, getRocksForRow(input[i]))
		rocksOnColumns = append(rocksOnColumns, getRocksForColumn(input, i))
	}

	iterations := 3
	for i := 1; i <= iterations; i++ {
		//_, err = f.WriteString("----- TILT TO N -----\n")
		//util.Boom(err)
		tiltAll(&input, rocksOnColumns, tiltToN)

		//_, err = f.WriteString(printMap(input))
		//util.Boom(err)

		//_, err = f.WriteString("\n\n----- TILT TO W -----\n")
		//util.Boom(err)
		tiltAll(&input, rocksOnRows, tiltToW)
		//_, err = f.WriteString(printMap(input))
		//util.Boom(err)

		//_, err = f.WriteString("\n\n----- TILT TO S -----\n")
		//util.Boom(err)
		tiltAll(&input, rocksOnColumns, tiltToS)
		//_, err = f.WriteString(printMap(input))
		//util.Boom(err)

		_, err = f.WriteString("\n\n----- TILT TO E -----\n")
		util.Boom(err)
		tiltAll(&input, rocksOnRows, tiltToE)
		_, err = f.WriteString(printMap(input))
		util.Boom(err)
	}

	return calculateResult(input)
}

func readInput(filename string) [][]string {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var input [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, strings.Split(line, ""))
	}

	return input
}

func tiltAll(input *[][]string, rocksOnColumns [][]int, tiltTo func(input *[][]string, column, start, end int)) {
	for i := 0; i < len(*input); i++ {
		tiltOne(input, i, rocksOnColumns[i], tiltTo)
	}
}

func tiltOne(input *[][]string, column int, cubeRockPositions []int, tiltTo func(input *[][]string, column, start, end int)) {
	startPosition := 0
	for i := 0; i < len(cubeRockPositions); i++ {
		endPosition := cubeRockPositions[i]

		if endPosition-startPosition <= 1 {
			continue
		}

		tiltTo(input, column, startPosition, endPosition)

		startPosition = endPosition
	}
}

// N = +1
// S = -1
func tiltingNS(direction int) func(input *[][]string, column, rowStart, rowEnd int) {
	return func(input *[][]string, column, rowStart, rowEnd int) {
		hasChanges := true

		for hasChanges {
			hasChanges = false

			for i := rowStart; i < rowEnd-1; i++ {
				if (*input)[i][column] == "." && (*input)[i+direction][column] == "O" {
					temp := (*input)[i][column]
					(*input)[i][column] = (*input)[i+direction][column]
					(*input)[i+direction][column] = temp
					hasChanges = true
				}
			}
		}
	}
}

// W = +1
// E = -1
func tiltToWE(direction int) func(input *[][]string, row, direction, rowStart, rowEnd int) {
	return func(input *[][]string, row, direction, rowStart, rowEnd int) {
		hasChanges := true

		for hasChanges {
			hasChanges = false

			for i := rowStart; i < rowEnd-1; i++ {
				if (*input)[row][i] == "." && (*input)[row][i+direction] == "O" {
					temp := (*input)[row][i]
					(*input)[row][i] = (*input)[row][i+direction]
					(*input)[row][i+direction] = temp
					hasChanges = true
				}
			}
		}
	}
}

func tiltToN(input *[][]string, column, rowStart, rowEnd int) {
	hasChanges := true

	for hasChanges {
		hasChanges = false

		for i := rowStart; i < rowEnd-1; i++ {
			if (*input)[i][column] == "." && (*input)[i+1][column] == "O" {
				temp := (*input)[i][column]
				(*input)[i][column] = (*input)[i+1][column]
				(*input)[i+1][column] = temp
				hasChanges = true
			}
		}
	}
}

func tiltToS(input *[][]string, column, rowStart, rowEnd int) {
	hasChanges := true

	for hasChanges {
		hasChanges = false

		for i := rowStart + 1; i < rowEnd; i++ {
			if (*input)[i][column] == "." && (*input)[i-1][column] == "O" {
				temp := (*input)[i][column]
				(*input)[i][column] = (*input)[i-1][column]
				(*input)[i-1][column] = temp
				hasChanges = true
			}
		}
	}
}
func tiltToW(input *[][]string, row, rowStart, rowEnd int) {
	hasChanges := true

	for hasChanges {
		hasChanges = false

		for i := rowStart; i < rowEnd-1; i++ {
			if (*input)[row][i] == "." && (*input)[row][i+1] == "O" {
				temp := (*input)[row][i]
				(*input)[row][i] = (*input)[row][i+1]
				(*input)[row][i+1] = temp
				hasChanges = true
			}
		}
	}
}
func tiltToE(input *[][]string, row, rowStart, rowEnd int) {
	hasChanges := true

	for hasChanges {
		hasChanges = false

		for i := rowStart + 1; i < rowEnd; i++ {
			if (*input)[row][i] == "." && (*input)[row][i-1] == "O" {
				temp := (*input)[row][i]
				(*input)[row][i] = (*input)[row][i-1]
				(*input)[row][i-1] = temp
				hasChanges = true
			}
		}
	}
}

func calculateResult(input [][]string) int {
	result := 0

	for i := 0; i < len(input); i++ {
		nrOfOccurrences := regexp.MustCompile("O").FindAllString(strings.Join(input[i], ""), -1)
		result += len(nrOfOccurrences) * (len(input) - i)
	}

	return result
}

func getRocksForColumn(input [][]string, column int) []int {
	var rocks []int

	for i := 0; i < len(input); i++ {
		if input[i][column] == "#" {
			rocks = append(rocks, i)
		}
	}

	if len(rocks) == 0 {
		rocks = append(rocks, len(input))
	}

	// If the last
	if rocks[len(rocks)-1] != len(input)-1 {
		rocks = append(rocks, len(input))
	}

	return rocks
}

func getRocksForRow(inputLine []string) []int {
	line := strings.Join(inputLine, "")
	rocksIndexes := regexp.MustCompile("#").FindAllStringIndex(line, -1)

	positions := []int{}

	positions = lo.Map(rocksIndexes, func(item []int, _ int) int {
		return item[0]
	})

	if len(positions) == 0 {
		positions = append(positions, len(inputLine))
		return positions
	}

	if positions[len(positions)-1] != len(inputLine)-1 {
		positions = append(positions, len(inputLine))
	}

	return positions
}

func printMap(input [][]string) string {
	var output string

	for i := 0; i < len(input); i++ {
		output += strings.Join(input[i], "") + "\n"
	}

	return output
}
