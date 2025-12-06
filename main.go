package main

import (
	"log"
	"os"
	"slices"
	"strconv"
)

type Solution interface {
	SolvePartA()
	SolvePartB()
}

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) < 1 {
		log.Fatal("Define a day to test")
	}

	day, err := strconv.Atoi(cliArgs[0])
	if err != nil {
		log.Fatal(err)
	}

	isTest := slices.Contains(cliArgs, "--test")

	var dailySolution Solution
	switch day {
	case 1:
		dailySolution = NewDayOne(isTest)
	case 2:
		dailySolution = NewDayTwo(isTest)
	case 3:
		dailySolution = NewDayThree(isTest)
	case 4:
		dailySolution = NewDayFour(isTest)
	case 5:
		dailySolution = NewDayFive(isTest)
	default:
		log.Fatal("Not Yet Implemented")

	}

	dailySolution.SolvePartA()
	dailySolution.SolvePartB()
}
