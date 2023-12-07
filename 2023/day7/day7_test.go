package day7

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
			description: "Should validate PROBLEM 1 against provided EXAMPLE 1",
			filename:    "example1.txt",
			expected:    6440,
		},
		{
			description: "Should validate PROBLEM 1 against provided EXAMPLE 2",
			filename:    "example2.txt",
			expected:    6592,
		},
		{
			description: "Should validate PROBLEM 1 against provided INPUT",
			filename:    "input1.txt",
			expected:    253313241,
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
			description: "Should validate PROBLEM 2 against provided EXAMPLE 1",
			filename:    "example1.txt",
			expected:    5905,
		},
		{
			description: "Should validate PROBLEM 2 against provided EXAMPLE 2",
			filename:    "example2.txt",
			expected:    6839,
		},
		{
			description: "Should validate PROBLEM 2 against provided INPUT",
			filename:    "input1.txt",
			expected:    253362743,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if actual := Problem2(tt.filename); actual != tt.expected {
				t.Errorf("Problem2() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestGetTypeForHandPart1(t *testing.T) {
	tests := []struct {
		description string
		hand        string
		expected    int
	}{
		{
			description: "Should map to Five of a Kind",
			hand:        "AAAAA",
			expected:    FiveOfAKind,
		},
		{
			description: "Should map to Four of a Kind",
			hand:        "AA8AA",
			expected:    FourOfAKind,
		},

		{
			description: "Should map to Full House",
			hand:        "23332",
			expected:    FullHouse,
		},
		{
			description: "Should map to Three of a Kind",
			hand:        "TTT98",
			expected:    ThreeOfAKind,
		},
		{
			description: "Should map to Two Pairs",
			hand:        "23432",
			expected:    TwoPairs,
		},
		{
			description: "Should map to One Pair",
			hand:        "A23A4",
			expected:    OnePair,
		},
		{
			description: "Should map to High Card",
			hand:        "23456",
			expected:    HighCard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if actual := getTypeForHandPart1(tt.hand); actual != tt.expected {
				t.Errorf("getTypeForHandPart1() = %v, want %v", actual, tt.expected)
			}
		})
	}
}

func TestSortCards(t *testing.T) {
	tests := []struct {
		description string
		lines       []string
		sortOrder   string
		expected    []string
	}{
		{
			description: "Test 1",
			lines:       []string{"KK677", "KTJJT"},
			sortOrder:   sortOrderPart1,
			expected:    []string{"KTJJT", "KK677"},
		},
		{
			description: "Test 2",
			lines:       []string{"QQQJA", "T55J5"},
			sortOrder:   sortOrderPart1,
			expected:    []string{"T55J5", "QQQJA"},
		},
		{
			description: "Test 3",
			lines:       []string{"23456", "12345"},
			sortOrder:   sortOrderPart1,
			expected:    []string{"12345", "23456"},
		},
		{
			description: "Test 4",
			lines:       []string{"T55J5", "KTJJT", "QQQJA"},
			sortOrder:   sortOrderPart2,
			expected:    []string{"T55J5", "QQQJA", "KTJJT"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual := sortCards(tt.lines, tt.sortOrder)

			for i := 0; i < len(tt.expected); i++ {
				if tt.expected[i] != actual[i] {
					t.Errorf("sortCards() = %v, want %v", actual, tt.expected)
				}
			}
		})
	}
}
