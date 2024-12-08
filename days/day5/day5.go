package day5

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"aoc2024/pkg/log"
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

func (r Rules) validUpdate(update Update) bool {
	page := update.page
	for i := 0; i < (len(page) - 1); i++ {
		log.Debug("Check valid page", log.Any("page", page[i:]), log.Int("index", i))

		// Check rules loop
		// Code logic edge case: if len of page is 2 then check rules loop condition will be false i+1 < len(page[i:])
		for j := i + 1; j < len(page); j++ {
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

		if r.validUpdate(update) {
			log.Debug("update is valid", log.Any("page", update.page))
			validUpates = append(validUpates, update)
		}
	}
	return validUpates
}

func (r Rules) patchPage(page []string) bool {
	for i := 0; i < (len(page) - 1); i++ {
		log.Debug("Check if patch required", log.Any("page", page[i:]), log.Int("index", i))
		// Check rules loop
		// Code logic edge case: if len of page is 2 then check rules loop condition will be false i+1 < len(page[i:])
		for j := i + 1; j < len(page); j++ {
			key := page[i] + "|" + page[j]

			log.Debug("Check key patch required",
				log.Any("page[i]", page[i]),
				log.Any("page[j]", page[j]),
				log.String("key", key),
				log.Int("index", i),
			)
			if _, ok := r[key]; !ok {
				// Swap key value
				page[i], page[j] = page[j], page[i]
				log.Debug("Swap keyed", log.String("key", key), log.Any("Page", page))
				return false
			}
		}
	}
	return true
}

func (r Rules) fixUpdate(update *Update) {
	var (
		page    = update.page
		patched = false
	)
	for !patched {
		log.Debug("Run update patch page", log.Any("page", page))
		if r.patchPage(page) {
			update.page = page
			log.Debug("Patched complete", log.Any("Page", page))
			patched = true
		}
	}
	return
}

func (r Rules) fixIncorrectlyUpdates(updates []Update) []Update {
	for _, update := range updates {
		r.fixUpdate(&update)
	}
	return updates
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

func FilterUpdate(updates []Update, remove []Update) []Update {
	filtered := []Update{}
	for _, update := range updates {
		add := true
		for _, filter := range remove {
			if reflect.DeepEqual(update, filter) {
				add = false
			}
		}
		if add {
			filtered = append(filtered, update)
		}
	}
	return filtered
}

func AdventSolveDay5(filename string) {
	log.Info("Start Part 1", log.String("filename", filename))

	rules, updates := extractUpdateManual(filename)
	updateOrdering := rules.updateOrdering(updates)
	sum := sumUpdates(updateOrdering)

	fmt.Println("correctly-ordered updates sum:", sum)
	log.Info("Done Part 1", log.String("filename", filename))

	incorrectlyUpdates := FilterUpdate(updates, updateOrdering)

	log.Info("Start Part 2", log.String("filename", filename))

	fixedUpdateOrdering := rules.fixIncorrectlyUpdates(incorrectlyUpdates)
	sum = sumUpdates(fixedUpdateOrdering)

	fmt.Println("sorted correctly-ordered updates sum:", sum)
	log.Info("Done Part 2", log.String("filename", filename))
}
