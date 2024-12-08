package day6

import (
	"bufio"
	"fmt"
	"os"

	"aoc2024/pkg/log"
)

type Directions int

const (
	Obstruction     = '#'
	ObstructionLoop = 'O'
	Empty           = '.'
	Marked          = 'X'
	GuardUp         = '^'

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

func (l Lab) Copy() Lab {
	copiedLab := make(Lab, len(l))

	for i := range l {
		copiedLab[i] = make([]byte, len(l[i]))
		copy(copiedLab[i], l[i])
	}
	return copiedLab
}

type Guard struct {
	Dir Directions
	X   int
	Y   int
}

// Rotate 90%
func (guard *Guard) Rotate() {
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
}

// Step take a step forward
func (guard *Guard) Step(nx, ny int) {
	guard.X = nx
	guard.Y = ny
}

func (guard *Guard) Copy() *Guard {
	return &Guard{
		Dir: guard.Dir,
		Y:   guard.Y,
		X:   guard.X,
	}
}

// SimulateGuardPatrol simulate guard patrol in laboratory where goal is for guard to exit map.
// Lab guard patrol protocol:
// - If there is something directly in front of you, turn right 90 degrees.
// - Otherwise, take a step forward.
// Returns will be the map of Lab with all position where guard have been marked by X and distinct positions count.
// If guard get stuck in a loop the final boolean variable will be true.
func (guard *Guard) SimulateGuardPatrol(lab Lab) (Lab, bool) {

	var (
		patrolMap = lab.Copy()
		visited   = map[Guard]bool{}
		looped    = false
	)
sim:
	for {
		dir := directions[guard.Dir]
		nx, ny := guard.X+dir.Dx, guard.Y+dir.Dy
		if ny < 0 || ny >= len(patrolMap[0]) ||
			guard.Y < 0 || guard.Y >= len(patrolMap) ||
			guard.X < 0 || guard.X >= len(patrolMap[0]) ||
			nx < 0 || nx >= len(patrolMap[0]) {
			break sim
		}

		switch patrolMap[ny][nx] {
		case Empty:
			// take a step forward
			guard.Step(nx, ny)
			// Marked as walked
			patrolMap[guard.Y][guard.X] = Marked
		case Marked:
			// take a step forward, we have already marked move which means previous Empty
			guard.Step(nx, ny)
		case Obstruction:
			// Rotate 90%
			guard.Rotate()
			log.Debug("Guard encountered obstacle", log.Any("Rotated", directions[guard.Dir]))
		case ObstructionLoop:
			log.Debug("Obstruction O encountered",
				log.Int("Y", guard.Y),
				log.Int("X", guard.X),
				log.Any("Rotated", directions[guard.Dir]),
			)
			// Rotate 90%
			guard.Rotate()
		}
		if _, ok := visited[*guard]; ok {
			looped = true
			break sim
		}
		// Add guard to visited position and direction
		visited[*guard] = true
		log.Debug("Guard new position",
			log.Int("Y", guard.Y),
			log.Int("X", guard.X),
			log.Any("Rotated", directions[guard.Dir]),
			log.Int("ny", ny),
			log.Int("nx", nx),
		)
	}
	return patrolMap, looped
}

func GuardLoopSimulation(guard *Guard, lab, patrolMap Lab) int {
	guardLoopedCount := 0

	// For every X mark try simlutate with a O marker to check if loop, expcept for guard current position.
	for y := 0; y < len(patrolMap); y++ {
		for x := 0; x < len(patrolMap); x++ {
			if guard.Y == y && guard.X == x {
				log.Debug("Guard Position encoutered",
					log.Int("Y", y),
					log.Int("Guard.Y", guard.Y),
					log.Int("X", x),
					log.Int("Guard.X", guard.X),
				)
				continue
			}
			if patrolMap[y][x] == Marked {
				// Make new map with ObstructionLoop 'O'
				simLab := lab.Copy()
				simLab[y][x] = ObstructionLoop
				// Copy guard
				simGuard := guard.Copy()

				log.Info("New Simulation of Guard patrol",
					log.Any("Guard", *guard),
					log.Int("O.y", y),
					log.Int("O.x", x),
				)
				simPatrolMap, looped := simGuard.SimulateGuardPatrol(simLab)
				simPatrolMap.PrintLab(simGuard)
				if looped {
					guardLoopedCount++
					log.Info("Successfully looped guard",
						log.Any("Guard", *guard),
						log.Int("O.y", y),
						log.Int("O.x", x),
						log.Int("Looped", guardLoopedCount),
					)
				} else {
					log.Info("Not successfully looped guard",
						log.Any("Guard", *guard),
						log.Int("O.y", y),
						log.Int("O.x", x),
						log.Int("Looped", guardLoopedCount),
					)
				}
			}
		}
	}

	return guardLoopedCount
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
	log.Info("Day 6 Extract Laboratory", log.String("filename", filename))
	lab, guard := extractLaboratory(filename)
	lab.PrintLab(guard)

	startGuard := &Guard{Dir: guard.Dir, X: guard.X, Y: guard.Y}

	log.Info("Start Part 1", log.String("filename", filename))
	patrolMap, _ := guard.SimulateGuardPatrol(lab)
	patrolMap.PrintLab(guard)

	distinctPositions := patrolMap.Count(Marked)

	fmt.Println("Distinct positions:", distinctPositions)
	log.Info("Done Part 1", log.String("filename", filename))

	log.Info("Start Part 2", log.String("filename", filename))
	guardLoopedCount := GuardLoopSimulation(startGuard, lab, patrolMap)

	fmt.Println("Guard Stuck in Looped:", guardLoopedCount)
	log.Info("Done Part 2", log.String("filename", filename))
}
