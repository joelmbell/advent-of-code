package day4

import (
	ds "aoc/datastructures"
	"strconv"
	"strings"
)

type rng struct {
	start int
	end   int
}

func parse(input string) []rng {
	rangeStrings := strings.Split(input, ",")

	var ranges []rng
	for _, strRng := range rangeStrings {
		points := strings.Split(strRng, "-")
		start, _ := strconv.Atoi(points[0])
		end, _ := strconv.Atoi(points[1])
		ranges = append(ranges, rng{start, end})
	}
	return ranges
}

func sliceForRng(input rng) []int {
	var output []int
	for i := input.start; i <= input.end; i++ {
		output = append(output, i)
	}
	return output
}

func Part1(input []string) int {
	count := 0
	for _, row := range input {
		pair := parse(row)
		first := ds.NewSet(sliceForRng(pair[0]))
		second := ds.NewSet(sliceForRng(pair[1]))
		all := first.Union(second)
		if first.Equals(all) || second.Equals(all) {
			count++
		}
	}
	return count
}

func Part2(input []string) int {
	count := 0
	for _, row := range input {
		pair := parse(row)
		first := ds.NewSet(sliceForRng(pair[0]))
		second := ds.NewSet(sliceForRng(pair[1]))
		inter := first.Intersection(second)
		if len(inter) > 0 {
			count++
		}
	}
	return count
}
