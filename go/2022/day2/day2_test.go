package day2

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 15},
		{loader.Load("input.txt"), 14069},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 12},
		{loader.Load("input.txt"), 12411},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
