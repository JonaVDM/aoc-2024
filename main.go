package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jonavdm/aoc-2024/day01"
	"github.com/jonavdm/aoc-2024/day02"
	"github.com/jonavdm/aoc-2024/day03"
)

type Runner struct {
	Day      int
	Function func(string) [2]interface{}
	File     string
}

func main() {
	onlyDay := flag.Int("day", -1, "Specify the day")
	replacedInput := flag.String("file", "", "Run with a different input")
	onlyTime := flag.Bool("time", false, "Show only the time to completion")
	flag.Parse()

	runners := []Runner{
		{1, day01.Run, "day01"},
		{2, day02.Run, "day02"},
		{3, day03.Run, "day03"},
	}

	printHeader(*onlyTime)

	var totalDuration time.Duration

	for _, runner := range runners {
		start := time.Now()
		var out [2]interface{}

		if *onlyDay > 0 && runner.Day == *onlyDay {
			file := runner.File
			if *replacedInput != "" {
				file = *replacedInput
			}
			out = runner.Function(file)
		}

		if *onlyDay == -1 {
			out = runner.Function(runner.File)
		}

		duration := time.Now().Sub(start)
		totalDuration += duration

		if out[0] != nil {
			printRow(runner.Day, out, duration, *onlyTime)
		}
	}
	printFooter(*onlyDay, totalDuration, *onlyTime)
}

func printHeader(onlyTime bool) {
	if !onlyTime {
		fmt.Println(" ╔═════════════════════════════════════════════════════════╗")
		fmt.Println(" ║                -- Advent Of Code 2024 --                ║")
		fmt.Println(" ╚═════════════════════════════════════════════════════════╝")
		fmt.Println()
		fmt.Println(" ┏━━━━━┯━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━┓")
		fmt.Println(" ┃ Day │ Part One         │ Part Two         │ Time        ┃")
		fmt.Println(" ┠─────┼──────────────────┼──────────────────┼─────────────┨")
	} else {
		fmt.Println(" ╔════════════════════════════╗")
		fmt.Println(" ║ -- Advent Of Code 2024 --  ║")
		fmt.Println(" ╚════════════════════════════╝")
		fmt.Println()
		fmt.Println(" ┏━━━━━━━┯━━━━━━━━━━━━━━━━━━━━┓")
		fmt.Println(" ┃ Day   │ Time               ┃")
		fmt.Println(" ┠───────┼────────────────────┨")
	}
}

func printRow(day int, out [2]interface{}, duration time.Duration, onlyTime bool) {
	if !onlyTime {
		fmt.Printf(" ┃ %2d  │ %-16v │ %-16v │ %-11v ┃\n", day, out[0], out[1], duration)
	} else {
		fmt.Printf(" ┃ %2d    │ %-18v ┃\n", day, duration)
	}
}

func printFooter(onlyDay int, duration time.Duration, onlyTime bool) {
	if !onlyTime {
		fmt.Println(" ┗━━━━━┷━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━┛")

		if onlyDay == -1 {
			fmt.Println("  Total duration", duration)
		}
	} else {
		fmt.Println(" ┠───────┼────────────────────┨")
		fmt.Printf(" ┃ Total │ %-18v ┃\n", duration)
		fmt.Println(" ┗━━━━━━━┷━━━━━━━━━━━━━━━━━━━━┛")
	}
}
