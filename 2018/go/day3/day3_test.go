package day3

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input  []string
		output int
	}{
		{[]string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}, 4},
		{[]string{"#1 @ 1,1: 2x2", "#2 @ 5,1: 2x2"}, 0},
	}

	for _, test := range tests {
		result := Part1(test.input)

		if result != test.output {
			t.Errorf("failed: %v != %v", result, test.output)
		}
	}
}

func TestPart2(t *testing.T) {
	input := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	expect := "3"

	result := Part2(input)

	if expect != result {
		t.Errorf("failed: %v != %v", result, expect)
	}
}

func TestPart1WithInput(t *testing.T) {
	result := Part1WithInput()
	if result != 100261 {
		t.Errorf("Failed: %v != %v", result, 100261)
	}
}

func TestPart2WithInput(t *testing.T) {
	result := Part2WithInput()
	expect := "251"
	if result != expect {
		t.Errorf("Failed: %v != %v", result, expect)
	}
}
