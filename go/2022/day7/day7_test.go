package day7

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 95437},
		{loader.Load("input.txt"), 1447046},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 24933642},
		{loader.Load("input.txt"), 578710},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
