package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	text := input2slice("day17/input")
	re := regexp.MustCompile(`x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)
	matches := re.FindStringSubmatch(text[0])

	xmin, _ := strconv.Atoi(matches[1])
	xmax, _ := strconv.Atoi(matches[2])
	ymin, _ := strconv.Atoi(matches[3])
	ymax, _ := strconv.Atoi(matches[4])

	fmt.Printf("Target Box from %d,%d to %d,%d\n", xmin, ymin, xmax, ymax)

	allH := make([]int, 0)
	maxVx := 0
	maxVy := 0
	maxH := 0
	for vx := 0; vx <= xmax; vx++ {
		for vy := ymin; vy < 10000; vy++ {
			h, hit := takeShot(vx, vy, xmin, xmax, ymin, ymax)
			if hit {
				allH = append(allH, h)
			}
			if h > maxH {
				maxH = h
				maxVx = vx
				maxVy = vy
			}
		}
	}

	fmt.Println(maxH, maxVx, maxVy, len(allH))
}

func takeShot(vx, vy int, xmin, xmax, ymin, ymax int) (int, bool) {
	hmax := 0
	for x, y := 0, 0; y > ymin; {
		x, y, vx, vy = step(x, y, vx, vy)
		if y > hmax {
			hmax = y
		}
		if x >= xmin && y >= ymin && x <= xmax && y <= ymax {
			return hmax, true
		}
	}
	return 0, false
}

func step(x, y, vx, vy int) (x1, y1, vx1, vy1 int) {
	x1 = x + vx
	y1 = y + vy
	vy1 = vy - 1
	if vx < 0 {
		vx1 = vx + 1
	} else if vx > 0 {
		vx1 = vx - 1
	} else {
		vx1 = 0
	}
	return
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
