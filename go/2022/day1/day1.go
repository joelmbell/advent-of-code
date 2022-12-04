package day1

import (
	"sort"
	"strconv"
)

func Part1(data []string) int64 {
	elvesCalories := getTotalCaloriesForEachElfSorted(data)
	return elvesCalories[0]
}

func Part2(data []string) int64 {
	elvesCalories := getTotalCaloriesForEachElfSorted(data)
	highest := elvesCalories[0:3]

	var total int64
	for _, count := range highest {
		total += count
	}
	return total
}

func getTotalCaloriesForEachElfSorted(data []string) []int64 {
	elvesCalories := make([]int64, 1)
	curElf := 0
	for _, row := range data {
		if row == "" {
			curElf += 1
			elvesCalories = append(elvesCalories, 0)
			continue
		}

		calorieCount, _ := strconv.ParseInt(row, 10, 32)
		elvesCalories[curElf] = elvesCalories[curElf] + calorieCount
	}

	sort.Slice(elvesCalories, func(i, j int) bool {
		return elvesCalories[i] > elvesCalories[j]
	})

	return elvesCalories
}
