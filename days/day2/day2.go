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
		log.Logger.Fatalw("Failed to open file",
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
				log.Logger.Fatalw("Cannot convert level to int",
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

func reportSafetySystemCheck(report []int) bool {
	// Either level increasing nor decreasing
	if report[0] == report[1] {
		log.Logger.Debugw("Level Unsafe Equal", log.Any("report", report))
		return false
	}

	// Determine if increasing
	if report[0] < report[1] {
		// Level is increasing
		for i := 1; i < len(report); i++ {
			if report[i-1] >= report[i] || (report[i]-report[i-1]) > 3 {
				log.Logger.Debugw("Level Unsafe Increasing",
					log.Any("report", report),
					log.Int("i-1", report[i-1]),
					log.Int("i", report[i]),
				)
				return false
			}
		}
	} else {
		// Level is decreasing
		for i := 1; i < len(report); i++ {
			// not decreasing or more than three increment
			if report[i-1] <= report[i] || (report[i-1]-report[i]) > 3 {
				log.Logger.Debugw("Level Unsafe Decreasing",
					log.Any("report", report),
					log.Int("i-1", report[i-1]),
					log.Int("i", report[i]),
				)
				return false
			}
		}
	}
	log.Logger.Debugw("Level Safe", log.Any("report", report))
	return true
}

func countSafeReports(filename string) int {
	reports := extractReports(filename)

	safeCount := 0
	for _, report := range reports {
		if reportSafetySystemCheck(report) {
			safeCount++
		}
	}

	return safeCount
}

func AdventSolveDay2(filename string) {
	safe := countSafeReports(filename)
	fmt.Println("Safe:", safe)
}
