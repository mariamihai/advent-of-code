package day11

import (
	"bufio"
	"github.com/samber/lo"
	"strings"

	"github.com/mariamihai/advent-of-code/util"
)

// -------------------------- Common --------------------------

type Point struct {
	X, Y int
}

// isEmptyRow check if the row contains only dots (no galaxy)
func isEmptyRow(row []string) bool {
	return len(lo.Uniq(row)) == 1
}

func isGalaxy(inputPosition string) bool {
	return inputPosition == "#"
}

func getAllGalaxies(input [][]string) []Point {
	galaxies := []Point{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if isGalaxy(input[i][j]) {
				galaxies = append(galaxies, Point{X: i, Y: j})
			}
		}
	}

	return galaxies
}

// -------------------------- Problem 1 --------------------------

func Problem1(filename string) int {
	input := getInputWithDuplicateEmptyRowsPart1(filename)
	input = getMapWithDuplicateEmptyColumnsPart1(input)

	allGalaxies := getAllGalaxies(input)

	var result int
	for _, galaxy := range allGalaxies {
		result += shortestPathsForGalaxyPart1(galaxy, allGalaxies)
	}

	return result / 2
}

func getInputWithDuplicateEmptyRowsPart1(filename string) [][]string {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var input [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		input = append(input, line)

		// Add a second empty row if one is found
		if isEmptyRow(line) {
			input = append(input, line)
		}
	}

	return input
}

func getMapWithDuplicateEmptyColumnsPart1(input [][]string) [][]string {
	var emptyColumnPositions []int
	for j := 0; j < len(input[0]); j++ {
		var isNotEmptyColumn bool

		for i := 0; i < len(input); i++ {
			if isGalaxy(input[i][j]) {
				isNotEmptyColumn = true
				break
			}
		}

		if !isNotEmptyColumn {
			emptyColumnPositions = append(emptyColumnPositions, j)
		}
	}

	for i := 0; i < len(input); i++ {
		for index, emptyPosition := range emptyColumnPositions {
			updatedRestOfTheRow := append([]string{"."}, input[i][emptyPosition+index:len(input[i])]...)

			input[i] = append(input[i][:emptyPosition+index], updatedRestOfTheRow...)
		}
	}

	return input
}

func shortestPathsForGalaxyPart1(current Point, allGalaxies []Point) int {
	isSameGalaxy := func(current, galaxy Point) bool { return current.X == galaxy.X && current.Y == galaxy.Y }

	var sum int

	for _, galaxy := range allGalaxies {
		if isSameGalaxy(current, galaxy) {
			continue
		}

		sum += util.AbsInt(current.X, galaxy.X) + util.AbsInt(current.Y, galaxy.Y)
	}

	return sum
}

// -------------------------- Problem 2 --------------------------

func Problem2(filename string, expansionMultiplier int) int {
	input, emptyRows := getInputAndEmptyRowsPart2(filename)
	emptyColumns := getEmptyColumnsPart2(input)

	allGalaxies := getAllGalaxies(input)

	var result int
	for _, p := range allGalaxies {
		result += shortestPathsForGalaxyPart2(p, allGalaxies, emptyRows, emptyColumns, expansionMultiplier)
	}

	return result / 2
}

func getInputAndEmptyRowsPart2(filename string) ([][]string, []int) {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var input [][]string
	var cnt int
	var emptyRows []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		input = append(input, line)

		if isEmptyRow(line) {
			emptyRows = append(emptyRows, cnt)
		}
		cnt++
	}

	return input, emptyRows
}

func getEmptyColumnsPart2(input [][]string) []int {
	var emptyColumns []int

	for j := 0; j < len(input[0]); j++ {
		var isNotEmpty bool

		for i := 0; i < len(input); i++ {
			if isGalaxy(input[i][j]) {
				isNotEmpty = true
				break
			}
		}

		if !isNotEmpty {
			emptyColumns = append(emptyColumns, j)
		}
	}

	return emptyColumns
}

func shortestPathsForGalaxyPart2(current Point, allGalaxies []Point, emptyRows, emptyColumns []int, expansionMultiplier int) int {
	isSameGalaxy := func(current, galaxy Point) bool { return current.X == galaxy.X && current.Y == galaxy.Y }

	var sum int

	for _, galaxy := range allGalaxies {
		if isSameGalaxy(current, galaxy) {
			continue
		}
		sum += (expansionMultiplier - 1) * areBetweenNumbers(current.X, galaxy.X, emptyRows...)
		sum += (expansionMultiplier - 1) * areBetweenNumbers(current.Y, galaxy.Y, emptyColumns...)

		sum += util.AbsInt(current.X, galaxy.X) + util.AbsInt(current.Y, galaxy.Y)
	}

	return sum
}

func areBetweenNumbers(a, b int, nrs ...int) int {
	var cnt int
	for _, nr := range nrs {
		if (a <= nr && nr <= b) ||
			(b <= nr && nr <= a) {
			cnt++
		}
	}

	return cnt
}
