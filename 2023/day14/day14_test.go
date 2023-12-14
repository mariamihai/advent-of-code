package day14

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
			expected:    136,
		},
		{
			description: "Should validate PROBLEM 1 against provided INPUT",
			filename:    "input1.txt",
			expected:    108857,
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
