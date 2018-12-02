package day2

import (
	"testing"
)

func TestDay2_1(t *testing.T) {

	tests := []struct {
		input  []string
		output int
	}{
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
	}

	for _, test := range tests {
		result := day2_01(test.input)
		if result != test.output {
			t.Fatalf("%v != %v -- for case: %v", result, test.output, test.input)
		}
	}
}

func TestDay2_2(t *testing.T) {

	tests := []struct {
		input  []string
		output string
	}{
		{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
	}

	for _, test := range tests {
		result := day2_02(test.input)
		if result != test.output {
			t.Fatalf("%v != %v -- for case: %v", result, test.output, test.input)
		}
	}
}

func TestDay2(t *testing.T) {
	result := Day2()

	if result != "vtnikorkulbfejvyznqgdxpaw" {
		t.Errorf("Failed: result should be %v", result)
	}
}
