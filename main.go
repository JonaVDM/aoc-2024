package main

import (
	"flag"
	"fmt"
	"log/slog"
	"time"

	"github.com/jonavdm/aoc-2024/day01"
	"github.com/jonavdm/aoc-2024/day02"
	"github.com/jonavdm/aoc-2024/day03"
	"github.com/jonavdm/aoc-2024/day04"
	"github.com/jonavdm/aoc-2024/day05"
	"github.com/jonavdm/aoc-2024/day06"
	"github.com/jonavdm/aoc-2024/day07"
	"github.com/jonavdm/aoc-2024/day08"
	"github.com/jonavdm/aoc-2024/day09"
	"github.com/jonavdm/aoc-2024/day10"
	"github.com/jonavdm/aoc-2024/day11"
	"github.com/jonavdm/aoc-2024/day12"
	"github.com/jonavdm/aoc-2024/day13"
	"github.com/jonavdm/aoc-2024/day14"
	"github.com/jonavdm/aoc-2024/day15"
	"github.com/jonavdm/aoc-2024/day16"
	"github.com/jonavdm/aoc-2024/day17"
	"github.com/jonavdm/aoc-2024/day18"
	"github.com/jonavdm/aoc-2024/day19"
	"github.com/jonavdm/aoc-2024/day20"
	"github.com/jonavdm/aoc-2024/day21"
	"github.com/jonavdm/aoc-2024/day22"
	"github.com/jonavdm/aoc-2024/day23"
	"github.com/jonavdm/aoc-2024/day24"
	"github.com/jonavdm/aoc-2024/day25"
	"github.com/jonavdm/aoc-2024/utils"
)

type Runner struct {
	Day      int
	Function func([]string) [2]interface{}
	File     string
}

func main() {
	onlyDay := flag.Int("day", -1, "Specify the day")
	replacedInput := flag.String("file", "", "Run with a different input")
	onlyTime := flag.Bool("time", false, "Show only the time to completion")
	debug := flag.Bool("debug", false, "Show debug messages")
	flag.Parse()

	if *debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	runners := []Runner{
		{1, day01.Run, "day01"},
		{2, day02.Run, "day02"},
		{3, day03.Run, "day03"},
		{4, day04.Run, "day04"},
		{5, day05.Run, "day05"},
		{6, day06.Run, "day06"},
		{7, day07.Run, "day07"},
		{8, day08.Run, "day08"},
		{9, day09.Run, "day09"},
		{10, day10.Run, "day10"},
		{11, day11.Run, "day11"},
		{12, day12.Run, "day12"},
		{13, day13.Run, "day13"},
		{14, day14.Run, "day14"},
		{15, day15.Run, "day15"},
		{16, day16.Run, "day16"},
		{17, day17.Run, "day17"},
		{18, day18.Run, "day18"},
		{19, day19.Run, "day19"},
		{20, day20.Run, "day20"},
		{21, day21.Run, "day21"},
		{22, day22.Run, "day22"},
		{23, day23.Run, "day23"},
		{24, day24.Run, "day24"},
		{25, day25.Run, "day25"},
	}

	printHeader(*onlyTime)

	var totalDuration time.Duration

	for _, runner := range runners {
		var out [2]interface{}
		file := runner.File

		if *onlyDay > 0 && runner.Day != *onlyDay {
			continue
		}

		if *onlyDay > 0 && *replacedInput != "" {
			file = *replacedInput
		}

		data := utils.ReadFile(file)

		start := time.Now()
		out = runner.Function(data)

		duration := time.Since(start)
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
		fmt.Printf(" ┃ %2d  │ %-16v │ %-16v │ %-11v ┃\n", day, out[0], out[1], duration.Round(time.Microsecond))
	} else {
		fmt.Printf(" ┃ %2d    │ %-18v ┃\n", day, duration.Round(time.Microsecond))
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
