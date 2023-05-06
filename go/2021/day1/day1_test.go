package day1_2019

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 7},
		{loader.Load("input.txt"), 1228},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 5},
		{loader.Load("input.txt"), 1257},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
