package day1_2018

import (
	loader "aoc/dataloader"
	"aoc/test"
	"testing"
)

func TestPart1(t *testing.T) {

	cases := []test.Case[[]int64, int64]{
		{[]int64{-1, 1}, 0},
		{[]int64{-5, -4, 9, 5, -4, -3, 5}, 3},
		{loader.LoadInt("input.txt"), 590},
	}

	err := test.Execute(cases, Part1)
	if err != nil {
		t.Error(err)
	}
}

func TestPart2(t *testing.T) {
	cases := []test.Case[[]int64, int64]{
		{[]int64{1, -1}, 0},
		{[]int64{3, 3, 4, -2, -4}, 10},
		{[]int64{-6, 3, 8, 5, -6}, 5},
		{[]int64{7, 7, -2, -7, -4}, 14},
		{loader.LoadInt("input.txt"), 83445},
	}

	err := test.Execute(cases, Part2)
	if err != nil {
		t.Error(err)
	}
}

//func TestDay1WithInput(t *testing.T) {
//	output := Day1()
//	if output != 83445 {
//		t.Errorf("output was %v", output)
//	}
//}
