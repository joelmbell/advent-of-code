package util

func ConvertToRunes(str string) []rune {
	output := make([]rune, 0)
	for _, char := range str {
		output = append(output, char)
	}
	return output
}
