package main

import (
	"fmt"
	"github.com/mariamihai/advent-of-code/util"
	"golang.org/x/exp/slices"
	"os"
)

var validYears []string

func init() {
	validYears = []string{"2021", "2022", "2023"}
}

func main() {
	year := validYear()

	fmt.Printf("Selected year %s\n", year)
}

func validYear() (year string) {
	if len(os.Args) < 2 {
		util.Boom(fmt.Errorf("year not provided"))
	}

	year = os.Args[1]

	if !slices.Contains(validYears, year) {
		util.Boom(fmt.Errorf("year [  %s  ] not supported", year))
	}

	return
}
