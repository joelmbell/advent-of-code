package day1

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int64]{
		{loader.Load("sample.txt"), 24000},
		{loader.Load("input.txt"), 71300},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int64]{
		{loader.Load("sample.txt"), 45000},
		{loader.Load("input.txt"), 209691},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
