package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DayFive struct {
	path string
	args []string
}

type IngredientRange struct {
	start int
	end   int
}

func NewDayFive(isTest bool) *DayFive {
	path := derivePath("day_five", isTest)
	args := loadFileArgs(path, "\n")
	return &DayFive{path, args}
}

func (d *DayFive) SolvePartA() {
	freshIngredients, ingredientList := d.processDatabaseInput(true)

	freshIngredientCount := 0
	for _, ingredient := range ingredientList {
		// Check if it exists in any range
		for _, ingredientRange := range freshIngredients {
			if ingredient >= ingredientRange.start && ingredient <= ingredientRange.end {
				freshIngredientCount++
				break
			}
		}
	}
	fmt.Printf("Part A: %d\n", freshIngredientCount)
}

func (d *DayFive) SolvePartB() {
	freshIngredientList, _ := d.processDatabaseInput(false)

	var modifiedFreshIngredientList []IngredientRange
	modifiedFreshIngredientList = freshIngredientList
	keepIterating := true
	for keepIterating {
		latestIteration, changeCount := d.finalizeFreshRanges(modifiedFreshIngredientList)
		modifiedFreshIngredientList = latestIteration
		if changeCount == 0 {
			keepIterating = false
		}
	}

	finalizedFreshRanges := []IngredientRange{}

	for _, tempFreshRange := range modifiedFreshIngredientList {
		contained := false

		for _, finalizedRange := range finalizedFreshRanges {
			if tempFreshRange.start == finalizedRange.start && tempFreshRange.end == finalizedRange.end {
				contained = true
				break
			}
		}

		if !contained {
			finalizedFreshRanges = append(finalizedFreshRanges, tempFreshRange)
		}

	}

	freshIngredientCount := 0
	for _, freshRange := range finalizedFreshRanges {
		freshIngredientCount += (freshRange.end - freshRange.start + 1)
	}

	fmt.Printf("Part B: %d\n", freshIngredientCount)
}

func (d *DayFive) processDatabaseInput(includeIngredientList bool) ([]IngredientRange, []int) {
	ingredientList := []int{}
	ingredientRanges := []IngredientRange{}
	processingFreshIngredientRanges := true
	for _, arg := range d.args {
		if arg == "" {
			processingFreshIngredientRanges = false
			continue
		}

		if processingFreshIngredientRanges {
			ingredientRange := strings.Split(arg, "-")
			if len(ingredientRange) != 2 {
				log.Fatal("Too many ingredients in range")
			} else {
				start, err := strconv.Atoi(ingredientRange[0])
				if err != nil {
					log.Fatal(err)
				}
				end, err := strconv.Atoi(ingredientRange[1])
				if err != nil {
					log.Fatal(err)
				}

				ingredientRanges = append(ingredientRanges, IngredientRange{start, end})
			}
		} else if includeIngredientList {
			if id, err := strconv.Atoi(arg); err == nil {
				ingredientList = append(ingredientList, id)
			}
		}
	}

	return ingredientRanges, ingredientList
}

func (d *DayFive) finalizeFreshRanges(startingFreshRanges []IngredientRange) ([]IngredientRange, int) {
	changeCount := 0
	tempFreshRanges := []IngredientRange{}

	for i, freshRange := range startingFreshRanges {
		// See if there's overlap with existing ranges
		tempRange := freshRange
		for j, otherFreshRange := range startingFreshRanges {
			if i == j {
				continue
			}

			// If there's overlap with the 'other' fresh range, update the current fresh range
			if tempRange.start <= otherFreshRange.start &&
				tempRange.end >= otherFreshRange.end {
				// Temp range is a superset, we don't need to worry about this other fresh range anymore
				continue
			} else if tempRange.start <= otherFreshRange.start &&
				tempRange.end >= otherFreshRange.start &&
				tempRange.end <= otherFreshRange.end {
				// Temp range end should be extended
				tempRange.end = otherFreshRange.end
				changeCount++
			} else if tempRange.start >= otherFreshRange.start &&
				tempRange.start <= otherFreshRange.end &&
				tempRange.end >= otherFreshRange.end {
				// Temp range start should be extended
				tempRange.start = otherFreshRange.start
				changeCount++
			}
		}

		tempFreshRanges = append(tempFreshRanges, tempRange)
	}

	return tempFreshRanges, changeCount
}
