package main

import (
	"fmt"
	"github.com/mariamihai/advent-of-code/2023/day1"
	"time"
)

func main() {
	fmt.Println("\n------- Day 1: Trebuchet?! -------")
	fmt.Println("Day 1, problem 1: ", day1.Problem1("./day1/input1.txt"))
	fmt.Println("Day 1, problem 2: ", day1.Problem2("./day1/input2.txt"))

	startTime := time.Now()
	day1.Problem1("./day1/input1.txt")
	fmt.Println("Problem 1 [linear solution]: ", time.Since(startTime))

	startTime = time.Now()
	day1.Problem1Concurrent("./day1/input1.txt")
	fmt.Println("Problem 1 [concurrent solution]: ", time.Since(startTime))

	startTime = time.Now()
	day1.Problem1Concurrent2("./day1/input1.txt")
	fmt.Println("Problem 1 [concurrent solution with semaphore]: ", time.Since(startTime))

}
