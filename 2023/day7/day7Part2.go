package day7

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/samber/lo"
	"strings"
)

func identifyTypeForPart2(handAsString string) int {
	countedJs := strings.Count(handAsString, "J")

	hand := strings.Split(handAsString, "")
	uniq := lo.Uniq(hand)
	length := len(uniq)

	if length == 1 {
		// Only the JJJJJ hand could include J
		return FiveOfAKind
	}

	if length == 2 {
		count := lo.Count(hand, hand[0])

		if count == 1 || count == 4 {
			// AAAAJ OR JJJJA
			if countedJs == 1 || countedJs == 4 {
				return FiveOfAKind
			}

			return FourOfAKind
		}

		// Full House with J => AAAJJJ => Five of a kind
		if countedJs != 0 {
			return FiveOfAKind
		}

		return FullHouse
	}

	if length == 3 {
		count0 := lo.Count(hand, uniq[0])
		count1 := lo.Count(hand, uniq[1])
		count2 := lo.Count(hand, uniq[2])

		if count0 == 3 || count1 == 3 || count2 == 3 {
			// JJJA3 or AAAJ1
			if countedJs == 3 || countedJs == 1 {
				return FourOfAKind
			}
			// countedJ2 == 2 => fullhouse if => will be 5 of a kind

			// countedJs == 0
			return ThreeOfAKind
		}

		// AA88J
		if countedJs == 1 {
			return FullHouse
		}

		// AAJJ1
		if countedJs == 2 {
			return FourOfAKind
		}

		// AA881, no Js
		return TwoPairs
	}

	if length == 4 {
		// JJ123
		if countedJs == 2 {
			return ThreeOfAKind
		}

		// AAJ12
		if countedJs == 1 {
			return ThreeOfAKind
		}

		// No Js
		return OnePair
	}

	//if length == 5

	// J1234
	if countedJs == 1 {
		return OnePair
	}

	return HighCard
}

func Problem2(filename string) int {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	allLines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		allLines = append(allLines, line)
	}

	allCards := make(map[int][]string)

	for _, line := range allLines {
		cardType := identifyTypeForPart2(strings.Split(line, " ")[0])

		//fmt.Println(strings.Split(line, " ")[0], "   ", types[cardType])

		if allCards[cardType] == nil {
			allCards[cardType] = []string{line}
			continue
		}
		allCards[cardType] = append(allCards[cardType], line)
	}

	var result int
	cnt := 1

	for i := 0; i <= 6; i++ {
		allCards[i] = sortCards(allCards[i], sortOrderPart2)

		for j := 0; j < len(allCards[i]); j++ {
			bid := util.StringToInt(strings.Split(allCards[i][j], " ")[1])
			result += cnt * bid
			cnt++
		}
	}

	return result
}
