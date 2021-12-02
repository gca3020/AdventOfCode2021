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
	text := input2slice("day2/input")

	part1(text)
	part2(text)
}

func part1(input []string) {
	distance := 0
	depth := 0

	for _, step := range input {
		s := strings.Fields(step)
		dist, _ := strconv.Atoi(s[1])

		switch s[0] {
		case "forward":
			distance += dist
			break
		case "up":
			depth -= dist
			break
		case "down":
			depth += dist
			break
		}
	}

	fmt.Printf("Part 1: Distance = %d, Depth = %d, Total = %d\n", distance, depth, distance*depth)
}

func part2(input []string) {
	aim := 0
	distance := 0
	depth := 0

	for _, step := range input {
		s := strings.Fields(step)
		dist, _ := strconv.Atoi(s[1])

		switch s[0] {
		case "forward":
			distance += dist
			depth += aim * dist
			break
		case "up":
			aim -= dist
			break
		case "down":
			aim += dist
			break
		}
	}

	fmt.Printf("Part 1: Distance = %d, Depth = %d, Total = %d\n", distance, depth, distance*depth)
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
