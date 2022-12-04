package day2_2018

func Part1(input []string) int {

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

func Part2(input []string) string {

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
