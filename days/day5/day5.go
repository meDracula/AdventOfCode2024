package day5

import (
	"aoc2024/pkg/log"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rules map[string]bool

type Update struct {
	page []string
}

const Two = 2

func extractUpdateManual(filename string) (Rules, []Update) {
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
		rules   = Rules{}
		updates = []Update{}
	)

	// extract page ordering rules
	for lines.Scan() {
		rule := lines.Text()
		if rule == "" {
			break
		}
		rules[rule] = true
	}

	// extract pages to produce in each update
	for lines.Scan() {
		update := lines.Text()
		u := strings.Split(update, ",")

		updates = append(updates, Update{page: u})
	}
	return rules, updates
}

func (r Rules) validPage(page []string) bool {
	for i := 0; i < (len(page) - 1); i++ {
		log.Debug("Check valid page", log.Any("page", page[i:]), log.Int("index", i))

		// Code logic edge case: if len of page is 2 then check rules loop condition will be false i+1 < len(page[i:])
		if len(page[i:]) == Two {
			key := page[i] + "|" + page[i+1]
			if _, ok := r[key]; !ok {
				log.Debug("Rule doesn't exist", log.String("key", key))
				return false
			}
		}

		// Check rules loop
		for j := i + 1; j < len(page[i:]); j++ {
			key := page[i] + "|" + page[j]
			if _, ok := r[key]; !ok {
				log.Debug("Rule doesn't exist", log.String("key", key))
				return false
			}
		}
	}
	log.Debug("Page exist", log.Any("page", page))
	return true
}

func (r Rules) updateOrdering(updates []Update) []Update {
	validUpates := []Update{}
	for _, update := range updates {
		log.Debug("Check valid Update", log.Any("page", update.page))

		if r.validPage(update.page) {
			log.Debug("update is valid", log.Any("page", update.page))
			validUpates = append(validUpates, update)
		}
	}
	return validUpates
}

func sumUpdates(updates []Update) int {
	sum := 0
	for i := 0; i < len(updates); i++ {
		middleIndex := len(updates[i].page) / Two
		middle := updates[i].page[middleIndex]

		num, err := strconv.Atoi(middle)
		if err != nil {
			log.Fatal("Failed to convert middle page number to int",
				log.String("error", err.Error()),
				log.String("middle", middle),
				log.Any("page", updates[i].page),
			)
		}
		sum += num
	}

	return sum
}

func AdventSolveDay5(filename string) {
	log.Info("Start Part 1", log.String("filename", filename))

	rules, updates := extractUpdateManual(filename)
	updateOrdering := rules.updateOrdering(updates)
	sum := sumUpdates(updateOrdering)

	fmt.Println("correctly-ordered updates sum:", sum)
	log.Info("Done Part 1", log.String("filename", filename))
}
