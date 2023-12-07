package day7

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/samber/lo"
	"sort"
	"strings"
)

const sortOrderPart1 = "23456789TJQKA"
const sortOrderPart2 = "J23456789TQKA"

const (
	FiveOfAKind  = 6
	FourOfAKind  = 5
	FullHouse    = 4
	ThreeOfAKind = 3
	TwoPairs     = 2
	OnePair      = 1
	HighCard     = 0
)

func Problem1(filename string) int {
	allLines := getInput(filename)
	cardsByType := mapCardsByType(allLines, getTypeForHandPart1)

	var result int
	cnt := 1

	for i := 0; i <= 6; i++ {
		cardsByType[i] = sortCards(cardsByType[i], sortOrderPart1)

		for j := 0; j < len(cardsByType[i]); j++ {
			bid := util.StringToInt(strings.Split(cardsByType[i][j], " ")[1])

			result += cnt * bid
			cnt++
		}
	}

	return result
}

func Problem2(filename string) int {
	allLines := getInput(filename)
	cardsByType := mapCardsByType(allLines, getTypeForHandPart2)

	var result int
	cnt := 1

	for i := 0; i <= 6; i++ {
		cardsByType[i] = sortCards(cardsByType[i], sortOrderPart2)

		for j := 0; j < len(cardsByType[i]); j++ {
			bid := util.StringToInt(strings.Split(cardsByType[i][j], " ")[1])
			result += cnt * bid
			cnt++
		}
	}

	return result
}

func getInput(filename string) []string {
	file := util.ReadFile(filename)
	defer util.CloseFile()(file)

	allLines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		allLines = append(allLines, line)
	}

	return allLines
}

func mapCardsByType(allLines []string, getTypeFunc func(string) int) map[int][]string {
	cardsByType := make(map[int][]string)

	for _, line := range allLines {
		cardType := getTypeFunc(strings.Split(line, " ")[0])

		if cardsByType[cardType] == nil {
			cardsByType[cardType] = []string{line}
			continue
		}
		cardsByType[cardType] = append(cardsByType[cardType], line)
	}

	return cardsByType
}

func getTypeForHandPart1(handAsString string) int {
	hand := strings.Split(handAsString, "")
	uniq := lo.Uniq(hand)

	switch len(uniq) {
	case 1:
		return FiveOfAKind
	case 2:
		count := lo.Count(hand, hand[0])

		if count == 1 || count == 4 {
			return FourOfAKind
		}

		return FullHouse
	case 3:
		count0 := lo.Count(hand, uniq[0])
		count1 := lo.Count(hand, uniq[1])
		count2 := lo.Count(hand, uniq[2])

		if count0 == 3 || count1 == 3 || count2 == 3 {
			return ThreeOfAKind
		}

		return TwoPairs
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func getTypeForHandPart2(handAsString string) int {
	numberOfJokers := strings.Count(handAsString, "J")

	if numberOfJokers == 0 {
		return getTypeForHandPart1(handAsString)
	}

	hand := strings.Split(handAsString, "")
	uniq := lo.Uniq(hand)

	switch len(uniq) {
	case 1, 2:
		// JJJJJ, AAAAJ, JJJJA, AAAJJJ
		return FiveOfAKind
	case 3:
		// JJJ12
		if numberOfJokers == 3 {
			return FourOfAKind
		}

		// AAAJ1
		if lo.Count(hand, uniq[0]) == 3 ||
			lo.Count(hand, uniq[1]) == 3 ||
			lo.Count(hand, uniq[2]) == 3 {
			return FourOfAKind
		}

		// AAJJ1
		if numberOfJokers == 2 {
			return FourOfAKind
		}

		// AA88J
		if numberOfJokers == 1 {
			return FullHouse
		}
	case 4:
		// JJ123, AAJ12
		return ThreeOfAKind
	default:
		// length == 5; J1234
		return OnePair
	}

	return HighCard
}

func sortCards(lines []string, sortOrder string) []string {
	sort.Slice(lines, func(i, j int) bool {
		card1 := strings.Split(lines[i], " ")[0]
		card2 := strings.Split(lines[j], " ")[0]

		for i := 0; i < 5; i++ {
			card1 := strings.IndexByte(sortOrder, card1[i])
			card2 := strings.IndexByte(sortOrder, card2[i])

			if card1 == card2 {
				continue
			}

			return card1 < card2
		}
		return true
	})

	return lines
}
