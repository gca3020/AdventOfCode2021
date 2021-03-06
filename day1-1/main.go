package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open the input file
	file, err := os.Open("day1-1/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read through it line by line
	increasing := 0
	total := 0
	scanner := bufio.NewScanner(file)

	// Read the first line
	scanner.Scan()
	prev, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		total++
		if num > prev {
			increasing++
		}
		prev = num
	}

	fmt.Printf("Read %d entries, found %d increasing measurements\n", total, increasing)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}