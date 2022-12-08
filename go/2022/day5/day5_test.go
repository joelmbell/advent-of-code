package day5

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, string]{
		{loader.Load("sample.txt"), "CMZ"},
		{loader.Load("input.txt"), "MQSHJMWNH"},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, string]{
		{loader.Load("sample.txt"), "MCD"},
		{loader.Load("input.txt"), "LLWJRBHVZ"},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
