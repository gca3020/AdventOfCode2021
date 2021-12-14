package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Fold struct {
	dir  string
	line int
}

func main() {
	text := input2slice("day13/input")
	points, folds := parseDirections(text)

	for _, fold := range folds {
		points = doFold(points, fold)
	}
	fmt.Println("Counted", len(points))
	printCanvas(points)
}

// doFold performs a "fold" by transposing any coordinates on the far side of it
func doFold(points []Point, fold Fold) []Point {
	newPoints := make([]Point, 0)
	for _, point := range points {
		newPoint := point
		if fold.dir == "y" && point.y > fold.line {
			newPoint.y = fold.line - (point.y - fold.line)
		} else if fold.dir == "x" && point.x > fold.line {
			newPoint.x = fold.line - (point.x - fold.line)
		}

		if !containsPoint(newPoints, newPoint) {
			newPoints = append(newPoints, newPoint)
		}
	}
	return newPoints
}

func containsPoint(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

// printCanvas returns a canvas with the points populated. 1 for present, 0 for absent
func printCanvas(points []Point) {
	// Get the dimensions required based on the largest coordinate values
	x, y := 0, 0
	for _, p := range points {
		if p.x > x {
			x = p.x
		}
		if p.y > y {
			y = p.y
		}
	}

	// Create the canvas big enough to hold the points
	canvas := make([][]int, y+1)
	for y := range canvas {
		canvas[y] = make([]int, x+1)
	}

	// Loop through the points, adding them to the canvas
	for _, point := range points {
		canvas[point.y][point.x] += 1
	}

	// Print the canvas
	for y := range canvas {
		for _, val := range canvas[y] {
			if val == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println("")
	}
}

// parseDirections parses the text input into slices of Point and Fold
func parseDirections(lines []string) ([]Point, []Fold) {
	points := make([]Point, 0)
	folds := make([]Fold, 0)
	for _, line := range lines {
		if strings.Contains(line, ",") {
			tokens := strings.Split(line, ",")
			x, _ := strconv.Atoi(tokens[0])
			y, _ := strconv.Atoi(tokens[1])
			points = append(points, Point{x, y})
		} else if strings.Contains(line, "fold") {
			tokens := strings.Fields(line)
			fold := strings.Split(tokens[2], "=")
			line, _ := strconv.Atoi(fold[1])
			folds = append(folds, Fold{fold[0], line})
		}
	}
	return points, folds
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
