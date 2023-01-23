package day2

import (
	"advent-of-code-2022/util"
	"bufio"
	"fmt"
	"strings"
)

func Problem1() func() {
	return func() {
		file := util.ReadFile("./day2/input2.txt")
		defer util.CloseFile()(file)

		totalScore := 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			playValues := strings.Split(line, " ")

			totalScore += scoreForShape(playValues[1]) + scoreForOutcome(mapEnemyShape(playValues[0]), playValues[1])
		}

		fmt.Printf("Part 1 - Total score: %d\n", totalScore)

		err := scanner.Err()
		util.Boom(err)
	}
}

func Problem2() func() {
	return func() {
		file := util.ReadFile("./day2/input2.txt")
		defer util.CloseFile()(file)

		totalScore := 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			playValues := strings.Split(line, " ")

			totalScore += scoreWithStrategy(mapEnemyShape(playValues[0]), playValues[1])
		}

		fmt.Printf("Part 2 - Total score with strategy: %d\n", totalScore)

		err := scanner.Err()
		util.Boom(err)
	}
}

func scoreForShape(shape string) int {
	switch shape {
	// Rock
	case "X":
		return 1
	// Paper
	case "Y":
		return 2
	// Scissors ("Z")
	default:
		return 3
	}
}

func mapEnemyShape(shape string) string {
	switch shape {
	// Rock
	case "A":
		return "X"
	// Paper
	case "B":
		return "Y"
	// Scissors ("C")
	default:
		return "Z"
	}
}

func scoreForOutcome(enemyShape, yourShape string) int {
	// Check if draw
	if enemyShape == yourShape {
		return 3
	}

	// Check if lost
	if (enemyShape == "Y" && yourShape == "X") ||
		(enemyShape == "Z" && yourShape == "Y") ||
		(enemyShape == "X" && yourShape == "Z") {
		return 0
	}

	// You won
	return 6
}

func scoreWithStrategy(enemyShape, yourShape string) int {
	// You need to lose
	if yourShape == "X" {
		switch enemyShape {
		case "X":
			return scoreForShape("Z") + 0
		case "Y":
			return scoreForShape("X") + 0
		// "Z"
		default:
			return scoreForShape("Y") + 0
		}
	}

	// Ends in draw - your shape is the same as your enemy
	if yourShape == "Y" {
		return scoreForShape(enemyShape) + 3
	}

	// You are going to win ("Z")
	switch enemyShape {
	case "X":
		return scoreForShape("Y") + 6
	case "Y":
		return scoreForShape("Z") + 6
	// "Z"
	default:
		return scoreForShape("X") + 6
	}
}
