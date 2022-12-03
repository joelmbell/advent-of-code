package day3

import (
	ds "2022/datastructures"
)

type rucksack struct {
	comp1 string
	comp2 string
}

func newRucksack(input string) rucksack {
	mid := len(input) / 2
	first := input[:mid]
	second := input[mid:]
	return rucksack{first, second}
}

func (r *rucksack) FindDupe() rune {
	a := ds.NewSet(convert(r.comp1))
	b := ds.NewSet(convert(r.comp2))

	dupes := a.Intersection(b)
	var dupe rune
	for key := range dupes {
		dupe = key
	}
	
	return dupe
}

func convert(str string) []rune {
	var output []rune
	for _, char := range str {
		output = append(output, char)
	}
	return output
}

func Part1(input []string) int {
	lowercase := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	total := 0
	for _, item := range input {
		rs := newRucksack(item)
		dupe := rs.FindDupe()

		for i, char := range lowercase {
			if dupe == char {
				total += i + 1
				break
			}
		}
	}

	return total
}
