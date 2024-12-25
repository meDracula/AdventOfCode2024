package day8

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"aoc2024/pkg/log"
)

const dot = '.'

type (
	MapBoarder struct {
		X, Y int
	}
	Node struct {
		X, Y int
	}
	FrequencyNodeMap map[rune][]Node
	AntiNodeMap      map[Node]bool
)

func (a AntiNodeMap) Unqiue() int {
	return len(a)
}

func (a AntiNodeMap) Print() {
	for node, _ := range a {
		fmt.Printf("{X: %d, Y: %d}\n", node.X, node.Y)
	}
}

func (f FrequencyNodeMap) Print() {
	for key, nodes := range f {
		builder := strings.Builder{}

		builder.WriteString(string(key))
		builder.WriteString(": { ")
		for i := 0; i < len(nodes); i++ {
			builder.WriteString(fmt.Sprintf("{X: %d, Y: %d}", nodes[i].X, nodes[i].Y))
			if i != (len(nodes) - 1) {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(" }")
		fmt.Println(builder.String())
	}
}

// Resonant Collinearity.
// Each antenna is tuned to a specific frequency
// indicated by a single lowercase letter, uppercase letter, or digit.
func ResonantCollinearity(frequencyNodes FrequencyNodeMap, mapBoarder MapBoarder) AntiNodeMap {
	antiNodes := AntiNodeMap{}

	for freq, nodes := range frequencyNodes {
		log.Debug("Comparing new antenna frequency",
			log.String("Frequency", string(freq)),
			log.Any("Nodes", nodes),
		)
		// Special case: lone antenna no antinodes will be created.
		if len(nodes) <= 1 {
			log.Debug("lone antenna therefore no antinodes to add!",
				log.String("Frequency", string(freq)),
				log.Any("Nodes", nodes),
			)
			continue
		}
		for selectedNodeIndex, selectedNode := range nodes {
			// Selected node compare with all other nodes
			for i := 0; i < len(nodes); i++ {
				// Skip over the selected node index
				if selectedNodeIndex == i {
					continue
				}
				// Diff nodes
				dx := nodes[i].X - selectedNode.X
				dy := nodes[i].Y - selectedNode.Y
				n := Node{X: nodes[i].X + dx, Y: nodes[i].Y + dy}

				log.Debug("Diff nodes",
					log.String("Frequency", string(freq)),
					log.Any("Selected Node", selectedNode),
					log.Any("Diff Node", n),
					log.Int("Diff.X", dx),
					log.Int("Diff.Y", dy),
				)

				// out bounds of the map
				if (n.X < 0 || n.X >= mapBoarder.X) || (n.Y < 0 || n.Y >= mapBoarder.Y) {
					log.Debug("AntiNode out-bound map", log.Any("node", n), log.Any("Boarder", mapBoarder))
					continue // Skip adding to antiNodes
				}

				// Make or re-add antinode
				antiNodes[n] = true
			}
		}
	}

	return antiNodes
}

func extractFile(filename string) (FrequencyNodeMap, MapBoarder) {
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

	frequencyNodes := FrequencyNodeMap{}
	xLength := 0
	y := 0
	for lines.Scan() {
		line := lines.Text()
		xLength = len(line)
		for x := 0; x < len(line); x++ {
			if line[x] != dot {
				freq := rune(line[x])
				frequencyNodes[freq] = append(frequencyNodes[freq], Node{X: x, Y: y})
			}
		}
		y++
	}
	return frequencyNodes, MapBoarder{X: xLength, Y: y}
}

func AdventSolveDay8(filename string) {
	log.Info("Start Part 1", log.String("filename", filename))

	frequencyNodes, mapBoarder := extractFile(filename)
	frequencyNodes.Print()

	antiNodes := ResonantCollinearity(frequencyNodes, mapBoarder)
	antiNodes.Print()
	locations := antiNodes.Unqiue()
	fmt.Println("total unique locations:", locations)

	log.Info("Done Part 1", log.String("filename", filename), log.Int("locations", locations))
}
