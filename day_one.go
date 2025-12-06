package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type DayOne struct {
	path string
	args []string
}

func NewDayOne(isTest bool) *DayOne {
	path := derivePath("day_one", isTest)
	args := loadFileArgs(path, "\n")
	return &DayOne{path, args}
}

func calculateNextTurn(currentSpot int, turnInstruction string) (int, int) {
	nextSpot := currentSpot
	timesCrossedZero := 0

	direction := string(turnInstruction[0])
	clicks, err := strconv.Atoi(turnInstruction[1:])
	if err != nil {
		log.Fatal(err)
	}

	// For every 100 clicks, we are guaranteed to cross zero
	timesCrossedZero += int(math.Floor(float64(clicks) / 100))
	clicks = clicks % 100

	// We're on our final turn; outside of [0,99], if we didn't start/end on 0 we crossed zero
	if direction == "L" {
		nextSpot -= clicks
		if nextSpot < 0 {
			nextSpot += 100
			if currentSpot != 0 && nextSpot != 0 {
				timesCrossedZero++
			}
		}
	} else {
		nextSpot += clicks
		if nextSpot >= 100 {
			nextSpot -= 100
			if currentSpot != 0 && nextSpot != 0 {
				timesCrossedZero++
			}
		}
	}

	return nextSpot, timesCrossedZero
}

func (d *DayOne) SolvePartA() {
	currentPosition := 50
	timesPointingAtZero := 0
	for _, arg := range d.args {
		currentPosition, _ = calculateNextTurn(currentPosition, arg)
		if currentPosition == 0 {
			timesPointingAtZero++
		}
	}

	fmt.Println(timesPointingAtZero)
}

func (d *DayOne) SolvePartB() {
	currentPosition := 50
	timesCrossedZero := 0
	for _, arg := range d.args {
		var crossedZeroCount int
		currentPosition, crossedZeroCount = calculateNextTurn(currentPosition, arg)
		timesCrossedZero += crossedZeroCount

		if currentPosition == 0 {
			timesCrossedZero++
		}
	}

	fmt.Println(timesCrossedZero)
}
