package day

import (
	"bufio"
	"fmt"

	"github.com/mariamihai/advent-of-code/util"
)

func Problem1() {
	file := util.ReadFile("./day/input1.txt")
	defer util.CloseFile()(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}
