package day19

import (
	"slices"
	"strings"
	"sync"
)

func Run(data []string) [2]interface{} {
	towels := strings.Split(data[0], ", ")

	var count int
	var total int
	countChan := make(chan int)
	closeChan := make(chan int)
	var wg sync.WaitGroup
	for _, pat := range data[2:] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := MatchCount(towels, pat)
			if c > 0 {
				countChan <- c
			}
		}()
	}
	go func() {
		wg.Wait()
		close(countChan)
		closeChan <- 1
	}()

OUTER:
	for {
		select {
		case c := <-countChan:
			count++
			total += c
		case <-closeChan:
			break OUTER
		}
	}

	close(closeChan)

	return [2]interface{}{
		count,
		total,
	}
}

func MatchCount(towels []string, pattern string) int {
	q := make([]string, 1)
	q[0] = pattern
	count := make(map[string]int)
	count[pattern] = 1

	for len(q) > 0 {
		slices.SortFunc(q, func(a, b string) int {
			return len(b) - len(a)
		})
		i := q[0]
		q = q[1:]

		for _, n := range FindNext(towels, i) {
			if _, ok := count[n]; !ok {
				q = append(q, n)
			}

			count[n] += count[i]
		}
	}

	return count[""]
}

func FindNext(towels []string, pattern string) []string {
	out := make([]string, 0)
	for _, p := range towels {
		if strings.HasPrefix(pattern, p) {
			out = append(out, strings.Replace(pattern, p, "", 1))
		}
	}
	return out
}
