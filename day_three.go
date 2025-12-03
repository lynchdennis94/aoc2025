package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type DayThree struct {
	path string
	args []string
}

func NewDayThree(isTest bool) *DayThree {
	path := "puzzles/day_three"
	if isTest {
		path += "_test.txt"
	} else {
		path += ".txt"
	}

	return &DayThree{path: path}
}

func (d *DayThree) LoadFile() {
	dat, err := os.ReadFile(d.path)
	if err != nil {
		log.Fatal(err)
	}

	d.args = strings.Split(string(dat), "\n")
}

func findLargestDigitInRange(bank string, start, end int) (int, int) {
	maxDigitPlace := -1
	maxDigit := -1
	for i := range end - start {
		currentDigitString := string(bank[i+start])
		currentDigit, err := strconv.Atoi(currentDigitString)
		if err != nil {
			log.Fatal(err)
		}
		// We want maximum search space for later searches, so only update if definitively larger
		if currentDigit > maxDigit {
			maxDigit = currentDigit
			maxDigitPlace = i
		}

		// We won't find higher than 9, so stop
		if currentDigit == 9 {
			break
		}
	}

	return maxDigit, maxDigitPlace + start
}

func findJoltage(bank string, batteryCount int) int {
	startingPlace := 0
	joltageString := ""

	// Loop through each battery we need to turn on
	for i := range batteryCount {
		// End is defined to allow for enough room for remaining digits to be found
		digit, place := findLargestDigitInRange(bank, startingPlace, len(bank)+i-batteryCount+1)

		// Start one _after_ where we just stopped on the next iteration
		startingPlace = place + 1
		joltageString += strconv.Itoa(digit)
	}

	// Convert the built-up string back into an int, and return
	if joltage, err := strconv.Atoi(joltageString); err != nil {
		log.Fatal(err)
	} else {
		return joltage
	}
	return 0
}

func (d *DayThree) SolvePartA() {
	totalOutputJoltage := 0
	for _, bank := range d.args {
		totalOutputJoltage += findJoltage(bank, 2)
	}
	fmt.Printf("Part A: %d\n", totalOutputJoltage)
}

func (d *DayThree) SolvePartB() {
	totalOutputJoltage := 0
	for _, bank := range d.args {
		totalOutputJoltage += findJoltage(bank, 12)
	}
	fmt.Printf("Part B: %d\n", totalOutputJoltage)
}
