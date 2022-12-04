package day2_2018

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {

	cases := []test.Case[[]string, int]{
		{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}, 12},
		{loader.Load("input.txt"), 3952},
	}

	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {

	cases := []test.Case[[]string, string]{
		{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}, "fgij"},
		{loader.Load("input.txt"), "vtnikorkulbfejvyznqgdxpaw"},
	}

	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}
