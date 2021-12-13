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
	text := input2slice("day11/input")

	// Parse the input into a board
	board := make([][]Octopus, 0, 10)
	for _, line := range text {
		row := make([]Octopus, 0, 10)
		levels := strings.Split(line, "")
		for _, levelStr := range levels {
			level, _ := strconv.Atoi(levelStr)
			row = append(row, Octopus{level, false})
		}
		board = append(board, row)
	}

	totalFlashes := 0
	done := false
	for i := 0; !done; i++ {
		flashes := step(board)
		totalFlashes += flashes
		if i == 99 {
			fmt.Printf("Got a total of %d flashes after %d steps\n", totalFlashes, i+1)
		}

		// If they all flashed, we synchronized, and we're done
		if flashes == 100 {
			done = true
			fmt.Printf("Got %d flashes on step %d", flashes, i+1)
		}
	}
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

type Octopus struct {
	level   int
	flashed bool
}

func step(board [][]Octopus) int {
	// Increment the energy level of every octopus by 1
	for y := range board {
		for x := range board[y] {
			board[y][x].level += 1
		}
	}

	// Recursively compute the number of flashes
	for y := range board {
		for x := range board[y] {
			flash(board, y, x)
		}
	}

	// Count the total number of flashes and reset each octopus who flashed to 0 with
	numFlashes := 0
	for y := range board {
		for x := range board[y] {
			if board[y][x].flashed {
				numFlashes++
				board[y][x].flashed = false
				board[y][x].level = 0
			}
		}
	}

	return numFlashes
}

func isInBounds(y, x int) bool {
	return y >= 0 && x >= 0 && y < 10 && x < 10
}

func flash(board [][]Octopus, y, x int) {
	// If this cell isn't in the bounds of the array, just return
	if !isInBounds(y, x) {
		return
	}

	// If this cell has already flashed, just return
	if board[y][x].flashed {
		return
	}

	// If this cell has a value greater than 9, it flashes, which increments the cells around it
	if board[y][x].level > 9 {
		board[y][x].flashed = true
		for ny := y - 1; ny <= y+1; ny++ {
			for nx := x - 1; nx <= x+1; nx++ {
				if isInBounds(ny, nx) {
					board[ny][nx].level += 1
					flash(board, ny, nx)
				}
			}
		}
	}
}
