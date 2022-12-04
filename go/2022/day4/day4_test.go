package day4

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 2},
		{loader.Load("input.txt"), 538},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}
