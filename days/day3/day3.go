package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"aoc2024/pkg/log"
)

func parseMul(mul string) int {
	// strip mul( and final )
	mul = mul[4 : len(mul)-1]
	digits := strings.Split(mul, ",")

	if len(digits) != 2 {
		log.Error("Parsing failed", log.String("mul-raw", mul), log.Any("digits", digits))
	}

	product := 1
	for _, d := range digits {
		factor, err := strconv.Atoi(d)
		if err != nil {
			log.Error("Parse",
				log.String("error", err.Error()),
				log.String("mul-raw", mul),
				log.Any("digits", digits),
			)
		}
		product *= factor
	}
	return product
}

func decorruptMemory(filename string) int {
	expr, err := regexp.Compile("mul\\(\\d+,\\d+\\)")
	if err != nil {
		log.Fatal("Failed to compile regex expression",
			log.String("filename", filename),
			log.String("error", err.Error()),
		)
	}

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

	sum := 0
	for lines.Scan() {
		memory := lines.Text()

		muls := expr.FindAllString(memory, -1)
		log.Debug("regex found mul", log.Any("muls-raw", muls), log.String("memory", memory))

		for _, m := range muls {
			s := parseMul(m)
			sum += s
			log.Debug("Multiplication added", log.String("mul-raw", m), log.Int("mul", s), log.Int("sum", s))
		}
	}

	return sum
}

func decorruptMemoryOperations(filename string) int {
	expr, err := regexp.Compile("mul\\(\\d+,\\d+\\)|don\\'t\\(\\)|do\\(\\)")
	if err != nil {
		log.Fatal("Failed to compile regex expression",
			log.String("filename", filename),
			log.String("error", err.Error()),
		)
	}

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

	var (
		sum = 0
		do  = true
	)
	for lines.Scan() {
		memory := lines.Text()

		operations := expr.FindAllString(memory, -1)
		log.Debug("regex found mul", log.Any("operations", operations), log.String("memory", memory))

		for _, op := range operations {
			if op == "do()" {
				do = true
				log.Debug("do() toggle on",
					log.String("operation", op),
					log.Bool("do", do),
				)
			} else if op == "don't()" {
				do = false
				log.Debug("don't() toggle off",
					log.String("operation", op),
					log.Bool("do", do),
				)
			} else {
				log.Debug("operation multiplication",
					log.String("operation", op),
					log.Bool("do", do),
				)
				if do {
					s := parseMul(op)
					sum += s
					log.Debug("multiplication added",
						log.String("operation", op),
						log.Int("mul", s),
						log.Int("sum", s),
						log.Bool("do", do),
					)
				}
			}
		}
	}

	return sum
}

func AdventSolveDay3(filename string) {
	log.Info("Start Part 1", log.String("filename", filename))
	sum := decorruptMemory(filename)
	fmt.Println("Part 1 multiply sum:", sum)
	log.Info("Done Part 1", log.String("filename", filename), log.Int("multiply-sum", sum))

	log.Info("Start Part 2", log.String("filename", filename))
	sum = decorruptMemoryOperations(filename)
	fmt.Println("Part 2 multiply sum:", sum)
	log.Info("Done Part 2", log.String("filename", filename), log.Int("multiply-sum", sum))
}
