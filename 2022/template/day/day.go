package day

import (
	"advent-of-code-2022/util"
	"bufio"
	"fmt"
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
