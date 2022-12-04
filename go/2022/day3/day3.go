package day3

import (
	ds "aoc/datastructures"
)

type rucksack struct {
	original string
}

func (r *rucksack) comp1() string {
	mid := len(r.original) / 2
	return r.original[:mid]
}

func (r *rucksack) comp2() string {
	mid := len(r.original) / 2
	return r.original[mid:]
}

func newRucksack(input string) rucksack {
	return rucksack{input}
}

func (r *rucksack) FindDupe() rune {
	a := ds.NewSet(convert(r.comp1()))
	b := ds.NewSet(convert(r.comp2()))

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

var key = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Part1(input []string) int {
	total := 0
	for _, item := range input {
		rs := newRucksack(item)
		dupe := rs.FindDupe()

		for i, char := range key {
			if dupe == char {
				total += i + 1
				break
			}
		}
	}

	return total
}

func Part2(input []string) int {
	total := 0
	for i := 0; i < len(input); i += 3 {
		elf1 := ds.NewSet(convert(input[i]))
		elf2 := ds.NewSet(convert(input[i+1]))
		elf3 := ds.NewSet(convert(input[i+2]))

		dupes := elf1.Intersection(elf2).Intersection(elf3)
		var dupe rune
		for key := range dupes {
			dupe = key
		}

		for i, char := range key {
			if dupe == char {
				total += i + 1
				break
			}
		}
	}

	return total
}
