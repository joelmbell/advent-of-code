package day3_2018

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{[]string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}, 4},
		{[]string{"#1 @ 1,1: 2x2", "#2 @ 5,1: 2x2"}, 0},
		{loader.Load("input.txt"), 100261},
	}

	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, string]{
		{[]string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}, "3"},
		{loader.Load("input.txt"), "251"},
	}

	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
