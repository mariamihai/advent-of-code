package day10

import (
	"bufio"
	"strings"

	"github.com/mariamihai/advent-of-code/util"
)

type Point struct {
	X, Y int
}

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var puzzle [][]string
	var cnt int
	var sPosition Point

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		sIndex := strings.Index(line, "S")
		if sIndex != -1 {
			sPosition.X = sIndex
			sPosition.Y = cnt
		}

		puzzle = append(puzzle, strings.Split(line, ""))
		cnt++
	}

	paths := getConnectionPipesForStartingPoint(puzzle, sPosition)

	// Clean up S -> already passed
	puzzle[sPosition.X][sPosition.Y] = "."

	cnt = 0

	for true {
		cnt++
		var newPaths []Point

		for _, path := range paths {
			newPaths = append(newPaths, getConnectionPipesForPoint(puzzle, path)...)
		}

		paths = newPaths

		if len(paths) == 2 && newPaths[0].X == newPaths[1].X && newPaths[1].Y == newPaths[1].Y {
			break
		}

		if cnt == 10 {
			break
		}
	}

	return cnt + 1
}

func getConnectionPipesForPoint(puzzle [][]string, p Point) []Point {
	var results []Point

	// Check left and right only
	if puzzle[p.X][p.Y] == "-" {
		// Check left
		isEConnection, point := checkE(puzzle, p.X, p.Y)
		if isEConnection {
			results = append(results, *point)
		}

		// Check right
		isWConnection, point := checkW(puzzle, p.X, p.Y)
		if isWConnection {
			results = append(results, *point)
		}
	}

	// Check up and down only
	if puzzle[p.X][p.Y] == "|" {
		// Check up
		isNConnection, point := checkN(puzzle, p.X, p.Y)
		if isNConnection {
			results = append(results, *point)
		}

		// Check down
		isSConnection, point := checkS(puzzle, p.X, p.Y)
		if isSConnection {
			results = append(results, *point)
		}
	}

	// Check up and right only
	if puzzle[p.X][p.Y] == "L" {
		// Check up
		isNConnection, point := checkN(puzzle, p.X, p.Y)
		if isNConnection {
			results = append(results, *point)
		}

		// Check right
		isWConnection, point := checkW(puzzle, p.X, p.Y)
		if isWConnection {
			results = append(results, *point)
		}
	}

	// Check up and left only
	if puzzle[p.X][p.Y] == "J" {
		// Check up
		isNConnection, point := checkN(puzzle, p.X, p.Y)
		if isNConnection {
			results = append(results, *point)
		}

		// Check left
		isEConnection, point := checkE(puzzle, p.X, p.Y)
		if isEConnection {
			results = append(results, *point)
		}
	}

	// Check down and right only
	if puzzle[p.X][p.Y] == "F" {
		// Check down
		isSConnection, point := checkS(puzzle, p.X, p.Y)
		if isSConnection {
			results = append(results, *point)
		}

		// Check right
		isWConnection, point := checkW(puzzle, p.X, p.Y)
		if isWConnection {
			results = append(results, *point)
		}
	}

	// Check down and left only
	if puzzle[p.X][p.Y] == "7" {
		// Check down
		isSConnection, point := checkS(puzzle, p.X, p.Y)
		if isSConnection {
			results = append(results, *point)
		}

		// Check left
		isEConnection, point := checkE(puzzle, p.X, p.Y)
		if isEConnection {
			results = append(results, *point)
		}
	}

	return results
}

func getConnectionPipesForStartingPoint(puzzle [][]string, p Point) []Point {
	var results []Point

	// Check up
	isNConnection, point := checkN(puzzle, p.X, p.Y)
	if isNConnection {
		results = append(results, *point)
	}

	// Check down
	isSConnection, point := checkS(puzzle, p.X, p.Y)
	if isSConnection {
		results = append(results, *point)
	}

	// Check left
	isEConnection, point := checkE(puzzle, p.X, p.Y)
	if isEConnection {
		results = append(results, *point)
	}

	// Check right
	isWConnection, point := checkW(puzzle, p.X, p.Y)
	if isWConnection {
		results = append(results, *point)
	}

	return results
}
