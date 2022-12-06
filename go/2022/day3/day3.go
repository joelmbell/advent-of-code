package day3

import (
	ds "aoc/datastructures"
	"aoc/util"
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
	a := ds.NewSet(util.ConvertToRunes(r.comp1()))
	b := ds.NewSet(util.ConvertToRunes(r.comp2()))

	dupes := a.Intersection(b)
	var dupe rune
	for key := range dupes {
		dupe = key
	}

	return dupe
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
		elf1 := ds.NewSet(util.ConvertToRunes(input[i]))
		elf2 := ds.NewSet(util.ConvertToRunes(input[i+1]))
		elf3 := ds.NewSet(util.ConvertToRunes(input[i+2]))

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
