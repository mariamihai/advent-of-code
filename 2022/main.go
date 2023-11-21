package main

import (
	"fmt"
	"github.com/mariamihai/advent-of-code/2022/day1"
	"github.com/mariamihai/advent-of-code/2022/day10"
	"github.com/mariamihai/advent-of-code/2022/day11"
	"github.com/mariamihai/advent-of-code/2022/day2"
	"github.com/mariamihai/advent-of-code/2022/day3"
	"github.com/mariamihai/advent-of-code/2022/day4"
	"github.com/mariamihai/advent-of-code/2022/day5"
	"github.com/mariamihai/advent-of-code/2022/day6"
	"github.com/mariamihai/advent-of-code/2022/day9"
	"github.com/mariamihai/advent-of-code/util"
	"os"
	"regexp"
	"strings"
)

type dayInfo struct {
	description string
	problems    []func()
}

var challenges map[string]dayInfo

func init() {
	challenges = map[string]dayInfo{
		"day1": {
			description: "\n--- Day 1: Calorie Counting ---",
			problems:    []func(){day1.Problem1(), day1.Problem2()},
		},
		"day2": {
			description: "\n--- Day 2: Rock Paper Scissors ---",
			problems:    []func(){day2.Problem1(), day2.Problem2()},
		},
		"day3": {
			description: "\n--- Day 3: Rucksack Reorganization ---",
			problems:    []func(){day3.Problem1(), day3.Problem2()},
		},
		"day4": {
			description: "\n--- Day 4: Camp Cleanup ---",
			problems:    []func(){day4.Problem()},
		},
		"day5": {
			description: "\n--- Day 5: Supply Stacks ---",
			problems:    []func(){day5.Problem1(), day5.Problem2()},
		},
		"day6": {
			description: "\n--- Day 6: Tuning Trouble ---",
			problems:    []func(){day6.Problem(4), day6.Problem(14)},
		},
		"day7": {
			description: "\n--- Day 7:  ---",
			problems:    []func(){},
		},
		"day8": {
			description: "\n--- Day 8:  ---",
			problems:    []func(){},
		},
		"day9": {
			description: "\n--- Day 9: Rope Bridge ---",
			problems:    []func(){day9.Problem1()},
		},
		"day10": {
			description: "\n--- Day 10: Cathode-Ray Tube ---",
			problems: []func(){
				//day10.Problem1FirstTry(),
				day10.Problem()},
		},
		"day11": {
			description: "\n--- Day 11: Monkey in the Middle ---",
			problems:    []func(){day11.Problem1(), day11.Problem2()},
		},
	}
}

func main() {
	validateDay()
	day := os.Args[1]

	fmt.Println(challenges[day].description)

	if noProblemSelected() {
		// Print all problems for the day
		for _, problem := range challenges[day].problems {
			problem()
		}
	} else {
		problemNumber := getProblemNumber()
		validateProblemNumber(problemNumber)
		challenges[day].problems[problemNumber-1]()
	}
}

func validateDay() {
	_, ok := challenges[os.Args[1]]
	if !strings.HasPrefix(os.Args[1], "day") || !ok {
		util.Boom(fmt.Errorf("invalid day: %s", os.Args[1]))
	}
}

func validateProblemNumber(problemNumber int) {
	day := challenges[os.Args[1]]

	if len(day.problems) < problemNumber {
		util.Boom(fmt.Errorf("invalid problem: %s", os.Args[2]))
	}
}

func noProblemSelected() bool {
	return len(os.Args) == 2
}

func getProblemNumber() int {
	re := regexp.MustCompile("problem(.*)")
	match := re.FindStringSubmatch(os.Args[2])

	return util.StringToInt(match[1])
}
