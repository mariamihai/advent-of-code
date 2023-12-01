package day1

import "testing"

func TestProblem1(t *testing.T) {
	testCases := []struct {
		description string
		filename    string
		expected    int
	}{
		{
			description: "Should validate PROBLEM 1 against provided EXAMPLE",
			filename:    "example1.txt",
			expected:    142,
		},
		{
			description: "Should validate PROBLEM 1 against provided INPUT",
			filename:    "input1.txt",
			expected:    54916,
		},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			if actual := Problem1(test.filename); actual != test.expected {
				t.Errorf("problem1() = %v, want %v", actual, test.expected)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	testCases := []struct {
		description string
		filename    string
		expected    int
	}{
		{
			description: "Should validate PROBLEM 2 against provided EXAMPLE",
			filename:    "example2.txt",
			expected:    281,
		},
		{
			description: "Should validate PROBLEM 2 against provided INPUT",
			filename:    "input2.txt",
			expected:    54728,
		},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			if actual := Problem2(test.filename); actual != test.expected {
				t.Errorf("problem2() = %v, want %v", actual, test.expected)
			}
		})
	}
}

func TestCalibratedValueForLine(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    int
	}{
		{
			description: "Should validate single digit x to xx",
			input:       "abc1abc",
			expected:    11,
		},
		{
			description: "Should set first and last characters when they are digits",
			input:       "1abc2",
			expected:    12,
		},
		{
			description: "Should set first and last digits when there are multiple digits",
			input:       "abc1abc2abc3abc4abc",
			expected:    14,
		},
		{
			description: "Should validate empty string to 0",
			input:       "",
			expected:    0,
		},
		{
			description: "Should validate string without digits to 0",
			input:       "abc",
			expected:    0,
		},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			if actual := calibratedValueForLine(test.input); actual != test.expected {
				t.Errorf("calibratedValueForLine() = %v, want %v", actual, test.expected)
			}
		})
	}
}

func TestReplacedLettersToDigits(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    int
	}{
		{
			description: "Should validate letters to digits",
			input:       "7pqrstsixteen",
			expected:    76,
		},
		{
			description: "Should validate overlapping letters of numbers",
			input:       "eightwothree",
			expected:    83,
		},
		{
			description: "Should validate combination of digits as letters and digits",
			input:       "zoneight234",
			expected:    14,
		},
	}
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			replacedInput := replacedLettersToDigits(test.input)
			actual := calibratedValueForLine(replacedInput)

			if actual != test.expected {
				t.Errorf("replaced letters and calibrated = %v, want %v", actual, test.expected)
			}
		})
	}
}
