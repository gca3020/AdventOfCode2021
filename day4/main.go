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
	text := input2slice("day4/input")

	// Read in the list of called numbers
	var numList []int
	numListText := strings.Split(text[0], ",")
	for _, numStr := range numListText {
		num, _ := strconv.Atoi(numStr)
		numList = append(numList, num)
	}

	// Construct the boards
	var boards []Board
	for i := 2; i+5 <= len(text); i += 6 {
		boards = append(boards, NewBoard(text[i:i+5]))
	}

	// Iterate over the input, marking boards until one is a winner
	for _, num := range numList {
		for i, _ := range boards {
			if boards[i].IsAWinner() {
				continue
			}
			if boards[i].Mark(num) {
				fmt.Println("Found a winning board!", boards[i].GetScore(num))
				boards[i].Dump()
			}
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
