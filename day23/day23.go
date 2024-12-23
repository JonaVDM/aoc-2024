package day23

import (
	"slices"
	"strings"
)

func Run(data []string) [2]interface{} {
	computers := Parse(data)
	return [2]interface{}{
		FindTriples(computers),
		0,
	}
}

func Parse(data []string) map[string][]string {
	computers := make(map[string][]string)

	for _, d := range data {
		spl := strings.Split(d, "-")

		if _, ok := computers[spl[0]]; !ok {
			computers[spl[0]] = make([]string, 0)
		}
		if _, ok := computers[spl[1]]; !ok {
			computers[spl[1]] = make([]string, 0)
		}

		computers[spl[0]] = append(computers[spl[0]], spl[1])
		computers[spl[1]] = append(computers[spl[1]], spl[0])
	}

	return computers
}

func FindTriples(m map[string][]string) int {
	match := make(map[string]int)

	for k, v := range m {
		for _, k1 := range v {
			for _, k2 := range m[k1] {
				if k2 == k {
					continue
				}

				s := []string{k, k1, k2}
				slices.Sort(s)
				if slices.Contains(m[k2], k) {
					match[strings.Join(s, "-")]++
				}
			}
		}
	}

	var count int
	for k := range match {
		if k[0] == 't' || k[3] == 't' || k[6] == 't' {
			count++
		}
	}

	return count
}
