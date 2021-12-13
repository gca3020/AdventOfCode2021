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
	text := input2slice("day10/input")

	// Compute the score for each line
	errorScore := 0
	completionScores := make([]int, 0)
	for _, line := range text {
		lineError, lineCompletion := computeSyntaxScores(line)
		errorScore += lineError
		if lineCompletion != 0 {
			completionScores = append(completionScores, lineCompletion)
		}
	}

	sort.Ints(completionScores)
	fmt.Println("Computed an error score of", errorScore)
	fmt.Println("Computed a completion score of", completionScores[len(completionScores)/2])
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

func computeSyntaxScores(line string) (errorScore int, completionScore int) {
	stack := make([]string, 0)
	for _, s := range strings.Split(line, "") {
		switch s {
		case "(", "[", "{", "<":
			stack = push(stack, s)
			break
		case ")", "]", "}", ">":
			if peek(stack) == matchingChar(s) {
				stack = pop(stack)
			} else {
				return charScore(s), 0
			}
			break
		default:
			fmt.Println("Error: Non-parseable character found:", s)
			break
		}
	}

	// Now that we have fully parsed the line and found it incomplete, we need to match the characters
	completionScore = 0
	for done := false; !done; {
		unmatched := peek(stack)
		if unmatched == "" {
			done = true
			break
		}

		completionScore = completionScore*5 + charScore(unmatched)
		stack = pop(stack)
	}

	return 0, completionScore
}

func matchingChar(character string) string {
	switch character {
	case ")":
		return "("
	case "]":
		return "["
	case "}":
		return "{"
	case ">":
		return "<"
	default:
		return ""
	}
}

func charScore(character string) int {
	switch character {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	case "(":
		return 1
	case "[":
		return 2
	case "{":
		return 3
	case "<":
		return 4
	default:
		return 0
	}
}

func push(stack []string, s string) []string {
	return append(stack, s)
}

func pop(stack []string) []string {
	if len(stack) > 0 {
		top := len(stack) - 1
		return stack[:top]
	}
	return stack
}

func peek(stack []string) string {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return ""
}
