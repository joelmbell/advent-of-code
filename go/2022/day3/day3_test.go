package day3

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 157},
		{loader.Load("input.txt"), 7821},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 70},
		{loader.Load("input.txt"), 2752},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
