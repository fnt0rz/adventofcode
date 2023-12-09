package day1

import (
	"fmt"
	"main/utils"
	"os"
	"regexp"
	"strings"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func Run() {
	file := utils.GetData("1", false)

	var lines = strings.Split(string(file), "\n")

	defer utils.Timer("Day1")()

	r := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	result := 0
	for _, line := range lines {
		var current []string

		for i := range line {
			found := r.FindString(line[i:])
			if found == "" {
				continue
			}

			current = append(current, found)
		}

		if first, ok := digits[current[0]]; ok {
			result += first * 10
		}

		if last, ok := digits[current[len(current)-1]]; ok {
			result += last
		}

	}

	fmt.Fprintln(os.Stdout, []any{"result:", result}...)
}
