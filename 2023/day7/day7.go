package day7

import (
	"bufio"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/samber/lo"
	"sort"
	"strings"
	"unicode"
)

const (
	FiveOfAKind  = 6
	FourOfAKind  = 5
	FullHouse    = 4
	ThreeOfAKind = 3
	TwoPairs     = 2
	OnePair      = 1
	HighCard     = 0
)

func identifyType(handAsString string) int {
	hand := strings.Split(handAsString, "")
	uniq := lo.Uniq(hand)

	length := len(uniq)

	if length == 1 {
		return FiveOfAKind
	}

	if length == 2 {
		count := lo.Count(hand, hand[0])

		if count == 1 || count == 4 {
			return FourOfAKind
		}

		return FullHouse
	}

	if length == 3 {
		count0 := lo.Count(hand, uniq[0])
		count1 := lo.Count(hand, uniq[1])
		count2 := lo.Count(hand, uniq[2])

		if count0 == 3 || count1 == 3 || count2 == 3 {
			return ThreeOfAKind
		}

		return TwoPairs
	}

	if length == 4 {
		return OnePair
	}

	//if length == 5
	return HighCard
}

func Problem1(filename string) int {
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
		cardType := identifyType(strings.Split(line, " ")[0])

		if allCards[cardType] == nil {
			allCards[cardType] = []string{line}
			continue
		}
		allCards[cardType] = append(allCards[cardType], line)
	}

	var result int
	cnt := 1

	for i := 0; i <= 6; i++ {
		allCards[i] = sortCards(allCards[i])

		for j := 0; j < len(allCards[i]); j++ {
			bid := util.StringToInt(strings.Split(allCards[i][j], " ")[1])

			result += cnt * bid
			cnt++
		}
	}

	return result
}

func sortCards(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		card1 := strings.Split(lines[i], " ")[0]
		card2 := strings.Split(lines[j], " ")[0]

		for i := 0; i < 5; i++ {
			rune1 := rune(card1[i])
			rune2 := rune(card2[i])

			string1 := string(card1[i])
			string2 := string(card2[i])

			isDigit1 := unicode.IsDigit(rune1)
			isDigit2 := unicode.IsDigit(rune2)

			if rune1 == rune2 {
				continue
			}

			if !isDigit1 && isDigit2 {
				return false
			}
			if isDigit1 && !isDigit2 {
				return true
			}

			if isDigit1 && isDigit2 {
				if util.StringToInt(string1) == util.StringToInt(string2) {
					continue
				}
				return util.StringToInt(string1) < util.StringToInt(string2)
			}

			if !isDigit1 && !isDigit2 {
				if rune1 == rune2 {
					continue
				}

				// A, K, Q, J, T
				// Compare A with anything else
				if rune1 == 'A' && rune2 != 'A' {
					return false
				}
				if rune1 != 'A' && rune2 == 'A' {
					return true
				}

				// Compare T with anything else
				if rune1 == 'T' && rune2 != 'T' {
					return true
				}
				if rune1 != 'T' && rune2 == 'T' {
					return false
				}

				// Compare K with anything else
				if rune1 == 'K' {
					return false
				}
				if rune2 == 'K' {
					return true
				}

				// Compare Q with anything else
				if rune1 == 'Q' {
					return false
				}
				if rune2 == 'Q' {
					return true
				}

				// Remaining J against T, already done
			}
		}
		return false // equal
	})

	return lines
}
