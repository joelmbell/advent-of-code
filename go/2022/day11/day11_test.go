package day11

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 10605},
		{loader.Load("input.txt"), 117624},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]string, int]{
		{loader.Load("sample.txt"), 2713310158},
		{loader.Load("input.txt"), 16792940265},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
