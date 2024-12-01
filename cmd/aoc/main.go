package main

import (
	"aoc2024/internal/run"
	"aoc2024/pkg/flags"
	"aoc2024/pkg/log"
)

func main() {
	opts := flags.Parse()

	// Setting debug level
	if opts.Debug {
		log.InitializeLogger(log.WithLevel(log.DebugLevel))
	}

	run.Day(opts.Day, opts.File)
}
