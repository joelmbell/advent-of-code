package day1_2018

func Part1(input []int64) int64 {
	var result int64
	for _, num := range input {
		result += num
	}
	return result
}

func Part2(input []int64) int64 {
	var store map[int64]bool
	store = make(map[int64]bool)
	store[0] = true

	var current int64
	for i := 0; true; i++ {
		if i == len(input) {
			i = 0
		}

		current += input[i]

		if store[current] {
			return current
		} else {
			store[current] = true
		}
	}

	return -10
}
