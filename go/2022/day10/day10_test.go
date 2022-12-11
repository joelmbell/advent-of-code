package day10

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 13140},
		{loader.Load("input.txt"), 17020},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, string]{
		{loader.Load("sample.txt"), ""}, // output is visual
		{loader.Load("input.txt"), ""},  // output is visual (RLEZFLGE)
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
