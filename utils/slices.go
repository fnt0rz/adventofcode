package utils

import (
	"strconv"
	"strings"
)

func ConvertStringSliceToIntSlice(s []string) (i []int) {

	i = make([]int, len(s))
	for index, v := range s {
		r, err := strconv.Atoi(strings.TrimSpace(v))

		if err != nil {
			panic(err.Error())
		}

		i[index] = r
	}
	return i
}
