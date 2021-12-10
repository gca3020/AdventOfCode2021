package main

import (
	"fmt"
	"strconv"
	"strings"
)

type space struct {
	num    int
	marked bool
}

type Board struct {
	spaces [][]space
	winner bool
}

// NewBoard returns a new board object constructed from the given input, a square grid of space-delimited integers
func NewBoard(text []string) Board {
	spaces := make([][]space, 0, 5)
	for _, rowText := range text {
		row := make([]space, 0, 5)
		for _, numText := range strings.Fields(rowText) {
			num, _ := strconv.Atoi(numText)
			row = append(row, space{num, false})
		}
		spaces = append(spaces, row)
	}
	return Board{spaces, false}
}

func (b *Board) Dump() {
	fmt.Println("Dumping board:")
	for _, row := range b.spaces {
		fmt.Println(row)
	}
}

func (b *Board) Mark(called int) bool {
	for i, row := range b.spaces {
		for j, space := range row {
			if space.num == called {
				b.spaces[i][j].marked = true
			}
		}
	}

	// Check for completed rows or columns
	for i := 0; i < 5; i++ {
		if (b.spaces[0][i].marked && b.spaces[1][i].marked && b.spaces[2][i].marked && b.spaces[3][i].marked && b.spaces[4][i].marked) ||
			(b.spaces[i][0].marked && b.spaces[i][1].marked && b.spaces[i][2].marked && b.spaces[i][3].marked && b.spaces[i][4].marked) {
			b.winner = true
		}
	}

	return b.IsAWinner()
}

func (b *Board) IsAWinner() bool {
	return b.winner
}

func (b *Board) GetScore(multiplier int) int {
	sum := 0
	for _, row := range b.spaces {
		for _, space := range row {
			if !space.marked {
				sum += space.num
			}
		}
	}
	return multiplier * sum
}
