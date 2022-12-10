package day9

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 13},
		{loader.Load("input.txt"), 6190},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 1},
		{loader.Load("sample2.txt"), 36},
		{loader.Load("input.txt"), 2516},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
