package day6

import (
	ds "aoc/datastructures"
	"aoc/util"
)

func findDistinctEnd(input string, length int) int {
	for i := 0; i < len(input)-length-1; i++ {
		sub := input[i : i+length]
		set := ds.NewSet(util.ConvertToRunes(sub))
		if len(set) == length {
			return i + length
		}
	}
	return -1
}

func Part1(input string) int {
	return findDistinctEnd(input, 4)
}

func Part2(input string) int {
	return findDistinctEnd(input, 14)
}
