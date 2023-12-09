package day4

import (
	"fmt"
	"main/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	file := utils.GetData("4", false)
	lines := strings.Split(file, "\n")

	defer utils.Timer("day4")()

	points := 0
	for gamenr, line := range lines {

		r := regexp.MustCompile(`: (.*?) [|] (.*)`)
		result := r.FindAllStringSubmatch(line, -1)

		cleaned := strings.ReplaceAll(strings.TrimSpace(result[0][1]), "  ", " ")
		winningNumbers := strings.Split(cleaned, " ")

		ints := make([]int, len(winningNumbers))
		for i, s := range winningNumbers {
			ints[i], _ = strconv.Atoi(s)
		}

		ownNumbers := strings.Split(result[0][2], " ")

		pointsPergame := 0
		for _, ownNumber := range ownNumbers {
			nr, _ := strconv.Atoi(strings.TrimSpace(ownNumber))
			if slices.Contains(ints, nr) {
				if pointsPergame == 0 {
					pointsPergame = 1
				} else {
					pointsPergame = pointsPergame * 2
				}
			}
		}

		fmt.Printf("Game %d: %d\n", gamenr, pointsPergame)
		points += pointsPergame
	}

	println("TotalPoints part1:", points)
}
