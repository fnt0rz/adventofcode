package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)

	var stringData = strings.Split(string(data), "\n")

	r, _ := regexp.Compile(`\d`)
	var result int
	for _, s := range stringData {
		var i, _ = strconv.Atoi(
			findCalibrationValues(
				r.FindAllString(s, 10)))

		result = result + i

	}

	fmt.Fprintln(os.Stdout, []any{"result:", result}...)
}

func findCalibrationValues(i []string) string {
	if len(i) <= 1 {
		return i[0] + i[0]
	}

	first := i[0]
	last := i[len(i)-1]

	return first + last
}
