package dataloader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Load(filename string) []string {
	data, err := loadData(filename)
	if err != nil {
		fmt.Printf("ERROR LOADING DATA for %s -- %v", filename, err)
		panic(err)
	}
	return data
}

func LoadInt(filename string) []int64 {
	data := Load(filename)
	var output []int64
	for _, item := range data {
		num, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			fmt.Printf("cannot convert `%v` to integer", item)
			panic(err)
		}
		output = append(output, num)
	}
	return output
}

func loadData(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	var result []string
	for scanner.Scan() {
		str := scanner.Text()
		//num, _ := strconv.ParseInt(string, 10, 32)

		result = append(result, str)
	}

	return result, err
}
