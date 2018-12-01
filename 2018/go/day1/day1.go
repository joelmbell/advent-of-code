package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func loadData(filename string) []int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []int64
	for scanner.Scan() {
		string := scanner.Text()
		num, _ := strconv.ParseInt(string, 10, 32)

		result = append(result, num)
	}

	return result
}

func Day1_01(input []int64) int64 {
	var result int64
	for _, num := range input {
		result += num
	}
	return result
}

func Day1_02(input []int64) int64 {
	var store map[int64]bool
	store = make(map[int64]bool)
	store[0] = true

	var current int64
	for i := 0; true; i++ {
		if i == len(input) {
			i = 0
		}

		current += input[i]
		fmt.Printf("%v ", current)

		if store[current] {
			return current
		} else {
			store[current] = true
		}
	}

	return -10
}

func Day1() int64 {
	data := loadData("input.txt")
	result := Day1_02(data)
	return result
}
