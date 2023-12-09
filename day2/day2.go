package day2

import (
	"fmt"
	"main/utils"
	"regexp"
	"strconv"
	"strings"
)

type setOfMarbles struct {
	redMarbles   int
	blueMarbles  int
	greenMarbles int
}

func Run() {

	file := utils.GetData("2", false)
	lines := strings.Split(string(file), "\n")

	defer utils.Timer("Day2")()

	total1, total2 := calculateNumberOfPossibleGames(lines)

	fmt.Println("Sum of possible games:", total1)
	fmt.Println("Power of games:", total2)
}

func calculateNumberOfPossibleGames(games []string) (totalPart1 int, totalPart2 int) {
	max := setOfMarbles{redMarbles: 12, blueMarbles: 14, greenMarbles: 13}

	r := regexp.MustCompile(`Game ([0-9]+): (.*)`)

	for _, game := range games {
		match := r.FindStringSubmatch(game)

		gamenr, _ := strconv.Atoi(match[1])
		sets := strings.Split(match[2], ";")

		errorsInSet := 0
		maxValuesForSet := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, set := range sets {

			m := mapOfSet(set)

			if m["red"] > max.redMarbles || m["green"] > max.greenMarbles || m["blue"] > max.blueMarbles {
				errorsInSet += 1
			}

			var colors = []string{"red", "green", "blue"}
			for _, color := range colors {
				if maxValuesForSet[color] < m[color] {
					maxValuesForSet[color] = m[color]
				}
			}
		}

		if errorsInSet == 0 {
			totalPart1 += gamenr
		}

		totalPart2 += (maxValuesForSet["red"] * maxValuesForSet["green"] * maxValuesForSet["blue"])
	}

	return totalPart1, totalPart2
}

func mapOfSet(set string) (mapOfSet map[string]int) {
	cubesInSet := map[string]int{}

	cubesByColor := strings.Split(set, ",")
	for _, color := range cubesByColor {
		cubeSplit := strings.Split(
			strings.Trim(color, " "), " ")
		value, _ := strconv.Atoi(cubeSplit[0])
		cubesInSet[cubeSplit[1]] = value
	}

	return cubesInSet
}
