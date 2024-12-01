package day1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"aoc2024/pkg/log"
	"aoc2024/pkg/math"
)

func sortList(wg *sync.WaitGroup, l []int) {
	defer wg.Done()
	// Sort descending order of list
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
}

func ExtractSplitList(filename string) ([]int, []int, error) {
	var (
		left  []int
		right []int
	)
	file, err := os.Open(filename)
	if err != nil {
		return left, right, err
	}

	defer file.Close()

	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)

	for filescanner.Scan() {
		line := filescanner.Text()
		v := strings.Split(line, "   ")

		if len(v) != 2 {
			log.Logger.Errorw("Error failed to split line",
				log.Any("Split-list", v),
				log.String("line", line),
			)
			return left, right, errors.New("Failed to splits list by two")
		}

		// Convert string to int and append to left
		l, err := strconv.Atoi(v[0])
		if err != nil {
			log.Logger.Errorw("Error failed convert integer",
				log.String("Value", v[0]),
				log.Any("Split-list", v),
			)
			return left, right, err
		}
		left = append(left, l)

		// Convert string to int and append to left
		r, err := strconv.Atoi(v[1])
		if err != nil {
			log.Logger.Errorw("Error failed convert integer",
				log.String("Value", v[1]),
				log.Any("Split-list", v),
			)
			return left, right, err
		}
		right = append(right, r)
	}
	return left, right, nil
}

func totalDistance(left, right []int) int {
	sum := 0
	log.Logger.Debugw("Total Distance calculation begins",
		log.Int("sum", sum),
		log.Int("left-length", len(left)),
		log.Int("right-length", len(right)),
	)

	for i := 0; i < len(left); i++ {
		dist := math.Abs(left[i] - right[i])
		sum += dist

		log.Logger.Debugw("Add distance",
			log.Int("left", left[i]),
			log.Int("right", right[i]),
			log.Int("distance", dist),
			log.Int("sum", sum),
		)
	}
	return sum
}

func AdventSolveDay1(filename string) {
	left, right, err := ExtractSplitList(filename)

	if err != nil {
		log.Logger.Fatalw(err.Error(), log.String("filename", filename))
	}

	// Go routine sort both lists
	var wg sync.WaitGroup

	wg.Add(2)
	go sortList(&wg, left)
	go sortList(&wg, right)
	wg.Wait()

	total := totalDistance(left, right)
	fmt.Println("TOTAL:", total)
}
