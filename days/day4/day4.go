package day4

import (
	"fmt"

	"aoc2024/pkg/log"
	"aoc2024/pkg/reader"
)

const XMAS = "XMAS"

type Xmas struct {
	text    []string
	yLength int
	xLength int
}

func New(text []string) *Xmas {
	return &Xmas{
		text:    text,
		yLength: len(text),
		xLength: len(text[0]),
	}
}

// searchForAllXMAS count for a word search of XMAS
// allows words to be horizontal, vertical, diagonal, backwards, or even overlapping other words
func (xmas *Xmas) searchForAllXMAS() int {
	directions := []struct {
		dx, dy int
	}{
		{0, 1},   // Horizontal right
		{1, 0},   // Vertical down
		{0, -1},  // Horizontal left
		{-1, 0},  // Vertical up
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{1, -1},  // Diagonal down-left
		{-1, 1},  // Diagonal up-right
	}

	count := 0

	// Helper function to check if the word can be found starting at (x, y) in the given direction
	isXMAS := func(x, y, dx, dy int) bool {
		for i := 0; i < len(XMAS); i++ {
			nx, ny := x+i*dx, y+i*dy
			if nx < 0 || ny < 0 || nx >= xmas.yLength || ny >= xmas.xLength || xmas.text[nx][ny] != XMAS[i] {
				return false
			}
		}
		return true
	}

	// Iterate through every cell in the xmas.text
	for y := 0; y < xmas.yLength; y++ {
		for x := 0; x < xmas.xLength; x++ {
			// Check in all 8 directions
			for _, dir := range directions {
				if isXMAS(x, y, dir.dx, dir.dy) {
					count++
					log.Debug("Found XMAS",
						log.Int("x", x),
						log.Int("y", y),
						log.Int("direction.x", dir.dx),
						log.Int("direction.y", dir.dy),
						log.Int("count", count),
					)
				}
			}
		}
	}

	return count
}

// TODO Work in progress
func (xmas *Xmas) searchAllXXMAS() int {
	grid := []struct {
		dx, dy int
	}{
		{1, 1},   // down-right
		{-1, -1}, // up-left
		{1, -1},  // down-left
		{-1, 1},  // up-right
	}

	isXXMAS := func(x, y int) bool {
		for _, dir := range grid {
			nx, ny := x+dir.dx, y+dir.dy
			log.Debug("Character",
				log.String("Rune", string(xmas.text[nx][ny])),
				log.Int("nx", nx),
				log.Int("ny", ny),
			)
			if xmas.text[nx][ny] != 'S' && xmas.text[nx][ny] != 'M' {
				return false
			}
		}

		return false
	}

	count := 0

	// skip line zero 0 and final line no match possible
	for y := 1; y < (xmas.yLength - 1); y++ {
		for x := 1; x < xmas.xLength-1; x++ {
			// Skip iteration if no A
			if xmas.text[y][x] != 'A' {
				continue
			}
			log.Debug("Found A",
				log.Int("x", x),
				log.Int("y", y),
			)

			if isXXMAS(x, y) {
				count++
				log.Debug("Found XMAS",
					log.Int("x", x),
					log.Int("y", y),
					log.Int("count", count),
				)
			}
		}
	}

	return count
}

func AdventSolveDay4(filename string) {
	text, err := reader.FileReadlines(filename)
	if err != nil {
		log.Fatal("Failed to read file",
			log.String("error", err.Error()),
			log.String("filename", filename),
		)
	}

	xmas := New(text)

	log.Info("Start Part 1", log.String("filename", filename))
	foundAllXMAS := xmas.searchForAllXMAS()
	fmt.Println("Found XMAS:", foundAllXMAS)
	log.Info("Done Part 1", log.String("filename", filename))

	// TODO Work in progress
	log.Info("Start Part 2", log.String("filename", filename))
	foundAllXXMAS := xmas.searchAllXXMAS()
	fmt.Println("Found X-MAS:", foundAllXXMAS)
	log.Info("Done Part 2", log.String("filename", filename))
}
