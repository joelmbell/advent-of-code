package day8

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 21},
		{loader.Load("input.txt"), 1715},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 8},
		{loader.Load("input.txt"), 374400},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
