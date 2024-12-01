package flags

import "flag"

type FlagOutputOpts struct {
	Day int
	File string
	Debug bool
}

func Parse() (opts *FlagOutputOpts) {
	flag.IntVar(opts.Day, "day", 0, "Select day to run")
	flag.StringVar(opts.File, "file", "", "Path to solve puzzle input")
	flag.BoolVar(opts.Debug, "debug", false, "log debug")

	flag.Parse()
	return opts
}
