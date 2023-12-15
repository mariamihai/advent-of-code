package day15

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
			expected:    1320,
		},
		{
			description: "Should validate PROBLEM 1 against provided INPUT",
			filename:    "input1.txt",
			expected:    513643,
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

func TestCustomHash(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    int
	}{
		{
			description: "Test 1",
			input:       "HASH",
			expected:    52,
		},
		{
			description: "Test 2",
			input:       "rn=1",
			expected:    30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if actual := customHash(tt.input); actual != tt.expected {
				t.Errorf("customHash() = %v, want %v", actual, tt.expected)
			}
		})
	}
}
