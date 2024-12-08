package day6

import (
	"bufio"
	"fmt"
	"os"

	"aoc2024/pkg/log"
)

type Directions int

const (
	Obstruction = '#'
	Empty       = '.'
	Marked      = 'X'
	GuardUp     = '^'

	Up    Directions = 1
	Down  Directions = 2
	Right Directions = 3
	Left  Directions = 4
)

var directions = map[Directions]struct{ Dx, Dy int }{
	Up:    {0, -1}, // Up
	Down:  {0, 1},  // Down
	Right: {1, 0},  // Right
	Left:  {-1, 0}, // Left
}

type Lab [][]byte

func (l Lab) PrintLab(guard *Guard) {
	fmt.Printf("\n")
	for y := 0; y < len(l); y++ {
		for x := 0; x < len(l); x++ {
			if guard.Y == y && guard.X == x {
				switch guard.Dir {
				case Up:
					fmt.Printf("^")
				case Down:
					fmt.Printf("⌄")
				case Left:
					fmt.Printf("<")
				case Right:
					fmt.Printf(">")
				}
			} else {
				fmt.Printf("%s", string(l[y][x]))
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (l Lab) Count(char byte) int {
	count := 0
	for y := 0; y < len(l); y++ {
		for x := 0; x < len(l[0]); x++ {
			if l[y][x] == char {
				count++
			}
		}
	}
	return count
}

type Guard struct {
	Dir Directions
	X   int
	Y   int
}

// SimulateGuardPatrol simulate guard patrol in laboratory where goal is for guard to exit map
// Lab guard patrol protocol:
// - If there is something directly in front of you, turn right 90 degrees.
// - Otherwise, take a step forward.
// Returns will be the map of Lab with all position where guard have been marked by X and distinct positions count
func (guard *Guard) SimulateGuardPatrol(lab Lab) Lab {
	exitedMap := false
	for !exitedMap {
		dir := directions[guard.Dir]
		nx, ny := guard.X+dir.Dx, guard.Y+dir.Dy
		if ny < 0 || ny >= len(lab[0]) ||
			guard.Y < 0 || guard.Y >= len(lab) ||
			guard.X < 0 || guard.X >= len(lab[0]) ||
			nx < 0 || nx >= len(lab[0]) {
			exitedMap = true
			break
		}

		switch lab[ny][nx] {
		case Empty:
			// take a step forward
			guard.X = nx
			guard.Y = ny
			// Marked as walked
			lab[guard.Y][guard.X] = Marked
		case Marked:
			// take a step forward, we have already marked move which means previous Empty
			guard.X = nx
			guard.Y = ny
		case Obstruction:
			// Rotate 90%
			switch guard.Dir {
			case Up:
				guard.Dir = Right
			case Right:
				guard.Dir = Down
			case Down:
				guard.Dir = Left
			case Left:
				guard.Dir = Up
			}
			log.Debug("Guard encountered obstacle", log.Any("Roted", directions[guard.Dir]))
		}
		log.Debug("Guard new position",
			log.Int("Y", guard.Y),
			log.Int("X", guard.X),
			log.Any("Roted", directions[guard.Dir]),
			log.Int("ny", ny),
			log.Int("nx", nx),
		)
	}
	return lab
}

func extractLaboratory(filename string) (Lab, *Guard) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file", log.String("filename", filename), log.String("error", err.Error()))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var (
		lab   = Lab{}
		guard = &Guard{Dir: Up}
	)
	y := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())

		for x := 0; x < len(line); x++ {
			if line[x] == GuardUp {
				log.Debug("Guard Position found in map", log.Int("Y", y), log.Int("X", x))
				guard.X = x
				guard.Y = y
				line[x] = Marked
			}
		}
		lab = append(lab, line)
		y++
	}
	return lab, guard
}

func AdventSolveDay6(filename string) {
	log.Info("Start Part 1", log.String("filename", filename))
	lab, guard := extractLaboratory(filename)
	lab.PrintLab(guard)

	patrolMap := guard.SimulateGuardPatrol(lab)
	patrolMap.PrintLab(guard)

	distinctPositions := patrolMap.Count(Marked)
	fmt.Println("Distinct positions:", distinctPositions)

	log.Info("Done Part 1", log.String("filename", filename))
}
