package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Loc struct {
	val        int
	basinFound bool
}

func main() {
	text := input2slice("day9/input")

	// Read in the map as a multi-dimensional slice of ints, tag anything with a value of 9 as not yet having its
	// basin found
	floor := make([][]Loc, 0)
	for _, textRow := range text {
		row := make([]Loc, 0, len(text))
		for _, char := range strings.Split(textRow, "") {
			val, _ := strconv.Atoi(char)
			row = append(row, Loc{val, false})
		}
		floor = append(floor, row)
	}

	fmt.Println("Calculated a risk level of", computeRiskLevel(floor))
	fmt.Println("The product of the size of the three largest basins is", findBasins(floor))
}

func computeRiskLevel(floor [][]Loc) (riskLevel int) {
	for y, row := range floor {
		for x, loc := range row {
			neighbors := make([]int, 0)
			// Append the bottom neighbor if possible
			if y+1 < len(floor) {
				neighbors = append(neighbors, floor[y+1][x].val)
			}
			// Append the top neighbor if possible
			if y-1 >= 0 {
				neighbors = append(neighbors, floor[y-1][x].val)
			}
			// Append the right neighbor if possible
			if x+1 < len(row) {
				neighbors = append(neighbors, floor[y][x+1].val)
			}
			// Append the left neighbor if possible
			if x-1 >= 0 {
				neighbors = append(neighbors, floor[y][x-1].val)
			}

			if isLocalMin(loc.val, neighbors) {
				riskLevel += loc.val + 1
			}
		}
	}
	return
}

func findBasins(floor [][]Loc) int {
	basinSizes := make([]int, 0)
	for y, row := range floor {
		for x := range row {
			// Recursively search for basins
			foundBasinSize := addToBasin(y, x, floor)
			if foundBasinSize > 0 {
				//fmt.Println("Found a basin of size", foundBasinSize)
				basinSizes = append(basinSizes, foundBasinSize)
			}
		}
	}

	// Sort the basin sizes
	sort.Ints(basinSizes)

	//fmt.Println("Found the following sizes", basinSizes)
	return basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
}

func addToBasin(y, x int, floor [][]Loc) int {
	if x < 0 || y < 0 || x >= len(floor[0]) || y >= len(floor) || floor[y][x].val == 9 || floor[y][x].basinFound {
		return 0
	}
	//fmt.Printf("Adding (%d) (%d, %d) to basin\n", floor[y][x].val, y, x)
	floor[y][x].basinFound = true
	return 1 + addToBasin(y, x-1, floor) + addToBasin(y, x+1, floor) + addToBasin(y-1, x, floor) + addToBasin(y+1, x, floor)
}

func isLocalMin(val int, neighbors []int) bool {
	for _, n := range neighbors {
		if val >= n {
			return false
		}
	}
	return true
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
