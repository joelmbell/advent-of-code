package util

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TowardZero(num int) int {
	if num == 0 {
		return 0
	} else {
		if num < 0 {
			return num + 1
		} else {
			return num - 1
		}
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
