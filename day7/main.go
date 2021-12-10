package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	text := input2slice("day7/input")

	max := 0
	crabs := make([]int, 0)
	for _, val := range strings.Split(text[0], ",") {
		n, _ := strconv.Atoi(val)
		crabs = append(crabs, n)
		if n > max {
			max = n
		}
	}

	positions := make([]int, max+1)
	for _, crab := range crabs {
		positions[crab]++
	}

	minFuel1 := math.MaxInt
	minFuel1Pos := 0
	minFuel2 := math.MaxInt
	minFuel2Pos := 0
	for i := range positions {
		fuel1, fuel2 := computeFuel(positions, i)
		if fuel1 < minFuel1 {
			minFuel1 = fuel1
			minFuel1Pos = i
		}
		if fuel2 < minFuel2 {
			minFuel2 = fuel2
			minFuel2Pos = i
		}
	}

	fmt.Println("Part 1: The minimum fuel of", minFuel1, "is at position", minFuel1Pos)
	fmt.Println("Part 2: The minimum fuel of", minFuel2, "is at position", minFuel2Pos)
}

func computeFuel(positionHistogram []int, destination int) (int, int) {
	fuelUsed1 := 0
	fuelUsed2 := 0
	for pos, numCrabs := range positionHistogram {
		if numCrabs != 0 {
			distance := Abs(destination - pos)
			fuelUsed1 += distance * numCrabs
			cost := (distance*distance + distance) / 2
			fuelUsed2 += cost * numCrabs
		}
	}
	fmt.Println("Destination", destination, fuelUsed1, fuelUsed2)
	return fuelUsed1, fuelUsed2
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

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
