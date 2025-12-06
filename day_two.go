package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DayTwo struct {
	path string
	args []string
}

func parseRangeString(rangeString string) (int, int) {
	startAndEnd := strings.Split(rangeString, "-")
	start, err := strconv.Atoi(startAndEnd[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(startAndEnd[1])
	if err != nil {
		log.Fatal(err)
	}
	return start, end
}

func isInvalidID(id string) bool {
	// Iterate through every possible chunk size - can stop at half, since we need the chunk to repeat
	for i := range int(len(id) / 2) {
		chunk := id[:i+1]
		isInvalid := true
		j := 0

		// Iterate through every chunk to see if it repeats the pattern
		for j <= len(id)-len(chunk) && isInvalid {
			if id[j:j+len(chunk)] != chunk {
				isInvalid = false
			} else {
				j += len(chunk)
			}
		}

		// Check if we still had an 'invalid' id (i.e. repeating) AND we got to the end of the id
		if j == len(id) && isInvalid {
			return true
		}
	}

	return false
}

func NewDayTwo(isTest bool) *DayTwo {
	path := derivePath("day_two", isTest)
	args := loadFileArgs(path, ",")
	return &DayTwo{path, args}
}

func (d *DayTwo) SolvePartA() {
	sumOfInvalidIds := 0
	for _, r := range d.args {
		start, end := parseRangeString(r)

		for i := range end - start + 1 {
			id := strconv.Itoa(start + i)
			if id[0:len(id)/2] == id[len(id)/2:] {
				sumOfInvalidIds += start + i
			}
		}
	}

	fmt.Printf("Part A: %d\n", sumOfInvalidIds)
}

func (d *DayTwo) SolvePartB() {
	sumOfInvalidIds := 0
	for _, r := range d.args {
		start, end := parseRangeString(r)

		for i := range end - start + 1 {
			id := strconv.Itoa(start + i)
			if isInvalidID(id) {
				sumOfInvalidIds += start + i
			}
		}
	}

	fmt.Printf("Part B: %d\n", sumOfInvalidIds)
}
