package day1_2019

import (
	"fmt"
	"strconv"
)

func Part1(input []string) int {

	prevNum := 0
	count := 0
	for _, val := range input {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic("invalid input")
		}
		if num > prevNum {
			count++
		}
		prevNum = num
	}

	return count - 1
}

func Part2(input []string) int {
	prevNum := 0
	count := 0
	for i := 0; i < len(input); i++ {
		slice := input[i : i+3]
		total := 0
		for _, piece := range slice {
			num, _ := strconv.Atoi(piece)
			total += num
		}
		fmt.Printf("\ntotal: %v", total)

		if total > prevNum {
			fmt.Printf("\n\tincreased count")
			count++
		}

		prevNum = total
	}
	return count - 1
}
