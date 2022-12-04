package day1_2018

import (
	"testing"
)

func TestDay1_01(t *testing.T) {

	tests := []struct {
		input  []int64
		output int64
	}{
		{[]int64{-1, 1}, 0},
		{[]int64{-5, -4, 9, 5, -4, -3, 5}, 3},
	}

	for _, test := range tests {
		output := Day1_01(test.input)
		if output != test.output {
			t.Errorf("%v != %v", output, test.output)
		}
	}
}

func TestDay1_02(t *testing.T) {
	tests := []struct {
		input  []int64
		output int64
	}{
		{[]int64{1, -1}, 0},
		{[]int64{3, 3, 4, -2, -4}, 10},
		{[]int64{-6, 3, 8, 5, -6}, 5},
		{[]int64{7, 7, -2, -7, -4}, 14},
	}

	for _, test := range tests {
		output := Day1_02(test.input)
		if output != test.output {
			t.Errorf("%v\n\t%v != %v", test.input, output, test.output)
		}
	}
}

//func TestDay1WithInput(t *testing.T) {
//	output := Day1()
//	if output != 83445 {
//		t.Errorf("output was %v", output)
//	}
//}
