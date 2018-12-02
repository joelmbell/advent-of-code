package day2

import (
	"testing"
)

func TestPart1(t *testing.T) {

	tests := []struct {
		input  []string
		output int
	}{
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}

	for _, test := range tests {
		result := part1(test.input)
		if result != test.output {
			t.Fatalf("%v != %v -- for case: %v", result, test.output, test.input)
		}
	}
}

func TestPart2(t *testing.T) {

	tests := []struct {
		input  []string
		output string
	}{
		{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
	}

	for _, test := range tests {
		result := part2(test.input)
		if result != test.output {
			t.Fatalf("%v != %v -- for case: %v", result, test.output, test.input)
		}
	}
}

func TestPart1WithInput(t *testing.T) {
	result := Part1WithInput()

	if result != 3952 {
		t.Errorf("Failed: result should be %v", result)
	}
}

func TestPart2WithInput(t *testing.T) {
	result := Part2WithInput()

	if result != "vtnikorkulbfejvyznqgdxpaw" {
		t.Errorf("Failed: result should be %v", result)
	}
}
