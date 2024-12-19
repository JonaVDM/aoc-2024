package day19

import (
	"log/slog"
	"strings"
	"time"
)

func Run(data []string) [2]interface{} {
	towels := strings.Split(data[0], ", ")

	var count int
	now := time.Now()
	for i, pat := range data[2:] {
		if len(FindMatch(towels, pat)) != 0 {
			count++
		}
		slog.Debug("checked a pattern", "index", i, "time", time.Since(now))
		now = time.Now()
	}

	return [2]interface{}{
		count,
		0,
	}
}

func FindNext(towels []string, pattern string) []int {
	out := make([]int, 0)
	for i, p := range towels {
		if strings.HasPrefix(pattern, p) {
			out = append(out, i)
		}
	}
	return out
}

func FindMatch(towels []string, pattern string) [][]int {
	out := make([][]int, 0)

	for _, n := range FindNext(towels, pattern) {
		if towels[n] == pattern {
			return [][]int{{n}}
		}

		next := FindMatch(towels, pattern[len(towels[n]):])
		for _, o := range next {
			out = append(out, append([]int{n}, o...))
		}
	}

	return out
}
