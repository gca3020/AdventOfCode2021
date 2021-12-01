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
	file, err := os.Open("day1-2/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read from input into a slice
	var readings []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val,_ := strconv.Atoi(scanner.Text())
		readings = append(readings, val)
	}

	total := 0
	prev := 0
	increasing := 0

	for i := 0; i < len(readings)-2; i++ {
		val := readings[i] + readings[i+1] + readings[i+2]
		total++
		if prev != 0 && val > prev {
			increasing++
		}
		prev = val
	}

	fmt.Printf("Read %d entries, found %d increasing measurements\n", total, increasing)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}