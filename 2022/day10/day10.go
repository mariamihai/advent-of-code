package day10

import (
	"advent-of-code-2022/util"
	"bufio"
	"fmt"
	"strings"
)

var initialRecordingCycleStep = 20
var step = 40
var lastRecordedStep = 220

// Problem1FirstTry - part 1 without saving each step's registry value
func Problem1FirstTry() {
	file := util.ReadFile("./day10/input3.txt")
	defer util.CloseFile()(file)

	// Starting with 0 cycles but with 1 in the registry
	currentCycleStep := 0
	registryValue := 1

	// Total value for result
	signalStrength := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if ok := isNoopInstruction(line); ok {
			currentCycleStep++

			if recordingCycleStep, ok := isRecordingCycleForNoop(currentCycleStep); ok {
				signalStrength += registryValue * recordingCycleStep
			}
		}

		if ok := isAddXInstruction(line); ok {
			currentCycleStep += 2

			if recordingCycleStep, ok := isRecordingCycleForAddX(currentCycleStep); ok {
				signalStrength += registryValue * recordingCycleStep
			}

			registryValue += addToRegistry(line)
		}
	}

	fmt.Printf("Part 1 - total sum of the six signal strengths (first try): %d\n", signalStrength)
}

// Problem - save each step's registry value; the solutions for both parts
func Problem() {
	file := util.ReadFile("./day10/input3.txt")
	defer util.CloseFile()(file)

	// Starting with 1 in the registry
	registryValue := 1

	// The registry for each cycle
	var registrySteps []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if ok := isNoopInstruction(line); ok {
			registrySteps = append(registrySteps, registryValue)
		}

		if ok := isAddXInstruction(line); ok {
			// Add the registryValue for each of the 2 cycle steps
			registrySteps = append(registrySteps, registryValue, registryValue)

			registryValue += addToRegistry(line)
		}
	}

	fmt.Printf("Part 1 - total sum of the six signal strengths (second way): %d\n", sumSignalStrengths(registrySteps))

	fmt.Print("\nPart 2: ")
	drawing(registrySteps)
}

func isNoopInstruction(line string) bool {
	return line == "noop"
}

func isAddXInstruction(line string) bool {
	return strings.Contains(line, "addx")
}

func addToRegistry(line string) int {
	valueAsString := strings.Split(line, " ")[1]

	return util.StringToInt(valueAsString)
}

func isRecordingCycleForNoop(currentCycleStep int) (int, bool) {
	for i := initialRecordingCycleStep; i <= currentCycleStep && i <= lastRecordedStep; i += step {
		if currentCycleStep == i {
			return i, true
		}
	}

	return 0, false
}

func isRecordingCycleForAddX(currentCycleStep int) (int, bool) {
	for i := initialRecordingCycleStep; i <= currentCycleStep && i <= lastRecordedStep; i += step {
		// Could have an addx instruction (2 cycles instruction) on step 19 or 59, etc
		// currentCycleStep is the record for the second cycle of the instruction
		if currentCycleStep-1 == i || currentCycleStep == i {
			return i, true
		}
	}

	return 0, false
}

func sumSignalStrengths(mapCycle []int) int {
	sum := 0

	for i := initialRecordingCycleStep; i < len(mapCycle); i += step {
		// starts with 0 => mapCycle[19] is actually step 20
		sum += mapCycle[i-1] * i
	}

	return sum
}

func drawing(mapCycle []int) {
	litPixel := "#"
	darkPixel := "."

	for index, cycleValue := range mapCycle {
		if isNewLine(index) {
			fmt.Println()
		}

		if isSpriteBeingDrawn(index, cycleValue) {
			fmt.Print(litPixel)
		} else {
			fmt.Print(darkPixel)
		}
	}
}

func isNewLine(index int) bool {
	return index%40 == 0
}

// isSpriteBeingDrawn - If the sprite is positioned such that one of its three pixels is the pixel currently being drawn
func isSpriteBeingDrawn(index, cycleValue int) bool {
	spriteStep := index % 40

	return spriteStep == cycleValue || spriteStep-1 == cycleValue || spriteStep+1 == cycleValue
}
