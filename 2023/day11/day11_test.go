package day11

import "testing"

func TestProblem1(t *testing.T) {
	tests := []struct {
		description string
		filename    string
		expected    int
	}{
		{
			description: "Should validate PROBLEM 1 against provided EXAMPLE",
			filename:    "example1.txt",
			expected:    374,
		},
		{
			description: "Should validate PROBLEM 1 against provided INPUT",
			filename:    "input1.txt",
			expected:    9403026,
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
		description         string
		filename            string
		expansionMultiplier int
		expected            int
	}{
		{
			description:         "Should validate PROBLEM 2 against provided EXAMPLE with expansion multiplier 2 (problem 1)",
			filename:            "example1.txt",
			expansionMultiplier: 2,
			expected:            374,
		},
		{
			description:         "Should validate PROBLEM 2 against provided INPUT with expansion multiplier 2 (problem 1)",
			filename:            "input1.txt",
			expansionMultiplier: 2,
			expected:            9403026,
		},
		{
			description:         "Should validate PROBLEM 2 against provided EXAMPLE with expansion multiplier 10",
			filename:            "example1.txt",
			expansionMultiplier: 10,
			expected:            1030,
		},
		{
			description:         "Should validate PROBLEM 2 against provided EXAMPLE with expansion multiplier 100",
			filename:            "example1.txt",
			expansionMultiplier: 100,
			expected:            8410,
		},
		{
			description:         "Should validate PROBLEM 2 against provided INPUT with expansion multiplier 1.000.000",
			filename:            "input1.txt",
			expansionMultiplier: 1000000,
			expected:            543018317006,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if actual := Problem2(tt.filename, tt.expansionMultiplier); actual != tt.expected {
				t.Errorf("Problem1() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
