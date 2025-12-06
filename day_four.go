package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

var deltas = []Point{
	{-1, 0},
	{-1, -1},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type DayFour struct {
	path string
	args []string
}

func NewDayFour(isTest bool) *DayFour {
	path := "puzzles/day_four"
	if isTest {
		path += "_test.txt"
	} else {
		path += ".txt"
	}

	return &DayFour{path: path}
}

func (d *DayFour) LoadFile() {
	dat, err := os.ReadFile(d.path)
	if err != nil {
		log.Fatal(err)
	}

	d.args = strings.Split(string(dat), "\n")
}

func (d *DayFour) SolvePartA() {
	grid := parseInput(d.args)
	sparseNeighborCount := 0
	setOfRolls := getRolls(grid)

	for roll := range setOfRolls {
		if get8DRollNeighbors(roll.X, roll.Y, grid) < 4 {
			sparseNeighborCount++
		}
	}

	fmt.Printf("Part A: %d\n", sparseNeighborCount)
}

func (d *DayFour) SolvePartB() {
	grid := parseInput(d.args)
	removedRollCount := 0
	setOfRolls := getRolls(grid)

	foundRolls := true
	for foundRolls {
		foundRolls = false
		rollsToRemove := map[Point]bool{}

		// Track how many rolls are removed
		for roll := range setOfRolls {
			if get8DRollNeighbors(roll.X, roll.Y, grid) < 4 {
				removedRollCount++
				foundRolls = true
				rollsToRemove[roll] = true
			}
		}

		// Do cleanup
		for roll := range rollsToRemove {
			grid[roll.X][roll.Y] = '.'
			delete(setOfRolls, roll)
		}
	}

	fmt.Printf("Part B: %d\n", removedRollCount)
}

func parseInput(input []string) [][]rune {
	output := make([][]rune, len(input))
	for i := range len(input) {
		output[i] = make([]rune, len(input[i]))
		for j := range len(input[i]) {
			output[i][j] = rune(input[i][j])
		}
	}
	return output
}

func getRolls(grid [][]rune) map[Point]bool {
	setOfRolls := map[Point]bool{}

	// Create a set of positions that have rolls
	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] == '@' {
				setOfRolls[Point{i, j}] = true
			}
		}
	}

	return setOfRolls
}

func get8DRollNeighbors(x, y int, grid [][]rune) int {
	neighborCount := 0
	for _, delta := range deltas {
		dx := x + delta.X
		dy := y + delta.Y
		if dx >= 0 &&
			dx < len(grid) &&
			dy >= 0 &&
			dy < len(grid[0]) &&
			grid[dx][dy] == '@' {
			neighborCount++
		}
	}
	return neighborCount
}
