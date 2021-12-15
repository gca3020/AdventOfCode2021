package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Loc defines a location in the X/Y grid for this particular cell
type Loc struct {
	x int
	y int
}

// Cell defines the metadata associated with a cell, including its risk and the previous best path, if any
type Cell struct {
	loc      Loc
	risk     int
	visited  bool
	pathRisk int
}

type CellMap map[Loc]*Cell

func main() {
	text := input2slice("day15/input")

	dy := len(text)
	dx := len(text[0])
	cells := make(CellMap, dx*dy)
	bigCells := make(CellMap, dx*dy*25)
	for y, line := range text {
		for x, c := range line {
			l := Loc{x, y}
			r, _ := strconv.Atoi(string(c))
			cells[l] = &Cell{loc: l, risk: r, visited: false, pathRisk: math.MaxInt}
			for y1 := 0; y1 < 5; y1++ {
				for x1 := 0; x1 < 5; x1++ {
					risk := r + x1 + y1
					if risk > 9 {
						risk -= 9
					}
					loc := Loc{x + (dx * x1), y + (dy * y1)}
					cell := &Cell{loc: loc, risk: risk, visited: false, pathRisk: math.MaxInt}
					if cell == nil {
						fmt.Println("Could not allocate a cell")
						os.Exit(1)
					}
					bigCells[loc] = cell
				}
			}
		}
	}

	fmt.Println("The lowest risk in the small map is", cells.getLowestRisk(dx, dy))
	fmt.Println("The lowest risk in the big map is", bigCells.getLowestRisk(dx*5, dy*5))
}

func (cm CellMap) getLowestRisk(dx, dy int) int {
	start := cm[Loc{0, 0}]
	end := cm[Loc{dx - 1, dy - 1}]
	start.pathRisk = 0

	count := 0
	onepct := len(cm) / 100
	for c := start; c != nil && c != end; c = cm.getNextUnvisited() {
		c.visited = true
		count++
		if count%onepct == 0 {
			fmt.Printf("Processed %d total entries: %f percent\n", count, float32(count)/float32(len(cm))*100.0)
		}

		// Compute the pathRisk for all the unvisited neighbors, storing the new value if it's better than before
		neighbors := cm.getUnvisitedNeighbors(c)
		for _, n := range neighbors {
			pathRisk := c.pathRisk + n.risk
			if pathRisk < n.pathRisk {
				n.pathRisk = pathRisk
			}
		}
	}

	return end.pathRisk
}

// getNextUnvisited will return the unvisited cell with the lowest total path risk
func (cm CellMap) getNextUnvisited() *Cell {
	var next *Cell = nil
	for _, v := range cm {
		if !v.visited && (next == nil || v.pathRisk < next.pathRisk) {
			next = v
		}
	}
	return next
}

// getUnvisitedNeighbors will return a slice containing all the neighbors of a given cell that haven't been visited
func (cm CellMap) getUnvisitedNeighbors(cell *Cell) []*Cell {
	locs := []Loc{
		{cell.loc.x, cell.loc.y - 1}, // Up
		{cell.loc.x, cell.loc.y + 1}, // Down
		{cell.loc.x - 1, cell.loc.y}, // Left
		{cell.loc.x + 1, cell.loc.y}, // Right
	}

	n := make([]*Cell, 0)
	for _, l := range locs {
		if c, ok := cm[l]; ok && !c.visited {
			n = append(n, c)
		}
	}
	return n
}

func input2slice(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var text []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		text = append(text, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return text
}
