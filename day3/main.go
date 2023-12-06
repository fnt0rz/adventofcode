package main

import (
	"fmt"
	"os"
	"strings"
)

type part struct {
	startIndex int
	endIndex   int
	lenght     int
}

func main() {

	file, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")

	partsMap := map[int][]part{}
	for index, line := range lines {
		partsInLine := partsPerLine(line)
		partsMap[index] = partsInLine
	}

	fmt.Printf("partsMap: %v\n", partsMap)
}

func partsPerLine(line string) (parts []part) {

	startIndex := 0
	for index, symbol := range line {
		if symbol != '.' && startIndex == 0 {
			startIndex = index
		}

		if symbol == '.' && startIndex != 0 {
			diff := index - startIndex
			parts = append(parts, part{startIndex: startIndex, endIndex: index, lenght: diff})
			startIndex = 0
		}
	}

	return parts
}
