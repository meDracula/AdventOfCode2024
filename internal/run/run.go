package run

import (
	"fmt"
)

type AdventSolveDay func(filename string)

var AdventOfDay = map[int]AdventSolveDay{
}

func Day(day int, file string) {
	if solve, ok := AdventOfDay[day]; ok {
		solve(file)
	} else {
		fmt.Printf("Unrecognized or Not sovled day %v\n", day)
	}
}
