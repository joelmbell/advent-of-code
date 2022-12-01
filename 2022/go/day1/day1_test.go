package day1

import (
	"2022/dataloader"
	"testing"
)

func TestPart1(t *testing.T) {

	tests := []struct {
		input  string
		output int64
	}{
		{"sample.txt", 24000},
		{"input.txt", 71300},
	}

	for _, test := range tests {
		data := dataloader.Load(test.input)
		output := Part1(data)

		if output != test.output {
			t.Errorf("%v != %v", output, test.output)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input  string
		output int64
	}{
		{"sample.txt", 45000},
		{"input.txt", 209691},
	}

	for _, test := range tests {
		data := dataloader.Load(test.input)
		output := Part2(data)

		if output != test.output {
			t.Errorf("%v != %v", output, test.output)
		}
	}
}
