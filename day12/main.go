package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var allPaths = make([][]string, 0)

func main() {
	text := input2slice("day12/input")
	connections := buildConnections(text)
	fmt.Println(connections)

	history := make([]string, 0)
	findPath(connections, history, "start", false)
	fmt.Println("Calculated", len(allPaths), "paths with default rules")

	allPaths = make([][]string, 0)
	history = make([]string, 0)
	findPath(connections, history, "start", true)
	fmt.Println("Calculated", len(allPaths), "where revisiting a single small path is allowed")
}

func findPath(connections map[string][]string, history []string, next string, allowTwice bool) {
	// If this is a small cave that we have already visited, just return early, since this is a dead end
	if isSmall(next) && alreadyVisited(next, history, allowTwice) {
		return
	}

	// Otherwise, we will append this node to our history
	history = append(history, next)

	// Check to see if we're at the end
	if isEnd(next) {
		allPaths = append(allPaths, history)
		return
	}

	// If we got this far, we're either at a new small node, or a big node. So we need to find all the possible next
	// hops from our current location, and recursively trace through them
	for _, n := range connections[next] {
		findPath(connections, history, n, allowTwice)
	}
}

// buildConnections returns a map of the connections between nodes.
func buildConnections(text []string) map[string][]string {
	conns := make(map[string][]string, 0)
	for _, edge := range text {
		nodes := strings.Split(edge, "-")
		conns[nodes[0]] = append(conns[nodes[0]], nodes[1])
		conns[nodes[1]] = append(conns[nodes[1]], nodes[0])
	}
	return conns
}

// isSmall returns whether a given cave is a small cave
func isSmall(name string) bool {
	return strings.ToLower(name) == name
}

// isEnd returns whether a given cave is the end destination
func isEnd(name string) bool {
	return name == "end"
}

// alreadyVisited returns whether a given cave has already been visited
func alreadyVisited(name string, history []string, allowTwice bool) bool {
	visited := alreadyVisitedStrict(name, history)
	if !allowTwice {
		return visited
	} else {
		if !visited {
			return false
		}
		if name == "start" || anySmallNodeVisitedTwice(history) {
			return true
		}
		return false
	}
}

func alreadyVisitedStrict(name string, history []string) bool {
	for _, n := range history {
		if n == name {
			return true
		}
	}
	return false
}

func anySmallNodeVisitedTwice(history []string) bool {
	for i, name := range history {
		if isSmall(name) {
			for _, name2 := range history[i+1:] {
				if name == name2 {
					return true
				}
			}
		}
	}
	return false
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
