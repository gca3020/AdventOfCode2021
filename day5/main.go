package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	text := input2slice("day5/input")

	lines := make([]Line, 0)
	dx, dy := 0, 0
	for _, l := range text {
		line := NewLine(l)
		x, y := line.MaxXY()
		if x > dx {
			dx = x
		}
		if y > dy {
			dy = y
		}
		lines = append(lines, line)
	}

	m := NewMap(dx+1, dy+1)
	for _, line := range lines {
		m.AddLine(line)
	}
	fmt.Println("Counted", m.Overlaps(), "Overlapping points")
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

type Point struct {
	X int
	Y int
}

func NewPoint(s string) Point {
	tokens := strings.Split(s, ",")
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])
	return Point{x, y}
}

// Line defines a line with a start point and an end point
type Line struct {
	p1 Point
	p2 Point
}

// NewLine creates a Line from a text string
func NewLine(s string) Line {
	tokens := strings.Split(s, " -> ")
	return Line{NewPoint(tokens[0]), NewPoint(tokens[1])}
}

// CoveredPoints returns a slice containing every Point covered by this Line
func (l *Line) CoveredPoints() []Point {
	points := make([]Point, 0)

	if l.p1.X == l.p2.X {
		ymin, ymax := minMax(l.p1.Y, l.p2.Y)
		for y := ymin; y <= ymax; y++ {
			points = append(points, Point{l.p1.X, y})
		}
	} else if l.p1.Y == l.p2.Y {
		xmin, xmax := minMax(l.p1.X, l.p2.X)
		for x := xmin; x <= xmax; x++ {
			points = append(points, Point{x, l.p1.Y})
		}
	} else {
		var x1, x2, y1, y2 int
		if l.p1.X < l.p2.X {
			x1, x2 = l.p1.X, l.p2.X
			y1, y2 = l.p1.Y, l.p2.Y
		} else {
			x1, x2 = l.p2.X, l.p1.X
			y1, y2 = l.p2.Y, l.p1.Y
		}

		y := y1
		for x := x1; x <= x2; x++ {
			points = append(points, Point{x, y})
			if y < y2 {
				y++
			} else {
				y--
			}
		}
	}

	return points
}

func (l *Line) MaxXY() (int, int) {
	_, x := minMax(l.p1.X, l.p2.X)
	_, y := minMax(l.p1.Y, l.p2.Y)
	return x, y
}

type Map struct {
	coordinates [][]int
}

func NewMap(dx int, dy int) Map {
	coords := make([][]int, dy)
	for r := range coords {
		coords[r] = make([]int, dx)
	}
	return Map{coords}
}

func (m *Map) AddLine(line Line) {
	for _, p := range line.CoveredPoints() {
		m.coordinates[p.Y][p.X]++
	}
}

func (m *Map) Overlaps() int {
	num := 0
	for _, row := range m.coordinates {
		for _, val := range row {
			if val >= 2 {
				num++
			}
		}
	}
	return num
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
