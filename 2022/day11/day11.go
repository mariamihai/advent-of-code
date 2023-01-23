package day11

import (
	util "advent-of-code-2022/util"
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"regexp"
	"sort"
	"strings"
)

type monkeyData struct {
	nr             int
	startingItems  []int
	operation      func(old int) int
	divisibleBy    int
	ifTrue         int
	ifFalse        int
	inspectedItems int
}

func Problem1() func() {
	return func() {
		file := util.ReadFile("./day11/input2.txt") // test input
		defer util.CloseFile()(file)

		lines := readAllLines(file)
		monkeys := mapToMonkeys(lines)

		inspectPart1(monkeys)
	}
}

func Problem2() func() {
	return func() {
		file := util.ReadFile("./day11/input2.txt") // test input
		defer util.CloseFile()(file)

		lines := readAllLines(file)
		monkeys := mapToMonkeys(lines)

		inspectPart2(monkeys)
	}
}

func readAllLines(file *os.File) []string {
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func mapToMonkeys(lines []string) map[int]*monkeyData {
	// 6 lines per monkey + an empty one
	numberOfLinesPerMonkey := 7
	//Monkey 0:
	monkeyLine := func(lines []string, i int) string { return lines[i] }
	//  Starting items: 79, 98
	startingItemsLine := func(lines []string, i int) string { return lines[i+1] }
	//  Operation: new = old * 19
	operationLine := func(lines []string, i int) string { return lines[i+2] }
	//  Test: divisible by 23
	divisionLine := func(lines []string, i int) string { return lines[i+3] }
	//    If true: throw to monkey 2
	trueLine := func(lines []string, i int) string { return lines[i+4] }
	//    If false: throw to monkey 3
	falseLine := func(lines []string, i int) string { return lines[i+5] }

	monkeys := make(map[int]*monkeyData)

	for i := 0; i < len(lines); i += numberOfLinesPerMonkey {
		monkey := &monkeyData{
			nr:            findMonkeyNumber(monkeyLine(lines, i)),
			startingItems: findStartingItems(startingItemsLine(lines, i)),
			operation:     findOperation(operationLine(lines, i)),
			divisibleBy:   findDivisibleBy(divisionLine(lines, i)),
			ifTrue:        findIfTrue(trueLine(lines, i)),
			ifFalse:       findIfFalse(falseLine(lines, i)),
		}

		monkeys[monkey.nr] = monkey
	}

	return monkeys
}

func findMonkeyNumber(line string) int {
	re := regexp.MustCompile("Monkey (.*):")
	match := re.FindStringSubmatch(line)

	return util.StringToInt(match[1])
}

func findStartingItems(line string) []int {
	items := strings.Split(line, "  Starting items: ")[1]

	return lo.Map[string, int](strings.Split(items, ", "), func(x string, _ int) int {
		return util.StringToInt(x)
	})
}

func findOperation(line string) func(old int) int {
	op := strings.Split(strings.Split(line, "  Operation: new = old ")[1], " ")

	withOld := false
	var value int
	if op[1] == "old" {
		withOld = true
	} else {
		value = util.StringToInt(op[1])
	}

	switch op[0] {
	case "+":
		return func(old int) int {
			if withOld {
				return old + old
			}
			return old + value
		}
	case "-":
		return func(old int) int {
			if withOld {
				return old - old
			}
			return old - value
		}
	case "*":
		return func(old int) int {
			if withOld {
				return old * old
			}
			return old * value
		}
	// "/"
	default:
		return func(old int) int {
			if withOld {
				return old / old
			}
			return old / value
		}
	}
}

func findDivisibleBy(line string) int {
	return util.StringToInt(strings.Split(line, "  Test: divisible by ")[1])
}

func findIfTrue(line string) int {
	return util.StringToInt(strings.Split(line, "    If true: throw to monkey ")[1])
}

func findIfFalse(line string) int {
	return util.StringToInt(strings.Split(line, "    If false: throw to monkey ")[1])
}

func inspectPart1(monkeys map[int]*monkeyData) {
	for round := 1; round <= 20; round++ {
		for i := 0; i < len(monkeys); i++ {
			itemsToCheck := monkeys[i].startingItems
			monkeys[i].inspectedItems += len(itemsToCheck)

			for _, item := range itemsToCheck {
				item = monkeys[i].operation(item) / 3

				if isDivisible(item, monkeys[i].divisibleBy) {
					monkeys[monkeys[i].ifTrue].startingItems = append(monkeys[monkeys[i].ifTrue].startingItems, item)
				} else {
					monkeys[monkeys[i].ifFalse].startingItems = append(monkeys[monkeys[i].ifFalse].startingItems, item)
				}
			}

			monkeys[i].startingItems = []int{}
		}
	}

	var inspectedTimes []int
	for i := 0; i < len(monkeys); i++ {
		inspectedTimes = append(inspectedTimes, monkeys[i].inspectedItems)
	}

	sort.Ints(inspectedTimes)
	length := len(inspectedTimes)
	fmt.Println("Part 1 - the monkey business of the two most active monkeys: ", inspectedTimes[length-1]*inspectedTimes[length-2])
}

func inspectPart2(monkeys map[int]*monkeyData) {
	for round := 1; round <= 10000; round++ {
		modForAll := findWorryLevel(monkeys)

		for i := 0; i < len(monkeys); i++ {
			itemsToCheck := monkeys[i].startingItems
			monkeys[i].inspectedItems += len(itemsToCheck)

			for _, item := range itemsToCheck {
				item = monkeys[i].operation(item) % modForAll

				if isDivisible(item, monkeys[i].divisibleBy) {
					monkeys[monkeys[i].ifTrue].startingItems = append(monkeys[monkeys[i].ifTrue].startingItems, item)
				} else {
					monkeys[monkeys[i].ifFalse].startingItems = append(monkeys[monkeys[i].ifFalse].startingItems, item)
				}
			}

			monkeys[i].startingItems = []int{}
		}
	}

	var inspectedTimes []int
	for i := 0; i < len(monkeys); i++ {
		inspectedTimes = append(inspectedTimes, monkeys[i].inspectedItems)
	}

	sort.Ints(inspectedTimes)
	ln := len(inspectedTimes)
	fmt.Println("Part 2 - the monkey business of the two most active monkeys: ", inspectedTimes[ln-1]*inspectedTimes[ln-2])
}

func isDivisible(item, testDivisible int) bool {
	return int(item%testDivisible) == 0
}

// findWorryLevel - find another way to keep your worry levels manageable (for part 2)
func findWorryLevel(monkeys map[int]*monkeyData) int {
	var dividingByForAll []int
	for i := 0; i < len(monkeys); i++ {
		dividingByForAll = append(dividingByForAll, monkeys[i].divisibleBy)
	}

	modForAll := lo.Reduce[int, int](dividingByForAll, func(agg int, item int, _ int) int {
		return agg * item
	}, 1)

	return modForAll
}
