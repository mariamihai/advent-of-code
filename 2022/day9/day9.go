package day9

import (
	"advent-of-code-2022/util"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"math"
	"strings"
)

type position struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

func Problem1() func() {
	return func() {
		file := util.ReadFile("./day9/input1.txt")
		defer util.CloseFile()(file)

		// starting position
		//s := position{X: 0, Y: 0}
		head := position{X: 0, Y: 0}
		tail := position{X: 0, Y: 0}

		tailPositions := []position{tail}

		scanner := bufio.NewScanner(file)
		scanner.Split(util.CustomSplit(createCustomData()))

		for scanner.Scan() {
			line := scanner.Bytes()

			var linePosition position
			err := json.Unmarshal(line, &linePosition)
			util.Boom(err)

			head, tail, positionsFoundSoFar := move(head, tail, linePosition)
			tailPositions = append(tailPositions, positionsFoundSoFar...)

			fmt.Printf("Head %v ; Tail %v\n", head, tail)
		}

		fmt.Println(tailPositions)
		fmt.Printf("Part 1 - number of unique positions: %d\n", len(lo.Uniq[position](tailPositions)))
	}
}

func createCustomData() func(data string) []byte {
	return func(data string) []byte {
		splittedLine := strings.Split(data, " ")
		number := util.StringToInt(splittedLine[1])

		stepPosition := getNewPosition(splittedLine[0], number)

		bytes, err := json.Marshal(stepPosition)
		util.Boom(err)

		return bytes
	}
}

func getNewPosition(motion string, numberOfSteps int) position {
	switch motion {
	// Right
	case "R":
		return position{X: numberOfSteps, Y: 0}
	// Left
	case "L":
		return position{X: -numberOfSteps, Y: 0}
	// Up
	case "U":
		return position{X: 0, Y: numberOfSteps}
	// Down
	default:
		return position{X: 0, Y: -numberOfSteps}
	}
}

func areKnotsAdjacent(head, tail position) bool {
	// overlapping
	if head.X == tail.X && head.Y == tail.Y {
		return true
	}

	// lazy++
	// same row
	if head.X == tail.X && int(math.Abs(float64(head.Y-tail.Y))) == 1 {
		return true
	}

	// lazy++
	// same column
	if head.Y == tail.Y && int(math.Abs(float64(head.X-tail.X))) == 1 {
		return true
	}

	if int(math.Abs(float64(head.Y-tail.Y))) == 1 &&
		int(math.Abs(float64(head.X-tail.X))) == 1 {
		return true
	}

	return false
}

func move(head, tail, newPosition position) (position, position, []position) {
	var tailPositions []position

	// Go right, x++
	if newPosition.X > 0 {
		//fmt.Printf("Right for %d steps\n", newPosition.X)

		for i := 1; i <= newPosition.X; i++ {
			head = position{head.X + 1, head.Y}

			if areKnotsAdjacent(head, tail) {
				continue
			}

			tail = moveTail(head, tail)
			tailPositions = append(tailPositions, tail)

			//fmt.Println("Moving the tail ", head, tail)
		}
	}

	// Go left, x--
	if newPosition.X < 0 {
		//fmt.Printf("Left for %d steps\n", newPosition.X)

		for i := 1; i <= int(math.Abs(float64(newPosition.X))); i++ {
			head = position{head.X - 1, head.Y}

			if areKnotsAdjacent(head, tail) {
				continue
			}

			tail = moveTail(head, tail)
			tailPositions = append(tailPositions, tail)

			//fmt.Println("Moving the tail ", head, tail)
		}
	}

	// Go up, y++
	if newPosition.Y > 0 {
		//fmt.Printf("Up for %d steps\n", newPosition.Y)

		for i := 1; i <= newPosition.Y; i++ {
			head = position{head.X, head.Y + 1}

			if areKnotsAdjacent(head, tail) {
				continue
			}

			tail = moveTail(head, tail)
			tailPositions = append(tailPositions, tail)

			//fmt.Println("Moving the tail ", head, tail)
		}
	}

	// Go down, y--
	if newPosition.Y < 0 {
		//fmt.Printf("Down for %d steps\n", newPosition.Y)

		for i := 1; i <= int(math.Abs(float64(newPosition.Y))); i++ {
			head = position{head.X, head.Y - 1}

			if areKnotsAdjacent(head, tail) {
				continue
			}

			tail = moveTail(head, tail)
			tailPositions = append(tailPositions, tail)

			//fmt.Println("Moving the tail ", head, tail)
		}
	}

	return head, tail, tailPositions
}

// moveTail move tail with the different between the Xs or Yx is maximum 2
func moveTail(head, tail position) position {
	upOne := func(pos position) position { return position{X: pos.X, Y: tail.Y + 1} }
	downOne := func(pos position) position { return position{X: pos.X, Y: tail.Y - 1} }
	rightOne := func(pos position) position { return position{X: pos.X + 1, Y: tail.Y} }
	leftOne := func(pos position) position { return position{X: pos.X - 1, Y: tail.Y} }

	if head.X > tail.X {
		tail = rightOne(tail)
	} else if head.X < tail.X {
		tail = leftOne(tail)
	}

	if head.Y > tail.Y {
		tail = upOne(tail)
	} else if head.Y < tail.Y {
		tail = downOne(tail)
	}

	return tail
}
