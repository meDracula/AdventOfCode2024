package run

import (
	"fmt"

	"aoc2024/days/day1"
	"aoc2024/days/day2"
)

type AdventSolveDay func(filename string)

var AdventOfDay = map[int]AdventSolveDay{
	1: day1.AdventSolveDay1,
	2: day2.AdventSolveDay2,
}

func Day(day int, file string) {
	if solve, ok := AdventOfDay[day]; ok {
		solve(file)
	} else {
		fmt.Printf("Unrecognized or Not sovled day %v\n", day)
	}
}
