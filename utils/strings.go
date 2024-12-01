package utils

func RuneInString(bit rune, str string) bool {
	for _, char := range str {
		if bit == char {
			return true
		}
	}

	return false
}
