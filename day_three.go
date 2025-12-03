package main

import (
	"fmt"
	"log"
	"os"
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

	d.args = strings.Split(string(dat), ",")
}

func (d *DayThree) SolvePartA() {
	fmt.Printf("Part A: \n")
}

func (d *DayThree) SolvePartB() {
	fmt.Printf("Part B: \n")
}
