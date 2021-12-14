package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	text := input2slice("day14/input")

	// Parse the input
	elements := strings.Split(text[0], "")
	rules := make(map[string]string, len(text)-2)
	for _, line := range text[2:] {
		tokens := strings.Split(line, " -> ")
		rules[tokens[0]] = tokens[1]
	}

	counts, pairs := seedHistograms(elements)
	for i := 0; i < 40; i++ {
		counts, pairs = step(counts, pairs, rules)
		if i == 9 {
			fmt.Println("The difference after 10 steps is", computeDifference(counts))
		}
	}
	fmt.Println("The difference after 40 steps is", computeDifference(counts))
}

// seedHistograms populates the count and pairs histograms so we can iterate with the step function
func seedHistograms(elements []string) (map[string]int, map[string]int) {
	counts := make(map[string]int)
	pairs := make(map[string]int)

	for i := range elements {
		if i+1 < len(elements) {
			pair := strings.Join(elements[i:i+2], "")
			pairs[pair] += 1
		}
		counts[elements[i]] += 1
	}

	return counts, pairs
}

// step performs a single step of the iteration, returning the updated histograms for the element and pair counts
func step(counts, pairs map[string]int, rules map[string]string) (map[string]int, map[string]int) {
	newPairs := make(map[string]int)
	for pair, cnt := range pairs {
		// Look up the new element created by this pair
		newElem := rules[pair]

		// Increment the element count of the new element by the number of times this pair appears
		counts[newElem] += cnt

		// Determine the two new pairs created from this pair, and increment their count in the new pairs histogram
		split := strings.Split(pair, "")
		newPairs[strings.Join([]string{split[0], newElem}, "")] += cnt
		newPairs[strings.Join([]string{newElem, split[1]}, "")] += cnt
	}
	return counts, newPairs
}

// computeDifference returns the difference in quantity between the most- and least-common element
func computeDifference(count map[string]int) int {
	c := make([]int, 0, len(count))
	for _, v := range count {
		c = append(c, v)
	}

	sort.Ints(c)
	return c[len(c)-1] - c[0]
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
