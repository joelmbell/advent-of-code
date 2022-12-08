package day8

import (
	"strconv"
)

func parse(input []string) [][]int {
	output := make([][]int, 0)
	for _, line := range input {
		row := make([]int, 0)
		for _, tree := range line {
			item, _ := strconv.Atoi(string(tree))
			row = append(row, item)
		}
		output = append(output, row)
	}
	return output
}

func Part1(input []string) int {
	data := parse(input)
	width := len(input[0])
	height := len(input)
	visibleCount := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || j == 0 || i == height-1 || j == width-1 {
				visibleCount++
				continue
			}

			// check left edge
			l := 0
			for ; l < j; l++ {
				if data[i][l] >= data[i][j] {
					break
				}
			}
			if l == j {
				visibleCount++
				continue
			}

			// check right edge
			r := width - 1
			for ; r > j; r-- {
				if data[i][r] >= data[i][j] {
					break
				}
			}
			if r == j {
				visibleCount++
				continue
			}

			// check top
			t := 0
			for ; t < i; t++ {
				if data[t][j] >= data[i][j] {
					break
				}
			}
			if t == i {
				visibleCount++
				continue
			}

			// check bottom
			b := height - 1
			for ; b > i; b-- {
				if data[b][j] >= data[i][j] {
					break
				}
			}
			if b == i {
				visibleCount++
				continue
			}
		}
	}
	return visibleCount
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Part2(input []string) int {
	data := parse(input)
	width := len(input[0])
	height := len(input)

	highestScenicScore := 1
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			scenicScore := 1
			// check left edge
			l := j - 1
			if j == 0 {
				l = j
			}
			for ; l > 0; l-- {
				if data[i][l] >= data[i][j] {
					break
				}
			}
			left := j - l

			// check right edge
			r := j + 1
			if r == width {
				r = j
			}
			for ; r < width-1; r++ {
				if data[i][r] >= data[i][j] {
					break
				}
			}
			right := r - j

			// check top
			t := i - 1
			if t < 0 {
				t = i
			}
			for ; t > 0; t-- {
				if data[t][j] >= data[i][j] {
					break
				}
			}
			top := max(i-t, 0)

			// check bottom
			b := i + 1
			if b == height {
				b = i
			}
			for ; b < height-1; b++ {
				if data[b][j] >= data[i][j] {
					break
				}
			}
			bottom := b - i
			scenicScore = left * right * bottom * top

			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}
	return highestScenicScore
}
