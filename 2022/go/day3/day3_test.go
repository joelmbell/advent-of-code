package day3

import (
	loader "2022/dataloader"
	"2022/test"
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
