package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc2024/pkg/log"
)

type Operator string

const (
	Addition       Operator = "+"
	Multiplication Operator = "*"
	Concatenation  Operator = "||"
)

func (o Operator) Eval(a, b int) int {
	switch o {
	case Addition:
		return a + b
	case Multiplication:
		return a * b
	case Concatenation:
		con, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		if err != nil {
			panic(err.Error())
		}
		return con
	}
	panic(fmt.Sprintf("Non supported operator %v", o))
	return 0
}

func CartesianProductOperators(operators []Operator, repeat int) [][]Operator {
	if repeat == 0 {
		return [][]Operator{}
	}

	if repeat == 1 {
		result := [][]Operator{}
		for _, op := range operators {
			result = append(result, []Operator{op})
		}
		return result
	}

	// Initialize result with single character combinations
	result := make([][]Operator, len(operators))
	for i, op := range operators {
		result[i] = []Operator{op}
	}

	// Iteratively build combinations of the desired repeat
	for i := 1; i < repeat; i++ {
		var newResult [][]Operator
		for _, combination := range result {
			for _, op := range operators {
				newCombination := append([]Operator{}, combination...)
				newCombination = append(newCombination, op)
				newResult = append(newResult, newCombination)
			}
		}
		result = newResult
	}

	return result
}

type CalibrationEquation struct {
	Test     int
	Equation []int
}

func (eq *CalibrationEquation) Evaluate(operators []Operator) bool {
	// Create the Cartesian product of the set of operators {+, *}
	product := CartesianProductOperators(operators, len(eq.Equation)-1)
	log.Debug("Cartesian Product Operators",
		log.Any("product", product),
		log.Any("Equation", eq.Equation),
	)

	for _, ops := range product {
		log.Debug("New Operators Equation",
			log.Any("Equation", eq.Equation),
			log.Any("operators", ops),
			log.Int("test", eq.Test),
		)
		// Evaluate equation
		total := eq.Equation[0]
		for i, op := range ops {
			res := op.Eval(total, eq.Equation[i+1])
			total = res
		}

		// Check if equation matches test
		if total == eq.Test {
			return true
		}
	}
	return false
}

func calibrationEquationsPatcher(calibrationEquations []CalibrationEquation, operators []Operator) int {
	total := 0
	for _, calibrationEquation := range calibrationEquations {
		if calibrationEquation.Evaluate(operators) {
			total += calibrationEquation.Test
		}
	}
	return total
}

func extractCalibrationEquations(filename string) []CalibrationEquation {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", log.String("filename", filename), log.String("error", err.Error()))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	calibrationEquations := []CalibrationEquation{}
	for scanner.Scan() {
		line := scanner.Text()

		// Cut out test digit
		sepIndex := strings.Index(line, ":")
		test, err := strconv.Atoi(line[:sepIndex])
		if err != nil {
			file.Close()
			log.Fatal("Failed to retrieve and covert test to int", //nolint:gocritic // manual operation set to avoid exitAfterDefer
				log.String("line", line),
				log.String("test", line[:sepIndex]),
				log.String("error", err.Error()),
			)
		}

		equationNumbers := strings.Split(line[sepIndex+2:], " ")
		log.Debug("Equation Numbers", log.Any("Equation", equationNumbers),
			log.String("line", line), log.Int("sepIndex", sepIndex),
		)
		equation := []int{}
		for _, num := range equationNumbers {
			number, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal("Failed to retrieve and covert test to int",
					log.String("line", line),
					log.String("number", num),
					log.String("error", err.Error()),
				)
			}
			equation = append(equation, number)
		}

		calibrationEquations = append(calibrationEquations, CalibrationEquation{
			Test:     test,
			Equation: equation,
		})
	}

	return calibrationEquations
}

func AdventSolveDay7(filename string) {
	log.Info("Start Part 1", log.String("filename", filename))
	calibrationEquations := extractCalibrationEquations(filename)
	calibrationTotal := calibrationEquationsPatcher(calibrationEquations, []Operator{Addition, Multiplication})
	fmt.Println("Calibration Equations Total:", calibrationTotal)
	log.Info("Done Part 1", log.String("filename", filename))

	log.Info("Start Part 2", log.String("filename", filename))
	calibrationTotal = calibrationEquationsPatcher(calibrationEquations, []Operator{Addition, Multiplication, Concatenation})
	fmt.Println("Calibration Equations Total:", calibrationTotal)
	log.Info("Done Part 2", log.String("filename", filename))
}
