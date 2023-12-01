package day5

import (
	"bufio"
	"fmt"
	"github.com/mariamihai/advent-of-code/util"
	"strings"
)

type command struct {
	amount    int // number of crates to move
	startLine int // the stack removing them from
	endLine   int // the stack where they end up in
}

func Problem1() func() {
	return func() {
		stackCrates, commands := processInputData()

		stacksAfterMove := moveCrates(stackCrates, commands)

		fmt.Println(getFinalResult(stacksAfterMove))
	}
}

func Problem2() func() {
	return func() {
		stackCrates, commands := processInputData()

		stacksAfterMove := moveCrates2(stackCrates, commands)

		fmt.Println(getFinalResult(stacksAfterMove))
	}
}

func processInputData() (map[int][]string, []command) {
	file := util.ReadFile("./day5/input2.txt")
	defer util.CloseFile()(file)

	// Check if the input is commands or the table containing the crates
	readingCommands := false
	tableLines := make(map[int][]string)
	countedLines := 0

	var commands []command
	//var tableLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Read the commands
		if readingCommands {
			commands = append(commands, getCommandFromLine(line))
			continue
		}

		if isSplitBetweenTableAndCommands(line) {
			// Start reading commands from now on
			readingCommands = true
			continue
		}

		// Read the first part of the input - the table containing the stack of crates
		if !readingCommands {
			// Draw each input line - get a slice with all characters from each line
			tableLines[countedLines] = readCratesLine(line)
			countedLines++
		}
	}

	stackCrates := drawEachStackFromLine(tableLines, getNumberOfStacks(tableLines))

	return stackCrates, commands
}

// isSplitBetweenTableAndCommands There is an empty line between the part of reading the stacks and the commands
func isSplitBetweenTableAndCommands(line string) bool {
	return line == ""
}

// getCommandFromLine Splitting commands, the lines like "move 3 from 5 to 2"
func getCommandFromLine(line string) command {
	withoutMove := strings.Split(line, "move ")[1]
	splitOnFrom := strings.Split(withoutMove, " from ")
	splitOnTo := strings.Split(splitOnFrom[1], " to ")

	return command{
		amount:    util.StringToInt(splitOnFrom[0]),
		startLine: util.StringToInt(splitOnTo[0]),
		endLine:   util.StringToInt(splitOnTo[1]),
	}
}

// readCratesLine Read the crates table line by line
func readCratesLine(line string) []string {
	// Using ____ for input1.txt out of laziness to fix the file
	// Eg. of line afterwards: .[D]
	test := strings.ReplaceAll(line, "____", ".")

	// Crates missing on a level are replaced with "."
	replaceEmptyWithDot := strings.ReplaceAll(test, "    ", ".")

	// Removing "[" and "]"
	// Eg. of line afterwards: .D
	replaceFirstParantheses := strings.ReplaceAll(replaceEmptyWithDot, "[", "")
	replaceSecondParantheses := strings.ReplaceAll(replaceFirstParantheses, "]", "")

	// Remove " " between crates
	// Eg. of line afterwards: NC.
	// This processes 1 2 3 ... line as well
	replaceLastSpaces := strings.ReplaceAll(replaceSecondParantheses, " ", "")

	// Get slice from each line
	return strings.Split(replaceLastSpaces, "")
}

// getNumberOfStacks Clean way to get the number of stacks from the last line of the table.
// This is still processed by readCratesLine and returns the numbers only.
func getNumberOfStacks(tableLines map[int][]string) int {
	// Data comes from line " 1   2   3   4   5   6   7   8   9 "
	return len(tableLines[len(tableLines)-1])
}

// drawEachStackFromLine Changing from line read to stack read (from horizontal to vertical)
func drawEachStackFromLine(cratesPerLine map[int][]string, totalStacks int) map[int][]string {
	result := make(map[int][]string)

	for i := 0; i < totalStacks; i++ {
		var stack []string

		// len(cratesPerLine)-1 => ignore last line containing the numbers of the stacks
		for j := 0; j < len(cratesPerLine)-1; j++ {
			if cratesPerLine[j][i] == "." {
				continue
			}
			stack = append(stack, cratesPerLine[j][i])
		}
		result[i] = stack
	}
	return result
}

func moveCrates(crates map[int][]string, commands []command) map[int][]string {
	for _, cmd := range commands {
		for i := 1; i <= cmd.amount; i++ {
			topCrate := crates[cmd.startLine-1][0]
			crates[cmd.startLine-1] = removeFirst(crates[cmd.startLine-1])
			crates[cmd.endLine-1] = addFirst(crates[cmd.endLine-1], topCrate)
		}
	}
	return crates
}

func moveCrates2(crates map[int][]string, commands []command) map[int][]string {
	for _, cmd := range commands {
		topCrates := crates[cmd.startLine-1][:cmd.amount]
		copyTopCrates := append([]string{}, topCrates...)

		crates[cmd.startLine-1] = crates[cmd.startLine-1][cmd.amount:]
		crates[cmd.endLine-1] = append(copyTopCrates, crates[cmd.endLine-1]...)
	}
	return crates
}

func removeFirst(slice []string) []string {
	return slice[1:]
}
func addFirst(slice []string, crate string) []string {
	return append([]string{crate}, slice...)
}

func getFinalResult(crates map[int][]string) string {
	var result string
	for i := 0; i < len(crates); i++ {
		result += crates[i][0]
	}
	return result
}
