package day5

import (
	"github.com/samber/lo"
	"reflect"
	"testing"
)

func TestMapSeedForProblem2(t *testing.T) {
	tests := []struct {
		description string
		inputSeed   seedWithRange
		converter   almanacLine
		expected    []seedWithRange
	}{
		{
			//    [ A, B]          seed
			//            [ C, D]  convert
			description: "Should not intersect (seed before convert range)",
			inputSeed:   seedWithRange{min: 79, rangeLength: 14},
			converter:   almanacLine{destination: -100, source: 100, rangeLength: 10},
			expected:    []seedWithRange{{min: 79, rangeLength: 14}},
		},
		{
			//             [ C, D]  seed
			//    [ A, B]           convert
			description: "Should not intersect (seed after convert range)",
			inputSeed:   seedWithRange{min: 79, rangeLength: 14},
			converter:   almanacLine{destination: -100, source: 60, rangeLength: 10},
			expected:    []seedWithRange{{min: 79, rangeLength: 14}},
		},
		{
			//    [ A,         D]  seed
			//        [ B, C]      convert
			description: "Should intersect [A, D] seed, [B, C] convert",
			inputSeed:   seedWithRange{min: 79, rangeLength: 14},
			converter:   almanacLine{destination: 50, source: 85, rangeLength: 5},
			expected:    []seedWithRange{{min: 50, rangeLength: 5}, {min: 79, rangeLength: 6}, {min: 90, rangeLength: 3}},
		},
		{
			//    [ A,    C]     seed
			//        [ B,   D]  convert
			description: "Should intersect [A, C] seed, [B, D] convert",
			inputSeed:   seedWithRange{min: 79, rangeLength: 14},
			converter:   almanacLine{destination: 50, source: 90, rangeLength: 10},
			expected:    []seedWithRange{{min: 50, rangeLength: 3}, {min: 79, rangeLength: 11}},
		},
		{
			//    	[ B,    D]  seed
			//    [ A,   C]     convert
			description: "Should intersect [B, D] seed, [A, C] convert",
			inputSeed:   seedWithRange{min: 79, rangeLength: 14},
			converter:   almanacLine{destination: 50, source: 60, rangeLength: 30},
			expected:    []seedWithRange{{min: 50, rangeLength: 11}, {min: 90, rangeLength: 3}},
		},
		{
			//        [ B, C]      seed
			//    [ A,         D]  convert
			description: "Should intersect [B, C] seed, [A, D] convert",
			inputSeed:   seedWithRange{min: 79, rangeLength: 14},
			converter:   almanacLine{destination: 50, source: 60, rangeLength: 40},
			expected:    []seedWithRange{{min: 50, rangeLength: 14}},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actual := mapSeedForProblem2(test.inputSeed, test.converter)

			subset := lo.Without(test.expected, actual...)
			if len(subset) != 0 {
				t.Errorf("mapSeedForProblem2() = %v, want %v", actual, test.expected)
			}
		})
	}
}

func TestMapSeedsForProblem2(t *testing.T) {
	tests := []struct {
		description string
		seeds       []seedWithRange
		converter   []almanacLine
		expected    []seedWithRange
	}{
		{
			description: "Should map",
			seeds:       []seedWithRange{{min: 79, rangeLength: 14}, {min: 55, rangeLength: 13}},
			converter: []almanacLine{
				{destination: 50, source: 98, rangeLength: 2},
				{destination: 52, source: 50, rangeLength: 48},
			},
			expected: []seedWithRange{{min: 52, rangeLength: 14}},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := mapSeedsForProblem2(test.seeds, test.converter); !reflect.DeepEqual(got, test.expected) {
				t.Errorf("mapSeedsForProblem2() = %v, want %v", got, test.expected)
			}
		})
	}
}
