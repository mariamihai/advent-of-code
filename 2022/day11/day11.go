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
	startingItems  []int64
	operation      func(old int64) int64
	divisibleBy    int64
	ifTrue         int
	ifFalse        int
	inspectedItems int
}

func Problem1() func() {
	return func() {
		file := util.ReadFile("./day11/input1.txt")
		defer util.CloseFile()(file)

		lines := readAllLines(file)
		monkeys := mapToMonkeys(lines)

		inspectPart1(monkeys)
	}
}

func Problem2() func() {
	return func() {
		file := util.ReadFile("./day11/input1.txt")
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
	monkeys := make(map[int]*monkeyData)

	for i := 0; i < len(lines); i += 7 {
		monkey := &monkeyData{
			nr:            findMonkeyNumber(lines[i]),
			startingItems: findStartingItems(lines[i+1]),
			operation:     findOperation(lines[i+2]),
			divisibleBy:   findDivisibleBy(lines[i+3]),
			ifTrue:        findIfTrue(lines[i+4]),
			ifFalse:       findIfFalse(lines[i+5]),
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

func findStartingItems(line string) []int64 {
	items := strings.Split(line, "  Starting items: ")[1]

	return lo.Map[string, int64](strings.Split(items, ", "), func(x string, _ int) int64 {
		return util.StringToInt64(x)
	})
}

func findOperation(line string) func(old int64) int64 {
	op := strings.Split(strings.Split(line, "  Operation: new = old ")[1], " ")

	withOld := false
	var value int64
	if op[1] == "old" {
		withOld = true
	} else {
		value = util.StringToInt64(op[1])
	}

	switch op[0] {
	case "+":
		return func(old int64) int64 {
			if withOld {
				return old + old
			}
			return old + value
		}
	case "-":
		return func(old int64) int64 {
			if withOld {
				return old - old
			}
			return old - value
		}
	case "*":
		return func(old int64) int64 {
			if withOld {
				return old * old
			}
			return old * value
		}
	// "/"
	default:
		return func(old int64) int64 {
			if withOld {
				return old / old
			}
			return old / value
		}
	}
}

func findDivisibleBy(line string) int64 {
	return util.StringToInt64(strings.Split(line, "  Test: divisible by ")[1])
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

			monkeys[i].startingItems = []int64{}
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

			monkeys[i].startingItems = []int64{}
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

func isDivisible(item, testDivisible int64) bool {
	return int(item%testDivisible) == 0
}

// findWorryLevel - find another way to keep your worry levels manageable (for part 2)
func findWorryLevel(monkeys map[int]*monkeyData) int64 {
	var dividingByForAll []int64
	for i := 0; i < len(monkeys); i++ {
		dividingByForAll = append(dividingByForAll, monkeys[i].divisibleBy)
	}

	modForAll := lo.Reduce[int64, int64](dividingByForAll, func(agg int64, item int64, _ int) int64 {
		return agg * item
	}, 1)

	return modForAll
}
