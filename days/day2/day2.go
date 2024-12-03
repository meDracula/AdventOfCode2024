package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc2024/pkg/log"
)

func extractReports(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file",
			log.String("filename", filename),
			log.String("error", err.Error()),
		)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)
	lines.Split(bufio.ScanLines)

	reports := [][]int{}
	for lines.Scan() {
		line := lines.Text()
		levels := strings.Split(line, " ")

		report := []int{}
		for _, level := range levels {
			v, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal("Cannot convert level to int",
					log.String("level", level),
					log.String("error", err.Error()),
					log.Any("levels", levels),
				)
			}
			report = append(report, v)
		}
		reports = append(reports, report)
	}
	return reports
}

type ReportSafetySystemFunc func(report []int) bool

func safetyLevelcheck(report []int, ruleBreak func(i, j int) bool) bool {
	for i := 1; i < len(report); i++ {
		if ruleBreak(report[i-1], report[i]) {
			return false
		}
	}
	return true
}

func reportSafetySystemCheck(report []int) bool {
	// Either level increasing nor decreasing
	if report[0] == report[1] {
		return false
	}

	// Rule function for decreasing order
	ruleFunc := func(i, j int) bool {
		return i <= j || (i-j) > 3
	}

	// Determine if increasing
	if report[0] < report[1] {
		// change rule function to increasing rule
		ruleFunc = func(i, j int) bool {
			return i >= j || (j-i) > 3
		}
	}
	return safetyLevelcheck(report, ruleFunc)
}

func countSafeReports(filename string, safetySystemFunc ReportSafetySystemFunc) int {
	reports := extractReports(filename)
	safeCountch := make(chan int, len(reports))

	for _, report := range reports {
		go func(report []int, safetySystemFunc ReportSafetySystemFunc, safeCountch chan<- int) {
			if safetySystemFunc(report) {
				log.Debug("Level Safe", log.Any("report", report))
				safeCountch <- 1
			} else {
				log.Debug("Report Unsafe Equal", log.Any("report", report))
				safeCountch <- 0
			}
		}(report, safetySystemFunc, safeCountch)
	}

	safeCount := 0
	for s := 0; s < len(reports); s++ {
		safeCount += <-safeCountch
	}

	return safeCount
}

func countSafeReportsWithDampener(filename string, safetySystemFunc ReportSafetySystemFunc) int {
	reports := extractReports(filename)
	safeCountch := make(chan int, len(reports))

	for _, report := range reports {
		go func(report []int, safetySystemFunc ReportSafetySystemFunc, safeCountch chan<- int) {
			// Generate report permutation of report
			for i := range report {
				subReport := make([]int, 0, len(report)-1)
				subReport = append(subReport, report[:i]...)
				subReport = append(subReport, report[i+1:]...)

				if safetySystemFunc(subReport) {
					safeCountch <- 1
					return
				}
			}
			safeCountch <- 0 // All permutation of report are unsafe
		}(report, safetySystemFunc, safeCountch)
	}

	safeCount := 0
	for s := 0; s < len(reports); s++ {
		safeCount += <-safeCountch
	}

	return safeCount
}

func AdventSolveDay2(filename string) {
	log.Info("Start Part 1", log.String("filename", filename))
	safe := countSafeReports(filename, reportSafetySystemCheck)
	fmt.Println("Safe Part 1:", safe)
	log.Info("Part 1 Done", log.String("filename", filename), log.Int("Safe", safe))

	log.Info("Start Part 2", log.String("filename", filename))
	safeWithDampeners := countSafeReportsWithDampener(filename, reportSafetySystemCheck)
	fmt.Println("Safe with Dampeners:", safeWithDampeners)
	log.Info("Part 2 Done", log.String("filename", filename), log.Int("Safe", safeWithDampeners))
}
