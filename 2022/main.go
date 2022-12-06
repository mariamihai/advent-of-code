package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"fmt"
)

func main() {
	fmt.Println("\n--- Day 1: Calorie Counting ---")
	day1.Problem1()
	day1.Problem2()

	fmt.Println("\n--- Day 2: Rock Paper Scissors ---")
	day2.Problem1()
	day2.Problem2()

	fmt.Println("\n--- Day 3: Rucksack Reorganization ---")
	day3.Problem1()
	day3.Problem2()

	fmt.Println("\n--- Day 4: Camp Cleanup ---")
	day4.Problem()

	fmt.Println("\n--- Day 5: Supply Stacks ---")
	day5.Problem1()
	day5.Problem2()

	fmt.Println("\n--- Day 6: Tuning Trouble ---")
	fmt.Println("Part 1:")
	day6.Problem(4)
	fmt.Println("Part 2:")
	day6.Problem(14)
}
