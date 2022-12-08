package day5_2018

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []test.Case[string, int]{
		{loader.Load("sample.txt")[0], 10},
		{loader.Load("input.txt")[0], 10888},
	}
	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[string, int]{
		{loader.Load("sample.txt")[0], 4},
		{loader.Load("input.txt")[0], 6952},
	}
	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
