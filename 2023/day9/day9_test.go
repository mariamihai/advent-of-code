package day9

import "testing"

func TestProblem1(t *testing.T) {
	tests := []struct {
		description string
		filename    string
		expected    int
	}{
		{
			description: "Should validate PROBLEM 1 against provided EXAMPLE 1",
			filename:    "example1.txt",
			expected:    114,
		},
		{
			description: "Should validate PROBLEM 1 against provided INPUT",
			filename:    "input1.txt",
			expected:    1641934234,
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
			description: "Should validate PROBLEM 2 against provided EXAMPLE 2",
			filename:    "example2.txt",
			expected:    5,
		},
		{
			description: "Should validate PROBLEM 2 against provided INPUT",
			filename:    "input1.txt",
			expected:    975,
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
