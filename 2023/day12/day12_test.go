package day12

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	tests := []struct {
		description string
		filename    string
		expected    int
	}{
		{
			description: "Should validate PROBLEM 1 against provided EXAMPLE",
			filename:    "example1.txt",
			expected:    21,
		},
		{
			description: "Should validate PROBLEM 1 against provided INPUT",
			filename:    "input1.txt",
			expected:    7204,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if actual := Problem1(tt.filename); actual != tt.expected {
				t.Errorf("Problem1() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	tests := []struct {
		description string
		filename    string
		expected    int
	}{
		{
			description: "Should validate PROBLEM 2 against provided EXAMPLE",
			filename:    "example1.txt",
			expected:    525152,
		},
		//{
		//	description: "Should validate PROBLEM 2 against provided INPUT",
		//	filename:    "input1.txt",
		//	expected:    0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if actual := Problem2(tt.filename); actual != tt.expected {
				t.Errorf("Problem2() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestCalculateNrOfPossibilities(t *testing.T) {
	tests := []struct {
		description    string
		input          string
		damagedSprings []int
		expected       int
	}{
		{
			description:    "Test 1",
			input:          "???.###",
			damagedSprings: []int{1, 1, 3},
			expected:       1,
		},
		{
			description:    "Test 2",
			input:          "????",
			damagedSprings: []int{1},
			expected:       4,
		},
		{
			description:    "Test 3",
			input:          "?????",
			damagedSprings: []int{1},
			expected:       5,
		},
		{
			description:    "Test 4",
			input:          "????.###",
			damagedSprings: []int{1, 1, 3},
			expected:       3,
		},
		{
			description:    "Test 5",
			input:          "#???.###",
			damagedSprings: []int{1, 1, 3},
			expected:       2,
		},
		{
			description:    "Test 6",
			input:          ".??..??...?##.",
			damagedSprings: []int{1, 1, 3},
			expected:       4,
		},
		{
			description:    "Test 7",
			input:          "?#?#?#?#?#?#?#?",
			damagedSprings: []int{1, 3, 1, 6},
			expected:       1,
		},
		{
			description:    "Test 8",
			input:          "????.#...#...",
			damagedSprings: []int{4, 1, 1},
			expected:       1,
		},
		{
			description:    "Test 9",
			input:          "????.######..#####.",
			damagedSprings: []int{1, 6, 5},
			expected:       4,
		},
		{
			description:    "Test 10",
			input:          "#??????",
			damagedSprings: []int{2, 1},
			expected:       4,
		},
		{
			description:    "Test 11",
			input:          "?###????????",
			damagedSprings: []int{3, 2, 1},
			expected:       10,
		},
		{
			description:    "Test 12",
			input:          "?????#??..",
			damagedSprings: []int{1, 3},
			expected:       9,
		},
		{
			description:    "Test 13",
			input:          "???#.?#..?",
			damagedSprings: []int{3, 1},
			expected:       1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if actual := evaluate(tt.input, tt.damagedSprings); actual != tt.expected {
				t.Errorf("Problem1() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
