package day8

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mariamihai/advent-of-code/util"
	"github.com/samber/lo"
	"strings"
)

type rowType []int

func Problem1() {
	file := util.ReadFile("./day8/input1.txt")
	defer util.CloseFile()(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(util.CustomSplit(createCustomData()))

	var row rowType
	var rows []rowType

	for scanner.Scan() {
		row = nil
		line := scanner.Bytes()

		err := json.Unmarshal(line, &row)
		util.Boom(err)

		rows = append(rows, row)
	}

	fmt.Printf("Part 1 - number of trees seen: %d\n", sumEdges(rows)+visibleTreesInsideTheEdges(rows))
	//fmt.Printf("Part 1 - number of trees seen: %d\n", sumEdges(rows)+visibleTreesInsideTheEdges2(rows))
}

func createCustomData() func(data string) []byte {
	return func(data string) []byte {
		mappedToInt := lo.Map[string, int](strings.Split(data, ""), func(x string, index int) int {
			return util.StringToInt(x)
		})

		bytes, err := json.Marshal(mappedToInt)
		util.Boom(err)

		return bytes
	}
}

func sumEdges(rows []rowType) int {
	return len(rows[0])*2 + (len(rows)-2)*2
}

func visibleTreesInsideTheEdges(rows []rowType) int {
	count := 0

	// Each row except first and last, from left to right:
	for i := 1; i < (len(rows) - 1); i++ {
		row := rows[i]

		maxHeightSeen := row[0]
		for j := 1; j < (len(row) - 1); j++ {
			if row[j] > maxHeightSeen {
				count++
				maxHeightSeen = row[j]
				row[j] = -1
			}
		}
	}

	// Each row except first and last, from right to left:
	for i := 1; i < (len(rows) - 1); i++ {
		row := rows[i]

		maxHeightSeen := row[len(row)-1]
		for j := len(row) - 2; j >= 1; j-- {
			if row[j] > maxHeightSeen {
				count++
				maxHeightSeen = row[j]
				row[j] = -1
			}
		}
	}

	// Each column except first and last, from top to bottom:
	for j := 1; j < (len(rows[0]) - 2); j++ {
		maxHeightSeen := rows[0][j]

		for i := 1; i < (len(rows) - 2); i++ {
			currentHeight := rows[i][j]

			if currentHeight > maxHeightSeen {
				count++
				maxHeightSeen = currentHeight
				rows[i][j] = -1
			}
		}
	}

	// Each column except first and last, from bottom to top:
	for j := 1; j < (len(rows[0]) - 2); j++ {
		maxHeightSeen := rows[len(rows)-1][j]

		for i := len(rows) - 2; i >= 1; i-- {
			currentHeight := rows[i][j]

			if currentHeight > maxHeightSeen {
				count++
				maxHeightSeen = currentHeight
				rows[i][j] = -1
			}
		}
	}

	return count
}

func visibleTreesInsideTheEdges2(rows []rowType) int {
	count := 0

	// Each row except first and last, from left to right:
	for i := 1; i < (len(rows) - 1); i++ {
		for j := 1; j < (len(rows[i]) - 1); j++ {
			// left to right
			if rows[i][j] > rows[i][j-1] {
				count++
				rows[i][j] = -1
			}
			// right to left
			if rows[i][j] > rows[i][j+1] {
				count++
				rows[i][j] = -1
			}
			// Top to bottom
			if rows[i][j] > rows[i-1][j] {
				count++
				rows[i][j] = -1
			}
			// Bottom to top
			if rows[i][j] > rows[i+1][j] {
				count++
				rows[i][j] = -1
			}
		}
	}

	return count
}
