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
			log.Error("Error failed to split line",
				log.Any("Split-list", v),
				log.String("line", line),
			)
			return left, right, errors.New("Failed to splits list by two")
		}

		// Convert string to int and append to left
		l, err := strconv.Atoi(v[0])
		if err != nil {
			log.Error("Error failed convert integer",
				log.String("Value", v[0]),
				log.Any("Split-list", v),
			)
			return left, right, err
		}
		left = append(left, l)

		// Convert string to int and append to left
		r, err := strconv.Atoi(v[1])
		if err != nil {
			log.Error("Error failed convert integer",
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
	log.Debug("Total Distance calculation begins",
		log.Int("sum", sum),
		log.Int("left-length", len(left)),
		log.Int("right-length", len(right)),
	)

	for i := 0; i < len(left); i++ {
		dist := math.Abs(left[i] - right[i])
		sum += dist

		log.Debug("Add distance",
			log.Int("left", left[i]),
			log.Int("right", right[i]),
			log.Int("distance", dist),
			log.Int("sum", sum),
		)
	}
	return sum
}

type frequencyMemorization struct {
	memo map[int]int
	List []int
}

func (f *frequencyMemorization) Count(v int) int {
	// Check memorization map for v
	if freq, ok := f.memo[v]; ok {
		return freq
	}
	// Create new index of v to memorization map
	f.memo[v] = 0
	for _, num := range f.List {
		if num == v {
			f.memo[v] = f.memo[v] + 1
		}
	}
	return f.memo[v]
}

func similarityScore(left, right []int) int {
	var (
		score     int                   = 0
		frequency frequencyMemorization = frequencyMemorization{List: right, memo: map[int]int{}}
	)

	log.Debug("similarity score calculation begins",
		log.Int("left-length", len(left)),
		log.Int("right-length", len(right)),
	)

	for i := 0; i < len(left); i++ {
		v := left[i]
		freq := frequency.Count(v)
		score += v * freq

		log.Debug("Add Score",
			log.Int("value", v),
			log.Int("frequency", freq),
			log.Int("score", score),
		)
	}
	return score
}

func AdventSolveDay1(filename string) {
	left, right, err := ExtractSplitList(filename)

	if err != nil {
		log.Fatal(err.Error(), log.String("filename", filename))
	}

	// Go routine sort both lists
	var wg sync.WaitGroup

	wg.Add(2)
	go sortList(&wg, left)
	go sortList(&wg, right)
	wg.Wait()

	log.Info("Start Part 1", log.String("filename", filename))
	total := totalDistance(left, right)
	fmt.Println("Part 1 Total Distance:", total)
	log.Info("Part 1 Done", log.String("filename", filename), log.Int("Total", total))

	log.Info("Start Part 2", log.String("filename", filename))
	score := similarityScore(left, right)
	fmt.Println("Part 2 Similarity Score:", score)
	log.Info("Part 2 Done", log.String("filename", filename), log.Int("Score", score))
}
