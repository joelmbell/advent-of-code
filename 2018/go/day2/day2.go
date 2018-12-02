package day2

import (
	"bufio"
	"log"
	"os"
)

func loadData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		string := scanner.Text()

		result = append(result, string)
	}

	return result
}

func day2_01(input []string) int {

	var dupesCount int
	var triplicatesCount int

	for _, s := range input {
		store := make(map[string]int)

		for _, c := range s {
			store[string(c)] = store[string(c)] + 1
		}

		var hasDupes bool
		var hasTriplicates bool
		for _, value := range store {
			if value == 2 {
				hasDupes = true
			} else if value == 3 {
				hasTriplicates = true
			}

			if hasDupes && hasTriplicates {
				break
			}
		}

		if hasDupes {
			dupesCount++
		}
		if hasTriplicates {
			triplicatesCount++
		}
	}

	return dupesCount * triplicatesCount
}

func day2_02(input []string) string {

	wordLength := len(input[0])

	for _, line := range input {
		matches := make(map[int][]rune)

		for i, char := range line {
			for l := 0; l < len(input); l++ {
				if char == rune(input[l][i]) {
					matches[l] = append(matches[l], char)
				}
			}
		}

		for _, v := range matches {
			if len(v) == wordLength-1 {
				return string(v)
			}
		}
	}

	return ""
}

func Day2() string {
	input := loadData("input.txt")
	return day2_02(input)
}
