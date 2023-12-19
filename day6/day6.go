package day6

import (
	"fmt"
	"main/utils"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	file := utils.GetData("6", false)
	lines := strings.Split(file, "\n")

	defer utils.Timer("day6")()

	times := getNumbers(lines[0])
	distances := getNumbers(lines[1])
	part1 := solve(times, distances)

	result := part1[0]
	for i := 1; i < len(part1); i++ {
		result = result * part1[i]
	}

	time := getNumber(lines[0])
	distance := getNumber(lines[1])

	part2 := solve(time, distance)

	fmt.Println(result)
	fmt.Println(part2[0])
}

func solve(times, distances []int) []int {
	totalOptions := []int{}

	for i := 0; i < len(times); i++ {
		allowedTime := times[i]
		record := distances[i]
		options := 0

		for j := 0; j < allowedTime; j++ {
			speed := j
			remainingTime := allowedTime - j
			distance := speed * remainingTime

			if distance > record {
				options++
			}
		}

		totalOptions = append(totalOptions, options)
	}
	return totalOptions
}

func getNumber(line string) []int {
	parsedLine := strings.ReplaceAll(line, " ", "")
	r := regexp.MustCompile(`[0-9]+`)
	result, _ := strconv.Atoi(
		r.FindString(parsedLine))

	return []int{result}
}

func getNumbers(line string) []int {
	r := regexp.MustCompile(`\s([0-9]+)`)
	parsedLine := r.FindAllStringSubmatch(line, -1)
	var s []string

	for _, match := range parsedLine {
		s = append(s, match[1])
	}

	return utils.ConvertStringSliceToIntSlice(s)
}
