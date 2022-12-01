package dataloader

import (
	"bufio"
	"fmt"
	"os"
)

func Load(filename string) []string {
	data, err := loadData(filename)
	if err != nil {
		fmt.Printf("ERROR LOADING DATA for %s -- %v", filename, err)
		panic(err)
	}
	return data
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
