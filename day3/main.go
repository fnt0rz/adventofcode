package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type neighbour struct {
	x    int
	y    int
	char rune
}

type symbol struct {
	char       rune
	x          int
	y          int
	neighbours []neighbour
}

func main() {

	file, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")

	parts := map[string]int{}
	gearValues := []int{}
	symbols := indexLines(lines)
	for _, symbol := range symbols {
		symbol.getNeighbours(lines)

		for _, neightbour := range symbol.neighbours {
			if unicode.IsDigit(neightbour.char) {
				key, value := getPartNumber(neightbour, lines)

				if _, exists := parts[key]; !exists {
					parts[key] = value
				}

			}
		}

		gearValue := 0
		if symbol.char == '*' {
			gearNeighboursMap := map[string]int{}
			gearParts := []int{}
			for _, neighbour := range symbol.neighbours {
				if unicode.IsDigit(neighbour.char) {
					key, value := getPartNumber(neighbour, lines)
					if _, exists := gearNeighboursMap[key]; !exists {
						gearNeighboursMap[key] = value
						gearParts = append(gearParts, value)
					}
				}
			}

			if len(gearParts) == 2 {
				gearValue = gearParts[0] * gearParts[1]
			}

			gearValues = append(gearValues, gearValue)
		}
	}

	totalParts := 0
	for _, part := range parts {
		totalParts += part
	}

	totalGears := 0
	for _, gearValue := range gearValues {
		totalGears += gearValue
	}

	fmt.Println(totalParts)
	fmt.Println(totalGears)
}

func getPartNumber(neightbour neighbour, lines []string) (key string, part int) {

	key = fmt.Sprintf("%d,%d", neightbour.x, neightbour.y)
	r := regexp.MustCompile(`[0-9]{1,3}`)
	pos := neightbour.y

	for i := neightbour.y - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(lines[neightbour.x][i])) {
			pos = i
		} else {
			break
		}
	}

	lineSubString := lines[neightbour.x][pos:]

	partAsString := r.FindString(lineSubString)
	if partAsString != "" {
		nr, err := strconv.Atoi(partAsString)

		if err != nil {
			panic(err.Error())
		}

		key = fmt.Sprintf("%d,%d", neightbour.x, pos)
		return key, nr
	}
	return key, 0
}

func (sym *symbol) getNeighbours(lines []string) {
	addNeightbour := func(x, y int) {
		neighbourToAdd := checkNeightbour(x, y, lines)
		sym.neighbours = append(sym.neighbours, neighbour{x: x, y: y, char: neighbourToAdd})
	}

	addNeightbour(sym.x-1, sym.y)
	addNeightbour(sym.x, sym.y-1)
	addNeightbour(sym.x+1, sym.y)
	addNeightbour(sym.x, sym.y+1)
	addNeightbour(sym.x-1, sym.y-1)
	addNeightbour(sym.x+1, sym.y+1)
	addNeightbour(sym.x-1, sym.y+1)
	addNeightbour(sym.x+1, sym.y-1)

}

func checkNeightbour(x, y int, lines []string) (neighbour rune) {
	neighbour = '.'
	withinGrid := (x >= 0 && x < len(lines[0]) && y >= 0 && y < len(lines))
	if withinGrid {
		neighbour = rune(lines[x][y])
	}

	return neighbour

}

func indexLines(lines []string) (symbols []symbol) {

	for x, line := range lines {
		for y, char := range line {
			if char == '.' {
				continue
			}

			if unicode.IsSymbol(char) || unicode.IsPunct(char) {
				symbols = append(symbols, symbol{x: x, y: y, char: char})
			}
		}
	}
	return symbols
}
