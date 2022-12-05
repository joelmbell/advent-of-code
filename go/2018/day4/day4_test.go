package day4_2018

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("test_input.txt"), 240},
		{loader.Load("input.txt"), 39584},
	}

	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("test_input.txt"), 4455},
		{loader.Load("input.txt"), 55053},
	}

	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
