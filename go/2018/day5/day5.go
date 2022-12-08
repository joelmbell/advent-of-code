package day5_2018

import (
	"aoc/util"
	"strings"
)

func find(haystack []rune, needle rune) int {
	for i, item := range haystack {
		if item == needle {
			return i
		}
	}
	return -1
}

func delete(s []rune, index int, count int) []rune {
	return append(s[0:index], s[index+count:]...)
}

var lowercase = util.ConvertToRunes("abcdefghijklmnopqrstuvwxyz")
var uppercase = util.ConvertToRunes("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var matchingCharacters = make(map[rune]rune)

func isPair(a rune, b rune) bool {
	if compare, ok := matchingCharacters[a]; ok {
		if compare == b {
			return true
		}
	}

	index := find(lowercase, a)
	var pair rune
	if index == -1 {
		index = find(uppercase, a)
		pair = lowercase[index]
	} else {
		pair = uppercase[index]
	}
	isPair := b == pair || a == pair
	if isPair {
		matchingCharacters[a] = b
		matchingCharacters[b] = a
	}
	return isPair
}

func Part1(input string) int {
	data := util.ConvertToRunes(input)

	for i := 0; true; {
		if len(data) <= i+1 {
			break
		}
		a := data[i]
		b := data[i+1]

		if isPair(a, b) {
			data = delete(data, i, 2)
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}

	return len(data)
}

func Part2(input string) int {
	lowercase := util.ConvertToRunes("abcdefghijklmnopqrstuvwxyz")
	uppercase := util.ConvertToRunes("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	inputs := make([]string, 0)
	for i := 0; i < len(lowercase); i++ {
		stripped := strings.Replace(input, string(lowercase[i]), "", -1)
		stripped = strings.Replace(stripped, string(uppercase[i]), "", -1)
		inputs = append(inputs, stripped)
	}
	shortest := len(input)
	for _, value := range inputs {
		processed := Part1(value)
		if shortest > processed {
			shortest = processed
		}
	}
	return shortest
}
