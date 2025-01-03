package run

import (
	"fmt"

	"aoc2024/days/day1"
	"aoc2024/days/day2"
	"aoc2024/days/day3"
	"aoc2024/days/day4"
	"aoc2024/days/day5"
	"aoc2024/days/day6"
	"aoc2024/days/day7"
	"aoc2024/days/day8"
)

type AdventSolveDay func(filename string)

var AdventOfDay = map[int]AdventSolveDay{
	1: day1.AdventSolveDay1,
	2: day2.AdventSolveDay2,
	3: day3.AdventSolveDay3,
	4: day4.AdventSolveDay4,
	5: day5.AdventSolveDay5,
	6: day6.AdventSolveDay6,
	7: day7.AdventSolveDay7,
	8: day8.AdventSolveDay8,
}

func Day(day int, file string) {
	if solve, ok := AdventOfDay[day]; ok {
		solve(file)
	} else {
		fmt.Printf("Unrecognized or Not sovled day %v\n", day)
	}
}
