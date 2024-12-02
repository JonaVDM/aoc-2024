package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(file string) []string {
	f, err := os.ReadFile(fmt.Sprintf("inputs/%s", file))
	if err != nil {
		panic("could not read input file")
	}

	inp := string(f)
	inp = strings.TrimSpace(inp)

	return strings.Split(inp, "\n")
}

func ReadSingleLineFile(file string) string {
	f, err := os.ReadFile(fmt.Sprintf("inputs/%s", file))
	if err != nil {
		panic("could not read input file")
	}

	inp := string(f)
	inp = strings.TrimSpace(inp)

	return inp
}

func ReadFileRaw(file string) []string {
	f, err := os.ReadFile(fmt.Sprintf("inputs/%s", file))
	if err != nil {
		panic("could not read input file")
	}

	inp := string(f)

	return strings.Split(inp, "\n")
}

// ReadGRidFileNum reads out `file` and returns the numbers present in it in a
// grid. The input is split on spaces, any additional empty values are skipped.
// A panic is triggered when a input could not be converted to numbers.
func ReadGridFileNum(file string) [][]int {
	f, err := os.ReadFile(fmt.Sprintf("inputs/%s", file))
	if err != nil {
		panic("could not read input file")
	}

	inp := string(f)
	inp = strings.TrimSpace(inp)

	data := strings.Split(inp, "\n")
	input := make([][]int, len(data))
	for i := range data {
		row := strings.Split(data[i], " ")
		input[i] = make([]int, 0)
		for _, c := range row {
			if c == "" {
				continue
			}

			num, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}

			input[i] = append(input[i], num)
		}
	}

	return input
}
