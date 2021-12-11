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
	text := input2slice("day8/input")

	decoders := make([]Decoder, 0)
	for _, entry := range text {
		decoders = append(decoders, NewDecoder(entry))
	}

	easyDigitCount := 0
	totalOutput := 0
	for _, decoder := range decoders {
		for _, digit := range decoder.GetEasyDigits() {
			// Part 1, just count the easy digits
			if digit == 1 || digit == 4 || digit == 7 || digit == 8 {
				easyDigitCount++
			}
		}

		// Part 2, Add up the full output values
		totalOutput += decoder.GetOutput()
	}

	fmt.Println("Counted", easyDigitCount, "easy to decode digits")
	fmt.Println("Total sum of entries is", totalOutput)
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

type Decoder struct {
	patterns  []string
	digits_ct []string
	ct_to_pt  map[rune]rune
}

func NewDecoder(line string) Decoder {
	s := strings.Split(line, "|")
	patterns := strings.Fields(s[0])
	digits_ct := strings.Fields(s[1])
	d := Decoder{patterns, digits_ct, make(map[rune]rune)}
	d.generateCtPtMap()
	return d
}

func (d *Decoder) GetEasyDigits() []int {
	digits := make([]int, 4)

	// First go through and extract the easy digits by their length
	for i, digit_ct := range d.digits_ct {
		if len(digit_ct) == 2 {
			digits[i] = 1
		} else if len(digit_ct) == 3 {
			digits[i] = 7
		} else if len(digit_ct) == 4 {
			digits[i] = 4
		} else if len(digit_ct) == 7 {
			digits[i] = 8
		} else {
			// TODO
			digits[i] = -1
		}
	}
	return digits
}

func (d *Decoder) GetOutput() int {
	digits := make([]int, 4)
	for i, digit_ct := range d.digits_ct {
		digit_pt := strings.ToUpper(digit_ct)
		for ct, pt := range d.ct_to_pt {
			digit_pt = strings.ReplaceAll(digit_pt, strings.ToUpper(string(ct)), string(pt))
		}
		digits[i] = d.lookupDigit(digit_pt)
	}
	// Debug - Print the digit sequence
	//fmt.Println(digits)

	return digits[0]*1000 + digits[1]*100 + digits[2]*10 + digits[3]
}

func (d *Decoder) generateCtPtMap() {
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	for _, ct := range letters {
		switch d.containedInPatterns(ct) {
		case 4:
			d.ct_to_pt[ct] = 'e'
			// Must be an E
			break
		case 6:
			d.ct_to_pt[ct] = 'b'
			// Must be a B
			break
		case 7:
			// Can be a D or a G
			if d.containedInPatternsOfLength(ct, 4) == 0 {
				// Segment D does not appear in '4', which is the only pattern of length 4
				d.ct_to_pt[ct] = 'g'
			} else {
				d.ct_to_pt[ct] = 'd'
			}
			break
		case 8:
			// Can be an A or a C
			if d.containedInPatternsOfLength(ct, 2) == 0 {
				// Segment A does not appear in '1' which is the only pattern of length 2
				d.ct_to_pt[ct] = 'a'
			} else {
				d.ct_to_pt[ct] = 'c'
			}
			break
		case 9:
			// Must be an F
			d.ct_to_pt[ct] = 'f'
			break
		}
	}
	/* -- Debug
	fmt.Printf("Dumping CT to PT Map: [")
	for _, c := range letters {
		fmt.Printf("%c:%c ", c, d.ct_to_pt[c])
	}
	fmt.Printf("]\n")
	*/
}

func (d *Decoder) containedInPatterns(r rune) int {
	count := 0
	for _, pattern := range d.patterns {
		if strings.ContainsRune(pattern, r) {
			count++
		}
	}
	return count
}

func (d *Decoder) containedInPatternsOfLength(r rune, l int) int {
	count := 0
	for _, pattern := range d.patterns {
		if len(pattern) == l && strings.ContainsRune(pattern, r) {
			count++
		}
	}
	return count
}

func (d *Decoder) lookupDigit(segments_pt string) int {
	// Sort the characters in the string
	chars := strings.Split(segments_pt, "")
	sort.Strings(chars)
	sorted_segment := strings.Join(chars, "")

	lut := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	return lut[sorted_segment]
}
