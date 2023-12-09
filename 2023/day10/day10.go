package day10

import (
	"bufio"
	"fmt"

	"github.com/mariamihai/advent-of-code/util"
)

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	return 0
}
