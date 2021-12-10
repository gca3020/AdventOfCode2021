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
	text := input2slice("day6/input")

	// Populate the sea with plenty of fish
	sea := make([]int, 9)
	for _, val := range strings.Split(text[0], ",") {
		day, _ := strconv.Atoi(val)
		sea[day]++
	}

	// Compute what tomorrow's sea will look like
	for i := 0; i <= 256; i++ {
		fmt.Println("Day", i, countFish(sea), sea)
		sea = computeTomorrow(sea)
	}
}

func countFish(sea []int) (sum int) {
	for _, n := range sea {
		sum += n
	}
	return
}

func computeTomorrow(sea []int) []int {
	tomorrow := make([]int, 9)

	// All the fish with a timer of 0 will be reproducing today
	reproducing := sea[0]

	// Every other fish gets decremented by a day (shift the histogram left)
	for i := 1; i < len(sea); i++ {
		tomorrow[i-1] = sea[i]
	}

	// The reproducing fish end up in day 6, and their spawn in day 8
	tomorrow[6] += reproducing
	tomorrow[8] = reproducing
	return tomorrow
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
