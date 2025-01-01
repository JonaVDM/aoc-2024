package utils

// GetKeys returns the indexing keys of the map m
func GetKeys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, len(m))

	i := 0
	for t := range m {
		keys[i] = t
		i++
	}

	return keys
}
