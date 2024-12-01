package utils

func Between(i, min, max int) bool {
	return (i >= min) && (i <= max)
}

func UniqueString(s string) bool {
	var a [256]bool
	for _, v := range s {
		if a[v] {
			return false
		}
		a[v] = true
	}
	return true
}
