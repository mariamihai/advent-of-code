package day9

import "testing"

func TestAreKnotsAdjacent(t *testing.T) {
	testCases := []struct {
		description string
		head        position
		tail        position
		result      bool
	}{
		{
			description: "Test overlapping",
			head:        position{X: 1, Y: 1},
			tail:        position{X: 1, Y: 1},
			result:      true,
		},
		{
			description: "Test same row and adjacent",
			head:        position{X: 1, Y: 1},
			tail:        position{X: 1, Y: 2},
			result:      true,
		},
		{
			description: "Test same row but NOT adjacent",
			head:        position{X: 1, Y: 0},
			tail:        position{X: 1, Y: 2},
			result:      false,
		},
		{
			description: "Test same column and adjacent",
			head:        position{X: 1, Y: 1},
			tail:        position{X: 2, Y: 1},
			result:      true,
		},
		{
			description: "Test same column but NOT adjacent",
			head:        position{X: 0, Y: 1},
			tail:        position{X: 2, Y: 1},
			result:      false,
		},
		//.....
		//..H..
		//.T...
		//.....
		{
			description: "Test diagonal positions and adjacent (test 1)",
			head:        position{X: 0, Y: 0},
			tail:        position{X: -1, Y: -1},
			result:      true,
		},
		//.....
		//..T..
		//.H...
		//.....
		{
			description: "Test diagonal positions and adjacent (test 2)",
			head:        position{X: 0, Y: 0},
			tail:        position{X: 1, Y: 1},
			result:      true,
		},
		//.....
		//..H..
		//...T.
		//.....
		{
			description: "Test diagonal positions and adjacent (test 3)",
			head:        position{X: 0, Y: 0},
			tail:        position{X: 1, Y: -1},
			result:      true,
		},
		//.....
		//..T..
		//...H.
		//.....
		{
			description: "Test diagonal positions and adjacent (test 4)",
			head:        position{X: 0, Y: 0},
			tail:        position{X: -1, Y: 1},
			result:      true,
		},
		{
			description: "Test not adjacent",
			head:        position{X: 0, Y: 2},
			tail:        position{X: 5, Y: 8},
			result:      false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			testResult := areKnotsAdjacent(testCase.head, testCase.tail)

			if testResult != testCase.result {
				t.Errorf("%s: expected %t, got %t", testCase.description, testCase.result, testResult)
			}
		})
	}
}

func TestMoveTail(t *testing.T) {
	testCases := []struct {
		description string
		head        position
		tail        position
		result      position
	}{
		{
			description: "Move on the same row to right",
			head:        position{X: 2, Y: 0},
			tail:        position{X: 0, Y: 0},
			result:      position{X: 1, Y: 0},
		},
		{
			description: "Move on the same row to left",
			head:        position{X: 0, Y: 0},
			tail:        position{X: 2, Y: 0},
			result:      position{X: 1, Y: 0},
		},
		{
			description: "Move on the same column up",
			head:        position{X: 0, Y: 2},
			tail:        position{X: 0, Y: 0},
			result:      position{X: 0, Y: 1},
		},
		{
			description: "Move on the same column down",
			head:        position{X: 0, Y: 0},
			tail:        position{X: 0, Y: 2},
			result:      position{X: 0, Y: 1},
		},
		{
			description: "Move up and right",
			head:        position{X: 1, Y: 2},
			tail:        position{X: 0, Y: 0},
			result:      position{X: 1, Y: 1},
		},
		{
			description: "Move up and left",
			head:        position{X: -1, Y: 2},
			tail:        position{X: 0, Y: 0},
			result:      position{X: -1, Y: 1},
		},
		{
			description: "Move down and right",
			head:        position{X: 1, Y: 2},
			tail:        position{X: 0, Y: 0},
			result:      position{X: 0, Y: 1},
		},
		{
			description: "Move down and left",
			head:        position{X: -1, Y: 2},
			tail:        position{X: 0, Y: 0},
			result:      position{X: 0, Y: 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			testResult := moveTail(testCase.head, testCase.tail)

			if testCase.result.X != testResult.X ||
				testCase.result.Y != testResult.Y {
				t.Errorf("%s: expected %v, got %v", testCase.description, testCase.result, testResult)
			}
		})
	}
}
