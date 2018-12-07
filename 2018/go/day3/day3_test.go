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

func TestPart1WithInput(t *testing.T) {
	result := Part1WithInput()

	if result != 100 {
		t.Errorf("Failed: %v != %v", result, 100261)
	}
}
