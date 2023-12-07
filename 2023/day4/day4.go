package day4

import (
	"bufio"
	"github.com/samber/lo"
	"math"
	"strconv"
	"strings"

	"github.com/mariamihai/advent-of-code/util"
)

func Problem1(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var sum float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		cardValue, _ := processCard(line)

		sum += cardValue
	}

	return int(sum)
}

type Card struct {
	cardNumber     int
	cardValue      int
	nrOfInstances  int
	winningNumbers []int
}

func Problem2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	var sum int

	var cnt int
	cardResults := make(map[int]Card)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		cardValue, winningNumbers := processCard(line)
		cardResults[cnt] = Card{
			cardNumber:     cnt,
			cardValue:      int(cardValue),
			nrOfInstances:  1,
			winningNumbers: winningNumbers,
		}
		cnt++
	}

	for i := 0; i < len(cardResults); i++ {
		card := cardResults[i]

		for j := 1; j <= len(card.winningNumbers); j++ {

			cardResults[card.cardNumber+j] = Card{
				cardNumber:     cardResults[card.cardNumber+j].cardNumber,
				cardValue:      cardResults[card.cardNumber+j].cardValue,
				nrOfInstances:  card.nrOfInstances + cardResults[card.cardNumber+j].nrOfInstances,
				winningNumbers: cardResults[card.cardNumber+j].winningNumbers,
			}
		}
	}

	for i := 0; i <= len(cardResults); i++ {
		sum += cardResults[i].nrOfInstances
	}

	return sum
}

func processCard(line string) (float64, []int) {
	split := strings.Split(line, ":")
	line = split[1]

	split = strings.Split(line, "|")

	winningNumbers := util.NumbersAsStringSlice(split[0])
	ownNumbers := util.NumbersAsStringSlice(split[1])

	intersect := lo.Intersect(winningNumbers, ownNumbers)

	intersectAsInts := lo.Map(intersect, func(x string, index int) int {
		result, err := strconv.Atoi(x)
		if err != nil {
			return 0
		}
		return result
	})

	if len(intersect) == 0 {
		return 0, []int{}
	}

	return math.Pow(2, (float64)(len(intersect)-1)), intersectAsInts
}

//fmt.Println("\n------- Day 4: Scratchcards -------")
//fmt.Println("Day 4, problem 1: ", day4.Problem1("./day4/input1.txt"))
//fmt.Println("Day 4, problem 2: ", day4.Problem2("./day4/input1.txt"))
