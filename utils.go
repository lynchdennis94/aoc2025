package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func derivePath(filePrefix string, isTest bool) string {
	path := fmt.Sprintf("puzzles/%s", filePrefix)
	if isTest {
		path += "_test.txt"
	} else {
		path += ".txt"
	}

	return path
}

func loadFileArgs(filePath, splitString string) []string {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(dat), splitString)
}
