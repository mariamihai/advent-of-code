package day2

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"strings"
)

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var horizontalPosition = 0
	var verticalPosition = 0 // depth

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		number := util.StringToInt(line[1])
		if line[0] == "forward" {
			horizontalPosition += number
		}
		if line[0] == "down" {
			verticalPosition += number
		}
		if line[0] == "up" {
			verticalPosition -= number
		}
	}

	util.Boom(scanner.Err())

	return horizontalPosition * verticalPosition
}

func Problem2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var horizontalPosition = 0
	var verticalPosition = 0 // depth
	var aim = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		number := util.StringToInt(line[1])
		if line[0] == "forward" {
			horizontalPosition += number
			verticalPosition += aim * number
		}
		if line[0] == "down" {
			aim += number
		}
		if line[0] == "up" {
			aim -= number
		}
	}
	util.Boom(scanner.Err())

	return horizontalPosition * verticalPosition
}
