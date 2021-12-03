package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	text := input2slice("day3/input")

	part1(text)
	part2(text)
}

func countOnes(input []string, column int) (ones int) {
	ones = 0
	for _, val := range input {
		if val[column] == '1' {
			ones++
		}
	}
	return ones
}

func part1(input []string) {
	gamma := 0
	mask := 0
	numBits := len(input[0])
	mostCommon := make([]int, numBits)

	for i := 0; i < numBits; i++ {
		if countOnes(input, i) > len(input)/2 {
			mostCommon[i] = 1
		} else {
			mostCommon[i] = 0
		}
	}

	for i := 0; i < len(mostCommon); i++ {
		mask |= 1 << i
		gamma <<= 1
		gamma |= mostCommon[i]
	}
	epsilon := ^gamma & mask

	fmt.Printf("Part 1: mostCpmmon: %v, gamma: 0b%b (%d), epsilon: 0b%b (%d), mask: 0x%x, result: %d\n",
		mostCommon, gamma, gamma, epsilon, epsilon, mask, gamma*epsilon)
}

func part2(input []string) {
	oxygenRating := computeOxygenRating(input, 0)
	scrubberRating := computeScrubberRating(input, 0)

	fmt.Printf("Part 2: O2 Generator = %d, Scrubber = %d, Life Support = %d",
		oxygenRating, scrubberRating, oxygenRating*scrubberRating)
}

func computeOxygenRating(input []string, col int) int64 {
	if len(input) == 1 {
		d, _ := strconv.ParseInt(input[0], 2, 64)
		return d
	}

	filtered := make([]string, 0)
	ones := countOnes(input, col)
	zeros := len(input) - ones

	keep := "0"
	if ones >= zeros {
		keep = "1"
	}

	for _, val := range input {
		if string(val[col]) == keep {
			filtered = append(filtered, val)
		}
	}
	return computeOxygenRating(filtered, col+1)
}

func computeScrubberRating(input []string, col int) int64 {
	if len(input) == 1 {
		d, _ := strconv.ParseInt(input[0], 2, 64)
		return d
	}

	filtered := make([]string, 0)
	ones := countOnes(input, col)
	zeros := len(input) - ones

	keep := "1"
	if zeros <= ones {
		keep = "0"
	}

	for _, val := range input {
		if string(val[col]) == keep {
			filtered = append(filtered, val)
		}
	}
	return computeScrubberRating(filtered, col+1)
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
